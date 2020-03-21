# GHActions

[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/ldez/ghactions.svg?label=release)](https://github.com/ldez/ghactions/releases)
[![Build Status](https://travis-ci.com/ldez/ghactions.svg?branch=master)](https://travis-ci.com/ldez/ghactions)
[![godoc](https://godoc.org/github.com/ldez/ghactions?status.svg)](https://godoc.org/github.com/ldez/ghactions)

[![Say Thanks!](https://img.shields.io/badge/Say%20Thanks-!-1EAEDB.svg)](https://saythanks.io/to/ldez)

Create a Github Action in 5 seconds!

- Environment variables: https://godoc.org/github.com/ldez/ghactions#pkg-constants
- Supported events: https://godoc.org/github.com/ldez/ghactions/event#pkg-constants

## Examples

```go
package main

import (
	"context"
	"log"

	"github.com/google/go-github/v30/github"
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
