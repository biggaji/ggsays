package helper

import "errors"

var (
	ErrPostNotFound            = errors.New("post not found")
	ErrCantGenerateAccessToken = errors.New("failed to generate access token")
)
