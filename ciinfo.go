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

func initialize() {
	vendorsIsCI = make(map[string]bool)
	IsPr = false
	IsCI = false
	Name = ""

	for _, vendor := range vendors {
		envKeys := vendor.env
		vendorsIsCI[vendor.constant] = false

		for _, env := range envKeys {
			if verifyCI(env) {
				vendorsIsCI[vendor.constant] = true
				break
			}
		}

		if !vendorsIsCI[vendor.constant] {
			continue
		}

		Name = vendor.name

		for _, pr := range vendor.pr {
			IsPr = verifyPr(pr)
			if IsPr {
				break
			}
		}
	}

	IsCI = isCI()
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

	return Name != ""
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
