// Package ghactions Creates a GitHub Actions in 5s.
package ghactions

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/go-github/v93/github"
	"golang.org/x/oauth2"
)

// GitHub Action environment variables.
// https://docs.github.com/en/actions/reference/workflows-and-actions/variables
const (
	GithubAction            = "GITHUB_ACTION"
	GithubActionPath        = "GITHUB_ACTION_PATH"
	GithubActionRepository  = "GITHUB_ACTION_REPOSITORY"
	GithubActions           = "GITHUB_ACTIONS"
	GithubActor             = "GITHUB_ACTOR"
	GithubActorID           = "GITHUB_ACTOR_ID"
	GithubAPIURL            = "GITHUB_API_URL"
	GithubArtifacts         = "GITHUB_ARTIFACTS"
	GithubArtifactsList     = "GITHUB_ARTIFACTS_LIST"
	GithubBaseRef           = "GITHUB_BASE_REF"
	GithubEnv               = "GITHUB_ENV"
	GithubEventName         = "GITHUB_EVENT_NAME"
	GithubEventPath         = "GITHUB_EVENT_PATH"
	GithubGraphqlURL        = "GITHUB_GRAPHQL_URL"
	GithubHeadRef           = "GITHUB_HEAD_REF"
	GithubJob               = "GITHUB_JOB"
	GithubOutput            = "GITHUB_OUTPUT"
	GithubPath              = "GITHUB_PATH"
	GithubRef               = "GITHUB_REF"
	GithubRefName           = "GITHUB_REF_NAME"
	GithubRefProtected      = "GITHUB_REF_PROTECTED"
	GithubRefType           = "GITHUB_REF_TYPE"
	GithubRepository        = "GITHUB_REPOSITORY"
	GithubRepositoryID      = "GITHUB_REPOSITORY_ID"
	GithubRepositoryOwner   = "GITHUB_REPOSITORY_OWNER"
	GithubRepositoryOwnerID = "GITHUB_REPOSITORY_OWNER_ID"
	GithubRetentionDays     = "GITHUB_RETENTION_DAYS"
	GithubRunAttempt        = "GITHUB_RUN_ATTEMPT"
	GithubRunID             = "GITHUB_RUN_ID"
	GithubRunNumber         = "GITHUB_RUN_NUMBER"
	GithubServerURL         = "GITHUB_SERVER_URL"
	GithubSha               = "GITHUB_SHA"
	GithubStepSummary       = "GITHUB_STEP_SUMMARY"
	GithubToken             = "GITHUB_TOKEN"
	GithubTriggeringActor   = "GITHUB_TRIGGERING_ACTOR"
	GithubWorkflow          = "GITHUB_WORKFLOW"
	GithubWorkflowRef       = "GITHUB_WORKFLOW_REF"
	GithubWorkflowSha       = "GITHUB_WORKFLOW_SHA"
	GithubWorkspace         = "GITHUB_WORKSPACE"
	RunnerArch              = "RUNNER_ARCH"
	RunnerDebug             = "RUNNER_DEBUG"
	RunnerEnvironment       = "RUNNER_ENVIRONMENT"
	RunnerName              = "RUNNER_NAME"
	RunnerOS                = "RUNNER_OS"
	RunnerTemp              = "RUNNER_TEMP"
	RunnerToolCache         = "RUNNER_TOOL_CACHE"
)

// Extra environment variables.
const (
	Home     = "HOME"
	Hostname = "HOSTNAME"
	Path     = "PATH"
	PWD      = "PWD"
)

