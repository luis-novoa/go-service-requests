package graphqlconfig

import "github.com/graphql-go/graphql"

clientType := graphql.NewObject(
	graphql.ObjectConfig {
		Name: "Client",
		Fields: graphql.Fields {
			"id": &graphql.Field { Type: graphql.Int },
			"name": &graphql.Field { Type: graphql.String },
			"auth_token": &graphql.Field { Type: graphql.String },
			"created_at": &graphql.Field { Type: graphql.String },
			"updated_at": &graphql.Field { Type: graphql.String }
		}
	}
)

technicianType := graphql.NewObject(
	graphql.ObjectConfig {
		Name: "Technician",
		Fields: graphql.Fields {
			"id": &graphql.Field { Type: graphql.Int },
			"name": &graphql.Field { Type: graphql.String },
			"auth_token": &graphql.Field { Type: graphql.String },
			"created_at": &graphql.Field { Type: graphql.String },
			"updated_at": &graphql.Field { Type: graphql.String }
		}
	}
)

serviceRequestType := graphql.NewObject(
	graphql.ObjectConfig {
		Name: "ServiceRequest"
		Fields: graphql.Fields {
			"id": &graphql.Field { Type: graphql.Int },
			"status": &graphql.Field { Type: graphql.String },
			"review": &graphql.Field { Type: graphql.Int },
			"user_id": &graphql.Field { Type: graphql.Int },
			"technician_id": &graphql.Field { Type: graphql.Int },
			"created_at": &graphql.Field { Type: graphql.String },
			"updated_at": &graphql.Field { Type: graphql.String }
		}
	}
)

userIDTokenInputType := graphql.NewInputObject(
	graphql.InputObjectConfig {
		Name: "UserIDTokenInput",
		Fields: graphql.InputObjectConfigFieldMap {
			"id": &graphql.InputObjectFieldConfig { Type: graphql.Int },
			"token": &graphql.InputObjectFieldConfig { Type: graphql.String }
		}
	}
)

createServiceRequestInputType := graphql.NewInputObject(
	graphql.InputObjectConfig {
		Name: "createServiceRequestInput",
		Fields: graphql.InputObjectConfigFieldMap {
			"client_id": &graphql.InputObjectFieldConfig { Type: graphql.Int },
			"token": &graphql.InputObjectFieldConfig { Type: graphql.String }
		}
	}
)

updateServiceRequestInputType := graphql.NewInputObject(
	graphql.InputObjectConfig {
		Name: "UpdateServiceRequestInput",
		Fields: graphql.InputObjectConfigFieldMap {
			"id": &graphql.InputObjectFieldConfig { Type: graphql.Int },
			"user_id": &graphql.InputObjectFieldConfig { Type: graphql.Int },
			"technician": &graphql.InputObjectFieldConfig { Type: graphql.Bool },
			"solved_request": &graphql.InputObjectFieldConfig { Type: graphql.Bool },
			"review": &graphql.InputObjectFieldConfig { Type: graphql.Int },
			"token": &graphql.InputObjectFieldConfig { Type: graphql.String }
		}
	}
)