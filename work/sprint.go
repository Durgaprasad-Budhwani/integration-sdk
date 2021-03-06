// DO NOT EDIT -- generated code

// Package work - the system which contains project work
package work

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/v10/datamodel"
	"github.com/pinpt/go-common/v10/datetime"
	"github.com/pinpt/go-common/v10/hash"
	pjson "github.com/pinpt/go-common/v10/json"
	"github.com/pinpt/go-common/v10/number"
	"github.com/pinpt/go-common/v10/slice"
	pstrings "github.com/pinpt/go-common/v10/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// SprintTopic is the default topic name
	SprintTopic datamodel.TopicNameType = "work_Sprint"

	// SprintTable is the default table name
	SprintTable datamodel.ModelNameType = "work_sprint"

	// SprintModelName is the model name
	SprintModelName datamodel.ModelNameType = "work.Sprint"
)

const (
	// SprintModelActiveColumn is the column json value active
	SprintModelActiveColumn = "active"
	// SprintModelBoardIDColumn is the column json value board_id
	SprintModelBoardIDColumn = "board_id"
	// SprintModelBoardIdsColumn is the column json value board_ids
	SprintModelBoardIdsColumn = "board_ids"
	// SprintModelColumnsColumn is the column json value columns
	SprintModelColumnsColumn = "columns"
	// SprintModelColumnsIssueIdsColumn is the column json value issue_ids
	SprintModelColumnsIssueIdsColumn = "issue_ids"
	// SprintModelColumnsNameColumn is the column json value name
	SprintModelColumnsNameColumn = "name"
	// SprintModelCompletedDateColumn is the column json value completed_date
	SprintModelCompletedDateColumn = "completed_date"
	// SprintModelCompletedDateEpochColumn is the column json value epoch
	SprintModelCompletedDateEpochColumn = "epoch"
	// SprintModelCompletedDateOffsetColumn is the column json value offset
	SprintModelCompletedDateOffsetColumn = "offset"
	// SprintModelCompletedDateRfc3339Column is the column json value rfc3339
	SprintModelCompletedDateRfc3339Column = "rfc3339"
	// SprintModelCustomerIDColumn is the column json value customer_id
	SprintModelCustomerIDColumn = "customer_id"
	// SprintModelDeletedColumn is the column json value deleted
	SprintModelDeletedColumn = "deleted"
	// SprintModelEndedDateColumn is the column json value ended_date
	SprintModelEndedDateColumn = "ended_date"
	// SprintModelEndedDateEpochColumn is the column json value epoch
	SprintModelEndedDateEpochColumn = "epoch"
	// SprintModelEndedDateOffsetColumn is the column json value offset
	SprintModelEndedDateOffsetColumn = "offset"
	// SprintModelEndedDateRfc3339Column is the column json value rfc3339
	SprintModelEndedDateRfc3339Column = "rfc3339"
	// SprintModelGoalColumn is the column json value goal
	SprintModelGoalColumn = "goal"
	// SprintModelIDColumn is the column json value id
	SprintModelIDColumn = "id"
	// SprintModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	SprintModelIntegrationInstanceIDColumn = "integration_instance_id"
	// SprintModelIssueIdsColumn is the column json value issue_ids
	SprintModelIssueIdsColumn = "issue_ids"
	// SprintModelNameColumn is the column json value name
	SprintModelNameColumn = "name"
	// SprintModelProjectIdsColumn is the column json value project_ids
	SprintModelProjectIdsColumn = "project_ids"
	// SprintModelRefIDColumn is the column json value ref_id
	SprintModelRefIDColumn = "ref_id"
	// SprintModelRefTypeColumn is the column json value ref_type
	SprintModelRefTypeColumn = "ref_type"
	// SprintModelStartedDateColumn is the column json value started_date
	SprintModelStartedDateColumn = "started_date"
	// SprintModelStartedDateEpochColumn is the column json value epoch
	SprintModelStartedDateEpochColumn = "epoch"
	// SprintModelStartedDateOffsetColumn is the column json value offset
	SprintModelStartedDateOffsetColumn = "offset"
	// SprintModelStartedDateRfc3339Column is the column json value rfc3339
	SprintModelStartedDateRfc3339Column = "rfc3339"
	// SprintModelStatusColumn is the column json value status
	SprintModelStatusColumn = "status"
	// SprintModelUpdatedDateColumn is the column json value updated_date
	SprintModelUpdatedDateColumn = "updated_date"
	// SprintModelUpdatedDateEpochColumn is the column json value epoch
	SprintModelUpdatedDateEpochColumn = "epoch"
	// SprintModelUpdatedDateOffsetColumn is the column json value offset
	SprintModelUpdatedDateOffsetColumn = "offset"
	// SprintModelUpdatedDateRfc3339Column is the column json value rfc3339
	SprintModelUpdatedDateRfc3339Column = "rfc3339"
	// SprintModelUpdatedAtColumn is the column json value updated_ts
	SprintModelUpdatedAtColumn = "updated_ts"
	// SprintModelURLColumn is the column json value url
	SprintModelURLColumn = "url"
)

