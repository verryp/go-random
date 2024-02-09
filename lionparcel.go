package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	sttManualCreateBulk()
}

func sttManualCreateBulk() {
	var listStt []string
	totalStt := 100
	//tn := time.Now()
	batch := 50
	for i := 1; i <= totalStt; i++ {
		//newTime := tn.Add(time.Millisecond * time.Duration(i))
		stt := RandStringRunes(5)
		listStt = append(listStt, stt)
	}

	//fmt.Println("listStt", listStt)

	for i := 0; i < len(listStt); i += batch {
		j := i + batch
		if j > len(listStt) {
			j = len(listStt)
		}

		//fmt.Println("ke-", i)
		fmt.Println("listStt[i:j]", listStt[i:j])
		//err = c.sttManualRepo.CreateBulk(selfCtx, &model.SttManualBulk{
		//	SttNo:              listStt[i:j],
		//	AccountRefID:       form.AccountRefID,
		//	AccountType:        form.AccountType,
		//	SttManualStatus:    model.SttManualUnused,
		//	SttManualMixpack:   form.IsMixpack,
		//	AccountRefName:     accountRefDetail.Name,
		//	AccountRefCode:     accountRefDetail.Code,
		//	AccountRefCityName: accountRefCityName,
		//})
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
