// DO NOT EDIT -- generated code

// Package agent - the agent communicates with pinpoint cloud to send integration data
package agent

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sync"
	"time"

	"github.com/bxcodec/faker"
	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/eventing"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	pstrings "github.com/pinpt/go-common/strings"
)

const (
	// CodequalityTriggerTopic is the default topic name
	CodequalityTriggerTopic datamodel.TopicNameType = "agent_CodequalityTrigger_topic"

	// CodequalityTriggerStream is the default stream name
	CodequalityTriggerStream datamodel.TopicNameType = "agent_CodequalityTrigger_stream"

	// CodequalityTriggerTable is the default table name
	CodequalityTriggerTable datamodel.TopicNameType = "agent_codequalitytrigger"

	// CodequalityTriggerModelName is the model name
	CodequalityTriggerModelName datamodel.ModelNameType = "agent.CodequalityTrigger"
)

const (
	// CodequalityTriggerCustomerIDColumn is the customer_id column name
	CodequalityTriggerCustomerIDColumn = "customer_id"
	// CodequalityTriggerIDColumn is the id column name
	CodequalityTriggerIDColumn = "id"
	// CodequalityTriggerIntegrationIDColumn is the integration_id column name
	CodequalityTriggerIntegrationIDColumn = "integration_id"
	// CodequalityTriggerRefIDColumn is the ref_id column name
	CodequalityTriggerRefIDColumn = "ref_id"
	// CodequalityTriggerRefTypeColumn is the ref_type column name
	CodequalityTriggerRefTypeColumn = "ref_type"
	// CodequalityTriggerUpdatedAtColumn is the updated_ts column name
	CodequalityTriggerUpdatedAtColumn = "updated_ts"
)

