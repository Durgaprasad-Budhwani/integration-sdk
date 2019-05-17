// DO NOT EDIT -- generated code

// Package customer - customer specific data
package customer

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

	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	number "github.com/pinpt/go-common/number"
	"github.com/pinpt/integration-sdk/util"
)

// CostCenterDefaultTopic is the default topic name
const CostCenterDefaultTopic = "customer_CostCenter_topic"

// CostCenterDefaultStream is the default stream name
const CostCenterDefaultStream = "customer_CostCenter_stream"

// CostCenterDefaultTable is the default table name
const CostCenterDefaultTable = "customer_CostCenter"

// CostCenter blah blah blah
type CostCenter struct {
	// built in types

	ID         string `json:"cost_center_id" yaml:"cost_center_id"`
	RefID      string `json:"ref_id" yaml:"ref_id"`
	RefType    string `json:"ref_type" yaml:"ref_type"`
	CustomerID string `json:"customer_id" yaml:"customer_id"`
	Hashcode   string `json:"hashcode" yaml:"hashcode"`

	// custom types

	// Name blah blah
	Name string `json:"name" yaml:"name"`
	// Description blah blah
	Description string `json:"description" yaml:"description"`
	// Cost blah blah
	Cost float64 `json:"cost" yaml:"cost"`
}

func toCostCenterObject(o interface{}, isavro bool) interface{} {
	if o == nil {
		return nil
	}
	switch v := o.(type) {
	case nil:
		return nil
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		return v
	case *string, *int, *int8, *int16, *int32, *int64, *float32, *float64, *bool:
		return v
	case map[string]interface{}:
		return o
	case *map[string]interface{}:
		return v
	case *CostCenter:
		return v.ToMap()
	case CostCenter:
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
			arr = append(arr, toCostCenterObject(av, isavro))
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of CostCenter
func (o *CostCenter) String() string {
	return fmt.Sprintf("customer.CostCenter<%s>", o.ID)
}

func (o *CostCenter) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *CostCenter) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("CostCenter", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetRefID returns the RefID for the object
func (o *CostCenter) GetRefID() string {
	return o.RefID
}

// MarshalJSON returns the bytes for marshaling to json
func (o *CostCenter) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *CostCenter) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	// make sure that these have values if empty
	o.setDefaults()
	return nil
}

var cachedCodecCostCenter *goavro.Codec

// ToAvroBinary returns the data as Avro binary data
func (o *CostCenter) ToAvroBinary() ([]byte, *goavro.Codec, error) {
	if cachedCodecCostCenter == nil {
		c, err := CreateCostCenterAvroSchema()
		if err != nil {
			return nil, nil, err
		}
		cachedCodecCostCenter = c
	}
	kv := o.ToMap(true)
	jbuf, _ := json.Marshal(kv)
	native, _, err := cachedCodecCostCenter.NativeFromTextual(jbuf)
	if err != nil {
		return nil, nil, err
	}
	// Convert native Go form to binary Avro data
	buf, err := cachedCodecCostCenter.BinaryFromNative(nil, native)
	return buf, cachedCodecCostCenter, err
}

// Stringify returns the object in JSON format as a string
func (o *CostCenter) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two CostCenter objects are equal
func (o *CostCenter) IsEqual(other *CostCenter) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *CostCenter) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	return map[string]interface{}{
		"cost_center_id": o.GetID(),
		"ref_id":         o.GetRefID(),
		"ref_type":       o.RefType,
		"customer_id":    o.CustomerID,
		"hashcode":       o.Hash(),
		"name":           toCostCenterObject(o.Name, isavro),
		"description":    toCostCenterObject(o.Description, isavro),
		"cost":           toCostCenterObject(o.Cost, isavro),
	}
}

