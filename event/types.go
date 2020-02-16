// Package event Event types.
package event

// Event types.
const (
	CheckRun                     = "check_run"
	CheckSuite                   = "check_suite"
	CommitComment                = "commit_comment"
	Create                       = "create"
	Delete                       = "delete"
	Deployment                   = "deployment"
	DeploymentStatus             = "deployment_status"
	Fork                         = "fork"
	Gollum                       = "gollum"
	IssueComment                 = "issue_comment"
	Issues                       = "issues"
	Label                        = "label"
	Member                       = "member" // Still available ?
	Milestone                    = "milestone"
	PageBuild                    = "page_build"
	Project                      = "project"
	ProjectCard                  = "project_card"
	ProjectColumn                = "project_column"
	Public                       = "public"
	PullRequest                  = "pull_request"
	PullRequestReviewComment     = "pull_request_review_comment"
	PullRequestReview            = "pull_request_review"
	Push                         = "push"
	RegistryPackage              = "registry_package"
	Release                      = "release"
	RepositoryDispatch           = "repository_dispatch"
	RepositoryVulnerabilityAlert = "repository_vulnerability_alert" // Still available ?
	Status                       = "status"
	Schedule                     = "schedule"
	Watch                        = "watch"
)
