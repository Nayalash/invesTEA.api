package external

import (
	"io/ioutil"
	"log"
)

func GetKey() string {
	content, err := ioutil.ReadFile("external/key.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	return text
}
