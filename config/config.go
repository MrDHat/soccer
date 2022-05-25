package config

import "strings"

type config struct {
	objectstorageConfig objectstorage
	postgresConfig      postgres
	apiConfig           api
}

var configuration = &config{}

// Load loads the config into the configuration object
func Load() {
	configuration.apiConfig.load()
	configuration.objectstorageConfig.load()
	configuration.postgresConfig.load()
}

// Postgres returns the postgres config
func Postgres() postgres {
	return configuration.postgresConfig
}

// ObjectStorageAccessKey returns the ObjectStorageAccessKey config
func ObjectStorageAccessKey() string {
	return configuration.objectstorageConfig.accessKey
}

// ObjectStorageSecretKey returns the ObjectStorageSecretKey config
func ObjectStorageSecretKey() string {
	return configuration.objectstorageConfig.secretKey
}

// ObjectStorageEndpoint returns the ObjectStorageEndpoint config
func ObjectStorageEndpoint() string {
	return configuration.objectstorageConfig.endpoint
}

// ObjectStorageIsSecure returns the ObjectStorageIsSecure config
func ObjectStorageIsSecure() bool {
	return configuration.objectstorageConfig.isSecure
}

// ObjectStorageBucketName returns the Minio bucket name for storing files
func ObjectStorageBucketName() string {
	return configuration.objectstorageConfig.bucketName
}

// ObjectStorageBucketLocation returns the Minio bucket location for storing files
func ObjectStorageBucketLocation() string {
	return configuration.objectstorageConfig.bucketLocation
}

//ObjectStorageSignedURLExpirySeconds returns the duration in seconds for which the signed url will be active.
func ObjectStorageSignedURLExpirySeconds() int64 {
	return configuration.objectstorageConfig.signedURLExpirySeconds
}

// APIPort returns the api port
func APIPort() string {
	return configuration.apiConfig.port
}

// JWTSigningKey returns the JWT signing key
func JWTSigningKey() string {
	return configuration.apiConfig.jwtKey
}

// BuildEnv returns the build env
func BuildEnv() string {
	return configuration.apiConfig.buildEnv
}

// AirbrakeProjectID returns the airbrake project id
func AirbrakeProjectID() int64 {
	return configuration.apiConfig.airbrakeProjectID
}

// AirbrakeProjectKey returns the airbrake project key
func AirbrakeProjectKey() string {
	return configuration.apiConfig.airbrakeProjectKey
}

func ServiceName() string {
	return configuration.apiConfig.serviceName
}

func ServiceVersion() string {
	return configuration.apiConfig.serviceVersion
}

func IsDevEnv() bool {
	return strings.Contains(strings.ToLower(configuration.apiConfig.buildEnv), "dev")
}

func IsStagingEnv() bool {
	return strings.Contains(strings.ToLower(configuration.apiConfig.buildEnv), "staging")
}

func IsDevLikeEnv() bool {
	return IsDevEnv() || IsStagingEnv()
}

func IsProdEnv() bool {
	return strings.Contains(strings.ToLower(configuration.apiConfig.buildEnv), "prod")
}
