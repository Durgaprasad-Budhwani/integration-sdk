// DO NOT EDIT -- generated code

// Package customer - customer specific data
package customer

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	"github.com/pinpt/go-common/slice"
	pstrings "github.com/pinpt/go-common/strings"
)

const (
	// TeamTopic is the default topic name
	TeamTopic datamodel.TopicNameType = "customer_Team_topic"

	// TeamTable is the default table name
	TeamTable datamodel.ModelNameType = "customer_team"

	// TeamModelName is the model name
	TeamModelName datamodel.ModelNameType = "customer.Team"
)

const (
	// TeamActiveColumn is the active column name
	TeamActiveColumn = "Active"
	// TeamChildrenIdsColumn is the children_ids column name
	TeamChildrenIdsColumn = "ChildrenIds"
	// TeamCreatedAtColumn is the created_ts column name
	TeamCreatedAtColumn = "CreatedAt"
	// TeamCustomerIDColumn is the customer_id column name
	TeamCustomerIDColumn = "CustomerID"
	// TeamDescriptionColumn is the description column name
	TeamDescriptionColumn = "Description"
	// TeamIDColumn is the id column name
	TeamIDColumn = "ID"
	// TeamLeafColumn is the leaf column name
	TeamLeafColumn = "Leaf"
	// TeamNameColumn is the name column name
	TeamNameColumn = "Name"
	// TeamParentIdsColumn is the parent_ids column name
	TeamParentIdsColumn = "ParentIds"
	// TeamRefIDColumn is the ref_id column name
	TeamRefIDColumn = "RefID"
	// TeamRefTypeColumn is the ref_type column name
	TeamRefTypeColumn = "RefType"
	// TeamTeamCostIDColumn is the team_cost_id column name
	TeamTeamCostIDColumn = "TeamCostID"
	// TeamUpdatedAtColumn is the updated_ts column name
	TeamUpdatedAtColumn = "UpdatedAt"
)

