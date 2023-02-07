package apihelper

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tejas-cogo/go-cogoport/models"
)

func RubyClient(body models.PartnerUserBody, rubyclient models.RubyClientInput) ([]byte, error) {
	hc := http.Client{}

	reqBody, err := json.Marshal(body)

	url := "https://api-apollo3.dev.cogoport.io/" + rubyclient.Endpoint
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(reqBody))
	if err != nil {
		var b []byte
		log.Printf("Request Failed: %s", err)
		return b, err
	}

	req.Header.Add("Authorization", "Bearer: 787b8f21-ca0a-4e79-af6e-81e3ca847909")
	req.Header.Add("AuthorizationScope", "partner")

	req.Header.Add("Content-type", "application/json; charset=utf-8")

	resp, err := hc.Do(req)

	if err == nil {

		defer resp.Body.Close()
	}

	resp_body, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(resp_body, "resp_body")
	// Log the request body
	// bodyString := string(resp_body)
	// log.Print(bodyString)
	// Unmarshal result
	// var post models.RubyClientOutput
	// post := models.RubyClientOutput{}
	// err = json.Unmarshal(resp_body, &post)
	// fmt.Println("---------post--------", post.List, "--------post----------")
	// if err != nil {
	// 	log.Printf("Reading body failed: %s", err)
	// 	return post, err
	// }
	return resp_body, err
}
