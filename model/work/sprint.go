// DO NOT EDIT -- generated code

// Package work - the system which contains project work
package work

import (
	"bufio"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"regexp"

	"github.com/bxcodec/faker"
	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/event"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
)

// SprintTopic is the default topic name
const SprintTopic datamodel.TopicNameType = "work_Sprint_topic"

// SprintStream is the default stream name
const SprintStream datamodel.TopicNameType = "work_Sprint_stream"

// SprintTable is the default table name
const SprintTable datamodel.TopicNameType = "work_Sprint"

// SprintModelName is the model name
const SprintModelName datamodel.ModelNameType = "work.Sprint"

// Sprint sprint details
type Sprint struct {
	// built in types

	ID         string `json:"sprint_id" yaml:"sprint_id" faker:"-"`
	RefID      string `json:"ref_id" yaml:"ref_id" faker:"-"`
	RefType    string `json:"ref_type" yaml:"ref_type" faker:"-"`
	CustomerID string `json:"customer_id" yaml:"customer_id" faker:"-"`
	Hashcode   string `json:"hashcode" yaml:"hashcode" faker:"-"`
	// custom types

	// Name the name of the field
	Name string `json:"name" yaml:"name" faker:"-"`
	// Identifier the common identifier for the sprint
	Identifier string `json:"identifier" yaml:"identifier" faker:"-"`
	// Status status of the sprint
	Status string `json:"status" yaml:"status" faker:"-"`
	// StartedAt the timestamp in UTC that the sprint was started
	StartedAt int64 `json:"started_ts" yaml:"started_ts" faker:"-"`
	// EndedAt the timestamp in UTC that the sprint was ended
	EndedAt *int64 `json:"ended_ts" yaml:"ended_ts" faker:"-"`
	// CompletedAt the timestamp in UTC that the sprint was completed
	CompletedAt *int64 `json:"completed_ts" yaml:"completed_ts" faker:"-"`
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
	if o == nil {
		return toSprintObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toSprintObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toSprintObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toSprintObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toSprintObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toSprintObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toSprintObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toSprintObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toSprintObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toSprintObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toSprintObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case map[string]interface{}:
		return o
	case *map[string]interface{}:
		return v
	case *Sprint:
		return v.ToMap()
	case Sprint:
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
			arr = append(arr, toSprintObject(av, isavro, false, ""))
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of Sprint
func (o *Sprint) String() string {
	return fmt.Sprintf("work.Sprint<%s>", o.ID)
}

func (o *Sprint) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Sprint) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Sprint", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetRefID returns the RefID for the object
func (o *Sprint) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Sprint) IsMaterialized() bool {
	return false
}

// MaterializedName returns the name of the materialized table
func (o *Sprint) MaterializedName() string {
	panic("work.Sprint is not a materialized table")
}

// Clone returns an exact copy of Sprint
func (o *Sprint) Clone() *Sprint {
	c := new(Sprint)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Sprint) Anon() *Sprint {
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
	// make sure that these have values if empty
	o.setDefaults()
	return nil
}

