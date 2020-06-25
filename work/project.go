// DO NOT EDIT -- generated code

// Package work - the system which contains project work
package work

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/v10/datamodel"
	"github.com/pinpt/go-common/v10/datetime"
	"github.com/pinpt/go-common/v10/hash"
	pjson "github.com/pinpt/go-common/v10/json"
	"github.com/pinpt/go-common/v10/number"
	pstrings "github.com/pinpt/go-common/v10/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

const (
	// ProjectTopic is the default topic name
	ProjectTopic datamodel.TopicNameType = "work_Project"

	// ProjectTable is the default table name
	ProjectTable datamodel.ModelNameType = "work_project"

	// ProjectModelName is the model name
	ProjectModelName datamodel.ModelNameType = "work.Project"
)

const (
	// ProjectModelActiveColumn is the column json value active
	ProjectModelActiveColumn = "active"
	// ProjectModelAffiliationColumn is the column json value affiliation
	ProjectModelAffiliationColumn = "affiliation"
	// ProjectModelCategoryColumn is the column json value category
	ProjectModelCategoryColumn = "category"
	// ProjectModelCustomerIDColumn is the column json value customer_id
	ProjectModelCustomerIDColumn = "customer_id"
	// ProjectModelDescriptionColumn is the column json value description
	ProjectModelDescriptionColumn = "description"
	// ProjectModelIDColumn is the column json value id
	ProjectModelIDColumn = "id"
	// ProjectModelIdentifierColumn is the column json value identifier
	ProjectModelIdentifierColumn = "identifier"
	// ProjectModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	ProjectModelIntegrationInstanceIDColumn = "integration_instance_id"
	// ProjectModelNameColumn is the column json value name
	ProjectModelNameColumn = "name"
	// ProjectModelRefIDColumn is the column json value ref_id
	ProjectModelRefIDColumn = "ref_id"
	// ProjectModelRefTypeColumn is the column json value ref_type
	ProjectModelRefTypeColumn = "ref_type"
	// ProjectModelUpdatedAtColumn is the column json value updated_ts
	ProjectModelUpdatedAtColumn = "updated_ts"
	// ProjectModelURLColumn is the column json value url
	ProjectModelURLColumn = "url"
	// ProjectModelVisibilityColumn is the column json value visibility
	ProjectModelVisibilityColumn = "visibility"
)

// ProjectAffiliation is the enumeration type for affiliation
type ProjectAffiliation int32

// UnmarshalBSONValue for unmarshaling value
func (v *ProjectAffiliation) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = ProjectAffiliation(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "ORGANIZATION":
			*v = ProjectAffiliation(0)
		case "USER":
			*v = ProjectAffiliation(1)
		case "THIRDPARTY":
			*v = ProjectAffiliation(2)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v ProjectAffiliation) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "ORGANIZATION":
		v = 0
	case "USER":
		v = 1
	case "THIRDPARTY":
		v = 2
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v ProjectAffiliation) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("ORGANIZATION")
	case 1:
		return json.Marshal("USER")
	case 2:
		return json.Marshal("THIRDPARTY")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for Affiliation
func (v ProjectAffiliation) String() string {
	switch int32(v) {
	case 0:
		return "ORGANIZATION"
	case 1:
		return "USER"
	case 2:
		return "THIRDPARTY"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *ProjectAffiliation) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case ProjectAffiliation:
		*v = val
	case int32:
		*v = ProjectAffiliation(int32(val))
	case int:
		*v = ProjectAffiliation(int32(val))
	case string:
		switch val {
		case "ORGANIZATION":
			*v = ProjectAffiliation(0)
		case "USER":
			*v = ProjectAffiliation(1)
		case "THIRDPARTY":
			*v = ProjectAffiliation(2)
		}
	}
	return nil
}

const (
	// ProjectAffiliationOrganization is the enumeration value for organization
	ProjectAffiliationOrganization ProjectAffiliation = 0
	// ProjectAffiliationUser is the enumeration value for user
	ProjectAffiliationUser ProjectAffiliation = 1
	// ProjectAffiliationThirdparty is the enumeration value for thirdparty
	ProjectAffiliationThirdparty ProjectAffiliation = 2
)

