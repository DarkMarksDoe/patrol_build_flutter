package validate

import (
	v "github.com/Masterminds/semver/v3"

	flutter "patrol_install/steps/validate/get_flutter_version"
)

type ValidatorRunner struct{}

func (p *ValidatorRunner) GetVersion() (*v.Version, error) {
	return flutter.GetVersion()
}
