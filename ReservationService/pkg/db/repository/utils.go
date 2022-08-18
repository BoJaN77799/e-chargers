package repository

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func VerifyUserUsernameAndVehicle(username string, vehicleId uint) error {

	endpoint := "http://localhost:50001/api/users/exist/" + username + "/" + strconv.Itoa(int(vehicleId))
	resp, err := http.Get(endpoint)

	if resp.StatusCode != http.StatusOK || err != nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatalln(err)
		}

		var errorMessage string
		json.Unmarshal(body, &errorMessage)
		return errors.New(errorMessage)
	}

	return nil
}

func VerifyChargerId(chargerId uint) (uint, error) {
	chargerIdStr := strconv.Itoa(int(chargerId))

	endpoint := "http://localhost:50002/api/chargers/exist/" + chargerIdStr
	resp, err := http.Get(endpoint)

	if resp.StatusCode != http.StatusOK || err != nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatalln(err)
		}

		var errorMessage string
		json.Unmarshal(body, &errorMessage)
		return 0, errors.New(errorMessage)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var capacity uint
	json.Unmarshal(body, &capacity)

	return capacity, nil
}
