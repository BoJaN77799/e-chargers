package repository

import (
	"encoding/json"
	"errors"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"log"
	"net/http"
	"reservation_service/pkg/entities"
)

const UserServiceUrl = "http://localhost:50001/api/usr"
const ChargerServiceUrl = "http://localhost:50002/api/chr"

func VerifyUserUsernameAndVehicle(userId uuid.UUID, vehicleId uuid.UUID) (entities.UserReservationDTO, error) {

	URL := UserServiceUrl + "/users/exist/" + userId.String() + "/" + vehicleId.String()
	response, err := http.Get(URL)

	var user entities.UserReservationDTO

	if response.StatusCode != http.StatusOK || err != nil {
		return user, errors.New("user doesn't exist or user isn't owner of vehicle")
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(body, &user)

	return user, nil
}

func GetChargerById(id uuid.UUID) (entities.ChargerReservationDTO, error) {
	URL := ChargerServiceUrl + "/chargers/reservation/" + id.String()
	resp, err := http.Get(URL)

	var charger entities.ChargerReservationDTO

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
