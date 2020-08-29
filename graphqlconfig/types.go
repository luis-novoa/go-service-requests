package graphqlconfig

import "github.com/graphql-go/graphql"

var userType = graphql.NewObject(
	graphql.ObjectConfig {
		Name: "User",
		Fields: graphql.Fields {
			"id": &graphql.Field { Type: graphql.Int },
			"name": &graphql.Field { Type: graphql.String },
			"technician": &graphql.Field { Type: graphql.Boolean },
			"auth_token": &graphql.Field { Type: graphql.String },
			"created_at": &graphql.Field { Type: graphql.String },
			"updated_at": &graphql.Field { Type: graphql.String },
		},
	},
)

var serviceRequestType = graphql.NewObject(
	graphql.ObjectConfig {
		Name: "ServiceRequest",
		Fields: graphql.Fields {
			"id": &graphql.Field { Type: graphql.Int },
			"status": &graphql.Field { Type: graphql.String },
			"review": &graphql.Field { Type: graphql.Int },
			"client_id": &graphql.Field { Type: graphql.Int },
			"technician_id": &graphql.Field { Type: graphql.Int },
			"created_at": &graphql.Field { Type: graphql.String },
			"updated_at": &graphql.Field { Type: graphql.String },
		},
	},
)

var createUserInputType = graphql.NewInputObject(
	graphql.InputObjectConfig {
		Name: "createUserInput",
		Fields: graphql.InputObjectConfigFieldMap {
			"name": &graphql.InputObjectFieldConfig { Type: graphql.NewNonNull(graphql.String) },
			"technician": &graphql.InputObjectFieldConfig { Type: graphql.NewNonNull(graphql.Boolean) },
		},
	},
)

var destroyUserInputType = graphql.NewInputObject(
	graphql.InputObjectConfig {
		Name: "destroyUserInput",
		Fields: graphql.InputObjectConfigFieldMap {
			"id": &graphql.InputObjectFieldConfig { Type: graphql.NewNonNull(graphql.Int) },
			"token": &graphql.InputObjectFieldConfig { Type: graphql.NewNonNull(graphql.String) },
		},
	},
)

var createAndIndexServiceRequestInputType = graphql.NewInputObject(
	graphql.InputObjectConfig {
		Name: "createAndIndexServiceRequestInput",
		Fields: graphql.InputObjectConfigFieldMap {
			"user_id": &graphql.InputObjectFieldConfig { Type: graphql.NewNonNull(graphql.Int) },
			"token": &graphql.InputObjectFieldConfig { Type: graphql.NewNonNull(graphql.String) },
		},
	},
)

var showServiceRequestInputType = graphql.NewInputObject(
	graphql.InputObjectConfig {
		Name: "showServiceRequestInput",
		Fields: graphql.InputObjectConfigFieldMap {
			"id": &graphql.InputObjectFieldConfig { Type: graphql.NewNonNull(graphql.Int) },
			"user_id": &graphql.InputObjectFieldConfig { Type: graphql.NewNonNull(graphql.Int) },
			"token": &graphql.InputObjectFieldConfig { Type: graphql.NewNonNull(graphql.String) },
		},
	},
)

var updateServiceRequestInputType = graphql.NewInputObject(
	graphql.InputObjectConfig {
		Name: "UpdateServiceRequestInput",
		Fields: graphql.InputObjectConfigFieldMap {
			"id": &graphql.InputObjectFieldConfig { Type: graphql.NewNonNull(graphql.Int) },
			"user_id": &graphql.InputObjectFieldConfig { Type: graphql.NewNonNull(graphql.Int) },
			"solved_request": &graphql.InputObjectFieldConfig { Type: graphql.Boolean },
			"review": &graphql.InputObjectFieldConfig { Type: graphql.Int },
			"token": &graphql.InputObjectFieldConfig { Type: graphql.NewNonNull(graphql.String) },
		},
	},
)