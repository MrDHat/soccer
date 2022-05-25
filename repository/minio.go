package repository

import (
	"net/url"
	"path"
	"time"

	"soccer-manager/config"
	"soccer-manager/constants"
	"soccer-manager/utils"

	"soccer-manager/logger"

	"github.com/minio/minio-go"
)

// MinioRepo is the repo for the Minio library
type MinioRepo interface {
	Upload() (string, error)
	GetUrl(objectName *string) (*string, error)
	GetObjectName(url *string) (*string, error)
}

type minioRepo struct {
	minioClient *minio.Client
}

//Upload returns a signed url that can be used to upload a file
func (repo *minioRepo) Upload() (string, error) {
	groupError := "UPLOAD_FILE"

	logger.Log.Info("Creating uuid for the file name")
	fileName, err := utils.GetRandomUUID()
	if err != nil {
		logger.Log.WithError(err).Error(constants.InternalServerError)
		return "", err
	}
	logger.Log.Info("generated uuid for the file name")

	u, err := repo.minioClient.PresignedPutObject(config.ObjectStorageBucketName(),
		fileName,
		time.Second*time.Duration(config.ObjectStorageSignedURLExpirySeconds()))

	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return "", err
	}

	return u.String(), nil
}

//GetUrl returns a signed url that can be used to get a file
func (repo *minioRepo) GetUrl(objectName *string) (*string, error) {
	if objectName == nil {
		return nil, nil
	}
	groupError := "GET_URL_FOR_FILE"
	var signedUrl *string

	logger.Log.Info("Getting URL for the file")
	reqParams := make(url.Values)

	u, err := repo.minioClient.PresignedGetObject(config.ObjectStorageBucketName(),
		*objectName,
		time.Second*time.Duration(config.ObjectStorageSignedURLExpirySeconds()),
		reqParams)

	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return signedUrl, err
	}
	logger.Log.Info("Done getting url for the file")
	sUrl := u.String()
	signedUrl = &sUrl
	return signedUrl, nil
}

//GetObjectName parses the object name from the signed url
func (*minioRepo) GetObjectName(s *string) (*string, error) {
	if s == nil {
		return nil, nil
	}
	groupError := "GET_OBJECT_NAME"
	var oName *string

	logger.Log.Info("Getting the object name from the PUT URL")
	u, err := url.Parse(*s)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return oName, err
	}
	logger.Log.Info("Done getting the object name from the PUT URL")
	base := path.Base(u.Path)
	oName = &base
	return oName, nil
}

// NewMinioRepo creates a new instance of the minio repo
func NewMinioRepo(minioClient *minio.Client) MinioRepo {
	return &minioRepo{
		minioClient: minioClient,
	}
}
