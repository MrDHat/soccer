package config

import (
	"soccer-manager/logger"

	"github.com/spf13/viper"
)

// api holds the config for the API
type api struct {
	buildEnv           string
	port               string
	jwtKey             string
	jwtExpirySeconds   int64
	airbrakeProjectID  int64
	airbrakeProjectKey string
	serviceName        string
	serviceVersion     string
}

// load returns the config for the API
func (apiConfig *api) load() {
	logger.Log.Info("Reading API config...")
	viper.SetEnvPrefix("api")
	viper.AutomaticEnv()

	apiConfig.buildEnv = viper.GetString("build_env")
	apiConfig.port = viper.GetString("port")
	apiConfig.jwtKey = viper.GetString("jwt_signing_key")
	apiConfig.jwtExpirySeconds = viper.GetInt64("jwt_expiry_seconds")

	apiConfig.airbrakeProjectID = viper.GetInt64("airbrake_project_id")
	apiConfig.airbrakeProjectKey = viper.GetString("airbrake_project_key")

	apiConfig.serviceName = viper.GetString("service_name")
	apiConfig.serviceVersion = viper.GetString("service_version")
}
