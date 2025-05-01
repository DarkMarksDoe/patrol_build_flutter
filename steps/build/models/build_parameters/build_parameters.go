package build_parameters

import (
	"fmt"
	"strings"
)

// BuildParameters holds validated and formatted build configuration.
type BuildParameters struct {
	Targets      string
	Platform     string
	BuildType    string
	Tags         string
	ExcludedTags string
	IsVerbose    string
	IsCoverage   string
}

// NewBuildParameters builds a BuildParameters struct from a map of environment variables.
func NewBuildParameters(envMap map[string]string) (*BuildParameters, error) {
	bp := &BuildParameters{}

	requiredFields := map[string]func(*BuildParameters, string) error{
		"platform":  SetPlatform,
		"targets":   SetTargets,
		"buildType": SetBuildType,
	}

	optionalFields := map[string]func(*BuildParameters, string) error{
		"tags":         SetTags,
		"excludedTags": SetExcludedTags,
		"verbose":      SetVerbose,
		"isCoverage":   SetCoverage,
	}

	// Apply required setters
	for key, setter := range requiredFields {
		val, ok := envMap[key]
		if !ok || strings.TrimSpace(val) == "" {
			return nil, fmt.Errorf("missing required field: %s", key)
		}
		if err := setter(bp, val); err != nil {
			return nil, err
		}
	}

	// Apply optional setters
	for key, setter := range optionalFields {
		if val, ok := envMap[key]; ok && strings.TrimSpace(val) != "" {
			if err := setter(bp, val); err != nil {
				return nil, err
			}
		}
	}

	return bp, nil
}

// Command constructs the final CLI command string based on the populated BuildParameters fields.
func (bp *BuildParameters) Command() string {
	var args []string

	args = append(args, "--platform", bp.Platform)
	args = append(args, "--targets", bp.Targets)
	args = append(args, "--buildType", bp.BuildType)

	if bp.Tags != "" {
		args = append(args, "--tags", bp.Tags)
	}
	if bp.ExcludedTags != "" {
		args = append(args, "--excludedTags", bp.ExcludedTags)
	}
	if bp.IsVerbose != "" {
		args = append(args, bp.IsVerbose)
	}
	if bp.IsCoverage != "" {
		args = append(args, bp.IsCoverage)
	}

	return "patrol build " + strings.Join(args, " ")
}
