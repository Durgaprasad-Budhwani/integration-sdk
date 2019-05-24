// DO NOT EDIT - this code is generated

package datamodel

import (
	"strings"

	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-datamodel/codequality"
	"github.com/pinpt/go-datamodel/customer"
	"github.com/pinpt/go-datamodel/sourcecode"
	"github.com/pinpt/go-datamodel/work"
)

// New returns a new instanceof from a ModelNameType
func New(name datamodel.ModelNameType) datamodel.Model {
	switch name {
	case "codequality.Metric":
		return new(codequality.Metric)
	case "codequality.Project":
		return new(codequality.Project)
	case "customer.User":
		return new(customer.User)
	case "customer.Team":
		return new(customer.Team)
	case "customer.CostCenter":
		return new(customer.CostCenter)
	case "sourcecode.Repo":
		return new(sourcecode.Repo)
	case "sourcecode.User":
		return new(sourcecode.User)
	case "sourcecode.Branch":
		return new(sourcecode.Branch)
	case "sourcecode.Changelog":
		return new(sourcecode.Changelog)
	case "sourcecode.Commit":
		return new(sourcecode.Commit)
	case "sourcecode.CommitFile":
		return new(sourcecode.CommitFile)
	case "sourcecode.PullRequest":
		return new(sourcecode.PullRequest)
	case "work.CustomField":
		return new(work.CustomField)
	case "work.Issue":
		return new(work.Issue)
	case "work.Project":
		return new(work.Project)
	case "work.Sprint":
		return new(work.Sprint)
	case "work.User":
		return new(work.User)
	case "work.Changelog":
		return new(work.Changelog)
	}
	panic("invalid type specific: " + name)
}

// NewFromTopic returns a new instanceof from a TopicNameType
func NewFromTopic(name datamodel.TopicNameType) datamodel.Model {
	switch name {
	case "customer_Team_topic":
		return new(customer.Team)
	case "customer_CostCenter_topic":
		return new(customer.CostCenter)
	case "customer_User_topic":
		return new(customer.User)
	case "sourcecode_Repo_topic":
		return new(sourcecode.Repo)
	case "sourcecode_User_topic":
		return new(sourcecode.User)
	case "sourcecode_Branch_topic":
		return new(sourcecode.Branch)
	case "sourcecode_Changelog_topic":
		return new(sourcecode.Changelog)
	case "sourcecode_Commit_topic":
		return new(sourcecode.Commit)
	case "sourcecode_CommitFile_topic":
		return new(sourcecode.CommitFile)
	case "sourcecode_PullRequest_topic":
		return new(sourcecode.PullRequest)
	case "work_CustomField_topic":
		return new(work.CustomField)
	case "work_Issue_topic":
		return new(work.Issue)
	case "work_Project_topic":
		return new(work.Project)
	case "work_Sprint_topic":
		return new(work.Sprint)
	case "work_User_topic":
		return new(work.User)
	case "work_Changelog_topic":
		return new(work.Changelog)
	case "codequality_Metric_topic":
		return new(codequality.Metric)
	case "codequality_Project_topic":
		return new(codequality.Project)
	}
	panic("invalid type specific: " + name)
}

// GetMaterializedTopics returns an array of topics to be materialized
func GetMaterializedTopics() []datamodel.TopicNameType {
	return []datamodel.TopicNameType{
		datamodel.TopicNameType(customer_Team_topic),
		datamodel.TopicNameType(customer_CostCenter_topic),
		datamodel.TopicNameType(customer_User_topic),
		datamodel.TopicNameType(sourcecode_CommitFile_topic),
		datamodel.TopicNameType(sourcecode_PullRequest_topic),
		datamodel.TopicNameType(sourcecode_Repo_topic),
		datamodel.TopicNameType(sourcecode_Branch_topic),
		datamodel.TopicNameType(sourcecode_Changelog_topic),
		datamodel.TopicNameType(sourcecode_Commit_topic),
		datamodel.TopicNameType(work_Project_topic),
	}
}
