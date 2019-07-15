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
)

const (
	// EnabledTopic is the default topic name
	EnabledTopic datamodel.TopicNameType = "agent_Enabled_topic"

	// EnabledStream is the default stream name
	EnabledStream datamodel.TopicNameType = "agent_Enabled_stream"

	// EnabledTable is the default table name
	EnabledTable datamodel.TopicNameType = "agent_Enabled"

	// EnabledModelName is the model name
	EnabledModelName datamodel.ModelNameType = "agent.Enabled"
)

const (
	// EnabledCustomerIDColumn is the customer_id column name
	EnabledCustomerIDColumn = "customer_id"
	// EnabledDateColumn is the date column name
	EnabledDateColumn = "date"
	// EnabledDateColumnEpochColumn is the epoch column property of the Date name
	EnabledDateColumnEpochColumn = "date->epoch"
	// EnabledDateColumnOffsetColumn is the offset column property of the Date name
	EnabledDateColumnOffsetColumn = "date->offset"
	// EnabledDateColumnRfc3339Column is the rfc3339 column property of the Date name
	EnabledDateColumnRfc3339Column = "date->rfc3339"
	// EnabledIDColumn is the id column name
	EnabledIDColumn = "id"
	// EnabledRefIDColumn is the ref_id column name
	EnabledRefIDColumn = "ref_id"
	// EnabledRefTypeColumn is the ref_type column name
	EnabledRefTypeColumn = "ref_type"
	// EnabledUUIDColumn is the uuid column name
	EnabledUUIDColumn = "uuid"
)

// EnabledDate represents the object structure for date
type EnabledDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func (o *EnabledDate) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": o.Epoch,
		// Offset the timezone offset from GMT
		"offset": o.Offset,
		// Rfc3339 the date in RFC3339 format
		"rfc3339": o.Rfc3339,
	}
}

