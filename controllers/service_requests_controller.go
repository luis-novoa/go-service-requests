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