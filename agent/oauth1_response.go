// DO NOT EDIT -- generated code

// Package agent - the agent communicates with pinpoint cloud to send integration data
package agent

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/v10/datamodel"
	"github.com/pinpt/go-common/v10/hash"
	pjson "github.com/pinpt/go-common/v10/json"
	"github.com/pinpt/go-common/v10/number"
	pstrings "github.com/pinpt/go-common/v10/strings"
)

const (

	// Oauth1ResponseTable is the default table name
	Oauth1ResponseTable datamodel.ModelNameType = "agent_oauth1response"

	// Oauth1ResponseModelName is the model name
	Oauth1ResponseModelName datamodel.ModelNameType = "agent.Oauth1Response"
)

const (
	// Oauth1ResponseModelCustomerIDColumn is the column json value customer_id
	Oauth1ResponseModelCustomerIDColumn = "customer_id"
	// Oauth1ResponseModelErrorColumn is the column json value error
	Oauth1ResponseModelErrorColumn = "error"
	// Oauth1ResponseModelIDColumn is the column json value id
	Oauth1ResponseModelIDColumn = "id"
	// Oauth1ResponseModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	Oauth1ResponseModelIntegrationInstanceIDColumn = "integration_instance_id"
	// Oauth1ResponseModelRefIDColumn is the column json value ref_id
	Oauth1ResponseModelRefIDColumn = "ref_id"
	// Oauth1ResponseModelRefTypeColumn is the column json value ref_type
	Oauth1ResponseModelRefTypeColumn = "ref_type"
	// Oauth1ResponseModelSecretColumn is the column json value secret
	Oauth1ResponseModelSecretColumn = "secret"
	// Oauth1ResponseModelSessionIDColumn is the column json value session_id
	Oauth1ResponseModelSessionIDColumn = "session_id"
	// Oauth1ResponseModelSuccessColumn is the column json value success
	Oauth1ResponseModelSuccessColumn = "success"
	// Oauth1ResponseModelTokenColumn is the column json value token
	Oauth1ResponseModelTokenColumn = "token"
)

// Oauth1Response An OAuth1 response sent from the agent with the temporary credentials
type Oauth1Response struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Error if the request was not successful, the description of the error
	Error *string `json:"error,omitempty" codec:"error,omitempty" bson:"error" yaml:"error,omitempty" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// Secret the temporary OAuth secret if successful
	Secret *string `json:"secret,omitempty" codec:"secret,omitempty" bson:"secret" yaml:"secret,omitempty" faker:"-"`
	// SessionID A unique session ID generated by the client which the response will include to capture it more easily
	SessionID string `json:"session_id" codec:"session_id" bson:"session_id" yaml:"session_id" faker:"-"`
	// Success if the request was successful or not
	Success bool `json:"success" codec:"success" bson:"success" yaml:"success" faker:"-"`
	// Token the temporary OAuth token if successful
	Token *string `json:"token,omitempty" codec:"token,omitempty" bson:"token" yaml:"token,omitempty" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Oauth1Response)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Oauth1Response)(nil)

func toOauth1ResponseObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Oauth1Response:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of Oauth1Response
func (o *Oauth1Response) String() string {
	return fmt.Sprintf("agent.Oauth1Response<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Oauth1Response) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *Oauth1Response) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Oauth1Response) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *Oauth1Response) GetModelName() datamodel.ModelNameType {
	return Oauth1ResponseModelName
}

// NewOauth1ResponseID provides a template for generating an ID field for Oauth1Response
func NewOauth1ResponseID(customerID string, refType string, refID string) string {
	return hash.Values("Oauth1Response", customerID, refType, refID)
}

func (o *Oauth1Response) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Oauth1Response", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Oauth1Response) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Oauth1Response) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Oauth1Response) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *Oauth1Response) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Oauth1Response) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Oauth1Response) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Oauth1Response) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Oauth1Response) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *Oauth1Response) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = Oauth1ResponseModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Oauth1Response) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *Oauth1Response) GetCustomerID() string {

	return o.CustomerID

}

// SetCustomerID will return the customer_id
func (o *Oauth1Response) SetCustomerID(id string) {

	o.CustomerID = id

}

// GetRefType will return the ref_type
func (o *Oauth1Response) GetRefType() string {
	return o.RefType
}

// SetRefType will return the ref_type
func (o *Oauth1Response) SetRefType(t string) {
	o.RefType = t
}

