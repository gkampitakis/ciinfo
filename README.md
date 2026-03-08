# CI Info

## Acknowledgement

This repository is based on work done in [watson/ci-info](https://github.com/watson/ci-info)
and the contributors.

I will do my best to keep this library up to date and in sync with changes in
[watson/ci-info](https://github.com/watson/ci-info).

---

Get details about the current Continuous Integration environment.

[![Tests](https://github.com/gkampitakis/ciinfo/actions/workflows/tests.yml/badge.svg)](https://github.com/gkampitakis/ciinfo/actions/workflows/tests.yml)

## Installation

```bash
go get github.com/gkampitakis/ciinfo
```

## Usage

```go
import (
  "fmt"

  "github.com/gkampitakis/ciinfo"
)

if ciinfo.IsCI {
  fmt.Printf("The name of the CI server is: %s", ciinfo.Name)
} else {
  fmt.Printf("This program is not running on a CI server")
}
```

## CLI Support

`ciinfo` can also be used as a CLI. You can install it with

```sh
go install github.com/gkampitakis/ciinfo/ciinfo@latest
```

Then `ciinfo` command will be successful ( code 0 ) if running on CI else error ( code -1 ).

```sh
#  will output isCI if running onCI
ciinfo && echo 'isCI'
```

`ciinfo` also has

```shell
Usage of ciinfo:
  -output string
    	you can output info [json, pretty].
  -pr
    	check if shell is running on CI for a Pull Request.
```

## Supported CI tools

Officially supported CI servers:

| Name                                                                            | Constant                    | isPR |
| ------------------------------------------------------------------------------- | --------------------------- | ---- |
| [Agola CI](https://agola.io/)                                                   | `ci.AGOLA`                  | ✅   |
| [Alpic](https://alpic.ai/)                                                      | `ci.ALPIC`                  | 🚫   |
| [AWS CodeBuild](https://aws.amazon.com/codebuild/)                              | `ciinfo.CODEBUILD`          | ✅   |
| [AppVeyor](http://www.appveyor.com)                                             | `ciinfo.APPVEYOR`           | ✅   |
| [Azure Pipelines](https://azure.microsoft.com/en-us/services/devops/pipelines/) | `ciinfo.AZURE_PIPELINES`    | ✅   |
| [Appcircle](https://appcircle.io/)                                              | `ciinfo.APPCIRCLE`          | ✅   |
| [Bamboo](https://www.atlassian.com/software/bamboo) by Atlassian                | `ciinfo.BAMBOO`             | 🚫   |
| [Bitbucket Pipelines](https://bitbucket.org/product/features/pipelines)         | `ciinfo.BITBUCKET`          | ✅   |
| [Bitrise](https://www.bitrise.io/)                                              | `ciinfo.BITRISE`            | ✅   |
| [Buddy](https://buddy.works/)                                                   | `ciinfo.BUDDY`              | ✅   |
| [Buildkite](https://buildkite.com)                                              | `ciinfo.BUILDKITE`          | ✅   |
| [CircleCI](http://circleci.com)                                                 | `ciinfo.CIRCLE`             | ✅   |
| [Cirrus CI](https://cirrus-ci.org)                                              | `ciinfo.CIRRUS`             | ✅   |
| [Cloudflare Pages](https://pages.cloudflare.com/)                               | `ci.CLOUDFLARE_PAGES`       | 🚫   |
| [Cloudflare Workers](https://pages.cloudflare.com/)                             | `ci.CLOUDFLARE_WORKERS`     | 🚫   |
| [Codefresh](https://codefresh.io/)                                              | `ciinfo.CODEFRESH`          | ✅   |
| [Codeship](https://codeship.com)                                                | `ciinfo.CODESHIP`           | 🚫   |
| [Drone](https://drone.io)                                                       | `ciinfo.DRONE`              | ✅   |
| [dsari](https://github.com/rfinnie/dsari)                                       | `ciinfo.DSARI`              | 🚫   |
| [Earthly CI](https://earthly.dev/)                                              | `ci.EARTHLY`                | 🚫   |
| [Expo Application Services](https://expo.dev/eas)                               | `ciinfo.EAS`                | 🚫   |
| [Gerrit CI](https://www.gerritcodereview.com)                                   | `ciinfo.GERRIT`             | 🚫   |
| [GitHub Actions](https://github.com/features/actions/)                          | `ciinfo.GITHUB_ACTIONS`     | ✅   |
| [GitLab CI](https://about.gitlab.com/gitlab-ci/)                                | `ciinfo.GITLAB`             | ✅   |
| [Gitea Actions](https://about.gitea.com/)                                       | `ci.GITEA_ACTIONS`          | 🚫   |
| [GoCD](https://www.go.cd/)                                                      | `ciinfo.GOCD`               | 🚫   |
| [Google Cloud Build](https://cloud.google.com/build)                            | `ciinfo.GOOGLE_CLOUD_BUILD` | 🚫   |
| [Harness CI](https://www.harness.io/products/continuous-integration)            | `ciinfo.HARNESS`            | 🚫   |
| [Heroku](https://www.heroku.com)                                                | `ciinfo.HEROKU`             | 🚫   |
| [Hudson](http://hudson-ci.org)                                                  | `ciinfo.HUDSON`             | 🚫   |
| [Jenkins CI](https://jenkins-ci.org)                                            | `ciinfo.JENKINS`            | ✅   |
| [LayerCI](https://layerci.com/)                                                 | `ciinfo.LAYERCI`            | ✅   |
| [Magnum CI](https://magnum-ci.com)                                              | `ciinfo.MAGNUM`             | 🚫   |
| [Netlify CI](https://www.netlify.com/)                                          | `ciinfo.NETLIFY`            | ✅   |
| [Nevercode](http://nevercode.io/)                                               | `ciinfo.NEVERCODE`          | ✅   |
| [Prow](https://docs.prow.k8s.io/)                                               | `ci.PROW`                   | 🚫   |
| [ReleaseHub](https://releasehub.com/)                                           | `ciinfo.RELEASEHUB`         | ✅   |
| [Render](https://render.com/)                                                   | `ciinfo.RENDER`             | ✅   |
| [Sail CI](https://sail.ci/)                                                     | `ciinfo.SAIL`               | ✅   |
| [Screwdriver](https://screwdriver.cd/)                                          | `ciinfo.SCREWDRIVER`        | ✅   |
| [Semaphore](https://semaphoreci.com)                                            | `ciinfo.SEMAPHORE`          | ✅   |
| [Sourcehut](https://sourcehut.org/)                                             | `ciinfo.SOURCEHUT`          | 🚫   |
| [Strider CD](https://strider-cd.github.io/)                                     | `ciinfo.STRIDER`            | 🚫   |
| [TaskCluster](http://docs.taskcluster.net)                                      | `ciinfo.TASKCLUSTER`        | 🚫   |
| [TeamCity](https://www.jetbrains.com/teamcity/) by JetBrains                    | `ciinfo.TEAMCITY`           | 🚫   |
| [Travis CI](http://travis-ci.org)                                               | `ciinfo.TRAVIS`             | ✅   |
| [Vela](https://go-vela.github.io/docs/)                                         | `ci.VELA`                   | ✅   |
| [Vercel](https://vercel.com/)                                                   | `ciinfo.VERCEL`             | ✅   |
| [Visual Studio App Center](https://appcenter.ms/)                               | `ciinfo.APPCENTER`          | 🚫   |
| [Woodpecker](https://woodpecker-ci.org/)                                        | `ciinfo.WOODPECKER`         | ✅   |

## API

### `ciinfo.Name`

Returns a string containing name of the CI server the code is running on.
If CI server is not detected, it returns empty string `""`.

Don't depend on the value of this string not to change for a specific
vendor. If you find your self writing `ciinfo.Name === "Travis CI"`, you
most likely want to use `ciinfo.IsVendor("TRAVIS")` instead.

### `ciinfo.IsCI`

Returns a boolean. Will be `true` if the code is running on a CI server,
otherwise `false`.

Some CI servers not listed here might still trigger the `ciinfo.isCI`
boolean to be set to `true` if they use certain vendor neutral
environment variables. In those cases `ciinfo.Name` will be `""` and no
vendor specific boolean will be set to `true`.

### `ciinfo.IsPR`

Returns a boolean if PR detection is supported for the current CI server. Will
be `true` if a PR is being tested, otherwise `false`. If PR detection is
not supported for the current CI server, the value will be `false`.

### `ciinfo.<VENDOR-CONSTANT>`

A vendor specific boolean constant is exposed for each support CI
vendor. A constant will be `true` if the code is determined to run on
the given CI server, otherwise `false`.

Examples of vendor constants are `ciinfo.TRAVIS` or `ciinfo.APPVEYOR`. For a
complete list, see the support table above.
