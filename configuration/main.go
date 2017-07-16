package configuration

import (
	"fmt"

	"github.com/spf13/viper"
)

// ----------------------------------------------------------------------------
// Interface object for viper.BindFlagValue(...)
// ----------------------------------------------------------------------------

type myFlag struct {
	MyName string
	Value  string
	Type   string
}

func (f myFlag) HasChanged() bool    { return false }
func (f myFlag) Name() string        { return f.MyName }
func (f myFlag) ValueString() string { return f.Value }
func (f myFlag) ValueType() string   { return f.Type }

// ----------------------------------------------------------------------------
// Safe argument extraction.
// ----------------------------------------------------------------------------

func getFlag(args map[string]interface{}, flagName string) string {
	result := ""
	flag := "--" + flagName
	resultRaw := args[flag]
	if resultRaw != nil {
		result = resultRaw.(string)
	}
	return result
}

// ----------------------------------------------------------------------------
// Load configuration.
// ----------------------------------------------------------------------------

func LoadConfig(args map[string]interface{}) {

	// Keys

	setKey := "setKey"
	setKeyOnly := "setKeyOnly"
	flagKey := "flagKey"
	flagKeyOnly := "flagKeyOnly"
	osKey := "osKey"
	osKeyOnly := "osKeyOnly"
	configurationKey := "configurationKey"
	defaultKey := "defaultKey"

	// ------------------------------------------------------------------------
	// Load defaults
	// ------------------------------------------------------------------------

	viper.SetDefault(defaultKey, "From SetDefault()")
	viper.SetDefault(configurationKey, "From SetDefault()")
	viper.SetDefault(osKey, "From SetDefault()")
	viper.SetDefault(flagKey, "From SetDefault()")
	viper.SetDefault(setKey, "From SetDefault()")

	// ------------------------------------------------------------------------
	// Load configuration file
	// ------------------------------------------------------------------------

	// Set paths where the configuration may be found.

	viper.SetConfigName("go-hello-viper")        // name of config file (without extension)
	viper.AddConfigPath("/etc/go-hello-viper/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.go-hello-viper") // call multiple times to add many search paths
	viper.AddConfigPath("$HOME/go/src/github.com/docktermj/go-hello-viper/")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// ------------------------------------------------------------------------
	// Load OS Environment Variables
	// ------------------------------------------------------------------------

	viper.BindEnv(osKey, "GOHELLOVIPER_OS_KEY")
	viper.BindEnv(osKeyOnly, "GOHELLOVIPER_OS_KEY_ONLY")
	viper.BindEnv(flagKey, "GOHELLOVIPER_FLAG_KEY")
	viper.BindEnv(setKey, "GOHELLOVIPER_SET_KEY")

	// ------------------------------------------------------------------------
	// Load commandline options
	// Note: this currently does not work.
	// Reported in https://github.com/spf13/viper/issues/369
	// ------------------------------------------------------------------------

	viper.BindFlagValue(flagKey, myFlag{MyName: flagKey, Value: getFlag(args, flagKey), Type: "string"})
	viper.BindFlagValue(flagKeyOnly, myFlag{MyName: flagKeyOnly, Value: getFlag(args, flagKeyOnly), Type: "string"})
	viper.BindFlagValue(setKey, myFlag{MyName: setKey, Value: getFlag(args, setKey), Type: "string"})

	// ------------------------------------------------------------------------
	// Set explicitly.
	// ------------------------------------------------------------------------

	viper.Set(setKey, "From Set()")
	viper.Set(setKeyOnly, "From Set()")

}
