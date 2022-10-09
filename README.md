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

## Supported CI tools

Officially supported CI servers:

| Name                                                                            | Constant             | isPR |
| ------------------------------------------------------------------------------- | -------------------- | ---- |
| [AWS CodeBuild](https://aws.amazon.com/codebuild/)                              | `CODEBUILD`       | 🚫   |
| [AppVeyor](http://www.appveyor.com)                                             | `APPVEYOR`        | ✅   |
| [Azure Pipelines](https://azure.microsoft.com/en-us/services/devops/pipelines/) | `AZURE_PIPELINES` | ✅   |
| [Appcircle](https://appcircle.io/)                                              | `APPCIRCLE`       | 🚫   |
| [Bamboo](https://www.atlassian.com/software/bamboo) by Atlassian                | `BAMBOO`          | 🚫   |
| [Bitbucket Pipelines](https://bitbucket.org/product/features/pipelines)         | `BITBUCKET`       | ✅   |
| [Bitrise](https://www.bitrise.io/)                                              | `BITRISE`         | ✅   |
| [Buddy](https://buddy.works/)                                                   | `BUDDY`           | ✅   |
| [Buildkite](https://buildkite.com)                                              | `BUILDKITE`       | ✅   |
| [CircleCI](http://circleci.com)                                                 | `CIRCLE`          | ✅   |
| [Cirrus CI](https://cirrus-ci.org)                                              | `CIRRUS`          | ✅   |
| [Codefresh](https://codefresh.io/)                                              | `CODEFRESH`       | ✅   |
| [Codeship](https://codeship.com)                                                | `CODESHIP`        | 🚫   |
| [Drone](https://drone.io)                                                       | `DRONE`           | ✅   |
| [dsari](https://github.com/rfinnie/dsari)                                       | `DSARI`           | 🚫   |
| [Expo Application Services](https://expo.dev/eas)                               | `EAS_BUILD`       | 🚫   |
| [GitHub Actions](https://github.com/features/actions/)                          | `GITHUB_ACTIONS`  | ✅   |
| [GitLab CI](https://about.gitlab.com/gitlab-ci/)                                | `GITLAB`          | ✅   |
| [GoCD](https://www.go.cd/)                                                      | `GOCD`            | 🚫   |
| [Hudson](http://hudson-ci.org)                                                  | `HUDSON`          | 🚫   |
| [Jenkins CI](https://jenkins-ci.org)                                            | `JENKINS`         | ✅   |
| [LayerCI](https://layerci.com/)                                                 | `LAYERCI`         | ✅   |
| [Magnum CI](https://magnum-ci.com)                                              | `MAGNUM`          | 🚫   |
| [Netlify CI](https://www.netlify.com/)                                          | `NETLIFY`         | ✅   |
| [Nevercode](http://nevercode.io/)                                               | `NEVERCODE`       | ✅   |
| [Render](https://render.com/)                                                   | `RENDER`          | ✅   |
| [Sail CI](https://sail.ci/)                                                     | `SAIL`            | ✅   |
| [Screwdriver](https://screwdriver.cd/)                                          | `SCREWDRIVER`     | ✅   |
| [Semaphore](https://semaphoreci.com)                                            | `SEMAPHORE`       | ✅   |
| [Shippable](https://www.shippable.com/)                                         | `SHIPPABLE`       | ✅   |
| [Solano CI](https://www.solanolabs.com/)                                        | `SOLANO`          | ✅   |
| [Strider CD](https://strider-cd.github.io/)                                     | `STRIDER`         | 🚫   |
| [TaskCluster](http://docs.taskcluster.net)                                      | `TASKCLUSTER`     | 🚫   |
| [TeamCity](https://www.jetbrains.com/teamcity/) by JetBrains                    | `TEAMCITY`        | 🚫   |
| [Travis CI](http://travis-ci.org)                                               | `TRAVIS`          | ✅   |
| [Vercel](https://vercel.com/)                                                   | `VERCEL`          | 🚫   |
| [Visual Studio App Center](https://appcenter.ms/)                               | `APPCENTER`       | 🚫   |
| [Woodpecker](https://woodpecker-ci.org/)                                        | `ci.WOODPECKER`   | ✅   |

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

### `ciinfo.IsVendor(<VENDOR-CONSTANT>)`

A vendor specific boolean constant is exposed for each support CI
vendor. A constant will be `true` if the code is determined to run on
the given CI server, otherwise `false`.

Examples of vendor constants are `ciinfo.IsVendor("TRAVIS")` or `ciinfo.IsVendor("APPVEYOR")`. For a
complete list, see the support table above.

Deprecated vendor constants that will be removed in the next major
release:

- `ciinfo.IsVendor("TDDIUM")` (Solano CI) This has been renamed `ciinfo.IsVendor("SOLANO")`

## License

[MIT](LICENSE)
