package helper

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const url = "https://pokeapi.co/api/v2/"

func MakeRequest() {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
		// return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(body)
}
