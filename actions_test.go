package ghactions

import (
	"context"
	"testing"

	"github.com/google/go-github/v68/github"
)

func TestAction(t *testing.T) {
	t.Setenv(GithubEventName, "issues")
	t.Setenv(GithubEventPath, "./fixtures/issues.json")

	ctx := context.Background()
	action := NewAction(ctx)
	action.SkipWhenNoHandler = false
	action.SkipWhenTypeUnknown = false

	err := action.
		OnPullRequest(func(_ *github.Client, _ *github.PullRequestEvent) error {
			return nil
		}).
		OnIssues(func(_ *github.Client, _ *github.IssuesEvent) error {
			return nil
		}).
		Run()
	if err != nil {
		t.Fatal(err)
	}
}
