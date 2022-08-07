package main

import (
	"fmt"
)

func main() {
	hKeys := []string{"a", "add", "dd"}

	var lenKeyCache int
	lenKeyCache = len(hKeys)

	hKeys = hKeys[:lenKeyCache]

	for _, v := range hKeys {
		fmt.Println(v)
	}
}