// Action GitHub Action executor.
type Action struct {
	client              *github.Client
	SkipWhenNoHandler   bool
	SkipWhenTypeUnknown bool

	onBranchProtectionConfiguration func(*github.Client, *github.BranchProtectionConfigurationEvent) error
	onBranchProtectionRule          func(*github.Client, *github.BranchProtectionRuleEvent) error
	onCheckRun                      func(*github.Client, *github.CheckRunEvent) error
	onCheckSuite                    func(*github.Client, *github.CheckSuiteEvent) error
	onCodeScanningAlert             func(*github.Client, *github.CodeScanningAlertEvent) error
	onCommitComment                 func(*github.Client, *github.CommitCommentEvent) error
	onContentReference              func(*github.Client, *github.ContentReferenceEvent) error
	onCreate                        func(*github.Client, *github.CreateEvent) error
	onCustomProperty                func(*github.Client, *github.CustomPropertyEvent) error
	onCustomPropertyValues          func(*github.Client, *github.CustomPropertyValuesEvent) error
	onDelete                        func(*github.Client, *github.DeleteEvent) error
	onDependabotAlert               func(*github.Client, *github.DependabotAlertEvent) error
	onDeployKey                     func(*github.Client, *github.DeployKeyEvent) error
	onDeployment                    func(*github.Client, *github.DeploymentEvent) error
	onDeploymentProtectionRule      func(*github.Client, *github.DeploymentProtectionRuleEvent) error
	onDeploymentReview              func(*github.Client, *github.DeploymentReviewEvent) error
	onDeploymentStatus              func(*github.Client, *github.DeploymentStatusEvent) error
	onDiscussion                    func(*github.Client, *github.DiscussionEvent) error
	onDiscussionComment             func(*github.Client, *github.DiscussionCommentEvent) error
	onFork                          func(*github.Client, *github.ForkEvent) error
	onGitHubAppAuthorization        func(*github.Client, *github.GitHubAppAuthorizationEvent) error
	onGollum                        func(*github.Client, *github.GollumEvent) error
	onInstallation                  func(*github.Client, *github.InstallationEvent) error
	onInstallationRepositories      func(*github.Client, *github.InstallationRepositoriesEvent) error
	onInstallationTarget            func(*github.Client, *github.InstallationTargetEvent) error
	onIssueComment                  func(*github.Client, *github.IssueCommentEvent) error
	onIssues                        func(*github.Client, *github.IssuesEvent) error
	onLabel                         func(*github.Client, *github.LabelEvent) error
	onMarketplacePurchase           func(*github.Client, *github.MarketplacePurchaseEvent) error
	onMember                        func(*github.Client, *github.MemberEvent) error
	onMembership                    func(*github.Client, *github.MembershipEvent) error
	onMergeGroup                    func(*github.Client, *github.MergeGroupEvent) error
	onMeta                          func(*github.Client, *github.MetaEvent) error
	onMilestone                     func(*github.Client, *github.MilestoneEvent) error
	onOrganization                  func(*github.Client, *github.OrganizationEvent) error
	onOrgBlock                      func(*github.Client, *github.OrgBlockEvent) error
	onPackage                       func(*github.Client, *github.PackageEvent) error
	onPageBuild                     func(*github.Client, *github.PageBuildEvent) error
	onPersonalAccessTokenRequest    func(*github.Client, *github.PersonalAccessTokenRequestEvent) error
	onPing                          func(*github.Client, *github.PingEvent) error
	onProject                       func(*github.Client, *github.ProjectV2Event) error
	onProjectItem                   func(*github.Client, *github.ProjectV2ItemEvent) error
	onPublic                        func(*github.Client, *github.PublicEvent) error
	onPullRequest                   func(*github.Client, *github.PullRequestEvent) error
	onPullRequestReview             func(*github.Client, *github.PullRequestReviewEvent) error
	onPullRequestReviewComment      func(*github.Client, *github.PullRequestReviewCommentEvent) error
	onPullRequestReviewThread       func(*github.Client, *github.PullRequestReviewThreadEvent) error
	onPullRequestTarget             func(*github.Client, *github.PullRequestTargetEvent) error
	onPush                          func(*github.Client, *github.PushEvent) error
	onRegistryPackage               func(*github.Client, *github.RegistryPackageEvent) error
	onRelease                       func(*github.Client, *github.ReleaseEvent) error
	onRepository                    func(*github.Client, *github.RepositoryEvent) error
	onRepositoryDispatch            func(*github.Client, *github.RepositoryDispatchEvent) error
	onRepositoryImport              func(*github.Client, *github.RepositoryImportEvent) error
	onRepositoryRuleset             func(*github.Client, *github.RepositoryRulesetEvent) error
	onRepositoryVulnerabilityAlert  func(*github.Client, *github.RepositoryVulnerabilityAlertEvent) error
	onSecretScanningAlert           func(*github.Client, *github.SecretScanningAlertEvent) error
	onSecretScanningAlertLocation   func(*github.Client, *github.SecretScanningAlertLocationEvent) error
	onSecurityAdvisory              func(*github.Client, *github.SecurityAdvisoryEvent) error
	onSecurityAndAnalysis           func(*github.Client, *github.SecurityAndAnalysisEvent) error
	onSponsorship                   func(*github.Client, *github.SponsorshipEvent) error
	onStar                          func(*github.Client, *github.StarEvent) error
	onStatus                        func(*github.Client, *github.StatusEvent) error
	onTeam                          func(*github.Client, *github.TeamEvent) error
	onTeamAdd                       func(*github.Client, *github.TeamAddEvent) error
	onUser                          func(*github.Client, *github.UserEvent) error
	onWatch                         func(*github.Client, *github.WatchEvent) error
	onWorkflowDispatch              func(*github.Client, *github.WorkflowDispatchEvent) error
	onWorkflowJob                   func(*github.Client, *github.WorkflowJobEvent) error
	onWorkflowRun                   func(*github.Client, *github.WorkflowRunEvent) error
}

