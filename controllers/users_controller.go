package controllers

import (
	"fmt"
	"crypto/rand"
	"crypto/sha256"
	"github.com/graphql-go/graphql"
	"github.com/luis-novoa/go-service-requests/models"
	"github.com/luis-novoa/go-service-requests/database"
)

func CreateUser(params graphql.ResolveParams) (interface{}, error) {
	db := database.Connect()
	defer db.Close()

	user := models.User {
		Name: params.Args["input"].(map[string]interface{})["name"].(string),
		Technician: params.Args["input"].(map[string]interface{})["technician"].(bool),
		AuthToken: generateToken(),
	}

	errors := db.Save(&user).Error
	if errors != nil {
		return nil, errors
	}
	return user, nil
}

func DestroyUser(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["input"].(map[string]interface{})["id"].(int)
	token := params.Args["input"].(map[string]interface{})["token"].(string)

	db := database.Connect()
	defer db.Close()

	var user models.User
	errors := db.First(&user, id).Error

	if errors != nil {
		return nil, errors
	}

	if user.AuthToken == token {
		db.Delete(&user)
		return fmt.Sprintf("%s was succesfully deleted from the database.", user.Name), nil
	} else {
		return nil, fmt.Errorf("This token doesn't correspond to this user. Verify if you're providing the right token or user id.") 
	}
}

// Utilities

func generateToken() string {
	randomByte := make([]byte, 10)
	rand.Read(randomByte)
	return fmt.Sprintf("%x", sha256.Sum256(randomByte))
}