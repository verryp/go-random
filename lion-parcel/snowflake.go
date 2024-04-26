package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	hn := GetHostname()
	fmt.Println(hn)
	fmt.Println(GetNodeID(hn))
}

// GetHostname  get machine hostname
func GetHostname() string {
	nodeName, err := os.Hostname()
	if err != nil {
		return ""
	}

	return nodeName
}

func GetNodeID(nodeName string) int {
	fmt.Println("ddddd", 1<<22)
	fmt.Println("d", time.Now().UnixNano()/(1<<22))
	nodeDefault := fmt.Sprint(time.Now().UnixNano() / (1 << 22))
	x := strings.Split(nodeName, ".")
	rgx, err := regexp.Compile(`[^0-9]+`)

	fmt.Println("nodedefault", nodeDefault)
	fmt.Println("nodedefault", SubString(nodeDefault, (len(nodeDefault)-3), 3))

	if err != nil || len(x) == 0 {
		return ParseToInt(SubString(nodeDefault, (len(nodeDefault) - 3), 3))
	}

	n := rgx.ReplaceAllString(x[0], "")

	if n == "" {
		n = fmt.Sprint(time.Now().UnixNano() / (1 << 22))
	}

	if len(n) <= 3 {
		return ParseToInt(n)
	}

	fmt.Println("n", n)
	return ParseToInt(SubString(n, (len(n) - 3), 3))
}

func ParseToInt(val string) int {
	i, _ := strconv.ParseInt(val, 0, 0)
	return int(i)
}

func SubString(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}
