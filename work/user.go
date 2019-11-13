// DO NOT EDIT -- generated code

// Package work - the system which contains project work
package work

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
)

const (
	// UserTopic is the default topic name
	UserTopic datamodel.TopicNameType = "work_User_topic"

	// UserTable is the default table name
	UserTable datamodel.ModelNameType = "work_user"

	// UserModelName is the model name
	UserModelName datamodel.ModelNameType = "work.User"
)

const (
	// UserAssociatedRefIDColumn is the associated_ref_id column name
	UserAssociatedRefIDColumn = "AssociatedRefID"
	// UserAvatarURLColumn is the avatar_url column name
	UserAvatarURLColumn = "AvatarURL"
	// UserCustomerIDColumn is the customer_id column name
	UserCustomerIDColumn = "CustomerID"
	// UserEmailColumn is the email column name
	UserEmailColumn = "Email"
	// UserIDColumn is the id column name
	UserIDColumn = "ID"
	// UserNameColumn is the name column name
	UserNameColumn = "Name"
	// UserRefIDColumn is the ref_id column name
	UserRefIDColumn = "RefID"
	// UserRefTypeColumn is the ref_type column name
	UserRefTypeColumn = "RefType"
	// UserUpdatedAtColumn is the updated_ts column name
	UserUpdatedAtColumn = "UpdatedAt"
	// UserUsernameColumn is the username column name
	UserUsernameColumn = "Username"
)

// User the work user
type User struct {
	// AssociatedRefID the ref id associated for this user in another system
	AssociatedRefID *string `json:"associated_ref_id,omitempty" codec:"associated_ref_id,omitempty" bson:"associated_ref_id" yaml:"associated_ref_id,omitempty" faker:"-"`
	// AvatarURL the url to users avatar
	AvatarURL *string `json:"avatar_url,omitempty" codec:"avatar_url,omitempty" bson:"avatar_url" yaml:"avatar_url,omitempty" faker:"avatar"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Email the email for the user
	Email *string `json:"email,omitempty" codec:"email,omitempty" bson:"email" yaml:"email,omitempty" faker:"email"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Name the name of the user
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"person"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Username the username of the user
	Username string `json:"username" codec:"username" bson:"username" yaml:"username" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*User)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*User)(nil)

func toUserObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *User:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of User
func (o *User) String() string {
	return fmt.Sprintf("work.User<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *User) GetTopicName() datamodel.TopicNameType {
	return UserTopic
}

// GetStreamName returns the name of the stream
func (o *User) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *User) GetTableName() string {
	return UserTable.String()
}

// GetModelName returns the name of the model
func (o *User) GetModelName() datamodel.ModelNameType {
	return UserModelName
}

// NewUserID provides a template for generating an ID field for User
func NewUserID(customerID string, refType string, refID string) string {
	return hash.Values("User", customerID, refType, refID)
}

func (o *User) setDefaults(frommap bool) {
	if o.AssociatedRefID == nil {
		o.AssociatedRefID = pstrings.Pointer("")
	}
	if o.AvatarURL == nil {
		o.AvatarURL = pstrings.Pointer("")
	}
	if o.Email == nil {
		o.Email = pstrings.Pointer("")
	}

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("User", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *User) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *User) GetTopicKey() string {
	var i interface{} = o.CustomerID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *User) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for User")
}

// GetRefID returns the RefID for the object
func (o *User) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *User) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *User) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *User) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *User) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *User) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = UserModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *User) GetTopicConfig() *datamodel.ModelTopicConfig {
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
		Timestamp:         "updated_ts",
		NumPartitions:     8,
		CleanupPolicy:     datamodel.CleanupPolicy("compact"),
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetCustomerID will return the customer_id
func (o *User) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of User
func (o *User) Clone() datamodel.Model {
	c := new(User)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *User) Anon() datamodel.Model {
	c := new(User)
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
func (o *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *User) UnmarshalJSON(data []byte) error {
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
func (o *User) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two User objects are equal
func (o *User) IsEqual(other *User) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *User) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"associated_ref_id": toUserObject(o.AssociatedRefID, true),
		"avatar_url":        toUserObject(o.AvatarURL, true),
		"customer_id":       toUserObject(o.CustomerID, false),
		"email":             toUserObject(o.Email, true),
		"id":                toUserObject(o.ID, false),
		"name":              toUserObject(o.Name, false),
		"ref_id":            toUserObject(o.RefID, false),
		"ref_type":          toUserObject(o.RefType, false),
		"updated_ts":        toUserObject(o.UpdatedAt, false),
		"username":          toUserObject(o.Username, false),
		"hashcode":          toUserObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *User) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["associated_ref_id"].(*string); ok {
		o.AssociatedRefID = val
	} else if val, ok := kv["associated_ref_id"].(string); ok {
		o.AssociatedRefID = &val
	} else {
		if val, ok := kv["associated_ref_id"]; ok {
			if val == nil {
				o.AssociatedRefID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.AssociatedRefID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["avatar_url"].(*string); ok {
		o.AvatarURL = val
	} else if val, ok := kv["avatar_url"].(string); ok {
		o.AvatarURL = &val
	} else {
		if val, ok := kv["avatar_url"]; ok {
			if val == nil {
				o.AvatarURL = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.AvatarURL = pstrings.Pointer(fmt.Sprintf("%v", val))
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.CustomerID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["email"].(*string); ok {
		o.Email = val
	} else if val, ok := kv["email"].(string); ok {
		o.Email = &val
	} else {
		if val, ok := kv["email"]; ok {
			if val == nil {
				o.Email = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Email = pstrings.Pointer(fmt.Sprintf("%v", val))
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

	if val, ok := kv["username"].(string); ok {
		o.Username = val
	} else {
		if val, ok := kv["username"]; ok {
			if val == nil {
				o.Username = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Username = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *User) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.AssociatedRefID)
	args = append(args, o.AvatarURL)
	args = append(args, o.CustomerID)
	args = append(args, o.Email)
	args = append(args, o.ID)
	args = append(args, o.Name)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.UpdatedAt)
	args = append(args, o.Username)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *User) GetEventAPIConfig() datamodel.EventAPIConfig {
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
