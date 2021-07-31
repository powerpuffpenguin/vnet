package vnet

import (
	"errors"

	"github.com/powerpuffpenguin/vnet/errs"
)

var ErrClosed = errors.New(`already closed`)
var ErrListenerClosed = errs.WrapError(ErrClosed, `listener already closed`)
var ErrDialerClosed = errs.WrapError(ErrClosed, `dialer already closed`)
