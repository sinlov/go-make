package go_make_config

const (
	FileName = "go-make.yml"
)

type Step struct {
	Name        string            `yaml:"name"`
	Omitted     bool              `yaml:"omitted,omitempty"`
	Environment map[string]string `yaml:"environment,omitempty"`
	Commands    []string          `yaml:"commands,omitempty"`
}

type GoMakeConfig struct {
	DisAble bool `yaml:"disable,omitempty"`

	MonoRepoPath []string `yaml:"mono_repo_path,omitempty"`

	GlobalEnv map[string]string `yaml:"global_env,omitempty"`

	Steps []Step `yaml:"steps,omitempty"`
}

func InitDefaultConfig() *GoMakeConfig {
	g := &GoMakeConfig{
		DisAble: false,
	}

	g = checkGoMakeConfig(g)

	return g
}

func checkGoMakeConfig(g *GoMakeConfig) *GoMakeConfig {
	if len(g.MonoRepoPath) == 0 {
		g.MonoRepoPath = []string{
			"",
		}
	}

	return g
}