// CodequalityTrigger used to trigger an agent.CodequalityRequest
type CodequalityTrigger struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationID the integration id
	IntegrationID string `json:"integration_id" bson:"integration_id" yaml:"integration_id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*CodequalityTrigger)(nil)

func toCodequalityTriggerObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toCodequalityTriggerObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *CodequalityTrigger:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

// String returns a string representation of CodequalityTrigger
func (o *CodequalityTrigger) String() string {
	return fmt.Sprintf("agent.CodequalityTrigger<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *CodequalityTrigger) GetTopicName() datamodel.TopicNameType {
	return CodequalityTriggerTopic
}

// GetModelName returns the name of the model
func (o *CodequalityTrigger) GetModelName() datamodel.ModelNameType {
	return CodequalityTriggerModelName
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
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *CodequalityTrigger) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for CodequalityTrigger")
}

// GetRefID returns the RefID for the object
func (o *CodequalityTrigger) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *CodequalityTrigger) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *CodequalityTrigger) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *CodequalityTrigger) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *CodequalityTrigger) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = CodequalityTriggerModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *CodequalityTrigger) GetTopicConfig() *datamodel.ModelTopicConfig {
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
		CleanupPolicy:     datamodel.CleanupPolicy("compact"),
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetStateKey returns a key for use in state store
func (o *CodequalityTrigger) GetStateKey() string {
	key := "id"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *CodequalityTrigger) GetCustomerID() string {

	return o.CustomerID

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

// MarshalBinary returns the bytes for marshaling to binary
func (o *CodequalityTrigger) MarshalBinary() ([]byte, error) {
	return o.MarshalJSON()
}

func (o *CodequalityTrigger) UnmarshalBinary(data []byte) error {
	return o.UnmarshalJSON(data)
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

var cachedCodecCodequalityTrigger *goavro.Codec
var cachedCodecCodequalityTriggerLock sync.Mutex

// GetAvroCodec returns the avro codec for this model
func (o *CodequalityTrigger) GetAvroCodec() *goavro.Codec {
	cachedCodecCodequalityTriggerLock.Lock()
	if cachedCodecCodequalityTrigger == nil {
		c, err := GetCodequalityTriggerAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecCodequalityTrigger = c
	}
	cachedCodecCodequalityTriggerLock.Unlock()
	return cachedCodecCodequalityTrigger
}

// ToAvroBinary returns the data as Avro binary data
func (o *CodequalityTrigger) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *CodequalityTrigger) FromAvroBinary(value []byte) error {
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
func (o *CodequalityTrigger) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two CodequalityTrigger objects are equal
func (o *CodequalityTrigger) IsEqual(other *CodequalityTrigger) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *CodequalityTrigger) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id":    toCodequalityTriggerObject(o.CustomerID, isavro, false, "string"),
		"id":             toCodequalityTriggerObject(o.ID, isavro, false, "string"),
		"integration_id": toCodequalityTriggerObject(o.IntegrationID, isavro, false, "string"),
		"ref_id":         toCodequalityTriggerObject(o.RefID, isavro, false, "string"),
		"ref_type":       toCodequalityTriggerObject(o.RefType, isavro, false, "string"),
		"updated_ts":     toCodequalityTriggerObject(o.UpdatedAt, isavro, false, "long"),
		"hashcode":       toCodequalityTriggerObject(o.Hashcode, isavro, false, "string"),
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
func (o *CodequalityTrigger) Hash() string {
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

// GetCodequalityTriggerAvroSchemaSpec creates the avro schema specification for CodequalityTrigger
func GetCodequalityTriggerAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "agent",
		"name":      "CodequalityTrigger",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name": "customer_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "integration_id",
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
				"name": "updated_ts",
				"type": "long",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *CodequalityTrigger) GetEventAPIConfig() datamodel.EventAPIConfig {
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

// GetCodequalityTriggerAvroSchema creates the avro schema for CodequalityTrigger
func GetCodequalityTriggerAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetCodequalityTriggerAvroSchemaSpec())
}

// TransformCodequalityTriggerFunc is a function for transforming CodequalityTrigger during processing
type TransformCodequalityTriggerFunc func(input *CodequalityTrigger) (*CodequalityTrigger, error)

// NewCodequalityTriggerPipe creates a pipe for processing CodequalityTrigger items
func NewCodequalityTriggerPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformCodequalityTriggerFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewCodequalityTriggerInputStream(input, errors)
	var stream chan CodequalityTrigger
	if len(transforms) > 0 {
		stream = make(chan CodequalityTrigger, 1000)
	} else {
		stream = inch
	}
	outdone := NewCodequalityTriggerOutputStream(output, stream, errors)
	go func() {
		if len(transforms) > 0 {
			var stop bool
			for item := range inch {
				input := &item
				for _, transform := range transforms {
					out, err := transform(input)
					if err != nil {
						stop = true
						errors <- err
						break
					}
					if out == nil {
						input = nil
						break
					} else {
						input = out
					}
				}
				if stop {
					break
				}
				if input != nil {
					stream <- *input
				}
			}
			close(stream)
		}
		<-indone
		<-outdone
		done <- true
	}()
	return done
}

// NewCodequalityTriggerInputStreamDir creates a channel for reading CodequalityTrigger as JSON newlines from a directory of files
func NewCodequalityTriggerInputStreamDir(dir string, errors chan<- error, transforms ...TransformCodequalityTriggerFunc) (chan CodequalityTrigger, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/agent/codequality_trigger\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan CodequalityTrigger)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for codequality_trigger")
		ch := make(chan CodequalityTrigger)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewCodequalityTriggerInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan CodequalityTrigger)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewCodequalityTriggerInputStreamFile creates an channel for reading CodequalityTrigger as JSON newlines from filename
func NewCodequalityTriggerInputStreamFile(filename string, errors chan<- error, transforms ...TransformCodequalityTriggerFunc) (chan CodequalityTrigger, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan CodequalityTrigger)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	var f io.ReadCloser = of
	if filepath.Ext(filename) == ".gz" {
		gz, err := gzip.NewReader(f)
		if err != nil {
			of.Close()
			errors <- err
			ch := make(chan CodequalityTrigger)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewCodequalityTriggerInputStream(f, errors, transforms...)
}

// NewCodequalityTriggerInputStream creates an channel for reading CodequalityTrigger as JSON newlines from stream
func NewCodequalityTriggerInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformCodequalityTriggerFunc) (chan CodequalityTrigger, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan CodequalityTrigger, 1000)
	go func() {
		defer func() { stream.Close(); close(ch); done <- true }()
		r := bufio.NewReader(stream)
		for {
			buf, err := r.ReadBytes('\n')
			if err != nil {
				if err == io.EOF {
					return
				}
				errors <- err
				return
			}
			var item CodequalityTrigger
			if err := json.Unmarshal(buf, &item); err != nil {
				errors <- err
				return
			}
			in := &item
			var skip bool
			for _, transform := range transforms {
				in, err = transform(in)
				if err != nil {
					errors <- err
					return
				}
				if in == nil {
					skip = true
					break
				}
			}
			if !skip {
				ch <- *in
			}
		}
	}()
	return ch, done
}

// NewCodequalityTriggerOutputStreamDir will output json newlines from channel and save in dir
func NewCodequalityTriggerOutputStreamDir(dir string, ch chan CodequalityTrigger, errors chan<- error, transforms ...TransformCodequalityTriggerFunc) <-chan bool {
	fp := filepath.Join(dir, "/agent/codequality_trigger\\.json(\\.gz)?$")
	os.MkdirAll(filepath.Dir(fp), 0777)
	of, err := os.Create(fp)
	if err != nil {
		errors <- err
		done := make(chan bool, 1)
		done <- true
		return done
	}
	gz, err := gzip.NewWriterLevel(of, gzip.BestCompression)
	if err != nil {
		errors <- err
		done := make(chan bool, 1)
		done <- true
		return done
	}
	return NewCodequalityTriggerOutputStream(gz, ch, errors, transforms...)
}

// NewCodequalityTriggerOutputStream will output json newlines from channel to the stream
func NewCodequalityTriggerOutputStream(stream io.WriteCloser, ch chan CodequalityTrigger, errors chan<- error, transforms ...TransformCodequalityTriggerFunc) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() {
			if gz, ok := stream.(*gzip.Writer); ok {
				gz.Flush()
				gz.Close()
			}
			stream.Close()
			done <- true
		}()
		for item := range ch {
			in := &item
			var skip bool
			var err error
			for _, transform := range transforms {
				in, err = transform(in)
				if err != nil {
					errors <- err
					return
				}
				if in == nil {
					skip = true
					break
				}
			}
			if !skip {
				buf, err := json.Marshal(in)
				if err != nil {
					errors <- err
					return
				}
				stream.Write(buf)
				stream.Write([]byte{'\n'})
			}
		}
	}()
	return done
}

// CodequalityTriggerSendEvent is an event detail for sending data
type CodequalityTriggerSendEvent struct {
	CodequalityTrigger *CodequalityTrigger
	headers            map[string]string
	time               time.Time
	key                string
}

var _ datamodel.ModelSendEvent = (*CodequalityTriggerSendEvent)(nil)

// Key is the key to use for the message
func (e *CodequalityTriggerSendEvent) Key() string {
	if e.key == "" {
		return e.CodequalityTrigger.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *CodequalityTriggerSendEvent) Object() datamodel.Model {
	return e.CodequalityTrigger
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *CodequalityTriggerSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *CodequalityTriggerSendEvent) Timestamp() time.Time {
	return e.time
}

// CodequalityTriggerSendEventOpts is a function handler for setting opts
type CodequalityTriggerSendEventOpts func(o *CodequalityTriggerSendEvent)

// WithCodequalityTriggerSendEventKey sets the key value to a value different than the object ID
func WithCodequalityTriggerSendEventKey(key string) CodequalityTriggerSendEventOpts {
	return func(o *CodequalityTriggerSendEvent) {
		o.key = key
	}
}

// WithCodequalityTriggerSendEventTimestamp sets the timestamp value
func WithCodequalityTriggerSendEventTimestamp(tv time.Time) CodequalityTriggerSendEventOpts {
	return func(o *CodequalityTriggerSendEvent) {
		o.time = tv
	}
}

// WithCodequalityTriggerSendEventHeader sets the timestamp value
func WithCodequalityTriggerSendEventHeader(key, value string) CodequalityTriggerSendEventOpts {
	return func(o *CodequalityTriggerSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewCodequalityTriggerSendEvent returns a new CodequalityTriggerSendEvent instance
func NewCodequalityTriggerSendEvent(o *CodequalityTrigger, opts ...CodequalityTriggerSendEventOpts) *CodequalityTriggerSendEvent {
	res := &CodequalityTriggerSendEvent{
		CodequalityTrigger: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewCodequalityTriggerProducer will stream data from the channel
func NewCodequalityTriggerProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
	done := make(chan bool, 1)
	emptyTime := time.Unix(0, 0)
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
				if object, ok := item.Object().(*CodequalityTrigger); ok {
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
					msg := eventing.Message{
						Encoding:  eventing.AvroEncoding,
						Key:       item.Key(),
						Value:     binary,
						Codec:     codec,
						Headers:   headers,
						Timestamp: tv,
						Partition: -1, // select any partition based on partitioner strategy in kafka
						Topic:     object.GetTopicName().String(),
					}
					if err := producer.Send(ctx, msg); err != nil {
						errors <- fmt.Errorf("error sending %s. %v", object.String(), err)
					}
				} else {
					errors <- fmt.Errorf("invalid event received. expected an object of type agent.CodequalityTrigger but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewCodequalityTriggerConsumer will stream data from the topic into the provided channel
func NewCodequalityTriggerConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object CodequalityTrigger
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into agent.CodequalityTrigger: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into agent.CodequalityTrigger: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for agent.CodequalityTrigger")
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

			ch <- &CodequalityTriggerReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object CodequalityTrigger
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &CodequalityTriggerReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// CodequalityTriggerReceiveEvent is an event detail for receiving data
type CodequalityTriggerReceiveEvent struct {
	CodequalityTrigger *CodequalityTrigger
	message            eventing.Message
	eof                bool
}

var _ datamodel.ModelReceiveEvent = (*CodequalityTriggerReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *CodequalityTriggerReceiveEvent) Object() datamodel.Model {
	return e.CodequalityTrigger
}

// Message returns the underlying message data for the event
func (e *CodequalityTriggerReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *CodequalityTriggerReceiveEvent) EOF() bool {
	return e.eof
}

// CodequalityTriggerProducer implements the datamodel.ModelEventProducer
type CodequalityTriggerProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*CodequalityTriggerProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *CodequalityTriggerProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *CodequalityTriggerProducer) Close() error {
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
func (o *CodequalityTrigger) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *CodequalityTrigger) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &CodequalityTriggerProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewCodequalityTriggerProducer(newctx, producer, ch, errors, empty),
	}
}

// NewCodequalityTriggerProducerChannel returns a channel which can be used for producing Model events
func NewCodequalityTriggerProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewCodequalityTriggerProducerChannelSize(producer, 0, errors)
}

// NewCodequalityTriggerProducerChannelSize returns a channel which can be used for producing Model events
func NewCodequalityTriggerProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &CodequalityTriggerProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewCodequalityTriggerProducer(newctx, producer, ch, errors, empty),
	}
}

// CodequalityTriggerConsumer implements the datamodel.ModelEventConsumer
type CodequalityTriggerConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*CodequalityTriggerConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *CodequalityTriggerConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *CodequalityTriggerConsumer) Close() error {
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
func (o *CodequalityTrigger) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &CodequalityTriggerConsumer{
		ch:       ch,
		callback: NewCodequalityTriggerConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewCodequalityTriggerConsumerChannel returns a consumer channel which can be used to consume Model events
func NewCodequalityTriggerConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &CodequalityTriggerConsumer{
		ch:       ch,
		callback: NewCodequalityTriggerConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}