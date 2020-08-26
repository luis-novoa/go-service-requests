package controllers

import (
	"fmt"
	"github.com/luis-novoa/utils"
	"github.com/luis-novoa/models"
	"github.com/graphql-go/graphql"
	"github.com/luis-novoa/database"
)

func createTechnician(params graphql.ResolveParams) (interface{}, error) {
	db := database.connect()
	defer db.Close()

	technician := models.Technician{
		Name: params.Args["name"].(string),
		Auth_token: utils.generateToken()
	}

	db.Create(&technician)
	return technician, technician.Error
}

func destroyTechnician(params graphql.ResolveParams) (string, error) {
	id, _ := params.Args["id"].(int)
	token, tokenOk := params.Args["token"].(string)

	if !tokenOk {
		return nil, fmt.Errorf("Please supply a token.")
	}

	db := database.connect()
	defer db.Close()

	var technician models.Technician
	db.First(&technician, id)

	if technician.Error {
		return nil, technician.Error
	}

	if technician.auth_token == token {
		db.Delete(&technician)
		return fmt.Sprintf("%s was succesfully deleted from the database.", technician.Name), nil
	} else {
		return nil, fmt.Errorf("This token doesn't correspond to this technician. Verify if you're providing the right token or technician id.") 
	}
}