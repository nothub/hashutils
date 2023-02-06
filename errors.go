package hashutils

import (
	"errors"
	"fmt"
	"io/fs"
)

var ErrFileIsDir = fmt.Errorf("%w, file is a directory", fs.ErrInvalid)
var ErrFileNotRegular = fmt.Errorf("%w, file is not regular", fs.ErrInvalid)
var ErrParseFail = errors.New("failed parsing")
var ErrMissingPrefix = fmt.Errorf("%w, missing $ prefix", ErrParseFail)
var ErrElementCount = fmt.Errorf("%w, wrong element count", ErrParseFail)
var ErrUnknownEncoding = errors.New("unknown encoding")
var ErrInvalidData = errors.New("invalid data")
