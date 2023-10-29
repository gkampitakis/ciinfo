package ciinfo

import (
	"os"
	"reflect"
	"testing"
)

func assertEqual(t *testing.T, expected, actual interface{}, s ...string) {
	t.Helper()

	if !reflect.DeepEqual(expected, actual) {
		message := ""
		if len(s) > 0 {
			message = " - " + s[0]
		}
		t.Errorf("\n[expected]: %v\n[actual]: %v%s", expected, actual, message)
	}
}

func isActualPr() bool {
	value, exists := os.LookupEnv("GITHUB_EVENT_NAME")

	return exists && value == "pull_request"
}

func assertVendorConstants(t *testing.T, expected string) {
	t.Helper()

	for _, vendor := range vendors {
		boolean := vendor.constant == expected
		boolean = expected == "SOLANO" && vendor.constant == "TDDIUM" ||
			boolean // support deprecated option

		assertEqual(t, boolean, vendorsIsCI[vendor.constant], "ci."+vendor.constant)
	}
}

type ScenarioExpected struct {
	isPR     bool
	name     string
	constant string
}
type TestScenario struct {
	description string
	setup       func(t *testing.T)
	expected    ScenarioExpected
}

func TestCI(t *testing.T) {
	t.Run("Known CI", func(t *testing.T) {
		t.Setenv("GITHUB_ACTIONS", "true")

		initialize()

		assertEqual(t, 47, len(vendors), "We should have 47 vendors")
		assertEqual(t, true, IsCI)
		assertEqual(t, isActualPr(), IsPr)
		assertEqual(t, "GitHub Actions", Name)
		assertVendorConstants(t, "GITHUB_ACTIONS")
	})

	t.Run("Not CI", func(t *testing.T) {
		t.Run("explicitly", func(t *testing.T) {
			os.Clearenv()
			// should ignore this and respect CI == false
			t.Setenv("BUILD_ID", "true")
			t.Setenv("CI", "false")

			initialize()

			assertEqual(t, false, IsCI)
			assertEqual(t, false, IsPr)
			assertEqual(t, "", Name)
			assertVendorConstants(t, "")
		})

		t.Run("implicitly", func(t *testing.T) {
			os.Clearenv()

			initialize()

			assertEqual(t, false, IsCI)
			assertEqual(t, false, IsPr)
			assertEqual(t, "", Name)
			assertVendorConstants(t, "")
		})
	})

	t.Run("Anonymous CI", func(t *testing.T) {
		envKeys := []string{
			"BUILD_ID",               // Jenkins, Cloudbees
			"BUILD_NUMBER",           // Jenkins, TeamCity
			"CI",                     // Travis CI, CircleCI, Cirrus CI, Gitlab CI, Appveyor, CodeShip, dsari
			"CI_APP_ID",              // Appflow
			"CI_BUILD_ID",            // Appflow
			"CI_BUILD_NUMBER",        // Appflow
			"CI_NAME",                // Codeship and others
			"CONTINUOUS_INTEGRATION", // Travis CI, Cirrus CI
			"RUN_ID",                 // TaskCluster, dsari
		}

		for _, key := range envKeys {
			t.Run(key, func(t *testing.T) {
				t.Setenv(key, "true")

				initialize()

				assertEqual(t, true, IsCI)
				assertEqual(t, false, IsPr)
				assertEqual(t, "", Name)
				assertVendorConstants(t, "")
			})
		}
	})

	for _, scenario := range []TestScenario{
		{
			description: "AppVeyor - PR",
			expected: ScenarioExpected{
				isPR:     true,
				name:     "AppVeyor",
				constant: "APPVEYOR",
			},
			setup: func(t *testing.T) {
				t.Setenv("APPVEYOR", "true")
				t.Setenv("APPVEYOR_PULL_REQUEST_NUMBER", "42")
			},
		},
		{
			description: "AppVeyor - Not PR",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "AppVeyor",
				constant: "APPVEYOR",
			},
			setup: func(t *testing.T) {
				t.Setenv("APPVEYOR", "true")
			},
		},
		{
			description: "Azure Pipelines - PR",
			expected: ScenarioExpected{
				isPR:     true,
				name:     "Azure Pipelines",
				constant: "AZURE_PIPELINES",
			},
			setup: func(t *testing.T) {
				t.Setenv("TF_BUILD", "true")
				t.Setenv("BUILD_REASON", "PullRequest")
			},
		},
		{
			description: "Azure Pipelines - Not PR",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "Azure Pipelines",
				constant: "AZURE_PIPELINES",
			},
			setup: func(t *testing.T) {
				t.Setenv("TF_BUILD", "true")
			},
		},
		{
			description: "Bitbucket Pipelines - PR",
			expected: ScenarioExpected{
				isPR:     true,
				name:     "Bitbucket Pipelines",
				constant: "BITBUCKET",
			},
			setup: func(t *testing.T) {
				t.Setenv("BITBUCKET_COMMIT", "true")
				t.Setenv("BITBUCKET_PR_ID", "42")
			},
		},
		{
			description: "Bitbucket Pipelines - Not PR",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "Bitbucket Pipelines",
				constant: "BITBUCKET",
			},
			setup: func(t *testing.T) {
				t.Setenv("BITBUCKET_COMMIT", "true")
			},
		},
		{
			description: "Buildkite - PR",
			expected: ScenarioExpected{
				isPR:     true,
				name:     "Buildkite",
				constant: "BUILDKITE",
			},
			setup: func(t *testing.T) {
				t.Setenv("BUILDKITE", "true")
				t.Setenv("BUILDKITE_PULL_REQUEST", "42")
			},
		},
		{
			description: "Buildkite - Not PR",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "Buildkite",
				constant: "BUILDKITE",
			},
			setup: func(t *testing.T) {
				t.Setenv("BUILDKITE", "true")
				t.Setenv("BUILDKITE_PULL_REQUEST", "false")
			},
		},
		{
			description: "CircleCI - PR",
			expected: ScenarioExpected{
				isPR:     true,
				name:     "CircleCI",
				constant: "CIRCLE",
			},
			setup: func(t *testing.T) {
				t.Setenv("CIRCLECI", "true")
				t.Setenv("CIRCLE_PULL_REQUEST", "42")
			},
		},
		{
			description: "CircleCI - Not PR",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "CircleCI",
				constant: "CIRCLE",
			},
			setup: func(t *testing.T) {
				t.Setenv("CIRCLECI", "true")
			},
		},
		{
			description: "Cirrus CI - PR",
			expected: ScenarioExpected{
				isPR:     true,
				name:     "Cirrus CI",
				constant: "CIRRUS",
			},
			setup: func(t *testing.T) {
				t.Setenv("CIRRUS_CI", "true")
				t.Setenv("CIRRUS_PR", "42")
			},
		},
		{
			description: "Cirrus CI - Not PR",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "Cirrus CI",
				constant: "CIRRUS",
			},
			setup: func(t *testing.T) {
				t.Setenv("CIRRUS_CI", "true")
			},
		},
		{
			description: "Codemagic - PR",
			expected: ScenarioExpected{
				isPR:     true,
				name:     "Codemagic",
				constant: "CODEMAGIC",
			},
			setup: func(t *testing.T) {
				t.Setenv("CM_BUILD_ID", "true")
				t.Setenv("CM_PULL_REQUEST", "42")
			},
		},
		{
			description: "Codemagic - Not PR",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "Codemagic",
				constant: "CODEMAGIC",
			},
			setup: func(t *testing.T) {
				t.Setenv("CM_BUILD_ID", "true")
			},
		},
		{
			description: "Codefresh - PR",
			expected: ScenarioExpected{
				isPR:     true,
				name:     "Codefresh",
				constant: "CODEFRESH",
			},
			setup: func(t *testing.T) {
				t.Setenv("CF_BUILD_ID", "true")
				t.Setenv("CF_PULL_REQUEST_ID", "42")
			},
		},
		{
			description: "Codefresh - PR 2",
			expected: ScenarioExpected{
				isPR:     true,
				name:     "Codefresh",
				constant: "CODEFRESH",
			},
			setup: func(t *testing.T) {
				t.Setenv("CF_BUILD_ID", "true")
				t.Setenv("CF_PULL_REQUEST_NUMBER", "42")
			},
		},
		{
			description: "Codeship",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "Codeship",
				constant: "CODESHIP",
			},
			setup: func(t *testing.T) {
				t.Setenv("CI_NAME", "codeship")
			},
		},
		{
			description: "Codefresh - Not PR",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "Codefresh",
				constant: "CODEFRESH",
			},
			setup: func(t *testing.T) {
				t.Setenv("CF_BUILD_ID", "true")
			},
		},
		{
			description: "Drone",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "Drone",
				constant: "DRONE",
			},
			setup: func(t *testing.T) {
				t.Setenv("DRONE", "true")
			},
		},
		{
			description: "Drone 2",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "Drone",
				constant: "DRONE",
			},
			setup: func(t *testing.T) {
				t.Setenv("DRONE", "true")
				t.Setenv("DRONE_BUILD_EVENT", "test")
			},
		},
		{
			description: "Drone - Pr",
			expected: ScenarioExpected{
				isPR:     true,
				name:     "Drone",
				constant: "DRONE",
			},
			setup: func(t *testing.T) {
				t.Setenv("DRONE", "true")
				t.Setenv("DRONE_BUILD_EVENT", "pull_request")
			},
		},
		{
			description: "Jenkins - PR",
			expected: ScenarioExpected{
				isPR:     true,
				name:     "Jenkins",
				constant: "JENKINS",
			},
			setup: func(t *testing.T) {
				t.Setenv("JENKINS_URL", "true")
				t.Setenv("BUILD_ID", "true")
				t.Setenv("ghprbPullId", "true")
			},
		},
		{
			description: "Jenkins - PR 2",
			expected: ScenarioExpected{
				isPR:     true,
				name:     "Jenkins",
				constant: "JENKINS",
			},
			setup: func(t *testing.T) {
				t.Setenv("JENKINS_URL", "true")
				t.Setenv("BUILD_ID", "true")
				t.Setenv("CHANGE_ID", "true")
			},
		},
		{
			description: "Jenkins - Not PR",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "Jenkins",
				constant: "JENKINS",
			},
			setup: func(t *testing.T) {
				t.Setenv("JENKINS_URL", "true")
				t.Setenv("BUILD_ID", "true")
			},
		},
		{
			description: "LayerCI - PR",
			expected: ScenarioExpected{
				isPR:     true,
				name:     "LayerCI",
				constant: "LAYERCI",
			},
			setup: func(t *testing.T) {
				t.Setenv("LAYERCI", "true")
				t.Setenv("LAYERCI_PULL_REQUEST", "LAYERCI_PULL_REQUEST")
			},
		},
		{
			description: "LayerCI - Not PR",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "LayerCI",
				constant: "LAYERCI",
			},
			setup: func(t *testing.T) {
				t.Setenv("LAYERCI", "true")
			},
		},
		{
			description: "Render - PR",
			expected: ScenarioExpected{
				isPR:     true,
				name:     "Render",
				constant: "RENDER",
			},
			setup: func(t *testing.T) {
				t.Setenv("RENDER", "true")
				t.Setenv("IS_PULL_REQUEST", "true")
			},
		},
		{
			description: "Render - Not PR",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "Render",
				constant: "RENDER",
			},
			setup: func(t *testing.T) {
				t.Setenv("RENDER", "true")
				t.Setenv("IS_PULL_REQUEST", "false")
			},
		},
		{
			description: "Semaphore - PR",
			expected: ScenarioExpected{
				isPR:     true,
				name:     "Semaphore",
				constant: "SEMAPHORE",
			},
			setup: func(t *testing.T) {
				t.Setenv("SEMAPHORE", "true")
				t.Setenv("PULL_REQUEST_NUMBER", "42")
			},
		},
		{
			description: "Semaphore - Not PR",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "Semaphore",
				constant: "SEMAPHORE",
			},
			setup: func(t *testing.T) {
				t.Setenv("SEMAPHORE", "true")
			},
		},
		{
			description: "Solano CI - PR",
			expected: ScenarioExpected{
				isPR:     true,
				name:     "Solano CI",
				constant: "SOLANO",
			},
			setup: func(t *testing.T) {
				t.Setenv("TDDIUM", "true")
				t.Setenv("TDDIUM_PR_ID", "42")
			},
		},
		{
			description: "Solano CI - Not PR",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "Solano CI",
				constant: "SOLANO",
			},
			setup: func(t *testing.T) {
				t.Setenv("TDDIUM", "true")
			},
		},
		{
			description: "Sourcehut",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "Sourcehut",
				constant: "SOURCEHUT",
			},
			setup: func(t *testing.T) {
				t.Setenv("CI_NAME", "sourcehut")
			},
		},
		{
			description: "Travis CI - PR",
			expected: ScenarioExpected{
				isPR:     true,
				name:     "Travis CI",
				constant: "TRAVIS",
			},
			setup: func(t *testing.T) {
				t.Setenv("TRAVIS", "true")
				t.Setenv("TRAVIS_PULL_REQUEST", "42")
			},
		},
		{
			description: "Travis CI - Not PR",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "Travis CI",
				constant: "TRAVIS",
			},
			setup: func(t *testing.T) {
				t.Setenv("TRAVIS", "true")
				t.Setenv("TRAVIS_PULL_REQUEST", "false")
			},
		},
		{
			description: "Netlify CI - PR",
			expected: ScenarioExpected{
				isPR:     true,
				name:     "Netlify CI",
				constant: "NETLIFY",
			},
			setup: func(t *testing.T) {
				t.Setenv("NETLIFY", "true")
				t.Setenv("PULL_REQUEST", "true")
			},
		},
		{
			description: "Netlify CI - Not PR",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "Netlify CI",
				constant: "NETLIFY",
			},
			setup: func(t *testing.T) {
				t.Setenv("NETLIFY", "true")
				t.Setenv("PULL_REQUEST", "false")
			},
		},
		{
			description: "ReleaseHub",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "ReleaseHub",
				constant: "RELEASEHUB",
			},
			setup: func(t *testing.T) {
				t.Setenv("RELEASE_BUILD_ID", "")
			},
		},
		{
			description: "Vercel - NOW_BUILDER",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "Vercel",
				constant: "VERCEL",
			},
			setup: func(t *testing.T) {
				t.Setenv("NOW_BUILDER", "1")
			},
		},
		{
			description: "Vercel - VERCEL",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "Vercel",
				constant: "VERCEL",
			},
			setup: func(t *testing.T) {
				t.Setenv("VERCEL", "1")
			},
		},
		{
			description: "Vercel - PR",
			expected: ScenarioExpected{
				isPR:     true,
				name:     "Vercel",
				constant: "VERCEL",
			},
			setup: func(t *testing.T) {
				t.Setenv("VERCEL", "1")
				t.Setenv("VERCEL_GIT_PULL_REQUEST_ID", "23")
			},
		},
		{
			description: "Nevercode - PR",
			expected: ScenarioExpected{
				isPR:     true,
				name:     "Nevercode",
				constant: "NEVERCODE",
			},
			setup: func(t *testing.T) {
				t.Setenv("NEVERCODE", "true")
				t.Setenv("NEVERCODE_PULL_REQUEST", "true")
			},
		},
		{
			description: "Nevercode - Not PR",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "Nevercode",
				constant: "NEVERCODE",
			},
			setup: func(t *testing.T) {
				t.Setenv("NEVERCODE", "true")
				t.Setenv("NEVERCODE_PULL_REQUEST", "false")
			},
		},
		{
			description: "Expo Application Services",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "Expo Application Services",
				constant: "EAS",
			},
			setup: func(t *testing.T) {
				t.Setenv("EAS_BUILD", "1")
			},
		},
		{
			description: "GitHub Actions - PR",
			expected: ScenarioExpected{
				isPR:     true,
				name:     "GitHub Actions",
				constant: "GITHUB_ACTIONS",
			},
			setup: func(t *testing.T) {
				t.Setenv("GITHUB_ACTIONS", "true")
				t.Setenv("GITHUB_EVENT_NAME", "pull_request")
			},
		},
		{
			description: "GitHub Actions - Not PR",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "GitHub Actions",
				constant: "GITHUB_ACTIONS",
			},
			setup: func(t *testing.T) {
				t.Setenv("GITHUB_ACTIONS", "true")
				t.Setenv("GITHUB_EVENT_NAME", "push")
			},
		},
		{
			description: "Screwdriver - PR",
			expected: ScenarioExpected{
				isPR:     true,
				name:     "Screwdriver",
				constant: "SCREWDRIVER",
			},
			setup: func(t *testing.T) {
				t.Setenv("SCREWDRIVER", "true")
				t.Setenv("SD_PULL_REQUEST", "1")
			},
		},
		{
			description: "Screwdriver - Not PR",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "Screwdriver",
				constant: "SCREWDRIVER",
			},
			setup: func(t *testing.T) {
				t.Setenv("SCREWDRIVER", "true")
				t.Setenv("SD_PULL_REQUEST", "false")
			},
		},
		{
			description: "Visual Studio App Center",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "Visual Studio App Center",
				constant: "APPCENTER",
			},
			setup: func(t *testing.T) {
				t.Setenv("APPCENTER_BUILD_ID", "1")
			},
		},
		{
			description: "Xcode Cloud - Not PR",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "Xcode Cloud",
				constant: "XCODE_CLOUD",
			},
			setup: func(t *testing.T) {
				t.Setenv("CI_XCODE_PROJECT", "1")
			},
		},
		{
			description: "Xcode Cloud - PR",
			expected: ScenarioExpected{
				isPR:     true,
				name:     "Xcode Cloud",
				constant: "XCODE_CLOUD",
			},
			setup: func(t *testing.T) {
				t.Setenv("CI_XCODE_PROJECT", "1")
				t.Setenv("CI_PULL_REQUEST_NUMBER", "1")
			},
		},
		{
			description: "Xcode Server",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "Xcode Server",
				constant: "XCODE_SERVER",
			},
			setup: func(t *testing.T) {
				t.Setenv("XCS", "1")
			},
		},
		{
			description: "Woodpecker - PR",
			expected: ScenarioExpected{
				isPR:     true,
				name:     "Woodpecker",
				constant: "WOODPECKER",
			},
			setup: func(t *testing.T) {
				t.Setenv("CI", "woodpecker")
				t.Setenv("CI_BUILD_EVENT", "pull_request")
			},
		},
		{
			description: "Woodpecker",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "Woodpecker",
				constant: "WOODPECKER",
			},
			setup: func(t *testing.T) {
				t.Setenv("CI", "woodpecker")
			},
		},
		{
			description: "Heroku",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "Heroku",
				constant: "HEROKU",
			},
			setup: func(t *testing.T) {
				t.Setenv("NODE", "/extra/content/app/.heroku/node/bin/node --extra --content")
			},
		},
		{
			description: "Gerrit",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "Gerrit",
				constant: "GERRIT",
			},
			setup: func(t *testing.T) {
				t.Setenv("GERRIT_PROJECT", "1")
			},
		},
		{
			description: "Google Cloud Build",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "Google Cloud Build",
				constant: "GOOGLE_CLOUD_BUILD",
			},
			setup: func(t *testing.T) {
				t.Setenv("BUILDER_OUTPUT", "1")
			},
		},
		{
			description: "Harness CI",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "Harness CI",
				constant: "HARNESS",
			},
			setup: func(t *testing.T) {
				t.Setenv("HARNESS_BUILD_ID", "1")
			},
		},
		{
			description: "Gitea Actions",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "Gitea Actions",
				constant: "GITEA_ACTIONS",
			},
			setup: func(t *testing.T) {
				t.Setenv("GITEA_ACTIONS", "")
			},
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			scenario.setup(t)

			initialize()

			assertEqual(t, true, IsCI)
			assertEqual(t, scenario.expected.isPR, IsPr)
			assertEqual(t, scenario.expected.name, Name)
			assertEqual(t, true, IsVendor(scenario.expected.constant))
			assertVendorConstants(t, scenario.expected.constant)
		})
	}
}
