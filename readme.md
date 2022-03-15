# GHActions

[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/ldez/ghactions.svg?label=release)](https://github.com/ldez/ghactions/releases)
[![Build Status](https://github.com/ldez/ghactions/workflows/Main/badge.svg?branch=master)](https://github.com/ldez/ghactions/actions)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/ldez/ghactions)](https://pkg.go.dev/github.com/ldez/ghactions)

[![Sponsor](https://img.shields.io/badge/Sponsor%20me-%E2%9D%A4%EF%B8%8F-pink)](https://github.com/sponsors/ldez)

Create a Github Action in 5 seconds!

- Environment variables: https://pkg.go.dev/github.com/ldez/ghactions#pkg-constants
- Supported events: https://pkg.go.dev/github.com/ldez/ghactions/event#pkg-constants

## Examples

```go
package main

import (
	"context"
	"log"

	"github.com/google/go-github/v43/github"
	"github.com/ldez/ghactions"
	"github.com/ldez/ghactions/event"
)

func main() {
	ctx := context.Background()
	action := ghactions.NewAction(ctx)
	// action.SkipWhenNoHandler = true
	// action.SkipWhenTypeUnknown = true

	err := action.
		OnPullRequest(func(client *github.Client, requestEvent *github.PullRequestEvent) error {
			// TODO add your code.
			return nil
		}).
		OnIssues(func(client *github.Client, issuesEvent *github.IssuesEvent) error {
			// TODO add your code.
			return nil
		}).
		Run()

	if err != nil {
		log.Fatal(err)
	}
}
```

## References

- https://help.github.com/en/actions
- https://github.com/marketplace/actions
- https://github-actions.explore-tech.org
- https://help.github.com/en/actions/configuring-and-managing-workflows/using-environment-variables
- https://help.github.com/en/actions/reference/events-that-trigger-workflows
