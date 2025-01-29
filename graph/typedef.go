package graph

import "github.com/graphql-go/graphql"

var DeviceType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Device",
	Fields: graphql.Fields{
		"deviceid":   &graphql.Field{Type: graphql.String},
		"devicemeta": &graphql.Field{Type: JSONType},
		"createdby":  &graphql.Field{Type: graphql.String},
		"createdat":  &graphql.Field{Type: BigIntScalar},
		"updatedat":  &graphql.Field{Type: BigIntScalar},
	},
})

// Group type
var GroupType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Group",
	Fields: graphql.Fields{
		"groupid":   &graphql.Field{Type: graphql.String},
		"groupname": &graphql.Field{Type: graphql.String},
		"isdeleted": &graphql.Field{Type: graphql.Boolean},
		"groupmeta": &graphql.Field{Type: JSONType},
		"createdby": &graphql.Field{Type: graphql.String},
		"createdat": &graphql.Field{Type: BigIntScalar},
		"updatedat": &graphql.Field{Type: BigIntScalar},
	},
})

// Org type
var OrgType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Org",
	Fields: graphql.Fields{
		"orgid":     &graphql.Field{Type: graphql.String},
		"orgname":   &graphql.Field{Type: graphql.String},
		"orgmeta":   &graphql.Field{Type: JSONType},
		"isenabled": &graphql.Field{Type: graphql.Boolean},
		"createdby": &graphql.Field{Type: graphql.String},
		"createdat": &graphql.Field{Type: BigIntScalar},
		"updatedat": &graphql.Field{Type: BigIntScalar},
	},
})

var GroupDevicesType = graphql.NewObject(graphql.ObjectConfig{
	Name: "GroupDevices",
	Fields: graphql.Fields{
		"orgid":     &graphql.Field{Type: graphql.String},
		"groupid":   &graphql.Field{Type: graphql.String},
		"deviceid":  &graphql.Field{Type: graphql.String},
		"isexist":   &graphql.Field{Type: graphql.Boolean},
		"updatedat": &graphql.Field{Type: BigIntScalar},
		"updatedby": &graphql.Field{Type: graphql.String},
	},
})
