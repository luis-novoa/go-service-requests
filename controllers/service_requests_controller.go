package controllers

import (
	"fmt"
	"time"
	"math/rand"
	"github.com/graphql-go/graphql"
	// "github.com/luis-novoa/go-service-requests/utils"
	"github.com/luis-novoa/go-service-requests/models"
	"github.com/luis-novoa/go-service-requests/database"
)

func CreateServiceRequest(params graphql.ResolveParams) (interface{}, error) {
	userID, userIDOk := params.Args["input"].(map[string]interface{})["user_id"].(int)
	token, tokenOk := params.Args["input"].(map[string]interface{})["token"].(string)

	if !tokenOk || !userIDOk {
		return nil, fmt.Errorf("Missing token and/or user_id.")
	}

	db := database.Connect()
	defer db.Close()

	var user models.User
	errors := db.First(&user, userID).Error

	if errors != nil {
		return nil, errors
	}

	if user.Technician {
		return nil, fmt.Errorf("Technicians can't create service requests.")
	}

	if user.AuthToken != token {
		return nil, fmt.Errorf("Incorrect token to access this request.")
	}

	var technicians []models.User
	db.Find(&technicians, models.User{ Technician: true })
	chosenTechnician := randomPick(technicians)

	serviceRequest := models.ServiceRequest{ ClientID: userID, TechnicianID: chosenTechnician.ID, Status: "Requested" }
	db.Create(&serviceRequest)
	db.Model(&user).Association("ServiceRequests").Append(&serviceRequest)
	db.Model(&chosenTechnician).Association("ServiceRequests").Append(&serviceRequest)
	return serviceRequest, nil
}

func ShowServiceRequest(params graphql.ResolveParams) (interface{}, error) {
	id, idOk := params.Args["input"].(map[string]interface{})["id"].(int)
	userID, userIDOk := params.Args["input"].(map[string]interface{})["user_id"].(int)
	token, tokenOk := params.Args["input"].(map[string]interface{})["token"].(string)

	if !idOk || !tokenOk || !userIDOk {
		return nil, fmt.Errorf("Missing user_id, token and/or id fields. Provide all the information required to proceed.")
	}

	db := database.Connect()
	defer db.Close()

	var user models.User
	errors := db.Preload("ServiceRequests").First(&user, userID).Error
	if errors != nil {
		return nil, errors
	}

	if user.AuthToken != token {
		return nil, fmt.Errorf("Wrong token for this user.")
	}

	var serviceRequest models.ServiceRequest
	errors = db.Model(&user.ServiceRequests).First(&serviceRequest, id).Error
	if errors != nil {
		return nil, errors
	} else {
		return serviceRequest, nil
	}
}

func IndexServiceRequests(params graphql.ResolveParams) (interface{}, error) {
	userID, userIDOk := params.Args["input"].(map[string]interface{})["user_id"].(int)
	token, tokenOk := params.Args["input"].(map[string]interface{})["token"].(string)

	if !userIDOk || !tokenOk {
		return nil, fmt.Errorf("Missing user_id and/or token fields. Provide all the information required to proceed.")
	}

	db := database.Connect()
	defer db.Close()

	var user models.User
	errors := db.Preload("ServiceRequests").Find(&user, userID).Error
	if errors != nil {
		return nil, errors
	}

	if user.AuthToken != token {
		return nil, fmt.Errorf("Wrong token for this user.")
	}
	
	return user.ServiceRequests, nil
}

func UpdateServiceRequest(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["input"].(map[string]interface{})["id"].(int)
	userID, userIDOk := params.Args["input"].(map[string]interface{})["user_id"].(int)
	solvedRequest, solvedRequestOk := params.Args["input"].(map[string]interface{})["solved_request"].(bool)
	review, reviewOk := params.Args["input"].(map[string]interface{})["review"].(int)
	token, tokenOk := params.Args["input"].(map[string]interface{})["token"].(string)

	if !userIDOk || !tokenOk {
		return nil, fmt.Errorf("Missing user_id, token and/or technician fields. Provide all the information required to proceed.")
	}

	if !solvedRequestOk && !reviewOk {
		return nil, fmt.Errorf("Missing solved_request and review fields. Please provide information about one of them.")
	}

	db := database.Connect()
	defer db.Close()

	var user models.User
	errors := db.Preload("ServiceRequests").Find(&user, userID).Error
	if errors != nil {
		return nil, errors
	}

	if user.AuthToken != token {
		return nil, fmt.Errorf("Wrong token for this user.")
	}

	if user.Technician && reviewOk {
		return nil, fmt.Errorf("Technicians aren't allowed to change the review of the service request.")
	}

	if !user.Technician && solvedRequestOk {
		return nil, fmt.Errorf("Users aren't allowed to change the solved_request field of the service request.")
	}

	if reviewOk && (review > 10 || review < 0) {
		return nil, fmt.Errorf("Your review should be a value between 0 and 10.")
	}

	var serviceRequest models.ServiceRequest
	errors = db.Model(&user.ServiceRequests).First(&serviceRequest, id).Error
	if errors != nil {
		return nil, errors
	} else {
		return serviceRequest, nil
	}

	var status string
	if user.Technician && solvedRequest {
		status = "Waiting for user's review"
	} else {
		if serviceRequest.Status == "Waiting for user's review" {
			status = "Solved"
		} else {
			return nil, fmt.Errorf("This service request isn't waiting for a review.")
		}
	}

	db.Model(&serviceRequest).Updates(models.ServiceRequest{ Status: status, Review: review })
	return serviceRequest, nil
}

// Utilities

func randomPick(users []models.User) models.User {
	rand.Seed(time.Now().UTC().UnixNano())
	randomIndex := rand.Intn(len(users))
	return users[randomIndex]
}