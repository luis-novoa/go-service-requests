package graphqlconfig

import (
	"github.com/graphql-go/graphql"
	"github.com/luis-novoa/go-service-requests/controllers"
)

var QueryType = graphql.NewObject(
	graphql.ObjectConfig {
		Name: "Query",
		Fields: graphql.Fields {
			"serviceRequest": &graphql.Field {
				Type: serviceRequestType,
				Description: "Get service request by id",
				Args: graphql.FieldConfigArgument {
					"input": &graphql.ArgumentConfig {
						Type: graphql.NewNonNull(showServiceRequestInputType),
					},
				},
				Resolve: controllers.ShowServiceRequest,
			},
			"serviceRequests": &graphql.Field {
				Type: graphql.NewList(serviceRequestType),
				Description: "Get service request by id",
				Args: graphql.FieldConfigArgument {
					"input": &graphql.ArgumentConfig {
						Type: graphql.NewNonNull(createAndIndexServiceRequestInputType),
					},
				},
				Resolve: controllers.IndexServiceRequests,
			},
		},
	},
)