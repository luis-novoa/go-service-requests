package graphqlconfig

import (
	"github.com/graphql-go/graphql"
	"github.com/luis-novoa/go-service-requests/controllers"
)

mutationType := graphql.NewObject(
	graphql.ObjectConfig {
		Name: "Mutation",
		Type: "Mutation",
		Fields: graphql.Fields {
			"createClient": &graphql.Field {
				Type: clientType,
				Description: "Create new client",
				Args: graphql.FieldConfigArgument {
					"name": &graphql.ArgumentConfig {
						Type: graphql.NewNonNull(graphql.String)
					}
				},
				Resolve: controllers.createClient()
			},
			"destroyClient": &graphql.Field {
				Type: graphql.String,
				Description: "Create new client",
				Args: graphql.FieldConfigArgument {
					"input": &graphql.ArgumentConfig {
						Type: graphql.NewNonNull(destroyUserInputType)
					}
				},
				Resolve: controllers.destroyClient()
			},
			"createTechnician": &graphql.Field {
				Type: technicianType,
				Description: "Create new technician",
				Args: graphql.FieldConfigArgument {
					"name": &graphql.ArgumentConfig {
						Type: graphql.NewNonNull(graphql.String)
					}
				},
				Resolve: controllers.createTechnician()
			},
			"destroyTechnician": &graphql.Field {
				Type: graphql.String,
				Description: "Create new technician",
				Args: graphql.FieldConfigArgument {
					"input": &graphql.ArgumentConfig {
						Type: graphql.NewNonNull(destroyUserInputType)
					}
				},
				Resolve: controllers.destroyTechnician()
			},
		}
	}
)