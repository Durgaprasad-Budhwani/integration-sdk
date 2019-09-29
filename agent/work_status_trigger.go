// DO NOT EDIT -- generated code

// Package agent - the agent communicates with pinpoint cloud to send integration data
package agent

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
)

const (
	// WorkStatusTriggerTopic is the default topic name
	WorkStatusTriggerTopic datamodel.TopicNameType = "agent_WorkStatusTrigger_topic"

	// WorkStatusTriggerTable is the default table name
	WorkStatusTriggerTable datamodel.ModelNameType = "agent_workstatustrigger"

	// WorkStatusTriggerModelName is the model name
	WorkStatusTriggerModelName datamodel.ModelNameType = "agent.WorkStatusTrigger"
)

const (
	// WorkStatusTriggerCustomerIDColumn is the customer_id column name
	WorkStatusTriggerCustomerIDColumn = "CustomerID"
	// WorkStatusTriggerIDColumn is the id column name
	WorkStatusTriggerIDColumn = "ID"
	// WorkStatusTriggerIntegrationIDColumn is the integration_id column name
	WorkStatusTriggerIntegrationIDColumn = "IntegrationID"
	// WorkStatusTriggerRefIDColumn is the ref_id column name
	WorkStatusTriggerRefIDColumn = "RefID"
	// WorkStatusTriggerRefTypeColumn is the ref_type column name
	WorkStatusTriggerRefTypeColumn = "RefType"
	// WorkStatusTriggerUpdatedAtColumn is the updated_ts column name
	WorkStatusTriggerUpdatedAtColumn = "UpdatedAt"
)

// WorkStatusTrigger used to trigger an agent.WorkStatusRequest
type WorkStatusTrigger struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationID the integration id
	IntegrationID string `json:"integration_id" codec:"integration_id" bson:"integration_id" yaml:"integration_id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*WorkStatusTrigger)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*WorkStatusTrigger)(nil)

func toWorkStatusTriggerObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WorkStatusTrigger:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of WorkStatusTrigger
func (o *WorkStatusTrigger) String() string {
	return fmt.Sprintf("agent.WorkStatusTrigger<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *WorkStatusTrigger) GetTopicName() datamodel.TopicNameType {
	return WorkStatusTriggerTopic
}

// GetStreamName returns the name of the stream
func (o *WorkStatusTrigger) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *WorkStatusTrigger) GetTableName() string {
	return WorkStatusTriggerTable.String()
}

// GetModelName returns the name of the model
func (o *WorkStatusTrigger) GetModelName() datamodel.ModelNameType {
	return WorkStatusTriggerModelName
}

// NewWorkStatusTriggerID provides a template for generating an ID field for WorkStatusTrigger
func NewWorkStatusTriggerID(customerID string, refType string, refID string) string {
	return hash.Values("WorkStatusTrigger", customerID, refType, refID)
}

func (o *WorkStatusTrigger) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("WorkStatusTrigger", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *WorkStatusTrigger) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *WorkStatusTrigger) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *WorkStatusTrigger) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for WorkStatusTrigger")
}

// GetRefID returns the RefID for the object
func (o *WorkStatusTrigger) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *WorkStatusTrigger) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *WorkStatusTrigger) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *WorkStatusTrigger) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *WorkStatusTrigger) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = WorkStatusTriggerModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *WorkStatusTrigger) GetTopicConfig() *datamodel.ModelTopicConfig {
	retention, err := time.ParseDuration("24h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 24h0m0s. " + err.Error())
	}

	ttl, err := time.ParseDuration("0s")
	if err != nil {
		ttl = 0
	}
	if ttl == 0 && retention != 0 {
		ttl = retention // they should be the same if not set
	}
	return &datamodel.ModelTopicConfig{
		Key:               "id",
		Timestamp:         "updated_ts",
		NumPartitions:     8,
		CleanupPolicy:     datamodel.CleanupPolicy("delete"),
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetCustomerID will return the customer_id
func (o *WorkStatusTrigger) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of WorkStatusTrigger
func (o *WorkStatusTrigger) Clone() datamodel.Model {
	c := new(WorkStatusTrigger)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *WorkStatusTrigger) Anon() datamodel.Model {
	c := new(WorkStatusTrigger)
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
func (o *WorkStatusTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *WorkStatusTrigger) UnmarshalJSON(data []byte) error {
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
func (o *WorkStatusTrigger) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two WorkStatusTrigger objects are equal
func (o *WorkStatusTrigger) IsEqual(other *WorkStatusTrigger) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *WorkStatusTrigger) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id":    toWorkStatusTriggerObject(o.CustomerID, false),
		"id":             toWorkStatusTriggerObject(o.ID, false),
		"integration_id": toWorkStatusTriggerObject(o.IntegrationID, false),
		"ref_id":         toWorkStatusTriggerObject(o.RefID, false),
		"ref_type":       toWorkStatusTriggerObject(o.RefType, false),
		"updated_ts":     toWorkStatusTriggerObject(o.UpdatedAt, false),
		"hashcode":       toWorkStatusTriggerObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *WorkStatusTrigger) FromMap(kv map[string]interface{}) {

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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.IntegrationID = fmt.Sprintf("%v", val)
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
func (o *WorkStatusTrigger) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.IntegrationID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.UpdatedAt)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *WorkStatusTrigger) GetEventAPIConfig() datamodel.EventAPIConfig {
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