// Enabled an agent event to indicate that it's enabled and ready for actions
type Enabled struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Date the date when the request was made
	Date EnabledDate `json:"date" bson:"date" yaml:"date" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Enabled)(nil)

func toEnabledObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toEnabledObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toEnabledObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toEnabledObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toEnabledObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toEnabledObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toEnabledObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toEnabledObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toEnabledObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toEnabledObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toEnabledObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toEnabledObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toEnabledObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case map[string]interface{}:
		return o
	case *map[string]interface{}:
		return v
	case map[string]string:
		return v
	case *map[string]string:
		return *v
	case *Enabled:
		return v.ToMap()
	case Enabled:
		return v.ToMap()
	case []string, []int64, []float64, []bool:
		return o
	case *[]string:
		return (*(o.(*[]string)))
	case *[]int64:
		return (*(o.(*[]int64)))
	case *[]float64:
		return (*(o.(*[]float64)))
	case *[]bool:
		return (*(o.(*[]bool)))
	case []interface{}:
		a := o.([]interface{})
		arr := make([]interface{}, 0)
		for _, av := range a {
			arr = append(arr, toEnabledObject(av, isavro, false, ""))
		}
		return arr

	case EnabledDate:
		vv := o.(EnabledDate)
		return vv.ToMap()
	case *EnabledDate:
		return (*o.(*EnabledDate)).ToMap()
	case []EnabledDate:
		arr := make([]interface{}, 0)
		for _, i := range o.([]EnabledDate) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]EnabledDate:
		arr := make([]interface{}, 0)
		vv := o.(*[]EnabledDate)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of Enabled
func (o *Enabled) String() string {
	return fmt.Sprintf("agent.Enabled<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Enabled) GetTopicName() datamodel.TopicNameType {
	return EnabledTopic
}

// GetModelName returns the name of the model
func (o *Enabled) GetModelName() datamodel.ModelNameType {
	return EnabledModelName
}

func (o *Enabled) setDefaults() {

	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Enabled) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Enabled", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Enabled) GetTopicKey() string {
	var i interface{} = o.UUID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Enabled) GetTimestamp() time.Time {
	var dt interface{} = o.Date
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
	case EnabledDate:
		return datetime.DateFromEpoch(v.Epoch)
	}
	panic("not sure how to handle the date time format for Enabled")
}

// GetRefID returns the RefID for the object
func (o *Enabled) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Enabled) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Enabled) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Enabled) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Enabled) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = EnabledModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Enabled) GetTopicConfig() *datamodel.ModelTopicConfig {
	retention, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}

	ttl, err := time.ParseDuration("0s")
	if err != nil {
		ttl = 0
	}
	return &datamodel.ModelTopicConfig{
		Key:               "uuid",
		Timestamp:         "date",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetStateKey returns a key for use in state store
func (o *Enabled) GetStateKey() string {
	key := "uuid"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *Enabled) GetCustomerID() string {
	return o.CustomerID
}

// Clone returns an exact copy of Enabled
func (o *Enabled) Clone() datamodel.Model {
	c := new(Enabled)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Enabled) Anon() datamodel.Model {
	c := new(Enabled)
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
func (o *Enabled) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Enabled) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

var cachedCodecEnabled *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *Enabled) GetAvroCodec() *goavro.Codec {
	if cachedCodecEnabled == nil {
		c, err := GetEnabledAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecEnabled = c
	}
	return cachedCodecEnabled
}

// ToAvroBinary returns the data as Avro binary data
func (o *Enabled) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *Enabled) FromAvroBinary(value []byte) error {
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
func (o *Enabled) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Enabled objects are equal
func (o *Enabled) IsEqual(other *Enabled) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Enabled) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	o.setDefaults()
	return map[string]interface{}{
		"customer_id": toEnabledObject(o.CustomerID, isavro, false, "string"),
		"date":        toEnabledObject(o.Date, isavro, false, "date"),
		"id":          toEnabledObject(o.ID, isavro, false, "string"),
		"ref_id":      toEnabledObject(o.RefID, isavro, false, "string"),
		"ref_type":    toEnabledObject(o.RefType, isavro, false, "string"),
		"uuid":        toEnabledObject(o.UUID, isavro, false, "string"),
		"hashcode":    toEnabledObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Enabled) FromMap(kv map[string]interface{}) {
	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["customer_id"].(string); ok {
		o.CustomerID = val
	} else {
		val := kv["customer_id"]
		if val == nil {
			o.CustomerID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.CustomerID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["date"].(EnabledDate); ok {
		o.Date = val
	} else {
		val := kv["date"]
		if val == nil {
			o.Date = EnabledDate{}
		} else {
			o.Date = EnabledDate{}
			b, _ := json.Marshal(val)
			json.Unmarshal(b, &o.Date)

		}
	}
	if val, ok := kv["id"].(string); ok {
		o.ID = val
	} else {
		val := kv["id"]
		if val == nil {
			o.ID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.ID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["ref_id"].(string); ok {
		o.RefID = val
	} else {
		val := kv["ref_id"]
		if val == nil {
			o.RefID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.RefID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["ref_type"].(string); ok {
		o.RefType = val
	} else {
		val := kv["ref_type"]
		if val == nil {
			o.RefType = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.RefType = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["uuid"].(string); ok {
		o.UUID = val
	} else {
		val := kv["uuid"]
		if val == nil {
			o.UUID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.UUID = fmt.Sprintf("%v", val)
		}
	}
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *Enabled) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.Date)
	args = append(args, o.ID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.UUID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEnabledAvroSchemaSpec creates the avro schema specification for Enabled
func GetEnabledAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "agent",
		"name":      "Enabled",
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
				"name": "date",
				"type": map[string]interface{}{"name": "date", "fields": []interface{}{map[string]interface{}{"type": "long", "name": "epoch", "doc": "the date in epoch format"}, map[string]interface{}{"type": "long", "name": "offset", "doc": "the timezone offset from GMT"}, map[string]interface{}{"type": "string", "name": "rfc3339", "doc": "the date in RFC3339 format"}}, "doc": "the date when the request was made", "type": "record"},
			},
			map[string]interface{}{
				"name": "id",
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
				"name": "uuid",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *Enabled) GetEventAPIConfig() datamodel.EventAPIConfig {
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

// GetEnabledAvroSchema creates the avro schema for Enabled
func GetEnabledAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetEnabledAvroSchemaSpec())
}

// TransformEnabledFunc is a function for transforming Enabled during processing
type TransformEnabledFunc func(input *Enabled) (*Enabled, error)

// NewEnabledPipe creates a pipe for processing Enabled items
func NewEnabledPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformEnabledFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewEnabledInputStream(input, errors)
	var stream chan Enabled
	if len(transforms) > 0 {
		stream = make(chan Enabled, 1000)
	} else {
		stream = inch
	}
	outdone := NewEnabledOutputStream(output, stream, errors)
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

// NewEnabledInputStreamDir creates a channel for reading Enabled as JSON newlines from a directory of files
func NewEnabledInputStreamDir(dir string, errors chan<- error, transforms ...TransformEnabledFunc) (chan Enabled, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/agent/enabled\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Enabled)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for enabled")
		ch := make(chan Enabled)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewEnabledInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Enabled)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewEnabledInputStreamFile creates an channel for reading Enabled as JSON newlines from filename
func NewEnabledInputStreamFile(filename string, errors chan<- error, transforms ...TransformEnabledFunc) (chan Enabled, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Enabled)
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
			ch := make(chan Enabled)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewEnabledInputStream(f, errors, transforms...)
}

// NewEnabledInputStream creates an channel for reading Enabled as JSON newlines from stream
func NewEnabledInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformEnabledFunc) (chan Enabled, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Enabled, 1000)
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
			var item Enabled
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

// NewEnabledOutputStreamDir will output json newlines from channel and save in dir
func NewEnabledOutputStreamDir(dir string, ch chan Enabled, errors chan<- error, transforms ...TransformEnabledFunc) <-chan bool {
	fp := filepath.Join(dir, "/agent/enabled\\.json(\\.gz)?$")
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
	return NewEnabledOutputStream(gz, ch, errors, transforms...)
}

// NewEnabledOutputStream will output json newlines from channel to the stream
func NewEnabledOutputStream(stream io.WriteCloser, ch chan Enabled, errors chan<- error, transforms ...TransformEnabledFunc) <-chan bool {
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

// EnabledSendEvent is an event detail for sending data
type EnabledSendEvent struct {
	Enabled *Enabled
	headers map[string]string
	time    time.Time
	key     string
}

var _ datamodel.ModelSendEvent = (*EnabledSendEvent)(nil)

// Key is the key to use for the message
func (e *EnabledSendEvent) Key() string {
	if e.key == "" {
		return e.Enabled.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *EnabledSendEvent) Object() datamodel.Model {
	return e.Enabled
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *EnabledSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *EnabledSendEvent) Timestamp() time.Time {
	return e.time
}

// EnabledSendEventOpts is a function handler for setting opts
type EnabledSendEventOpts func(o *EnabledSendEvent)

// WithEnabledSendEventKey sets the key value to a value different than the object ID
func WithEnabledSendEventKey(key string) EnabledSendEventOpts {
	return func(o *EnabledSendEvent) {
		o.key = key
	}
}

// WithEnabledSendEventTimestamp sets the timestamp value
func WithEnabledSendEventTimestamp(tv time.Time) EnabledSendEventOpts {
	return func(o *EnabledSendEvent) {
		o.time = tv
	}
}

// WithEnabledSendEventHeader sets the timestamp value
func WithEnabledSendEventHeader(key, value string) EnabledSendEventOpts {
	return func(o *EnabledSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewEnabledSendEvent returns a new EnabledSendEvent instance
func NewEnabledSendEvent(o *Enabled, opts ...EnabledSendEventOpts) *EnabledSendEvent {
	res := &EnabledSendEvent{
		Enabled: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewEnabledProducer will stream data from the channel
func NewEnabledProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
	done := make(chan bool, 1)
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
				if object, ok := item.Object().(*Enabled); ok {
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
					if tv.IsZero() {
						tv = time.Now() // if its still zero, use the ingest time
					}
					msg := eventing.Message{
						Encoding:  eventing.AvroEncoding,
						Key:       item.Key(),
						Value:     binary,
						Codec:     codec,
						Headers:   headers,
						Timestamp: tv,
						Topic:     object.GetTopicName().String(),
					}
					if err := producer.Send(ctx, msg); err != nil {
						errors <- fmt.Errorf("error sending %s. %v", object.String(), err)
					}
				} else {
					errors <- fmt.Errorf("invalid event received. expected an object of type agent.Enabled but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewEnabledConsumer will stream data from the topic into the provided channel
func NewEnabledConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object Enabled
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into agent.Enabled: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into agent.Enabled: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for agent.Enabled")
			}

			// ignore messages that have exceeded the TTL
			cfg := object.GetTopicConfig()
			if cfg != nil && cfg.TTL != 0 && msg.Timestamp.UTC().Add(cfg.TTL).Sub(time.Now().UTC()) < 0 {
				return nil
			}
			msg.Codec = object.GetAvroCodec() // match the codec

			ch <- &EnabledReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object Enabled
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &EnabledReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// EnabledReceiveEvent is an event detail for receiving data
type EnabledReceiveEvent struct {
	Enabled *Enabled
	message eventing.Message
	eof     bool
}

var _ datamodel.ModelReceiveEvent = (*EnabledReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *EnabledReceiveEvent) Object() datamodel.Model {
	return e.Enabled
}

// Message returns the underlying message data for the event
func (e *EnabledReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *EnabledReceiveEvent) EOF() bool {
	return e.eof
}

// EnabledProducer implements the datamodel.ModelEventProducer
type EnabledProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*EnabledProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *EnabledProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *EnabledProducer) Close() error {
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
func (o *Enabled) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *Enabled) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &EnabledProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewEnabledProducer(newctx, producer, ch, errors, empty),
	}
}

// NewEnabledProducerChannel returns a channel which can be used for producing Model events
func NewEnabledProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewEnabledProducerChannelSize(producer, 0, errors)
}

// NewEnabledProducerChannelSize returns a channel which can be used for producing Model events
func NewEnabledProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &EnabledProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewEnabledProducer(newctx, producer, ch, errors, empty),
	}
}

// EnabledConsumer implements the datamodel.ModelEventConsumer
type EnabledConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*EnabledConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *EnabledConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *EnabledConsumer) Close() error {
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
func (o *Enabled) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &EnabledConsumer{
		ch:       ch,
		callback: NewEnabledConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewEnabledConsumerChannel returns a consumer channel which can be used to consume Model events
func NewEnabledConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &EnabledConsumer{
		ch:       ch,
		callback: NewEnabledConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
