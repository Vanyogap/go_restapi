package go_pack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func fetch(accountId string) {

	var url bytes.Buffer
	url.WriteString("http://localhost:8080/v1/organisation/accounts/")
	url.WriteString(accountId)

	resp, err := http.Get(url.String())
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Printf(sb)

}

func fetchAll() {

	resp, err := http.Get("http://localhost:8080/v1/organisation/accounts/")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Printf(sb)

}

func create(values string) {

	var acc Data

	json.Unmarshal([]byte(values), &acc)

	json_data, err := json.Marshal(acc)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(string(json_data))
	fmt.Printf("%s", json_data)

	resp, err := http.Post("http://localhost:8080/v1/organisation/accounts/", "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

}
func delete(accountId string) {

	var url bytes.Buffer
	url.WriteString("http://localhost:8080/v1/organisation/accounts/")
	url.WriteString(accountId)
	url.WriteString("?version=0")
	req, err := http.NewRequest(http.MethodDelete, url.String(), nil)
	client := &http.Client{}
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

}
