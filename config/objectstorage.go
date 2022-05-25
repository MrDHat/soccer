package config

import (
	"soccer-manager/logger"

	"github.com/spf13/viper"
)

type objectstorage struct {
	accessKey              string
	secretKey              string
	endpoint               string
	isSecure               bool
	baseURL                string
	bucketName             string
	bucketLocation         string
	signedURLExpirySeconds int64
}

func (os *objectstorage) load() {
	logger.Log.Info("Reading object storage config...")
	viper.SetEnvPrefix("obj_storage")
	viper.AutomaticEnv()

	os.accessKey = viper.GetString("access_key")
	os.secretKey = viper.GetString("secret_key")
	os.endpoint = viper.GetString("endpoint")
	os.isSecure = viper.GetBool("secure")
	os.bucketName = viper.GetString("bucket_name")
	os.bucketLocation = viper.GetString("bucket_location")
	os.signedURLExpirySeconds = viper.GetInt64("signed_url_expiry_seconds")
	os.baseURL = viper.GetString("base_url")
}
