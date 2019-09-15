// DO NOT EDIT -- generated code

// Package work - the system which contains project work
package work

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"sync"
	"time"

	"github.com/bxcodec/faker"
	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/eventing"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	pstrings "github.com/pinpt/go-common/strings"
)

const (
	// SprintTopic is the default topic name
	SprintTopic datamodel.TopicNameType = "work_Sprint_topic"

	// SprintStream is the default stream name
	SprintStream datamodel.TopicNameType = "work_Sprint_stream"

	// SprintTable is the default table name
	SprintTable datamodel.TopicNameType = "work_sprint"

	// SprintModelName is the model name
	SprintModelName datamodel.ModelNameType = "work.Sprint"
)

const (
	// SprintCompletedDateColumn is the completed_date column name
	SprintCompletedDateColumn = "completed_date"
	// SprintCompletedDateColumnEpochColumn is the epoch column property of the CompletedDate name
	SprintCompletedDateColumnEpochColumn = "completed_date->epoch"
	// SprintCompletedDateColumnOffsetColumn is the offset column property of the CompletedDate name
	SprintCompletedDateColumnOffsetColumn = "completed_date->offset"
	// SprintCompletedDateColumnRfc3339Column is the rfc3339 column property of the CompletedDate name
	SprintCompletedDateColumnRfc3339Column = "completed_date->rfc3339"
	// SprintCustomerIDColumn is the customer_id column name
	SprintCustomerIDColumn = "customer_id"
	// SprintEndedDateColumn is the ended_date column name
	SprintEndedDateColumn = "ended_date"
	// SprintEndedDateColumnEpochColumn is the epoch column property of the EndedDate name
	SprintEndedDateColumnEpochColumn = "ended_date->epoch"
	// SprintEndedDateColumnOffsetColumn is the offset column property of the EndedDate name
	SprintEndedDateColumnOffsetColumn = "ended_date->offset"
	// SprintEndedDateColumnRfc3339Column is the rfc3339 column property of the EndedDate name
	SprintEndedDateColumnRfc3339Column = "ended_date->rfc3339"
	// SprintFetchedDateColumn is the fetched_date column name
	SprintFetchedDateColumn = "fetched_date"
	// SprintFetchedDateColumnEpochColumn is the epoch column property of the FetchedDate name
	SprintFetchedDateColumnEpochColumn = "fetched_date->epoch"
	// SprintFetchedDateColumnOffsetColumn is the offset column property of the FetchedDate name
	SprintFetchedDateColumnOffsetColumn = "fetched_date->offset"
	// SprintFetchedDateColumnRfc3339Column is the rfc3339 column property of the FetchedDate name
	SprintFetchedDateColumnRfc3339Column = "fetched_date->rfc3339"
	// SprintGoalColumn is the goal column name
	SprintGoalColumn = "goal"
	// SprintIDColumn is the id column name
	SprintIDColumn = "id"
	// SprintNameColumn is the name column name
	SprintNameColumn = "name"
	// SprintRefIDColumn is the ref_id column name
	SprintRefIDColumn = "ref_id"
	// SprintRefTypeColumn is the ref_type column name
	SprintRefTypeColumn = "ref_type"
	// SprintStartedDateColumn is the started_date column name
	SprintStartedDateColumn = "started_date"
	// SprintStartedDateColumnEpochColumn is the epoch column property of the StartedDate name
	SprintStartedDateColumnEpochColumn = "started_date->epoch"
	// SprintStartedDateColumnOffsetColumn is the offset column property of the StartedDate name
	SprintStartedDateColumnOffsetColumn = "started_date->offset"
	// SprintStartedDateColumnRfc3339Column is the rfc3339 column property of the StartedDate name
	SprintStartedDateColumnRfc3339Column = "started_date->rfc3339"
	// SprintStatusColumn is the status column name
	SprintStatusColumn = "status"
	// SprintUpdatedAtColumn is the updated_ts column name
	SprintUpdatedAtColumn = "updated_ts"
)

