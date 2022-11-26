package ciinfo

type pr struct {
	key string
	ne  string
	eq  string
}

type env struct {
	key      string
	eq       string
	includes string
}

type vendor struct {
	name     string
	constant string
	// if multiple keys on env only one is needed
	anyEnv bool
	env    []env
	pr     []pr
}

var vendors = []vendor{
	{
		name:     "Appcircle",
		constant: "APPCIRCLE",
		env:      []env{{key: "AC_APPCIRCLE"}},
	},
	{
		name:     "AppVeyor",
		constant: "APPVEYOR",
		env:      []env{{key: "APPVEYOR"}},
		pr: []pr{
			{
				key: "APPVEYOR_PULL_REQUEST_NUMBER",
			},
		},
	},
	{
		name:     "AWS CodeBuild",
		constant: "CODEBUILD",
		env:      []env{{key: "CODEBUILD_BUILD_ARN"}},
	},
	{
		name:     "Azure Pipelines",
		constant: "AZURE_PIPELINES",
		env:      []env{{key: "SYSTEM_TEAMFOUNDATIONCOLLECTIONURI"}},
		pr:       []pr{{key: "SYSTEM_PULLREQUEST_PULLREQUESTID"}},
	},
	{
		name:     "Bamboo",
		constant: "BAMBOO",
		env:      []env{{key: "bamboo_planKey"}},
	},
	{
		name:     "Bitbucket Pipelines",
		constant: "BITBUCKET",
		env:      []env{{key: "BITBUCKET_COMMIT"}},
		pr: []pr{
			{
				key: "BITBUCKET_PR_ID",
			},
		},
	},
	{
		name:     "Bitrise",
		constant: "BITRISE",
		env:      []env{{key: "BITRISE_IO"}},
		pr: []pr{
			{
				key: "BITRISE_PULL_REQUEST",
			},
		},
	},
	{
		name:     "Buddy",
		constant: "BUDDY",
		env:      []env{{key: "BUDDY_WORKSPACE_ID"}},
		pr: []pr{
			{
				key: "BUDDY_EXECUTION_PULL_REQUEST_ID",
			},
		},
	},
	{
		name:     "Buildkite",
		constant: "BUILDKITE",
		env:      []env{{key: "BUILDKITE"}},
		pr: []pr{
			{
				key: "BUILDKITE_PULL_REQUEST",
				ne:  "false",
			},
		},
	},
	{
		name:     "CircleCI",
		constant: "CIRCLE",
		env:      []env{{key: "CIRCLECI"}},
		pr: []pr{
			{
				key: "CIRCLE_PULL_REQUEST",
			},
		},
	},
	{
		name:     "Cirrus CI",
		constant: "CIRRUS",
		env:      []env{{key: "CIRRUS_CI"}},
		pr:       []pr{{key: "CIRRUS_PR"}},
	},
	{
		name:     "Codefresh",
		constant: "CODEFRESH",
		env:      []env{{key: "CF_BUILD_ID"}},
		pr: []pr{
			{key: "CF_PULL_REQUEST_NUMBER"},
			{key: "CF_PULL_REQUEST_ID"},
		},
	},
	{
		name:     "Codemagic",
		constant: "CODEMAGIC",
		env:      []env{{key: "CM_BUILD_ID"}},
		pr:       []pr{{key: "CM_PULL_REQUEST"}},
	},
	{
		name:     "Codeship",
		constant: "CODESHIP",
		env:      []env{{key: "CI_NAME", eq: "codeship"}},
	},
	{
		name:     "Drone",
		constant: "DRONE",
		env:      []env{{key: "DRONE"}},
		pr: []pr{
			{
				key: "DRONE_BUILD_EVENT",
				eq:  "pull_request",
			},
		},
	},
	{
		name:     "dsari",
		constant: "DSARI",
		env:      []env{{key: "DSARI"}},
	},
	{
		name:     "Expo Application Services",
		constant: "EAS",
		env:      []env{{key: "EAS_BUILD"}},
	},
	{
		name:     "Gerrit",
		constant: "GERRIT",
		env:      []env{{key: "GERRIT_PROJECT"}},
	},
	{
		name:     "GitHub Actions",
		constant: "GITHUB_ACTIONS",
		env:      []env{{key: "GITHUB_ACTIONS"}},
		pr: []pr{
			{
				key: "GITHUB_EVENT_NAME",
				eq:  "pull_request",
			},
		},
	},
	{
		name:     "GitLab CI",
		constant: "GITLAB",
		env:      []env{{key: "GITLAB_CI"}},
		pr: []pr{
			{
				key: "CI_MERGE_REQUEST_ID",
			},
		},
	},
	{
		name:     "GoCD",
		constant: "GOCD",
		env:      []env{{key: "GO_PIPELINE_LABEL"}},
	},
	{
		name:     "Google Cloud Build",
		constant: "GOOGLE_CLOUD_BUILD",
		env:      []env{{key: "BUILDER_OUTPUT"}},
	},
	{
		name:     "Hudson",
		constant: "HUDSON",
		env:      []env{{key: "HUDSON_URL"}},
	},
	{
		name:     "Heroku",
		constant: "HEROKU",
		env:      []env{{key: "NODE", includes: "/app/.heroku/node/bin/node"}},
	},
	{
		name:     "Jenkins",
		constant: "JENKINS",
		env:      []env{{key: "JENKINS_URL"}, {key: "BUILD_ID"}},
		pr: []pr{
			{
				key: "ghprbPullId",
			},
			{
				key: "CHANGE_ID",
			},
		},
	},
	{
		name:     "LayerCI",
		constant: "LAYERCI",
		env:      []env{{key: "LAYERCI"}},
		pr: []pr{
			{
				key: "LAYERCI_PULL_REQUEST",
			},
		},
	},
	{
		name:     "Magnum CI",
		constant: "MAGNUM",
		env:      []env{{key: "MAGNUM"}},
	},
	{
		name:     "Netlify CI",
		constant: "NETLIFY",
		env:      []env{{key: "NETLIFY"}},
		pr: []pr{
			{
				key: "PULL_REQUEST",
				ne:  "false",
			},
		},
	},
	{
		name:     "ReleaseHub",
		constant: "RELEASEHUB",
		env:      []env{{key: "RELEASE_BUILD_ID"}},
	},
	{
		name:     "Nevercode",
		constant: "NEVERCODE",
		env:      []env{{key: "NEVERCODE"}},
		pr: []pr{
			{
				key: "NEVERCODE_PULL_REQUEST",
				ne:  "false",
			},
		},
	},
	{
		name:     "Render",
		constant: "RENDER",
		env:      []env{{key: "RENDER"}},
		pr: []pr{
			{
				key: "IS_PULL_REQUEST",
				eq:  "true",
			},
		},
	},
	{
		name:     "Sail CI",
		constant: "SAIL",
		env:      []env{{key: "SAILCI"}},
		pr: []pr{
			{
				key: "SAIL_PULL_REQUEST_NUMBER",
			},
		},
	},
	{
		name:     "Screwdriver",
		constant: "SCREWDRIVER",
		env:      []env{{key: "SCREWDRIVER"}},
		pr: []pr{
			{
				key: "SD_PULL_REQUEST",
				ne:  "false",
			},
		},
	},
	{
		name:     "Semaphore",
		constant: "SEMAPHORE",
		env:      []env{{key: "SEMAPHORE"}},
		pr: []pr{
			{
				key: "PULL_REQUEST_NUMBER",
			},
		},
	},
	{
		name:     "Shippable",
		constant: "SHIPPABLE",
		env:      []env{{key: "SHIPPABLE"}},
		pr: []pr{
			{
				key: "IS_PULL_REQUEST",
				eq:  "true",
			},
		},
	},
	{
		name:     "Solano CI",
		constant: "SOLANO",
		env:      []env{{key: "TDDIUM"}},
		pr: []pr{
			{
				key: "TDDIUM_PR_ID",
			},
		},
	},
	{
		name:     "Sourcehut",
		constant: "SOURCEHUT",
		env:      []env{{key: "CI_NAME", eq: "sourcehut"}},
	},
	{
		name:     "Strider CD",
		constant: "STRIDER",
		env:      []env{{key: "STRIDER"}},
	},
	{
		name:     "TaskCluster",
		constant: "TASKCLUSTER",
		env:      []env{{key: "TASK_ID"}, {key: "RUN_ID"}},
	},
	{
		name:     "TeamCity",
		constant: "TEAMCITY",
		env:      []env{{key: "TEAMCITY_VERSION"}},
	},
	{
		name:     "Travis CI",
		constant: "TRAVIS",
		env:      []env{{key: "TRAVIS"}},
		pr: []pr{
			{
				key: "TRAVIS_PULL_REQUEST",
				ne:  "false",
			},
		},
	},
	{
		name:     "Vercel",
		constant: "VERCEL",
		anyEnv:   true,
		env:      []env{{key: "NOW_BUILDER"}, {key: "VERCEL"}},
	},
	{
		name:     "Visual Studio App Center",
		constant: "APPCENTER",
		env:      []env{{key: "APPCENTER_BUILD_ID"}},
	},
	{
		name:     "Woodpecker",
		constant: "WOODPECKER",
		env:      []env{{key: "CI", eq: "woodpecker"}},
		pr:       []pr{{key: "CI_BUILD_EVENT", eq: "pull_request"}},
	},
	{
		name:     "Xcode Cloud",
		constant: "XCODE_CLOUD",
		env:      []env{{key: "CI_XCODE_PROJECT"}},
		pr:       []pr{{key: "CI_PULL_REQUEST_NUMBER"}},
	},
	{
		name:     "Xcode Server",
		constant: "XCODE_SERVER",
		env:      []env{{key: "XCS"}},
	},
}
