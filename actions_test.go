package ghactions

import (
	"context"
	"testing"

	"github.com/google/go-github/v43/github"
	"github.com/ldez/ghactions/event"
)

func TestAction(t *testing.T) {
	t.Setenv(GithubEventName, event.Issues)
	t.Setenv(GithubEventPath, "./fixtures/issues.json")

	ctx := context.Background()
	action := NewAction(ctx)
	action.SkipWhenNoHandler = false
	action.SkipWhenTypeUnknown = false

	err := action.
		OnPullRequest(func(client *github.Client, requestEvent *github.PullRequestEvent) error {
			return nil
		}).
		OnIssues(func(client *github.Client, issuesEvent *github.IssuesEvent) error {
			return nil
		}).
		Run()
	if err != nil {
		t.Fatal(err)
	}
}
