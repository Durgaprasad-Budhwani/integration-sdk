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

	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	"github.com/pinpt/integration-sdk/util"
)

// ChangelogDefaultTopic is the default topic name
const ChangelogDefaultTopic = "work_Changelog_topic"

// ChangelogDefaultStream is the default stream name
const ChangelogDefaultStream = "work_Changelog_stream"

// ChangelogDefaultTable is the default table name
const ChangelogDefaultTable = "work_Changelog"

// Changelog change log
type Changelog struct {
	// built in types

	ID         string `json:"changelog_id" yaml:"changelog_id"`
	RefID      string `json:"ref_id" yaml:"ref_id"`
	RefType    string `json:"ref_type" yaml:"ref_type"`
	CustomerID string `json:"customer_id" yaml:"customer_id"`
	Hashcode   string `json:"hashcode" yaml:"hashcode"`

	// custom types

	// IssueID id of the issue
	IssueID string `json:"issue_id" yaml:"issue_id"`
	// CreatedAt the timestamp in UTC when this change was created
	CreatedAt int64 `json:"created_ts" yaml:"created_ts"`
	// Ordinal so we can order correctly in queries since dates could be equal
	Ordinal int64 `json:"ordinal" yaml:"ordinal"`
	// UserID id of the user of this change
	UserID string `json:"user_id" yaml:"user_id"`
	// Field name of the field that was changed
	Field string `json:"field" yaml:"field"`
	// FieldType type of the field that was changed
	FieldType string `json:"field_type" yaml:"field_type"`
	// From id of the change from
	From string `json:"from" yaml:"from"`
	// FromString name of the change from
	FromString string `json:"from_string" yaml:"from_string"`
	// To id of the change to
	To string `json:"to" yaml:"to"`
	// ToString name of the change to
	ToString string `json:"to_string" yaml:"to_string"`
}

func toChangelogObject(o interface{}, isavro bool) interface{} {
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
	case *Changelog:
		return v.ToMap()
	case Changelog:
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
			arr = append(arr, toChangelogObject(av, isavro))
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of Changelog
func (o *Changelog) String() string {
	return fmt.Sprintf("work.Changelog<%s>", o.ID)
}

func (o *Changelog) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Changelog) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Changelog", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetRefID returns the RefID for the object
func (o *Changelog) GetRefID() string {
	return o.RefID
}

// MarshalJSON returns the bytes for marshaling to json
func (o *Changelog) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Changelog) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	// make sure that these have values if empty
	o.setDefaults()
	return nil
}

var cachedCodecChangelog *goavro.Codec

// ToAvroBinary returns the data as Avro binary data
func (o *Changelog) ToAvroBinary() ([]byte, *goavro.Codec, error) {
	if cachedCodecChangelog == nil {
		c, err := CreateChangelogAvroSchema()
		if err != nil {
			return nil, nil, err
		}
		cachedCodecChangelog = c
	}
	kv := o.ToMap(true)
	jbuf, _ := json.Marshal(kv)
	native, _, err := cachedCodecChangelog.NativeFromTextual(jbuf)
	if err != nil {
		return nil, nil, err
	}
	// Convert native Go form to binary Avro data
	buf, err := cachedCodecChangelog.BinaryFromNative(nil, native)
	return buf, cachedCodecChangelog, err
}

// Stringify returns the object in JSON format as a string
func (o *Changelog) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Changelog objects are equal
func (o *Changelog) IsEqual(other *Changelog) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Changelog) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	return map[string]interface{}{
		"changelog_id": o.GetID(),
		"ref_id":       o.GetRefID(),
		"ref_type":     o.RefType,
		"customer_id":  o.CustomerID,
		"hashcode":     o.Hash(),
		"issue_id":     toChangelogObject(o.IssueID, isavro),
		"created_ts":   toChangelogObject(o.CreatedAt, isavro),
		"ordinal":      toChangelogObject(o.Ordinal, isavro),
		"user_id":      toChangelogObject(o.UserID, isavro),
		"field":        toChangelogObject(o.Field, isavro),
		"field_type":   toChangelogObject(o.FieldType, isavro),
		"from":         toChangelogObject(o.From, isavro),
		"from_string":  toChangelogObject(o.FromString, isavro),
		"to":           toChangelogObject(o.To, isavro),
		"to_string":    toChangelogObject(o.ToString, isavro),
	}
}

