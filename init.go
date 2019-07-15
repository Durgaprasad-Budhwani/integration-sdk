// DO NOT EDIT - this code is generated

package datamodel

import (
	"github.com/pinpt/go-common/datamodel"

	dm_admin "github.com/pinpt/go-datamodel/admin"
	dm_agent "github.com/pinpt/go-datamodel/agent"
	dm_auth "github.com/pinpt/go-datamodel/auth"
	dm_codequality "github.com/pinpt/go-datamodel/codequality"
	dm_customer "github.com/pinpt/go-datamodel/customer"
	dm_ops_db "github.com/pinpt/go-datamodel/ops/db"
	dm_pipeline_activity "github.com/pinpt/go-datamodel/pipeline/activity"
	dm_pipeline_customer "github.com/pinpt/go-datamodel/pipeline/customer"
	dm_pipeline_integration "github.com/pinpt/go-datamodel/pipeline/integration"
	dm_pipeline_signal "github.com/pinpt/go-datamodel/pipeline/signal"
	dm_pipeline_sourcecode "github.com/pinpt/go-datamodel/pipeline/sourcecode"
	dm_pipeline_work "github.com/pinpt/go-datamodel/pipeline/work"
	dm_sourcecode "github.com/pinpt/go-datamodel/sourcecode"
	dm_web "github.com/pinpt/go-datamodel/web"
	dm_work "github.com/pinpt/go-datamodel/work"
)

// New returns a new instanceof from a ModelNameType
func New(name datamodel.ModelNameType) datamodel.Model {
	switch name {
	case "admin.Agent":
		return new(dm_admin.Agent)
	case "pipeline.customer.User":
		return new(dm_pipeline_customer.User)
	case "pipeline.integration.User":
		return new(dm_pipeline_integration.User)
	case "work.Changelog":
		return new(dm_work.Changelog)
	case "work.CustomField":
		return new(dm_work.CustomField)
	case "work.Issue":
		return new(dm_work.Issue)
	case "work.Project":
		return new(dm_work.Project)
	case "work.Sprint":
		return new(dm_work.Sprint)
	case "work.User":
		return new(dm_work.User)
	case "agent.ExportResponse":
		return new(dm_agent.ExportResponse)
	case "agent.ExportTrigger":
		return new(dm_agent.ExportTrigger)
	case "agent.Event":
		return new(dm_agent.Event)
	case "agent.IntegrationRequest":
		return new(dm_agent.IntegrationRequest)
	case "agent.ProjectRequest":
		return new(dm_agent.ProjectRequest)
	case "agent.Enabled":
		return new(dm_agent.Enabled)
	case "agent.EnrollRequest":
		return new(dm_agent.EnrollRequest)
	case "agent.ExportRequest":
		return new(dm_agent.ExportRequest)
	case "agent.IntegrationResponse":
		return new(dm_agent.IntegrationResponse)
	case "agent.ProjectResponse":
		return new(dm_agent.ProjectResponse)
	case "agent.EnrollResponse":
		return new(dm_agent.EnrollResponse)
	case "sourcecode.Branch":
		return new(dm_sourcecode.Branch)
	case "sourcecode.Commit":
		return new(dm_sourcecode.Commit)
	case "sourcecode.PullRequest":
		return new(dm_sourcecode.PullRequest)
	case "sourcecode.PullRequestComment":
		return new(dm_sourcecode.PullRequestComment)
	case "sourcecode.PullRequestReview":
		return new(dm_sourcecode.PullRequestReview)
	case "sourcecode.Repo":
		return new(dm_sourcecode.Repo)
	case "sourcecode.User":
		return new(dm_sourcecode.User)
	case "sourcecode.Blame":
		return new(dm_sourcecode.Blame)
	case "web.Hook":
		return new(dm_web.Hook)
	case "auth.ACLGrant":
		return new(dm_auth.ACLGrant)
	case "ops.db.Change":
		return new(dm_ops_db.Change)
	case "pipeline.work.CustomField":
		return new(dm_pipeline_work.CustomField)
	case "codequality.Metric":
		return new(dm_codequality.Metric)
	case "codequality.Project":
		return new(dm_codequality.Project)
	case "customer.User":
		return new(dm_customer.User)
	case "customer.CostCenter":
		return new(dm_customer.CostCenter)
	case "customer.Team":
		return new(dm_customer.Team)
	case "pipeline.activity.Activity":
		return new(dm_pipeline_activity.Activity)
	case "pipeline.signal.Signal":
		return new(dm_pipeline_signal.Signal)
	case "pipeline.sourcecode.Commit":
		return new(dm_pipeline_sourcecode.Commit)
	}
	return nil
}

