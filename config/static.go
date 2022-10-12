package config

func init() {
	if BuildType == "development" {
		DevelopmentBuild = true
	} else if BuildType != "production" {
		panic("Invalid build type: " + BuildType)
	}
}

var (
	Version          = "0.0.0"
	BuildType        = "development"
	DevelopmentBuild bool
)