// Clone returns an exact copy of Oauth1Response
func (o *Oauth1Response) Clone() datamodel.Model {
	c := new(Oauth1Response)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Oauth1Response) Anon() datamodel.Model {
	c := new(Oauth1Response)
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
func (o *Oauth1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Oauth1Response) UnmarshalJSON(data []byte) error {
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
func (o *Oauth1Response) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *Oauth1Response) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two Oauth1Response objects are equal
func (o *Oauth1Response) IsEqual(other *Oauth1Response) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Oauth1Response) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id":             toOauth1ResponseObject(o.CustomerID, false),
		"error":                   toOauth1ResponseObject(o.Error, true),
		"id":                      toOauth1ResponseObject(o.ID, false),
		"integration_instance_id": toOauth1ResponseObject(o.IntegrationInstanceID, true),
		"ref_id":                  toOauth1ResponseObject(o.RefID, false),
		"ref_type":                toOauth1ResponseObject(o.RefType, false),
		"secret":                  toOauth1ResponseObject(o.Secret, true),
		"session_id":              toOauth1ResponseObject(o.SessionID, false),
		"success":                 toOauth1ResponseObject(o.Success, false),
		"token":                   toOauth1ResponseObject(o.Token, true),
		"hashcode":                toOauth1ResponseObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Oauth1Response) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["customer_id"].(string); ok {
		o.CustomerID = val
	} else {
		if val, ok := kv["customer_id"]; ok {
			if val == nil {
				o.CustomerID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.CustomerID = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["error"].(*string); ok {
		o.Error = val
	} else if val, ok := kv["error"].(string); ok {
		o.Error = &val
	} else {
		if val, ok := kv["error"]; ok {
			if val == nil {
				o.Error = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Error = pstrings.Pointer(fmt.Sprintf("%v", val))
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.ID = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["integration_instance_id"].(*string); ok {
		o.IntegrationInstanceID = val
	} else if val, ok := kv["integration_instance_id"].(string); ok {
		o.IntegrationInstanceID = &val
	} else {
		if val, ok := kv["integration_instance_id"]; ok {
			if val == nil {
				o.IntegrationInstanceID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.IntegrationInstanceID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["ref_id"].(string); ok {
		o.RefID = val
	} else {
		if val, ok := kv["ref_id"]; ok {
			if val == nil {
				o.RefID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.RefType = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["secret"].(*string); ok {
		o.Secret = val
	} else if val, ok := kv["secret"].(string); ok {
		o.Secret = &val
	} else {
		if val, ok := kv["secret"]; ok {
			if val == nil {
				o.Secret = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Secret = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["session_id"].(string); ok {
		o.SessionID = val
	} else {
		if val, ok := kv["session_id"]; ok {
			if val == nil {
				o.SessionID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.SessionID = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["success"].(bool); ok {
		o.Success = val
	} else {
		if val, ok := kv["success"]; ok {
			if val == nil {
				o.Success = false
			} else {
				o.Success = number.ToBoolAny(val)
			}
		}
	}
	if val, ok := kv["token"].(*string); ok {
		o.Token = val
	} else if val, ok := kv["token"].(string); ok {
		o.Token = &val
	} else {
		if val, ok := kv["token"]; ok {
			if val == nil {
				o.Token = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Token = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *Oauth1Response) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.Error)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.Secret)
	args = append(args, o.SessionID)
	args = append(args, o.Success)
	args = append(args, o.Token)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *Oauth1Response) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *Oauth1Response) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// Oauth1ResponsePartial is a partial struct for upsert mutations for Oauth1Response
type Oauth1ResponsePartial struct {
	// Error if the request was not successful, the description of the error
	Error *string `json:"error,omitempty"`
	// Secret the temporary OAuth secret if successful
	Secret *string `json:"secret,omitempty"`
	// SessionID A unique session ID generated by the client which the response will include to capture it more easily
	SessionID *string `json:"session_id,omitempty"`
	// Success if the request was successful or not
	Success *bool `json:"success,omitempty"`
	// Token the temporary OAuth token if successful
	Token *string `json:"token,omitempty"`
}

var _ datamodel.PartialModel = (*Oauth1ResponsePartial)(nil)

// GetModelName returns the name of the model
func (o *Oauth1ResponsePartial) GetModelName() datamodel.ModelNameType {
	return Oauth1ResponseModelName
}

// ToMap returns the object as a map
func (o *Oauth1ResponsePartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"error":      toOauth1ResponseObject(o.Error, true),
		"secret":     toOauth1ResponseObject(o.Secret, true),
		"session_id": toOauth1ResponseObject(o.SessionID, true),
		"success":    toOauth1ResponseObject(o.Success, true),
		"token":      toOauth1ResponseObject(o.Token, true),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {
		}
	}
	return kv
}

// Stringify returns the object in JSON format as a string
func (o *Oauth1ResponsePartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *Oauth1ResponsePartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Oauth1ResponsePartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *Oauth1ResponsePartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *Oauth1ResponsePartial) FromMap(kv map[string]interface{}) {
	if val, ok := kv["error"].(*string); ok {
		o.Error = val
	} else if val, ok := kv["error"].(string); ok {
		o.Error = &val
	} else {
		if val, ok := kv["error"]; ok {
			if val == nil {
				o.Error = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Error = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["secret"].(*string); ok {
		o.Secret = val
	} else if val, ok := kv["secret"].(string); ok {
		o.Secret = &val
	} else {
		if val, ok := kv["secret"]; ok {
			if val == nil {
				o.Secret = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Secret = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["session_id"].(*string); ok {
		o.SessionID = val
	} else if val, ok := kv["session_id"].(string); ok {
		o.SessionID = &val
	} else {
		if val, ok := kv["session_id"]; ok {
			if val == nil {
				o.SessionID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.SessionID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["success"].(*bool); ok {
		o.Success = val
	} else if val, ok := kv["success"].(bool); ok {
		o.Success = &val
	} else {
		if val, ok := kv["success"]; ok {
			if val == nil {
				o.Success = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Success = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	if val, ok := kv["token"].(*string); ok {
		o.Token = val
	} else if val, ok := kv["token"].(string); ok {
		o.Token = &val
	} else {
		if val, ok := kv["token"]; ok {
			if val == nil {
				o.Token = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Token = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	o.setDefaults(false)
}