// FromMap attempts to load data into object from a map
func (o *Changelog) FromMap(kv map[string]interface{}) {
	if val, ok := kv["changelog_id"].(string); ok {
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
	if val, ok := kv["issue_id"].(string); ok {
		o.IssueID = val
	} else {
		val := kv["issue_id"]
		if val == nil {
			o.IssueID = ""
		} else {
			o.IssueID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["created_ts"].(int64); ok {
		o.CreatedAt = val
	} else {
		val := kv["created_ts"]
		if val == nil {
			o.CreatedAt = number.ToInt64Any(nil)
		} else {
			o.CreatedAt = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["ordinal"].(int64); ok {
		o.Ordinal = val
	} else {
		val := kv["ordinal"]
		if val == nil {
			o.Ordinal = number.ToInt64Any(nil)
		} else {
			o.Ordinal = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["user_id"].(string); ok {
		o.UserID = val
	} else {
		val := kv["user_id"]
		if val == nil {
			o.UserID = ""
		} else {
			o.UserID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["field"].(string); ok {
		o.Field = val
	} else {
		val := kv["field"]
		if val == nil {
			o.Field = ""
		} else {
			o.Field = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["field_type"].(string); ok {
		o.FieldType = val
	} else {
		val := kv["field_type"]
		if val == nil {
			o.FieldType = ""
		} else {
			o.FieldType = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["from"].(string); ok {
		o.From = val
	} else {
		val := kv["from"]
		if val == nil {
			o.From = ""
		} else {
			o.From = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["from_string"].(string); ok {
		o.FromString = val
	} else {
		val := kv["from_string"]
		if val == nil {
			o.FromString = ""
		} else {
			o.FromString = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["to"].(string); ok {
		o.To = val
	} else {
		val := kv["to"]
		if val == nil {
			o.To = ""
		} else {
			o.To = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["to_string"].(string); ok {
		o.ToString = val
	} else {
		val := kv["to_string"]
		if val == nil {
			o.ToString = ""
		} else {
			o.ToString = fmt.Sprintf("%v", val)
		}
	}
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *Changelog) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.IssueID)
	args = append(args, o.CreatedAt)
	args = append(args, o.Ordinal)
	args = append(args, o.UserID)
	args = append(args, o.Field)
	args = append(args, o.FieldType)
	args = append(args, o.From)
	args = append(args, o.FromString)
	args = append(args, o.To)
	args = append(args, o.ToString)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// CreateChangelogAvroSchemaSpec creates the avro schema specification for Changelog
func CreateChangelogAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "work",
		"name":         "Changelog",
		"connect.name": "work.Changelog",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "changelog_id",
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
				"name": "issue_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "created_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "ordinal",
				"type": "long",
			},
			map[string]interface{}{
				"name": "user_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "field",
				"type": "string",
			},
			map[string]interface{}{
				"name": "field_type",
				"type": "string",
			},
			map[string]interface{}{
				"name": "from",
				"type": "string",
			},
			map[string]interface{}{
				"name": "from_string",
				"type": "string",
			},
			map[string]interface{}{
				"name": "to",
				"type": "string",
			},
			map[string]interface{}{
				"name": "to_string",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// CreateChangelogAvroSchema creates the avro schema for Changelog
func CreateChangelogAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(CreateChangelogAvroSchemaSpec())
}

// TransformChangelogFunc is a function for transforming Changelog during processing
type TransformChangelogFunc func(input *Changelog) (*Changelog, error)

// CreateChangelogPipe creates a pipe for processing Changelog items
func CreateChangelogPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformChangelogFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := CreateChangelogInputStream(input, errors)
	var stream chan Changelog
	if len(transforms) > 0 {
		stream = make(chan Changelog, 1000)
	} else {
		stream = inch
	}
	outdone := CreateChangelogOutputStream(output, stream, errors)
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

// CreateChangelogInputStreamDir creates a channel for reading Changelog as JSON newlines from a directory of files
func CreateChangelogInputStreamDir(dir string, errors chan<- error, transforms ...TransformChangelogFunc) (chan Changelog, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/work/changelog\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Changelog)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for changelog")
		ch := make(chan Changelog)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return CreateChangelogInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Changelog)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// CreateChangelogInputStreamFile creates an channel for reading Changelog as JSON newlines from filename
func CreateChangelogInputStreamFile(filename string, errors chan<- error, transforms ...TransformChangelogFunc) (chan Changelog, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Changelog)
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
			ch := make(chan Changelog)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return CreateChangelogInputStream(f, errors, transforms...)
}

// CreateChangelogInputStream creates an channel for reading Changelog as JSON newlines from stream
func CreateChangelogInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformChangelogFunc) (chan Changelog, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Changelog, 1000)
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
			var item Changelog
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

// CreateChangelogOutputStreamDir will output json newlines from channel and save in dir
func CreateChangelogOutputStreamDir(dir string, ch chan Changelog, errors chan<- error, transforms ...TransformChangelogFunc) <-chan bool {
	fp := filepath.Join(dir, "/work/changelog\\.json(\\.gz)?$")
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
	return CreateChangelogOutputStream(gz, ch, errors, transforms...)
}

// CreateChangelogOutputStream will output json newlines from channel to the stream
func CreateChangelogOutputStream(stream io.WriteCloser, ch chan Changelog, errors chan<- error, transforms ...TransformChangelogFunc) <-chan bool {
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

// CreateChangelogProducer will stream data from the channel
func CreateChangelogProducer(producer util.Producer, ch chan Changelog, errors chan<- error) <-chan bool {
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

// CreateChangelogConsumer will stream data from the default topic into the provided channel
func CreateChangelogConsumer(factory util.ConsumerFactory, topic string, ch chan Changelog, errors chan<- error) (<-chan bool, chan<- bool) {
	return CreateChangelogConsumerForTopic(factory, ChangelogDefaultTopic, ch, errors)
}

// CreateChangelogConsumerForTopic will stream data from the topic into the provided channel
func CreateChangelogConsumerForTopic(factory util.ConsumerFactory, topic string, ch chan Changelog, errors chan<- error) (<-chan bool, chan<- bool) {
	done := make(chan bool, 1)
	closed := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		callback := util.ConsumerCallback{
			OnDataReceived: func(key []byte, value []byte) error {
				var object Changelog
				if err := json.Unmarshal(value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into Changelog: %s", err)
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
