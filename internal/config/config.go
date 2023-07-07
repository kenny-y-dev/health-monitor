package config

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type MonitorConfig struct {
	// TODO look at turning monitor target and notify into slices for multiple targets/notices
	// TODO multiple targets/notices may be good use for async; though ping may block
	MonitorTarget      string
	MonitorCheckStrict bool
	MonitorTimeout     string
	NotifyTarget       string
	NotifyMethod       string
	NotifyDownJSON     []byte
	NotifyUpJSON       []byte
}

func Build() MonitorConfig {
	// TODO make sure monitor host is an IP or hostname
	monitorTarget, err := GetEnvString("MONITOR_TARGET")
	if err != nil {
		log.Fatalf("MONITOR_TARGET: %v", err)
	}
	monitorCheckStrict, err := GetEnvBoolDefault("MONITOR_STRICT", true)
	if err != nil {
		log.Printf("Using default value of %v for MONITOR_HOST_STRICT", monitorCheckStrict)
	}

	monitorTimeout, err := GetEnvStringDefault("MONITOR_TIMEOUT", "5s")
	if err != nil {
		log.Printf("Using default value of %v for MONITOR_TIMEOUT", monitorTimeout)
	}

	notifyTarget, err := GetEnvString("NOTIFY_TARGET")
	if err != nil {
		log.Fatalf("NOTIFY_TARGET: %v", err)
	}

	notifyMethod, err := GetEnvString("NOTIFY_METHOD")
	if err != nil {
		log.Fatalf("NOTIFY_METHOD: %v", err)
	}
	notifyMethod = strings.ToUpper(notifyMethod)
	if !ValidateNotifyMethod(notifyMethod) {
		log.Fatalf("NOTIFY_METHOD not set to valid or implemented HTTP method")
	}
	notifyDownJSONStr, err := GetEnvString("NOTIFY_DOWN_JSON")
	if err != nil {
		log.Fatalf("NOTIFY_DOWN_JSON is not valid")
	}
	notifyDownJSON := []byte(notifyDownJSONStr)

	notifyUpJSONStr, err := GetEnvString("NOTIFY_UP_JSON")
	if err != nil {
		log.Fatalf("NOTIFY_UP_JSON is not valid")
	}
	notifyUpJSON := []byte(notifyUpJSONStr)

	return MonitorConfig{
		MonitorTarget:      monitorTarget,
		MonitorCheckStrict: monitorCheckStrict,
		MonitorTimeout:     monitorTimeout,
		NotifyTarget:       notifyTarget,
		NotifyMethod:       notifyMethod,
		NotifyDownJSON:     notifyDownJSON,
		NotifyUpJSON:       notifyUpJSON,
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
	env = strings.ToLower(env)
	if env != "true" && env != "false" {
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

func ValidateNotifyMethod(value string) bool {
	methods := map[string]bool{"GET": true, "POST": true, "PUT": true}
	return CheckValidValue(value, methods)
}

func CheckValidValue(value string, valid map[string]bool) bool {
	_, found := valid[value]
	if found {
		return true
	} else {
		return false
	}
}

func (mc MonitorConfig) PrintConfig() {
	log.Printf("Starting monitoring with the following config:")
	log.Printf("Monitor target: %v", mc.MonitorTarget)
	log.Printf("Strict monitor checking: %v", mc.MonitorCheckStrict)
	log.Printf("Monitor timeout: %v", mc.MonitorTimeout)
	log.Printf("Notify target: %v", mc.NotifyTarget)
	log.Printf("Notify method: %v", mc.NotifyMethod)
}
