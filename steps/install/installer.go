package install

import (
	"github.com/Masterminds/semver/v3"

	"patrol_install/utils/print"
)

type Installer interface {
	GetVersion() (*semver.Version, error)
	Install() error
}

func Run(installer Installer) error {
	print.StepIniciated("--- Checking if Patrol CLI is already installed ---\n")

	version, err := installer.GetVersion()
	if err != nil {
		print.Warning("CLI is not installed, attempting installation...")
		if err := installer.Install(); err != nil {
			print.Error("❌ Installation failed: " + err.Error())
			return err
		}

		version, err = installer.GetVersion()
		if err != nil {
			print.Error("❌ Failed to verify version after install: " + err.Error())
			return err
		}

		print.StepCompleted("✅ PATROL CLI installed successfully. Version: " + version.String())
		return nil
	}

	print.StepCompleted("✅ Tool already installed. Version: " + version.String())
	return nil
}
