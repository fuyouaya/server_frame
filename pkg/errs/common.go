package errs

import (
	"net/http"
)

var (
	ErrUnknown       = NewError(http.StatusInternalServerError, 1)
	ErrUnauthorized  = NewError(http.StatusUnauthorized, 11010)
	ErrForbidden     = NewError(http.StatusForbidden, 1)
	ErrNotFound      = NewError(http.StatusNotFound, 1)
	ErrParamInvalid  = NewError(http.StatusBadRequest, 2)
	ErrTokenExpire   = NewError(http.StatusGone, 1)
	ErrContextCancel = NewError(499, 1)

	ErrRateLimit          = NewError(http.StatusTooManyRequests, 110001)
	ErrRateLimitAfterTime = NewError(http.StatusTooManyRequests, 110002)

	ErrLinkNotFound   = NewError(http.StatusBadRequest, 1)
	ErrPatchLinkError = NewError(http.StatusBadRequest, 1)
	ErrLinkExpired    = NewError(http.StatusBadRequest, 1)

	ErrUserWechatBinded = NewError(http.StatusConflict, 1)
	ErrUserExists       = NewError(http.StatusConflict, 1)
	ErrUserLocked       = NewError(http.StatusLocked, 0)
)