// FromMap attempts to load data into object from a map
func (o *CostCenter) FromMap(kv map[string]interface{}) {
	if val, ok := kv["cost_center_id"].(string); ok {
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
	if val, ok := kv["description"].(string); ok {
		o.Description = val
	} else {
		val := kv["description"]
		if val == nil {
			o.Description = ""
		} else {
			o.Description = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["cost"].(float64); ok {
		o.Cost = val
	} else {
		val := kv["cost"]
		if val == nil {
			o.Cost = number.ToFloat64Any(nil)
		} else {
			o.Cost = number.ToFloat64Any(val)
		}
	}
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *CostCenter) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.Name)
	args = append(args, o.Description)
	args = append(args, o.Cost)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// CreateCostCenterAvroSchemaSpec creates the avro schema specification for CostCenter
func CreateCostCenterAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "customer",
		"name":         "CostCenter",
		"connect.name": "customer.CostCenter",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "cost_center_id",
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
				"name": "description",
				"type": "string",
			},
			map[string]interface{}{
				"name": "cost",
				"type": "float",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// CreateCostCenterAvroSchema creates the avro schema for CostCenter
func CreateCostCenterAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(CreateCostCenterAvroSchemaSpec())
}

// TransformCostCenterFunc is a function for transforming CostCenter during processing
type TransformCostCenterFunc func(input *CostCenter) (*CostCenter, error)

// CreateCostCenterPipe creates a pipe for processing CostCenter items
func CreateCostCenterPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformCostCenterFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := CreateCostCenterInputStream(input, errors)
	var stream chan CostCenter
	if len(transforms) > 0 {
		stream = make(chan CostCenter, 1000)
	} else {
		stream = inch
	}
	outdone := CreateCostCenterOutputStream(output, stream, errors)
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

// CreateCostCenterInputStreamDir creates a channel for reading CostCenter as JSON newlines from a directory of files
func CreateCostCenterInputStreamDir(dir string, errors chan<- error, transforms ...TransformCostCenterFunc) (chan CostCenter, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/customer/cost_center\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan CostCenter)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for cost_center")
		ch := make(chan CostCenter)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return CreateCostCenterInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan CostCenter)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// CreateCostCenterInputStreamFile creates an channel for reading CostCenter as JSON newlines from filename
func CreateCostCenterInputStreamFile(filename string, errors chan<- error, transforms ...TransformCostCenterFunc) (chan CostCenter, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan CostCenter)
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
			ch := make(chan CostCenter)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return CreateCostCenterInputStream(f, errors, transforms...)
}

// CreateCostCenterInputStream creates an channel for reading CostCenter as JSON newlines from stream
func CreateCostCenterInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformCostCenterFunc) (chan CostCenter, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan CostCenter, 1000)
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
			var item CostCenter
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

// CreateCostCenterOutputStreamDir will output json newlines from channel and save in dir
func CreateCostCenterOutputStreamDir(dir string, ch chan CostCenter, errors chan<- error, transforms ...TransformCostCenterFunc) <-chan bool {
	fp := filepath.Join(dir, "/customer/cost_center\\.json(\\.gz)?$")
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
	return CreateCostCenterOutputStream(gz, ch, errors, transforms...)
}

// CreateCostCenterOutputStream will output json newlines from channel to the stream
func CreateCostCenterOutputStream(stream io.WriteCloser, ch chan CostCenter, errors chan<- error, transforms ...TransformCostCenterFunc) <-chan bool {
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

// CreateCostCenterProducer will stream data from the channel
func CreateCostCenterProducer(producer util.Producer, ch chan CostCenter, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		ctx := context.Background()
		for item := range ch {
			binary, codec, err := item.ToAvroBinary()
			if err != nil {
				errors <- fmt.Errorf("error encoding %s to avro binary data. %v", item.String(), err)
				return
			}
			if err := producer.Send(ctx, codec, []byte(item.ID), binary); err != nil {
				errors <- fmt.Errorf("error sending %s. %v", item.String(), err)
			}
		}
	}()
	return done
}

// CreateCostCenterConsumer will stream data from the default topic into the provided channel
func CreateCostCenterConsumer(factory util.ConsumerFactory, topic string, ch chan CostCenter, errors chan<- error) (<-chan bool, chan<- bool) {
	return CreateCostCenterConsumerForTopic(factory, CostCenterDefaultTopic, ch, errors)
}

// CreateCostCenterConsumerForTopic will stream data from the topic into the provided channel
func CreateCostCenterConsumerForTopic(factory util.ConsumerFactory, topic string, ch chan CostCenter, errors chan<- error) (<-chan bool, chan<- bool) {
	done := make(chan bool, 1)
	closed := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		callback := util.ConsumerCallback{
			OnDataReceived: func(key []byte, value []byte) error {
				var object CostCenter
				if err := json.Unmarshal(value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into CostCenter: %s", err)
				}
				ch <- object
				return nil
			},
			OnErrorReceived: func(err error) {
				errors <- err
			},
		}
		consumer, err := factory.CreateConsumer(topic, callback)
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
