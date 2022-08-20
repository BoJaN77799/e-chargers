package repository

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"reservation_service/pkg/models"
	"strconv"
)

func VerifyUserUsernameAndVehicle(username string, vehicleId uint) (models.UserReservationDTO, error) {

	endpoint := "http://localhost:50001/api/users/exist/" + username + "/" + strconv.Itoa(int(vehicleId))
	resp, err := http.Get(endpoint)

	var user models.UserReservationDTO

	if resp.StatusCode != http.StatusOK || err != nil {
		return user, errors.New("user doesn't exist or isn't owner of vehicle")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(body, &user)

	return user, nil
}

func VerifyChargerId(chargerId uint) (models.ChargerReservationDTO, error) {
	chargerIdStr := strconv.Itoa(int(chargerId))

	endpoint := "http://localhost:50002/api/chargers/exist/" + chargerIdStr
	resp, err := http.Get(endpoint)

	var charger models.ChargerReservationDTO

	if resp.StatusCode != http.StatusOK || err != nil {

		return charger, errors.New("charger doesn't exist")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(body, &charger)

	return charger, nil
}
