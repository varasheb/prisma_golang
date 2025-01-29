package graph

import (
	"context"
	"demo/db"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

type Resolver struct {
	Client *db.PrismaClient
}

func (r *Resolver) Devices(ctx context.Context) ([]map[string]interface{}, error) {
	devices, err := r.Client.Devices.FindMany().Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch devices: %v", err)
	}
	var result []map[string]interface{}
	for _, device := range devices {
		result = append(result, map[string]interface{}{
			"deviceid":   device.Deviceid,
			"devicemeta": device.Devicemeta,
			"createdby":  device.Createdby,
			"createdat":  fmt.Sprintf("%d", device.Createdat),
			"updatedat":  fmt.Sprintf("%d", device.Updatedat),
		})
	}

	return result, nil
}
func (r *Resolver) Groups(ctx context.Context) ([]map[string]interface{}, error) {
	groups, err := r.Client.Groups.FindMany().Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch devices: %v", err)
	}
	var result []map[string]interface{}
	for _, group := range groups {
		value, ok := group.Groupmeta()
		if ok {
			groupMetaStr := string(value)

			result = append(result, map[string]interface{}{
				"orgid":     group.Orgid,
				"groupid":   group.Groupid,
				"groupname": group.Groupname,
				"isdeleted": group.Isdeleted,
				"groupmeta": groupMetaStr,
				"createdat": fmt.Sprintf("%d", group.Createdat),
				"updatedat": fmt.Sprintf("%d", group.Updatedat),
				"createdby": group.Createdby,
				"updatedby": group.Updatedby,
			})
		}
	}

	return result, nil
}

func (r *Resolver) Orgs(ctx context.Context) ([]map[string]interface{}, error) {
	orgs, err := r.Client.Orgs.FindMany().Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch organizations: %v", err)
	}

	var result []map[string]interface{}
	for _, org := range orgs {
		result = append(result, map[string]interface{}{
			"orgid":     org.Orgid,
			"orgname":   org.Orgname,
			"orgmeta":   org.Orgmeta,
			"isenabled": org.Isenabled,
			"createdby": org.Createdby,
			"updatedby": org.Updatedby,
			"createdat": fmt.Sprintf("%d", org.Createdat),
			"updatedat": fmt.Sprintf("%d", org.Updatedat),
		})
	}

	return result, nil
}

func (r *Resolver) GroupDevices(ctx context.Context) ([]map[string]interface{}, error) {
	groupDevices, err := r.Client.Groupdevices.FindMany().Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch group devices: %v", err)
	}

	var result []map[string]interface{}
	for _, groupDevice := range groupDevices {

		result = append(result, map[string]interface{}{
			"orgid":     groupDevice.Orgid,
			"groupid":   groupDevice.Groupid,
			"deviceid":  groupDevice.Deviceid,
			"isexist":   groupDevice.Isexist,
			"updatedat": fmt.Sprintf("%d", groupDevice.Updatedat),
			"updatedby": groupDevice.Updatedby,
		})
	}

	return result, nil
}

