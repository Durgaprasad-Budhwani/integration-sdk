// DO NOT EDIT -- generated code

// Package agent -
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
	pstrings "github.com/pinpt/go-common/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

const (

	// IntegrationDataResponseTable is the default table name
	IntegrationDataResponseTable datamodel.ModelNameType = "agent_integrationdataresponse"

	// IntegrationDataResponseModelName is the model name
	IntegrationDataResponseModelName datamodel.ModelNameType = "agent.IntegrationDataResponse"
)

const (
	// IntegrationDataResponseModelArchitectureColumn is the column json value architecture
	IntegrationDataResponseModelArchitectureColumn = "architecture"
	// IntegrationDataResponseModelCustomerIDColumn is the column json value customer_id
	IntegrationDataResponseModelCustomerIDColumn = "customer_id"
	// IntegrationDataResponseModelDataColumn is the column json value data
	IntegrationDataResponseModelDataColumn = "data"
	// IntegrationDataResponseModelDistroColumn is the column json value distro
	IntegrationDataResponseModelDistroColumn = "distro"
	// IntegrationDataResponseModelErrorColumn is the column json value error
	IntegrationDataResponseModelErrorColumn = "error"
	// IntegrationDataResponseModelEventDateColumn is the column json value event_date
	IntegrationDataResponseModelEventDateColumn = "event_date"
	// IntegrationDataResponseModelEventDateEpochColumn is the column json value epoch
	IntegrationDataResponseModelEventDateEpochColumn = "epoch"
	// IntegrationDataResponseModelEventDateOffsetColumn is the column json value offset
	IntegrationDataResponseModelEventDateOffsetColumn = "offset"
	// IntegrationDataResponseModelEventDateRfc3339Column is the column json value rfc3339
	IntegrationDataResponseModelEventDateRfc3339Column = "rfc3339"
	// IntegrationDataResponseModelFreeSpaceColumn is the column json value free_space
	IntegrationDataResponseModelFreeSpaceColumn = "free_space"
	// IntegrationDataResponseModelGoVersionColumn is the column json value go_version
	IntegrationDataResponseModelGoVersionColumn = "go_version"
	// IntegrationDataResponseModelHostnameColumn is the column json value hostname
	IntegrationDataResponseModelHostnameColumn = "hostname"
	// IntegrationDataResponseModelIDColumn is the column json value id
	IntegrationDataResponseModelIDColumn = "id"
	// IntegrationDataResponseModelJobIDColumn is the column json value job_id
	IntegrationDataResponseModelJobIDColumn = "job_id"
	// IntegrationDataResponseModelLastExportDateColumn is the column json value last_export_date
	IntegrationDataResponseModelLastExportDateColumn = "last_export_date"
	// IntegrationDataResponseModelLastExportDateEpochColumn is the column json value epoch
	IntegrationDataResponseModelLastExportDateEpochColumn = "epoch"
	// IntegrationDataResponseModelLastExportDateOffsetColumn is the column json value offset
	IntegrationDataResponseModelLastExportDateOffsetColumn = "offset"
	// IntegrationDataResponseModelLastExportDateRfc3339Column is the column json value rfc3339
	IntegrationDataResponseModelLastExportDateRfc3339Column = "rfc3339"
	// IntegrationDataResponseModelMemoryColumn is the column json value memory
	IntegrationDataResponseModelMemoryColumn = "memory"
	// IntegrationDataResponseModelMessageColumn is the column json value message
	IntegrationDataResponseModelMessageColumn = "message"
	// IntegrationDataResponseModelNumCPUColumn is the column json value num_cpu
	IntegrationDataResponseModelNumCPUColumn = "num_cpu"
	// IntegrationDataResponseModelOSColumn is the column json value os
	IntegrationDataResponseModelOSColumn = "os"
	// IntegrationDataResponseModelRefIDColumn is the column json value ref_id
	IntegrationDataResponseModelRefIDColumn = "ref_id"
	// IntegrationDataResponseModelRefTypeColumn is the column json value ref_type
	IntegrationDataResponseModelRefTypeColumn = "ref_type"
	// IntegrationDataResponseModelRequestIDColumn is the column json value request_id
	IntegrationDataResponseModelRequestIDColumn = "request_id"
	// IntegrationDataResponseModelSuccessColumn is the column json value success
	IntegrationDataResponseModelSuccessColumn = "success"
	// IntegrationDataResponseModelSystemIDColumn is the column json value system_id
	IntegrationDataResponseModelSystemIDColumn = "system_id"
	// IntegrationDataResponseModelTypeColumn is the column json value type
	IntegrationDataResponseModelTypeColumn = "type"
	// IntegrationDataResponseModelUptimeColumn is the column json value uptime
	IntegrationDataResponseModelUptimeColumn = "uptime"
	// IntegrationDataResponseModelUUIDColumn is the column json value uuid
	IntegrationDataResponseModelUUIDColumn = "uuid"
	// IntegrationDataResponseModelVersionColumn is the column json value version
	IntegrationDataResponseModelVersionColumn = "version"
)

