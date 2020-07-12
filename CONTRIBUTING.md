# Contributing
When contributing to this repository, please first discuss the change you wish to make via issue, email, or any other method with the owners of this repository before making a change.

## Prerequisites

1. [Install Go][go-install].
2. Setup go `$GOPATH` and add `$GOPATH/bin` to your `$PATH`.

## Submitting a Pull Request

A typical workflow is:

1. [Fork the repository.][fork].
2. [Create a topic branch.][branch]
3. Add tests for your change.
4. Run `go test`. If your tests pass, return to the step 3.
5. Implement the change and ensure the steps from the previous step pass.
6. Run `goimports -w .`, to ensure the new code conforms to Go formatting guideline.
7. [Add, commit and push your changes.][git-help]
8. [Submit a pull request.][pull-req]

[go-install]: https://golang.org/doc/install
[go-fork-tip]: http://blog.campoy.cat/2014/03/github-and-go-forking-pull-requests-and.html
[fork]: https://help.github.com/articles/fork-a-repo
[branch]: http://learn.github.com/p/branching.html
[git-help]: https://guides.github.com
[pull-req]: https://help.github.com/articles/using-pull-requests