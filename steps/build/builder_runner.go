package builder

import (
	"fmt"

	getEnv "patrol_install/steps/build/steps/create_parameters"
	"patrol_install/utils/print"
)

type BuilderRunner struct{}

func (p *BuilderRunner) BuildParametersFromEnv() (string, error) {
	command, err := getEnv.BuildParametersFromEnv()
	if err != nil {
		print.Error(fmt.Sprintf("Build failed: %s", err))
		return "", err
	}
	return command.Command(), nil
}