// IntegrationDataResponseEventDate represents the object structure for event_date
type IntegrationDataResponseEventDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationDataResponseEventDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationDataResponseEventDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *IntegrationDataResponseEventDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationDataResponseEventDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationDataResponseEventDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationDataResponseEventDateObject(o.Rfc3339, false),
	}
}

func (o *IntegrationDataResponseEventDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationDataResponseEventDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = 0
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
				o.Offset = 0
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// IntegrationDataResponseLastExportDate represents the object structure for last_export_date
type IntegrationDataResponseLastExportDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationDataResponseLastExportDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationDataResponseLastExportDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *IntegrationDataResponseLastExportDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationDataResponseLastExportDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationDataResponseLastExportDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationDataResponseLastExportDateObject(o.Rfc3339, false),
	}
}

func (o *IntegrationDataResponseLastExportDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationDataResponseLastExportDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = 0
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
				o.Offset = 0
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// IntegrationDataResponseType is the enumeration type for type
type IntegrationDataResponseType int32

// UnmarshalBSONValue for unmarshaling value
func (v *IntegrationDataResponseType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = IntegrationDataResponseType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "ENROLL":
			*v = IntegrationDataResponseType(0)
		case "PING":
			*v = IntegrationDataResponseType(1)
		case "CRASH":
			*v = IntegrationDataResponseType(2)
		case "LOG":
			*v = IntegrationDataResponseType(3)
		case "INTEGRATION":
			*v = IntegrationDataResponseType(4)
		case "EXPORT":
			*v = IntegrationDataResponseType(5)
		case "PROJECT":
			*v = IntegrationDataResponseType(6)
		case "REPO":
			*v = IntegrationDataResponseType(7)
		case "USER":
			*v = IntegrationDataResponseType(8)
		case "UNINSTALL":
			*v = IntegrationDataResponseType(9)
		case "UPGRADE":
			*v = IntegrationDataResponseType(10)
		case "START":
			*v = IntegrationDataResponseType(11)
		case "STOP":
			*v = IntegrationDataResponseType(12)
		case "PAUSE":
			*v = IntegrationDataResponseType(13)
		case "RESUME":
			*v = IntegrationDataResponseType(14)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v IntegrationDataResponseType) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "ENROLL":
		v = 0
	case "PING":
		v = 1
	case "CRASH":
		v = 2
	case "LOG":
		v = 3
	case "INTEGRATION":
		v = 4
	case "EXPORT":
		v = 5
	case "PROJECT":
		v = 6
	case "REPO":
		v = 7
	case "USER":
		v = 8
	case "UNINSTALL":
		v = 9
	case "UPGRADE":
		v = 10
	case "START":
		v = 11
	case "STOP":
		v = 12
	case "PAUSE":
		v = 13
	case "RESUME":
		v = 14
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v IntegrationDataResponseType) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("ENROLL")
	case 1:
		return json.Marshal("PING")
	case 2:
		return json.Marshal("CRASH")
	case 3:
		return json.Marshal("LOG")
	case 4:
		return json.Marshal("INTEGRATION")
	case 5:
		return json.Marshal("EXPORT")
	case 6:
		return json.Marshal("PROJECT")
	case 7:
		return json.Marshal("REPO")
	case 8:
		return json.Marshal("USER")
	case 9:
		return json.Marshal("UNINSTALL")
	case 10:
		return json.Marshal("UPGRADE")
	case 11:
		return json.Marshal("START")
	case 12:
		return json.Marshal("STOP")
	case 13:
		return json.Marshal("PAUSE")
	case 14:
		return json.Marshal("RESUME")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for Type
