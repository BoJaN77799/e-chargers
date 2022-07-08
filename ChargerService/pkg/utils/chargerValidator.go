package utils

import (
	"charger_service/pkg/models"
	"errors"
	"fmt"
)

const MaxChargerCapacity = 10

func CheckChargersInfo(charger *models.Charger) error {

	err := CheckAddressInfo(charger.Address)

	if err != nil {
		return err
	}

	if len(charger.Name) == 0 {
		return errors.New("charger name is empty")
	}

	if len(charger.WorkTime) == 0 {
		return errors.New("charger name is empty")
	}

	if len(charger.Description) == 0 {
		charger.Description = "Empty description."
	}

	if charger.Capacity < 0 || charger.Capacity > MaxChargerCapacity {
		charger.Description = fmt.Sprintf("charger capacity is out of bounds (0, %d)", MaxChargerCapacity)
	}

	if len(charger.Plugs) == 0 {
		return errors.New("charger is without plugs")
	}

	err = CheckPlugsInfo(charger.Plugs)

	if err != nil {
		return err
	}

	return nil
}

func CheckAddressInfo(address models.Address) error {

	if len(address.Street) == 0 {
		return errors.New("address street is empty")
	}

	if len(address.City) == 0 {
		return errors.New("address city is empty")
	}

	if len(address.Country) == 0 {
		return errors.New("address country is empty")
	}

	if address.PostalCode < 500 || address.PostalCode > 99950 {
		return errors.New(fmt.Sprintf("address postal code is out of bounds (%d, %d)", 500, 99950))
	}

	if address.Latitude < -85 || address.Latitude > 85 {
		return errors.New(fmt.Sprintf("address latitude is out of bounds (%d, %d)", -85, 85))
	}

	if address.Longitude < -180 || address.Longitude > 180 {
		return errors.New(fmt.Sprintf("address longitude is out of bounds (%d, %d)", -180, 180))
	}

	return nil
}

func CheckPlugsInfo(plugs []models.Plug) error {

	for _, plug := range plugs {
		err := CheckPlugInfo(plug)
		if err != nil {
			return err
		}
	}
	return nil
}

func CheckPlugInfo(plug models.Plug) error {

	if len(plug.PricePerHour) == 0 {
		return errors.New("charger price is empty")
	}

	if len(plug.Type) == 0 {
		return errors.New("plug type is empty")
	}

	if len(plug.ChargingSpeedPerMinute) == 0 {
		return errors.New("plug charging speed per minute is empty")
	}

	return nil
}
