package etc

import (
	"fmt"

	"github.com/spf13/viper"
)

func TrySomething() {

	// Keys

	setKey := "setKey"
	setKeyOnly := "setKeyOnly"
	flagKey := "flagKey"
	flagKeyOnly := "flagKeyOnly"
	osKey := "osKey"
	osKeyOnly := "osKeyOnly"
	configurationKey := "configurationKey"
	configurationKeyOnly := "configurationKeyOnly"
	defaultKey := "defaultKey"

	// Print from different configuration sources.

	fmt.Println("               set:", viper.GetString(setKey))
	fmt.Println("          set only:", viper.GetString(setKeyOnly))
	fmt.Println("              flag:", viper.GetString(flagKey))
	fmt.Println("         flag only:", viper.GetString(flagKeyOnly))
	fmt.Println("                os:", viper.GetString(osKey))
	fmt.Println("           os only:", viper.GetString(osKeyOnly))
	fmt.Println("     configuration:", viper.GetString(configurationKey))
	fmt.Println("configuration only:", viper.GetString(configurationKeyOnly))
	fmt.Println("           default:", viper.GetString(defaultKey))
}