// NewAction Creates a new GitHub Action executor.
func NewAction(ctx context.Context) (*Action, error) {
	client, err := newGitHubClient(ctx, os.Getenv(GithubToken))
	if err != nil {
		return nil, err
	}

	return &Action{client: client}, nil
}

// Run Executes action.
func (a *Action) Run() error {
	eventName := os.Getenv(GithubEventName)
	eventPath := os.Getenv(GithubEventPath)

	content, err := os.ReadFile(filepath.Clean(eventPath))
	if err != nil {
		return err
	}

	rawEvent, err := github.ParseWebHook(eventName, content)
	if err != nil {
		return err
	}

	switch evt := rawEvent.(type) {
	case *github.BranchProtectionConfigurationEvent:
		if a.onBranchProtectionConfiguration != nil {
			return a.onBranchProtectionConfiguration(a.client, evt)
		}

	case *github.BranchProtectionRuleEvent:
		if a.onBranchProtectionRule != nil {
			return a.onBranchProtectionRule(a.client, evt)
		}

	case *github.CheckRunEvent:
		if a.onCheckRun != nil {
			return a.onCheckRun(a.client, evt)
		}

	case *github.CheckSuiteEvent:
		if a.onCheckSuite != nil {
			return a.onCheckSuite(a.client, evt)
		}

	case *github.CodeScanningAlertEvent:
		if a.onCodeScanningAlert != nil {
			return a.onCodeScanningAlert(a.client, evt)
		}

	case *github.CommitCommentEvent:
		if a.onCommitComment != nil {
			return a.onCommitComment(a.client, evt)
		}

	case *github.ContentReferenceEvent:
		if a.onContentReference != nil {
			return a.onContentReference(a.client, evt)
		}

	case *github.CreateEvent:
		if a.onCreate != nil {
			return a.onCreate(a.client, evt)
		}

	case *github.CustomPropertyEvent:
		if a.onCustomProperty != nil {
			return a.onCustomProperty(a.client, evt)
		}

	case *github.CustomPropertyValuesEvent:
		if a.onCustomPropertyValues != nil {
			return a.onCustomPropertyValues(a.client, evt)
		}

	case *github.DeleteEvent:
		if a.onDelete != nil {
			return a.onDelete(a.client, evt)
		}

	case *github.DependabotAlertEvent:
		if a.onDependabotAlert != nil {
			return a.onDependabotAlert(a.client, evt)
		}

	case *github.DeployKeyEvent:
		if a.onDeployKey != nil {
			return a.onDeployKey(a.client, evt)
		}

	case *github.DeploymentEvent:
		if a.onDeployment != nil {
			return a.onDeployment(a.client, evt)
		}

	case *github.DeploymentProtectionRuleEvent:
		if a.onDeploymentProtectionRule != nil {
			return a.onDeploymentProtectionRule(a.client, evt)
		}

	case *github.DeploymentReviewEvent:
		if a.onDeploymentReview != nil {
			return a.onDeploymentReview(a.client, evt)
		}

	case *github.DeploymentStatusEvent:
		if a.onDeploymentStatus != nil {
			return a.onDeploymentStatus(a.client, evt)
		}

	case *github.DiscussionCommentEvent:
		if a.onDiscussionComment != nil {
			return a.onDiscussionComment(a.client, evt)
		}

	case *github.DiscussionEvent:
		if a.onDiscussion != nil {
			return a.onDiscussion(a.client, evt)
		}

	case *github.ForkEvent:
		if a.onFork != nil {
			return a.onFork(a.client, evt)
		}

	case *github.GitHubAppAuthorizationEvent:
		if a.onGitHubAppAuthorization != nil {
			return a.onGitHubAppAuthorization(a.client, evt)
		}

	case *github.GollumEvent:
		if a.onGollum != nil {
			return a.onGollum(a.client, evt)
		}

	case *github.InstallationEvent:
		if a.onInstallation != nil {
			return a.onInstallation(a.client, evt)
		}

	case *github.InstallationRepositoriesEvent:
		if a.onInstallationRepositories != nil {
			return a.onInstallationRepositories(a.client, evt)
		}

	case *github.InstallationTargetEvent:
		if a.onInstallationTarget != nil {
			return a.onInstallationTarget(a.client, evt)
		}

	case *github.IssueCommentEvent:
		if a.onIssueComment != nil {
			return a.onIssueComment(a.client, evt)
		}

	case *github.IssuesEvent:
		if a.onIssues != nil {
			return a.onIssues(a.client, evt)
		}

	case *github.LabelEvent:
		if a.onLabel != nil {
			return a.onLabel(a.client, evt)
		}

	case *github.MarketplacePurchaseEvent:
		if a.onMarketplacePurchase != nil {
			return a.onMarketplacePurchase(a.client, evt)
		}

	case *github.MemberEvent:
		if a.onMember != nil {
			return a.onMember(a.client, evt)
		}

	case *github.MembershipEvent:
		if a.onMembership != nil {
			return a.onMembership(a.client, evt)
		}

	case *github.MergeGroupEvent:
		if a.onMergeGroup != nil {
			return a.onMergeGroup(a.client, evt)
		}

	case *github.MetaEvent:
		if a.onMeta != nil {
			return a.onMeta(a.client, evt)
		}

	case *github.MilestoneEvent:
		if a.onMilestone != nil {
			return a.onMilestone(a.client, evt)
		}

	case *github.OrganizationEvent:
		if a.onOrganization != nil {
			return a.onOrganization(a.client, evt)
		}

	case *github.OrgBlockEvent:
		if a.onOrgBlock != nil {
			return a.onOrgBlock(a.client, evt)
		}

	case *github.PackageEvent:
		if a.onPackage != nil {
			return a.onPackage(a.client, evt)
		}

	case *github.PageBuildEvent:
		if a.onPageBuild != nil {
			return a.onPageBuild(a.client, evt)
		}

	case *github.PersonalAccessTokenRequestEvent:
		if a.onPersonalAccessTokenRequest != nil {
			return a.onPersonalAccessTokenRequest(a.client, evt)
		}

	case *github.PingEvent:
		if a.onPing != nil {
			return a.onPing(a.client, evt)
		}

	case *github.ProjectV2Event:
		if a.onProject != nil {
			return a.onProject(a.client, evt)
		}

	case *github.ProjectV2ItemEvent:
		if a.onProjectItem != nil {
			return a.onProjectItem(a.client, evt)
		}

	case *github.PublicEvent:
		if a.onPublic != nil {
			return a.onPublic(a.client, evt)
		}

	case *github.PullRequestEvent:
		if a.onPullRequest != nil {
			return a.onPullRequest(a.client, evt)
		}

	case *github.PullRequestReviewCommentEvent:
		if a.onPullRequestReviewComment != nil {
			return a.onPullRequestReviewComment(a.client, evt)
		}

	case *github.PullRequestReviewEvent:
		if a.onPullRequestReview != nil {
			return a.onPullRequestReview(a.client, evt)
		}

	case *github.PullRequestReviewThreadEvent:
		if a.onPullRequestReviewThread != nil {
			return a.onPullRequestReviewThread(a.client, evt)
		}

	case *github.PullRequestTargetEvent:
		if a.onPullRequestTarget != nil {
			return a.onPullRequestTarget(a.client, evt)
		}

	case *github.PushEvent:
		if a.onPush != nil {
			return a.onPush(a.client, evt)
		}

	case *github.RegistryPackageEvent:
		if a.onRegistryPackage != nil {
			return a.onRegistryPackage(a.client, evt)
		}

	case *github.ReleaseEvent:
		if a.onRelease != nil {
			return a.onRelease(a.client, evt)
		}

	case *github.RepositoryDispatchEvent:
		// noop

	case *github.RepositoryEvent:
		if a.onRepository != nil {
			return a.onRepository(a.client, evt)
		}

	case *github.RepositoryImportEvent:
		if a.onRepositoryImport != nil {
			return a.onRepositoryImport(a.client, evt)
		}

	case *github.RepositoryRulesetEvent:
		if a.onRepositoryRuleset != nil {
			return a.onRepositoryRuleset(a.client, evt)
		}

	case *github.RepositoryVulnerabilityAlertEvent:
		if a.onRepositoryVulnerabilityAlert != nil {
			return a.onRepositoryVulnerabilityAlert(a.client, evt)
		}

	case *github.SecretScanningAlertEvent:
		if a.onSecretScanningAlert != nil {
			return a.onSecretScanningAlert(a.client, evt)
		}

	case *github.SecretScanningAlertLocationEvent:
		if a.onSecretScanningAlertLocation != nil {
			return a.onSecretScanningAlertLocation(a.client, evt)
		}

	case *github.SecurityAdvisoryEvent:
		if a.onSecurityAdvisory != nil {
			return a.onSecurityAdvisory(a.client, evt)
		}

	case *github.SecurityAndAnalysisEvent:
		if a.onSecurityAndAnalysis != nil {
			return a.onSecurityAndAnalysis(a.client, evt)
		}

	case *github.SponsorshipEvent:
		if a.onSponsorship != nil {
			return a.onSponsorship(a.client, evt)
		}

	case *github.StarEvent:
		if a.onStar != nil {
			return a.onStar(a.client, evt)
		}

	case *github.StatusEvent:
		if a.onStatus != nil {
			return a.onStatus(a.client, evt)
		}

	case *github.TeamAddEvent:
		if a.onTeamAdd != nil {
			return a.onTeamAdd(a.client, evt)
		}

	case *github.TeamEvent:
		if a.onTeam != nil {
			return a.onTeam(a.client, evt)
		}

	case *github.UserEvent:
		if a.onUser != nil {
			return a.onUser(a.client, evt)
		}

	case *github.WatchEvent:
		if a.onWatch != nil {
			return a.onWatch(a.client, evt)
		}

	case *github.WorkflowDispatchEvent:
		if a.onWorkflowDispatch != nil {
			return a.onWorkflowDispatch(a.client, evt)
		}

	case *github.WorkflowJobEvent:
		if a.onWorkflowJob != nil {
			return a.onWorkflowJob(a.client, evt)
		}

	case *github.WorkflowRunEvent:
		if a.onWorkflowRun != nil {
			return a.onWorkflowRun(a.client, evt)
		}

	default:
		if a.SkipWhenTypeUnknown {
			return nil
		}

		return fmt.Errorf("unsupported event type: %q", eventName)
	}

	if a.SkipWhenNoHandler {
		return nil
	}

	return fmt.Errorf("no handler for the received event type %q", eventName)
}

