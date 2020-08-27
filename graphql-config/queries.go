package graphqlconfig

import (
	"github.com/graphql-go/graphql"
	"github.com/luis-novoa/go-service-requests/controllers"
)

serviceRequestQueryType := graphql.NewObject(
	graphql.ObjectConfig {
		Name: "ServiceRequestQuery",
		Type: "Query",
		Fields: graphql.Fields {
			"service_request": &graphql.Field {
				Type: serviceRequestType,
				Description: "Get service request by id",
				Args: graphql.FieldConfigArgument {
					"id": &graphql.ArgumentConfig { Type: graphql.Int },
					"technician": &graphql.ArgumentConfig { Type: graphql.Bool },
					"token": &graphql.ArgumentConfig { Type: graphql.String }
				}
			},
			"service_requests": &graphql.Field {
				Type: graphql.NewList(serviceRequestType),
				Description: "Get service request by id",
				Args: graphql.FieldConfigArgument {
					"user_id": &graphql.ArgumentConfig { Type: graphql.Int },
					"technician": &graphql.ArgumentConfig { Type: graphql.Bool },
					"token": &graphql.ArgumentConfig { Type: graphql.String }
				}
			}
		}
	}
)