package RecensionService

import (
	"ApiGateway/pkg/models/RecensionService"
	"ApiGateway/pkg/utils"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func AddRecension(w http.ResponseWriter, r *http.Request) {

	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	var chargerDTO RecensionService.RecensionDTO
	data, _ := ioutil.ReadAll(r.Body)
	json.NewDecoder(bytes.NewReader(data)).Decode(&chargerDTO)

	req, _ := http.NewRequest(http.MethodPost, utils.BaseRecensionsServicePath.Next().Host, bytes.NewReader(data))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}
