package get_flutter_version

import (
	"bufio"
	"fmt"
	"strings"

	v "github.com/Masterminds/semver/v3"

	"patrol_install/commands"
	"patrol_install/utils/exec"
)

var flutterPubDeps = commands.FlutterPubDependencies

func GetVersion() (*v.Version, error) {
	output, err := exec.Command(flutterPubDeps)
	if err != nil {
		return nil, err
	}

	version, err := getPatrolVersionFromLog(output)
	if err != nil {
		return nil, fmt.Errorf("could not find version in output")
	}

	return version, nil
}

func getPatrolVersionFromLog(log string) (*v.Version, error) {
	scanner := bufio.NewScanner(strings.NewReader(log))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(strings.TrimSpace(line), "- patrol ") {
			fields := strings.Fields(line)
			if len(fields) >= 3 {
				rawVersion := fields[2]
				version, err := v.NewVersion(rawVersion)
				if err != nil {
					return nil, fmt.Errorf("invalid version format: %v", err)
				}
				return version, nil
			}
		}
	}
	return nil, fmt.Errorf("patrol package not found")
}