// OnCheckRun CheckRun handler.
func (a *Action) OnCheckRun(eventHandler func(*github.Client, *github.CheckRunEvent) error) *Action {
	a.onCheckRun = eventHandler
	return a
}

// OnCheckSuite CheckSuite handler.
func (a *Action) OnCheckSuite(eventHandler func(*github.Client, *github.CheckSuiteEvent) error) *Action {
	a.onCheckSuite = eventHandler
	return a
}

// OnCommitComment CommitComment handler.
func (a *Action) OnCommitComment(eventHandler func(*github.Client, *github.CommitCommentEvent) error) *Action {
	a.onCommitComment = eventHandler
	return a
}

// OnCreate Create handler.
func (a *Action) OnCreate(eventHandler func(*github.Client, *github.CreateEvent) error) *Action {
	a.onCreate = eventHandler
	return a
}

// OnDelete Delete handler.
func (a *Action) OnDelete(eventHandler func(*github.Client, *github.DeleteEvent) error) *Action {
	a.onDelete = eventHandler
	return a
}

// OnDeployment Deployment handler.
func (a *Action) OnDeployment(eventHandler func(*github.Client, *github.DeploymentEvent) error) *Action {
	a.onDeployment = eventHandler
	return a
}

// OnDeploymentStatus DeploymentStatus handler.
func (a *Action) OnDeploymentStatus(eventHandler func(*github.Client, *github.DeploymentStatusEvent) error) *Action {
	a.onDeploymentStatus = eventHandler
	return a
}

