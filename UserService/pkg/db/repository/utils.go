package repository

import (
	"errors"
	"net/http"
	"strconv"
)

func BanRecension(recensionId uint) error {

	endpoint := "http://localhost:50005/api/recensions/" + strconv.Itoa(int(recensionId))
	req, _ := http.NewRequest(http.MethodDelete, endpoint, nil)

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)

	if response.StatusCode != http.StatusOK || err != nil {
		return errors.New("recension doesn't exist")
	}

	return nil
}
