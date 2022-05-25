package instance

import (
	"sync"

	"soccer-manager/config"
	"soccer-manager/db"
	"soccer-manager/logger"

	"github.com/airbrake/gobrake/v4"

	// pg driver for postgres
	_ "github.com/lib/pq"
	"github.com/minio/minio-go"
)

type instance struct {
	dbInstance  db.DBInstance
	minioClient *minio.Client
	airbrake    *gobrake.Notifier
}

var singleton = &instance{}
var once sync.Once

func Init() {
	postgresConfig := config.Postgres()
	once.Do(func() {

		minioClient, err := minio.New(config.ObjectStorageEndpoint(), config.ObjectStorageAccessKey(), config.ObjectStorageSecretKey(), config.ObjectStorageIsSecure())
		if err != nil {
			logger.Log.Fatal(err)
		}
		singleton.minioClient = minioClient

		logger.Log.Info("Connecting to postgres...")
		singleton.dbInstance = db.NewDBInstance(db.PgInstanceConfig{
			ConnURL:  postgresConfig.ConnURL(),
			BuildEnv: config.BuildEnv(),
		})

		logger.Log.Info("Connected to postgres successfully...")

		if !config.IsDevEnv() {
			singleton.airbrake = gobrake.NewNotifierWithOptions(&gobrake.NotifierOptions{
				ProjectId:  config.AirbrakeProjectID(),
				ProjectKey: config.AirbrakeProjectKey(),
			})
		}

	})
}

// DB returns the database object
func DB() db.DBInstance {
	return singleton.dbInstance
}

// MinioClient returns the MinioClient
func MinioClient() *minio.Client {
	return singleton.minioClient
}

// Airbrake returns the airbrake object in singleton
func Airbrake() *gobrake.Notifier {
	return singleton.airbrake
}

// Destroy closes the connections & cleans up the instance
func Destroy() error {
	if !config.IsDevEnv() {
		singleton.airbrake.Close()
	}
	return nil
}