func (v IntegrationDataResponseType) String() string {
	switch int32(v) {
	case 0:
		return "ENROLL"
	case 1:
		return "PING"
	case 2:
		return "CRASH"
	case 3:
		return "LOG"
	case 4:
		return "INTEGRATION"
	case 5:
		return "EXPORT"
	case 6:
		return "PROJECT"
	case 7:
		return "REPO"
	case 8:
		return "USER"
	case 9:
		return "UNINSTALL"
	case 10:
		return "UPGRADE"
	case 11:
		return "START"
	case 12:
		return "STOP"
	case 13:
		return "PAUSE"
	case 14:
		return "RESUME"
	}
	return "unset"
}

const (
	// IntegrationDataResponseTypeEnroll is the enumeration value for enroll
	IntegrationDataResponseTypeEnroll IntegrationDataResponseType = 0
	// IntegrationDataResponseTypePing is the enumeration value for ping
	IntegrationDataResponseTypePing IntegrationDataResponseType = 1
	// IntegrationDataResponseTypeCrash is the enumeration value for crash
	IntegrationDataResponseTypeCrash IntegrationDataResponseType = 2
	// IntegrationDataResponseTypeLog is the enumeration value for log
	IntegrationDataResponseTypeLog IntegrationDataResponseType = 3
	// IntegrationDataResponseTypeIntegration is the enumeration value for integration
	IntegrationDataResponseTypeIntegration IntegrationDataResponseType = 4
	// IntegrationDataResponseTypeExport is the enumeration value for export
	IntegrationDataResponseTypeExport IntegrationDataResponseType = 5
	// IntegrationDataResponseTypeProject is the enumeration value for project
	IntegrationDataResponseTypeProject IntegrationDataResponseType = 6
	// IntegrationDataResponseTypeRepo is the enumeration value for repo
	IntegrationDataResponseTypeRepo IntegrationDataResponseType = 7
	// IntegrationDataResponseTypeUser is the enumeration value for user
	IntegrationDataResponseTypeUser IntegrationDataResponseType = 8
	// IntegrationDataResponseTypeUninstall is the enumeration value for uninstall
	IntegrationDataResponseTypeUninstall IntegrationDataResponseType = 9
	// IntegrationDataResponseTypeUpgrade is the enumeration value for upgrade
	IntegrationDataResponseTypeUpgrade IntegrationDataResponseType = 10
	// IntegrationDataResponseTypeStart is the enumeration value for start
	IntegrationDataResponseTypeStart IntegrationDataResponseType = 11
	// IntegrationDataResponseTypeStop is the enumeration value for stop
	IntegrationDataResponseTypeStop IntegrationDataResponseType = 12
	// IntegrationDataResponseTypePause is the enumeration value for pause
	IntegrationDataResponseTypePause IntegrationDataResponseType = 13
	// IntegrationDataResponseTypeResume is the enumeration value for resume
	IntegrationDataResponseTypeResume IntegrationDataResponseType = 14
)