// SprintCompletedDate represents the object structure for completed_date
type SprintCompletedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toSprintCompletedDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toSprintCompletedDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *SprintCompletedDate:
		return v.ToMap(isavro)

	default:
		return o
	}
}

func (o *SprintCompletedDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toSprintCompletedDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toSprintCompletedDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toSprintCompletedDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *SprintCompletedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *SprintCompletedDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Epoch = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["offset"].(int64); ok {
		o.Offset = val
	} else {
		if val, ok := kv["offset"]; ok {
			if val == nil {
				o.Offset = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Offset = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["rfc3339"].(string); ok {
		o.Rfc3339 = val
	} else {
		if val, ok := kv["rfc3339"]; ok {
			if val == nil {
				o.Rfc3339 = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// SprintEndedDate represents the object structure for ended_date
type SprintEndedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toSprintEndedDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toSprintEndedDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *SprintEndedDate:
		return v.ToMap(isavro)

	default:
		return o
	}
}

func (o *SprintEndedDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toSprintEndedDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toSprintEndedDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toSprintEndedDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *SprintEndedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *SprintEndedDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Epoch = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["offset"].(int64); ok {
		o.Offset = val
	} else {
		if val, ok := kv["offset"]; ok {
			if val == nil {
				o.Offset = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Offset = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["rfc3339"].(string); ok {
		o.Rfc3339 = val
	} else {
		if val, ok := kv["rfc3339"]; ok {
			if val == nil {
				o.Rfc3339 = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// SprintFetchedDate represents the object structure for fetched_date
type SprintFetchedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toSprintFetchedDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toSprintFetchedDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *SprintFetchedDate:
		return v.ToMap(isavro)

	default:
		return o
	}
}

func (o *SprintFetchedDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toSprintFetchedDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toSprintFetchedDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toSprintFetchedDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *SprintFetchedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *SprintFetchedDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Epoch = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["offset"].(int64); ok {
		o.Offset = val
	} else {
		if val, ok := kv["offset"]; ok {
			if val == nil {
				o.Offset = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Offset = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["rfc3339"].(string); ok {
		o.Rfc3339 = val
	} else {
		if val, ok := kv["rfc3339"]; ok {
			if val == nil {
				o.Rfc3339 = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// SprintStartedDate represents the object structure for started_date
type SprintStartedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toSprintStartedDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toSprintStartedDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *SprintStartedDate:
		return v.ToMap(isavro)

	default:
		return o
	}
}

func (o *SprintStartedDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toSprintStartedDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toSprintStartedDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toSprintStartedDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *SprintStartedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *SprintStartedDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Epoch = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["offset"].(int64); ok {
		o.Offset = val
	} else {
		if val, ok := kv["offset"]; ok {
			if val == nil {
				o.Offset = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Offset = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["rfc3339"].(string); ok {
		o.Rfc3339 = val
	} else {
		if val, ok := kv["rfc3339"]; ok {
			if val == nil {
				o.Rfc3339 = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// SprintStatus is the enumeration type for status
type SprintStatus int32

// String returns the string value for Status
func (v SprintStatus) String() string {
	switch int32(v) {
	case 0:
		return "ACTIVE"
	case 1:
		return "FUTURE"
	case 2:
		return "CLOSED"
	}
	return "unset"
}

const (
	// StatusActive is the enumeration value for active
	SprintStatusActive SprintStatus = 0
	// StatusFuture is the enumeration value for future
	SprintStatusFuture SprintStatus = 1
	// StatusClosed is the enumeration value for closed
	SprintStatusClosed SprintStatus = 2
)

// Sprint sprint details
type Sprint struct {
	// CompletedDate the date that the sprint was completed
	CompletedDate SprintCompletedDate `json:"completed_date" bson:"completed_date" yaml:"completed_date" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-" badgerhold:"index"`
	// EndedDate the date that the sprint was ended
	EndedDate SprintEndedDate `json:"ended_date" bson:"ended_date" yaml:"ended_date" faker:"-"`
	// FetchedDate date in when the api was called
	FetchedDate SprintFetchedDate `json:"fetched_date" bson:"fetched_date" yaml:"fetched_date" faker:"-"`
	// Goal the goal for the sprint
	Goal string `json:"goal" bson:"goal" yaml:"goal" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// Name the name of the field
	Name string `json:"name" bson:"name" yaml:"name" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// StartedDate the date that the sprint was started
	StartedDate SprintStartedDate `json:"started_date" bson:"started_date" yaml:"started_date" faker:"-"`
	// Status status of the sprint
	Status SprintStatus `json:"status" bson:"status" yaml:"status" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Sprint)(nil)

func toSprintObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toSprintObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *Sprint:
		return v.ToMap(isavro)

	case SprintCompletedDate:
		return v.ToMap(isavro)

	case SprintEndedDate:
		return v.ToMap(isavro)

	case SprintFetchedDate:
		return v.ToMap(isavro)

	case SprintStartedDate:
		return v.ToMap(isavro)

	case SprintStatus:
		return v.String()

	default:
		return o
	}
}

// String returns a string representation of Sprint
func (o *Sprint) String() string {
	return fmt.Sprintf("work.Sprint<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Sprint) GetTopicName() datamodel.TopicNameType {
	return SprintTopic
}

// GetModelName returns the name of the model
func (o *Sprint) GetModelName() datamodel.ModelNameType {
	return SprintModelName
}

// GetStreamName returns the name of the stream
func (o *Sprint) GetStreamName() string {
	return SprintStream.String()
}

// GetTableName returns the name of the table
func (o *Sprint) GetTableName() string {
	return SprintTable.String()
}

// NewSprintID provides a template for generating an ID field for Sprint
func NewSprintID(customerID string, refID string, refType string) string {
	return hash.Values(customerID, refID, refType)
}

func (o *Sprint) setDefaults(frommap bool) {

	if o.ID == "" {
		o.ID = hash.Values(o.CustomerID, o.RefID, o.RefType)
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Sprint) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Sprint) GetTopicKey() string {
	var i interface{} = o.CustomerID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Sprint) GetTimestamp() time.Time {
	var dt interface{} = o.StartedDate
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
	case SprintStartedDate:
		return datetime.DateFromEpoch(v.Epoch)
	}
	panic("not sure how to handle the date time format for Sprint")
}

// GetRefID returns the RefID for the object
func (o *Sprint) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Sprint) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Sprint) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Sprint) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Sprint) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = SprintModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Sprint) GetTopicConfig() *datamodel.ModelTopicConfig {
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
		Timestamp:         "started_date",
		NumPartitions:     8,
		CleanupPolicy:     datamodel.CleanupPolicy("compact"),
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetStateKey returns a key for use in state store
func (o *Sprint) GetStateKey() string {
	key := "customer_id"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *Sprint) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of Sprint
func (o *Sprint) Clone() datamodel.Model {
	c := new(Sprint)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Sprint) Anon() datamodel.Model {
	c := new(Sprint)
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

// MarshalBinary returns the bytes for marshaling to binary
func (o *Sprint) MarshalBinary() ([]byte, error) {
	return o.MarshalJSON()
}

func (o *Sprint) UnmarshalBinary(data []byte) error {
	return o.UnmarshalJSON(data)
}

// MarshalJSON returns the bytes for marshaling to json
func (o *Sprint) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Sprint) UnmarshalJSON(data []byte) error {
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

var cachedCodecSprint *goavro.Codec
var cachedCodecSprintLock sync.Mutex

// GetAvroCodec returns the avro codec for this model
func (o *Sprint) GetAvroCodec() *goavro.Codec {
	cachedCodecSprintLock.Lock()
	if cachedCodecSprint == nil {
		c, err := GetSprintAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecSprint = c
	}
	cachedCodecSprintLock.Unlock()
	return cachedCodecSprint
}

// ToAvroBinary returns the data as Avro binary data
func (o *Sprint) ToAvroBinary() ([]byte, *goavro.Codec, error) {
	kv := o.ToMap(true)
	jbuf, _ := json.Marshal(kv)
	codec := o.GetAvroCodec()
	native, _, err := codec.NativeFromTextual(jbuf)
	if err != nil {
		return nil, nil, err
	}
	// Convert native Go form to binary Avro data
	buf, err := codec.BinaryFromNative(nil, native)
	return buf, codec, err
}

// FromAvroBinary will convert from Avro binary data into data in this object
func (o *Sprint) FromAvroBinary(value []byte) error {
	var nullHeader = []byte{byte(0)}
	// if this still has the schema encoded in the header, move past it to the avro payload
	if bytes.HasPrefix(value, nullHeader) {
		value = value[5:]
	}
	kv, _, err := o.GetAvroCodec().NativeFromBinary(value)
	if err != nil {
		return err
	}
	o.FromMap(kv.(map[string]interface{}))
	return nil
}

// Stringify returns the object in JSON format as a string
func (o *Sprint) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Sprint objects are equal
func (o *Sprint) IsEqual(other *Sprint) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Sprint) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	o.setDefaults(false)
	return map[string]interface{}{
		"completed_date": toSprintObject(o.CompletedDate, isavro, false, "completed_date"),
		"customer_id":    toSprintObject(o.CustomerID, isavro, false, "string"),
		"ended_date":     toSprintObject(o.EndedDate, isavro, false, "ended_date"),
		"fetched_date":   toSprintObject(o.FetchedDate, isavro, false, "fetched_date"),
		"goal":           toSprintObject(o.Goal, isavro, false, "string"),
		"id":             toSprintObject(o.ID, isavro, false, "string"),
		"name":           toSprintObject(o.Name, isavro, false, "string"),
		"ref_id":         toSprintObject(o.RefID, isavro, false, "string"),
		"ref_type":       toSprintObject(o.RefType, isavro, false, "string"),
		"started_date":   toSprintObject(o.StartedDate, isavro, false, "started_date"),
		"status":         toSprintObject(o.Status, isavro, false, "status"),
		"updated_ts":     toSprintObject(o.UpdatedAt, isavro, false, "long"),
		"hashcode":       toSprintObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Sprint) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["completed_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CompletedDate.FromMap(kv)
		} else if sv, ok := val.(SprintCompletedDate); ok {
			// struct
			o.CompletedDate = sv
		} else if sp, ok := val.(*SprintCompletedDate); ok {
			// struct pointer
			o.CompletedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.CompletedDate.Epoch = dt.Epoch
			o.CompletedDate.Rfc3339 = dt.Rfc3339
			o.CompletedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.CompletedDate.Epoch = dt.Epoch
			o.CompletedDate.Rfc3339 = dt.Rfc3339
			o.CompletedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.CompletedDate.Epoch = dt.Epoch
				o.CompletedDate.Rfc3339 = dt.Rfc3339
				o.CompletedDate.Offset = dt.Offset
			}
		}
	} else {
		o.CompletedDate.FromMap(map[string]interface{}{})
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

	if val, ok := kv["ended_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.EndedDate.FromMap(kv)
		} else if sv, ok := val.(SprintEndedDate); ok {
			// struct
			o.EndedDate = sv
		} else if sp, ok := val.(*SprintEndedDate); ok {
			// struct pointer
			o.EndedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.EndedDate.Epoch = dt.Epoch
			o.EndedDate.Rfc3339 = dt.Rfc3339
			o.EndedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.EndedDate.Epoch = dt.Epoch
			o.EndedDate.Rfc3339 = dt.Rfc3339
			o.EndedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.EndedDate.Epoch = dt.Epoch
				o.EndedDate.Rfc3339 = dt.Rfc3339
				o.EndedDate.Offset = dt.Offset
			}
		}
	} else {
		o.EndedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["fetched_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.FetchedDate.FromMap(kv)
		} else if sv, ok := val.(SprintFetchedDate); ok {
			// struct
			o.FetchedDate = sv
		} else if sp, ok := val.(*SprintFetchedDate); ok {
			// struct pointer
			o.FetchedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.FetchedDate.Epoch = dt.Epoch
			o.FetchedDate.Rfc3339 = dt.Rfc3339
			o.FetchedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.FetchedDate.Epoch = dt.Epoch
			o.FetchedDate.Rfc3339 = dt.Rfc3339
			o.FetchedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.FetchedDate.Epoch = dt.Epoch
				o.FetchedDate.Rfc3339 = dt.Rfc3339
				o.FetchedDate.Offset = dt.Offset
			}
		}
	} else {
		o.FetchedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["goal"].(string); ok {
		o.Goal = val
	} else {
		if val, ok := kv["goal"]; ok {
			if val == nil {
				o.Goal = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Goal = fmt.Sprintf("%v", val)
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

	if val, ok := kv["started_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.StartedDate.FromMap(kv)
		} else if sv, ok := val.(SprintStartedDate); ok {
			// struct
			o.StartedDate = sv
		} else if sp, ok := val.(*SprintStartedDate); ok {
			// struct pointer
			o.StartedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.StartedDate.Epoch = dt.Epoch
			o.StartedDate.Rfc3339 = dt.Rfc3339
			o.StartedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.StartedDate.Epoch = dt.Epoch
			o.StartedDate.Rfc3339 = dt.Rfc3339
			o.StartedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.StartedDate.Epoch = dt.Epoch
				o.StartedDate.Rfc3339 = dt.Rfc3339
				o.StartedDate.Offset = dt.Offset
			}
		}
	} else {
		o.StartedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["status"].(SprintStatus); ok {
		o.Status = val
	} else {
		if em, ok := kv["status"].(map[string]interface{}); ok {
			ev := em["work.status"].(string)
			switch ev {
			case "active", "ACTIVE":
				o.Status = 0
			case "future", "FUTURE":
				o.Status = 1
			case "closed", "CLOSED":
				o.Status = 2
			}
		}
		if em, ok := kv["status"].(string); ok {
			switch em {
			case "active", "ACTIVE":
				o.Status = 0
			case "future", "FUTURE":
				o.Status = 1
			case "closed", "CLOSED":
				o.Status = 2
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
func (o *Sprint) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CompletedDate)
	args = append(args, o.CustomerID)
	args = append(args, o.EndedDate)
	args = append(args, o.FetchedDate)
	args = append(args, o.Goal)
	args = append(args, o.ID)
	args = append(args, o.Name)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.StartedDate)
	args = append(args, o.Status)
	args = append(args, o.UpdatedAt)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetSprintAvroSchemaSpec creates the avro schema specification for Sprint
func GetSprintAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "work",
		"name":      "Sprint",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name": "completed_date",
				"type": map[string]interface{}{"doc": "the date that the sprint was completed", "fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "name": "epoch", "type": "long"}, map[string]interface{}{"doc": "the timezone offset from GMT", "name": "offset", "type": "long"}, map[string]interface{}{"doc": "the date in RFC3339 format", "name": "rfc3339", "type": "string"}}, "name": "completed_date", "type": "record"},
			},
			map[string]interface{}{
				"name": "customer_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "ended_date",
				"type": map[string]interface{}{"doc": "the date that the sprint was ended", "fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "name": "epoch", "type": "long"}, map[string]interface{}{"doc": "the timezone offset from GMT", "name": "offset", "type": "long"}, map[string]interface{}{"doc": "the date in RFC3339 format", "name": "rfc3339", "type": "string"}}, "name": "ended_date", "type": "record"},
			},
			map[string]interface{}{
				"name": "fetched_date",
				"type": map[string]interface{}{"doc": "date in when the api was called", "fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "name": "epoch", "type": "long"}, map[string]interface{}{"doc": "the timezone offset from GMT", "name": "offset", "type": "long"}, map[string]interface{}{"doc": "the date in RFC3339 format", "name": "rfc3339", "type": "string"}}, "name": "fetched_date", "type": "record"},
			},
			map[string]interface{}{
				"name": "goal",
				"type": "string",
			},
			map[string]interface{}{
				"name": "id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "name",
				"type": "string",
			},
			map[string]interface{}{
				"name": "ref_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "ref_type",
				"type": "string",
			},
			map[string]interface{}{
				"name": "started_date",
				"type": map[string]interface{}{"doc": "the date that the sprint was started", "fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "name": "epoch", "type": "long"}, map[string]interface{}{"doc": "the timezone offset from GMT", "name": "offset", "type": "long"}, map[string]interface{}{"doc": "the date in RFC3339 format", "name": "rfc3339", "type": "string"}}, "name": "started_date", "type": "record"},
			},
			map[string]interface{}{
				"name": "status",
				"type": map[string]interface{}{
					"type":    "enum",
					"name":    "status",
					"symbols": []interface{}{"ACTIVE", "FUTURE", "CLOSED"},
				},
			},
			map[string]interface{}{
				"name": "updated_ts",
				"type": "long",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *Sprint) GetEventAPIConfig() datamodel.EventAPIConfig {
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

// GetSprintAvroSchema creates the avro schema for Sprint
func GetSprintAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetSprintAvroSchemaSpec())
}

// SprintSendEvent is an event detail for sending data
type SprintSendEvent struct {
	Sprint  *Sprint
	headers map[string]string
	time    time.Time
	key     string
}

var _ datamodel.ModelSendEvent = (*SprintSendEvent)(nil)

// Key is the key to use for the message
func (e *SprintSendEvent) Key() string {
	if e.key == "" {
		return e.Sprint.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *SprintSendEvent) Object() datamodel.Model {
	return e.Sprint
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *SprintSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *SprintSendEvent) Timestamp() time.Time {
	return e.time
}

// SprintSendEventOpts is a function handler for setting opts
type SprintSendEventOpts func(o *SprintSendEvent)

// WithSprintSendEventKey sets the key value to a value different than the object ID
func WithSprintSendEventKey(key string) SprintSendEventOpts {
	return func(o *SprintSendEvent) {
		o.key = key
	}
}

// WithSprintSendEventTimestamp sets the timestamp value
func WithSprintSendEventTimestamp(tv time.Time) SprintSendEventOpts {
	return func(o *SprintSendEvent) {
		o.time = tv
	}
}

// WithSprintSendEventHeader sets the timestamp value
func WithSprintSendEventHeader(key, value string) SprintSendEventOpts {
	return func(o *SprintSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewSprintSendEvent returns a new SprintSendEvent instance
func NewSprintSendEvent(o *Sprint, opts ...SprintSendEventOpts) *SprintSendEvent {
	res := &SprintSendEvent{
		Sprint: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewSprintProducer will stream data from the channel
func NewSprintProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
	done := make(chan bool, 1)
	emptyTime := time.Unix(0, 0)
	var numPartitions int
	go func() {
		defer func() { done <- true }()
		for {
			select {
			case <-ctx.Done():
				return
			case item := <-ch:
				if item == nil {
					empty <- true
					return
				}
				if object, ok := item.Object().(*Sprint); ok {
					if numPartitions == 0 {
						numPartitions = object.GetTopicConfig().NumPartitions
					}
					binary, codec, err := object.ToAvroBinary()
					if err != nil {
						errors <- fmt.Errorf("error encoding %s to avro binary data. %v", object.String(), err)
						return
					}
					headers := map[string]string{}
					object.SetEventHeaders(headers)
					for k, v := range item.Headers() {
						headers[k] = v
					}
					tv := item.Timestamp()
					if tv.IsZero() {
						tv = object.GetTimestamp() // if not provided in the message, use the objects value
					}
					if tv.IsZero() || tv.Equal(emptyTime) {
						tv = time.Now() // if its still zero, use the ingest time
					}
					// add generated message headers
					headers["message-id"] = pstrings.NewUUIDV4()
					headers["message-ts"] = fmt.Sprintf("%v", datetime.EpochNow())
					// determine the partition selection by using the partition key
					// and taking the modulo over the number of partitions for the topic
					partition := hash.Modulo(item.Key(), numPartitions)
					msg := eventing.Message{
						Encoding:  eventing.AvroEncoding,
						Key:       object.GetID(),
						Value:     binary,
						Codec:     codec,
						Headers:   headers,
						Timestamp: tv,
						Partition: int32(partition),
						Topic:     object.GetTopicName().String(),
					}
					if err := producer.Send(ctx, msg); err != nil {
						errors <- fmt.Errorf("error sending %s. %v", object.String(), err)
					}
				} else {
					errors <- fmt.Errorf("invalid event received. expected an object of type work.Sprint but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewSprintConsumer will stream data from the topic into the provided channel
func NewSprintConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object Sprint
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into work.Sprint: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into work.Sprint: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for work.Sprint")
			}

			// ignore messages that have exceeded the TTL
			cfg := object.GetTopicConfig()
			if cfg != nil && cfg.TTL != 0 && msg.Timestamp.UTC().Add(cfg.TTL).Sub(time.Now().UTC()) < 0 {
				// if disable auto and we're skipping, we need to commit the message
				if !msg.IsAutoCommit() {
					msg.Commit()
				}
				return nil
			}
			msg.Codec = object.GetAvroCodec() // match the codec

			ch <- &SprintReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object Sprint
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &SprintReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// SprintReceiveEvent is an event detail for receiving data
type SprintReceiveEvent struct {
	Sprint  *Sprint
	message eventing.Message
	eof     bool
}

var _ datamodel.ModelReceiveEvent = (*SprintReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *SprintReceiveEvent) Object() datamodel.Model {
	return e.Sprint
}

// Message returns the underlying message data for the event
func (e *SprintReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *SprintReceiveEvent) EOF() bool {
	return e.eof
}

// SprintProducer implements the datamodel.ModelEventProducer
type SprintProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*SprintProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *SprintProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *SprintProducer) Close() error {
	p.mu.Lock()
	closed := p.closed
	p.closed = true
	p.mu.Unlock()
	if !closed {
		close(p.ch)
		<-p.empty
		p.cancel()
		<-p.done
	}
	return nil
}

// NewProducerChannel returns a channel which can be used for producing Model events
func (o *Sprint) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *Sprint) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &SprintProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewSprintProducer(newctx, producer, ch, errors, empty),
	}
}

// NewSprintProducerChannel returns a channel which can be used for producing Model events
func NewSprintProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewSprintProducerChannelSize(producer, 0, errors)
}

// NewSprintProducerChannelSize returns a channel which can be used for producing Model events
func NewSprintProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &SprintProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewSprintProducer(newctx, producer, ch, errors, empty),
	}
}

// SprintConsumer implements the datamodel.ModelEventConsumer
type SprintConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*SprintConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *SprintConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *SprintConsumer) Close() error {
	c.mu.Lock()
	closed := c.closed
	c.closed = true
	c.mu.Unlock()
	var err error
	if !closed {
		c.callback.Close()
		err = c.consumer.Close()
	}
	return err
}

// NewConsumerChannel returns a consumer channel which can be used to consume Model events
func (o *Sprint) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &SprintConsumer{
		ch:       ch,
		callback: NewSprintConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewSprintConsumerChannel returns a consumer channel which can be used to consume Model events
func NewSprintConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &SprintConsumer{
		ch:       ch,
		callback: NewSprintConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
