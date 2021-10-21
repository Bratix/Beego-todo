package global

import (
	"log"

	"github.com/spf13/viper"
)

func EnviromentVariable(key string) string {

	/* Viper function to read /conf/auth.evn data */

	// name of config file (without extension)
	viper.SetConfigName("auth")
	// look for config in the working directory
	viper.AddConfigPath("./conf")

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	// viper.Get() returns an empty interface{}
	// to get the underlying type of the key,
	// we have to do the type assertion, we know the underlying value is string
	// if we type assert to other type it will throw an error
	value, ok := viper.Get(key).(string)

	// If the type is a string then ok will be true
	// ok will make sure the program not break
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}
