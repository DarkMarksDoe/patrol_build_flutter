package validate

import (
	v "github.com/Masterminds/semver/v3"

	"patrol_install/utils/print"
)

type Validator interface {
	GetVersion() (*v.Version, error)
	GetPatrolVersion() (*v.Version, error)
}

type ValidatorRunParams struct {
	Runner     Validator
	CliVersion *v.Version
}

func Run(params ValidatorRunParams) error {
	runner := params.Runner

	print.StepIniciated("--- Getting Flutter Version ---")

	version, err := runner.GetVersion()
	if err != nil {
		print.Warning("❌ Failed to get Flutter version")
		print.Error(err.Error())
	}

	print.StepCompleted("✅ Flutter Version: " + version.String() + "\n")

	print.StepIniciated("--- Getting Patrol Version ---")
	patrolVersion, patrolErr := runner.GetPatrolVersion()

	if patrolErr != nil {
		print.Warning("❌ Failed to get Patrol version")
		print.Error(patrolErr.Error())
	}

	print.StepCompleted("✅ Patrol Version: " + patrolVersion.String())

	return nil
}