// SprintColumns represents the object structure for columns
type SprintColumns struct {
	// IssueIds the issue ids for the column
	IssueIds []string `json:"issue_ids" codec:"issue_ids" bson:"issue_ids" yaml:"issue_ids" faker:"-"`
	// Name the name of the column
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
}

func toSprintColumnsObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *SprintColumns:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *SprintColumns) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// IssueIds the issue ids for the column
		"issue_ids": toSprintColumnsObject(o.IssueIds, false),
		// Name the name of the column
		"name": toSprintColumnsObject(o.Name, false),
	}
}

func (o *SprintColumns) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *SprintColumns) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["issue_ids"]; ok {
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
								panic("unsupported type for issue_ids field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for issue_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for issue_ids field")
				}
			}
			o.IssueIds = na
		}
	}
	if o.IssueIds == nil {
		o.IssueIds = make([]string, 0)
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
	o.setDefaults(false)
}

// SprintCompletedDate represents the object structure for completed_date
type SprintCompletedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toSprintCompletedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *SprintCompletedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *SprintCompletedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toSprintCompletedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toSprintCompletedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toSprintCompletedDateObject(o.Rfc3339, false),
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

// SprintEndedDate represents the object structure for ended_date
type SprintEndedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toSprintEndedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *SprintEndedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *SprintEndedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toSprintEndedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toSprintEndedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toSprintEndedDateObject(o.Rfc3339, false),
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

// SprintStartedDate represents the object structure for started_date
type SprintStartedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toSprintStartedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *SprintStartedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *SprintStartedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toSprintStartedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toSprintStartedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toSprintStartedDateObject(o.Rfc3339, false),
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

// SprintStatus is the enumeration type for status
type SprintStatus int32

// toSprintStatusPointer is the enumeration pointer type for status
func toSprintStatusPointer(v int32) *SprintStatus {
	nv := SprintStatus(v)
	return &nv
}

// toSprintStatusEnum is the enumeration pointer wrapper for status
func toSprintStatusEnum(v *SprintStatus) string {
	if v == nil {
		return toSprintStatusPointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *SprintStatus) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = SprintStatus(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "ACTIVE":
			*v = SprintStatus(0)
		case "FUTURE":
			*v = SprintStatus(1)
		case "CLOSED":
			*v = SprintStatus(2)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v *SprintStatus) UnmarshalJSON(buf []byte) error {
	var val string
	if err := json.Unmarshal(buf, &val); err != nil {
		return err
	}
	switch val {
	case "ACTIVE":
		*v = 0
	case "FUTURE":
		*v = 1
	case "CLOSED":
		*v = 2
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v SprintStatus) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("ACTIVE")
	case 1:
		return json.Marshal("FUTURE")
	case 2:
		return json.Marshal("CLOSED")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

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

// FromInterface for decoding from an interface
func (v *SprintStatus) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case SprintStatus:
		*v = val
	case int32:
		*v = SprintStatus(int32(val))
	case int:
		*v = SprintStatus(int32(val))
	case string:
		switch val {
		case "ACTIVE":
			*v = SprintStatus(0)
		case "FUTURE":
			*v = SprintStatus(1)
		case "CLOSED":
			*v = SprintStatus(2)
		}
	}
	return nil
}

const (
	// SprintStatusActive is the enumeration value for active
	SprintStatusActive SprintStatus = 0
	// SprintStatusFuture is the enumeration value for future
	SprintStatusFuture SprintStatus = 1
	// SprintStatusClosed is the enumeration value for closed
	SprintStatusClosed SprintStatus = 2
)

