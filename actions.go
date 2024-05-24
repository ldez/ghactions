// Package ghactions Creates a GitHub Actions in 5s.
package ghactions

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/go-github/v62/github"
	"golang.org/x/oauth2"
)

// GitHub Action environment variables.
const (
	Home             = "HOME"
	Hostname         = "HOSTNAME"
	PWD              = "PWD"
	Path             = "PATH"
	GithubAction     = "GITHUB_ACTION"
	GithubActions    = "GITHUB_ACTIONS"
	GithubActor      = "GITHUB_ACTOR"
	GithubToken      = "GITHUB_TOKEN"
	GithubWorkflow   = "GITHUB_WORKFLOW"
	GithubRunID      = "GITHUB_RUN_ID"
	GithubRunNumber  = "GITHUB_RUN_NUMBER"
	GithubRepository = "GITHUB_REPOSITORY"
	GithubEventName  = "GITHUB_EVENT_NAME"
	GithubEventPath  = "GITHUB_EVENT_PATH"
	GithubWorkspace  = "GITHUB_WORKSPACE"
	GithubSha        = "GITHUB_SHA"
	GithubRef        = "GITHUB_REF"
	GithubHeadRef    = "GITHUB_HEAD_REF"
	GithubBaseRef    = "GITHUB_BASE_REF"
)

// Action GitHub Action executor.
type Action struct {
	client                         *github.Client
	SkipWhenNoHandler              bool
	SkipWhenTypeUnknown            bool
	onCheckRun                     func(*github.Client, *github.CheckRunEvent) error
	onCheckSuite                   func(*github.Client, *github.CheckSuiteEvent) error
	onCommitComment                func(*github.Client, *github.CommitCommentEvent) error
	onCreate                       func(*github.Client, *github.CreateEvent) error
	onDelete                       func(*github.Client, *github.DeleteEvent) error
	onDeployment                   func(*github.Client, *github.DeploymentEvent) error
	onDeploymentStatus             func(*github.Client, *github.DeploymentStatusEvent) error
	onFork                         func(*github.Client, *github.ForkEvent) error
	onGollum                       func(*github.Client, *github.GollumEvent) error
	onIssueComment                 func(*github.Client, *github.IssueCommentEvent) error
	onIssues                       func(*github.Client, *github.IssuesEvent) error
	onLabel                        func(*github.Client, *github.LabelEvent) error
	onMember                       func(*github.Client, *github.MemberEvent) error
	onMilestone                    func(*github.Client, *github.MilestoneEvent) error
	onPageBuild                    func(*github.Client, *github.PageBuildEvent) error
	onProjectCard                  func(*github.Client, *github.ProjectCardEvent) error
	onProjectColumn                func(*github.Client, *github.ProjectColumnEvent) error
	onProject                      func(*github.Client, *github.ProjectEvent) error
	onPublic                       func(*github.Client, *github.PublicEvent) error
	onPullRequest                  func(*github.Client, *github.PullRequestEvent) error
	onPullRequestTarget            func(*github.Client, *github.PullRequestTargetEvent) error
	onPullRequestReview            func(*github.Client, *github.PullRequestReviewEvent) error
	onPullRequestReviewComment     func(*github.Client, *github.PullRequestReviewCommentEvent) error
	onPush                         func(*github.Client, *github.PushEvent) error
	onRelease                      func(*github.Client, *github.ReleaseEvent) error
	onRepositoryVulnerabilityAlert func(*github.Client, *github.RepositoryVulnerabilityAlertEvent) error
	onStatus                       func(*github.Client, *github.StatusEvent) error
	onWatch                        func(*github.Client, *github.WatchEvent) error
	// onRepositoryDispatch func(*github.Client, interface{}) error
}

