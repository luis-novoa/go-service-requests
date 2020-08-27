package controllers

import (
	"fmt"
	"github.com/luis-novoa/utils"
	"github.com/luis-novoa/models"
	"github.com/graphql-go/graphql"
	"github.com/luis-novoa/database"
)

func createServiceRequest(params graphql.ResolveParams) (interface{}, error) {
	clientID, clientIDOk := params.Args["client_id"].(int)
	token, tokenOk := params.Args["token"].(string)

	if !tokenOk || !clientIDOk {
		return nil, fmt.Errorf("Missing token and/or client_id.")
	}

	db := database.connect()
	defer db.Close()

	var client models.Client
	db.First(&client, clientID)

	if client.Error {
		return nil, client.Error
	}

	var technicians []models.Technician
	db.Find(&technicians)
	chosenTechnician = utils.randomPick(technicians)

	serviceRequest := models.ServiceRequest{ ClientID: clientID, TechnicianID: chosenTechnician.ID }
	db.Create(&serviceRequest)
	return serviceRequest, serviceRequest.Error
}

func showServiceRequest(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(int)
	technician, technicianOk := params.Args["technician"].(bool)
	token, tokenOk := params.Args["token"].(string)

	if !technicianOk || !tokenOk {
		return nil, fmt.Errorf("Missing token and/or technician fields. Provide all the information required to proceed.")
	}

	db := database.connect()
	defer db.Close()

	var serviceRequest models.ServiceRequest
	db.Find(&serviceRequest, id)
	if serviceRequest.Error {
		return nil, serviceRequest.Error
	}

	var authToken string
	if technician {
		authToken = db.Model(models.Technician).Related(&serviceRequest).AuthToken
	} else {
		authToken = db.Model(models.Client).Related(&serviceRequest).AuthToken
	}

	if authToken == token {
		return serviceRequest, nil
	} else {
		return nil, fmt.Errorf("Incorrect token to access this request.")
	}
}