// ProjectVisibility is the enumeration type for visibility
type ProjectVisibility int32

// UnmarshalBSONValue for unmarshaling value
func (v *ProjectVisibility) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = ProjectVisibility(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "PRIVATE":
			*v = ProjectVisibility(0)
		case "PUBLIC":
			*v = ProjectVisibility(1)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v ProjectVisibility) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "PRIVATE":
		v = 0
	case "PUBLIC":
		v = 1
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v ProjectVisibility) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("PRIVATE")
	case 1:
		return json.Marshal("PUBLIC")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for Visibility
func (v ProjectVisibility) String() string {
	switch int32(v) {
	case 0:
		return "PRIVATE"
	case 1:
		return "PUBLIC"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *ProjectVisibility) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case ProjectVisibility:
		*v = val
	case int32:
		*v = ProjectVisibility(int32(val))
	case int:
		*v = ProjectVisibility(int32(val))
	case string:
		switch val {
		case "PRIVATE":
			*v = ProjectVisibility(0)
		case "PUBLIC":
			*v = ProjectVisibility(1)
		}
	}
	return nil
}

const (
	// ProjectVisibilityPrivate is the enumeration value for private
	ProjectVisibilityPrivate ProjectVisibility = 0
	// ProjectVisibilityPublic is the enumeration value for public
	ProjectVisibilityPublic ProjectVisibility = 1
)

// Project the project holds work
type Project struct {
	// Active the status of the project
	Active bool `json:"active" codec:"active" bson:"active" yaml:"active" faker:"-"`
	// Affiliation the affiliation to the project owner
	Affiliation ProjectAffiliation `json:"affiliation" codec:"affiliation" bson:"affiliation" yaml:"affiliation" faker:"-"`
	// Category the project category
	Category *string `json:"category,omitempty" codec:"category,omitempty" bson:"category" yaml:"category,omitempty" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Description the description of the project
	Description *string `json:"description,omitempty" codec:"description,omitempty" bson:"description" yaml:"description,omitempty" faker:"sentence"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Identifier the common identifier for the project
	Identifier string `json:"identifier" codec:"identifier" bson:"identifier" yaml:"identifier" faker:"abbreviation"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// Name the name of the project
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"project"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// URL the url to the project home page
	URL string `json:"url" codec:"url" bson:"url" yaml:"url" faker:"url"`
	// Visibility the visibility of the project
	Visibility ProjectVisibility `json:"visibility" codec:"visibility" bson:"visibility" yaml:"visibility" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Project)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Project)(nil)

func toProjectObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Project:
		return v.ToMap()

	case ProjectAffiliation:
		return v.String()

	case ProjectVisibility:
		return v.String()
	default:
		return o
	}
}

// String returns a string representation of Project
func (o *Project) String() string {
	return fmt.Sprintf("work.Project<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Project) GetTopicName() datamodel.TopicNameType {
	return ProjectTopic
}

// GetStreamName returns the name of the stream
func (o *Project) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Project) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *Project) GetModelName() datamodel.ModelNameType {
	return ProjectModelName
}

// NewProjectID provides a template for generating an ID field for Project
func NewProjectID(customerID string, refID string, refType string) string {
	return hash.Values(customerID, refID, refType)
}

func (o *Project) setDefaults(frommap bool) {

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
func (o *Project) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Project) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Project) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for Project")
}

