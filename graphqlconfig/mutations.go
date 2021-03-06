package graphqlconfig

import (
	"github.com/graphql-go/graphql"
	"github.com/luis-novoa/go-service-requests/controllers"
)

var MutationType = graphql.NewObject(
	graphql.ObjectConfig {
		Name: "Mutation",
		Fields: graphql.Fields {
			"createUser": &graphql.Field {
				Type: userType,
				Description: "Create new user",
				Args: graphql.FieldConfigArgument {
					"input": &graphql.ArgumentConfig {
						Type: graphql.NewNonNull(createUserInputType),
					},
				},
				Resolve: controllers.CreateUser,
			},
			"destroyUser": &graphql.Field {
				Type: graphql.String,
				Description: "Destroy user by ID",
				Args: graphql.FieldConfigArgument {
					"input": &graphql.ArgumentConfig {
						Type: graphql.NewNonNull(destroyUserInputType),
					},
				},
				Resolve: controllers.DestroyUser,
			},
			"createServiceRequest": &graphql.Field {
				Type: serviceRequestType,
				Description: "Create new service request",
				Args: graphql.FieldConfigArgument {
					"input": &graphql.ArgumentConfig {
						Type: graphql.NewNonNull(createAndIndexServiceRequestInputType),
					},
				},
				Resolve: controllers.CreateServiceRequest,
			},
			"updateServiceRequest": &graphql.Field {
				Type: serviceRequestType,
				Description: "Update characteristics of existing service request",
				Args: graphql.FieldConfigArgument {
					"input": &graphql.ArgumentConfig {
						Type: graphql.NewNonNull(updateServiceRequestInputType),
					},
				},
				Resolve: controllers.UpdateServiceRequest,
			},
		},
	},
)