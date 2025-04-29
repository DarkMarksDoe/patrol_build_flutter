package install

import (
	"github.com/Masterminds/semver/v3"

	cli_info "patrol_install/steps/install/get_cli_version"
	cli_install "patrol_install/steps/install/install_cli_tool"
)

type PatrolInstaller struct{}

func (p *PatrolInstaller) GetVersion() (*semver.Version, error) {
	return cli_info.GetCLIVersion()
}

func (p *PatrolInstaller) Install() error {
	_, err := cli_install.Install()
	return err
}
