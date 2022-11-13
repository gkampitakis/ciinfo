package ciinfo

import (
	"os"
	"strings"
)

var (
	Name        string
	IsPr        bool
	IsCI        bool
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

		for _, pr := range vendor.pr {
			if IsPr = verifyPr(pr); IsPr {
				break
			}
		}
	}

	IsCI = Name != "" || isCI()
}

func isCI() bool {
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
