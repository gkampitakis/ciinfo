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
		name:     "Agola CI",
		constant: "AGOLA",
		env:      []env{{key: "AGOLA_GIT_REF"}},
		pr:       []pr{{key: "AGOLA_PULL_REQUEST_ID"}},
	},
	{
		name:     "Appcircle",
		constant: "APPCIRCLE",
		env:      []env{{key: "AC_APPCIRCLE"}},
		pr:       []pr{{key: "AC_GIT_PR", ne: "false"}},
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
		pr: []pr{
			{key: "CODEBUILD_WEBHOOK_EVENT", eq: "PULL_REQUEST_CREATED"},
			{key: "CODEBUILD_WEBHOOK_EVENT", eq: "PULL_REQUEST_UPDATED"},
			{key: "CODEBUILD_WEBHOOK_EVENT", eq: "PULL_REQUEST_REOPENED"},
		},
	},
	{
		name:     "Azure Pipelines",
		constant: "AZURE_PIPELINES",
		env:      []env{{key: "TF_BUILD"}},
		pr:       []pr{{key: "BUILD_REASON", eq: "PullRequest"}},
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
		name:     "Cloudflare Pages",
		constant: "CLOUDFLARE_PAGES",
		env:      []env{{key: "CF_PAGES"}},
	},
	{
		name:     "Cloudflare Workers",
		constant: "CLOUDFLARE_WORKERS",
		env:      []env{{key: "WORKERS_CI"}},
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
		name:     "Earthly",
		constant: "EARTHLY",
		env:      []env{{key: "EARTHLY_CI"}},
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
		name:     "Gitea Actions",
		constant: "GITEA_ACTIONS",
		env:      []env{{key: "GITEA_ACTIONS"}},
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
		name:     "Harness CI",
		constant: "HARNESS",
		env:      []env{{key: "HARNESS_BUILD_ID"}},
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
		name:     "Prow",
		constant: "PROW",
		env:      []env{{key: "PROW_JOB_ID"}},
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
		name:     "Vela",
		constant: "VELA",
		env:      []env{{key: "VELA"}},
		pr:       []pr{{key: "VELA_PULL_REQUEST", eq: "1"}},
	},
	{
		name:     "Vercel",
		constant: "VERCEL",
		anyEnv:   true,
		env:      []env{{key: "NOW_BUILDER"}, {key: "VERCEL"}},
		pr: []pr{
			{key: "VERCEL_GIT_PULL_REQUEST_ID"},
		},
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
