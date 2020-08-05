// DO NOT EDIT -- generated code

// Package agent -
package agent

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/v10/datamodel"
	"github.com/pinpt/go-common/v10/datetime"
	"github.com/pinpt/go-common/v10/hash"
	pjson "github.com/pinpt/go-common/v10/json"
	"github.com/pinpt/go-common/v10/number"
	pstrings "github.com/pinpt/go-common/v10/strings"
)

const (

	// MutationResponseTable is the default table name
	MutationResponseTable datamodel.ModelNameType = "agent_mutationresponse"

	// MutationResponseModelName is the model name
	MutationResponseModelName datamodel.ModelNameType = "agent.MutationResponse"
)

const (
	// MutationResponseModelCustomerIDColumn is the column json value customer_id
	MutationResponseModelCustomerIDColumn = "customer_id"
	// MutationResponseModelErrorColumn is the column json value error
	MutationResponseModelErrorColumn = "error"
	// MutationResponseModelIDColumn is the column json value id
	MutationResponseModelIDColumn = "id"
	// MutationResponseModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	MutationResponseModelIntegrationInstanceIDColumn = "integration_instance_id"
	// MutationResponseModelRefIDColumn is the column json value ref_id
	MutationResponseModelRefIDColumn = "ref_id"
	// MutationResponseModelRefTypeColumn is the column json value ref_type
	MutationResponseModelRefTypeColumn = "ref_type"
	// MutationResponseModelSuccessColumn is the column json value success
	MutationResponseModelSuccessColumn = "success"
)

// MutationResponse request to change data in integration going from agent service to agent
type MutationResponse struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Error the error message if success is false
	Error *string `json:"error,omitempty" codec:"error,omitempty" bson:"error" yaml:"error,omitempty" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// Success if the mutation was successful.
	Success bool `json:"success" codec:"success" bson:"success" yaml:"success" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*MutationResponse)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*MutationResponse)(nil)

func toMutationResponseObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *MutationResponse:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of MutationResponse
func (o *MutationResponse) String() string {
	return fmt.Sprintf("agent.MutationResponse<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *MutationResponse) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *MutationResponse) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *MutationResponse) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *MutationResponse) GetModelName() datamodel.ModelNameType {
	return MutationResponseModelName
}

// NewMutationResponseID provides a template for generating an ID field for MutationResponse
func NewMutationResponseID(customerID string) string {
	return hash.Values(customerID, datetime.EpochNow(), randomString(64))
}

func (o *MutationResponse) setDefaults(frommap bool) {

	if o.ID == "" {
		o.ID = hash.Values(o.CustomerID, datetime.EpochNow(), randomString(64))
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *MutationResponse) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *MutationResponse) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *MutationResponse) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *MutationResponse) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *MutationResponse) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *MutationResponse) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *MutationResponse) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *MutationResponse) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *MutationResponse) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = MutationResponseModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *MutationResponse) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *MutationResponse) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of MutationResponse
func (o *MutationResponse) Clone() datamodel.Model {
	c := new(MutationResponse)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *MutationResponse) Anon() datamodel.Model {
	c := new(MutationResponse)
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
func (o *MutationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *MutationResponse) UnmarshalJSON(data []byte) error {
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
func (o *MutationResponse) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *MutationResponse) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two MutationResponse objects are equal
func (o *MutationResponse) IsEqual(other *MutationResponse) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *MutationResponse) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id":             toMutationResponseObject(o.CustomerID, false),
		"error":                   toMutationResponseObject(o.Error, true),
		"id":                      toMutationResponseObject(o.ID, false),
		"integration_instance_id": toMutationResponseObject(o.IntegrationInstanceID, true),
		"ref_id":                  toMutationResponseObject(o.RefID, false),
		"ref_type":                toMutationResponseObject(o.RefType, false),
		"success":                 toMutationResponseObject(o.Success, false),
		"hashcode":                toMutationResponseObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *MutationResponse) FromMap(kv map[string]interface{}) {

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
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *MutationResponse) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.Error)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.Success)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// MutationResponsePartial is a partial struct for upsert mutations for MutationResponse
type MutationResponsePartial struct {
	// Error the error message if success is false
	Error *string `json:"error,omitempty"`
	// Success if the mutation was successful.
	Success *bool `json:"success,omitempty"`
}

var _ datamodel.PartialModel = (*MutationResponsePartial)(nil)

// GetModelName returns the name of the model
func (o *MutationResponsePartial) GetModelName() datamodel.ModelNameType {
	return MutationResponseModelName
}

// ToMap returns the object as a map
func (o *MutationResponsePartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"error":   toMutationResponseObject(o.Error, true),
		"success": toMutationResponseObject(o.Success, true),
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
func (o *MutationResponsePartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *MutationResponsePartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *MutationResponsePartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *MutationResponsePartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *MutationResponsePartial) FromMap(kv map[string]interface{}) {
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
	o.setDefaults(false)
}