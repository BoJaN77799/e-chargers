package UserService

import (
	"ApiGateway/pkg/handlers"
	"ApiGateway/pkg/models/RecensionService"
	"ApiGateway/pkg/utils"
	"bytes"
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

type UserBaseInfoDTO struct {
	Id        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
}

type UsersBatchDTO struct {
	UserIds []uuid.UUID `json:"userIds"`
}

func GetUsersBatch(recensions []RecensionService.RecensionDTO) ([]RecensionService.RecensionWithUserDTO, error) {
	userIds := getUniqueUserIdsFromRecensions(recensions)
	response, err := getUsersByUserIds(userIds)
	if err != nil {
		return []RecensionService.RecensionWithUserDTO{}, err
	}

	defer response.Body.Close()

	var users []UserBaseInfoDTO
	if err := json.NewDecoder(response.Body).Decode(&users); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return []RecensionService.RecensionWithUserDTO{}, err
	}

	return mergeRecensionsWithUsers(recensions, users), nil
}

func getUniqueUserIdsFromRecensions(recensions []RecensionService.RecensionDTO) []uuid.UUID {
	tempMap := make(map[uuid.UUID]bool)

	for _, recension := range recensions {
		tempMap[recension.UserId] = true
	}

	var userIds []uuid.UUID
	for userId := range tempMap {
		userIds = append(userIds, userId)
	}
	return userIds
}

func getUsersByUserIds(userIds []uuid.UUID) (*http.Response, error) {
	var usersBatch = UsersBatchDTO{
		UserIds: userIds,
	}
	requestBody, _ := json.Marshal(usersBatch)
	URL := utils.BaseUserServicePath.Next().Host + "/users/batch"
	response, err := handlers.DoRequest(http.MethodPost, URL, bytes.NewBuffer(requestBody))

	return response, err
}

func mergeRecensionsWithUsers(recensions []RecensionService.RecensionDTO, users []UserBaseInfoDTO) []RecensionService.RecensionWithUserDTO {
	userMap := make(map[uuid.UUID]UserBaseInfoDTO)
	for _, user := range users {
		userMap[user.Id] = user
	}

	var recensionsWithUser []RecensionService.RecensionWithUserDTO
	for _, recension := range recensions {
		user := userMap[recension.UserId]
		merged := RecensionService.RecensionWithUserDTO{
			RecensionDTO: recension,
			Email:        user.Email,
			Firstname:    user.Firstname,
			Lastname:     user.Lastname,
		}
		recensionsWithUser = append(recensionsWithUser, merged)
	}
	return recensionsWithUser
}
