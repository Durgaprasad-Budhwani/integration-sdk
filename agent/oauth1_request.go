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

	// Oauth1RequestTable is the default table name
	Oauth1RequestTable datamodel.ModelNameType = "agent_oauth1request"

	// Oauth1RequestModelName is the model name
	Oauth1RequestModelName datamodel.ModelNameType = "agent.Oauth1Request"
)

const (
	// Oauth1RequestModelCallbackURLColumn is the column json value callback_url
	Oauth1RequestModelCallbackURLColumn = "callback_url"
	// Oauth1RequestModelConsumerKeyColumn is the column json value consumer_key
	Oauth1RequestModelConsumerKeyColumn = "consumer_key"
	// Oauth1RequestModelConsumerSecretColumn is the column json value consumer_secret
	Oauth1RequestModelConsumerSecretColumn = "consumer_secret"
	// Oauth1RequestModelCustomerIDColumn is the column json value customer_id
	Oauth1RequestModelCustomerIDColumn = "customer_id"
	// Oauth1RequestModelEnrollmentIDColumn is the column json value enrollment_id
	Oauth1RequestModelEnrollmentIDColumn = "enrollment_id"
	// Oauth1RequestModelIDColumn is the column json value id
	Oauth1RequestModelIDColumn = "id"
	// Oauth1RequestModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	Oauth1RequestModelIntegrationInstanceIDColumn = "integration_instance_id"
	// Oauth1RequestModelPrivateKeyColumn is the column json value private_key
	Oauth1RequestModelPrivateKeyColumn = "private_key"
	// Oauth1RequestModelRefIDColumn is the column json value ref_id
	Oauth1RequestModelRefIDColumn = "ref_id"
	// Oauth1RequestModelRefTypeColumn is the column json value ref_type
	Oauth1RequestModelRefTypeColumn = "ref_type"
	// Oauth1RequestModelSessionIDColumn is the column json value session_id
	Oauth1RequestModelSessionIDColumn = "session_id"
	// Oauth1RequestModelURLColumn is the column json value url
	Oauth1RequestModelURLColumn = "url"
)

// Oauth1Request An OAuth1 request sent to the agent to fetch temporary credentials
type Oauth1Request struct {
	// CallbackURL The callback url to use.
	CallbackURL string `json:"callback_url" codec:"callback_url" bson:"callback_url" yaml:"callback_url" faker:"-"`
	// ConsumerKey The Consumer Key
	ConsumerKey string `json:"consumer_key" codec:"consumer_key" bson:"consumer_key" yaml:"consumer_key" faker:"-"`
	// ConsumerSecret The Consumer Secret
	ConsumerSecret string `json:"consumer_secret" codec:"consumer_secret" bson:"consumer_secret" yaml:"consumer_secret" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// EnrollmentID The id of the on premise agent enrollment to route this to. Leave null for cloud agent.
	EnrollmentID *string `json:"enrollment_id,omitempty" codec:"enrollment_id,omitempty" bson:"enrollment_id" yaml:"enrollment_id,omitempty" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// PrivateKey A private key for signing the oauth request.
	PrivateKey string `json:"private_key" codec:"private_key" bson:"private_key" yaml:"private_key" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// SessionID A unique session ID generated by the client which the response will include to capture it more easily
	SessionID string `json:"session_id" codec:"session_id" bson:"session_id" yaml:"session_id" faker:"-"`
	// URL The URL to the endpoint for fetching the temporary token
	URL string `json:"url" codec:"url" bson:"url" yaml:"url" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Oauth1Request)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Oauth1Request)(nil)