// NewAction Creates a new GitHub Action executor.
func NewAction(ctx context.Context) *Action {
	return &Action{
		client: newGitHubClient(ctx, os.Getenv(GithubToken)),
	}
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
	case *github.CheckRunEvent:
		if a.onCheckRun != nil {
			return a.onCheckRun(a.client, evt)
		}

	case *github.CheckSuiteEvent:
		if a.onCheckSuite != nil {
			return a.onCheckSuite(a.client, evt)
		}

	case *github.CommitCommentEvent:
		if a.onCommitComment != nil {
			return a.onCommitComment(a.client, evt)
		}

	case *github.CreateEvent:
		if a.onCreate != nil {
			return a.onCreate(a.client, evt)
		}

	case *github.DeleteEvent:
		if a.onDelete != nil {
			return a.onDelete(a.client, evt)
		}

	case *github.DeploymentEvent:
		if a.onDeployment != nil {
			return a.onDeployment(a.client, evt)
		}

	case *github.DeploymentStatusEvent:
		if a.onDeploymentStatus != nil {
			return a.onDeploymentStatus(a.client, evt)
		}

	case *github.ForkEvent:
		if a.onFork != nil {
			return a.onFork(a.client, evt)
		}

	case *github.GollumEvent:
		if a.onGollum != nil {
			return a.onGollum(a.client, evt)
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

	case *github.MemberEvent:
		if a.onMember != nil {
			return a.onMember(a.client, evt)
		}

	case *github.MilestoneEvent:
		if a.onMilestone != nil {
			return a.onMilestone(a.client, evt)
		}

	case *github.PageBuildEvent:
		if a.onPageBuild != nil {
			return a.onPageBuild(a.client, evt)
		}

	case *github.ProjectCardEvent:
		if a.onProjectCard != nil {
			return a.onProjectCard(a.client, evt)
		}

	case *github.ProjectColumnEvent:
		if a.onProjectColumn != nil {
			return a.onProjectColumn(a.client, evt)
		}

	case *github.ProjectEvent:
		if a.onProject != nil {
			return a.onProject(a.client, evt)
		}

	case *github.PublicEvent:
		if a.onPublic != nil {
			return a.onPublic(a.client, evt)
		}

	case *github.PullRequestEvent:
		if a.onPullRequest != nil {
			return a.onPullRequest(a.client, evt)
		}

	case *github.PullRequestTargetEvent:
		if a.onPullRequestTarget != nil {
			return a.onPullRequestTarget(a.client, evt)
		}

	case *github.PullRequestReviewEvent:
		if a.onPullRequestReview != nil {
			return a.onPullRequestReview(a.client, evt)
		}

	case *github.PullRequestReviewCommentEvent:
		if a.onPullRequestReviewComment != nil {
			return a.onPullRequestReviewComment(a.client, evt)
		}

	case *github.PushEvent:
		if a.onPush != nil {
			return a.onPush(a.client, evt)
		}

	case *github.ReleaseEvent:
		if a.onRelease != nil {
			return a.onRelease(a.client, evt)
		}

	case *github.RepositoryVulnerabilityAlertEvent:
		if a.onRepositoryVulnerabilityAlert != nil {
			return a.onRepositoryVulnerabilityAlert(a.client, evt)
		}

	case *github.RepositoryDispatchEvent:
		// noop

	case *github.StatusEvent:
		if a.onStatus != nil {
			return a.onStatus(a.client, evt)
		}

	case *github.WatchEvent:
		if a.onWatch != nil {
			return a.onWatch(a.client, evt)
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

// OnProjectCard ProjectCard handler.
func (a *Action) OnProjectCard(eventHandler func(*github.Client, *github.ProjectCardEvent) error) *Action {
	a.onProjectCard = eventHandler
	return a
}

// OnProjectColumn ProjectColumn handler.
func (a *Action) OnProjectColumn(eventHandler func(*github.Client, *github.ProjectColumnEvent) error) *Action {
	a.onProjectColumn = eventHandler
	return a
}

// OnProject Project handler.
func (a *Action) OnProject(eventHandler func(*github.Client, *github.ProjectEvent) error) *Action {
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

func newGitHubClient(ctx context.Context, token string) *github.Client {
	var client *github.Client
	if token == "" {
		client = github.NewClient(nil)
	} else {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)
		client = github.NewClient(tc)
	}
	return client
}

// GetRepoInfo Split "GITHUB_REPOSITORY" to [owner, repoName].
func GetRepoInfo() (owner, repoName string) {
	githubRepository := os.Getenv(GithubRepository)

	parts := strings.SplitN(githubRepository, "/", 2)

	return parts[0], parts[1]
}
