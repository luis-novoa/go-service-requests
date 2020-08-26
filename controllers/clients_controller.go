package controllers

import (
	"fmt"
	"github.com/luis-novoa/utils"
	"github.com/luis-novoa/models"
	"github.com/graphql-go/graphql"
	"github.com/luis-novoa/database"
)

func createClient(params graphql.ResolveParams) (interface{}, error) {
	db := database.connect()
	defer db.Close()

	client := models.Client{
		Name: params.Args["name"].(string),
		Auth_token: utils.generateToken()
	}

	db.Create(&client)
	return client, client.Error
}

func destroyClient(params graphql.ResolveParams) (string, error) {
	id, _ := params.Args["id"].(int)
	token, tokenOk := params.Args["token"].(string)

	if !tokenOk {
		return nil, fmt.Errorf("Please supply a token.")
	}

	db := database.connect()
	defer db.Close()

	var client models.Client
	db.First(&client, id)

	if client.Error {
		return nil, client.Error
	}

	if client.auth_token == token {
		db.Delete(&client)
		return fmt.Sprintf("%s was succesfully deleted from the database.", client.Name), nil
	} else {
		return nil, fmt.Errorf("This token doesn't correspond to this client. Verify if you're providing the right token or client id.") 
	}
}