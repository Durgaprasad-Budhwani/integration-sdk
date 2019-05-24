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

// CommitTopic is the default topic name
const CommitTopic datamodel.TopicNameType = "sourcecode_Commit_topic"

// CommitStream is the default stream name
const CommitStream datamodel.TopicNameType = "sourcecode_Commit_stream"

// CommitTable is the default table name
const CommitTable datamodel.TopicNameType = "sourcecode_Commit"

// CommitModelName is the model name
const CommitModelName datamodel.ModelNameType = "sourcecode.Commit"

// Commit the commit is a specific change in a repo
type Commit struct {
	// built in types

	ID         string `json:"commit_id" yaml:"commit_id" faker:"-"`
	RefID      string `json:"ref_id" yaml:"ref_id" faker:"-"`
	RefType    string `json:"ref_type" yaml:"ref_type" faker:"-"`
	CustomerID string `json:"customer_id" yaml:"customer_id" faker:"-"`
	Hashcode   string `json:"hashcode" yaml:"hashcode" faker:"-"`
	// custom types

	// RepoID the unique id for the repo
	RepoID string `json:"repo_id" yaml:"repo_id" faker:"-"`
	// Sha the unique sha for the commit
	Sha string `json:"sha" yaml:"sha" faker:"sha"`
	// Message the commit message
	Message string `json:"message" yaml:"message" faker:"commit_message"`
	// URL the url to the commit detail
	URL string `json:"url" yaml:"url" faker:"url"`
	// CreatedAt the timestamp in UTC that the commit was created
	CreatedAt int64 `json:"created_ts" yaml:"created_ts" faker:"-"`
	// Branch the branch that the commit was made to
	Branch string `json:"branch" yaml:"branch" faker:"-"`
	// Additions the number of additions for the commit
	Additions int64 `json:"additions" yaml:"additions" faker:"-"`
	// Deletions the number of deletions for the commit
	Deletions int64 `json:"deletions" yaml:"deletions" faker:"-"`
	// FilesChanged the number of files changed for the commit
	FilesChanged int64 `json:"files_changed" yaml:"files_changed" faker:"-"`
	// AuthorRefID the author ref_id in the source system
	AuthorRefID string `json:"author_ref_id" yaml:"author_ref_id" faker:"-"`
	// Ordinal the order of the commit in the commit stream
	Ordinal int64 `json:"ordinal" yaml:"ordinal" faker:"-"`
	// Loc the number of lines in the commit
	Loc int64 `json:"loc" yaml:"loc" faker:"-"`
	// Sloc the number of source lines in the commit
	Sloc int64 `json:"sloc" yaml:"sloc" faker:"-"`
	// Comments the number of comment lines in the commit
	Comments int64 `json:"comments" yaml:"comments" faker:"-"`
	// Blanks the number of blank lines in the commit
	Blanks int64 `json:"blanks" yaml:"blanks" faker:"-"`
	// Size the size of all files in the commit
	Size int64 `json:"size" yaml:"size" faker:"-"`
	// Complexity the complexity value for the change
	Complexity int64 `json:"complexity" yaml:"complexity" faker:"-"`
	// GpgSigned if the commit was signed
	GpgSigned bool `json:"gpg_signed" yaml:"gpg_signed" faker:"-"`
	// Excluded if the commit was excluded
	Excluded bool `json:"excluded" yaml:"excluded" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Commit)(nil)

func toCommitObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toCommitObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toCommitObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toCommitObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toCommitObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toCommitObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toCommitObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toCommitObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toCommitObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toCommitObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toCommitObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toCommitObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toCommitObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case map[string]interface{}:
		return o
	case *map[string]interface{}:
		return v
	case *Commit:
		return v.ToMap()
	case Commit:
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
			arr = append(arr, toCommitObject(av, isavro, false, ""))
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of Commit
func (o *Commit) String() string {
	return fmt.Sprintf("sourcecode.Commit<%s>", o.ID)
}

func (o *Commit) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Commit) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Commit", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetRefID returns the RefID for the object
func (o *Commit) GetRefID() string {
	o.RefID = o.Sha
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Commit) IsMaterialized() bool {
	return true
}

// MaterializedName returns the name of the materialized table
func (o *Commit) MaterializedName() string {
	return "sourcecode_commit"
}

// Clone returns an exact copy of Commit
func (o *Commit) Clone() *Commit {
	c := new(Commit)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Commit) Anon() *Commit {
	c := new(Commit)
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
func (o *Commit) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Commit) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	// make sure that these have values if empty
	o.setDefaults()
	return nil
}

var cachedCodecCommit *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *Commit) GetAvroCodec() *goavro.Codec {
	if cachedCodecCommit == nil {
		c, err := CreateCommitAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecCommit = c
	}
	return cachedCodecCommit
}

// ToAvroBinary returns the data as Avro binary data
func (o *Commit) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *Commit) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Commit objects are equal
func (o *Commit) IsEqual(other *Commit) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Commit) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	return map[string]interface{}{
		"commit_id":     o.GetID(),
		"ref_id":        o.GetRefID(),
		"ref_type":      o.RefType,
		"customer_id":   o.CustomerID,
		"hashcode":      o.Hash(),
		"repo_id":       toCommitObject(o.RepoID, isavro, false, "string"),
		"sha":           toCommitObject(o.Sha, isavro, false, "string"),
		"message":       toCommitObject(o.Message, isavro, false, "string"),
		"url":           toCommitObject(o.URL, isavro, false, "string"),
		"created_ts":    toCommitObject(o.CreatedAt, isavro, false, "long"),
		"branch":        toCommitObject(o.Branch, isavro, false, "string"),
		"additions":     toCommitObject(o.Additions, isavro, false, "long"),
		"deletions":     toCommitObject(o.Deletions, isavro, false, "long"),
		"files_changed": toCommitObject(o.FilesChanged, isavro, false, "long"),
		"author_ref_id": toCommitObject(o.AuthorRefID, isavro, false, "string"),
		"ordinal":       toCommitObject(o.Ordinal, isavro, false, "long"),
		"loc":           toCommitObject(o.Loc, isavro, false, "long"),
		"sloc":          toCommitObject(o.Sloc, isavro, false, "long"),
		"comments":      toCommitObject(o.Comments, isavro, false, "long"),
		"blanks":        toCommitObject(o.Blanks, isavro, false, "long"),
		"size":          toCommitObject(o.Size, isavro, false, "long"),
		"complexity":    toCommitObject(o.Complexity, isavro, false, "long"),
		"gpg_signed":    toCommitObject(o.GpgSigned, isavro, false, "boolean"),
		"excluded":      toCommitObject(o.Excluded, isavro, false, "boolean"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Commit) FromMap(kv map[string]interface{}) {
	if val, ok := kv["commit_id"].(string); ok {
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
		if val == nil {
			o.RepoID = ""
		} else {
			o.RepoID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["sha"].(string); ok {
		o.Sha = val
	} else {
		val := kv["sha"]
		if val == nil {
			o.Sha = ""
		} else {
			o.Sha = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["message"].(string); ok {
		o.Message = val
	} else {
		val := kv["message"]
		if val == nil {
			o.Message = ""
		} else {
			o.Message = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["url"].(string); ok {
		o.URL = val
	} else {
		val := kv["url"]
		if val == nil {
			o.URL = ""
		} else {
			o.URL = fmt.Sprintf("%v", val)
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
	if val, ok := kv["branch"].(string); ok {
		o.Branch = val
	} else {
		val := kv["branch"]
		if val == nil {
			o.Branch = ""
		} else {
			o.Branch = fmt.Sprintf("%v", val)
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
	if val, ok := kv["files_changed"].(int64); ok {
		o.FilesChanged = val
	} else {
		val := kv["files_changed"]
		if val == nil {
			o.FilesChanged = number.ToInt64Any(nil)
		} else {
			o.FilesChanged = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["author_ref_id"].(string); ok {
		o.AuthorRefID = val
	} else {
		val := kv["author_ref_id"]
		if val == nil {
			o.AuthorRefID = ""
		} else {
			o.AuthorRefID = fmt.Sprintf("%v", val)
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
	if val, ok := kv["size"].(int64); ok {
		o.Size = val
	} else {
		val := kv["size"]
		if val == nil {
			o.Size = number.ToInt64Any(nil)
		} else {
			o.Size = number.ToInt64Any(val)
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
	if val, ok := kv["gpg_signed"].(bool); ok {
		o.GpgSigned = val
	} else {
		val := kv["gpg_signed"]
		if val == nil {
			o.GpgSigned = number.ToBoolAny(nil)
		} else {
			o.GpgSigned = number.ToBoolAny(val)
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
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *Commit) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.CustomerID)
	args = append(args, o.RepoID)
	args = append(args, o.Sha)
	args = append(args, o.Message)
	args = append(args, o.URL)
	args = append(args, o.CreatedAt)
	args = append(args, o.Branch)
	args = append(args, o.Additions)
	args = append(args, o.Deletions)
	args = append(args, o.FilesChanged)
	args = append(args, o.AuthorRefID)
	args = append(args, o.Ordinal)
	args = append(args, o.Loc)
	args = append(args, o.Sloc)
	args = append(args, o.Comments)
	args = append(args, o.Blanks)
	args = append(args, o.Size)
	args = append(args, o.Complexity)
	args = append(args, o.GpgSigned)
	args = append(args, o.Excluded)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// CreateCommitAvroSchemaSpec creates the avro schema specification for Commit
func CreateCommitAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "sourcecode",
		"name":         "Commit",
		"connect.name": "sourcecode.Commit",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "commit_id",
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
				"name": "sha",
				"type": "string",
			},
			map[string]interface{}{
				"name": "message",
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
				"name": "branch",
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
				"name": "files_changed",
				"type": "long",
			},
			map[string]interface{}{
				"name": "author_ref_id",
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
				"name": "comments",
				"type": "long",
			},
			map[string]interface{}{
				"name": "blanks",
				"type": "long",
			},
			map[string]interface{}{
				"name": "size",
				"type": "long",
			},
			map[string]interface{}{
				"name": "complexity",
				"type": "long",
			},
			map[string]interface{}{
				"name": "gpg_signed",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "excluded",
				"type": "boolean",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// CreateCommitAvroSchema creates the avro schema for Commit
func CreateCommitAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(CreateCommitAvroSchemaSpec())
}

// TransformCommitFunc is a function for transforming Commit during processing
type TransformCommitFunc func(input *Commit) (*Commit, error)

// CreateCommitPipe creates a pipe for processing Commit items
func CreateCommitPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformCommitFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := CreateCommitInputStream(input, errors)
	var stream chan Commit
	if len(transforms) > 0 {
		stream = make(chan Commit, 1000)
	} else {
		stream = inch
	}
	outdone := CreateCommitOutputStream(output, stream, errors)
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

// CreateCommitInputStreamDir creates a channel for reading Commit as JSON newlines from a directory of files
func CreateCommitInputStreamDir(dir string, errors chan<- error, transforms ...TransformCommitFunc) (chan Commit, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/sourcecode/commit\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Commit)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for commit")
		ch := make(chan Commit)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return CreateCommitInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Commit)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// CreateCommitInputStreamFile creates an channel for reading Commit as JSON newlines from filename
func CreateCommitInputStreamFile(filename string, errors chan<- error, transforms ...TransformCommitFunc) (chan Commit, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Commit)
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
			ch := make(chan Commit)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return CreateCommitInputStream(f, errors, transforms...)
}

// CreateCommitInputStream creates an channel for reading Commit as JSON newlines from stream
func CreateCommitInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformCommitFunc) (chan Commit, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Commit, 1000)
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
			var item Commit
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

// CreateCommitOutputStreamDir will output json newlines from channel and save in dir
func CreateCommitOutputStreamDir(dir string, ch chan Commit, errors chan<- error, transforms ...TransformCommitFunc) <-chan bool {
	fp := filepath.Join(dir, "/sourcecode/commit\\.json(\\.gz)?$")
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
	return CreateCommitOutputStream(gz, ch, errors, transforms...)
}

// CreateCommitOutputStream will output json newlines from channel to the stream
func CreateCommitOutputStream(stream io.WriteCloser, ch chan Commit, errors chan<- error, transforms ...TransformCommitFunc) <-chan bool {
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

// CommitSendEvent is an event detail for sending data
type CommitSendEvent struct {
	Commit  Commit
	Headers map[string]string
}

// CreateCommitProducer will stream data from the channel
func CreateCommitProducer(producer event.Producer, ch chan CommitSendEvent, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		ctx := context.Background()
		for item := range ch {
			binary, codec, err := item.Commit.ToAvroBinary()
			if err != nil {
				errors <- fmt.Errorf("error encoding %s to avro binary data. %v", item.Commit.String(), err)
				return
			}
			headers := map[string]string{
				"customer_id": item.Commit.CustomerID,
			}
			if item.Headers != nil {
				for k, v := range item.Headers {
					headers[k] = v
				}
			}
			msg := event.Message{
				Key:     item.Commit.ID,
				Value:   binary,
				Codec:   codec,
				Headers: headers,
			}
			if err := producer.Send(ctx, msg); err != nil {
				errors <- fmt.Errorf("error sending %s. %v", item.Commit.String(), err)
			}
		}
	}()
	return done
}

// CommitReceiveEvent is an event detail for receiving data
type CommitReceiveEvent struct {
	Commit  Commit
	Message event.Message
}

// CreateCommitConsumer will stream data from the topic into the provided channel
func CreateCommitConsumer(factory event.ConsumerFactory, topic datamodel.TopicNameType, ch chan CommitReceiveEvent, errors chan<- error) (<-chan bool, chan<- bool) {
	done := make(chan bool, 1)
	closed := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		callback := event.ConsumerCallback{
			OnDataReceived: func(msg event.Message) error {
				var object Commit
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into sourcecode.Commit: %s", err)
				}
				msg.Codec = object.GetAvroCodec() // match the codec
				ch <- CommitReceiveEvent{object, msg}
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
