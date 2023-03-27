package errors

import "errors"

var ErrInvalidAccessToken = errors.New("invalid auth token")
var ErrUserDoesNotExist = errors.New("user does not exist")
var ErrUserAlreadyExists = errors.New("user with such credentials already exist")
var ErrUserDoesNotHaveAccess = errors.New("user does not have access")
var ErrFileDoesNotExist = errors.New("file does not exist")