// OnFork Fork handler.
func (a *Action) OnFork(eventHandler func(*github.Client, *github.ForkEvent) error) *Action {
	a.onFork = eventHandler
	return a
}

// OnGollum Gollum handler.
func (a *Action) OnGollum(eventHandler func(*github.Client, *github.GollumEvent) error) *Action {
	a.onGollum = eventHandler
	return a
}

// OnIssueComment IssueComment handler.
func (a *Action) OnIssueComment(eventHandler func(*github.Client, *github.IssueCommentEvent) error) *Action {
	a.onIssueComment = eventHandler
	return a
}

// OnIssues Issues handler.
func (a *Action) OnIssues(eventHandler func(*github.Client, *github.IssuesEvent) error) *Action {
	a.onIssues = eventHandler
	return a
}

// OnLabel Label handler.
func (a *Action) OnLabel(eventHandler func(*github.Client, *github.LabelEvent) error) *Action {
	a.onLabel = eventHandler
	return a
}

// OnMember Member handler.
func (a *Action) OnMember(eventHandler func(*github.Client, *github.MemberEvent) error) *Action {
	a.onMember = eventHandler
	return a
}

// OnMilestone Milestone handler.
func (a *Action) OnMilestone(eventHandler func(*github.Client, *github.MilestoneEvent) error) *Action {
	a.onMilestone = eventHandler
	return a
}