var cachedCodecSprint *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *Sprint) GetAvroCodec() *goavro.Codec {
	if cachedCodecSprint == nil {
		c, err := CreateSprintAvroSchema()
		if err != nil {
			return nil, nil, err
		}
		cachedCodecSprint = c
	}
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

// Stringify returns the object in JSON format as a string
func (o *Sprint) Stringify() string {
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
	return map[string]interface{}{
		"sprint_id":    o.GetID(),
		"ref_id":       o.GetRefID(),
		"ref_type":     o.RefType,
		"customer_id":  o.CustomerID,
		"hashcode":     o.Hash(),
		"name":         toSprintObject(o.Name, isavro, false, "string"),
		"identifier":   toSprintObject(o.Identifier, isavro, false, "string"),
		"status":       toSprintObject(o.Status, isavro, false, "string"),
		"started_ts":   toSprintObject(o.StartedAt, isavro, false, "long"),
		"ended_ts":     toSprintObject(o.EndedAt, isavro, true, "long"),
		"completed_ts": toSprintObject(o.CompletedAt, isavro, true, "long"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Sprint) FromMap(kv map[string]interface{}) {
	if val, ok := kv["sprint_id"].(string); ok {
		o.ID = val
	}
	if val, ok := kv["ref_id"].(string); ok {
		o.RefID = val
	}
	if val, ok := kv["ref_type"].(string); ok {
		o.RefType = val
	}
	if val, ok := kv["customer_id"].(string); ok {
		o.CustomerID = val
	}
	if val, ok := kv["name"].(string); ok {
		o.Name = val
	} else {
		val := kv["name"]
		if val == nil {
			o.Name = ""
		} else {
			o.Name = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["identifier"].(string); ok {
		o.Identifier = val
	} else {
		val := kv["identifier"]
		if val == nil {
			o.Identifier = ""
		} else {
			o.Identifier = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["status"].(string); ok {
		o.Status = val
	} else {
		val := kv["status"]
		if val == nil {
			o.Status = ""
		} else {
			o.Status = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["started_ts"].(int64); ok {
		o.StartedAt = val
	} else {
		val := kv["started_ts"]
		if val == nil {
			o.StartedAt = number.ToInt64Any(nil)
		} else {
			o.StartedAt = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["ended_ts"].(*int64); ok {
		o.EndedAt = val
	} else if val, ok := kv["ended_ts"].(int64); ok {
		o.EndedAt = &val
	} else {
		val := kv["ended_ts"]
		if val == nil {
			o.EndedAt = number.Int64Pointer(number.ToInt64Any(nil))
		} else {
			o.EndedAt = number.Int64Pointer(number.ToInt64Any(val))
		}
	}
	if val, ok := kv["completed_ts"].(*int64); ok {
		o.CompletedAt = val
	} else if val, ok := kv["completed_ts"].(int64); ok {
		o.CompletedAt = &val
	} else {
		val := kv["completed_ts"]
		if val == nil {
			o.CompletedAt = number.Int64Pointer(number.ToInt64Any(nil))
		} else {
			o.CompletedAt = number.Int64Pointer(number.ToInt64Any(val))
		}
	}
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *Sprint) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.CustomerID)
	args = append(args, o.Name)
	args = append(args, o.Identifier)
	args = append(args, o.Status)
	args = append(args, o.StartedAt)
	args = append(args, o.EndedAt)
	args = append(args, o.CompletedAt)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// CreateSprintAvroSchemaSpec creates the avro schema specification for Sprint
func CreateSprintAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "work",
		"name":         "Sprint",
		"connect.name": "work.Sprint",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "sprint_id",
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
				"name": "customer_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name": "name",
				"type": "string",
			},
			map[string]interface{}{
				"name": "identifier",
				"type": "string",
			},
			map[string]interface{}{
				"name": "status",
				"type": "string",
			},
			map[string]interface{}{
				"name": "started_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name":    "ended_ts",
				"type":    []interface{}{"null", "long"},
				"default": nil,
			},
			map[string]interface{}{
				"name":    "completed_ts",
				"type":    []interface{}{"null", "long"},
				"default": nil,
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// CreateSprintAvroSchema creates the avro schema for Sprint
func CreateSprintAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(CreateSprintAvroSchemaSpec())
}

// TransformSprintFunc is a function for transforming Sprint during processing
type TransformSprintFunc func(input *Sprint) (*Sprint, error)

// CreateSprintPipe creates a pipe for processing Sprint items
func CreateSprintPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformSprintFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := CreateSprintInputStream(input, errors)
	var stream chan Sprint
	if len(transforms) > 0 {
		stream = make(chan Sprint, 1000)
	} else {
		stream = inch
	}
	outdone := CreateSprintOutputStream(output, stream, errors)
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

// CreateSprintInputStreamDir creates a channel for reading Sprint as JSON newlines from a directory of files
func CreateSprintInputStreamDir(dir string, errors chan<- error, transforms ...TransformSprintFunc) (chan Sprint, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/work/sprint\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Sprint)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for sprint")
		ch := make(chan Sprint)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return CreateSprintInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Sprint)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// CreateSprintInputStreamFile creates an channel for reading Sprint as JSON newlines from filename
func CreateSprintInputStreamFile(filename string, errors chan<- error, transforms ...TransformSprintFunc) (chan Sprint, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Sprint)
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
			ch := make(chan Sprint)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return CreateSprintInputStream(f, errors, transforms...)
}

// CreateSprintInputStream creates an channel for reading Sprint as JSON newlines from stream
func CreateSprintInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformSprintFunc) (chan Sprint, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Sprint, 1000)
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
			var item Sprint
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

// CreateSprintOutputStreamDir will output json newlines from channel and save in dir
func CreateSprintOutputStreamDir(dir string, ch chan Sprint, errors chan<- error, transforms ...TransformSprintFunc) <-chan bool {
	fp := filepath.Join(dir, "/work/sprint\\.json(\\.gz)?$")
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
	return CreateSprintOutputStream(gz, ch, errors, transforms...)
}

// CreateSprintOutputStream will output json newlines from channel to the stream
func CreateSprintOutputStream(stream io.WriteCloser, ch chan Sprint, errors chan<- error, transforms ...TransformSprintFunc) <-chan bool {
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

// SprintSendEvent is an event detail for sending data
type SprintSendEvent struct {
	Sprint  Sprint
	Headers map[string]string
}

// CreateSprintProducer will stream data from the channel
func CreateSprintProducer(producer event.Producer, ch chan SprintSendEvent, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		ctx := context.Background()
		for item := range ch {
			binary, codec, err := item.Sprint.ToAvroBinary()
			if err != nil {
				errors <- fmt.Errorf("error encoding %s to avro binary data. %v", item.String(), err)
				return
			}
			headers := map[string]string{
				"customer_id": item.Sprint.CustomerID,
			}
			if item.Headers != nil {
				for k, v := range item.Headers {
					headers[k] = v
				}
			}
			msg := event.Message{
				Key:     item.Sprint.ID,
				Value:   binary,
				Codec:   codec,
				Headers: headers,
			}
			if err := producer.Send(ctx, msg); err != nil {
				errors <- fmt.Errorf("error sending %s. %v", item.Sprint.String(), err)
			}
		}
	}()
	return done
}

// SprintReceiveEvent is an event detail for receiving data
type SprintReceiveEvent struct {
	Sprint  Sprint
	Message event.Message
}

// CreateSprintConsumer will stream data from the topic into the provided channel
func CreateSprintConsumer(factory event.ConsumerFactory, topic datamodel.TopicNameType, ch chan SprintReceiveEvent, errors chan<- error) (<-chan bool, chan<- bool) {
	done := make(chan bool, 1)
	closed := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		callback := event.ConsumerCallback{
			OnDataReceived: func(msg event.Message) error {
				var object Sprint
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into work.Sprint: %s", err)
				}
				msg.Codec = object.GetAvroCodec() // match the codec
				ch <- SprintReceiveEvent{object, msg}
				return nil
			},
			OnErrorReceived: func(err error) {
				errors <- err
			},
		}
		consumer, err := factory.CreateConsumer(string(topic), callback)
		if err != nil {
			errors <- err
			return
		}
		select {
		case <-closed:
			consumer.Close()
		}
	}()
	return done, closed
}
