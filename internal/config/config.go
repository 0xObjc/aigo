package config

type Config struct {
	IncludeAllFiles bool
}

func NewConfig(args []string) Config {
	includeAllFiles := false
	if len(args) == 3 && args[2] == "-all" {
		includeAllFiles = true
	}
	return Config{
		IncludeAllFiles: includeAllFiles,
	}
}
