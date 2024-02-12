package main

import "fmt"

func main() {
	lang := map[string]string{"en": "english", "fr": "french", "es": "spanish", "de": "german", "it": "italian"}
	//unordered collection of key/value pairs
	lang["jp"] = "japanese"
	for k, v := range lang {
		fmt.Println(k, v)
	}

}