// OnPageBuild PageBuild handler.
func (a *Action) OnPageBuild(eventHandler func(*github.Client, *github.PageBuildEvent) error) *Action {
	a.onPageBuild = eventHandler
	return a
}

// OnProjectItem ProjectItem handler.
func (a *Action) OnProjectItem(eventHandler func(*github.Client, *github.ProjectV2ItemEvent) error) *Action {
	a.onProjectItem = eventHandler
	return a
}

// OnProject Project handler.
func (a *Action) OnProject(eventHandler func(*github.Client, *github.ProjectV2Event) error) *Action {
	a.onProject = eventHandler
	return a
}

// OnPublic Public handler.
func (a *Action) OnPublic(eventHandler func(*github.Client, *github.PublicEvent) error) *Action {
	a.onPublic = eventHandler
	return a
}

// OnPullRequest PullRequest handler.
func (a *Action) OnPullRequest(eventHandler func(*github.Client, *github.PullRequestEvent) error) *Action {
	a.onPullRequest = eventHandler
	return a
}

// OnPullRequestTarget PullRequestTarget handler.
func (a *Action) OnPullRequestTarget(eventHandler func(*github.Client, *github.PullRequestTargetEvent) error) *Action {
	a.onPullRequestTarget = eventHandler
	return a
}

// OnPullRequestReview PullRequestReview handler.
func (a *Action) OnPullRequestReview(eventHandler func(*github.Client, *github.PullRequestReviewEvent) error) *Action {
	a.onPullRequestReview = eventHandler
	return a
}

// OnPullRequestReviewComment PullRequestReviewComment handler.
func (a *Action) OnPullRequestReviewComment(eventHandler func(*github.Client, *github.PullRequestReviewCommentEvent) error) *Action {
	a.onPullRequestReviewComment = eventHandler
	return a
}

// OnPush Push handler.
func (a *Action) OnPush(eventHandler func(*github.Client, *github.PushEvent) error) *Action {
	a.onPush = eventHandler
	return a
}

// OnRelease Release handler.
func (a *Action) OnRelease(eventHandler func(*github.Client, *github.ReleaseEvent) error) *Action {
	a.onRelease = eventHandler
	return a
}

// OnRepositoryVulnerabilityAlert RepositoryVulnerabilityAlert handler.
func (a *Action) OnRepositoryVulnerabilityAlert(eventHandler func(*github.Client, *github.RepositoryVulnerabilityAlertEvent) error) *Action {
	a.onRepositoryVulnerabilityAlert = eventHandler
	return a
}

// OnStatus Status handler.
func (a *Action) OnStatus(eventHandler func(*github.Client, *github.StatusEvent) error) *Action {
	a.onStatus = eventHandler
	return a
}

// OnWatch Watch handler.
func (a *Action) OnWatch(eventHandler func(*github.Client, *github.WatchEvent) error) *Action {
	a.onWatch = eventHandler
	return a
}

// OnBranchProtectionRule BranchProtectionRule handler.
func (a *Action) OnBranchProtectionRule(eventHandler func(*github.Client, *github.BranchProtectionRuleEvent) error) *Action {
	a.onBranchProtectionRule = eventHandler
	return a
}

// OnBranchProtectionConfiguration BranchProtectionConfiguration handler.
func (a *Action) OnBranchProtectionConfiguration(eventHandler func(*github.Client, *github.BranchProtectionConfigurationEvent) error) *Action {
	a.onBranchProtectionConfiguration = eventHandler
	return a
}

// OnContentReference ContentReference handler.
func (a *Action) OnContentReference(eventHandler func(*github.Client, *github.ContentReferenceEvent) error) *Action {
	a.onContentReference = eventHandler
	return a
}

// OnCustomProperty CustomProperty handler.
func (a *Action) OnCustomProperty(eventHandler func(*github.Client, *github.CustomPropertyEvent) error) *Action {
	a.onCustomProperty = eventHandler
	return a
}

// OnCustomPropertyValues CustomPropertyValues handler.
func (a *Action) OnCustomPropertyValues(eventHandler func(*github.Client, *github.CustomPropertyValuesEvent) error) *Action {
	a.onCustomPropertyValues = eventHandler
	return a
}

