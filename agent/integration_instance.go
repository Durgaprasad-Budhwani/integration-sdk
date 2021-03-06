// DO NOT EDIT -- generated code

// Package agent - the agent communicates with pinpoint cloud to send integration data
package agent

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/v10/datamodel"
	"github.com/pinpt/go-common/v10/datetime"
	"github.com/pinpt/go-common/v10/graphql"
	"github.com/pinpt/go-common/v10/hash"
	pjson "github.com/pinpt/go-common/v10/json"
	"github.com/pinpt/go-common/v10/number"
	pstrings "github.com/pinpt/go-common/v10/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (

	// IntegrationInstanceTable is the default table name
	IntegrationInstanceTable datamodel.ModelNameType = "agent_integrationinstance"

	// IntegrationInstanceModelName is the model name
	IntegrationInstanceModelName datamodel.ModelNameType = "agent.IntegrationInstance"
)

const (
	// IntegrationInstanceModelActiveColumn is the column json value active
	IntegrationInstanceModelActiveColumn = "active"
	// IntegrationInstanceModelAutoConfigureColumn is the column json value auto_configure
	IntegrationInstanceModelAutoConfigureColumn = "auto_configure"
	// IntegrationInstanceModelConfigColumn is the column json value config
	IntegrationInstanceModelConfigColumn = "config"
	// IntegrationInstanceModelCreatedByProfileIDColumn is the column json value created_by_profile_id
	IntegrationInstanceModelCreatedByProfileIDColumn = "created_by_profile_id"
	// IntegrationInstanceModelCreatedByUserIDColumn is the column json value created_by_user_id
	IntegrationInstanceModelCreatedByUserIDColumn = "created_by_user_id"
	// IntegrationInstanceModelCreatedDateColumn is the column json value created_date
	IntegrationInstanceModelCreatedDateColumn = "created_date"
	// IntegrationInstanceModelCreatedDateEpochColumn is the column json value epoch
	IntegrationInstanceModelCreatedDateEpochColumn = "epoch"
	// IntegrationInstanceModelCreatedDateOffsetColumn is the column json value offset
	IntegrationInstanceModelCreatedDateOffsetColumn = "offset"
	// IntegrationInstanceModelCreatedDateRfc3339Column is the column json value rfc3339
	IntegrationInstanceModelCreatedDateRfc3339Column = "rfc3339"
	// IntegrationInstanceModelCreatedAtColumn is the column json value created_ts
	IntegrationInstanceModelCreatedAtColumn = "created_ts"
	// IntegrationInstanceModelCustomerIDColumn is the column json value customer_id
	IntegrationInstanceModelCustomerIDColumn = "customer_id"
	// IntegrationInstanceModelDeletedColumn is the column json value deleted
	IntegrationInstanceModelDeletedColumn = "deleted"
	// IntegrationInstanceModelDeletedByProfileIDColumn is the column json value deleted_by_profile_id
	IntegrationInstanceModelDeletedByProfileIDColumn = "deleted_by_profile_id"
	// IntegrationInstanceModelDeletedByUserIDColumn is the column json value deleted_by_user_id
	IntegrationInstanceModelDeletedByUserIDColumn = "deleted_by_user_id"
	// IntegrationInstanceModelDeletedDateColumn is the column json value deleted_date
	IntegrationInstanceModelDeletedDateColumn = "deleted_date"
	// IntegrationInstanceModelDeletedDateEpochColumn is the column json value epoch
	IntegrationInstanceModelDeletedDateEpochColumn = "epoch"
	// IntegrationInstanceModelDeletedDateOffsetColumn is the column json value offset
	IntegrationInstanceModelDeletedDateOffsetColumn = "offset"
	// IntegrationInstanceModelDeletedDateRfc3339Column is the column json value rfc3339
	IntegrationInstanceModelDeletedDateRfc3339Column = "rfc3339"
	// IntegrationInstanceModelEnrollmentIDColumn is the column json value enrollment_id
	IntegrationInstanceModelEnrollmentIDColumn = "enrollment_id"
	// IntegrationInstanceModelErrorDateColumn is the column json value error_date
	IntegrationInstanceModelErrorDateColumn = "error_date"
	// IntegrationInstanceModelErrorDateEpochColumn is the column json value epoch
	IntegrationInstanceModelErrorDateEpochColumn = "epoch"
	// IntegrationInstanceModelErrorDateOffsetColumn is the column json value offset
	IntegrationInstanceModelErrorDateOffsetColumn = "offset"
	// IntegrationInstanceModelErrorDateRfc3339Column is the column json value rfc3339
	IntegrationInstanceModelErrorDateRfc3339Column = "rfc3339"
	// IntegrationInstanceModelErrorMessageColumn is the column json value error_message
	IntegrationInstanceModelErrorMessageColumn = "error_message"
	// IntegrationInstanceModelErroredColumn is the column json value errored
	IntegrationInstanceModelErroredColumn = "errored"
	// IntegrationInstanceModelExportAcknowledgedColumn is the column json value export_acknowledged
	IntegrationInstanceModelExportAcknowledgedColumn = "export_acknowledged"
	// IntegrationInstanceModelExportLivenessColumn is the column json value export_liveness
	IntegrationInstanceModelExportLivenessColumn = "export_liveness"
	// IntegrationInstanceModelIDColumn is the column json value id
	IntegrationInstanceModelIDColumn = "id"
	// IntegrationInstanceModelIntegrationIDColumn is the column json value integration_id
	IntegrationInstanceModelIntegrationIDColumn = "integration_id"
	// IntegrationInstanceModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	IntegrationInstanceModelIntegrationInstanceIDColumn = "integration_instance_id"
	// IntegrationInstanceModelIntervalColumn is the column json value interval
	IntegrationInstanceModelIntervalColumn = "interval"
	// IntegrationInstanceModelJobIDColumn is the column json value job_id
	IntegrationInstanceModelJobIDColumn = "job_id"
	// IntegrationInstanceModelLocationColumn is the column json value location
	IntegrationInstanceModelLocationColumn = "location"
	// IntegrationInstanceModelNameColumn is the column json value name
	IntegrationInstanceModelNameColumn = "name"
	// IntegrationInstanceModelPausedColumn is the column json value paused
	IntegrationInstanceModelPausedColumn = "paused"
	// IntegrationInstanceModelPrivateKeyColumn is the column json value private_key
	IntegrationInstanceModelPrivateKeyColumn = "private_key"
	// IntegrationInstanceModelProcessedColumn is the column json value processed
	IntegrationInstanceModelProcessedColumn = "processed"
	// IntegrationInstanceModelRefIDColumn is the column json value ref_id
	IntegrationInstanceModelRefIDColumn = "ref_id"
	// IntegrationInstanceModelRefTypeColumn is the column json value ref_type
	IntegrationInstanceModelRefTypeColumn = "ref_type"
	// IntegrationInstanceModelRequiresHistoricalColumn is the column json value requires_historical
	IntegrationInstanceModelRequiresHistoricalColumn = "requires_historical"
	// IntegrationInstanceModelSetupColumn is the column json value setup
	IntegrationInstanceModelSetupColumn = "setup"
	// IntegrationInstanceModelStateColumn is the column json value state
	IntegrationInstanceModelStateColumn = "state"
	// IntegrationInstanceModelThrottledColumn is the column json value throttled
	IntegrationInstanceModelThrottledColumn = "throttled"
	// IntegrationInstanceModelThrottledUntilColumn is the column json value throttled_until
	IntegrationInstanceModelThrottledUntilColumn = "throttled_until"
	// IntegrationInstanceModelThrottledUntilEpochColumn is the column json value epoch
	IntegrationInstanceModelThrottledUntilEpochColumn = "epoch"
	// IntegrationInstanceModelThrottledUntilOffsetColumn is the column json value offset
	IntegrationInstanceModelThrottledUntilOffsetColumn = "offset"
	// IntegrationInstanceModelThrottledUntilRfc3339Column is the column json value rfc3339
	IntegrationInstanceModelThrottledUntilRfc3339Column = "rfc3339"
	// IntegrationInstanceModelUpdatedAtColumn is the column json value updated_ts
	IntegrationInstanceModelUpdatedAtColumn = "updated_ts"
	// IntegrationInstanceModelUpgradeDateAtColumn is the column json value upgrade_date_ts
	IntegrationInstanceModelUpgradeDateAtColumn = "upgrade_date_ts"
	// IntegrationInstanceModelUpgradeExpiresDateAtColumn is the column json value upgrade_expires_date_ts
	IntegrationInstanceModelUpgradeExpiresDateAtColumn = "upgrade_expires_date_ts"
	// IntegrationInstanceModelUpgradeMessageColumn is the column json value upgrade_message
	IntegrationInstanceModelUpgradeMessageColumn = "upgrade_message"
	// IntegrationInstanceModelUpgradeRequiredColumn is the column json value upgrade_required
	IntegrationInstanceModelUpgradeRequiredColumn = "upgrade_required"
	// IntegrationInstanceModelWebhooksColumn is the column json value webhooks
	IntegrationInstanceModelWebhooksColumn = "webhooks"
	// IntegrationInstanceModelWebhooksEnabledColumn is the column json value enabled
	IntegrationInstanceModelWebhooksEnabledColumn = "enabled"
	// IntegrationInstanceModelWebhooksErrorMessageColumn is the column json value error_message
	IntegrationInstanceModelWebhooksErrorMessageColumn = "error_message"
	// IntegrationInstanceModelWebhooksErroredColumn is the column json value errored
	IntegrationInstanceModelWebhooksErroredColumn = "errored"
	// IntegrationInstanceModelWebhooksRefIDColumn is the column json value ref_id
	IntegrationInstanceModelWebhooksRefIDColumn = "ref_id"
	// IntegrationInstanceModelWebhooksURLColumn is the column json value url
	IntegrationInstanceModelWebhooksURLColumn = "url"
)

// IntegrationInstanceCreatedDate represents the object structure for created_date
type IntegrationInstanceCreatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationInstanceCreatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationInstanceCreatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationInstanceCreatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationInstanceCreatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationInstanceCreatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationInstanceCreatedDateObject(o.Rfc3339, false),
	}
}

func (o *IntegrationInstanceCreatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationInstanceCreatedDate) FromMap(kv map[string]interface{}) {

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

// IntegrationInstanceDeletedDate represents the object structure for deleted_date
type IntegrationInstanceDeletedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationInstanceDeletedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationInstanceDeletedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationInstanceDeletedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationInstanceDeletedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationInstanceDeletedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationInstanceDeletedDateObject(o.Rfc3339, false),
	}
}

func (o *IntegrationInstanceDeletedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationInstanceDeletedDate) FromMap(kv map[string]interface{}) {

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

// IntegrationInstanceErrorDate represents the object structure for error_date
type IntegrationInstanceErrorDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationInstanceErrorDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationInstanceErrorDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationInstanceErrorDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationInstanceErrorDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationInstanceErrorDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationInstanceErrorDateObject(o.Rfc3339, false),
	}
}

func (o *IntegrationInstanceErrorDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationInstanceErrorDate) FromMap(kv map[string]interface{}) {

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

// IntegrationInstanceLocation is the enumeration type for location
type IntegrationInstanceLocation int32

// toIntegrationInstanceLocationPointer is the enumeration pointer type for location
func toIntegrationInstanceLocationPointer(v int32) *IntegrationInstanceLocation {
	nv := IntegrationInstanceLocation(v)
	return &nv
}

// toIntegrationInstanceLocationEnum is the enumeration pointer wrapper for location
func toIntegrationInstanceLocationEnum(v *IntegrationInstanceLocation) string {
	if v == nil {
		return toIntegrationInstanceLocationPointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *IntegrationInstanceLocation) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = IntegrationInstanceLocation(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "PRIVATE":
			*v = IntegrationInstanceLocation(0)
		case "CLOUD":
			*v = IntegrationInstanceLocation(1)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v *IntegrationInstanceLocation) UnmarshalJSON(buf []byte) error {
	var val string
	if err := json.Unmarshal(buf, &val); err != nil {
		return err
	}
	switch val {
	case "PRIVATE":
		*v = 0
	case "CLOUD":
		*v = 1
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v IntegrationInstanceLocation) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("PRIVATE")
	case 1:
		return json.Marshal("CLOUD")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for Location
func (v IntegrationInstanceLocation) String() string {
	switch int32(v) {
	case 0:
		return "PRIVATE"
	case 1:
		return "CLOUD"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *IntegrationInstanceLocation) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case IntegrationInstanceLocation:
		*v = val
	case int32:
		*v = IntegrationInstanceLocation(int32(val))
	case int:
		*v = IntegrationInstanceLocation(int32(val))
	case string:
		switch val {
		case "PRIVATE":
			*v = IntegrationInstanceLocation(0)
		case "CLOUD":
			*v = IntegrationInstanceLocation(1)
		}
	}
	return nil
}

const (
	// IntegrationInstanceLocationPrivate is the enumeration value for private
	IntegrationInstanceLocationPrivate IntegrationInstanceLocation = 0
	// IntegrationInstanceLocationCloud is the enumeration value for cloud
	IntegrationInstanceLocationCloud IntegrationInstanceLocation = 1
)

// IntegrationInstanceSetup is the enumeration type for setup
type IntegrationInstanceSetup int32

// toIntegrationInstanceSetupPointer is the enumeration pointer type for setup
func toIntegrationInstanceSetupPointer(v int32) *IntegrationInstanceSetup {
	nv := IntegrationInstanceSetup(v)
	return &nv
}

// toIntegrationInstanceSetupEnum is the enumeration pointer wrapper for setup
func toIntegrationInstanceSetupEnum(v *IntegrationInstanceSetup) string {
	if v == nil {
		return toIntegrationInstanceSetupPointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *IntegrationInstanceSetup) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = IntegrationInstanceSetup(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "CONFIG":
			*v = IntegrationInstanceSetup(0)
		case "READY":
			*v = IntegrationInstanceSetup(1)
		case "RUNNING":
			*v = IntegrationInstanceSetup(2)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v *IntegrationInstanceSetup) UnmarshalJSON(buf []byte) error {
	var val string
	if err := json.Unmarshal(buf, &val); err != nil {
		return err
	}
	switch val {
	case "CONFIG":
		*v = 0
	case "READY":
		*v = 1
	case "RUNNING":
		*v = 2
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v IntegrationInstanceSetup) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("CONFIG")
	case 1:
		return json.Marshal("READY")
	case 2:
		return json.Marshal("RUNNING")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for Setup
func (v IntegrationInstanceSetup) String() string {
	switch int32(v) {
	case 0:
		return "CONFIG"
	case 1:
		return "READY"
	case 2:
		return "RUNNING"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *IntegrationInstanceSetup) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case IntegrationInstanceSetup:
		*v = val
	case int32:
		*v = IntegrationInstanceSetup(int32(val))
	case int:
		*v = IntegrationInstanceSetup(int32(val))
	case string:
		switch val {
		case "CONFIG":
			*v = IntegrationInstanceSetup(0)
		case "READY":
			*v = IntegrationInstanceSetup(1)
		case "RUNNING":
			*v = IntegrationInstanceSetup(2)
		}
	}
	return nil
}

const (
	// IntegrationInstanceSetupConfig is the enumeration value for config
	IntegrationInstanceSetupConfig IntegrationInstanceSetup = 0
	// IntegrationInstanceSetupReady is the enumeration value for ready
	IntegrationInstanceSetupReady IntegrationInstanceSetup = 1
	// IntegrationInstanceSetupRunning is the enumeration value for running
	IntegrationInstanceSetupRunning IntegrationInstanceSetup = 2
)

// IntegrationInstanceState is the enumeration type for state
type IntegrationInstanceState int32

// toIntegrationInstanceStatePointer is the enumeration pointer type for state
func toIntegrationInstanceStatePointer(v int32) *IntegrationInstanceState {
	nv := IntegrationInstanceState(v)
	return &nv
}

// toIntegrationInstanceStateEnum is the enumeration pointer wrapper for state
func toIntegrationInstanceStateEnum(v *IntegrationInstanceState) string {
	if v == nil {
		return toIntegrationInstanceStatePointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *IntegrationInstanceState) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = IntegrationInstanceState(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "IDLE":
			*v = IntegrationInstanceState(0)
		case "EXPORTING":
			*v = IntegrationInstanceState(1)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v *IntegrationInstanceState) UnmarshalJSON(buf []byte) error {
	var val string
	if err := json.Unmarshal(buf, &val); err != nil {
		return err
	}
	switch val {
	case "IDLE":
		*v = 0
	case "EXPORTING":
		*v = 1
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v IntegrationInstanceState) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("IDLE")
	case 1:
		return json.Marshal("EXPORTING")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for State
func (v IntegrationInstanceState) String() string {
	switch int32(v) {
	case 0:
		return "IDLE"
	case 1:
		return "EXPORTING"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *IntegrationInstanceState) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case IntegrationInstanceState:
		*v = val
	case int32:
		*v = IntegrationInstanceState(int32(val))
	case int:
		*v = IntegrationInstanceState(int32(val))
	case string:
		switch val {
		case "IDLE":
			*v = IntegrationInstanceState(0)
		case "EXPORTING":
			*v = IntegrationInstanceState(1)
		}
	}
	return nil
}

const (
	// IntegrationInstanceStateIdle is the enumeration value for idle
	IntegrationInstanceStateIdle IntegrationInstanceState = 0
	// IntegrationInstanceStateExporting is the enumeration value for exporting
	IntegrationInstanceStateExporting IntegrationInstanceState = 1
)

// IntegrationInstanceThrottledUntil represents the object structure for throttled_until
type IntegrationInstanceThrottledUntil struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationInstanceThrottledUntilObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationInstanceThrottledUntil:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationInstanceThrottledUntil) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationInstanceThrottledUntilObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationInstanceThrottledUntilObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationInstanceThrottledUntilObject(o.Rfc3339, false),
	}
}

func (o *IntegrationInstanceThrottledUntil) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationInstanceThrottledUntil) FromMap(kv map[string]interface{}) {

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

// IntegrationInstanceWebhooks represents the object structure for webhooks
type IntegrationInstanceWebhooks struct {
	// Enabled if webhooks are enabled for this instance
	Enabled bool `json:"enabled" codec:"enabled" bson:"enabled" yaml:"enabled" faker:"-"`
	// ErrorMessage the error message
	ErrorMessage *string `json:"error_message,omitempty" codec:"error_message,omitempty" bson:"error_message" yaml:"error_message,omitempty" faker:"-"`
	// Errored if the webhook has an error
	Errored bool `json:"errored" codec:"errored" bson:"errored" yaml:"errored" faker:"-"`
	// RefID an optional reference id related to the webhook and what its related to
	RefID *string `json:"ref_id,omitempty" codec:"ref_id,omitempty" bson:"ref_id" yaml:"ref_id,omitempty" faker:"-"`
	// URL the url the webhook for the webhook
	URL *string `json:"url,omitempty" codec:"url,omitempty" bson:"url" yaml:"url,omitempty" faker:"-"`
}

func toIntegrationInstanceWebhooksObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationInstanceWebhooks:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationInstanceWebhooks) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Enabled if webhooks are enabled for this instance
		"enabled": toIntegrationInstanceWebhooksObject(o.Enabled, false),
		// ErrorMessage the error message
		"error_message": toIntegrationInstanceWebhooksObject(o.ErrorMessage, true),
		// Errored if the webhook has an error
		"errored": toIntegrationInstanceWebhooksObject(o.Errored, false),
		// RefID an optional reference id related to the webhook and what its related to
		"ref_id": toIntegrationInstanceWebhooksObject(o.RefID, true),
		// URL the url the webhook for the webhook
		"url": toIntegrationInstanceWebhooksObject(o.URL, true),
	}
}

