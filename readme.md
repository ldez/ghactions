# GHActions

[![release](https://img.shields.io/github/release/ldez/ghactions.svg?style=flat)](https://github.com/ldez/ghactions/releases)
[![Build Status](https://travis-ci.org/ldez/ghactions.svg?branch=master)](https://travis-ci.org/ldez/ghactions)
[![godoc](https://godoc.org/github.com/ldez/ghactions?status.svg)](https://godoc.org/github.com/ldez/ghactions)

Create a Github Action in 5 seconds!

- Environment variables: https://godoc.org/github.com/ldez/ghactions#pkg-constants
- Supported events: https://godoc.org/github.com/ldez/ghactions/event#pkg-constants

## Examples

```go
package main

import (
	"context"
	"log"

	"github.com/google/go-github/v24/github"
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

- https://developer.github.com/actions/creating-github-actions/accessing-the-runtime-environment/
- https://developer.github.com/actions/creating-github-actions/creating-a-docker-container/
- https://github.com/marketplace/actions
- https://github-actions.explore-tech.org