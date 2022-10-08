package config

func init() {
	if BuildType == "development" {
		DevelopmentBuild = true
	} else if BuildType == "production" {
		ProductionBuild = true
	} else {
		panic("Invalid build type: " + BuildType)
	}
}

var (
	Version          = "0.0.0"
	BuildType        = "development"
	DevelopmentBuild bool
	ProductionBuild  bool
)
