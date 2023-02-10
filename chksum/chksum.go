package chksum

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/nothub/hashutils"
	"github.com/nothub/hashutils/encoding"
	"github.com/nothub/hashutils/encoding/b64"
	"hash"
	"io"
	"log"
	"os"
)

func Create(reader io.Reader, algo hash.Hash, enco encoding.Scheme) (string, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}

	algo.Write(b)
	defer algo.Reset()

	switch enco {
	case encoding.Hex:
		return hex.EncodeToString(algo.Sum(nil)), nil
	case encoding.B64:
		return b64.Encode(algo.Sum(nil)), nil
	default:
		return "", hashutils.ErrUnknownEncoding
	}
}

func Verify(reader io.Reader, chksum string, algo hash.Hash, enco encoding.Scheme) (bool, error) {
	c, err := Create(reader, algo, enco)
	if err != nil {
		return false, err
	}

	return c == chksum, nil
}

func VerifyFile(path string, chksum string, algo hash.Hash, enco encoding.Scheme) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, err
		}
		log.Fatalln(err.Error())
	}
	if info.IsDir() {
		return false, hashutils.ErrFileIsDir
	}
	if !info.Mode().IsRegular() {
		return false, hashutils.ErrFileNotRegular
	}

	file, err := os.Open(path)
	if err != nil {
		return false, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(file)

	return Verify(file, chksum, algo, enco)
}
