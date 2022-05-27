package utils

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"soccer-manager/constants"
	"soccer-manager/logger"

	"github.com/gin-gonic/gin"
)

// ContextKey is an internal type used for holding values in a `context.Context`
type ContextKey string

const (
	keyGinContextKey ContextKey = "soccer-manager-gin-context"
	keyRequest       ContextKey = "soccer-manager-req"
	keyRawRequest    ContextKey = "soccer-manager-raw-req"
	keyUserID        ContextKey = "soccer-manager-user-id"
	keyUserToken     ContextKey = "soccer-manager-user-token"
)

// contextOrBackground returns the given context if it is not nil.
// Returns context.Background() otherwise.
func contextOrBackground(ctx context.Context) context.Context {
	if ctx != nil {
		return ctx
	}
	return context.Background()
}

// CreateContextFromGinContext creates the context object from gin context
func CreateContextFromGinContext(c *gin.Context) context.Context {
	return context.WithValue(c.Request.Context(), keyGinContextKey, c)
}

// GinContextFromContext returns the gin context from context
func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(keyGinContextKey)
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}

// WithUserID is used to configure a context to add user id
func WithUserID(parent context.Context, userID int64) context.Context {
	return context.WithValue(contextOrBackground(parent), keyUserID, userID)
}

// WithUserToken is used to configure a context to add user token
func WithUserToken(parent context.Context, token string) context.Context {
	return context.WithValue(contextOrBackground(parent), keyUserToken, token)
}

// UserIDFromContext returns the user id from the context
func UserIDFromContext(ctx context.Context) (int64, error) {
	userID, ok := ctx.Value(keyUserID).(int64)
	if !ok {
		err := errors.New(constants.NoUserIdInContext)
		logger.Log.Error(err)
		return 0, err
	}
	return userID, nil
}

// UserTokenFromContext returns the user Token from the context
func UserTokenFromContext(ctx context.Context) (string, error) {
	userToken, ok := ctx.Value(keyUserToken).(string)
	if !ok {
		err := errors.New(constants.NoUserTokenInContext)
		logger.Log.Error(err)
		return "", err
	}
	return userToken, nil
}

// WithRequest adds the user data to context
func WithRequest(ctx context.Context, c *gin.Context) context.Context {
	body, _ := ioutil.ReadAll(c.Request.Body)
	ctx = context.WithValue(ctx, keyRequest, string(body))
	ctx = context.WithValue(ctx, keyRawRequest, *c.Request)
	c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
	return ctx

}

// RequestFromContext returns the user token from the context
func RequestFromContext(ctx context.Context) (string, error) {
	req, ok := ctx.Value(keyRequest).(string)
	if !ok {
		logger.Log.Error("No request present in context...")
		return "", errors.New("no request present in context")

	}
	return req, nil

}

// RawRequestFromContext returns the user token from the context
func RawRequestFromContext(ctx context.Context) (*http.Request, error) {
	req, ok := ctx.Value(keyRawRequest).(http.Request)
	if !ok {
		logger.Log.Error("No request present in context...")
		return nil, errors.New("no request present in context")

	}
	return &req, nil

}
