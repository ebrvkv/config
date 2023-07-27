package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type cnf struct {
	Property1 string `env:"PROPERTY1" yaml:"property1" env-required:""`
}

type testConfigStruct[T any] struct {
	cnfStruct     T
	envVariables  map[string]string
	cnfPath       string
	err           string
	checkProperty string
	description   string
}

func TestGet(t *testing.T) {
	tests := []testConfigStruct[any]{
		{&cnf{}, map[string]string{}, "", "value is not provided",
			"", "check if value is required but not provided",
		},
		{
			&cnf{}, map[string]string{
				"PROPERTY1": "localhost",
			}, "", "", "localhost", "check getting data from env vars",
		},
		{&cnf{}, nil,
			"configs/config.yml", "",
			"test123", "check env variable overwriting",
		},
		{&cnf{}, map[string]string{
			"PROPERTY1": "localhost123",
		}, "configs/config.yml", "",
			"localhost123", "check env variable overwriting",
		},
	}

	for _, test := range tests {
		var err error
		if test.envVariables != nil {
			for k, v := range test.envVariables {
				if err := os.Setenv(k, v); err != nil {
					t.Error(err)
				}
			}
		} else {
			os.Clearenv()
		}
		if test.cnfPath != "" {
			err = Get(test.cnfStruct, test.cnfPath)
		} else {
			err = Get(test.cnfStruct)
		}
		if test.err == "" {
			assert.NoError(t, err, test.description)
		} else {
			assert.ErrorContains(t, err, test.err, test.description)
		}
		if test.checkProperty != "" {
			assert.EqualValues(
				t, test.checkProperty,
				test.cnfStruct.(*cnf).Property1, test.description,
			)
		}
	}
}
