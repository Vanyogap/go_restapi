package go_pack

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// )

type Data struct {
	Data *AccountData `json:"data,omitempty"`
}
type AccountData struct {
	ID             string             `json:"id,omitempty"`
	OrganisationID string             `json:"organisation_id,omitempty"`
	Type           string             `json:"type,omitempty"`
	Version        *int64             `json:"version,omitempty"`
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
}

type AccountAttributes struct {
	AccountClassification   *string  `json:"account_classification,omitempty"`
	AccountMatchingOptOut   *bool    `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	AlternativeNames        []string `json:"alternative_names,omitempty"`
	BankID                  string   `json:"bank_id,omitempty"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Country                 *string  `json:"country,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	JointAccount            *bool    `json:"joint_account,omitempty"`
	Name                    []string `json:"name,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Status                  *string  `json:"status,omitempty"`
	Switched                *bool    `json:"switched,omitempty"`
}

// func fetch(accountId string) {

// 	var url bytes.Buffer
// 	url.WriteString("http://localhost:8080/v1/organisation/accounts/")
// 	url.WriteString(accountId)

// 	resp, err := http.Get(url.String())
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	sb := string(body)
// 	log.Printf(sb)

// }

// func fetchAll() {

// 	resp, err := http.Get("http://localhost:8080/v1/organisation/accounts/")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	sb := string(body)
// 	log.Printf(sb)

// }

// func create(values string) {

// 	var acc Data

// 	json.Unmarshal([]byte(values), &acc)

// 	json_data, err := json.Marshal(acc)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// fmt.Println(string(json_data))
// 	fmt.Printf("%s", json_data)

// 	resp, err := http.Post("http://localhost:8080/v1/organisation/accounts/", "application/json", bytes.NewBuffer(json_data))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer resp.Body.Close()
// 	bodyBytes, _ := ioutil.ReadAll(resp.Body)

// 	// Convert response body to string
// 	bodyString := string(bodyBytes)
// 	fmt.Println(bodyString)

// }
// func delete(accountId string) {

// 	var url bytes.Buffer
// 	url.WriteString("http://localhost:8080/v1/organisation/accounts/")
// 	url.WriteString(accountId)
// 	url.WriteString("?version=0")
// 	req, err := http.NewRequest(http.MethodDelete, url.String(), nil)
// 	client := &http.Client{}
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	defer resp.Body.Close()
// 	bodyBytes, _ := ioutil.ReadAll(resp.Body)

// 	// Convert response body to string
// 	bodyString := string(bodyBytes)
// 	fmt.Println(bodyString)

// }
