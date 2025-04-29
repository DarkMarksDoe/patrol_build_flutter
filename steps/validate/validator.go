package validate

import (
	v "github.com/Masterminds/semver/v3"

	"patrol_install/utils/print"
)

type Validator interface {
	GetVersion() (*v.Version, error)
}

type ValidatorRunParams struct {
	Runner     Validator
	CliVersion *v.Version
}

func Run(params ValidatorRunParams) error {
	runner := params.Runner
	cliVersion := params.CliVersion

	print.StepIniciated("--- Getting Flutter Version ---\n")

	version, err := runner.GetVersion()
	if err != nil {
		print.Warning("❌ Failed to get Flutter version")
		print.Error(err.Error())
	}

	print.StepCompleted("✅ Flutter Version: " + version.String())
	return nil
}
