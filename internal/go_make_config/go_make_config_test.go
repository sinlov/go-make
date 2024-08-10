package go_make_config

import (
	"github.com/sebdah/goldie/v2"
	"gopkg.in/yaml.v3"
	"testing"
)

var (
	testGlobalEnv = map[string]string{
		"GONOPROXY":   "*.gitlab.com,*.gitee.com",
		"GONOSUMDB":   "*.gitlab.com,*.gitee.com",
		"GOPRIVATE":   "*.gitlab.com,*.gitee.com",
		"GO111MODULE": "on",
	}
)

func TestInitDefaultConfig(t *testing.T) {

	// basic
	basicConfig := InitDefaultConfig()
	basicConfig.GlobalEnv = testGlobalEnv

	// env
	envConfig := InitDefaultConfig()
	envConfig.GlobalEnv = testGlobalEnv
	envConfig.Steps = append(basicConfig.Steps, Step{
		Name:    "env",
		Omitted: true,
		Commands: []string{
			"go env",
			"go version",
		},
	})

	// dep
	depConfig := InitDefaultConfig()
	depConfig.GlobalEnv = testGlobalEnv
	depConfig.Steps = append(basicConfig.Steps, Step{
		Name:    "env",
		Omitted: true,
		Commands: []string{
			"go env",
			"go version",
		},
	})
	depConfig.Steps = append(depConfig.Steps, Step{
		Name: "dep",
		Commands: []string{
			"go mod verify",
			"go mod tidy -v",
			"go mod download -x",
		},
	})

	// mock InitDefaultConfig
	type args struct {
		cfg GoMakeConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "basic", // testdata/TestInitDefaultConfig/basic.golden
			args: args{
				cfg: *basicConfig,
			},
		},
		{
			name: "env", // testdata/TestInitDefaultConfig/env.golden
			args: args{
				cfg: *envConfig,
			},
		},
		{
			name: "dep", // testdata/TestInitDefaultConfig/dep.golden
			args: args{
				cfg: *depConfig,
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)

			// do Marshal GoMakeConfig
			out, errYaml := yaml.Marshal(&tc.args.cfg)
			if errYaml != nil {
				t.Fatal(errYaml)
			}
			// verify GoMakeConfig
			g.Assert(t, t.Name(), out)
		})
	}
}
