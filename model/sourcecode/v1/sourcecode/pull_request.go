// DO NOT EDIT -- generated code

// Package sourcecode - the system which contains source code
package sourcecode

import (
	"bufio"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	cluster "github.com/bsm/sarama-cluster"
	kafka "github.com/dangkaka/go-kafka-avro"
	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	"github.com/pinpt/integration-sdk/util"
)

// PullRequestDefaultTopic is the default topic name
const PullRequestDefaultTopic = "sourcecode.PullRequest.topic"

// PullRequestDefaultStream is the default stream name
const PullRequestDefaultStream = "sourcecode.PullRequest.topicstream"

// PullRequestDefaultTable is the default table name
const PullRequestDefaultTable = "sourcecode.PullRequest"

// PullRequest the pull request for a given repo
type PullRequest struct {
	// built in types

	ID         string `json:"id" yaml:"id"`
	RefID      string `json:"ref_id" yaml:"ref_id"`
	RefType    string `json:"ref_type" yaml:"ref_type"`
	CustomerID string `json:"customer_id" yaml:"customer_id"`
	Hashcode   string `json:"hashcode" yaml:"hashcode"`

	// custom types

	RepoID      string `json:"repo_id" yaml:"repo_id"`
	Title       string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`
	URL         string `json:"url" yaml:"url"`
	CreatedAt   int64  `json:"created_ts" yaml:"created_ts"`
	MergedAt    *int64 `json:"merged_ts" yaml:"merged_ts"`
	ClosedAt    *int64 `json:"closed_ts" yaml:"closed_ts"`
	Status      string `json:"status" yaml:"status"`
	UserRefID   string `json:"user_ref_id" yaml:"user_ref_id"`
}

// String returns a string representation of PullRequest
func (o *PullRequest) String() string {
	return fmt.Sprintf("sourcecode.v1.PullRequest<%s>", o.ID)
}

func (o *PullRequest) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *PullRequest) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("PullRequest", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetRefID returns the RefID for the object
func (o *PullRequest) GetRefID() string {
	return o.RefID
}

// MarshalJSON returns the bytes for marshaling to json
func (o *PullRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *PullRequest) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	// make sure that these have values if empty
	o.setDefaults()
	return nil
}

// Stringify returns the object in JSON format as a string
func (o *PullRequest) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two PullRequest objects are equal
func (o *PullRequest) IsEqual(other *PullRequest) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *PullRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":          o.GetID(),
		"ref_id":      o.GetRefID(),
		"ref_type":    o.RefType,
		"customer_id": o.CustomerID,
		"hashcode":    o.Hash(),
		"repo_id":     o.RepoID,
		"title":       o.Title,
		"description": o.Description,
		"url":         o.URL,
		"created_ts":  o.CreatedAt,
		"merged_ts":   o.MergedAt,
		"closed_ts":   o.ClosedAt,
		"status":      o.Status,
		"user_ref_id": o.UserRefID,
	}
}

// FromMap attempts to load data into object from a map
func (o *PullRequest) FromMap(kv map[string]interface{}) {
	if val, ok := kv["id"].(string); ok {
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
	if val, ok := kv["repo_id"].(string); ok {
		o.RepoID = val
	} else {
		val := kv["repo_id"]
		o.RepoID = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["title"].(string); ok {
		o.Title = val
	} else {
		val := kv["title"]
		o.Title = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["description"].(string); ok {
		o.Description = val
	} else {
		val := kv["description"]
		o.Description = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["url"].(string); ok {
		o.URL = val
	} else {
		val := kv["url"]
		o.URL = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["created_ts"].(int64); ok {
		o.CreatedAt = val
	} else {
		val := kv["created_ts"]
		o.CreatedAt = number.ToInt64Any(val)
	}
	if val, ok := kv["merged_ts"].(*int64); ok {
		o.MergedAt = val
	} else if val, ok := kv["merged_ts"].(int64); ok {
		o.MergedAt = &val
	} else {
		val := kv["merged_ts"]
		o.MergedAt = number.Int64Pointer(number.ToInt64Any(val))
	}
	if val, ok := kv["closed_ts"].(*int64); ok {
		o.ClosedAt = val
	} else if val, ok := kv["closed_ts"].(int64); ok {
		o.ClosedAt = &val
	} else {
		val := kv["closed_ts"]
		o.ClosedAt = number.Int64Pointer(number.ToInt64Any(val))
	}
	if val, ok := kv["status"].(string); ok {
		o.Status = val
	} else {
		val := kv["status"]
		o.Status = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["user_ref_id"].(string); ok {
		o.UserRefID = val
	} else {
		val := kv["user_ref_id"]
		o.UserRefID = fmt.Sprintf("%v", val)
	}
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *PullRequest) Hash() string {
	if o.Hashcode == "" {
		args := make([]interface{}, 0)
		args = append(args, o.GetID())
		args = append(args, o.GetRefID())
		args = append(args, o.RefType)
		args = append(args, o.RepoID)
		args = append(args, o.Title)
		args = append(args, o.Description)
		args = append(args, o.URL)
		args = append(args, o.CreatedAt)
		args = append(args, o.MergedAt)
		args = append(args, o.ClosedAt)
		args = append(args, o.Status)
		args = append(args, o.UserRefID)
		o.Hashcode = hash.Values(args...)
	}
	return o.Hashcode
}

// CreatePullRequestAvroSchemaSpec creates the avro schema specification for PullRequest
func CreatePullRequestAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "sourcecode.v1",
		"name":         "PullRequest",
		"connect.name": "sourcecode.v1.PullRequest",
		"fields": []map[string]interface{}{
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
				"name": "customer_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name": "repo_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "title",
				"type": "string",
			},
			map[string]interface{}{
				"name": "description",
				"type": "string",
			},
			map[string]interface{}{
				"name": "url",
				"type": "string",
			},
			map[string]interface{}{
				"name": "created_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "merged_ts",
				"type": []string{"null", "long"},
			},
			map[string]interface{}{
				"name": "closed_ts",
				"type": []string{"null", "long"},
			},
			map[string]interface{}{
				"name": "status",
				"type": "string",
			},
			map[string]interface{}{
				"name": "user_ref_id",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// CreatePullRequestAvroSchema creates the avro schema for PullRequest
func CreatePullRequestAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(CreatePullRequestAvroSchemaSpec())
}

// CreatePullRequestKQLStreamSQL creates KQL Stream SQL for PullRequest
func CreatePullRequestKQLStreamSQL() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("CREATE STREAM %s ", PullRequestDefaultStream))
	builder.WriteString(fmt.Sprintf("WITH (KAFKA_TOPIC='%s', VALUE_FORMAT='AVRO', KEY='id'", PullRequestDefaultTopic))
	builder.WriteString(");")
	return builder.String()
}

// CreatePullRequestKQLTableSQL creates KQL Table SQL for PullRequest
func CreatePullRequestKQLTableSQL() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("CREATE TABLE %s ", PullRequestDefaultTable))
	builder.WriteString(fmt.Sprintf("WITH (KAFKA_TOPIC='%s', VALUE_FORMAT='AVRO', KEY='id'", PullRequestDefaultTopic))
	builder.WriteString(");")
	return builder.String()
}

// TransformPullRequestFunc is a function for transforming PullRequest during processing
type TransformPullRequestFunc func(input *PullRequest) (*PullRequest, error)

// CreatePullRequestPipe creates a pipe for processing PullRequest items
func CreatePullRequestPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformPullRequestFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := CreatePullRequestInputStream(input, errors)
	var stream chan PullRequest
	if len(transforms) > 0 {
		stream = make(chan PullRequest, 1000)
	} else {
		stream = inch
	}
	outdone := CreatePullRequestOutputStream(output, stream, errors)
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

// CreatePullRequestInputStreamDir creates a channel for reading PullRequest as JSON newlines from a directory of files
func CreatePullRequestInputStreamDir(dir string, errors chan<- error, transforms ...TransformPullRequestFunc) (chan PullRequest, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/sourcecode/v1/pull_request\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan PullRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for pull_request")
		ch := make(chan PullRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return CreatePullRequestInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan PullRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// CreatePullRequestInputStreamFile creates an channel for reading PullRequest as JSON newlines from filename
func CreatePullRequestInputStreamFile(filename string, errors chan<- error, transforms ...TransformPullRequestFunc) (chan PullRequest, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan PullRequest)
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
			ch := make(chan PullRequest)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return CreatePullRequestInputStream(f, errors, transforms...)
}

// CreatePullRequestInputStream creates an channel for reading PullRequest as JSON newlines from stream
func CreatePullRequestInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformPullRequestFunc) (chan PullRequest, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan PullRequest, 1000)
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
			var item PullRequest
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

// CreatePullRequestOutputStreamDir will output json newlines from channel and save in dir
func CreatePullRequestOutputStreamDir(dir string, ch chan PullRequest, errors chan<- error, transforms ...TransformPullRequestFunc) <-chan bool {
	fp := filepath.Join(dir, "/sourcecode/v1/pull_request\\.json(\\.gz)?$")
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
	return CreatePullRequestOutputStream(gz, ch, errors, transforms...)
}

// CreatePullRequestOutputStream will output json newlines from channel to the stream
func CreatePullRequestOutputStream(stream io.WriteCloser, ch chan PullRequest, errors chan<- error, transforms ...TransformPullRequestFunc) <-chan bool {
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

// CreatePullRequestProducer will stream data from the channel
func CreatePullRequestProducer(producer *util.KafkaProducer, ch chan PullRequest, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		schemaspec := CreatePullRequestAvroSchemaSpec()
		ctx := context.Background()
		for item := range ch {
			if err := producer.Send(ctx, schemaspec, []byte(item.ID), []byte(item.Stringify())); err != nil {
				errors <- fmt.Errorf("error sending %s. %v", item.String(), err)
			}
		}
	}()
	return done
}

// CreatePullRequestConsumer will stream data from the default topic into the provided channel
func CreatePullRequestConsumer(kafkaServers []string, schemaRegistryServers []string, topic string, consumerGroupID string, ch chan PullRequest, errors chan<- error) (<-chan bool, chan<- bool) {
	return CreatePullRequestConsumerForTopic(kafkaServers, schemaRegistryServers, PullRequestDefaultTopic, consumerGroupID, ch, errors)
}

// CreatePullRequestConsumerForTopic will stream data from the topic into the provided channel
func CreatePullRequestConsumerForTopic(kafkaServers []string, schemaRegistryServers []string, topic string, consumerGroupID string, ch chan PullRequest, errors chan<- error) (<-chan bool, chan<- bool) {
	done := make(chan bool, 1)
	closed := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		consumerCallbacks := kafka.ConsumerCallbacks{
			OnDataReceived: func(msg kafka.Message) {
				var object PullRequest
				if err := json.Unmarshal([]byte(msg.Value), &object); err != nil {
					errors <- fmt.Errorf("error unmarshaling json data into PullRequest: %s", err)
					return
				}
				ch <- object
			},
			OnError: func(err error) {
				errors <- err
			},
			OnNotification: func(notification *cluster.Notification) {
			},
		}
		consumer, err := kafka.NewAvroConsumer(kafkaServers, schemaRegistryServers, topic, consumerGroupID, consumerCallbacks)
		if err != nil {
			errors <- err
			return
		}
		go consumer.Consume()
		select {
		case <-closed:
			consumer.Close()
		}
	}()
	return done, closed
}
