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

	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	pstrings "github.com/pinpt/go-common/strings"
	"github.com/pinpt/integration-sdk/util"
)

// CommitFileDefaultTopic is the default topic name
const CommitFileDefaultTopic = "sourcecode_CommitFile_topic"

// CommitFileDefaultStream is the default stream name
const CommitFileDefaultStream = "sourcecode_CommitFile_stream"

// CommitFileDefaultTable is the default table name
const CommitFileDefaultTable = "sourcecode_CommitFile"

// CommitFile the file change for a specific commit
type CommitFile struct {
	// built in types

	ID         string `json:"commit_file_id" yaml:"commit_file_id"`
	RefID      string `json:"ref_id" yaml:"ref_id"`
	RefType    string `json:"ref_type" yaml:"ref_type"`
	CustomerID string `json:"customer_id" yaml:"customer_id"`
	Hashcode   string `json:"hashcode" yaml:"hashcode"`

	// custom types

	// CommitID the unique id for the commit
	CommitID string `json:"commit_id" yaml:"commit_id"`
	// RepoID the unique id for the repo
	RepoID string `json:"repo_id" yaml:"repo_id"`
	// Filename the filename
	Filename string `json:"filename" yaml:"filename"`
	// Additions the number of additions for the commit file
	Additions int64 `json:"additions" yaml:"additions"`
	// Deletions the number of deletions for the commit file
	Deletions int64 `json:"deletions" yaml:"deletions"`
	// Status the status of the change
	Status string `json:"status" yaml:"status"`
	// Binary indicates if the file was detected to be a binary file
	Binary bool `json:"binary" yaml:"binary"`
	// Language the language that was detected for the file
	Language string `json:"language" yaml:"language"`
	// Excluded if the file was excluded from processing
	Excluded bool `json:"excluded" yaml:"excluded"`
	// ExcludedReason if the file was excluded, the reason
	ExcludedReason string `json:"excluded_reason" yaml:"excluded_reason"`
	// Ordinal the order value for the file in the change set
	Ordinal int64 `json:"ordinal" yaml:"ordinal"`
	// Loc the number of lines in the file
	Loc int64 `json:"loc" yaml:"loc"`
	// Sloc the number of source lines in the file
	Sloc int64 `json:"sloc" yaml:"sloc"`
	// Blanks the number of blank lines in the file
	Blanks int64 `json:"blanks" yaml:"blanks"`
	// Comments the number of comment lines in the file
	Comments int64 `json:"comments" yaml:"comments"`
	// Complexity the complexity value for the file change
	Complexity int64 `json:"complexity" yaml:"complexity"`
	// License the license which was detected for the file
	License *string `json:"license" yaml:"license"`
	// LicenseConfidence the license confidence from the detection engine
	LicenseConfidence float64 `json:"license_confidence" yaml:"license_confidence"`
}

// String returns a string representation of CommitFile
func (o *CommitFile) String() string {
	return fmt.Sprintf("sourcecode.CommitFile<%s>", o.ID)
}

func (o *CommitFile) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *CommitFile) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("CommitFile", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetRefID returns the RefID for the object
func (o *CommitFile) GetRefID() string {
	return o.RefID
}

// MarshalJSON returns the bytes for marshaling to json
func (o *CommitFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *CommitFile) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	// make sure that these have values if empty
	o.setDefaults()
	return nil
}

var cachedCodecCommitFile *goavro.Codec

// ToAvroBinary returns the data as Avro binary data
func (o *CommitFile) ToAvroBinary() ([]byte, *goavro.Codec, error) {
	if cachedCodecCommitFile == nil {
		c, err := CreateCommitFileAvroSchema()
		if err != nil {
			return nil, nil, err
		}
		cachedCodecCommitFile = c
	}
	// Convert native Go form to binary Avro data
	buf, err := cachedCodecCommitFile.BinaryFromNative(nil, o.ToMap())
	return buf, cachedCodecCommitFile, err
}

// Stringify returns the object in JSON format as a string
func (o *CommitFile) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two CommitFile objects are equal
func (o *CommitFile) IsEqual(other *CommitFile) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *CommitFile) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"commit_file_id":     o.GetID(),
		"ref_id":             o.GetRefID(),
		"ref_type":           o.RefType,
		"customer_id":        o.CustomerID,
		"hashcode":           o.Hash(),
		"commit_id":          o.CommitID,
		"repo_id":            o.RepoID,
		"filename":           o.Filename,
		"additions":          o.Additions,
		"deletions":          o.Deletions,
		"status":             o.Status,
		"binary":             o.Binary,
		"language":           o.Language,
		"excluded":           o.Excluded,
		"excluded_reason":    o.ExcludedReason,
		"ordinal":            o.Ordinal,
		"loc":                o.Loc,
		"sloc":               o.Sloc,
		"blanks":             o.Blanks,
		"comments":           o.Comments,
		"complexity":         o.Complexity,
		"license":            o.License,
		"license_confidence": o.LicenseConfidence,
	}
}

