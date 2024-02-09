package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type WebhookResponse struct {
	//
}

type Response struct {
	Revision           string        `json:"revision"`
	BundleId           string        `json:"bundleId"`
	Environment        string        `json:"environment"`
	HasMore            bool          `json:"hasMore"`
	SignedTransactions []interface{} `json:"signedTransactions"`
}

// type SignedTransactionsPayload struct {
// 	AppAccountToken             string `json:"appAccountToken"`
// 	BundleId                    string `json:"bundleId"`
// 	Environment                 string `json:"environment"`
// 	ExpiresDate                 int64  `json:"expiresDate"`
// 	InAppOwnershipType          string `json:"inAppOwnershipType"`
// 	IsUpgraded                  bool   `json:"isUpgraded"`
// 	OfferIdentifier             string `json:"offerIdentifier"`
// 	OfferType                   string `json:"offerType"`
// 	OriginalPurchaseDate        int64  `json:"originalPurchaseDate"`
// 	OriginalTransactionId       string `json:"originalTransactionId"`
// 	ProductId                   string `json:"productId"`
// 	PurchaseDate                int64  `json:"purchaseDate"`
// 	Quantity                    int    `json:"quantity"`
// 	RevocationDate              int64  `json:"revocationDate"`
// 	RevocationReason            string `json:"revocationReason"`
// 	SignedDate                  int64  `json:"signedDate"`
// 	SubscriptionGroupIdentifier string `json:"subscriptionGroupIdentifier"`
// 	TransactionId               string `json:"transactionId"`
// 	Type                        string `json:"type"`
// 	WebOrderLineItemId          string `json:"webOrderLineItemId"`
// }

type ResponseSubs struct {
	Environment string        `json:"environment"`
	BundleId    string        `json:"BundleId"`
	Data        []SubsPayload `json:"data"`
}

type SubsPayload struct {
	SubscriptionGroupIdentifier string                    `json:"subscriptionGroupIdentifier"`
	LastTransactions            []LastTransactionsPayload `json:"lastTransactions"`
}

type LastTransactionsPayload struct {
	OriginalTransactionId string      `json:"originalTransactionId"`
	Status                int         `json:"status"`
	SignedTransactionInfo interface{} `json:"signedTransactionInfo"`
	SignedRenewalInfo     interface{} `json:"signedRenewalInfo"`
}

type CallbackPayload struct {
	SignedPayload interface{} `json:"signedPayload"`
}

func getCallback(w http.ResponseWriter, r *http.Request) {
	var test CallbackPayload

	json.NewDecoder(r.Body).Decode(&test)

	token := strings.Split(test.SignedPayload.(string), ".")
	// decodedPayload, err := base64.URLEncoding.DecodeString(token[1])
	decodedPayload, _ := jwt.DecodeSegment(token[1])

	var resp interface{}
	json.Unmarshal(decodedPayload, &resp)

	test.SignedPayload = resp
	b, _ := json.Marshal(test)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func getTxHistory(w http.ResponseWriter, r *http.Request) {
	var test Response

	json.NewDecoder(r.Body).Decode(&test)

	var responses []interface{}
	for i, t := range test.SignedTransactions {
		token := strings.Split(t.(string), ".")
		// decodedPayload, err := base64.URLEncoding.DecodeString(token[1])
		decodedPayload, err := jwt.DecodeSegment(token[1])
		if err != nil {
			fmt.Println("errr", i, err.Error())
		}

		var resp interface{}
		json.Unmarshal(decodedPayload, &resp)

		responses = append(responses, resp)
	}

	test.SignedTransactions = responses

	b, _ := json.Marshal(test)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func getSubsStatuses(w http.ResponseWriter, r *http.Request) {
	var test ResponseSubs

	json.NewDecoder(r.Body).Decode(&test)

	var responses []SubsPayload
	for _, data := range test.Data {
		var lastTx []LastTransactionsPayload
		for _, tx := range data.LastTransactions {
			tokenSub := strings.Split(tx.SignedTransactionInfo.(string), ".")
			tokenRenew := strings.Split(tx.SignedTransactionInfo.(string), ".")
			// decodedPayload, err := base64.URLEncoding.DecodeString(token[1])
			dTokenSub, _ := jwt.DecodeSegment(tokenSub[1])
			dTokenRenew, _ := jwt.DecodeSegment(tokenRenew[1])

			var respSub interface{}
			json.Unmarshal(dTokenSub, &respSub)
			var respRenew interface{}
			json.Unmarshal(dTokenRenew, &respRenew)

			tx.SignedTransactionInfo = respSub
			tx.SignedRenewalInfo = respRenew

			lastTx = append(lastTx, tx)
		}

		data.LastTransactions = lastTx
		responses = append(responses, data)
	}

	test.Data = responses
	b, _ := json.Marshal(test)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/callback", getCallback)

	mux.HandleFunc("/get-subscription-statuses", getSubsStatuses)

	mux.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		token := `eyJ0cmFuc2FjdGlvbklkIjoiMjAwMDAwMDE5NzQxMTcwNiIsIm9yaWdpbmFsVHJhbnNhY3Rpb25JZCI6IjIwMDAwMDAxOTc0MTE3MDYiLCJidW5kbGVJZCI6ImNvbS5yYWxhbGkuYXBwLmlvcy5kZWJ1ZyIsInByb2R1Y3RJZCI6Im5vbnJlbmV3aW5nNDUiLCJwdXJjaGFzZURhdGUiOjE2Njc5Nzk5NTgwMDAsIm9yaWdpbmFsUHVyY2hhc2VEYXRlIjoxNjY3OTc5OTU4MDAwLCJxdWFudGl0eSI6MSwidHlwZSI6Ik5vbi1SZW5ld2luZyBTdWJzY3JpcHRpb24iLCJpbkFwcE93bmVyc2hpcFR5cGUiOiJQVVJDSEFTRUQiLCJzaWduZWREYXRlIjoxNjY3OTgwNjA3MzYwLCJlbnZpcm9ubWVudCI6IlNhbmRib3gifQ`
		// token = strings.Replace(token, " ", "", -1)
		decodedPayload, err := jwt.DecodeSegment(token)
		if err != nil {
			fmt.Println("errr", err.Error())
		}

		var resp interface{}
		json.Unmarshal(decodedPayload, &resp)

		b, _ := json.Marshal(resp)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(b)
	})

	mux.HandleFunc("/get-transaction-history", getTxHistory)

	err := http.ListenAndServe(":3333", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
