package utils

import (
	"fmt"
	"strings"
)

const (
	cnfHostKey          = "server.host"
	cnfPortKey          = "server.port"
	cnfProxyProtocolKey = "server.http.proxy_protocol"
	cnfProxyHostKey     = "server.http.proxy_host"
	cnfProxyPortKey     = "server.http.proxy_port"

	ConfigLoadedSuccessMsg    = "Configuration was loaded successfully from %s"
	ConfigChangeDetectedMsg   = "A configuration change was detected in the config file '%s'"
	ConfigUpdateSuccessMsg    = "Configuration was updated"
	ConfigUpdateFailedMsg     = "Configuration was not updated because of a validation failure"
	ConfigValidationFailedMsg = "Config validation failed"
	readConfigFileErrMsg      = "Unable to read configuration from config file '%s'. Error: %v"
)

// ReadConfigFileError represents an error when the config file cannot be read
type ReadConfigFileError struct {
	File string
	Err  error
}

// Error returns the formatted ReadConfigFileError
func (rcf ReadConfigFileError) Error() string {
	return fmt.Sprintf(readConfigFileErrMsg, rcf.File, rcf.Err)
}

// LoggerCnf represents the Logger settings
type LoggerCnf struct {
	Level string `json:"level" yaml:"level" mapstructure:"level"`
}

// LogLevel is the global variable to set the log level. Default value is INFO log level
var LogLevel = infoLogLevel

// Validate checks if the values in the LoggerCnf struct are valid
// The method returns an error if the configuration is not valid
func (l *LoggerCnf) Validate() error {
	if !EntryExists(validLogLevels, l.Level) {
		return InvalidLogLevelError(l.Level)
	}
	return nil
}

// Set sets the log level
func (l *LoggerCnf) Set() {
	LogLevel = l.Level
}

// HTTPCnf represents the servers http configuration
type HTTPCnf struct {
	SkipTLS       bool   `yaml:"skip_tls" mapstructure:"skip_tls"`
	ProxyEnable   bool   `yaml:"proxy_enable" mapstructure:"proxy_enable"`
	ProxyProtocol string `yaml:"proxy_protocol" mapstructure:"proxy_protocol"`
	ProxyHost     string `yaml:"proxy_host" mapstructure:"proxy_host"`
	ProxyPort     string `yaml:"proxy_port" mapstructure:"proxy_port"`
}

// Validate checks if the values in the HTTPCnf are valid
func (hc *HTTPCnf) Validate() error {

	var missingParams []string

	if hc.ProxyEnable {
		if strings.TrimSpace(hc.ProxyProtocol) == "" {
			missingParams = append(missingParams, cnfProxyProtocolKey)
		}
		if strings.TrimSpace(hc.ProxyHost) == "" {
			missingParams = append(missingParams, cnfProxyHostKey)
		}
		if strings.TrimSpace(hc.ProxyPort) == "" {
			missingParams = append(missingParams, cnfProxyPortKey)
		}

		if len(missingParams) != 0 {
			return MissingMandatoryParamError(missingParams)
		}

		if !EntryExists(validProtocols, hc.ProxyProtocol) {
			return InvalidProxyProtocolError(hc.ProxyProtocol)
		}
	}

	return nil
}

// GetProxyUrl returns the formatted proxy URL
func (hc *HTTPCnf) GetProxyUrl() string {
	return fmt.Sprintf("%s://%s:%s", hc.ProxyProtocol, hc.ProxyHost, hc.ProxyPort)
}

// Set sets the http config
func (hc *HTTPCnf) Set() {
	SkipTLS = hc.SkipTLS
	ProxyEnabled = hc.ProxyEnable
	ProxyProtocol = hc.ProxyProtocol
	ProxyHost = hc.ProxyHost
	ProxyPort = hc.ProxyPort
}

// ServerCnf represents the server configuration
// It includes the basic host + port config along with the Logger and HTTP configurations
type ServerCnf struct {
	Host      string `yaml:"host" mapstructure:"host"`
	Port      string `yaml:"port" mapstructure:"port"`
	LoggerCnf `yaml:"logging" mapstructure:"logging"`
	HTTPCnf   `yaml:"http" mapstructure:"http"`
}

// Validate validates if the sonar configuration provided in the configuration file is valid
// The method returns an error if any configuration defined in the receiver is not valid
func (sc *ServerCnf) Validate() error {
	var missingParams []string

	if strings.TrimSpace(sc.Host) == "" {
		missingParams = append(missingParams, cnfHostKey)
	}
	if strings.TrimSpace(sc.Port) == "" {
		missingParams = append(missingParams, cnfPortKey)
	}

	if len(missingParams) != 0 {
		return MissingMandatoryParamError(missingParams)
	}

	if err := sc.LoggerCnf.Validate(); err != nil {
		return err
	}

	if err := sc.HTTPCnf.Validate(); err != nil {
		return err
	}

	return nil
}
