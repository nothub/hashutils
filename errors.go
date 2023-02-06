package hashutils

import (
	"errors"
	"fmt"
	"io/fs"
)

var ErrParseFail = errors.New("failed parsing")
var ErrFileIsDir = fmt.Errorf("%s, file is a directory", fs.ErrInvalid)
var ErrFileNotRegular = fmt.Errorf("%s, file is not regular", fs.ErrInvalid)
var ErrUnknownEncoding = errors.New("unknown encoding")