// FromMap attempts to load data into object from a map
func (o *CommitFile) FromMap(kv map[string]interface{}) {
	if val, ok := kv["commit_file_id"].(string); ok {
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
	if val, ok := kv["commit_id"].(string); ok {
		o.CommitID = val
	} else {
		val := kv["commit_id"]
		if val == nil {
			o.CommitID = ""
		} else {
			o.CommitID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["repo_id"].(string); ok {
		o.RepoID = val
	} else {
		val := kv["repo_id"]
		if val == nil {
			o.RepoID = ""
		} else {
			o.RepoID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["filename"].(string); ok {
		o.Filename = val
	} else {
		val := kv["filename"]
		if val == nil {
			o.Filename = ""
		} else {
			o.Filename = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["additions"].(int64); ok {
		o.Additions = val
	} else {
		val := kv["additions"]
		if val == nil {
			o.Additions = number.ToInt64Any(nil)
		} else {
			o.Additions = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["deletions"].(int64); ok {
		o.Deletions = val
	} else {
		val := kv["deletions"]
		if val == nil {
			o.Deletions = number.ToInt64Any(nil)
		} else {
			o.Deletions = number.ToInt64Any(val)
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
	if val, ok := kv["binary"].(bool); ok {
		o.Binary = val
	} else {
		val := kv["binary"]
		if val == nil {
			o.Binary = number.ToBoolAny(nil)
		} else {
			o.Binary = number.ToBoolAny(val)
		}
	}
	if val, ok := kv["language"].(string); ok {
		o.Language = val
	} else {
		val := kv["language"]
		if val == nil {
			o.Language = ""
		} else {
			o.Language = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["excluded"].(bool); ok {
		o.Excluded = val
	} else {
		val := kv["excluded"]
		if val == nil {
			o.Excluded = number.ToBoolAny(nil)
		} else {
			o.Excluded = number.ToBoolAny(val)
		}
	}
	if val, ok := kv["excluded_reason"].(string); ok {
		o.ExcludedReason = val
	} else {
		val := kv["excluded_reason"]
		if val == nil {
			o.ExcludedReason = ""
		} else {
			o.ExcludedReason = fmt.Sprintf("%v", val)
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
	if val, ok := kv["loc"].(int64); ok {
		o.Loc = val
	} else {
		val := kv["loc"]
		if val == nil {
			o.Loc = number.ToInt64Any(nil)
		} else {
			o.Loc = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["sloc"].(int64); ok {
		o.Sloc = val
	} else {
		val := kv["sloc"]
		if val == nil {
			o.Sloc = number.ToInt64Any(nil)
		} else {
			o.Sloc = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["blanks"].(int64); ok {
		o.Blanks = val
	} else {
		val := kv["blanks"]
		if val == nil {
			o.Blanks = number.ToInt64Any(nil)
		} else {
			o.Blanks = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["comments"].(int64); ok {
		o.Comments = val
	} else {
		val := kv["comments"]
		if val == nil {
			o.Comments = number.ToInt64Any(nil)
		} else {
			o.Comments = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["complexity"].(int64); ok {
		o.Complexity = val
	} else {
		val := kv["complexity"]
		if val == nil {
			o.Complexity = number.ToInt64Any(nil)
		} else {
			o.Complexity = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["license"].(*string); ok {
		o.License = val
	} else if val, ok := kv["license"].(string); ok {
		o.License = &val
	} else {
		val := kv["license"]
		if val == nil {
			o.License = pstrings.Pointer("")
		} else {
			o.License = pstrings.Pointer(fmt.Sprintf("%v", val))
		}
	}
	if val, ok := kv["license_confidence"].(float64); ok {
		o.LicenseConfidence = val
	} else {
		val := kv["license_confidence"]
		if val == nil {
			o.LicenseConfidence = number.ToFloat64Any(nil)
		} else {
			o.LicenseConfidence = number.ToFloat64Any(val)
		}
	}
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *CommitFile) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.CommitID)
	args = append(args, o.RepoID)
	args = append(args, o.Filename)
	args = append(args, o.Additions)
	args = append(args, o.Deletions)
	args = append(args, o.Status)
	args = append(args, o.Binary)
	args = append(args, o.Language)
	args = append(args, o.Excluded)
	args = append(args, o.ExcludedReason)
	args = append(args, o.Ordinal)
	args = append(args, o.Loc)
	args = append(args, o.Sloc)
	args = append(args, o.Blanks)
	args = append(args, o.Comments)
	args = append(args, o.Complexity)
	args = append(args, o.License)
	args = append(args, o.LicenseConfidence)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// CreateCommitFileAvroSchemaSpec creates the avro schema specification for CommitFile
func CreateCommitFileAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "sourcecode",
		"name":         "CommitFile",
		"connect.name": "sourcecode.CommitFile",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "commit_file_id",
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
				"name": "commit_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "repo_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "filename",
				"type": "string",
			},
			map[string]interface{}{
				"name": "additions",
				"type": "long",
			},
			map[string]interface{}{
				"name": "deletions",
				"type": "long",
			},
			map[string]interface{}{
				"name": "status",
				"type": "string",
			},
			map[string]interface{}{
				"name": "binary",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "language",
				"type": "string",
			},
			map[string]interface{}{
				"name": "excluded",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "excluded_reason",
				"type": "string",
			},
			map[string]interface{}{
				"name": "ordinal",
				"type": "long",
			},
			map[string]interface{}{
				"name": "loc",
				"type": "long",
			},
			map[string]interface{}{
				"name": "sloc",
				"type": "long",
			},
			map[string]interface{}{
				"name": "blanks",
				"type": "long",
			},
			map[string]interface{}{
				"name": "comments",
				"type": "long",
			},
			map[string]interface{}{
				"name": "complexity",
				"type": "long",
			},
			map[string]interface{}{
				"name":    "license",
				"type":    []string{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name": "license_confidence",
				"type": "float",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// CreateCommitFileAvroSchema creates the avro schema for CommitFile
func CreateCommitFileAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(CreateCommitFileAvroSchemaSpec())
}

// TransformCommitFileFunc is a function for transforming CommitFile during processing
type TransformCommitFileFunc func(input *CommitFile) (*CommitFile, error)

// CreateCommitFilePipe creates a pipe for processing CommitFile items
func CreateCommitFilePipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformCommitFileFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := CreateCommitFileInputStream(input, errors)
	var stream chan CommitFile
	if len(transforms) > 0 {
		stream = make(chan CommitFile, 1000)
	} else {
		stream = inch
	}
	outdone := CreateCommitFileOutputStream(output, stream, errors)
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

// CreateCommitFileInputStreamDir creates a channel for reading CommitFile as JSON newlines from a directory of files
func CreateCommitFileInputStreamDir(dir string, errors chan<- error, transforms ...TransformCommitFileFunc) (chan CommitFile, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/sourcecode/commit_file\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan CommitFile)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for commit_file")
		ch := make(chan CommitFile)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return CreateCommitFileInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan CommitFile)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// CreateCommitFileInputStreamFile creates an channel for reading CommitFile as JSON newlines from filename
func CreateCommitFileInputStreamFile(filename string, errors chan<- error, transforms ...TransformCommitFileFunc) (chan CommitFile, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan CommitFile)
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
			ch := make(chan CommitFile)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return CreateCommitFileInputStream(f, errors, transforms...)
}

// CreateCommitFileInputStream creates an channel for reading CommitFile as JSON newlines from stream
func CreateCommitFileInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformCommitFileFunc) (chan CommitFile, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan CommitFile, 1000)
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
			var item CommitFile
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

// CreateCommitFileOutputStreamDir will output json newlines from channel and save in dir
func CreateCommitFileOutputStreamDir(dir string, ch chan CommitFile, errors chan<- error, transforms ...TransformCommitFileFunc) <-chan bool {
	fp := filepath.Join(dir, "/sourcecode/commit_file\\.json(\\.gz)?$")
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
	return CreateCommitFileOutputStream(gz, ch, errors, transforms...)
}

// CreateCommitFileOutputStream will output json newlines from channel to the stream
func CreateCommitFileOutputStream(stream io.WriteCloser, ch chan CommitFile, errors chan<- error, transforms ...TransformCommitFileFunc) <-chan bool {
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

// CreateCommitFileProducer will stream data from the channel
func CreateCommitFileProducer(producer util.Producer, ch chan CommitFile, errors chan<- error) <-chan bool {
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

// CreateCommitFileConsumer will stream data from the default topic into the provided channel
func CreateCommitFileConsumer(factory util.ConsumerFactory, topic string, ch chan CommitFile, errors chan<- error) (<-chan bool, chan<- bool) {
	return CreateCommitFileConsumerForTopic(factory, CommitFileDefaultTopic, ch, errors)
}

// CreateCommitFileConsumerForTopic will stream data from the topic into the provided channel
func CreateCommitFileConsumerForTopic(factory util.ConsumerFactory, topic string, ch chan CommitFile, errors chan<- error) (<-chan bool, chan<- bool) {
	done := make(chan bool, 1)
	closed := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		callback := util.ConsumerCallback{
			OnDataReceived: func(key []byte, value []byte) error {
				var object CommitFile
				if err := json.Unmarshal(value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into CommitFile: %s", err)
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