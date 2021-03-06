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
	pstrings "github.com/pinpt/go-common/v10/strings"
)

const (

	// CodequalityTriggerTable is the default table name
	CodequalityTriggerTable datamodel.ModelNameType = "agent_codequalitytrigger"

	// CodequalityTriggerModelName is the model name
	CodequalityTriggerModelName datamodel.ModelNameType = "agent.CodequalityTrigger"
)

const (
	// CodequalityTriggerModelCustomerIDColumn is the column json value customer_id
	CodequalityTriggerModelCustomerIDColumn = "customer_id"
	// CodequalityTriggerModelIDColumn is the column json value id
	CodequalityTriggerModelIDColumn = "id"
	// CodequalityTriggerModelIntegrationIDColumn is the column json value integration_id
	CodequalityTriggerModelIntegrationIDColumn = "integration_id"
	// CodequalityTriggerModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	CodequalityTriggerModelIntegrationInstanceIDColumn = "integration_instance_id"
	// CodequalityTriggerModelRefIDColumn is the column json value ref_id
	CodequalityTriggerModelRefIDColumn = "ref_id"
	// CodequalityTriggerModelRefTypeColumn is the column json value ref_type
	CodequalityTriggerModelRefTypeColumn = "ref_type"
)

// CodequalityTrigger used to trigger an agent.CodequalityRequest
type CodequalityTrigger struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationID the integration id
	IntegrationID string `json:"integration_id" codec:"integration_id" bson:"integration_id" yaml:"integration_id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*CodequalityTrigger)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*CodequalityTrigger)(nil)

func toCodequalityTriggerObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CodequalityTrigger:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of CodequalityTrigger
func (o *CodequalityTrigger) String() string {
	return fmt.Sprintf("agent.CodequalityTrigger<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *CodequalityTrigger) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *CodequalityTrigger) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *CodequalityTrigger) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *CodequalityTrigger) GetModelName() datamodel.ModelNameType {
	return CodequalityTriggerModelName
}

// NewCodequalityTriggerID provides a template for generating an ID field for CodequalityTrigger
func NewCodequalityTriggerID(customerID string, refType string, refID string) string {
	return hash.Values("CodequalityTrigger", customerID, refType, refID)
}

func (o *CodequalityTrigger) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("CodequalityTrigger", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *CodequalityTrigger) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *CodequalityTrigger) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *CodequalityTrigger) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *CodequalityTrigger) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *CodequalityTrigger) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *CodequalityTrigger) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *CodequalityTrigger) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *CodequalityTrigger) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *CodequalityTrigger) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = CodequalityTriggerModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *CodequalityTrigger) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *CodequalityTrigger) GetCustomerID() string {
	return o.CustomerID
}

// SetCustomerID will return the customer_id
func (o *CodequalityTrigger) SetCustomerID(id string) {
	o.CustomerID = id
}

// GetRefType will return the ref_type
func (o *CodequalityTrigger) GetRefType() string {
	return o.RefType
}

// SetRefType will return the ref_type
func (o *CodequalityTrigger) SetRefType(t string) {
	o.RefType = t
}

// Clone returns an exact copy of CodequalityTrigger
func (o *CodequalityTrigger) Clone() datamodel.Model {
	c := new(CodequalityTrigger)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *CodequalityTrigger) Anon() datamodel.Model {
	c := new(CodequalityTrigger)
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
func (o *CodequalityTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *CodequalityTrigger) UnmarshalJSON(data []byte) error {
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
func (o *CodequalityTrigger) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *CodequalityTrigger) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two CodequalityTrigger objects are equal
func (o *CodequalityTrigger) IsEqual(other *CodequalityTrigger) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *CodequalityTrigger) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id":             toCodequalityTriggerObject(o.CustomerID, false),
		"id":                      toCodequalityTriggerObject(o.ID, false),
		"integration_id":          toCodequalityTriggerObject(o.IntegrationID, false),
		"integration_instance_id": toCodequalityTriggerObject(o.IntegrationInstanceID, true),
		"ref_id":                  toCodequalityTriggerObject(o.RefID, false),
		"ref_type":                toCodequalityTriggerObject(o.RefType, false),
		"hashcode":                toCodequalityTriggerObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *CodequalityTrigger) FromMap(kv map[string]interface{}) {

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
	if val, ok := kv["integration_id"].(string); ok {
		o.IntegrationID = val
	} else {
		if val, ok := kv["integration_id"]; ok {
			if val == nil {
				o.IntegrationID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.IntegrationID = fmt.Sprintf("%v", val)
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
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *CodequalityTrigger) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.IntegrationID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *CodequalityTrigger) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *CodequalityTrigger) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// CodequalityTriggerPartial is a partial struct for upsert mutations for CodequalityTrigger
type CodequalityTriggerPartial struct {
	// IntegrationID the integration id
	IntegrationID *string `json:"integration_id,omitempty"`
}

var _ datamodel.PartialModel = (*CodequalityTriggerPartial)(nil)

// GetModelName returns the name of the model
func (o *CodequalityTriggerPartial) GetModelName() datamodel.ModelNameType {
	return CodequalityTriggerModelName
}

// ToMap returns the object as a map
func (o *CodequalityTriggerPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"integration_id": toCodequalityTriggerObject(o.IntegrationID, true),
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
func (o *CodequalityTriggerPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *CodequalityTriggerPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *CodequalityTriggerPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *CodequalityTriggerPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *CodequalityTriggerPartial) FromMap(kv map[string]interface{}) {
	if val, ok := kv["integration_id"].(*string); ok {
		o.IntegrationID = val
	} else if val, ok := kv["integration_id"].(string); ok {
		o.IntegrationID = &val
	} else {
		if val, ok := kv["integration_id"]; ok {
			if val == nil {
				o.IntegrationID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.IntegrationID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	o.setDefaults(false)
}
