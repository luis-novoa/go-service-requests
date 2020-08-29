package controllers

import (
	"fmt"
	"time"
	"math/rand"
	"github.com/graphql-go/graphql"
	"github.com/luis-novoa/go-service-requests/models"
	"github.com/luis-novoa/go-service-requests/database"
)

func CreateServiceRequest(params graphql.ResolveParams) (interface{}, error) {
	userID := params.Args["input"].(map[string]interface{})["user_id"].(int)
	token := params.Args["input"].(map[string]interface{})["token"].(string)

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
	return serviceRequest, nil
}

func ShowServiceRequest(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["input"].(map[string]interface{})["id"].(int)
	userID := params.Args["input"].(map[string]interface{})["user_id"].(int)
	token := params.Args["input"].(map[string]interface{})["token"].(string)

	db := database.Connect()
	defer db.Close()

	var user models.User
	errors := db.First(&user, userID).Error
	if errors != nil {
		return nil, errors
	}

	if user.AuthToken != token {
		return nil, fmt.Errorf("Wrong token for this user.")
	}

	var serviceRequest models.ServiceRequest
	if user.Technician {
		errors = db.Where(&models.ServiceRequest{ TechnicianID: user.ID }).First(&serviceRequest, id).Error
	} else {
		errors = db.Where(&models.ServiceRequest{ ClientID: user.ID }).First(&serviceRequest, id).Error
	}
	if errors != nil {
		return nil, errors
	} else {
		return serviceRequest, nil
	}
}

func IndexServiceRequests(params graphql.ResolveParams) (interface{}, error) {
	userID := params.Args["input"].(map[string]interface{})["user_id"].(int)
	token := params.Args["input"].(map[string]interface{})["token"].(string)

	db := database.Connect()
	defer db.Close()

	var user models.User
	errors := db.First(&user, userID).Error
	if errors != nil {
		return nil, errors
	}

	if user.AuthToken != token {
		return nil, fmt.Errorf("Wrong token for this user.")
	}

	var serviceRequests []models.ServiceRequest
	if user.Technician {
		errors = db.Where(&models.ServiceRequest{ TechnicianID: user.ID }).Find(&serviceRequests).Error
	} else {
		errors = db.Where(&models.ServiceRequest{ ClientID: user.ID }).Find(&serviceRequests).Error
	}
	
	if errors != nil {
		return nil, errors
	} else {
		return serviceRequests, nil
	}
}

func UpdateServiceRequest(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["input"].(map[string]interface{})["id"].(int)
	userID := params.Args["input"].(map[string]interface{})["user_id"].(int)
	solvedRequest, solvedRequestOk := params.Args["input"].(map[string]interface{})["solved_request"].(bool)
	review, reviewOk := params.Args["input"].(map[string]interface{})["review"].(int)
	token := params.Args["input"].(map[string]interface{})["token"].(string)


	if solvedRequestOk && reviewOk {
		return nil, fmt.Errorf("It is not allowed to change status and review at the same time.")
	}

	db := database.Connect()
	defer db.Close()

	var user models.User
	errors := db.First(&user, userID).Error
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
	if user.Technician {
		errors = db.Where(&models.ServiceRequest{ TechnicianID: user.ID }).First(&serviceRequest, id).Error
	} else {
		errors = db.Where(&models.ServiceRequest{ ClientID: user.ID }).First(&serviceRequest, id).Error
	}
	if errors != nil {
		return nil, errors
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