// NewFromTopic returns a new instanceof from a TopicNameType
func NewFromTopic(name datamodel.TopicNameType) datamodel.Model {
	switch name {
	case "agent_ExportRequest_topic":
		return new(dm_agent.ExportRequest)
	case "agent_IntegrationResponse_topic":
		return new(dm_agent.IntegrationResponse)
	case "agent_ProjectResponse_topic":
		return new(dm_agent.ProjectResponse)
	case "agent_EnrollResponse_topic":
		return new(dm_agent.EnrollResponse)
	case "agent_IntegrationRequest_topic":
		return new(dm_agent.IntegrationRequest)
	case "agent_ProjectRequest_topic":
		return new(dm_agent.ProjectRequest)
	case "agent_Enabled_topic":
		return new(dm_agent.Enabled)
	case "agent_EnrollRequest_topic":
		return new(dm_agent.EnrollRequest)
	case "agent_Event_topic":
		return new(dm_agent.Event)
	case "agent_ExportResponse_topic":
		return new(dm_agent.ExportResponse)
	case "agent_ExportTrigger_topic":
		return new(dm_agent.ExportTrigger)
	case "sourcecode_Blame_topic":
		return new(dm_sourcecode.Blame)
	case "sourcecode_Branch_topic":
		return new(dm_sourcecode.Branch)
	case "sourcecode_Commit_topic":
		return new(dm_sourcecode.Commit)
	case "sourcecode_PullRequest_topic":
		return new(dm_sourcecode.PullRequest)
	case "sourcecode_PullRequestComment_topic":
		return new(dm_sourcecode.PullRequestComment)
	case "sourcecode_PullRequestReview_topic":
		return new(dm_sourcecode.PullRequestReview)
	case "sourcecode_Repo_topic":
		return new(dm_sourcecode.Repo)
	case "sourcecode_User_topic":
		return new(dm_sourcecode.User)
	case "web_Hook_topic":
		return new(dm_web.Hook)
	case "auth_ACLGrant_topic":
		return new(dm_auth.ACLGrant)
	case "ops_db_Change_topic":
		return new(dm_ops_db.Change)
	case "pipeline_work_CustomField_topic":
		return new(dm_pipeline_work.CustomField)
	case "codequality_Metric_topic":
		return new(dm_codequality.Metric)
	case "codequality_Project_topic":
		return new(dm_codequality.Project)
	case "customer_User_topic":
		return new(dm_customer.User)
	case "customer_CostCenter_topic":
		return new(dm_customer.CostCenter)
	case "customer_Team_topic":
		return new(dm_customer.Team)
	case "pipeline_activity_Activity_topic":
		return new(dm_pipeline_activity.Activity)
	case "pipeline_signal_Signal_topic":
		return new(dm_pipeline_signal.Signal)
	case "pipeline_sourcecode_Commit_topic":
		return new(dm_pipeline_sourcecode.Commit)
	case "admin_Agent_topic":
		return new(dm_admin.Agent)
	case "pipeline_customer_User_topic":
		return new(dm_pipeline_customer.User)
	case "pipeline_integration_User_topic":
		return new(dm_pipeline_integration.User)
	case "work_Issue_topic":
		return new(dm_work.Issue)
	case "work_Project_topic":
		return new(dm_work.Project)
	case "work_Sprint_topic":
		return new(dm_work.Sprint)
	case "work_User_topic":
		return new(dm_work.User)
	case "work_Changelog_topic":
		return new(dm_work.Changelog)
	case "work_CustomField_topic":
		return new(dm_work.CustomField)
	}
	return nil
}

// GetMaterializedTopics returns an array of topics to be materialized
func GetMaterializedTopics() []datamodel.TopicNameType {
	return []datamodel.TopicNameType{
		datamodel.TopicNameType("admin_Agent_topic"),
		datamodel.TopicNameType("auth_ACLGrant_topic"),
		datamodel.TopicNameType("customer_User_topic"),
		datamodel.TopicNameType("customer_CostCenter_topic"),
		datamodel.TopicNameType("customer_Team_topic"),
		datamodel.TopicNameType("pipeline_activity_Activity_topic"),
		datamodel.TopicNameType("pipeline_integration_User_topic"),
		datamodel.TopicNameType("pipeline_signal_Signal_topic"),
		datamodel.TopicNameType("pipeline_sourcecode_Commit_topic"),
		datamodel.TopicNameType("pipeline_work_CustomField_topic"),
		datamodel.TopicNameType("sourcecode_User_topic"),
		datamodel.TopicNameType("sourcecode_Branch_topic"),
		datamodel.TopicNameType("sourcecode_PullRequest_topic"),
		datamodel.TopicNameType("sourcecode_PullRequestComment_topic"),
		datamodel.TopicNameType("sourcecode_PullRequestReview_topic"),
		datamodel.TopicNameType("sourcecode_Repo_topic"),
		datamodel.TopicNameType("work_Project_topic"),
		datamodel.TopicNameType("work_User_topic"),
	}
}

