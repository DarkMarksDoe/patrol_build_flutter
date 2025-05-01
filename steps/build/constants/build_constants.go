package build_constants

const (
	Platform     = "BUILD_PLATFORM"
	Targets      = "BUILD_TARGETS" // comma-separated
	BuildType    = "BUILD_TYPE"
	Tags         = "BUILD_TAGS"          // optional, comma-separated
	ExcludedTags = "BUILD_EXCLUDED_TAGS" // optional, comma-separated
	Verbose      = "BUILD_VERBOSE"
	IsCovered    = "IS_COVERED"
)
