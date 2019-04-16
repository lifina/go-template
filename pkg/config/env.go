package config

import (
	"github.com/spf13/viper"
)

const key = "ENV"

func IsEnvDeveloping() bool {
	return viper.GetString("ENV") == "develop"
}

func IsEnvStaging() bool {
	return viper.GetString("ENV") == "staging"
}

func IsEnvProduction() bool {
	return viper.GetString("ENV") == "production"
}