func (o *IntegrationInstanceWebhooks) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationInstanceWebhooks) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["enabled"].(bool); ok {
		o.Enabled = val
	} else {
		if val, ok := kv["enabled"]; ok {
			if val == nil {
				o.Enabled = false
			} else {
				o.Enabled = number.ToBoolAny(val)
			}
		}
	}
	if val, ok := kv["error_message"].(*string); ok {
		o.ErrorMessage = val
	} else if val, ok := kv["error_message"].(string); ok {
		o.ErrorMessage = &val
	} else {
		if val, ok := kv["error_message"]; ok {
			if val == nil {
				o.ErrorMessage = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.ErrorMessage = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["errored"].(bool); ok {
		o.Errored = val
	} else {
		if val, ok := kv["errored"]; ok {
			if val == nil {
				o.Errored = false
			} else {
				o.Errored = number.ToBoolAny(val)
			}
		}
	}
	if val, ok := kv["ref_id"].(*string); ok {
		o.RefID = val
	} else if val, ok := kv["ref_id"].(string); ok {
		o.RefID = &val
	} else {
		if val, ok := kv["ref_id"]; ok {
			if val == nil {
				o.RefID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.RefID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["url"].(*string); ok {
		o.URL = val
	} else if val, ok := kv["url"].(string); ok {
		o.URL = &val
	} else {
		if val, ok := kv["url"]; ok {
			if val == nil {
				o.URL = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.URL = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	o.setDefaults(false)
}

// IntegrationInstance The integration instance for a customer
type IntegrationInstance struct {
	// Active If true, the integration is still active
	Active bool `json:"active" codec:"active" bson:"active" yaml:"active" faker:"-"`
	// AutoConfigure if this integration should be auto configured by the integration before activating
	AutoConfigure bool `json:"auto_configure" codec:"auto_configure" bson:"auto_configure" yaml:"auto_configure" faker:"-"`
	// Config the integration configuration controlled by the integration itself
	Config *string `json:"config,omitempty" codec:"config,omitempty" bson:"config" yaml:"config,omitempty" faker:"-"`
	// CreatedByProfileID The id of the profile for the user that created the integration
	CreatedByProfileID *string `json:"created_by_profile_id,omitempty" codec:"created_by_profile_id,omitempty" bson:"created_by_profile_id" yaml:"created_by_profile_id,omitempty" faker:"-"`
	// CreatedByUserID The id of the user that created the integration
	CreatedByUserID *string `json:"created_by_user_id,omitempty" codec:"created_by_user_id,omitempty" bson:"created_by_user_id" yaml:"created_by_user_id,omitempty" faker:"-"`
	// CreatedDate when the integration was created
	CreatedDate IntegrationInstanceCreatedDate `json:"created_date" codec:"created_date" bson:"created_date" yaml:"created_date" faker:"-"`
	// CreatedAt the date the record was created in Epoch time
	CreatedAt int64 `json:"created_ts" codec:"created_ts" bson:"created_ts" yaml:"created_ts" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Deleted If true, the integration has been deleted
	Deleted bool `json:"deleted" codec:"deleted" bson:"deleted" yaml:"deleted" faker:"-"`
	// DeletedByProfileID The id of the profile for the user that deleted the integration
	DeletedByProfileID *string `json:"deleted_by_profile_id,omitempty" codec:"deleted_by_profile_id,omitempty" bson:"deleted_by_profile_id" yaml:"deleted_by_profile_id,omitempty" faker:"-"`
	// DeletedByUserID The id of the user that deleted the integration
	DeletedByUserID *string `json:"deleted_by_user_id,omitempty" codec:"deleted_by_user_id,omitempty" bson:"deleted_by_user_id" yaml:"deleted_by_user_id,omitempty" faker:"-"`
	// DeletedDate when the integration was deleted
	DeletedDate IntegrationInstanceDeletedDate `json:"deleted_date" codec:"deleted_date" bson:"deleted_date" yaml:"deleted_date" faker:"-"`
	// EnrollmentID if the integration is linked to a self-managed agent, it will have the enrollment_id set otherwise will be null
	EnrollmentID *string `json:"enrollment_id,omitempty" codec:"enrollment_id,omitempty" bson:"enrollment_id" yaml:"enrollment_id,omitempty" faker:"-"`
	// ErrorDate The date of the error from an export run
	ErrorDate *IntegrationInstanceErrorDate `json:"error_date,omitempty" codec:"error_date,omitempty" bson:"error_date" yaml:"error_date,omitempty" faker:"-"`
	// ErrorMessage The error message from an export run
	ErrorMessage *string `json:"error_message,omitempty" codec:"error_message,omitempty" bson:"error_message" yaml:"error_message,omitempty" faker:"-"`
	// Errored If authorization failed by the agent or any other error
	Errored *bool `json:"errored,omitempty" codec:"errored,omitempty" bson:"errored" yaml:"errored,omitempty" faker:"-"`
	// ExportAcknowledged Set to true an export has been received by the agent.
	//
	// Deprecated: no longer used, see export_liveness
	ExportAcknowledged *bool `json:"export_acknowledged,omitempty" codec:"export_acknowledged,omitempty" bson:"export_acknowledged" yaml:"export_acknowledged,omitempty" faker:"-"`
	// ExportLiveness Is true when agent is still exporting.
	ExportLiveness *bool `json:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationID The unique id for the integration
	IntegrationID string `json:"integration_id" codec:"integration_id" bson:"integration_id" yaml:"integration_id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// Interval the interval in milliseconds for how often an export job is scheduled
	Interval int64 `json:"interval" codec:"interval" bson:"interval" yaml:"interval" faker:"-"`
	// JobID the current (if EXPORTING) or previous (if IDLE) job id which is useful for debugging
	JobID *string `json:"job_id,omitempty" codec:"job_id,omitempty" bson:"job_id" yaml:"job_id,omitempty" faker:"-"`
	// Location The location of this integration (on-premise / private or cloud)
	Location IntegrationInstanceLocation `json:"location" codec:"location" bson:"location" yaml:"location" faker:"-"`
	// Name The user friendly name of the integration
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// Paused true if the agent is paused and should not start new scheduled jobs
	Paused bool `json:"paused" codec:"paused" bson:"paused" yaml:"paused" faker:"-"`
	// PrivateKey the private key for the instance if needed by an integration
	PrivateKey *string `json:"private_key,omitempty" codec:"private_key,omitempty" bson:"private_key" yaml:"private_key,omitempty" faker:"-"`
	// Processed If the integration has been processed at least once
	//
	// Deprecated: this field should not be used anymore. use stat to get details about processing
	Processed *bool `json:"processed,omitempty" codec:"processed,omitempty" bson:"processed" yaml:"processed,omitempty" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RequiresHistorical flag which can be set to trigger a historical on the next scheduler visit
	RequiresHistorical bool `json:"requires_historical" codec:"requires_historical" bson:"requires_historical" yaml:"requires_historical" faker:"-"`
	// Setup the setup state of the integration
	Setup IntegrationInstanceSetup `json:"setup" codec:"setup" bson:"setup" yaml:"setup" faker:"-"`
	// State the current state of the integration
	State IntegrationInstanceState `json:"state" codec:"state" bson:"state" yaml:"state" faker:"-"`
	// Throttled Set to true when integration is throttled.
	Throttled *bool `json:"throttled,omitempty" codec:"throttled,omitempty" bson:"throttled" yaml:"throttled,omitempty" faker:"-"`
	// ThrottledUntil After throttling integration, we set this field for estimated resume date.
	ThrottledUntil *IntegrationInstanceThrottledUntil `json:"throttled_until,omitempty" codec:"throttled_until,omitempty" bson:"throttled_until" yaml:"throttled_until,omitempty" faker:"-"`
	// UpdatedAt the date the record was updated in Epoch time
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// UpgradeDateAt If upgrade is required, the date when the upgrade was set
	UpgradeDateAt *int64 `json:"upgrade_date_ts,omitempty" codec:"upgrade_date_ts,omitempty" bson:"upgrade_date_ts" yaml:"upgrade_date_ts,omitempty" faker:"-"`
	// UpgradeExpiresDateAt If upgrade is required and there's a due date this will be set
	UpgradeExpiresDateAt *int64 `json:"upgrade_expires_date_ts,omitempty" codec:"upgrade_expires_date_ts,omitempty" bson:"upgrade_expires_date_ts" yaml:"upgrade_expires_date_ts,omitempty" faker:"-"`
	// UpgradeMessage If upgrade is required, the message to display to the user
	UpgradeMessage *string `json:"upgrade_message,omitempty" codec:"upgrade_message,omitempty" bson:"upgrade_message" yaml:"upgrade_message,omitempty" faker:"-"`
	// UpgradeRequired If true, the integration requires a manual upgrade
	UpgradeRequired bool `json:"upgrade_required" codec:"upgrade_required" bson:"upgrade_required" yaml:"upgrade_required" faker:"-"`
	// Webhooks for any webhooks installed at the integration instance
	Webhooks []IntegrationInstanceWebhooks `json:"webhooks" codec:"webhooks" bson:"webhooks" yaml:"webhooks" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*IntegrationInstance)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*IntegrationInstance)(nil)

func toIntegrationInstanceObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationInstance:
		return v.ToMap()

	case IntegrationInstanceCreatedDate:
		return v.ToMap()

	case IntegrationInstanceDeletedDate:
		return v.ToMap()

	case *IntegrationInstanceErrorDate:
		return v.ToMap()

	case IntegrationInstanceLocation:
		return v.String()

	case IntegrationInstanceSetup:
		return v.String()

	case IntegrationInstanceState:
		return v.String()

	case *IntegrationInstanceThrottledUntil:
		return v.ToMap()

	case []IntegrationInstanceWebhooks:
		arr := make([]interface{}, 0)
		for _, i := range v {
			arr = append(arr, i.ToMap())
		}
		return arr

	default:
		return o
	}
}

// String returns a string representation of IntegrationInstance
func (o *IntegrationInstance) String() string {
	return fmt.Sprintf("agent.IntegrationInstance<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *IntegrationInstance) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *IntegrationInstance) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *IntegrationInstance) GetTableName() string {
	return IntegrationInstanceTable.String()
}

// GetModelName returns the name of the model
func (o *IntegrationInstance) GetModelName() datamodel.ModelNameType {
	return IntegrationInstanceModelName
}

// NewIntegrationInstanceID provides a template for generating an ID field for IntegrationInstance
func NewIntegrationInstanceID(customerID string) string {
	return hash.Values(customerID, datetime.EpochNow(), randomString(64))
}

func (o *IntegrationInstance) setDefaults(frommap bool) {
	if o.ErrorDate == nil {
		o.ErrorDate = &IntegrationInstanceErrorDate{}
	}
	if o.ThrottledUntil == nil {
		o.ThrottledUntil = &IntegrationInstanceThrottledUntil{}
	}
	if o.Webhooks == nil {
		o.Webhooks = make([]IntegrationInstanceWebhooks, 0)
	}

	if o.ID == "" {
		// set the id from the spec provided in the model
		o.ID = hash.Values(o.CustomerID, datetime.EpochNow(), randomString(64))
	}

	if o.Errored == nil {
		var v bool
		o.Errored = &v
	}

	if o.ExportAcknowledged == nil {
		var v bool
		o.ExportAcknowledged = &v
	}

	if o.ExportLiveness == nil {
		var v bool
		o.ExportLiveness = &v
	}

	{

	}

	if o.Processed == nil {
		var v bool
		o.Processed = &v
	}

	if o.Throttled == nil {
		var v bool
		o.Throttled = &v
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *IntegrationInstance) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *IntegrationInstance) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *IntegrationInstance) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *IntegrationInstance) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *IntegrationInstance) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *IntegrationInstance) IsMutable() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *IntegrationInstance) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *IntegrationInstance) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *IntegrationInstance) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = IntegrationInstanceModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *IntegrationInstance) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *IntegrationInstance) GetCustomerID() string {
	return o.CustomerID
}

// SetCustomerID will return the customer_id
func (o *IntegrationInstance) SetCustomerID(id string) {
	o.CustomerID = id
}

// GetRefType will return the ref_type
func (o *IntegrationInstance) GetRefType() string {
	return o.RefType
}

// SetRefType will return the ref_type
func (o *IntegrationInstance) SetRefType(t string) {
	o.RefType = t
}

// Clone returns an exact copy of IntegrationInstance
func (o *IntegrationInstance) Clone() datamodel.Model {
	c := new(IntegrationInstance)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *IntegrationInstance) Anon() datamodel.Model {
	c := new(IntegrationInstance)
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
func (o *IntegrationInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *IntegrationInstance) UnmarshalJSON(data []byte) error {
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
func (o *IntegrationInstance) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *IntegrationInstance) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two IntegrationInstance objects are equal
func (o *IntegrationInstance) IsEqual(other *IntegrationInstance) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *IntegrationInstance) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"active":                  toIntegrationInstanceObject(o.Active, false),
		"auto_configure":          toIntegrationInstanceObject(o.AutoConfigure, false),
		"config":                  toIntegrationInstanceObject(o.Config, true),
		"created_by_profile_id":   toIntegrationInstanceObject(o.CreatedByProfileID, true),
		"created_by_user_id":      toIntegrationInstanceObject(o.CreatedByUserID, true),
		"created_date":            toIntegrationInstanceObject(o.CreatedDate, false),
		"created_ts":              toIntegrationInstanceObject(o.CreatedAt, false),
		"customer_id":             toIntegrationInstanceObject(o.CustomerID, false),
		"deleted":                 toIntegrationInstanceObject(o.Deleted, false),
		"deleted_by_profile_id":   toIntegrationInstanceObject(o.DeletedByProfileID, true),
		"deleted_by_user_id":      toIntegrationInstanceObject(o.DeletedByUserID, true),
		"deleted_date":            toIntegrationInstanceObject(o.DeletedDate, false),
		"enrollment_id":           toIntegrationInstanceObject(o.EnrollmentID, true),
		"error_date":              toIntegrationInstanceObject(o.ErrorDate, true),
		"error_message":           toIntegrationInstanceObject(o.ErrorMessage, true),
		"errored":                 toIntegrationInstanceObject(o.Errored, true),
		"export_acknowledged":     toIntegrationInstanceObject(o.ExportAcknowledged, true),
		"export_liveness":         toIntegrationInstanceObject(o.ExportLiveness, true),
		"id":                      toIntegrationInstanceObject(o.ID, false),
		"integration_id":          toIntegrationInstanceObject(o.IntegrationID, false),
		"integration_instance_id": toIntegrationInstanceObject(o.IntegrationInstanceID, true),
		"interval":                toIntegrationInstanceObject(o.Interval, false),
		"job_id":                  toIntegrationInstanceObject(o.JobID, true),

		"location":            o.Location.String(),
		"name":                toIntegrationInstanceObject(o.Name, false),
		"paused":              toIntegrationInstanceObject(o.Paused, false),
		"private_key":         toIntegrationInstanceObject(o.PrivateKey, true),
		"processed":           toIntegrationInstanceObject(o.Processed, true),
		"ref_id":              toIntegrationInstanceObject(o.RefID, false),
		"ref_type":            toIntegrationInstanceObject(o.RefType, false),
		"requires_historical": toIntegrationInstanceObject(o.RequiresHistorical, false),

		"setup": o.Setup.String(),

		"state":                   o.State.String(),
		"throttled":               toIntegrationInstanceObject(o.Throttled, true),
		"throttled_until":         toIntegrationInstanceObject(o.ThrottledUntil, true),
		"updated_ts":              toIntegrationInstanceObject(o.UpdatedAt, false),
		"upgrade_date_ts":         toIntegrationInstanceObject(o.UpgradeDateAt, true),
		"upgrade_expires_date_ts": toIntegrationInstanceObject(o.UpgradeExpiresDateAt, true),
		"upgrade_message":         toIntegrationInstanceObject(o.UpgradeMessage, true),
		"upgrade_required":        toIntegrationInstanceObject(o.UpgradeRequired, false),
		"webhooks":                toIntegrationInstanceObject(o.Webhooks, false),
		"hashcode":                toIntegrationInstanceObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationInstance) FromMap(kv map[string]interface{}) {

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
	if val, ok := kv["auto_configure"].(bool); ok {
		o.AutoConfigure = val
	} else {
		if val, ok := kv["auto_configure"]; ok {
			if val == nil {
				o.AutoConfigure = false
			} else {
				o.AutoConfigure = number.ToBoolAny(val)
			}
		}
	}
	if val, ok := kv["config"].(*string); ok {
		o.Config = val
	} else if val, ok := kv["config"].(string); ok {
		o.Config = &val
	} else {
		if val, ok := kv["config"]; ok {
			if val == nil {
				o.Config = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Config = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["created_by_profile_id"].(*string); ok {
		o.CreatedByProfileID = val
	} else if val, ok := kv["created_by_profile_id"].(string); ok {
		o.CreatedByProfileID = &val
	} else {
		if val, ok := kv["created_by_profile_id"]; ok {
			if val == nil {
				o.CreatedByProfileID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.CreatedByProfileID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["created_by_user_id"].(*string); ok {
		o.CreatedByUserID = val
	} else if val, ok := kv["created_by_user_id"].(string); ok {
		o.CreatedByUserID = &val
	} else {
		if val, ok := kv["created_by_user_id"]; ok {
			if val == nil {
				o.CreatedByUserID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.CreatedByUserID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["created_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CreatedDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationInstanceCreatedDate); ok {
			// struct
			o.CreatedDate = sv
		} else if sp, ok := val.(*IntegrationInstanceCreatedDate); ok {
			// struct pointer
			o.CreatedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.CreatedDate.Epoch = dt.Epoch
			o.CreatedDate.Rfc3339 = dt.Rfc3339
			o.CreatedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.CreatedDate.Epoch = dt.Epoch
			o.CreatedDate.Rfc3339 = dt.Rfc3339
			o.CreatedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.CreatedDate.Epoch = dt.Epoch
				o.CreatedDate.Rfc3339 = dt.Rfc3339
				o.CreatedDate.Offset = dt.Offset
			}
		}
	} else {
		o.CreatedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["created_ts"].(int64); ok {
		o.CreatedAt = val
	} else {
		if val, ok := kv["created_ts"]; ok {
			if val == nil {
				o.CreatedAt = 0
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
	if val, ok := kv["deleted"].(bool); ok {
		o.Deleted = val
	} else {
		if val, ok := kv["deleted"]; ok {
			if val == nil {
				o.Deleted = false
			} else {
				o.Deleted = number.ToBoolAny(val)
			}
		}
	}
	if val, ok := kv["deleted_by_profile_id"].(*string); ok {
		o.DeletedByProfileID = val
	} else if val, ok := kv["deleted_by_profile_id"].(string); ok {
		o.DeletedByProfileID = &val
	} else {
		if val, ok := kv["deleted_by_profile_id"]; ok {
			if val == nil {
				o.DeletedByProfileID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.DeletedByProfileID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["deleted_by_user_id"].(*string); ok {
		o.DeletedByUserID = val
	} else if val, ok := kv["deleted_by_user_id"].(string); ok {
		o.DeletedByUserID = &val
	} else {
		if val, ok := kv["deleted_by_user_id"]; ok {
			if val == nil {
				o.DeletedByUserID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.DeletedByUserID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["deleted_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.DeletedDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationInstanceDeletedDate); ok {
			// struct
			o.DeletedDate = sv
		} else if sp, ok := val.(*IntegrationInstanceDeletedDate); ok {
			// struct pointer
			o.DeletedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.DeletedDate.Epoch = dt.Epoch
			o.DeletedDate.Rfc3339 = dt.Rfc3339
			o.DeletedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.DeletedDate.Epoch = dt.Epoch
			o.DeletedDate.Rfc3339 = dt.Rfc3339
			o.DeletedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.DeletedDate.Epoch = dt.Epoch
				o.DeletedDate.Rfc3339 = dt.Rfc3339
				o.DeletedDate.Offset = dt.Offset
			}
		}
	} else {
		o.DeletedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["enrollment_id"].(*string); ok {
		o.EnrollmentID = val
	} else if val, ok := kv["enrollment_id"].(string); ok {
		o.EnrollmentID = &val
	} else {
		if val, ok := kv["enrollment_id"]; ok {
			if val == nil {
				o.EnrollmentID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.EnrollmentID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if o.ErrorDate == nil {
		o.ErrorDate = &IntegrationInstanceErrorDate{}
	}

	if val, ok := kv["error_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.ErrorDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationInstanceErrorDate); ok {
			// struct
			o.ErrorDate = &sv
		} else if sp, ok := val.(*IntegrationInstanceErrorDate); ok {
			// struct pointer
			o.ErrorDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.ErrorDate.Epoch = dt.Epoch
			o.ErrorDate.Rfc3339 = dt.Rfc3339
			o.ErrorDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.ErrorDate.Epoch = dt.Epoch
			o.ErrorDate.Rfc3339 = dt.Rfc3339
			o.ErrorDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.ErrorDate.Epoch = dt.Epoch
				o.ErrorDate.Rfc3339 = dt.Rfc3339
				o.ErrorDate.Offset = dt.Offset
			}
		}
	} else {
		o.ErrorDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["error_message"].(*string); ok {
		o.ErrorMessage = val
	} else if val, ok := kv["error_message"].(string); ok {
		o.ErrorMessage = &val
	} else {
		if val, ok := kv["error_message"]; ok {
			if val == nil {
				o.ErrorMessage = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.ErrorMessage = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["errored"].(*bool); ok {
		o.Errored = val
	} else if val, ok := kv["errored"].(bool); ok {
		o.Errored = &val
	} else {
		if val, ok := kv["errored"]; ok {
			if val == nil {
				o.Errored = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Errored = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	// Deprecated
	if val, ok := kv["export_acknowledged"].(*bool); ok {
		o.ExportAcknowledged = val
	} else if val, ok := kv["export_acknowledged"].(bool); ok {
		o.ExportAcknowledged = &val
	} else {
		if val, ok := kv["export_acknowledged"]; ok {
			if val == nil {
				o.ExportAcknowledged = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.ExportAcknowledged = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	if val, ok := kv["export_liveness"].(*bool); ok {
		o.ExportLiveness = val
	} else if val, ok := kv["export_liveness"].(bool); ok {
		o.ExportLiveness = &val
	} else {
		if val, ok := kv["export_liveness"]; ok {
			if val == nil {
				o.ExportLiveness = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.ExportLiveness = number.BoolPointer(number.ToBoolAny(val))
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
	if val, ok := kv["interval"].(int64); ok {
		o.Interval = val
	} else {
		if val, ok := kv["interval"]; ok {
			if val == nil {
				o.Interval = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Interval = number.ToInt64Any(val)
			}
		}
	}
	if val, ok := kv["job_id"].(*string); ok {
		o.JobID = val
	} else if val, ok := kv["job_id"].(string); ok {
		o.JobID = &val
	} else {
		if val, ok := kv["job_id"]; ok {
			if val == nil {
				o.JobID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.JobID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["location"].(IntegrationInstanceLocation); ok {
		o.Location = val
	} else {
		if em, ok := kv["location"].(map[string]interface{}); ok {

			ev := em["agent.location"].(string)
			switch ev {
			case "private", "PRIVATE":
				o.Location = 0
			case "cloud", "CLOUD":
				o.Location = 1
			}
		}
		if em, ok := kv["location"].(string); ok {
			switch em {
			case "private", "PRIVATE":
				o.Location = 0
			case "cloud", "CLOUD":
				o.Location = 1
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
	if val, ok := kv["paused"].(bool); ok {
		o.Paused = val
	} else {
		if val, ok := kv["paused"]; ok {
			if val == nil {
				o.Paused = false
			} else {
				o.Paused = number.ToBoolAny(val)
			}
		}
	}
	if val, ok := kv["private_key"].(*string); ok {
		o.PrivateKey = val
	} else if val, ok := kv["private_key"].(string); ok {
		o.PrivateKey = &val
	} else {
		if val, ok := kv["private_key"]; ok {
			if val == nil {
				o.PrivateKey = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.PrivateKey = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	// Deprecated
	if val, ok := kv["processed"].(*bool); ok {
		o.Processed = val
	} else if val, ok := kv["processed"].(bool); ok {
		o.Processed = &val
	} else {
		if val, ok := kv["processed"]; ok {
			if val == nil {
				o.Processed = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Processed = number.BoolPointer(number.ToBoolAny(val))
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
	if val, ok := kv["requires_historical"].(bool); ok {
		o.RequiresHistorical = val
	} else {
		if val, ok := kv["requires_historical"]; ok {
			if val == nil {
				o.RequiresHistorical = false
			} else {
				o.RequiresHistorical = number.ToBoolAny(val)
			}
		}
	}
	if val, ok := kv["setup"].(IntegrationInstanceSetup); ok {
		o.Setup = val
	} else {
		if em, ok := kv["setup"].(map[string]interface{}); ok {

			ev := em["agent.setup"].(string)
			switch ev {
			case "config", "CONFIG":
				o.Setup = 0
			case "ready", "READY":
				o.Setup = 1
			case "running", "RUNNING":
				o.Setup = 2
			}
		}
		if em, ok := kv["setup"].(string); ok {
			switch em {
			case "config", "CONFIG":
				o.Setup = 0
			case "ready", "READY":
				o.Setup = 1
			case "running", "RUNNING":
				o.Setup = 2
			}
		}
	}
	if val, ok := kv["state"].(IntegrationInstanceState); ok {
		o.State = val
	} else {
		if em, ok := kv["state"].(map[string]interface{}); ok {

			ev := em["agent.state"].(string)
			switch ev {
			case "idle", "IDLE":
				o.State = 0
			case "exporting", "EXPORTING":
				o.State = 1
			}
		}
		if em, ok := kv["state"].(string); ok {
			switch em {
			case "idle", "IDLE":
				o.State = 0
			case "exporting", "EXPORTING":
				o.State = 1
			}
		}
	}
	if val, ok := kv["throttled"].(*bool); ok {
		o.Throttled = val
	} else if val, ok := kv["throttled"].(bool); ok {
		o.Throttled = &val
	} else {
		if val, ok := kv["throttled"]; ok {
			if val == nil {
				o.Throttled = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Throttled = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}

	if o.ThrottledUntil == nil {
		o.ThrottledUntil = &IntegrationInstanceThrottledUntil{}
	}

	if val, ok := kv["throttled_until"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.ThrottledUntil.FromMap(kv)
		} else if sv, ok := val.(IntegrationInstanceThrottledUntil); ok {
			// struct
			o.ThrottledUntil = &sv
		} else if sp, ok := val.(*IntegrationInstanceThrottledUntil); ok {
			// struct pointer
			o.ThrottledUntil = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.ThrottledUntil.Epoch = dt.Epoch
			o.ThrottledUntil.Rfc3339 = dt.Rfc3339
			o.ThrottledUntil.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.ThrottledUntil.Epoch = dt.Epoch
			o.ThrottledUntil.Rfc3339 = dt.Rfc3339
			o.ThrottledUntil.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.ThrottledUntil.Epoch = dt.Epoch
				o.ThrottledUntil.Rfc3339 = dt.Rfc3339
				o.ThrottledUntil.Offset = dt.Offset
			}
		}
	} else {
		o.ThrottledUntil.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["updated_ts"].(int64); ok {
		o.UpdatedAt = val
	} else {
		if val, ok := kv["updated_ts"]; ok {
			if val == nil {
				o.UpdatedAt = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.UpdatedAt = number.ToInt64Any(val)
			}
		}
	}
	if val, ok := kv["upgrade_date_ts"].(*int64); ok {
		o.UpgradeDateAt = val
	} else if val, ok := kv["upgrade_date_ts"].(int64); ok {
		o.UpgradeDateAt = &val
	} else {
		if val, ok := kv["upgrade_date_ts"]; ok {
			if val == nil {
				o.UpgradeDateAt = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["long"]
				}
				o.UpgradeDateAt = number.Int64Pointer(number.ToInt64Any(val))
			}
		}
	}
	if val, ok := kv["upgrade_expires_date_ts"].(*int64); ok {
		o.UpgradeExpiresDateAt = val
	} else if val, ok := kv["upgrade_expires_date_ts"].(int64); ok {
		o.UpgradeExpiresDateAt = &val
	} else {
		if val, ok := kv["upgrade_expires_date_ts"]; ok {
			if val == nil {
				o.UpgradeExpiresDateAt = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["long"]
				}
				o.UpgradeExpiresDateAt = number.Int64Pointer(number.ToInt64Any(val))
			}
		}
	}
	if val, ok := kv["upgrade_message"].(*string); ok {
		o.UpgradeMessage = val
	} else if val, ok := kv["upgrade_message"].(string); ok {
		o.UpgradeMessage = &val
	} else {
		if val, ok := kv["upgrade_message"]; ok {
			if val == nil {
				o.UpgradeMessage = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.UpgradeMessage = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["upgrade_required"].(bool); ok {
		o.UpgradeRequired = val
	} else {
		if val, ok := kv["upgrade_required"]; ok {
			if val == nil {
				o.UpgradeRequired = false
			} else {
				o.UpgradeRequired = number.ToBoolAny(val)
			}
		}
	}

	if o == nil {

		o.Webhooks = make([]IntegrationInstanceWebhooks, 0)

	}
	if val, ok := kv["webhooks"]; ok {
		if sv, ok := val.([]IntegrationInstanceWebhooks); ok {
			o.Webhooks = sv
		} else if sp, ok := val.([]*IntegrationInstanceWebhooks); ok {
			o.Webhooks = o.Webhooks[:0]
			for _, e := range sp {
				o.Webhooks = append(o.Webhooks, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(IntegrationInstanceWebhooks); ok {
					o.Webhooks = append(o.Webhooks, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm IntegrationInstanceWebhooks
					fm.FromMap(av)
					o.Webhooks = append(o.Webhooks, fm)
				} else {
					b, _ := json.Marshal(ae)
					bkv := make(map[string]interface{})
					json.Unmarshal(b, &bkv)
					var av IntegrationInstanceWebhooks
					av.FromMap(bkv)
					o.Webhooks = append(o.Webhooks, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(IntegrationInstanceWebhooks); ok {
					o.Webhooks = append(o.Webhooks, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm IntegrationInstanceWebhooks
					fm.FromMap(r)
					o.Webhooks = append(o.Webhooks, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := IntegrationInstanceWebhooks{}
					fm.FromMap(r)
					o.Webhooks = append(o.Webhooks, fm)
				}
			}
		} else {
			arr := reflect.ValueOf(val)
			if arr.Kind() == reflect.Slice {
				for i := 0; i < arr.Len(); i++ {
					item := arr.Index(i)
					if item.CanAddr() {
						v := item.Addr().MethodByName("ToMap")
						if !v.IsNil() {
							m := v.Call([]reflect.Value{})
							var fm IntegrationInstanceWebhooks
							fm.FromMap(m[0].Interface().(map[string]interface{}))
							o.Webhooks = append(o.Webhooks, fm)
						}
					}
				}
			}
		}
	}

	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *IntegrationInstance) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Active)
	args = append(args, o.AutoConfigure)
	args = append(args, o.Config)
	args = append(args, o.CreatedByProfileID)
	args = append(args, o.CreatedByUserID)
	args = append(args, o.CreatedDate)
	args = append(args, o.CreatedAt)
	args = append(args, o.CustomerID)
	args = append(args, o.Deleted)
	args = append(args, o.DeletedByProfileID)
	args = append(args, o.DeletedByUserID)
	args = append(args, o.DeletedDate)
	args = append(args, o.EnrollmentID)
	args = append(args, o.ErrorDate)
	args = append(args, o.ErrorMessage)
	args = append(args, o.Errored)
	args = append(args, o.ExportAcknowledged)
	args = append(args, o.ExportLiveness)
	args = append(args, o.ID)
	args = append(args, o.IntegrationID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.Interval)
	args = append(args, o.JobID)
	args = append(args, o.Location)
	args = append(args, o.Name)
	args = append(args, o.Paused)
	args = append(args, o.PrivateKey)
	args = append(args, o.Processed)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RequiresHistorical)
	args = append(args, o.Setup)
	args = append(args, o.State)
	args = append(args, o.Throttled)
	args = append(args, o.ThrottledUntil)
	args = append(args, o.UpdatedAt)
	args = append(args, o.UpgradeDateAt)
	args = append(args, o.UpgradeExpiresDateAt)
	args = append(args, o.UpgradeMessage)
	args = append(args, o.UpgradeRequired)
	args = append(args, o.Webhooks)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *IntegrationInstance) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *IntegrationInstance) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// GetHydrationQuery returns a query for all fields, and one level deep of relations.
// This query requires "id" to be in the query variables.
func (o *IntegrationInstance) GetHydrationQuery() string {
	return `query GoIntegrationInstanceQuery($id: ID, $nocache: Boolean, $etag: String) {
	agent {
		IntegrationInstance(_id: $id, nocache: $nocache, etag: $etag) {
			active
			auto_configure
			config
			created_by_profile_id
			created_by_user_id
			created_date {
			epoch
			offset
			rfc3339
			}
			created_ts
			customer_id
			deleted
			deleted_by_profile_id
			deleted_by_user_id
			deleted_date {
			epoch
			offset
			rfc3339
			}
			enrollment_id
			error_date {
			epoch
			offset
			rfc3339
			}
			error_message
			errored
			export_acknowledged
			export_liveness
			_id
			integration_id
			integration_instance_id
			interval
			job_id
			location
			name
			paused
			private_key
			processed
			ref_id
			ref_type
			requires_historical
			setup
			state
			throttled
			throttled_until {
			epoch
			offset
			rfc3339
			}
			updated_ts
			upgrade_date_ts
			upgrade_expires_date_ts
			upgrade_message
			upgrade_required
			webhooks {
			enabled
			error_message
			errored
			ref_id
			url
			}
		}
	}
}
`
}

func getIntegrationInstanceQueryFields() string {
	var sb strings.Builder

	// scalar
	sb.WriteString("\t\t\tactive\n")
	// scalar
	sb.WriteString("\t\t\tauto_configure\n")
	// scalar
	sb.WriteString("\t\t\tconfig\n")
	// scalar
	sb.WriteString("\t\t\tcreated_by_profile_id\n")
	// scalar
	sb.WriteString("\t\t\tcreated_by_user_id\n")
	// object with fields
	sb.WriteString("\t\t\tcreated_date {\n")

	// scalar
	sb.WriteString("\t\t\tepoch\n")
	// scalar
	sb.WriteString("\t\t\toffset\n")
	// scalar
	sb.WriteString("\t\t\trfc3339\n")
	sb.WriteString("\t\t\t}\n")
	// scalar
	sb.WriteString("\t\t\tcreated_ts\n")
	// scalar
	sb.WriteString("\t\t\tcustomer_id\n")
	// scalar
	sb.WriteString("\t\t\tdeleted\n")
	// scalar
	sb.WriteString("\t\t\tdeleted_by_profile_id\n")
	// scalar
	sb.WriteString("\t\t\tdeleted_by_user_id\n")
	// object with fields
	sb.WriteString("\t\t\tdeleted_date {\n")

	// scalar
	sb.WriteString("\t\t\tepoch\n")
	// scalar
	sb.WriteString("\t\t\toffset\n")
	// scalar
	sb.WriteString("\t\t\trfc3339\n")
	sb.WriteString("\t\t\t}\n")
	// scalar
	sb.WriteString("\t\t\tenrollment_id\n")
	// object with fields
	sb.WriteString("\t\t\terror_date {\n")

	// scalar
	sb.WriteString("\t\t\tepoch\n")
	// scalar
	sb.WriteString("\t\t\toffset\n")
	// scalar
	sb.WriteString("\t\t\trfc3339\n")
	sb.WriteString("\t\t\t}\n")
	// scalar
	sb.WriteString("\t\t\terror_message\n")
	// scalar
	sb.WriteString("\t\t\terrored\n")
	// scalar
	sb.WriteString("\t\t\texport_acknowledged\n")
	// scalar
	sb.WriteString("\t\t\texport_liveness\n")
	// id
	sb.WriteString("\t\t\t_id\n")
	// scalar
	sb.WriteString("\t\t\tintegration_id\n")
	// scalar
	sb.WriteString("\t\t\tintegration_instance_id\n")
	// scalar
	sb.WriteString("\t\t\tinterval\n")
	// scalar
	sb.WriteString("\t\t\tjob_id\n")
	// scalar
	sb.WriteString("\t\t\tlocation\n")
	// scalar
	sb.WriteString("\t\t\tname\n")
	// scalar
	sb.WriteString("\t\t\tpaused\n")
	// scalar
	sb.WriteString("\t\t\tprivate_key\n")
	// scalar
	sb.WriteString("\t\t\tprocessed\n")
	// scalar
	sb.WriteString("\t\t\tref_id\n")
	// scalar
	sb.WriteString("\t\t\tref_type\n")
	// scalar
	sb.WriteString("\t\t\trequires_historical\n")
	// scalar
	sb.WriteString("\t\t\tsetup\n")
	// scalar
	sb.WriteString("\t\t\tstate\n")
	// scalar
	sb.WriteString("\t\t\tthrottled\n")
	// object with fields
	sb.WriteString("\t\t\tthrottled_until {\n")

	// scalar
	sb.WriteString("\t\t\tepoch\n")
	// scalar
	sb.WriteString("\t\t\toffset\n")
	// scalar
	sb.WriteString("\t\t\trfc3339\n")
	sb.WriteString("\t\t\t}\n")
	// scalar
	sb.WriteString("\t\t\tupdated_ts\n")
	// scalar
	sb.WriteString("\t\t\tupgrade_date_ts\n")
	// scalar
	sb.WriteString("\t\t\tupgrade_expires_date_ts\n")
	// scalar
	sb.WriteString("\t\t\tupgrade_message\n")
	// scalar
	sb.WriteString("\t\t\tupgrade_required\n")
	// object with fields
	sb.WriteString("\t\t\twebhooks {\n")

	// scalar
	sb.WriteString("\t\t\tenabled\n")
	// scalar
	sb.WriteString("\t\t\terror_message\n")
	// scalar
	sb.WriteString("\t\t\terrored\n")
	// scalar
	sb.WriteString("\t\t\tref_id\n")
	// scalar
	sb.WriteString("\t\t\turl\n")
	sb.WriteString("\t\t\t}\n")
	return sb.String()
}

// IntegrationInstancePageInfo is a grapqhl PageInfo
type IntegrationInstancePageInfo struct {
	StartCursor     string `json:"startCursor,omitempty"`
	EndCursor       string `json:"endCursor,omitempty"`
	HasNextPage     bool   `json:"hasNextPage,omitempty"`
	HasPreviousPage bool   `json:"hasPreviousPage,omitempty"`
}

// IntegrationInstanceCacheInfo is the grapqhl CacheInfo
type IntegrationInstanceCacheInfo struct {
	Cached bool    `json:"cached,omitempty"`
	ID     *string `json:"id,omitempty"`
	ETag   *string `json:"etag,omitempty"`
}

// IntegrationInstanceConnection is a grapqhl connection
type IntegrationInstanceConnection struct {
	Edges      []*IntegrationInstanceEdge   `json:"edges,omitempty"`
	PageInfo   IntegrationInstancePageInfo  `json:"pageInfo,omitempty"`
	CacheInfo  IntegrationInstanceCacheInfo `json:"cacheInfo,omitempty"`
	TotalCount *int64                       `json:"totalCount,omitempty"`
}

// IntegrationInstanceEdge is a grapqhl edge
type IntegrationInstanceEdge struct {
	Node *IntegrationInstance `json:"node,omitempty"`
}

// QueryManyIntegrationInstanceNode is a grapqhl query many node
type QueryManyIntegrationInstanceNode struct {
	Object *IntegrationInstanceConnection `json:"IntegrationInstances,omitempty"`
}

// QueryManyIntegrationInstanceData is a grapqhl query many data node
type QueryManyIntegrationInstanceData struct {
	Data *QueryManyIntegrationInstanceNode `json:"agent,omitempty"`
}

// QueryOneIntegrationInstanceNode is a grapqhl query one node
type QueryOneIntegrationInstanceNode struct {
	Object *IntegrationInstance `json:"IntegrationInstance,omitempty"`
}

// QueryOneIntegrationInstanceData is a grapqhl query one data node
type QueryOneIntegrationInstanceData struct {
	Data *QueryOneIntegrationInstanceNode `json:"agent,omitempty"`
}

// IntegrationInstanceQuery is query struct
type IntegrationInstanceQuery struct {
	Filters []string      `json:"filters,omitempty"`
	Params  []interface{} `json:"params,omitempty"`
}

// IntegrationInstanceOrder is the order direction
type IntegrationInstanceOrder *string

var (
	// IntegrationInstanceOrderDesc is descending
	IntegrationInstanceOrderDesc IntegrationInstanceOrder = pstrings.Pointer("DESC")
	// IntegrationInstanceOrderAsc is ascending
	IntegrationInstanceOrderAsc IntegrationInstanceOrder = pstrings.Pointer("ASC")
)

// IntegrationInstanceQueryInput is query input struct
type IntegrationInstanceQueryInput struct {
	First   *int64                    `json:"first,omitempty"`
	Last    *int64                    `json:"last,omitempty"`
	Before  *string                   `json:"before,omitempty"`
	After   *string                   `json:"after,omitempty"`
	Query   *IntegrationInstanceQuery `json:"query,omitempty"`
	OrderBy *string                   `json:"orderBy,omitempty"`
	Order   IntegrationInstanceOrder  `json:"order,omitempty"`
	NoCache bool                      `json:"nocache,omitempty"`
	ETag    *string                   `json:"etag,omitempty"`
}

// NewIntegrationInstanceQuery is a convenience for building a *IntegrationInstanceQuery
func NewIntegrationInstanceQuery(params ...interface{}) *IntegrationInstanceQueryInput {
	if len(params)%2 != 0 {
		panic("incorrect number of arguments passed")
	}
	q := &IntegrationInstanceQuery{
		Filters: make([]string, 0),
		Params:  make([]interface{}, 0),
	}
	for i := 0; i < len(params); i += 2 {
		q.Filters = append(q.Filters, params[i].(string))
		q.Params = append(q.Params, params[i+1])
	}
	return &IntegrationInstanceQueryInput{
		Query: q,
	}
}

// FindIntegrationInstance will query an IntegrationInstance by id
func FindIntegrationInstance(client graphql.Client, id string) (*IntegrationInstance, error) {
	variables := make(graphql.Variables)
	variables["id"] = id
	var sb strings.Builder
	sb.WriteString("query GoIntegrationInstanceQuery($id: ID) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tIntegrationInstance(_id: $id) {\n")
	sb.WriteString(getIntegrationInstanceQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryOneIntegrationInstanceData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	if res.Data != nil {
		return res.Data.Object, nil
	}
	return nil, nil
}

// FindIntegrationInstanceWithoutCache will query an IntegrationInstance by id avoiding the cache
func FindIntegrationInstanceWithoutCache(client graphql.Client, id string) (*IntegrationInstance, error) {
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["nocache"] = true
	var sb strings.Builder
	sb.WriteString("query GoIntegrationInstanceQuery($id: ID, $nocache: Boolean) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tIntegrationInstance(_id: $id, nocache: $nocache) {\n")
	sb.WriteString(getIntegrationInstanceQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryOneIntegrationInstanceData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	if res.Data != nil {
		return res.Data.Object, nil
	}
	return nil, nil
}

// FindIntegrationInstances will query for any IntegrationInstances matching the query
func FindIntegrationInstances(client graphql.Client, input *IntegrationInstanceQueryInput) (*IntegrationInstanceConnection, error) {
	variables := make(graphql.Variables)
	if input != nil {
		variables["first"] = input.First
		variables["last"] = input.Last
		variables["before"] = input.Before
		variables["after"] = input.After
		variables["query"] = input.Query
		if input.OrderBy != nil {
			variables["orderBy"] = *input.OrderBy
		}
		if input.Order != nil {
			variables["order"] = *input.Order
		}
		variables["nocache"] = input.NoCache
	}
	var sb strings.Builder
	sb.WriteString("query GoIntegrationInstanceQueryMany($first: Int, $last: Int, $before: Cursor, $after: Cursor, $query: QueryInput, $order: SortOrderEnum, $orderBy: AgentIntegrationInstanceColumnEnum, $nocache: Boolean) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tIntegrationInstances(first: $first last: $last before: $before after: $after query: $query orderBy: $orderBy order: $order nocache: $nocache) {\n")
	sb.WriteString("\t\t\tpageInfo {\n")
	sb.WriteString("\t\t\t\tstartCursor\n")
	sb.WriteString("\t\t\t\tendCursor\n")
	sb.WriteString("\t\t\t\thasNextPage\n")
	sb.WriteString("\t\t\t\thasPreviousPage\n")
	sb.WriteString("\t\t\t}\n")
	sb.WriteString("\t\t\tcacheInfo {\n")
	sb.WriteString("\t\t\t\tcached\n")
	sb.WriteString("\t\t\t\tid\n")
	sb.WriteString("\t\t\t\tetag\n")
	sb.WriteString("\t\t\t}\n")
	sb.WriteString("\t\t\ttotalCount\n")
	sb.WriteString("\t\t\tedges {\n")
	sb.WriteString("\t\t\t\tnode {\n")
	sb.WriteString(getIntegrationInstanceQueryFields())
	sb.WriteString("\t\t\t\t}\n")
	sb.WriteString("\t\t\t}\n")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryManyIntegrationInstanceData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	return res.Data.Object, nil
}

// FindIntegrationInstancesPaginatedCallback is a callback function for handling each page
type FindIntegrationInstancesPaginatedCallback func(conn *IntegrationInstanceConnection) (bool, error)

// FindIntegrationInstancesPaginated will query for any IntegrationInstances matching the query and return each page callback
func FindIntegrationInstancesPaginated(client graphql.Client, query *IntegrationInstanceQuery, pageSize int64, callback FindIntegrationInstancesPaginatedCallback) error {
	input := &IntegrationInstanceQueryInput{
		First: &pageSize,
		Query: query,
	}
	for {
		res, err := FindIntegrationInstances(client, input)
		if err != nil {
			return err
		}
		if res == nil {
			break
		}
		ok, err := callback(res)
		if err != nil {
			return err
		}
		if !ok || !res.PageInfo.HasNextPage {
			break
		}
		input.After = &res.PageInfo.EndCursor
	}
	return nil
}

// FindIntegrationInstancesPaginatedWithoutCache will query for any IntegrationInstances matching the query and return each page callback
func FindIntegrationInstancesPaginatedWithoutCache(client graphql.Client, query *IntegrationInstanceQuery, pageSize int64, callback FindIntegrationInstancesPaginatedCallback) error {
	input := &IntegrationInstanceQueryInput{
		First:   &pageSize,
		Query:   query,
		NoCache: true,
	}
	for {
		res, err := FindIntegrationInstances(client, input)
		if err != nil {
			return err
		}
		if res == nil {
			break
		}
		ok, err := callback(res)
		if err != nil {
			return err
		}
		if !ok || !res.PageInfo.HasNextPage {
			break
		}
		input.After = &res.PageInfo.EndCursor
	}
	return nil
}

// UpdateIntegrationInstanceNode is a grapqhl update node
type UpdateIntegrationInstanceNode struct {
	Object *IntegrationInstance `json:"updateIntegrationInstance,omitempty"`
}

// UpdateIntegrationInstanceData is a grapqhl update data node
type UpdateIntegrationInstanceData struct {
	Data *UpdateIntegrationInstanceNode `json:"agent,omitempty"`
}

// ExecIntegrationInstanceUpdateMutation returns a graphql update mutation result for IntegrationInstance
func ExecIntegrationInstanceUpdateMutation(client graphql.Client, id string, input graphql.Variables, upsert bool) (*IntegrationInstance, error) {
	if !upsert && id == "" {
		return nil, fmt.Errorf("id is required with upsert false")
	}
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["upsert"] = upsert
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation GoIntegrationInstanceUpdateMutation($id: String, $input: UpdateAgentIntegrationInstanceInput, $upsert: Boolean) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tupdateIntegrationInstance(_id: $id, input: $input, upsert: $upsert) {\n")
	sb.WriteString(getIntegrationInstanceQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateIntegrationInstanceData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	return res.Data.Object, nil
}

// ExecIntegrationInstanceSilentUpdateMutation returns a graphql update mutation result for IntegrationInstance
func ExecIntegrationInstanceSilentUpdateMutation(client graphql.Client, id string, input graphql.Variables, upsert bool) error {
	if !upsert && id == "" {
		return fmt.Errorf("id is required with upsert false")
	}
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["upsert"] = upsert
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation GoIntegrationInstanceUpdateMutation($id: String, $input: UpdateAgentIntegrationInstanceInput, $upsert: Boolean) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tupdateIntegrationInstance(_id: $id, input: $input, upsert: $upsert) {\n")
	sb.WriteString("\t\t\t_id")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateIntegrationInstanceData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

// ExecIntegrationInstanceDeleteMutation executes a graphql delete mutation for IntegrationInstance
func ExecIntegrationInstanceDeleteMutation(client graphql.Client, id string) error {
	variables := make(graphql.Variables)
	variables["id"] = id
	var sb strings.Builder
	sb.WriteString("mutation GoIntegrationInstanceDeleteMutation($id: String!) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tdeleteIntegrationInstance(_id: $id)\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res interface{}
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

func CreateIntegrationInstance(client graphql.Client, model IntegrationInstance) error {
	variables := make(graphql.Variables)
	input := model.ToMap()
	delete(input, "hashcode")
	delete(input, "customer_id")
	delete(input, "created_ts")
	delete(input, "id")
	input["_id"] = model.GetID()
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation GoCreateIntegrationInstance($input: CreateAgentIntegrationInstanceInput!) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tcreateIntegrationInstance(input: $input) {\n")
	sb.WriteString("\t\t\t_id")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateIntegrationInstanceData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

// IntegrationInstancePartial is a partial struct for upsert mutations for IntegrationInstance
type IntegrationInstancePartial struct {
	// Active If true, the integration is still active
	Active *bool `json:"active,omitempty"`
	// AutoConfigure if this integration should be auto configured by the integration before activating
	AutoConfigure *bool `json:"auto_configure,omitempty"`
	// Config the integration configuration controlled by the integration itself
	Config *string `json:"config,omitempty"`
	// CreatedByProfileID The id of the profile for the user that created the integration
	CreatedByProfileID *string `json:"created_by_profile_id,omitempty"`
	// CreatedByUserID The id of the user that created the integration
	CreatedByUserID *string `json:"created_by_user_id,omitempty"`
	// CreatedDate when the integration was created
	CreatedDate *IntegrationInstanceCreatedDate `json:"created_date,omitempty"`
	// Deleted If true, the integration has been deleted
	Deleted *bool `json:"deleted,omitempty"`
	// DeletedByProfileID The id of the profile for the user that deleted the integration
	DeletedByProfileID *string `json:"deleted_by_profile_id,omitempty"`
	// DeletedByUserID The id of the user that deleted the integration
	DeletedByUserID *string `json:"deleted_by_user_id,omitempty"`
	// DeletedDate when the integration was deleted
	DeletedDate *IntegrationInstanceDeletedDate `json:"deleted_date,omitempty"`
	// EnrollmentID if the integration is linked to a self-managed agent, it will have the enrollment_id set otherwise will be null
	EnrollmentID *string `json:"enrollment_id,omitempty"`
	// ErrorDate The date of the error from an export run
	ErrorDate *IntegrationInstanceErrorDate `json:"error_date,omitempty"`
	// ErrorMessage The error message from an export run
	ErrorMessage *string `json:"error_message,omitempty"`
	// Errored If authorization failed by the agent or any other error
	Errored *bool `json:"errored,omitempty"`
	// ExportAcknowledged Set to true an export has been received by the agent.
	//
	// Deprecated: no longer used, see export_liveness
	ExportAcknowledged *bool `json:"export_acknowledged,omitempty"`
	// IntegrationID The unique id for the integration
	IntegrationID *string `json:"integration_id,omitempty"`
	// Interval the interval in milliseconds for how often an export job is scheduled
	Interval *int64 `json:"interval,omitempty"`
	// JobID the current (if EXPORTING) or previous (if IDLE) job id which is useful for debugging
	JobID *string `json:"job_id,omitempty"`
	// Location The location of this integration (on-premise / private or cloud)
	Location *IntegrationInstanceLocation `json:"location,omitempty"`
	// Name The user friendly name of the integration
	Name *string `json:"name,omitempty"`
	// Paused true if the agent is paused and should not start new scheduled jobs
	Paused *bool `json:"paused,omitempty"`
	// PrivateKey the private key for the instance if needed by an integration
	PrivateKey *string `json:"private_key,omitempty"`
	// Processed If the integration has been processed at least once
	//
	// Deprecated: this field should not be used anymore. use stat to get details about processing
	Processed *bool `json:"processed,omitempty"`
	// RequiresHistorical flag which can be set to trigger a historical on the next scheduler visit
	RequiresHistorical *bool `json:"requires_historical,omitempty"`
	// Setup the setup state of the integration
	Setup *IntegrationInstanceSetup `json:"setup,omitempty"`
	// State the current state of the integration
	State *IntegrationInstanceState `json:"state,omitempty"`
	// Throttled Set to true when integration is throttled.
	Throttled *bool `json:"throttled,omitempty"`
	// ThrottledUntil After throttling integration, we set this field for estimated resume date.
	ThrottledUntil *IntegrationInstanceThrottledUntil `json:"throttled_until,omitempty"`
	// UpgradeDateAt If upgrade is required, the date when the upgrade was set
	UpgradeDateAt *int64 `json:"upgrade_date_ts,omitempty"`
	// UpgradeExpiresDateAt If upgrade is required and there's a due date this will be set
	UpgradeExpiresDateAt *int64 `json:"upgrade_expires_date_ts,omitempty"`
	// UpgradeMessage If upgrade is required, the message to display to the user
	UpgradeMessage *string `json:"upgrade_message,omitempty"`
	// UpgradeRequired If true, the integration requires a manual upgrade
	UpgradeRequired *bool `json:"upgrade_required,omitempty"`
	// Webhooks for any webhooks installed at the integration instance
	Webhooks []IntegrationInstanceWebhooks `json:"webhooks,omitempty"`
}

var _ datamodel.PartialModel = (*IntegrationInstancePartial)(nil)

// GetModelName returns the name of the model
func (o *IntegrationInstancePartial) GetModelName() datamodel.ModelNameType {
	return IntegrationInstanceModelName
}

// ToMap returns the object as a map
func (o *IntegrationInstancePartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"active":                toIntegrationInstanceObject(o.Active, true),
		"auto_configure":        toIntegrationInstanceObject(o.AutoConfigure, true),
		"config":                toIntegrationInstanceObject(o.Config, true),
		"created_by_profile_id": toIntegrationInstanceObject(o.CreatedByProfileID, true),
		"created_by_user_id":    toIntegrationInstanceObject(o.CreatedByUserID, true),
		"created_date":          toIntegrationInstanceObject(o.CreatedDate, true),
		"deleted":               toIntegrationInstanceObject(o.Deleted, true),
		"deleted_by_profile_id": toIntegrationInstanceObject(o.DeletedByProfileID, true),
		"deleted_by_user_id":    toIntegrationInstanceObject(o.DeletedByUserID, true),
		"deleted_date":          toIntegrationInstanceObject(o.DeletedDate, true),
		"enrollment_id":         toIntegrationInstanceObject(o.EnrollmentID, true),
		"error_date":            toIntegrationInstanceObject(o.ErrorDate, true),
		"error_message":         toIntegrationInstanceObject(o.ErrorMessage, true),
		"errored":               toIntegrationInstanceObject(o.Errored, true),
		"export_acknowledged":   toIntegrationInstanceObject(o.ExportAcknowledged, true),
		"integration_id":        toIntegrationInstanceObject(o.IntegrationID, true),
		"interval":              toIntegrationInstanceObject(o.Interval, true),
		"job_id":                toIntegrationInstanceObject(o.JobID, true),

		"location":            toIntegrationInstanceLocationEnum(o.Location),
		"name":                toIntegrationInstanceObject(o.Name, true),
		"paused":              toIntegrationInstanceObject(o.Paused, true),
		"private_key":         toIntegrationInstanceObject(o.PrivateKey, true),
		"processed":           toIntegrationInstanceObject(o.Processed, true),
		"requires_historical": toIntegrationInstanceObject(o.RequiresHistorical, true),

		"setup": toIntegrationInstanceSetupEnum(o.Setup),

		"state":                   toIntegrationInstanceStateEnum(o.State),
		"throttled":               toIntegrationInstanceObject(o.Throttled, true),
		"throttled_until":         toIntegrationInstanceObject(o.ThrottledUntil, true),
		"upgrade_date_ts":         toIntegrationInstanceObject(o.UpgradeDateAt, true),
		"upgrade_expires_date_ts": toIntegrationInstanceObject(o.UpgradeExpiresDateAt, true),
		"upgrade_message":         toIntegrationInstanceObject(o.UpgradeMessage, true),
		"upgrade_required":        toIntegrationInstanceObject(o.UpgradeRequired, true),
		"webhooks":                toIntegrationInstanceObject(o.Webhooks, true),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {
			if k == "created_date" {
				if dt, ok := v.(*IntegrationInstanceCreatedDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
			if k == "deleted_date" {
				if dt, ok := v.(*IntegrationInstanceDeletedDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
			if k == "error_date" {
				if dt, ok := v.(*IntegrationInstanceErrorDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
			if k == "throttled_until" {
				if dt, ok := v.(*IntegrationInstanceThrottledUntil); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}

			if k == "webhooks" {
				if arr, ok := v.([]IntegrationInstanceWebhooks); ok {
					if len(arr) == 0 {
						delete(kv, k)
					}
				}
			}
		}
	}
	return kv
}

// Stringify returns the object in JSON format as a string
func (o *IntegrationInstancePartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *IntegrationInstancePartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *IntegrationInstancePartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *IntegrationInstancePartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *IntegrationInstancePartial) FromMap(kv map[string]interface{}) {
	if val, ok := kv["active"].(*bool); ok {
		o.Active = val
	} else if val, ok := kv["active"].(bool); ok {
		o.Active = &val
	} else {
		if val, ok := kv["active"]; ok {
			if val == nil {
				o.Active = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Active = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	if val, ok := kv["auto_configure"].(*bool); ok {
		o.AutoConfigure = val
	} else if val, ok := kv["auto_configure"].(bool); ok {
		o.AutoConfigure = &val
	} else {
		if val, ok := kv["auto_configure"]; ok {
			if val == nil {
				o.AutoConfigure = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.AutoConfigure = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	if val, ok := kv["config"].(*string); ok {
		o.Config = val
	} else if val, ok := kv["config"].(string); ok {
		o.Config = &val
	} else {
		if val, ok := kv["config"]; ok {
			if val == nil {
				o.Config = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Config = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["created_by_profile_id"].(*string); ok {
		o.CreatedByProfileID = val
	} else if val, ok := kv["created_by_profile_id"].(string); ok {
		o.CreatedByProfileID = &val
	} else {
		if val, ok := kv["created_by_profile_id"]; ok {
			if val == nil {
				o.CreatedByProfileID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.CreatedByProfileID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["created_by_user_id"].(*string); ok {
		o.CreatedByUserID = val
	} else if val, ok := kv["created_by_user_id"].(string); ok {
		o.CreatedByUserID = &val
	} else {
		if val, ok := kv["created_by_user_id"]; ok {
			if val == nil {
				o.CreatedByUserID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.CreatedByUserID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if o.CreatedDate == nil {
		o.CreatedDate = &IntegrationInstanceCreatedDate{}
	}

	if val, ok := kv["created_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CreatedDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationInstanceCreatedDate); ok {
			// struct
			o.CreatedDate = &sv
		} else if sp, ok := val.(*IntegrationInstanceCreatedDate); ok {
			// struct pointer
			o.CreatedDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.CreatedDate.Epoch = dt.Epoch
			o.CreatedDate.Rfc3339 = dt.Rfc3339
			o.CreatedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.CreatedDate.Epoch = dt.Epoch
			o.CreatedDate.Rfc3339 = dt.Rfc3339
			o.CreatedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.CreatedDate.Epoch = dt.Epoch
				o.CreatedDate.Rfc3339 = dt.Rfc3339
				o.CreatedDate.Offset = dt.Offset
			}
		}
	} else {
		o.CreatedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["deleted"].(*bool); ok {
		o.Deleted = val
	} else if val, ok := kv["deleted"].(bool); ok {
		o.Deleted = &val
	} else {
		if val, ok := kv["deleted"]; ok {
			if val == nil {
				o.Deleted = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Deleted = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	if val, ok := kv["deleted_by_profile_id"].(*string); ok {
		o.DeletedByProfileID = val
	} else if val, ok := kv["deleted_by_profile_id"].(string); ok {
		o.DeletedByProfileID = &val
	} else {
		if val, ok := kv["deleted_by_profile_id"]; ok {
			if val == nil {
				o.DeletedByProfileID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.DeletedByProfileID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["deleted_by_user_id"].(*string); ok {
		o.DeletedByUserID = val
	} else if val, ok := kv["deleted_by_user_id"].(string); ok {
		o.DeletedByUserID = &val
	} else {
		if val, ok := kv["deleted_by_user_id"]; ok {
			if val == nil {
				o.DeletedByUserID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.DeletedByUserID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if o.DeletedDate == nil {
		o.DeletedDate = &IntegrationInstanceDeletedDate{}
	}

	if val, ok := kv["deleted_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.DeletedDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationInstanceDeletedDate); ok {
			// struct
			o.DeletedDate = &sv
		} else if sp, ok := val.(*IntegrationInstanceDeletedDate); ok {
			// struct pointer
			o.DeletedDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.DeletedDate.Epoch = dt.Epoch
			o.DeletedDate.Rfc3339 = dt.Rfc3339
			o.DeletedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.DeletedDate.Epoch = dt.Epoch
			o.DeletedDate.Rfc3339 = dt.Rfc3339
			o.DeletedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.DeletedDate.Epoch = dt.Epoch
				o.DeletedDate.Rfc3339 = dt.Rfc3339
				o.DeletedDate.Offset = dt.Offset
			}
		}
	} else {
		o.DeletedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["enrollment_id"].(*string); ok {
		o.EnrollmentID = val
	} else if val, ok := kv["enrollment_id"].(string); ok {
		o.EnrollmentID = &val
	} else {
		if val, ok := kv["enrollment_id"]; ok {
			if val == nil {
				o.EnrollmentID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.EnrollmentID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if o.ErrorDate == nil {
		o.ErrorDate = &IntegrationInstanceErrorDate{}
	}

	if val, ok := kv["error_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.ErrorDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationInstanceErrorDate); ok {
			// struct
			o.ErrorDate = &sv
		} else if sp, ok := val.(*IntegrationInstanceErrorDate); ok {
			// struct pointer
			o.ErrorDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.ErrorDate.Epoch = dt.Epoch
			o.ErrorDate.Rfc3339 = dt.Rfc3339
			o.ErrorDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.ErrorDate.Epoch = dt.Epoch
			o.ErrorDate.Rfc3339 = dt.Rfc3339
			o.ErrorDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.ErrorDate.Epoch = dt.Epoch
				o.ErrorDate.Rfc3339 = dt.Rfc3339
				o.ErrorDate.Offset = dt.Offset
			}
		}
	} else {
		o.ErrorDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["error_message"].(*string); ok {
		o.ErrorMessage = val
	} else if val, ok := kv["error_message"].(string); ok {
		o.ErrorMessage = &val
	} else {
		if val, ok := kv["error_message"]; ok {
			if val == nil {
				o.ErrorMessage = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.ErrorMessage = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["errored"].(*bool); ok {
		o.Errored = val
	} else if val, ok := kv["errored"].(bool); ok {
		o.Errored = &val
	} else {
		if val, ok := kv["errored"]; ok {
			if val == nil {
				o.Errored = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Errored = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	// Deprecated
	if val, ok := kv["export_acknowledged"].(*bool); ok {
		o.ExportAcknowledged = val
	} else if val, ok := kv["export_acknowledged"].(bool); ok {
		o.ExportAcknowledged = &val
	} else {
		if val, ok := kv["export_acknowledged"]; ok {
			if val == nil {
				o.ExportAcknowledged = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.ExportAcknowledged = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
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
	if val, ok := kv["interval"].(*int64); ok {
		o.Interval = val
	} else if val, ok := kv["interval"].(int64); ok {
		o.Interval = &val
	} else {
		if val, ok := kv["interval"]; ok {
			if val == nil {
				o.Interval = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["long"]
				}
				o.Interval = number.Int64Pointer(number.ToInt64Any(val))
			}
		}
	}
	if val, ok := kv["job_id"].(*string); ok {
		o.JobID = val
	} else if val, ok := kv["job_id"].(string); ok {
		o.JobID = &val
	} else {
		if val, ok := kv["job_id"]; ok {
			if val == nil {
				o.JobID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.JobID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["location"].(*IntegrationInstanceLocation); ok {
		o.Location = val
	} else if val, ok := kv["location"].(IntegrationInstanceLocation); ok {
		o.Location = &val
	} else {
		if val, ok := kv["location"]; ok {
			if val == nil {
				o.Location = toIntegrationInstanceLocationPointer(0)
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["IntegrationInstanceLocation"]
				}
				// this is an enum pointer
				if em, ok := val.(string); ok {
					switch em {
					case "private", "PRIVATE":
						o.Location = toIntegrationInstanceLocationPointer(0)
					case "cloud", "CLOUD":
						o.Location = toIntegrationInstanceLocationPointer(1)
					}
				}
			}
		}
	}
	if val, ok := kv["name"].(*string); ok {
		o.Name = val
	} else if val, ok := kv["name"].(string); ok {
		o.Name = &val
	} else {
		if val, ok := kv["name"]; ok {
			if val == nil {
				o.Name = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Name = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["paused"].(*bool); ok {
		o.Paused = val
	} else if val, ok := kv["paused"].(bool); ok {
		o.Paused = &val
	} else {
		if val, ok := kv["paused"]; ok {
			if val == nil {
				o.Paused = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Paused = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	if val, ok := kv["private_key"].(*string); ok {
		o.PrivateKey = val
	} else if val, ok := kv["private_key"].(string); ok {
		o.PrivateKey = &val
	} else {
		if val, ok := kv["private_key"]; ok {
			if val == nil {
				o.PrivateKey = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.PrivateKey = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	// Deprecated
	if val, ok := kv["processed"].(*bool); ok {
		o.Processed = val
	} else if val, ok := kv["processed"].(bool); ok {
		o.Processed = &val
	} else {
		if val, ok := kv["processed"]; ok {
			if val == nil {
				o.Processed = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Processed = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	if val, ok := kv["requires_historical"].(*bool); ok {
		o.RequiresHistorical = val
	} else if val, ok := kv["requires_historical"].(bool); ok {
		o.RequiresHistorical = &val
	} else {
		if val, ok := kv["requires_historical"]; ok {
			if val == nil {
				o.RequiresHistorical = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.RequiresHistorical = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	if val, ok := kv["setup"].(*IntegrationInstanceSetup); ok {
		o.Setup = val
	} else if val, ok := kv["setup"].(IntegrationInstanceSetup); ok {
		o.Setup = &val
	} else {
		if val, ok := kv["setup"]; ok {
			if val == nil {
				o.Setup = toIntegrationInstanceSetupPointer(0)
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["IntegrationInstanceSetup"]
				}
				// this is an enum pointer
				if em, ok := val.(string); ok {
					switch em {
					case "config", "CONFIG":
						o.Setup = toIntegrationInstanceSetupPointer(0)
					case "ready", "READY":
						o.Setup = toIntegrationInstanceSetupPointer(1)
					case "running", "RUNNING":
						o.Setup = toIntegrationInstanceSetupPointer(2)
					}
				}
			}
		}
	}
	if val, ok := kv["state"].(*IntegrationInstanceState); ok {
		o.State = val
	} else if val, ok := kv["state"].(IntegrationInstanceState); ok {
		o.State = &val
	} else {
		if val, ok := kv["state"]; ok {
			if val == nil {
				o.State = toIntegrationInstanceStatePointer(0)
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["IntegrationInstanceState"]
				}
				// this is an enum pointer
				if em, ok := val.(string); ok {
					switch em {
					case "idle", "IDLE":
						o.State = toIntegrationInstanceStatePointer(0)
					case "exporting", "EXPORTING":
						o.State = toIntegrationInstanceStatePointer(1)
					}
				}
			}
		}
	}
	if val, ok := kv["throttled"].(*bool); ok {
		o.Throttled = val
	} else if val, ok := kv["throttled"].(bool); ok {
		o.Throttled = &val
	} else {
		if val, ok := kv["throttled"]; ok {
			if val == nil {
				o.Throttled = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Throttled = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}

	if o.ThrottledUntil == nil {
		o.ThrottledUntil = &IntegrationInstanceThrottledUntil{}
	}

	if val, ok := kv["throttled_until"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.ThrottledUntil.FromMap(kv)
		} else if sv, ok := val.(IntegrationInstanceThrottledUntil); ok {
			// struct
			o.ThrottledUntil = &sv
		} else if sp, ok := val.(*IntegrationInstanceThrottledUntil); ok {
			// struct pointer
			o.ThrottledUntil = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.ThrottledUntil.Epoch = dt.Epoch
			o.ThrottledUntil.Rfc3339 = dt.Rfc3339
			o.ThrottledUntil.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.ThrottledUntil.Epoch = dt.Epoch
			o.ThrottledUntil.Rfc3339 = dt.Rfc3339
			o.ThrottledUntil.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.ThrottledUntil.Epoch = dt.Epoch
				o.ThrottledUntil.Rfc3339 = dt.Rfc3339
				o.ThrottledUntil.Offset = dt.Offset
			}
		}
	} else {
		o.ThrottledUntil.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["upgrade_date_ts"].(*int64); ok {
		o.UpgradeDateAt = val
	} else if val, ok := kv["upgrade_date_ts"].(int64); ok {
		o.UpgradeDateAt = &val
	} else {
		if val, ok := kv["upgrade_date_ts"]; ok {
			if val == nil {
				o.UpgradeDateAt = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["long"]
				}
				o.UpgradeDateAt = number.Int64Pointer(number.ToInt64Any(val))
			}
		}
	}
	if val, ok := kv["upgrade_expires_date_ts"].(*int64); ok {
		o.UpgradeExpiresDateAt = val
	} else if val, ok := kv["upgrade_expires_date_ts"].(int64); ok {
		o.UpgradeExpiresDateAt = &val
	} else {
		if val, ok := kv["upgrade_expires_date_ts"]; ok {
			if val == nil {
				o.UpgradeExpiresDateAt = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["long"]
				}
				o.UpgradeExpiresDateAt = number.Int64Pointer(number.ToInt64Any(val))
			}
		}
	}
	if val, ok := kv["upgrade_message"].(*string); ok {
		o.UpgradeMessage = val
	} else if val, ok := kv["upgrade_message"].(string); ok {
		o.UpgradeMessage = &val
	} else {
		if val, ok := kv["upgrade_message"]; ok {
			if val == nil {
				o.UpgradeMessage = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.UpgradeMessage = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["upgrade_required"].(*bool); ok {
		o.UpgradeRequired = val
	} else if val, ok := kv["upgrade_required"].(bool); ok {
		o.UpgradeRequired = &val
	} else {
		if val, ok := kv["upgrade_required"]; ok {
			if val == nil {
				o.UpgradeRequired = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.UpgradeRequired = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}

	if o == nil {

		o.Webhooks = make([]IntegrationInstanceWebhooks, 0)

	}
	if val, ok := kv["webhooks"]; ok {
		if sv, ok := val.([]IntegrationInstanceWebhooks); ok {
			o.Webhooks = sv
		} else if sp, ok := val.([]*IntegrationInstanceWebhooks); ok {
			o.Webhooks = o.Webhooks[:0]
			for _, e := range sp {
				o.Webhooks = append(o.Webhooks, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(IntegrationInstanceWebhooks); ok {
					o.Webhooks = append(o.Webhooks, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm IntegrationInstanceWebhooks
					fm.FromMap(av)
					o.Webhooks = append(o.Webhooks, fm)
				} else {
					b, _ := json.Marshal(ae)
					bkv := make(map[string]interface{})
					json.Unmarshal(b, &bkv)
					var av IntegrationInstanceWebhooks
					av.FromMap(bkv)
					o.Webhooks = append(o.Webhooks, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(IntegrationInstanceWebhooks); ok {
					o.Webhooks = append(o.Webhooks, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm IntegrationInstanceWebhooks
					fm.FromMap(r)
					o.Webhooks = append(o.Webhooks, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := IntegrationInstanceWebhooks{}
					fm.FromMap(r)
					o.Webhooks = append(o.Webhooks, fm)
				}
			}
		} else {
			arr := reflect.ValueOf(val)
			if arr.Kind() == reflect.Slice {
				for i := 0; i < arr.Len(); i++ {
					item := arr.Index(i)
					if item.CanAddr() {
						v := item.Addr().MethodByName("ToMap")
						if !v.IsNil() {
							m := v.Call([]reflect.Value{})
							var fm IntegrationInstanceWebhooks
							fm.FromMap(m[0].Interface().(map[string]interface{}))
							o.Webhooks = append(o.Webhooks, fm)
						}
					}
				}
			}
		}
	}

	o.setDefaults(false)
}
