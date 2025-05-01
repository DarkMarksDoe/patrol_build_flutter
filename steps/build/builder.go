package builder

import (
	"fmt"

	"patrol_install/utils/print"
)

type Builder interface {
	BuildParametersFromEnv() (string, error)
}

func Run(installer Builder) error {
	print.StepIniciated("--- Building Command ---")

	// Final build
	_, err := installer.BuildParametersFromEnv()
	if err != nil {
		print.Error(fmt.Sprintf("Build failed: %s", err))
		return err
	}
	print.StepCompleted("Build completed successfully")
	return nil
}
