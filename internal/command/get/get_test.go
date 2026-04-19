package get

import (
	"bytes"
	"errors"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestGet_Success(t *testing.T) {
	originalExit := osExit
	defer func() { osExit = originalExit }()
	osExit = func(code int) {}

	runGet := func(key string) (string, error) {
		return "my-secret-value", nil
	}

	s := &getImpl{getCmd: &cobra.Command{Use: "get"}}
	s.attachFlags()

	var buf bytes.Buffer
	s.getCmd.SetOut(&buf)

	s.get(runGet, []string{"db/password"})

	assert.Contains(t, buf.String(), "my-secret-value")
}

func TestGet_JSON(t *testing.T) {
	originalExit := osExit
	defer func() { osExit = originalExit }()
	osExit = func(code int) {}

	runGet := func(key string) (string, error) {
		return "secret", nil
	}

	s := &getImpl{getCmd: &cobra.Command{Use: "get"}}
	s.attachFlags()
	s.getCmd.Flags().Set("output", "json")

	var buf bytes.Buffer
	s.getCmd.SetOut(&buf)

	s.get(runGet, []string{"db/password"})

	assert.Contains(t, buf.String(), `"key":"db/password"`)
	assert.Contains(t, buf.String(), `"value":"secret"`)
}

func TestGet_YAML(t *testing.T) {
	originalExit := osExit
	defer func() { osExit = originalExit }()
	osExit = func(code int) {}

	runGet := func(key string) (string, error) {
		return "secret", nil
	}

	s := &getImpl{getCmd: &cobra.Command{Use: "get"}}
	s.attachFlags()
	s.getCmd.Flags().Set("output", "yaml")

	var buf bytes.Buffer
	s.getCmd.SetOut(&buf)

	s.get(runGet, []string{"db/password"})

	assert.Contains(t, buf.String(), "key: db/password")
	assert.Contains(t, buf.String(), "value: secret")
}

func TestGet_Error(t *testing.T) {
	originalExit := osExit
	defer func() { osExit = originalExit }()
	var exitCode int
	osExit = func(code int) { exitCode = code }

	runGet := func(key string) (string, error) {
		return "", errors.New("vault sealed")
	}

	s := &getImpl{getCmd: &cobra.Command{Use: "get"}}
	s.attachFlags()

	var buf bytes.Buffer
	s.getCmd.SetOut(&buf)
	s.getCmd.SetErr(&buf)

	s.get(runGet, []string{"db/password"})

	assert.Equal(t, 1, exitCode)
}

func TestGet_NoArgs(t *testing.T) {
	originalExit := osExit
	defer func() { osExit = originalExit }()
	var exitCode int
	osExit = func(code int) { exitCode = code }

	runGet := func(key string) (string, error) {
		return "", nil
	}

	s := &getImpl{getCmd: &cobra.Command{Use: "get"}}
	s.attachFlags()

	var buf bytes.Buffer
	s.getCmd.SetOut(&buf)
	s.getCmd.SetErr(&buf)

	s.get(runGet, []string{})

	assert.Equal(t, 2, exitCode)
}

func TestNew(t *testing.T) {
	runGet := func(key string) (string, error) {
		return "value", nil
	}

	svc := New(runGet)

	assert.NotNil(t, svc)
	assert.NotNil(t, svc.GetCmd())
	assert.Equal(t, "get [key]", svc.GetCmd().Use)
}

func TestGetCmd(t *testing.T) {
	cmd := &cobra.Command{Use: "get"}
	s := &getImpl{getCmd: cmd}
	assert.Equal(t, cmd, s.GetCmd())
}
