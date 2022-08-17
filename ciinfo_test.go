package ciinfo

import (
	"os"
	"reflect"
	"testing"
)

// NOTE: backporting from 1.17
func setEnv(t *testing.T, key, value string) {
	t.Helper()

	if prevVal, exists := os.LookupEnv(key); exists {
		t.Cleanup(func() {
			os.Setenv(key, prevVal)
		})
	} else {
		t.Cleanup(func() {
			os.Unsetenv(key)
		})
	}

	os.Setenv(key, value)
}

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
		boolean = expected == "SOLANO" && vendor.constant == "TDDIUM" || boolean // support deprecated option

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
		setEnv(t, "GITHUB_ACTIONS", "true")

		initialize()

		assertEqual(t, 37, len(vendors), "We should have 37 vendors")
		assertEqual(t, true, IsCI)
		assertEqual(t, isActualPr(), IsPr)
		assertEqual(t, "GitHub Actions", Name)
		assertVendorConstants(t, "GITHUB_ACTIONS")
	})

	t.Run("Not CI", func(t *testing.T) {
		os.Clearenv()

		initialize()

		assertEqual(t, false, IsCI)
		assertEqual(t, false, IsPr)
		assertEqual(t, "", Name)
		assertVendorConstants(t, "")
	})

	t.Run("Unknown CI", func(t *testing.T) {
		setEnv(t, "CI", "true")

		initialize()

		assertEqual(t, true, IsCI)
		assertEqual(t, false, IsPr)
		assertEqual(t, "", Name)
		assertVendorConstants(t, "")
	})

	t.Run("Not Codeship", func(t *testing.T) {
		setEnv(t, "CI_NAME", "invalid")

		initialize()

		assertEqual(t, false, IsCI)
		assertEqual(t, false, IsPr)
		assertEqual(t, "", Name)
		assertEqual(t, false, IsVendor("CODESHIP"))
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
				setEnv(t, "APPVEYOR", "true")
				setEnv(t, "APPVEYOR_PULL_REQUEST_NUMBER", "42")
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
				setEnv(t, "APPVEYOR", "true")
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
				setEnv(t, "SYSTEM_TEAMFOUNDATIONCOLLECTIONURI", "https://dev.azure.com/Contoso")
				setEnv(t, "SYSTEM_PULLREQUEST_PULLREQUESTID", "42")
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
				setEnv(t, "SYSTEM_TEAMFOUNDATIONCOLLECTIONURI", "https://dev.azure.com/Contoso")
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
				setEnv(t, "BITBUCKET_COMMIT", "true")
				setEnv(t, "BITBUCKET_PR_ID", "42")
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
				setEnv(t, "BITBUCKET_COMMIT", "true")
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
				setEnv(t, "BUILDKITE", "true")
				setEnv(t, "BUILDKITE_PULL_REQUEST", "42")
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
				setEnv(t, "BUILDKITE", "true")
				setEnv(t, "BUILDKITE_PULL_REQUEST", "false")
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
				setEnv(t, "CIRCLECI", "true")
				setEnv(t, "CIRCLE_PULL_REQUEST", "42")
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
				setEnv(t, "CIRCLECI", "true")
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
				setEnv(t, "CIRRUS_CI", "true")
				setEnv(t, "CIRRUS_PR", "42")
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
				setEnv(t, "CIRRUS_CI", "true")
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
				setEnv(t, "CF_BUILD_ID", "true")
				setEnv(t, "CF_PULL_REQUEST_ID", "42")
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
				setEnv(t, "CF_BUILD_ID", "true")
				setEnv(t, "CF_PULL_REQUEST_NUMBER", "42")
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
				setEnv(t, "CI_NAME", "codeship")
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
				setEnv(t, "CF_BUILD_ID", "true")
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
				setEnv(t, "DRONE", "true")
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
				setEnv(t, "DRONE", "true")
				setEnv(t, "DRONE_BUILD_EVENT", "test")
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
				setEnv(t, "DRONE", "true")
				setEnv(t, "DRONE_BUILD_EVENT", "pull_request")
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
				setEnv(t, "JENKINS_URL", "true")
				setEnv(t, "ghprbPullId", "true")
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
				setEnv(t, "BUILD_ID", "true")
				setEnv(t, "CHANGE_ID", "true")
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
				setEnv(t, "BUILD_ID", "true")
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
				setEnv(t, "LAYERCI", "true")
				setEnv(t, "LAYERCI_PULL_REQUEST", "LAYERCI_PULL_REQUEST")
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
				setEnv(t, "LAYERCI", "true")
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
				setEnv(t, "RENDER", "true")
				setEnv(t, "IS_PULL_REQUEST", "true")
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
				setEnv(t, "RENDER", "true")
				setEnv(t, "IS_PULL_REQUEST", "false")
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
				setEnv(t, "SEMAPHORE", "true")
				setEnv(t, "PULL_REQUEST_NUMBER", "42")
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
				setEnv(t, "SEMAPHORE", "true")
			},
		},
		{
			description: "Shippable - PR",
			expected: ScenarioExpected{
				isPR:     true,
				name:     "Shippable",
				constant: "SHIPPABLE",
			},
			setup: func(t *testing.T) {
				setEnv(t, "SHIPPABLE", "true")
				setEnv(t, "IS_PULL_REQUEST", "true")
			},
		},
		{
			description: "Shippable - Not PR",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "Shippable",
				constant: "SHIPPABLE",
			},
			setup: func(t *testing.T) {
				setEnv(t, "SHIPPABLE", "true")
				setEnv(t, "IS_PULL_REQUEST", "false")
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
				setEnv(t, "TDDIUM", "true")
				setEnv(t, "TDDIUM_PR_ID", "42")
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
				setEnv(t, "TDDIUM", "true")
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
				setEnv(t, "TRAVIS", "true")
				setEnv(t, "TRAVIS_PULL_REQUEST", "42")
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
				setEnv(t, "TRAVIS", "true")
				setEnv(t, "TRAVIS_PULL_REQUEST", "false")
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
				setEnv(t, "NETLIFY", "true")
				setEnv(t, "PULL_REQUEST", "true")
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
				setEnv(t, "NETLIFY", "true")
				setEnv(t, "PULL_REQUEST", "false")
			},
		},
		{
			description: "Vercel",
			expected: ScenarioExpected{
				isPR:     false,
				name:     "Vercel",
				constant: "VERCEL",
			},
			setup: func(t *testing.T) {
				setEnv(t, "NOW_BUILDER", "1")
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
				setEnv(t, "NEVERCODE", "true")
				setEnv(t, "NEVERCODE_PULL_REQUEST", "true")
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
				setEnv(t, "NEVERCODE", "true")
				setEnv(t, "NEVERCODE_PULL_REQUEST", "false")
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
				setEnv(t, "EAS_BUILD", "1")
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
				setEnv(t, "GITHUB_ACTIONS", "true")
				setEnv(t, "GITHUB_EVENT_NAME", "pull_request")
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
				setEnv(t, "GITHUB_ACTIONS", "true")
				setEnv(t, "GITHUB_EVENT_NAME", "push")
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
				setEnv(t, "SCREWDRIVER", "true")
				setEnv(t, "SD_PULL_REQUEST", "1")
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
				setEnv(t, "SCREWDRIVER", "true")
				setEnv(t, "SD_PULL_REQUEST", "false")
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
				setEnv(t, "APPCENTER_BUILD_ID", "1")
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