// SprintUpdatedDate represents the object structure for updated_date
type SprintUpdatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toSprintUpdatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *SprintUpdatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *SprintUpdatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toSprintUpdatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toSprintUpdatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toSprintUpdatedDateObject(o.Rfc3339, false),
	}
}

func (o *SprintUpdatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *SprintUpdatedDate) FromMap(kv map[string]interface{}) {

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

// Sprint sprint details
type Sprint struct {
	// Active indicates that this model is displayed in a source system, false if the model is deleted
	Active bool `json:"active" codec:"active" bson:"active" yaml:"active" faker:"-"`
	// BoardID the board that this kanban board is related to
	//
	// Deprecated: this field needs to be an array, added board_ids instead
	BoardID *string `json:"board_id,omitempty" codec:"board_id,omitempty" bson:"board_id" yaml:"board_id,omitempty" faker:"-"`
	// BoardIds the boards that this kanban board is related to
	BoardIds []string `json:"board_ids" codec:"board_ids" bson:"board_ids" yaml:"board_ids" faker:"-"`
	// Columns the columns for the board in the order that they are displayed and the issue ids for each
	Columns []SprintColumns `json:"columns" codec:"columns" bson:"columns" yaml:"columns" faker:"-"`
	// CompletedDate the date that the sprint was completed
	CompletedDate SprintCompletedDate `json:"completed_date" codec:"completed_date" bson:"completed_date" yaml:"completed_date" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Deleted indicates that this board was deleted by the user and shouldn't be shown in the app
	Deleted bool `json:"deleted" codec:"deleted" bson:"deleted" yaml:"deleted" faker:"-"`
	// EndedDate the date that the sprint was ended
	EndedDate SprintEndedDate `json:"ended_date" codec:"ended_date" bson:"ended_date" yaml:"ended_date" faker:"-"`
	// Goal the goal for the sprint
	Goal string `json:"goal" codec:"goal" bson:"goal" yaml:"goal" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// IssueIds ids of the all issues in the sprint in rank order
	IssueIds []string `json:"issue_ids" codec:"issue_ids" bson:"issue_ids" yaml:"issue_ids" faker:"-"`
	// Name the name of the field
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// ProjectIds ids of the projects used in this board
	ProjectIds []string `json:"project_ids" codec:"project_ids" bson:"project_ids" yaml:"project_ids" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// StartedDate the date that the sprint was started
	StartedDate SprintStartedDate `json:"started_date" codec:"started_date" bson:"started_date" yaml:"started_date" faker:"-"`
	// Status status of the sprint
	Status SprintStatus `json:"status" codec:"status" bson:"status" yaml:"status" faker:"-"`
	// UpdatedDate the timestamp that the sprint was updated
	UpdatedDate SprintUpdatedDate `json:"updated_date" codec:"updated_date" bson:"updated_date" yaml:"updated_date" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// URL the url to the sprint board
	URL *string `json:"url,omitempty" codec:"url,omitempty" bson:"url" yaml:"url,omitempty" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Sprint)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Sprint)(nil)

func toSprintObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Sprint:
		return v.ToMap()

	case []SprintColumns:
		arr := make([]interface{}, 0)
		for _, i := range v {
			arr = append(arr, i.ToMap())
		}
		return arr

	case SprintCompletedDate:
		return v.ToMap()

	case SprintEndedDate:
		return v.ToMap()

	case SprintStartedDate:
		return v.ToMap()

	case SprintStatus:
		return v.String()

	case SprintUpdatedDate:
		return v.ToMap()

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

// GetStreamName returns the name of the stream
func (o *Sprint) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Sprint) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *Sprint) GetModelName() datamodel.ModelNameType {
	return SprintModelName
}

// NewSprintID provides a template for generating an ID field for Sprint
func NewSprintID(customerID string, refID string, refType string) string {
	return hash.Values(customerID, refID, refType)
}

func (o *Sprint) setDefaults(frommap bool) {
	if o.BoardIds == nil {
		o.BoardIds = make([]string, 0)
	}
	if o.Columns == nil {
		o.Columns = make([]SprintColumns, 0)
	}
	if o.IssueIds == nil {
		o.IssueIds = make([]string, 0)
	}
	if o.ProjectIds == nil {
		o.ProjectIds = make([]string, 0)
	}

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
	var i interface{} = o.ID
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

// IsMutable returns true if the model is mutable
func (o *Sprint) IsMutable() bool {
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
		Key:               "id",
		Timestamp:         "started_date",
		NumPartitions:     128,
		CleanupPolicy:     datamodel.CleanupPolicy("compact"),
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetCustomerID will return the customer_id
func (o *Sprint) GetCustomerID() string {
	return o.CustomerID
}

// SetCustomerID will return the customer_id
func (o *Sprint) SetCustomerID(id string) {
	o.CustomerID = id
}

// GetRefType will return the ref_type
func (o *Sprint) GetRefType() string {
	return o.RefType
}

// SetRefType will return the ref_type
func (o *Sprint) SetRefType(t string) {
	o.RefType = t
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

// Stringify returns the object in JSON format as a string
func (o *Sprint) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *Sprint) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two Sprint objects are equal
func (o *Sprint) IsEqual(other *Sprint) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Sprint) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"active":                  toSprintObject(o.Active, false),
		"board_id":                toSprintObject(o.BoardID, true),
		"board_ids":               toSprintObject(o.BoardIds, false),
		"columns":                 toSprintObject(o.Columns, false),
		"completed_date":          toSprintObject(o.CompletedDate, false),
		"customer_id":             toSprintObject(o.CustomerID, false),
		"deleted":                 toSprintObject(o.Deleted, false),
		"ended_date":              toSprintObject(o.EndedDate, false),
		"goal":                    toSprintObject(o.Goal, false),
		"id":                      toSprintObject(o.ID, false),
		"integration_instance_id": toSprintObject(o.IntegrationInstanceID, true),
		"issue_ids":               toSprintObject(o.IssueIds, false),
		"name":                    toSprintObject(o.Name, false),
		"project_ids":             toSprintObject(o.ProjectIds, false),
		"ref_id":                  toSprintObject(o.RefID, false),
		"ref_type":                toSprintObject(o.RefType, false),
		"started_date":            toSprintObject(o.StartedDate, false),

		"status":       o.Status.String(),
		"updated_date": toSprintObject(o.UpdatedDate, false),
		"updated_ts":   toSprintObject(o.UpdatedAt, false),
		"url":          toSprintObject(o.URL, true),
		"hashcode":     toSprintObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Sprint) FromMap(kv map[string]interface{}) {

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
	// Deprecated
	if val, ok := kv["board_id"].(*string); ok {
		o.BoardID = val
	} else if val, ok := kv["board_id"].(string); ok {
		o.BoardID = &val
	} else {
		if val, ok := kv["board_id"]; ok {
			if val == nil {
				o.BoardID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.BoardID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["board_ids"]; ok {
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
								panic("unsupported type for board_ids field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for board_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for board_ids field")
				}
			}
			o.BoardIds = na
		}
	}
	if o.BoardIds == nil {
		o.BoardIds = make([]string, 0)
	}

	if o == nil {

		o.Columns = make([]SprintColumns, 0)

	}
	if val, ok := kv["columns"]; ok {
		if sv, ok := val.([]SprintColumns); ok {
			o.Columns = sv
		} else if sp, ok := val.([]*SprintColumns); ok {
			o.Columns = o.Columns[:0]
			for _, e := range sp {
				o.Columns = append(o.Columns, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(SprintColumns); ok {
					o.Columns = append(o.Columns, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm SprintColumns
					fm.FromMap(av)
					o.Columns = append(o.Columns, fm)
				} else {
					b, _ := json.Marshal(ae)
					bkv := make(map[string]interface{})
					json.Unmarshal(b, &bkv)
					var av SprintColumns
					av.FromMap(bkv)
					o.Columns = append(o.Columns, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(SprintColumns); ok {
					o.Columns = append(o.Columns, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm SprintColumns
					fm.FromMap(r)
					o.Columns = append(o.Columns, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := SprintColumns{}
					fm.FromMap(r)
					o.Columns = append(o.Columns, fm)
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
							var fm SprintColumns
							fm.FromMap(m[0].Interface().(map[string]interface{}))
							o.Columns = append(o.Columns, fm)
						}
					}
				}
			}
		}
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
			dt := datetime.NewDateWithTime(tv)
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
			dt := datetime.NewDateWithTime(tv)
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

	if val, ok := kv["goal"].(string); ok {
		o.Goal = val
	} else {
		if val, ok := kv["goal"]; ok {
			if val == nil {
				o.Goal = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
	if val, ok := kv["issue_ids"]; ok {
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
								panic("unsupported type for issue_ids field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for issue_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for issue_ids field")
				}
			}
			o.IssueIds = na
		}
	}
	if o.IssueIds == nil {
		o.IssueIds = make([]string, 0)
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
	if val, ok := kv["project_ids"]; ok {
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
								panic("unsupported type for project_ids field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for project_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for project_ids field")
				}
			}
			o.ProjectIds = na
		}
	}
	if o.ProjectIds == nil {
		o.ProjectIds = make([]string, 0)
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
			dt := datetime.NewDateWithTime(tv)
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

	if val, ok := kv["updated_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.UpdatedDate.FromMap(kv)
		} else if sv, ok := val.(SprintUpdatedDate); ok {
			// struct
			o.UpdatedDate = sv
		} else if sp, ok := val.(*SprintUpdatedDate); ok {
			// struct pointer
			o.UpdatedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.UpdatedDate.Epoch = dt.Epoch
			o.UpdatedDate.Rfc3339 = dt.Rfc3339
			o.UpdatedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.UpdatedDate.Epoch = dt.Epoch
			o.UpdatedDate.Rfc3339 = dt.Rfc3339
			o.UpdatedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.UpdatedDate.Epoch = dt.Epoch
				o.UpdatedDate.Rfc3339 = dt.Rfc3339
				o.UpdatedDate.Offset = dt.Offset
			}
		}
	} else {
		o.UpdatedDate.FromMap(map[string]interface{}{})
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

// Hash will return a hashcode for the object
func (o *Sprint) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Active)
	args = append(args, o.BoardID)
	args = append(args, o.BoardIds)
	args = append(args, o.Columns)
	args = append(args, o.CompletedDate)
	args = append(args, o.CustomerID)
	args = append(args, o.Deleted)
	args = append(args, o.EndedDate)
	args = append(args, o.Goal)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.IssueIds)
	args = append(args, o.Name)
	args = append(args, o.ProjectIds)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.StartedDate)
	args = append(args, o.Status)
	args = append(args, o.UpdatedDate)
	args = append(args, o.UpdatedAt)
	args = append(args, o.URL)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *Sprint) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *Sprint) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// SprintPartial is a partial struct for upsert mutations for Sprint
type SprintPartial struct {
	// Active indicates that this model is displayed in a source system, false if the model is deleted
	Active *bool `json:"active,omitempty"`
	// BoardID the board that this kanban board is related to
	//
	// Deprecated: this field needs to be an array, added board_ids instead
	BoardID *string `json:"board_id,omitempty"`
	// BoardIds the boards that this kanban board is related to
	BoardIds []string `json:"board_ids,omitempty"`
	// Columns the columns for the board in the order that they are displayed and the issue ids for each
	Columns []SprintColumns `json:"columns,omitempty"`
	// CompletedDate the date that the sprint was completed
	CompletedDate *SprintCompletedDate `json:"completed_date,omitempty"`
	// Deleted indicates that this board was deleted by the user and shouldn't be shown in the app
	Deleted *bool `json:"deleted,omitempty"`
	// EndedDate the date that the sprint was ended
	EndedDate *SprintEndedDate `json:"ended_date,omitempty"`
	// Goal the goal for the sprint
	Goal *string `json:"goal,omitempty"`
	// IssueIds ids of the all issues in the sprint in rank order
	IssueIds []string `json:"issue_ids,omitempty"`
	// Name the name of the field
	Name *string `json:"name,omitempty"`
	// ProjectIds ids of the projects used in this board
	ProjectIds []string `json:"project_ids,omitempty"`
	// StartedDate the date that the sprint was started
	StartedDate *SprintStartedDate `json:"started_date,omitempty"`
	// Status status of the sprint
	Status *SprintStatus `json:"status,omitempty"`
	// UpdatedDate the timestamp that the sprint was updated
	UpdatedDate *SprintUpdatedDate `json:"updated_date,omitempty"`
	// URL the url to the sprint board
	URL *string `json:"url,omitempty"`
}

var _ datamodel.PartialModel = (*SprintPartial)(nil)

// GetModelName returns the name of the model
func (o *SprintPartial) GetModelName() datamodel.ModelNameType {
	return SprintModelName
}

// ToMap returns the object as a map
func (o *SprintPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"active":         toSprintObject(o.Active, true),
		"board_id":       toSprintObject(o.BoardID, true),
		"board_ids":      toSprintObject(o.BoardIds, true),
		"columns":        toSprintObject(o.Columns, true),
		"completed_date": toSprintObject(o.CompletedDate, true),
		"deleted":        toSprintObject(o.Deleted, true),
		"ended_date":     toSprintObject(o.EndedDate, true),
		"goal":           toSprintObject(o.Goal, true),
		"issue_ids":      toSprintObject(o.IssueIds, true),
		"name":           toSprintObject(o.Name, true),
		"project_ids":    toSprintObject(o.ProjectIds, true),
		"started_date":   toSprintObject(o.StartedDate, true),

		"status":       toSprintStatusEnum(o.Status),
		"updated_date": toSprintObject(o.UpdatedDate, true),
		"url":          toSprintObject(o.URL, true),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {

			if k == "board_ids" {
				if arr, ok := v.([]string); ok {
					if len(arr) == 0 {
						delete(kv, k)
					}
				}
			}

			if k == "columns" {
				if arr, ok := v.([]SprintColumns); ok {
					if len(arr) == 0 {
						delete(kv, k)
					}
				}
			}
			if k == "completed_date" {
				if dt, ok := v.(*SprintCompletedDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
			if k == "ended_date" {
				if dt, ok := v.(*SprintEndedDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}

			if k == "issue_ids" {
				if arr, ok := v.([]string); ok {
					if len(arr) == 0 {
						delete(kv, k)
					}
				}
			}

			if k == "project_ids" {
				if arr, ok := v.([]string); ok {
					if len(arr) == 0 {
						delete(kv, k)
					}
				}
			}
			if k == "started_date" {
				if dt, ok := v.(*SprintStartedDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
			if k == "updated_date" {
				if dt, ok := v.(*SprintUpdatedDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
		}
	}
	return kv
}

// Stringify returns the object in JSON format as a string
func (o *SprintPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *SprintPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *SprintPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *SprintPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *SprintPartial) FromMap(kv map[string]interface{}) {
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
	// Deprecated
	if val, ok := kv["board_id"].(*string); ok {
		o.BoardID = val
	} else if val, ok := kv["board_id"].(string); ok {
		o.BoardID = &val
	} else {
		if val, ok := kv["board_id"]; ok {
			if val == nil {
				o.BoardID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.BoardID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["board_ids"]; ok {
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
								panic("unsupported type for board_ids field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for board_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for board_ids field")
				}
			}
			o.BoardIds = na
		}
	}
	if o.BoardIds == nil {
		o.BoardIds = make([]string, 0)
	}

	if o == nil {

		o.Columns = make([]SprintColumns, 0)

	}
	if val, ok := kv["columns"]; ok {
		if sv, ok := val.([]SprintColumns); ok {
			o.Columns = sv
		} else if sp, ok := val.([]*SprintColumns); ok {
			o.Columns = o.Columns[:0]
			for _, e := range sp {
				o.Columns = append(o.Columns, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(SprintColumns); ok {
					o.Columns = append(o.Columns, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm SprintColumns
					fm.FromMap(av)
					o.Columns = append(o.Columns, fm)
				} else {
					b, _ := json.Marshal(ae)
					bkv := make(map[string]interface{})
					json.Unmarshal(b, &bkv)
					var av SprintColumns
					av.FromMap(bkv)
					o.Columns = append(o.Columns, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(SprintColumns); ok {
					o.Columns = append(o.Columns, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm SprintColumns
					fm.FromMap(r)
					o.Columns = append(o.Columns, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := SprintColumns{}
					fm.FromMap(r)
					o.Columns = append(o.Columns, fm)
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
							var fm SprintColumns
							fm.FromMap(m[0].Interface().(map[string]interface{}))
							o.Columns = append(o.Columns, fm)
						}
					}
				}
			}
		}
	}

	if o.CompletedDate == nil {
		o.CompletedDate = &SprintCompletedDate{}
	}

	if val, ok := kv["completed_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CompletedDate.FromMap(kv)
		} else if sv, ok := val.(SprintCompletedDate); ok {
			// struct
			o.CompletedDate = &sv
		} else if sp, ok := val.(*SprintCompletedDate); ok {
			// struct pointer
			o.CompletedDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.CompletedDate.Epoch = dt.Epoch
			o.CompletedDate.Rfc3339 = dt.Rfc3339
			o.CompletedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
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

	if o.EndedDate == nil {
		o.EndedDate = &SprintEndedDate{}
	}

	if val, ok := kv["ended_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.EndedDate.FromMap(kv)
		} else if sv, ok := val.(SprintEndedDate); ok {
			// struct
			o.EndedDate = &sv
		} else if sp, ok := val.(*SprintEndedDate); ok {
			// struct pointer
			o.EndedDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.EndedDate.Epoch = dt.Epoch
			o.EndedDate.Rfc3339 = dt.Rfc3339
			o.EndedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
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

	if val, ok := kv["goal"].(*string); ok {
		o.Goal = val
	} else if val, ok := kv["goal"].(string); ok {
		o.Goal = &val
	} else {
		if val, ok := kv["goal"]; ok {
			if val == nil {
				o.Goal = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Goal = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["issue_ids"]; ok {
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
								panic("unsupported type for issue_ids field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for issue_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for issue_ids field")
				}
			}
			o.IssueIds = na
		}
	}
	if o.IssueIds == nil {
		o.IssueIds = make([]string, 0)
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
	if val, ok := kv["project_ids"]; ok {
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
								panic("unsupported type for project_ids field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for project_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for project_ids field")
				}
			}
			o.ProjectIds = na
		}
	}
	if o.ProjectIds == nil {
		o.ProjectIds = make([]string, 0)
	}

	if o.StartedDate == nil {
		o.StartedDate = &SprintStartedDate{}
	}

	if val, ok := kv["started_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.StartedDate.FromMap(kv)
		} else if sv, ok := val.(SprintStartedDate); ok {
			// struct
			o.StartedDate = &sv
		} else if sp, ok := val.(*SprintStartedDate); ok {
			// struct pointer
			o.StartedDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.StartedDate.Epoch = dt.Epoch
			o.StartedDate.Rfc3339 = dt.Rfc3339
			o.StartedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
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

	if val, ok := kv["status"].(*SprintStatus); ok {
		o.Status = val
	} else if val, ok := kv["status"].(SprintStatus); ok {
		o.Status = &val
	} else {
		if val, ok := kv["status"]; ok {
			if val == nil {
				o.Status = toSprintStatusPointer(0)
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["SprintStatus"]
				}
				// this is an enum pointer
				if em, ok := val.(string); ok {
					switch em {
					case "active", "ACTIVE":
						o.Status = toSprintStatusPointer(0)
					case "future", "FUTURE":
						o.Status = toSprintStatusPointer(1)
					case "closed", "CLOSED":
						o.Status = toSprintStatusPointer(2)
					}
				}
			}
		}
	}

	if o.UpdatedDate == nil {
		o.UpdatedDate = &SprintUpdatedDate{}
	}

	if val, ok := kv["updated_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.UpdatedDate.FromMap(kv)
		} else if sv, ok := val.(SprintUpdatedDate); ok {
			// struct
			o.UpdatedDate = &sv
		} else if sp, ok := val.(*SprintUpdatedDate); ok {
			// struct pointer
			o.UpdatedDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.UpdatedDate.Epoch = dt.Epoch
			o.UpdatedDate.Rfc3339 = dt.Rfc3339
			o.UpdatedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.UpdatedDate.Epoch = dt.Epoch
			o.UpdatedDate.Rfc3339 = dt.Rfc3339
			o.UpdatedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.UpdatedDate.Epoch = dt.Epoch
				o.UpdatedDate.Rfc3339 = dt.Rfc3339
				o.UpdatedDate.Offset = dt.Offset
			}
		}
	} else {
		o.UpdatedDate.FromMap(map[string]interface{}{})
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