// OnDependabotAlert DependabotAlert handler.
func (a *Action) OnDependabotAlert(eventHandler func(*github.Client, *github.DependabotAlertEvent) error) *Action {
	a.onDependabotAlert = eventHandler
	return a
}

// OnDeployKey DeployKey handler.
func (a *Action) OnDeployKey(eventHandler func(*github.Client, *github.DeployKeyEvent) error) *Action {
	a.onDeployKey = eventHandler
	return a
}

// OnDeploymentProtectionRule DeploymentProtectionRule handler.
func (a *Action) OnDeploymentProtectionRule(eventHandler func(*github.Client, *github.DeploymentProtectionRuleEvent) error) *Action {
	a.onDeploymentProtectionRule = eventHandler
	return a
}

// OnDeploymentReview DeploymentReview handler.
func (a *Action) OnDeploymentReview(eventHandler func(*github.Client, *github.DeploymentReviewEvent) error) *Action {
	a.onDeploymentReview = eventHandler
	return a
}

// OnDiscussionComment DiscussionComment handler.
func (a *Action) OnDiscussionComment(eventHandler func(*github.Client, *github.DiscussionCommentEvent) error) *Action {
	a.onDiscussionComment = eventHandler
	return a
}

// OnDiscussion Discussion handler.
func (a *Action) OnDiscussion(eventHandler func(*github.Client, *github.DiscussionEvent) error) *Action {
	a.onDiscussion = eventHandler
	return a
}

// OnGitHubAppAuthorization GitHubAppAuthorization handler.
func (a *Action) OnGitHubAppAuthorization(eventHandler func(*github.Client, *github.GitHubAppAuthorizationEvent) error) *Action {
	a.onGitHubAppAuthorization = eventHandler
	return a
}

// OnInstallation Installation handler.
func (a *Action) OnInstallation(eventHandler func(*github.Client, *github.InstallationEvent) error) *Action {
	a.onInstallation = eventHandler
	return a
}

// OnInstallationRepositories InstallationRepositories handler.
func (a *Action) OnInstallationRepositories(eventHandler func(*github.Client, *github.InstallationRepositoriesEvent) error) *Action {
	a.onInstallationRepositories = eventHandler
	return a
}

// OnInstallationTarget InstallationTarget handler.
func (a *Action) OnInstallationTarget(eventHandler func(*github.Client, *github.InstallationTargetEvent) error) *Action {
	a.onInstallationTarget = eventHandler
	return a
}

// OnMarketplacePurchase MarketplacePurchase handler.
func (a *Action) OnMarketplacePurchase(eventHandler func(*github.Client, *github.MarketplacePurchaseEvent) error) *Action {
	a.onMarketplacePurchase = eventHandler
	return a
}

// OnMembership Membership handler.
func (a *Action) OnMembership(eventHandler func(*github.Client, *github.MembershipEvent) error) *Action {
	a.onMembership = eventHandler
	return a
}

// OnMergeGroup MergeGroup handler.
func (a *Action) OnMergeGroup(eventHandler func(*github.Client, *github.MergeGroupEvent) error) *Action {
	a.onMergeGroup = eventHandler
	return a
}

// OnMeta Meta handler.
func (a *Action) OnMeta(eventHandler func(*github.Client, *github.MetaEvent) error) *Action {
	a.onMeta = eventHandler
	return a
}

// OnOrganization Organization handler.
func (a *Action) OnOrganization(eventHandler func(*github.Client, *github.OrganizationEvent) error) *Action {
	a.onOrganization = eventHandler
	return a
}

// OnOrgBlock OrgBlock handler.
func (a *Action) OnOrgBlock(eventHandler func(*github.Client, *github.OrgBlockEvent) error) *Action {
	a.onOrgBlock = eventHandler
	return a
}

// OnPackage Package handler.
func (a *Action) OnPackage(eventHandler func(*github.Client, *github.PackageEvent) error) *Action {
	a.onPackage = eventHandler
	return a
}

// OnPersonalAccessTokenRequest PersonalAccessTokenRequest handler.
func (a *Action) OnPersonalAccessTokenRequest(eventHandler func(*github.Client, *github.PersonalAccessTokenRequestEvent) error) *Action {
	a.onPersonalAccessTokenRequest = eventHandler
	return a
}

// OnPing Ping handler.
func (a *Action) OnPing(eventHandler func(*github.Client, *github.PingEvent) error) *Action {
	a.onPing = eventHandler
	return a
}

