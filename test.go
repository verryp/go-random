package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	//p := "pizza"
	//p2 := "pizza"
	//fmt.Println([]byte(p))
	//fmt.Println([]byte(p2))
	//fmt.Println(compressedString("abaabbbc"))
	fmt.Println(getTopicCount("https://en.wikipedia.org/w/api.php?action=parse&format=json&page=Pizza"))
}

func compressedString(message string) string {
	// Write your code here

	concur := 1

	d := []rune{}
	for i := 0; i < len(message)-1; i++ {
		if i == 5 {
			fmt.Println(i, string(message[i]))
			fmt.Println(i, string(message[i+1]))
			fmt.Println("compare", string(message[i]) == string(message[i+1]))
		}
		if message[i] == message[i+1] {
			//message = message[:i] + message[i+1:]
			concur++
			fmt.Println(i)
			message = fmt.Sprintf("%s%d%s", message[:i+1], message[i+concur:])
			fmt.Println("1", message)
		} else {
			//fmt.Println("out", i)
			message = fmt.Sprintf("%s%d%s", message[:i+1], concur, message[i+concur:])
			fmt.Println("2", message)
			concur = 1
		}
	}

	return message
}

func getTopicCount(topic string) (count int) {
	res, _ := http.Get(topic)
	defer res.Body.Close()

	rawQ := res.Request.URL.RawQuery
	split1 := strings.Split(rawQ, "&")
	split2 := strings.Split(split1[len(split1)-1], "=")[1]

	body, _ := ioutil.ReadAll(res.Body)

	obj := struct {
		Parse struct {
			Title  string `json:"title"`
			Pageid int    `json:"pageid"`
			Text   struct {
				Val string `json:"*"`
			} `json:"text"`
		} `json:"parse"`
	}{}

	json.Unmarshal(body, &obj)

	textB := []byte(obj.Parse.Text.Val)
	topicB := []byte(split2)

	for i, v := range textB {
		if v == topicB[0] {
			fmt.Println(string(textB[i:5]))
			count++
		}
		if count == len(topicB) {
			break
		}
	}
	return
}
