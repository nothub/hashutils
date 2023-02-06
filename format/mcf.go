package format

import (
	"fmt"
	"github.com/nothub/hashutils"
	"strings"
)

// Modular Crypt Format
// https://passlib.readthedocs.io/en/stable/modular_crypt_format.html

type MCF struct {
	id   id
	hash hash
}

func (mcf *MCF) Id() string {
	return string(mcf.id)
}

func (mcf *MCF) Hash() string {
	return string(mcf.hash)
}

func (mcf *MCF) Serialize() string {
	return fmt.Sprintf("$%s$%s", mcf.id, mcf.hash)
}

func ParseMCF(str string) (*MCF, error) {
	if !strings.HasPrefix(str, "$") {
		return nil, fmt.Errorf("%s, missing $ prefix", hashutils.ErrParseFail.Error())
	}

	str = strings.TrimPrefix(str, "$")
	str = strings.TrimSuffix(str, "$")

	split := strings.Split(str, "$")
	if len(split) != 2 {
		return nil, hashutils.ErrParseFail
	}

	return &MCF{
		id:   id(split[0]),
		hash: hash(split[1]),
	}, nil
}