// OnRepository Repository handler.
func (a *Action) OnRepository(eventHandler func(*github.Client, *github.RepositoryEvent) error) *Action {
	a.onRepository = eventHandler
	return a
}

// OnRepositoryDispatch RepositoryDispatch handler.
func (a *Action) OnRepositoryDispatch(eventHandler func(*github.Client, *github.RepositoryDispatchEvent) error) *Action {
	a.onRepositoryDispatch = eventHandler
	return a
}

// OnRepositoryImport RepositoryImport handler.
func (a *Action) OnRepositoryImport(eventHandler func(*github.Client, *github.RepositoryImportEvent) error) *Action {
	a.onRepositoryImport = eventHandler
	return a
}

// OnRepositoryRuleset RepositoryRuleset handler.
func (a *Action) OnRepositoryRuleset(eventHandler func(*github.Client, *github.RepositoryRulesetEvent) error) *Action {
	a.onRepositoryRuleset = eventHandler
	return a
}

// OnSecretScanningAlert SecretScanningAlert handler.
func (a *Action) OnSecretScanningAlert(eventHandler func(*github.Client, *github.SecretScanningAlertEvent) error) *Action {
	a.onSecretScanningAlert = eventHandler
	return a
}

// OnSecretScanningAlertLocation SecretScanningAlertLocation handler.
func (a *Action) OnSecretScanningAlertLocation(eventHandler func(*github.Client, *github.SecretScanningAlertLocationEvent) error) *Action {
	a.onSecretScanningAlertLocation = eventHandler
	return a
}

// OnSecurityAndAnalysis SecurityAndAnalysis handler.
func (a *Action) OnSecurityAndAnalysis(eventHandler func(*github.Client, *github.SecurityAndAnalysisEvent) error) *Action {
	a.onSecurityAndAnalysis = eventHandler
	return a
}

// OnStar Star handler.
func (a *Action) OnStar(eventHandler func(*github.Client, *github.StarEvent) error) *Action {
	a.onStar = eventHandler
	return a
}

// OnTeam Team handler.
func (a *Action) OnTeam(eventHandler func(*github.Client, *github.TeamEvent) error) *Action {
	a.onTeam = eventHandler
	return a
}

// OnTeamAdd Team Add handler.
func (a *Action) OnTeamAdd(eventHandler func(*github.Client, *github.TeamAddEvent) error) *Action {
	a.onTeamAdd = eventHandler
	return a
}

// OnUser User handler.
func (a *Action) OnUser(eventHandler func(*github.Client, *github.UserEvent) error) *Action {
	a.onUser = eventHandler
	return a
}

// OnWorkflowDispatch Workflow Dispatch handler.
func (a *Action) OnWorkflowDispatch(eventHandler func(*github.Client, *github.WorkflowDispatchEvent) error) *Action {
	a.onWorkflowDispatch = eventHandler
	return a
}

// OnWorkflowJob Workflow Job handler.
func (a *Action) OnWorkflowJob(eventHandler func(*github.Client, *github.WorkflowJobEvent) error) *Action {
	a.onWorkflowJob = eventHandler
	return a
}

// OnWorkflowRun Workflow Run handler.
func (a *Action) OnWorkflowRun(eventHandler func(*github.Client, *github.WorkflowRunEvent) error) *Action {
	a.onWorkflowRun = eventHandler
	return a
}

// OnSecurityAdvisory Security Advisory handler.
func (a *Action) OnSecurityAdvisory(eventHandler func(*github.Client, *github.SecurityAdvisoryEvent) error) *Action {
	a.onSecurityAdvisory = eventHandler
	return a
}

// OnCodeScanningAlert CodeScanning Alert handler.
func (a *Action) OnCodeScanningAlert(eventHandler func(*github.Client, *github.CodeScanningAlertEvent) error) *Action {
	a.onCodeScanningAlert = eventHandler
	return a
}

// OnSponsorship Sponsorship handler.
func (a *Action) OnSponsorship(eventHandler func(*github.Client, *github.SponsorshipEvent) error) *Action {
	a.onSponsorship = eventHandler
	return a
}

func newGitHubClient(ctx context.Context, token string) (*github.Client, error) {
	if token == "" {
		return github.NewClient()
	}

	return github.NewClient(github.WithHTTPClient(oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	))))
}

// GetRepoInfo Split "GITHUB_REPOSITORY" to [owner, repoName].
func GetRepoInfo() (owner, repoName string) {
	githubRepository := os.Getenv(GithubRepository)

	parts := strings.SplitN(githubRepository, "/", 2)

	return parts[0], parts[1]
}
