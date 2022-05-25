package utils

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"soccer-manager/logger"

	"github.com/gin-gonic/gin"
)

// ContextKey is an internal type used for holding values in a `context.Context`
type ContextKey string

const (
	keyGinContextKey ContextKey = "soccer-manager-gin-context"
	keyRequest       ContextKey = "soccer-manager-req"
	keyRawRequest    ContextKey = "soccer-manager-raw-req"
)

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
