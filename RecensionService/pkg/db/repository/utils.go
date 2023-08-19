package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"log"
	"net/http"
)

const UserServiceURL = "http://localhost:50001/api/usr"

const ChargerServiceURL = "http://localhost:50002/api"

func VerifyUserUsername(userId uuid.UUID) error {
	URL := UserServiceURL + "/users/exist/" + userId.String()
	resp, err := http.Get(URL)
	if resp.StatusCode != http.StatusOK || err != nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		var errorMessage string
		json.Unmarshal(body, &errorMessage)

		log.Fatalln(errorMessage)
		return errors.New(errorMessage)
	}
	return nil
}

func VerifyCharger(chargerId uuid.UUID) error {
	URL := ChargerServiceURL + "/chargers/exist/" + chargerId.String()
	resp, err := http.Get(URL)
	if resp.StatusCode != http.StatusOK || err != nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatalln(err)
		}

		var errorMessage string
		json.Unmarshal(body, &errorMessage)

		log.Fatalln(errorMessage)
		return errors.New(errorMessage)
	}
	return nil
}

func GetRecensionToxicity(recensionText string) ([]float32, error) {

	var result []float32

	req, _ := http.NewRequest(http.MethodPost, "http://localhost:50006/prediction", bytes.NewReader([]byte(recensionText)))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)

	if response.StatusCode != http.StatusOK || err != nil {
		return result, errors.New("strikes AI not work")
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(body, &result)

	return result, nil
}