func (r *Resolver) GetGroupByGroupID(ctx context.Context, groupID string) ([]map[string]interface{}, error) {
	groups, err := r.Client.Groups.FindMany(
		db.Groups.Groupid.Equals(groupID),
	).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch groups: %v", err)
	}

	var result []map[string]interface{}

	for _, group := range groups {
		value, ok := group.Groupmeta()
		if ok {
			groupMetaStr := string(value)

			result = append(result, map[string]interface{}{
				"orgid":     group.Orgid,
				"groupid":   group.Groupid,
				"groupname": group.Groupname,
				"isdeleted": group.Isdeleted,
				"groupmeta": groupMetaStr,
				"createdat": fmt.Sprintf("%d", group.Createdat),
				"updatedat": fmt.Sprintf("%d", group.Updatedat),
				"createdby": group.Createdby,
				"updatedby": group.Updatedby,
			})
		}
	}

	return result, nil
}
func (r *Resolver) GetGroupByOrgID(ctx context.Context, orgID string) ([]map[string]interface{}, error) {
	groups, err := r.Client.Groups.FindMany(
		db.Groups.Orgid.Equals(orgID),
	).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch groups: %v", err)
	}

	var result []map[string]interface{}
	for _, group := range groups {
		value, ok := group.Groupmeta()
		if ok {
			groupMetaStr := string(value)

			result = append(result, map[string]interface{}{
				"orgid":     group.Orgid,
				"groupid":   group.Groupid,
				"groupname": group.Groupname,
				"isdeleted": group.Isdeleted,
				"groupmeta": groupMetaStr,
				"createdat": fmt.Sprintf("%d", group.Createdat),
				"updatedat": fmt.Sprintf("%d", group.Updatedat),
				"createdby": group.Createdby,
				"updatedby": group.Updatedby,
			})
		}
	}

	return result, nil
}

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"devices": &graphql.Field{
			Type: graphql.NewList(DeviceType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				resolver := p.Context.Value("resolver").(*Resolver)
				return resolver.Devices(p.Context)
			},
		},
		"groups": &graphql.Field{
			Type: graphql.NewList(GroupType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				resolver := p.Context.Value("resolver").(*Resolver)
				return resolver.Groups(p.Context)
			},
		},
		"orgs": &graphql.Field{
			Type: graphql.NewList(OrgType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				resolver := p.Context.Value("resolver").(*Resolver)
				return resolver.Orgs(p.Context)
			},
		},
		"groupdevices": &graphql.Field{
			Type: graphql.NewList(GroupDevicesType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				resolver := p.Context.Value("resolver").(*Resolver)
				return resolver.GroupDevices(p.Context)
			},
		},
		"groupbygroupid": &graphql.Field{
			Type: graphql.NewList(GroupType),
			Args: graphql.FieldConfigArgument{
				"groupID": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				groupID, ok := p.Args["groupID"].(string)
				if !ok {
					return nil, fmt.Errorf("groupID argument is required")
				}

				resolver := p.Context.Value("resolver").(*Resolver)
				return resolver.GetGroupByGroupID(p.Context, groupID)
			},
		},
		"groupbyorgid": &graphql.Field{
			Type: graphql.NewList(GroupType),
			Args: graphql.FieldConfigArgument{
				"orgID": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {

				orgID, ok := p.Args["orgID"].(string)
				if !ok {
					return nil, fmt.Errorf("orgID argument is required")
				}

				resolver := p.Context.Value("resolver").(*Resolver)
				return resolver.GetGroupByOrgID(p.Context, orgID)
			},
		},
	},
})

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createDevice": &graphql.Field{
			Type: DeviceType,
			Args: graphql.FieldConfigArgument{
				"deviceid": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"devicemeta": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(JSONType),
				},
				"createdby": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				deviceid, _ := p.Args["deviceid"].(string)
				devicemeta, _ := p.Args["devicemeta"].(map[string]interface{})
				createdby, _ := p.Args["createdby"].(string)

				metaJSON, err := json.Marshal(devicemeta)
				if err != nil {
					return nil, fmt.Errorf("failed to parse devicemeta: %v", err)
				}
				resolver := p.Context.Value("resolver").(*Resolver)
				return resolver.CreateDevice(p.Context, deviceid, metaJSON, createdby)
			},
		},
	},
})

func (r *Resolver) CreateDevice(ctx context.Context, deviceid string, devicemeta json.RawMessage, createdby string) (map[string]interface{}, error) {
	newDevice, err := r.Client.Devices.CreateOne(
		db.Devices.Deviceid.Set(deviceid),
		db.Devices.Devicemeta.Set(db.JSON(devicemeta)),
		db.Devices.Createdby.Set(createdby),
		db.Devices.Updatedby.Set(createdby),
	).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create device: %v", err)
	}

	return map[string]interface{}{
		"deviceid":   newDevice.Deviceid,
		"devicemeta": newDevice.Devicemeta,
		"createdat":  newDevice.Createdat,
		"createdby":  newDevice.Createdby,
		"updatedat":  newDevice.Updatedat,
		"updatedby":  newDevice.Updatedby,
	}, nil
}

var JSONType = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "JSON",
	Description: "The `JSON` scalar type represents JSON values.",
	Serialize: func(value interface{}) interface{} {
		return value
	},
	ParseValue: func(value interface{}) interface{} {
		return value
	},
	ParseLiteral: func(astValue ast.Value) interface{} {
		switch value := astValue.(type) {
		case *ast.ObjectValue:
			return parseObject(value)
		case *ast.ListValue:
			return parseArray(value)
		case *ast.StringValue:
			return value.Value
		default:
			return nil
		}
	},
})

func parseObject(obj *ast.ObjectValue) map[string]interface{} {
	result := map[string]interface{}{}
	for _, field := range obj.Fields {
		result[field.Name.Value] = field.Value.GetValue()
	}
	return result
}

func parseArray(array *ast.ListValue) []interface{} {
	var result []interface{}
	for _, value := range array.Values {
		result = append(result, value.GetValue())
	}
	return result
}

var BigIntScalar = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "BigInt",
	Description: "BigInt custom scalar",
	Serialize: func(value interface{}) interface{} {
		if v, ok := value.(int64); ok {
			return v
		}
		if v, ok := value.(string); ok {
			return v
		}
		return nil
	},
	ParseValue: func(value interface{}) interface{} {
		if v, ok := value.(int64); ok {
			return v
		}
		if v, ok := value.(string); ok {
			return v
		}
		return nil
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		if v, ok := valueAST.(*ast.IntValue); ok {
			return v.Value
		}
		if v, ok := valueAST.(*ast.StringValue); ok {
			return v.Value
		}
		return nil
	},
})
