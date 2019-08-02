package ghactions

import (
	"context"
	"os"
	"testing"

	"github.com/google/go-github/v27/github"
	"github.com/ldez/ghactions/event"
)

func TestAction(t *testing.T) {
	err := os.Setenv(GithubEventName, event.Issues)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Setenv(GithubEventPath, "./fixtures/issues.json")
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	action := NewAction(ctx)
	action.SkipWhenNoHandler = false
	action.SkipWhenTypeUnknown = false

	err = action.
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
