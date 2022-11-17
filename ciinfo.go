package ciinfo

import (
	"os"
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

func initialize() {
	vendorsIsCI = make(map[string]bool)
	IsPr = false
	IsCI = false
	Name = ""

	for _, vendor := range vendors {
		isVendor := everyEnv(vendor.env, verifyCI)
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
		"CI",                     // Travis CI, CircleCI, Cirrus CI, Gitlab CI, Appveyor, CodeShip, dsari
		"CONTINUOUS_INTEGRATION", // Travis CI, Cirrus CI
		"BUILD_NUMBER",           // Jenkins, TeamCity
		"RUN_ID",                 // TaskCluster, dsari
		"CI_APP_ID",              // Applfow
		"CI_BUILD_ID",            // Applfow
		"CI_BUILD_NUMBER",        // Applfow
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
