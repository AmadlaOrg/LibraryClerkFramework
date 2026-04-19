package doorman

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"doorman-test", "version"}

	assert.NotPanics(t, func() {
		New("doorman-test", "Doorman Test", "1.0.0",
			[]string{"amadla.org/entity/infrastructure@^v1.0.0"},
			"Test doorman plugin",
			func(key string) (string, error) {
				return "value", nil
			})
	})
}

func TestNew_Info(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"doorman-test", "info"}

	assert.NotPanics(t, func() {
		New("doorman-test", "Doorman Test", "1.0.0",
			[]string{"amadla.org/entity/infrastructure@^v1.0.0"},
			"Test doorman plugin",
			func(key string) (string, error) {
				return "value", nil
			})
	})
}