// GetRefID returns the RefID for the object
func (o *Project) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Project) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Project) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Project) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Project) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Project) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = ProjectModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Project) GetTopicConfig() *datamodel.ModelTopicConfig {
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
		Timestamp:         "updated_ts",
		NumPartitions:     128,
		CleanupPolicy:     datamodel.CleanupPolicy("compact"),
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetCustomerID will return the customer_id
func (o *Project) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of Project
func (o *Project) Clone() datamodel.Model {
	c := new(Project)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Project) Anon() datamodel.Model {
	c := new(Project)
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
func (o *Project) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Project) UnmarshalJSON(data []byte) error {
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
func (o *Project) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *Project) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two Project objects are equal
func (o *Project) IsEqual(other *Project) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Project) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"active": toProjectObject(o.Active, false),

		"affiliation":             o.Affiliation.String(),
		"category":                toProjectObject(o.Category, true),
		"customer_id":             toProjectObject(o.CustomerID, false),
		"description":             toProjectObject(o.Description, true),
		"id":                      toProjectObject(o.ID, false),
		"identifier":              toProjectObject(o.Identifier, false),
		"integration_instance_id": toProjectObject(o.IntegrationInstanceID, true),
		"name":                    toProjectObject(o.Name, false),
		"ref_id":                  toProjectObject(o.RefID, false),
		"ref_type":                toProjectObject(o.RefType, false),
		"updated_ts":              toProjectObject(o.UpdatedAt, false),
		"url":                     toProjectObject(o.URL, false),

		"visibility": o.Visibility.String(),
		"hashcode":   toProjectObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Project) FromMap(kv map[string]interface{}) {

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
	if val, ok := kv["affiliation"].(ProjectAffiliation); ok {
		o.Affiliation = val
	} else {
		if em, ok := kv["affiliation"].(map[string]interface{}); ok {

			ev := em["work.affiliation"].(string)
			switch ev {
			case "organization", "ORGANIZATION":
				o.Affiliation = 0
			case "user", "USER":
				o.Affiliation = 1
			case "thirdparty", "THIRDPARTY":
				o.Affiliation = 2
			}
		}
		if em, ok := kv["affiliation"].(string); ok {
			switch em {
			case "organization", "ORGANIZATION":
				o.Affiliation = 0
			case "user", "USER":
				o.Affiliation = 1
			case "thirdparty", "THIRDPARTY":
				o.Affiliation = 2
			}
		}
	}
	if val, ok := kv["category"].(*string); ok {
		o.Category = val
	} else if val, ok := kv["category"].(string); ok {
		o.Category = &val
	} else {
		if val, ok := kv["category"]; ok {
			if val == nil {
				o.Category = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Category = pstrings.Pointer(fmt.Sprintf("%v", val))
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
	if val, ok := kv["description"].(*string); ok {
		o.Description = val
	} else if val, ok := kv["description"].(string); ok {
		o.Description = &val
	} else {
		if val, ok := kv["description"]; ok {
			if val == nil {
				o.Description = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Description = pstrings.Pointer(fmt.Sprintf("%v", val))
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
	if val, ok := kv["identifier"].(string); ok {
		o.Identifier = val
	} else {
		if val, ok := kv["identifier"]; ok {
			if val == nil {
				o.Identifier = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Identifier = fmt.Sprintf("%v", val)
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
	if val, ok := kv["url"].(string); ok {
		o.URL = val
	} else {
		if val, ok := kv["url"]; ok {
			if val == nil {
				o.URL = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.URL = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["visibility"].(ProjectVisibility); ok {
		o.Visibility = val
	} else {
		if em, ok := kv["visibility"].(map[string]interface{}); ok {

			ev := em["work.visibility"].(string)
			switch ev {
			case "private", "PRIVATE":
				o.Visibility = 0
			case "public", "PUBLIC":
				o.Visibility = 1
			}
		}
		if em, ok := kv["visibility"].(string); ok {
			switch em {
			case "private", "PRIVATE":
				o.Visibility = 0
			case "public", "PUBLIC":
				o.Visibility = 1
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *Project) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Active)
	args = append(args, o.Affiliation)
	args = append(args, o.Category)
	args = append(args, o.CustomerID)
	args = append(args, o.Description)
	args = append(args, o.ID)
	args = append(args, o.Identifier)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.Name)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.UpdatedAt)
	args = append(args, o.URL)
	args = append(args, o.Visibility)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}