// GetTopics returns an array of topics that are configured
func GetTopics() []datamodel.TopicNameType {
	return []datamodel.TopicNameType{
		datamodel.TopicNameType("admin_Agent_topic"),
		datamodel.TopicNameType("agent_EnrollResponse_topic"),
		datamodel.TopicNameType("agent_Event_topic"),
		datamodel.TopicNameType("agent_ExportResponse_topic"),
		datamodel.TopicNameType("agent_IntegrationRequest_topic"),
		datamodel.TopicNameType("agent_ProjectRequest_topic"),
		datamodel.TopicNameType("agent_Enabled_topic"),
		datamodel.TopicNameType("agent_EnrollRequest_topic"),
		datamodel.TopicNameType("agent_ExportRequest_topic"),
		datamodel.TopicNameType("agent_ExportTrigger_topic"),
		datamodel.TopicNameType("agent_IntegrationResponse_topic"),
		datamodel.TopicNameType("agent_ProjectResponse_topic"),
		datamodel.TopicNameType("auth_ACLGrant_topic"),
		datamodel.TopicNameType("codequality_Metric_topic"),
		datamodel.TopicNameType("codequality_Project_topic"),
		datamodel.TopicNameType("customer_User_topic"),
		datamodel.TopicNameType("customer_CostCenter_topic"),
		datamodel.TopicNameType("customer_Team_topic"),
		datamodel.TopicNameType("ops_db_Change_topic"),
		datamodel.TopicNameType("pipeline_activity_Activity_topic"),
		datamodel.TopicNameType("pipeline_customer_User_topic"),
		datamodel.TopicNameType("pipeline_integration_User_topic"),
		datamodel.TopicNameType("pipeline_signal_Signal_topic"),
		datamodel.TopicNameType("pipeline_sourcecode_Commit_topic"),
		datamodel.TopicNameType("pipeline_work_CustomField_topic"),
		datamodel.TopicNameType("sourcecode_User_topic"),
		datamodel.TopicNameType("sourcecode_Blame_topic"),
		datamodel.TopicNameType("sourcecode_Branch_topic"),
		datamodel.TopicNameType("sourcecode_Commit_topic"),
		datamodel.TopicNameType("sourcecode_PullRequest_topic"),
		datamodel.TopicNameType("sourcecode_PullRequestComment_topic"),
		datamodel.TopicNameType("sourcecode_PullRequestReview_topic"),
		datamodel.TopicNameType("sourcecode_Repo_topic"),
		datamodel.TopicNameType("web_Hook_topic"),
		datamodel.TopicNameType("work_Changelog_topic"),
		datamodel.TopicNameType("work_CustomField_topic"),
		datamodel.TopicNameType("work_Issue_topic"),
		datamodel.TopicNameType("work_Project_topic"),
		datamodel.TopicNameType("work_Sprint_topic"),
		datamodel.TopicNameType("work_User_topic"),
	}
}

// GetModelNames returns an array of model names that are configured
func GetModelNames() []datamodel.ModelNameType {
	return []datamodel.ModelNameType{
		datamodel.ModelNameType("admin.Agent"),
		datamodel.ModelNameType("agent.EnrollResponse"),
		datamodel.ModelNameType("agent.Event"),
		datamodel.ModelNameType("agent.ExportResponse"),
		datamodel.ModelNameType("agent.IntegrationRequest"),
		datamodel.ModelNameType("agent.ProjectRequest"),
		datamodel.ModelNameType("agent.Enabled"),
		datamodel.ModelNameType("agent.EnrollRequest"),
		datamodel.ModelNameType("agent.ExportRequest"),
		datamodel.ModelNameType("agent.ExportTrigger"),
		datamodel.ModelNameType("agent.IntegrationResponse"),
		datamodel.ModelNameType("agent.ProjectResponse"),
		datamodel.ModelNameType("auth.ACLGrant"),
		datamodel.ModelNameType("codequality.Metric"),
		datamodel.ModelNameType("codequality.Project"),
		datamodel.ModelNameType("customer.User"),
		datamodel.ModelNameType("customer.CostCenter"),
		datamodel.ModelNameType("customer.Team"),
		datamodel.ModelNameType("ops.db.Change"),
		datamodel.ModelNameType("pipeline.activity.Activity"),
		datamodel.ModelNameType("pipeline.customer.User"),
		datamodel.ModelNameType("pipeline.integration.User"),
		datamodel.ModelNameType("pipeline.signal.Signal"),
		datamodel.ModelNameType("pipeline.sourcecode.Commit"),
		datamodel.ModelNameType("pipeline.work.CustomField"),
		datamodel.ModelNameType("sourcecode.User"),
		datamodel.ModelNameType("sourcecode.Blame"),
		datamodel.ModelNameType("sourcecode.Branch"),
		datamodel.ModelNameType("sourcecode.Commit"),
		datamodel.ModelNameType("sourcecode.PullRequest"),
		datamodel.ModelNameType("sourcecode.PullRequestComment"),
		datamodel.ModelNameType("sourcecode.PullRequestReview"),
		datamodel.ModelNameType("sourcecode.Repo"),
		datamodel.ModelNameType("web.Hook"),
		datamodel.ModelNameType("work.Changelog"),
		datamodel.ModelNameType("work.CustomField"),
		datamodel.ModelNameType("work.Issue"),
		datamodel.ModelNameType("work.Project"),
		datamodel.ModelNameType("work.Sprint"),
		datamodel.ModelNameType("work.User"),
	}
}
