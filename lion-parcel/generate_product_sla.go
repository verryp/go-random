package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	ONEPACK  = `ONEPACK`
	BOSSPACK = `BOSSPACK`

	IsAllowGenerateEstimateSLAProductType = map[string]bool{
		ONEPACK:  true,
		BOSSPACK: true,
	}
)

func main() {
	fmt.Println(generateEstimateSlaByProductType("2-4 Hari", "ONEPACK", time.Now()))

	apiKeyPlain := GenerateApiKey(32)
	encryptionKey := "cee2e0f11be53a9dc1c400f45b9b113d"

	fmt.Println("apiKeyPlain", apiKeyPlain)
	secret, _ := Encrypt(encryptionKey, apiKeyPlain)
	fmt.Println("secret", secret)

	decryptedApiKey, _ := Decrypt(encryptionKey, secret)
	fmt.Println("decryptedApiKey", decryptedApiKey)
}

func generateEstimateSlaByProductType(estimateSLA, productType string, now time.Time) string {

	// Split Estimate SLA
	estimateSlaMin := ""
	estimateSlaMax := ""
	trimEstimateSlString := strings.Replace(estimateSLA, " ", "", -1)
	trimEstimateSlString = strings.TrimSpace(trimEstimateSlString)
	splitEstimateSlaTrim := strings.Split(trimEstimateSlString, "-")
	fmt.Println("trimEstimateSlString", trimEstimateSlString)
	fmt.Println("trimEstimateSlString", trimEstimateSlString)
	fmt.Println("splitEstimateSlaTrim", splitEstimateSlaTrim)
	if len(splitEstimateSlaTrim) > 1 {
		estimateSlaMin = splitEstimateSlaTrim[0]
		estimateSlaMax = strings.Replace(splitEstimateSlaTrim[1], "Hari", "", -1)
	}

	fmt.Println("estimateSlaMin", estimateSlaMin)
	fmt.Println("estimateSlaMax", estimateSlaMax)

	if estimateSlaMin == "" {
		return estimateSLA
	}

	// Check Estimate SLA for ONEPACK
	if IsAllowGenerateEstimateSLAProductType[productType] {
		syyyy, smm, sdd := now.Date()
		customStartTimeFormated := time.Date(syyyy, smm, sdd, 17, 0, 0, 0, now.Location())
		if now.After(customStartTimeFormated) {
			estimateSlaMaxInt, _ := strconv.Atoi(estimateSlaMax)
			estimateSlaMax = fmt.Sprint(estimateSlaMaxInt + 1)
		}
	}

	return fmt.Sprintf("%s - %s Hari", estimateSlaMin, estimateSlaMax)
}
