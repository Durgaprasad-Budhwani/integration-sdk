// DO NOT EDIT -- generated code

// Package calendar - public calendar data models
package calendar

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	pstrings "github.com/pinpt/go-common/strings"
)

const (

	// CalendarTable is the default table name
	CalendarTable datamodel.ModelNameType = "calendar_calendar"

	// CalendarModelName is the model name
	CalendarModelName datamodel.ModelNameType = "calendar.Calendar"
)

const (
	// CalendarModelActiveColumn is the column json value active
	CalendarModelActiveColumn = "active"
	// CalendarModelCustomerIDColumn is the column json value customer_id
	CalendarModelCustomerIDColumn = "customer_id"
	// CalendarModelDescriptionColumn is the column json value description
	CalendarModelDescriptionColumn = "description"
	// CalendarModelIDColumn is the column json value id
	CalendarModelIDColumn = "id"
	// CalendarModelNameColumn is the column json value name
	CalendarModelNameColumn = "name"
	// CalendarModelRefIDColumn is the column json value ref_id
	CalendarModelRefIDColumn = "ref_id"
	// CalendarModelRefTypeColumn is the column json value ref_type
	CalendarModelRefTypeColumn = "ref_type"
)

// Calendar details for the given integration calendarsitory
type Calendar struct {
	// Active the status of the calendar
	Active bool `json:"active" codec:"active" bson:"active" yaml:"active" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Description the description of the calendar
	Description string `json:"description" codec:"description" bson:"description" yaml:"description" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Name the name of the calendar
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// RefID the calendar ID
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the record type
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Calendar)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Calendar)(nil)

func toCalendarObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Calendar:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of Calendar
func (o *Calendar) String() string {
	return fmt.Sprintf("calendar.Calendar<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Calendar) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *Calendar) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Calendar) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *Calendar) GetModelName() datamodel.ModelNameType {
	return CalendarModelName
}

// NewCalendarID provides a template for generating an ID field for Calendar
func NewCalendarID(customerID string, refType string, refID string) string {
	return hash.Values("Calendar", customerID, refType, refID)
}

func (o *Calendar) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Calendar", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Calendar) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Calendar) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Calendar) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *Calendar) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Calendar) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Calendar) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Calendar) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Calendar) IsEvented() bool {
	return false
}

// GetTopicConfig returns the topic config object
func (o *Calendar) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *Calendar) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of Calendar
func (o *Calendar) Clone() datamodel.Model {
	c := new(Calendar)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Calendar) Anon() datamodel.Model {
	c := new(Calendar)
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
func (o *Calendar) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Calendar) UnmarshalJSON(data []byte) error {
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
func (o *Calendar) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *Calendar) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two Calendar objects are equal
func (o *Calendar) IsEqual(other *Calendar) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Calendar) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"active":      toCalendarObject(o.Active, false),
		"customer_id": toCalendarObject(o.CustomerID, false),
		"description": toCalendarObject(o.Description, false),
		"id":          toCalendarObject(o.ID, false),
		"name":        toCalendarObject(o.Name, false),
		"ref_id":      toCalendarObject(o.RefID, false),
		"ref_type":    toCalendarObject(o.RefType, false),
		"hashcode":    toCalendarObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Calendar) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["active"].(bool); ok {
		o.Active = val
	} else {
		if val, ok := kv["active"]; ok {
			if val == nil {
				o.Active = false
			} else {
				o.Active = number.ToBoolAny(val)
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

	if val, ok := kv["description"].(string); ok {
		o.Description = val
	} else {
		if val, ok := kv["description"]; ok {
			if val == nil {
				o.Description = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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

	if val, ok := kv["name"].(string); ok {
		o.Name = val
	} else {
		if val, ok := kv["name"]; ok {
			if val == nil {
				o.Name = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Name = fmt.Sprintf("%v", val)
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
func (o *Calendar) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Active)
	args = append(args, o.CustomerID)
	args = append(args, o.Description)
	args = append(args, o.ID)
	args = append(args, o.Name)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *Calendar) GetEventAPIConfig() datamodel.EventAPIConfig {
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