// Team a team is a grouping of one or more users
type Team struct {
	// Active whether the team is tracked in pinpoint
	Active *bool `json:"active,omitempty" yaml:"active,omitempty" faker:"-"`
	// ChildrenIds the children_ids for this team
	ChildrenIds []string `json:"children_ids" yaml:"children_ids" faker:"-"`
	// CreatedAt the date the record was created in Epoch time
	CreatedAt int64 `json:"created_ts" yaml:"created_ts" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" yaml:"customer_id" faker:"-"`
	// Description the description of the team
	Description string `json:"description" yaml:"description" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// Leaf True when team has no children_ids
	Leaf bool `json:"leaf" yaml:"leaf" faker:"-"`
	// Name the name of the team
	Name string `json:"name" yaml:"name" faker:"team"`
	// ParentIds the parent_ids for this team
	ParentIds []string `json:"parent_ids" yaml:"parent_ids" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" yaml:"ref_type" faker:"-"`
	// TeamCostID the team cost id
	TeamCostID *string `json:"team_cost_id,omitempty" yaml:"team_cost_id,omitempty" faker:"-"`
	// UpdatedAt the date the record was updated in Epoch time
	UpdatedAt int64 `json:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Team)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Team)(nil)

func toTeamObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Team:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of Team
func (o *Team) String() string {
	return fmt.Sprintf("customer.Team<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Team) GetTopicName() datamodel.TopicNameType {
	return TeamTopic
}

// GetStreamName returns the name of the stream
func (o *Team) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Team) GetTableName() string {
	return TeamTable.String()
}

// GetModelName returns the name of the model
func (o *Team) GetModelName() datamodel.ModelNameType {
	return TeamModelName
}

// NewTeamID provides a template for generating an ID field for Team
func NewTeamID(Name string, customerID string) string {
	return hash.Values(Name, customerID)
}

func (o *Team) setDefaults(frommap bool) {
	if o.Active == nil {
		o.Active = &emptyBool
	}
	if o.ChildrenIds == nil {
		o.ChildrenIds = make([]string, 0)
	}
	if o.ParentIds == nil {
		o.ParentIds = make([]string, 0)
	}
	if o.TeamCostID == nil {
		o.TeamCostID = &emptyString
	}

	if o.ID == "" {
		// set the id from the spec provided in the model
		o.ID = hash.Values(o.Name, o.CustomerID)
	}

	if o.Active == nil {
		v := false
		o.Active = &v
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Team) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Team) GetTopicKey() string {
	var i interface{} = o.CustomerID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Team) GetTimestamp() time.Time {
	var dt interface{} = o.UpdatedAt
	switch v := dt.(type) {
	case int64:
		return datetime.DateFromEpoch(v).UTC()
	case string:
		tv, err := datetime.ISODateToTime(v)
		if err != nil {
			panic(err)
		}
		return tv.UTC()
	case time.Time:
		return v.UTC()
	}
	panic("not sure how to handle the date time format for Team")
}

// GetRefID returns the RefID for the object
func (o *Team) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Team) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Team) IsMutable() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Team) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Team) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Team) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = TeamModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Team) GetTopicConfig() *datamodel.ModelTopicConfig {
	retention, err := time.ParseDuration("87360h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 87360h0m0s. " + err.Error())
	}

	ttl, err := time.ParseDuration("0s")
	if err != nil {
		ttl = 0
	}
	if ttl == 0 && retention != 0 {
		ttl = retention // they should be the same if not set
	}
	return &datamodel.ModelTopicConfig{
		Key:               "customer_id",
		Timestamp:         "updated_ts",
		NumPartitions:     8,
		CleanupPolicy:     datamodel.CleanupPolicy("compact"),
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetCustomerID will return the customer_id
func (o *Team) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of Team
func (o *Team) Clone() datamodel.Model {
	c := new(Team)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Team) Anon() datamodel.Model {
	c := new(Team)
	if err := faker.FakeData(c); err != nil {
		panic("couldn't create anon version of object: " + err.Error())
	}
	kv := c.ToMap()
	for k, v := range o.ToMap() {
		if _, ok := kv[k]; !ok {
			kv[k] = v
		}
	}
	c.FromMap(kv)
	return c
}

// MarshalJSON returns the bytes for marshaling to json
func (o *Team) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Team) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	if idstr, ok := kv["id"].(string); ok {
		o.ID = idstr
	}
	return nil
}

// Stringify returns the object in JSON format as a string
func (o *Team) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Team objects are equal
func (o *Team) IsEqual(other *Team) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Team) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"active":       toTeamObject(o.Active, true),
		"children_ids": toTeamObject(o.ChildrenIds, false),
		"created_ts":   toTeamObject(o.CreatedAt, false),
		"customer_id":  toTeamObject(o.CustomerID, false),
		"description":  toTeamObject(o.Description, false),
		"id":           toTeamObject(o.ID, false),
		"leaf":         toTeamObject(o.Leaf, false),
		"name":         toTeamObject(o.Name, false),
		"parent_ids":   toTeamObject(o.ParentIds, false),
		"ref_id":       toTeamObject(o.RefID, false),
		"ref_type":     toTeamObject(o.RefType, false),
		"team_cost_id": toTeamObject(o.TeamCostID, true),
		"updated_ts":   toTeamObject(o.UpdatedAt, false),
		"hashcode":     toTeamObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Team) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["active"].(*bool); ok {
		o.Active = val
	} else if val, ok := kv["active"].(bool); ok {
		o.Active = &val
	} else {
		if val, ok := kv["active"]; ok {
			if val == nil {
				o.Active = number.BoolPointer(number.ToBoolAny(nil))
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Active = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}

	if val, ok := kv["children_ids"]; ok {
		if val != nil {
			na := make([]string, 0)
			if a, ok := val.([]string); ok {
				na = append(na, a...)
			} else {
				if a, ok := val.([]interface{}); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							if badMap, ok := ae.(map[interface{}]interface{}); ok {
								ae = slice.ConvertToStringToInterface(badMap)
							}
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for children_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else if s, ok := val.(string); ok {
					for _, sv := range strings.Split(s, ",") {
						na = append(na, strings.TrimSpace(sv))
					}
				} else if a, ok := val.(primitive.A); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for children_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for children_ids field")
				}
			}
			o.ChildrenIds = na
		}
	}
	if o.ChildrenIds == nil {
		o.ChildrenIds = make([]string, 0)
	}

	if val, ok := kv["created_ts"].(int64); ok {
		o.CreatedAt = val
	} else {
		if val, ok := kv["created_ts"]; ok {
			if val == nil {
				o.CreatedAt = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.CreatedAt = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["customer_id"].(string); ok {
		o.CustomerID = val
	} else {
		if val, ok := kv["customer_id"]; ok {
			if val == nil {
				o.CustomerID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.CustomerID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["description"].(string); ok {
		o.Description = val
	} else {
		if val, ok := kv["description"]; ok {
			if val == nil {
				o.Description = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Description = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["id"].(string); ok {
		o.ID = val
	} else {
		if val, ok := kv["id"]; ok {
			if val == nil {
				o.ID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.ID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["leaf"].(bool); ok {
		o.Leaf = val
	} else {
		if val, ok := kv["leaf"]; ok {
			if val == nil {
				o.Leaf = number.ToBoolAny(nil)
			} else {
				o.Leaf = number.ToBoolAny(val)
			}
		}
	}

	if val, ok := kv["name"].(string); ok {
		o.Name = val
	} else {
		if val, ok := kv["name"]; ok {
			if val == nil {
				o.Name = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Name = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["parent_ids"]; ok {
		if val != nil {
			na := make([]string, 0)
			if a, ok := val.([]string); ok {
				na = append(na, a...)
			} else {
				if a, ok := val.([]interface{}); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							if badMap, ok := ae.(map[interface{}]interface{}); ok {
								ae = slice.ConvertToStringToInterface(badMap)
							}
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for parent_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else if s, ok := val.(string); ok {
					for _, sv := range strings.Split(s, ",") {
						na = append(na, strings.TrimSpace(sv))
					}
				} else if a, ok := val.(primitive.A); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for parent_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for parent_ids field")
				}
			}
			o.ParentIds = na
		}
	}
	if o.ParentIds == nil {
		o.ParentIds = make([]string, 0)
	}

	if val, ok := kv["ref_id"].(string); ok {
		o.RefID = val
	} else {
		if val, ok := kv["ref_id"]; ok {
			if val == nil {
				o.RefID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.RefID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["ref_type"].(string); ok {
		o.RefType = val
	} else {
		if val, ok := kv["ref_type"]; ok {
			if val == nil {
				o.RefType = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.RefType = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["team_cost_id"].(*string); ok {
		o.TeamCostID = val
	} else if val, ok := kv["team_cost_id"].(string); ok {
		o.TeamCostID = &val
	} else {
		if val, ok := kv["team_cost_id"]; ok {
			if val == nil {
				o.TeamCostID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.TeamCostID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["updated_ts"].(int64); ok {
		o.UpdatedAt = val
	} else {
		if val, ok := kv["updated_ts"]; ok {
			if val == nil {
				o.UpdatedAt = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.UpdatedAt = number.ToInt64Any(val)
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *Team) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Active)
	args = append(args, o.ChildrenIds)
	args = append(args, o.CreatedAt)
	args = append(args, o.CustomerID)
	args = append(args, o.Description)
	args = append(args, o.ID)
	args = append(args, o.Leaf)
	args = append(args, o.Name)
	args = append(args, o.ParentIds)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.TeamCostID)
	args = append(args, o.UpdatedAt)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *Team) GetEventAPIConfig() datamodel.EventAPIConfig {
	return datamodel.EventAPIConfig{
		Publish: datamodel.EventAPIPublish{
			Public: false,
		},
		Subscribe: datamodel.EventAPISubscribe{
			Public: false,
			Key:    "",
		},
	}
}
