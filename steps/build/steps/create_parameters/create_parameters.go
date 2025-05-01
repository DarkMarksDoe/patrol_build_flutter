package create_parameters

import (
	constants "patrol_install/steps/build/constants"
	bp "patrol_install/steps/build/models/build_parameters"
	env "patrol_install/utils/environment"
)

func BuildParametersFromEnv() (*bp.BuildParameters, error) {
	envMap := map[string]string{
		"platform":     env.GetEnvironmentValue(constants.Platform),
		"targets":      env.GetEnvironmentValue(constants.Targets),
		"buildType":    env.GetEnvironmentValue(constants.BuildType),
		"tags":         env.GetEnvironmentValue(constants.Tags),
		"excludedTags": env.GetEnvironmentValue(constants.ExcludedTags),
		"verbose":      env.GetEnvironmentValue(constants.Verbose),
		"isCovered":    env.GetEnvironmentValue(constants.IsCovered),
	}

	// Final build
	return bp.NewBuildParameters(envMap)
}
