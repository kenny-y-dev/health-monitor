package config

import (
	"os"
	"testing"
)

func TestGetEnvBoolDefault(t *testing.T) {
	os.Unsetenv("TESTBOOL")
	defer os.Unsetenv("TESTBOOL")

	var want bool
	var env bool
	var err error

	// Unset
	want = true
	env, err = GetEnvBoolDefault("TESTBOOL", true)
	if err == nil {
		t.Errorf("Empty error not thrown")
	}
	if env != want {
		t.Errorf("Expected %v, got %v", want, env)
	}

	// Set but empty
	os.Setenv("TESTBOOL", "")
	want = true
	env, err = GetEnvBoolDefault("TESTBOOL", true)
	if err == nil {
		t.Errorf("Empty error not thrown")
	}
	if env != want {
		t.Errorf("Expected %v, got %v", want, env)
	}

	// Set but invalid
	os.Setenv("TESTBOOL", "testing123")
	want = true
	env, err = GetEnvBoolDefault("TESTBOOL", true)
	if err == nil {
		t.Errorf("Empty error not thrown")
	}
	if env != want {
		t.Errorf("Expected %v, got %v", want, env)
	}

	// Set with True
	os.Setenv("TESTBOOL", "True")
	want = true
	env, err = GetEnvBoolDefault("TESTBOOL", false)
	if err != nil {
		t.Errorf("Error unexpectatntly thrown: %v", err)
	}
	if env != want {
		t.Errorf("Expected %v, got %v", want, env)
	}

	// Set with False
	os.Setenv("TESTBOOL", "False")
	want = false
	env, err = GetEnvBoolDefault("TESTBOOL", true)
	if err != nil {
		t.Errorf("Error unexpectatntly thrown: %v", err)
	}
	if env != want {
		t.Errorf("Expected %v, got %v", want, env)
	}

}

func TestGetEnvString(t *testing.T) {
	os.Unsetenv("TESTSTR")
	defer os.Unsetenv("TESTSTR")
	var env string
	var err error
	var want string

	// Unset
	want = ""
	env, err = GetEnvString("TESTSTR")
	if err == nil {
		t.Errorf("Unset error not thrown")
	}
	if want != env {
		t.Errorf("Expected %v, got %v", want, env)
	}

	// Empty
	os.Setenv("TESTSTR", "")
	want = ""
	env, err = GetEnvString("TESTSTR")
	if err == nil {
		t.Errorf("Empty error not thrown")
	}
	if want != env {
		t.Errorf("Expected %v, got %v", want, env)
	}

	// Set
	os.Setenv("TESTSTR", "Get")
	want = "Get"
	env, err = GetEnvString("TESTSTR")
	if err != nil {
		t.Errorf("Error unexpectatntly thrown: %v", err)
	}
	if want != env {
		t.Errorf("Expected %v, got %v", want, env)
	}

}
