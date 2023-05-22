package config

import (
	"fmt"
	"log"
	"os"
)

type MonitorConfig struct {
	MonitorHost       string
	MonitorHostStrict bool
	MonitorTimeout    string
	NotifyHost        string
}

func Build() MonitorConfig {
	// TODO make sure monitor host is an IP or hostname
	monitorHost, err := GetEnvString("MONITOR_HOST")
	if err != nil {
		log.Fatalf("MONITOR_HOST: %v", err)
	}
	monitorHostStrict, err := GetEnvBoolDefault("MONITOR_STRICT", true)
	if err != nil {
		log.Printf("Using default value of %v for MONITOR_HOST_STRICT", monitorHostStrict)
	}

	monitorTimeout, err := GetEnvStringDefault("MONITOR_TIMEOUT", "5s")
	if err != nil {
		log.Printf("Using default value of %v for MONITOR_TIMEOUT", monitorTimeout)
	}

	notifyHost, err := GetEnvString("NOTIFY_HOST")
	if err != nil {
		log.Fatalf("NOTIFY_HOST: %v", err)
	}

	return MonitorConfig{
		MonitorHost:       monitorHost,
		MonitorHostStrict: monitorHostStrict,
		MonitorTimeout:    monitorTimeout,
		NotifyHost:        notifyHost,
	}
}

func GetEnvBoolDefault(name string, def bool) (bool, error) {
	env, set := os.LookupEnv(name)
	if !set {
		return def, fmt.Errorf("Environment variable %v is unset, returning default", name)
	}
	if env == "" {
		return def, fmt.Errorf("Environment variable %v is set but empty, returning default", name)
	}
	if env != "True" && env != "true" && env != "False" && env != "false" {
		return def, fmt.Errorf("Environment variable %v has invalid value %v, returning default", name, env)
	}
	if env == "True" || env == "true" {
		return true, nil
	} else {
		return false, nil
	}
}

func GetEnvString(name string) (string, error) {
	env, set := os.LookupEnv(name)
	if !set {
		return env, fmt.Errorf("Environment variable %v is unset", name)
	}
	if env == "" {
		return env, fmt.Errorf("Environment variable %v is set, but empty", name)
	}
	return env, nil
}

func GetEnvStringDefault(name string, def string) (string, error) {
	env, set := os.LookupEnv(name)
	if !set {
		return def, fmt.Errorf("Environment variable %v is unset, returning default", name)
	}
	if env == "" {
		return def, fmt.Errorf("Environment variable %v is set but empty, returning default", name)
	}
	return env, nil
}
