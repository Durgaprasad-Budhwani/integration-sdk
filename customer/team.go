// DO NOT EDIT -- generated code

// Package customer - customer specific data
package customer

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	"github.com/pinpt/go-common/slice"
	pstrings "github.com/pinpt/go-common/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// TeamModelName is the model name
	TeamModelName datamodel.ModelNameType = "customer.Team"
)

// Team a team is a grouping of one or more users
type Team struct {
	// Active whether the team is tracked in pinpoint
	Active *bool `json:"active,omitempty" codec:"active,omitempty" bson:"active" yaml:"active,omitempty" faker:"-"`
	// ChildrenIds the children_ids for this team
	ChildrenIds []string `json:"children_ids" codec:"children_ids" bson:"children_ids" yaml:"children_ids" faker:"-"`
	// CreatedAt the date the record was created in Epoch time
	CreatedAt int64 `json:"created_ts" codec:"created_ts" bson:"created_ts" yaml:"created_ts" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Description the description of the team
	Description string `json:"description" codec:"description" bson:"description" yaml:"description" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Leaf True when team has no children_ids
	Leaf bool `json:"leaf" codec:"leaf" bson:"leaf" yaml:"leaf" faker:"-"`
	// Name the name of the team
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"team"`
	// ParentIds the parent_ids for this team
	ParentIds []string `json:"parent_ids" codec:"parent_ids" bson:"parent_ids" yaml:"parent_ids" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// TeamCostID the team cost id
	TeamCostID *string `json:"team_cost_id,omitempty" codec:"team_cost_id,omitempty" bson:"team_cost_id" yaml:"team_cost_id,omitempty" faker:"-"`
	// UpdatedAt the date the record was updated in Epoch time
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Team)(nil)

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

// GetModelName returns the name of the model
func (o *Team) GetModelName() datamodel.ModelNameType {
	return TeamModelName
}

// NewTeamID provides a template for generating an ID field for Team
func NewTeamID(customerID string) string {
	return hash.Values(customerID, randomString(64))
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
		o.ID = hash.Values(o.CustomerID, randomString(64))
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

// GetRefID returns the RefID for the object
func (o *Team) GetRefID() string {
	return o.RefID
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
