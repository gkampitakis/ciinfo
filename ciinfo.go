package ciinfo

import (
	"os"
	"strings"
)

var (
	// Returns a boolean if PR detection is supported for the current CI server.
	// Will be `true` if a PR is being tested, otherwise `false`. If PR detection is
	// not supported for the current CI server, the value will be `null`.
	IsPr bool
	//  Returns a boolean. Will be `true` if the code is running on a CI server,
	//  otherwise `false`.
	//
	//  Some CI servers not listed here might still trigger the `ci.isCI`
	//  boolean to be set to `true` if they use certain vendor neutral environment
	//  variables. In those cases `ci.name` will be `null` and no vendor specific
	//  boolean will be set to `true`.
	IsCI bool
	// Returns a string containing name of the CI server the code is running on. If
	// CI server is not detected, it returns `null`.
	//
	// Don't depend on the value of this string not to change for a specific vendor.
	// If you find your self writing `ci.name === 'Travis CI'`, you most likely want
	// to use `ci.TRAVIS` instead.
	Name string
	// Returns a string containing the identifier of the CI server the code is running on. If
	// CI server is not detected, it returns `null`.
	ID          string
	vendorsIsCI map[string]bool
)

func init() {
	initialize()
}

func everyEnv(envs []env, check func(env) bool) bool {
	for _, e := range envs {
		if !check(e) {
			return false
		}
	}
	return true
}

func anyEnv(envs []env, check func(env) bool) bool {
	for _, env := range envs {
		if check(env) {
			return true
		}
	}
	return false
}

func initialize() {
	vendorsIsCI = make(map[string]bool)
	IsPr = false
	IsCI = false
	Name = ""

	for _, vendor := range vendors {
		checkEnvs := everyEnv
		if vendor.anyEnv {
			checkEnvs = anyEnv
		}
		isVendor := checkEnvs(vendor.env, verifyCI)
		vendorsIsCI[vendor.constant] = isVendor
		if !isVendor {
			continue
		}

		Name = vendor.name
		IsPr = anyPr(vendor.pr)
		ID = vendor.constant
	}

	IsCI = os.Getenv("CI") != "false" && (Name != "" || isCI())
}

func isCI() bool {
	envKeys := []string{
		"BUILD_ID",               // Jenkins, Cloudbees
		"BUILD_NUMBER",           // Jenkins, TeamCity
		"CI",                     // Travis CI, CircleCI, Cirrus CI, Gitlab CI, Appveyor, CodeShip, dsari, Cloudflare Pages/Workers
		"CI_APP_ID",              // Appflow
		"CI_BUILD_ID",            // Appflow
		"CI_BUILD_NUMBER",        // Appflow
		"CI_NAME",                // Codeship and others
		"CONTINUOUS_INTEGRATION", // Travis CI, Cirrus CI
		"RUN_ID",                 // TaskCluster, dsari
	}

	for _, key := range envKeys {
		if _, exists := os.LookupEnv(key); exists {
			return true
		}
	}

	return false
}

func IsVendor(vendor string) bool {
	return vendorsIsCI[vendor]
}

func verifyCI(e env) bool {
	value, exists := os.LookupEnv(e.key)

	if !exists {
		return false
	}

	if e.includes != "" && !strings.Contains(value, e.includes) {
		return false
	}

	if e.eq != "" && value != e.eq {
		return false
	}

	return true
}

func anyPr(pr []pr) bool {
	for _, p := range pr {
		if verifyPr(p) {
			return true
		}
	}

	return false
}

func verifyPr(p pr) bool {
	value, exists := os.LookupEnv(p.key)
	if !exists {
		return false
	}

	// if we have eq constraint then env value must be == with the eq
	if p.eq != "" {
		return value == p.eq
	}

	// if we have ne constraint then env value must be != with the ne
	if p.ne != "" {
		return value != p.ne
	}

	return true
}