// IntegrationDataResponse response for data request going back from agent (used by webapp)
type IntegrationDataResponse struct {
	// Architecture the architecture of the agent machine
	Architecture string `json:"architecture" codec:"architecture" bson:"architecture" yaml:"architecture" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Data Arbitrary response stored as json
	Data string `json:"data" codec:"data" bson:"data" yaml:"data" faker:"-"`
	// Distro the agent os distribution
	Distro string `json:"distro" codec:"distro" bson:"distro" yaml:"distro" faker:"-"`
	// Error an error message related to this event
	Error *string `json:"error,omitempty" codec:"error,omitempty" bson:"error" yaml:"error,omitempty" faker:"-"`
	// EventDate the date of the event
	EventDate IntegrationDataResponseEventDate `json:"event_date" codec:"event_date" bson:"event_date" yaml:"event_date" faker:"-"`
	// FreeSpace the amount of free space in bytes for the agent machine
	FreeSpace int64 `json:"free_space" codec:"free_space" bson:"free_space" yaml:"free_space" faker:"-"`
	// GoVersion the go version that the agent build was built with
	GoVersion string `json:"go_version" codec:"go_version" bson:"go_version" yaml:"go_version" faker:"-"`
	// Hostname the agent hostname
	Hostname string `json:"hostname" codec:"hostname" bson:"hostname" yaml:"hostname" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// JobID The job ID
	JobID string `json:"job_id" codec:"job_id" bson:"job_id" yaml:"job_id" faker:"-"`
	// LastExportDate the last export date
	LastExportDate IntegrationDataResponseLastExportDate `json:"last_export_date" codec:"last_export_date" bson:"last_export_date" yaml:"last_export_date" faker:"-"`
	// Memory the amount of memory in bytes for the agent machine
	Memory int64 `json:"memory" codec:"memory" bson:"memory" yaml:"memory" faker:"-"`
	// Message a message related to this event
	Message string `json:"message" codec:"message" bson:"message" yaml:"message" faker:"-"`
	// NumCPU the number of CPU the agent is running
	NumCPU int64 `json:"num_cpu" codec:"num_cpu" bson:"num_cpu" yaml:"num_cpu" faker:"-"`
	// OS the agent operating system
	OS string `json:"os" codec:"os" bson:"os" yaml:"os" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RequestID the request id that this response is correlated to
	RequestID string `json:"request_id" codec:"request_id" bson:"request_id" yaml:"request_id" faker:"-"`
	// Success if the response was successful
	Success bool `json:"success" codec:"success" bson:"success" yaml:"success" faker:"-"`
	// SystemID system unique device ID
	SystemID string `json:"system_id" codec:"system_id" bson:"system_id" yaml:"system_id" faker:"-"`
	// Type the type of event
	Type IntegrationDataResponseType `json:"type" codec:"type" bson:"type" yaml:"type" faker:"-"`
	// Uptime the uptime in milliseconds since the agent started
	Uptime int64 `json:"uptime" codec:"uptime" bson:"uptime" yaml:"uptime" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" codec:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Version the agent version
	Version string `json:"version" codec:"version" bson:"version" yaml:"version" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*IntegrationDataResponse)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*IntegrationDataResponse)(nil)

func toIntegrationDataResponseObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationDataResponse:
		return v.ToMap()

	case IntegrationDataResponseEventDate:
		return v.ToMap()

	case IntegrationDataResponseLastExportDate:
		return v.ToMap()

	case IntegrationDataResponseType:
		return v.String()

	default:
		return o
	}
}

// String returns a string representation of IntegrationDataResponse
func (o *IntegrationDataResponse) String() string {
	return fmt.Sprintf("agent.IntegrationDataResponse<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *IntegrationDataResponse) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *IntegrationDataResponse) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *IntegrationDataResponse) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *IntegrationDataResponse) GetModelName() datamodel.ModelNameType {
	return IntegrationDataResponseModelName
}

// NewIntegrationDataResponseID provides a template for generating an ID field for IntegrationDataResponse
func NewIntegrationDataResponseID(customerID string, refType string, refID string) string {
	return hash.Values("IntegrationDataResponse", customerID, refType, refID)
}

func (o *IntegrationDataResponse) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("IntegrationDataResponse", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *IntegrationDataResponse) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *IntegrationDataResponse) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *IntegrationDataResponse) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *IntegrationDataResponse) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *IntegrationDataResponse) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *IntegrationDataResponse) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *IntegrationDataResponse) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *IntegrationDataResponse) IsEvented() bool {
	return false
}

// GetTopicConfig returns the topic config object
func (o *IntegrationDataResponse) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *IntegrationDataResponse) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of IntegrationDataResponse
func (o *IntegrationDataResponse) Clone() datamodel.Model {
	c := new(IntegrationDataResponse)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *IntegrationDataResponse) Anon() datamodel.Model {
	c := new(IntegrationDataResponse)
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
func (o *IntegrationDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *IntegrationDataResponse) UnmarshalJSON(data []byte) error {
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
func (o *IntegrationDataResponse) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *IntegrationDataResponse) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two IntegrationDataResponse objects are equal
func (o *IntegrationDataResponse) IsEqual(other *IntegrationDataResponse) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *IntegrationDataResponse) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"architecture":     toIntegrationDataResponseObject(o.Architecture, false),
		"customer_id":      toIntegrationDataResponseObject(o.CustomerID, false),
		"data":             toIntegrationDataResponseObject(o.Data, false),
		"distro":           toIntegrationDataResponseObject(o.Distro, false),
		"error":            toIntegrationDataResponseObject(o.Error, true),
		"event_date":       toIntegrationDataResponseObject(o.EventDate, false),
		"free_space":       toIntegrationDataResponseObject(o.FreeSpace, false),
		"go_version":       toIntegrationDataResponseObject(o.GoVersion, false),
		"hostname":         toIntegrationDataResponseObject(o.Hostname, false),
		"id":               toIntegrationDataResponseObject(o.ID, false),
		"job_id":           toIntegrationDataResponseObject(o.JobID, false),
		"last_export_date": toIntegrationDataResponseObject(o.LastExportDate, false),
		"memory":           toIntegrationDataResponseObject(o.Memory, false),
		"message":          toIntegrationDataResponseObject(o.Message, false),
		"num_cpu":          toIntegrationDataResponseObject(o.NumCPU, false),
		"os":               toIntegrationDataResponseObject(o.OS, false),
		"ref_id":           toIntegrationDataResponseObject(o.RefID, false),
		"ref_type":         toIntegrationDataResponseObject(o.RefType, false),
		"request_id":       toIntegrationDataResponseObject(o.RequestID, false),
		"success":          toIntegrationDataResponseObject(o.Success, false),
		"system_id":        toIntegrationDataResponseObject(o.SystemID, false),

		"type":     o.Type.String(),
		"uptime":   toIntegrationDataResponseObject(o.Uptime, false),
		"uuid":     toIntegrationDataResponseObject(o.UUID, false),
		"version":  toIntegrationDataResponseObject(o.Version, false),
		"hashcode": toIntegrationDataResponseObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationDataResponse) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["architecture"].(string); ok {
		o.Architecture = val
	} else {
		if val, ok := kv["architecture"]; ok {
			if val == nil {
				o.Architecture = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Architecture = fmt.Sprintf("%v", val)
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

	if val, ok := kv["data"].(string); ok {
		o.Data = val
	} else {
		if val, ok := kv["data"]; ok {
			if val == nil {
				o.Data = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Data = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["distro"].(string); ok {
		o.Distro = val
	} else {
		if val, ok := kv["distro"]; ok {
			if val == nil {
				o.Distro = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Distro = fmt.Sprintf("%v", val)
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

	if val, ok := kv["event_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.EventDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationDataResponseEventDate); ok {
			// struct
			o.EventDate = sv
		} else if sp, ok := val.(*IntegrationDataResponseEventDate); ok {
			// struct pointer
			o.EventDate = *sp
		}
	} else {
		o.EventDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["free_space"].(int64); ok {
		o.FreeSpace = val
	} else {
		if val, ok := kv["free_space"]; ok {
			if val == nil {
				o.FreeSpace = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.FreeSpace = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["go_version"].(string); ok {
		o.GoVersion = val
	} else {
		if val, ok := kv["go_version"]; ok {
			if val == nil {
				o.GoVersion = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.GoVersion = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["hostname"].(string); ok {
		o.Hostname = val
	} else {
		if val, ok := kv["hostname"]; ok {
			if val == nil {
				o.Hostname = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Hostname = fmt.Sprintf("%v", val)
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

	if val, ok := kv["job_id"].(string); ok {
		o.JobID = val
	} else {
		if val, ok := kv["job_id"]; ok {
			if val == nil {
				o.JobID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.JobID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["last_export_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastExportDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationDataResponseLastExportDate); ok {
			// struct
			o.LastExportDate = sv
		} else if sp, ok := val.(*IntegrationDataResponseLastExportDate); ok {
			// struct pointer
			o.LastExportDate = *sp
		}
	} else {
		o.LastExportDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["memory"].(int64); ok {
		o.Memory = val
	} else {
		if val, ok := kv["memory"]; ok {
			if val == nil {
				o.Memory = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Memory = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["message"].(string); ok {
		o.Message = val
	} else {
		if val, ok := kv["message"]; ok {
			if val == nil {
				o.Message = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Message = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["num_cpu"].(int64); ok {
		o.NumCPU = val
	} else {
		if val, ok := kv["num_cpu"]; ok {
			if val == nil {
				o.NumCPU = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.NumCPU = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["os"].(string); ok {
		o.OS = val
	} else {
		if val, ok := kv["os"]; ok {
			if val == nil {
				o.OS = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.OS = fmt.Sprintf("%v", val)
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

	if val, ok := kv["request_id"].(string); ok {
		o.RequestID = val
	} else {
		if val, ok := kv["request_id"]; ok {
			if val == nil {
				o.RequestID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.RequestID = fmt.Sprintf("%v", val)
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

	if val, ok := kv["system_id"].(string); ok {
		o.SystemID = val
	} else {
		if val, ok := kv["system_id"]; ok {
			if val == nil {
				o.SystemID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.SystemID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["type"].(IntegrationDataResponseType); ok {
		o.Type = val
	} else {
		if em, ok := kv["type"].(map[string]interface{}); ok {
			ev := em["agent.type"].(string)
			switch ev {
			case "enroll", "ENROLL":
				o.Type = 0
			case "ping", "PING":
				o.Type = 1
			case "crash", "CRASH":
				o.Type = 2
			case "log", "LOG":
				o.Type = 3
			case "integration", "INTEGRATION":
				o.Type = 4
			case "export", "EXPORT":
				o.Type = 5
			case "project", "PROJECT":
				o.Type = 6
			case "repo", "REPO":
				o.Type = 7
			case "user", "USER":
				o.Type = 8
			case "uninstall", "UNINSTALL":
				o.Type = 9
			case "upgrade", "UPGRADE":
				o.Type = 10
			case "start", "START":
				o.Type = 11
			case "stop", "STOP":
				o.Type = 12
			case "pause", "PAUSE":
				o.Type = 13
			case "resume", "RESUME":
				o.Type = 14
			}
		}
		if em, ok := kv["type"].(string); ok {
			switch em {
			case "enroll", "ENROLL":
				o.Type = 0
			case "ping", "PING":
				o.Type = 1
			case "crash", "CRASH":
				o.Type = 2
			case "log", "LOG":
				o.Type = 3
			case "integration", "INTEGRATION":
				o.Type = 4
			case "export", "EXPORT":
				o.Type = 5
			case "project", "PROJECT":
				o.Type = 6
			case "repo", "REPO":
				o.Type = 7
			case "user", "USER":
				o.Type = 8
			case "uninstall", "UNINSTALL":
				o.Type = 9
			case "upgrade", "UPGRADE":
				o.Type = 10
			case "start", "START":
				o.Type = 11
			case "stop", "STOP":
				o.Type = 12
			case "pause", "PAUSE":
				o.Type = 13
			case "resume", "RESUME":
				o.Type = 14
			}
		}
	}

	if val, ok := kv["uptime"].(int64); ok {
		o.Uptime = val
	} else {
		if val, ok := kv["uptime"]; ok {
			if val == nil {
				o.Uptime = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Uptime = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["uuid"].(string); ok {
		o.UUID = val
	} else {
		if val, ok := kv["uuid"]; ok {
			if val == nil {
				o.UUID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.UUID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["version"].(string); ok {
		o.Version = val
	} else {
		if val, ok := kv["version"]; ok {
			if val == nil {
				o.Version = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Version = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *IntegrationDataResponse) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Architecture)
	args = append(args, o.CustomerID)
	args = append(args, o.Data)
	args = append(args, o.Distro)
	args = append(args, o.Error)
	args = append(args, o.EventDate)
	args = append(args, o.FreeSpace)
	args = append(args, o.GoVersion)
	args = append(args, o.Hostname)
	args = append(args, o.ID)
	args = append(args, o.JobID)
	args = append(args, o.LastExportDate)
	args = append(args, o.Memory)
	args = append(args, o.Message)
	args = append(args, o.NumCPU)
	args = append(args, o.OS)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RequestID)
	args = append(args, o.Success)
	args = append(args, o.SystemID)
	args = append(args, o.Type)
	args = append(args, o.Uptime)
	args = append(args, o.UUID)
	args = append(args, o.Version)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *IntegrationDataResponse) GetEventAPIConfig() datamodel.EventAPIConfig {
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