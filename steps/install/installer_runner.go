package install

import (
	v "github.com/Masterminds/semver/v3"

	cli_info "patrol_install/steps/install/get_cli_version"
	cli_install "patrol_install/steps/install/install_cli_tool"
)

type InstallerRunner struct{}

func (p *InstallerRunner) GetVersion() (*v.Version, error) {
	return cli_info.GetCLIVersion()
}

func (p *InstallerRunner) Install() error {
	_, err := cli_install.Install()
	return err
}
