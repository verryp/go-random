package main

import (
	"fmt"
	"encoding/base64"
)
func main() {
	// create function from string to base64
	cacert := `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvKNTHBcZgR1K10ESr+iQ
7y/aQdxZR9dfW0l25q8yr73anWAt3TI4Mk5d9ZVlICcfmUT1SqpDsxMDqzcXpds/
XLPP5Odj7ZMkbIeyUXhkTf3NSonZWozP+zzMMznpzZIyCK3oSGOR10ovJYaJW7uP
0/bnSV5RizF7rWl8qtaStZGKCYW+RwtE4UPreErxNy4hQSxiPuS1nUbUZdtt9mq3
HO0fnTbBqbRKs+MlHYcsV8Uw7cytOz2bj/4FTwggV8TVJt0tTq7QefvBlkkstvc5
iABDjQVulFkv/wgimgkOFn8wIiWJzGJ9w45eDbw0q4ipIcTEKk8WhbKBR0CKCNgU
kwIDAQAB
-----END PUBLIC KEY-----`
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(cacert)))
}
