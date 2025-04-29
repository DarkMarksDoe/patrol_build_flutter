package validate

import (
	v "github.com/Masterminds/semver/v3"

	flutter "patrol_install/steps/validate/get_flutter_version"
	patrol "patrol_install/steps/validate/get_patrol_version"
)

type ValidatorRunner struct{}

func (p *ValidatorRunner) GetVersion() (*v.Version, error) {
	return flutter.GetVersion()
}

func (p *ValidatorRunner) GetPatrolVersion() (*v.Version, error) {
	return patrol.GetVersion()
}