func toOauth1RequestObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Oauth1Request:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of Oauth1Request
func (o *Oauth1Request) String() string {
	return fmt.Sprintf("agent.Oauth1Request<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Oauth1Request) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *Oauth1Request) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Oauth1Request) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *Oauth1Request) GetModelName() datamodel.ModelNameType {
	return Oauth1RequestModelName
}

// NewOauth1RequestID provides a template for generating an ID field for Oauth1Request
func NewOauth1RequestID(customerID string, refType string, refID string) string {
	return hash.Values("Oauth1Request", customerID, refType, refID)
}

func (o *Oauth1Request) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Oauth1Request", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Oauth1Request) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Oauth1Request) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Oauth1Request) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *Oauth1Request) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Oauth1Request) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Oauth1Request) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Oauth1Request) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Oauth1Request) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *Oauth1Request) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = Oauth1RequestModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Oauth1Request) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *Oauth1Request) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of Oauth1Request
func (o *Oauth1Request) Clone() datamodel.Model {
	c := new(Oauth1Request)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Oauth1Request) Anon() datamodel.Model {
	c := new(Oauth1Request)
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
func (o *Oauth1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Oauth1Request) UnmarshalJSON(data []byte) error {
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
func (o *Oauth1Request) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *Oauth1Request) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two Oauth1Request objects are equal
func (o *Oauth1Request) IsEqual(other *Oauth1Request) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Oauth1Request) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"callback_url":            toOauth1RequestObject(o.CallbackURL, false),
		"consumer_key":            toOauth1RequestObject(o.ConsumerKey, false),
		"consumer_secret":         toOauth1RequestObject(o.ConsumerSecret, false),
		"customer_id":             toOauth1RequestObject(o.CustomerID, false),
		"enrollment_id":           toOauth1RequestObject(o.EnrollmentID, true),
		"id":                      toOauth1RequestObject(o.ID, false),
		"integration_instance_id": toOauth1RequestObject(o.IntegrationInstanceID, true),
		"private_key":             toOauth1RequestObject(o.PrivateKey, false),
		"ref_id":                  toOauth1RequestObject(o.RefID, false),
		"ref_type":                toOauth1RequestObject(o.RefType, false),
		"session_id":              toOauth1RequestObject(o.SessionID, false),
		"url":                     toOauth1RequestObject(o.URL, false),
		"hashcode":                toOauth1RequestObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Oauth1Request) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["callback_url"].(string); ok {
		o.CallbackURL = val
	} else {
		if val, ok := kv["callback_url"]; ok {
			if val == nil {
				o.CallbackURL = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.CallbackURL = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["consumer_key"].(string); ok {
		o.ConsumerKey = val
	} else {
		if val, ok := kv["consumer_key"]; ok {
			if val == nil {
				o.ConsumerKey = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.ConsumerKey = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["consumer_secret"].(string); ok {
		o.ConsumerSecret = val
	} else {
		if val, ok := kv["consumer_secret"]; ok {
			if val == nil {
				o.ConsumerSecret = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.ConsumerSecret = fmt.Sprintf("%v", val)
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
	if val, ok := kv["private_key"].(string); ok {
		o.PrivateKey = val
	} else {
		if val, ok := kv["private_key"]; ok {
			if val == nil {
				o.PrivateKey = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.PrivateKey = fmt.Sprintf("%v", val)
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
	if val, ok := kv["session_id"].(string); ok {
		o.SessionID = val
	} else {
		if val, ok := kv["session_id"]; ok {
			if val == nil {
				o.SessionID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.SessionID = fmt.Sprintf("%v", val)
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
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *Oauth1Request) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CallbackURL)
	args = append(args, o.ConsumerKey)
	args = append(args, o.ConsumerSecret)
	args = append(args, o.CustomerID)
	args = append(args, o.EnrollmentID)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.PrivateKey)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.SessionID)
	args = append(args, o.URL)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// Oauth1RequestPartial is a partial struct for upsert mutations for Oauth1Request
type Oauth1RequestPartial struct {
	// CallbackURL The callback url to use.
	CallbackURL *string `json:"callback_url,omitempty"`
	// ConsumerKey The Consumer Key
	ConsumerKey *string `json:"consumer_key,omitempty"`
	// ConsumerSecret The Consumer Secret
	ConsumerSecret *string `json:"consumer_secret,omitempty"`
	// EnrollmentID The id of the on premise agent enrollment to route this to. Leave null for cloud agent.
	EnrollmentID *string `json:"enrollment_id,omitempty"`
	// PrivateKey A private key for signing the oauth request.
	PrivateKey *string `json:"private_key,omitempty"`
	// SessionID A unique session ID generated by the client which the response will include to capture it more easily
	SessionID *string `json:"session_id,omitempty"`
	// URL The URL to the endpoint for fetching the temporary token
	URL *string `json:"url,omitempty"`
}

var _ datamodel.PartialModel = (*Oauth1RequestPartial)(nil)

// GetModelName returns the name of the model
func (o *Oauth1RequestPartial) GetModelName() datamodel.ModelNameType {
	return Oauth1RequestModelName
}

// ToMap returns the object as a map
func (o *Oauth1RequestPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"callback_url":    toOauth1RequestObject(o.CallbackURL, true),
		"consumer_key":    toOauth1RequestObject(o.ConsumerKey, true),
		"consumer_secret": toOauth1RequestObject(o.ConsumerSecret, true),
		"enrollment_id":   toOauth1RequestObject(o.EnrollmentID, true),
		"private_key":     toOauth1RequestObject(o.PrivateKey, true),
		"session_id":      toOauth1RequestObject(o.SessionID, true),
		"url":             toOauth1RequestObject(o.URL, true),
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
func (o *Oauth1RequestPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *Oauth1RequestPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Oauth1RequestPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *Oauth1RequestPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *Oauth1RequestPartial) FromMap(kv map[string]interface{}) {
	if val, ok := kv["callback_url"].(*string); ok {
		o.CallbackURL = val
	} else if val, ok := kv["callback_url"].(string); ok {
		o.CallbackURL = &val
	} else {
		if val, ok := kv["callback_url"]; ok {
			if val == nil {
				o.CallbackURL = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.CallbackURL = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["consumer_key"].(*string); ok {
		o.ConsumerKey = val
	} else if val, ok := kv["consumer_key"].(string); ok {
		o.ConsumerKey = &val
	} else {
		if val, ok := kv["consumer_key"]; ok {
			if val == nil {
				o.ConsumerKey = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.ConsumerKey = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["consumer_secret"].(*string); ok {
		o.ConsumerSecret = val
	} else if val, ok := kv["consumer_secret"].(string); ok {
		o.ConsumerSecret = &val
	} else {
		if val, ok := kv["consumer_secret"]; ok {
			if val == nil {
				o.ConsumerSecret = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.ConsumerSecret = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
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
	if val, ok := kv["session_id"].(*string); ok {
		o.SessionID = val
	} else if val, ok := kv["session_id"].(string); ok {
		o.SessionID = &val
	} else {
		if val, ok := kv["session_id"]; ok {
			if val == nil {
				o.SessionID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.SessionID = pstrings.Pointer(fmt.Sprintf("%v", val))
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
