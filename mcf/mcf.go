package mcf

import (
	"fmt"
	"github.com/nothub/hashutils"
	"strings"
)

// MCF (Modular Crypt Format) specifies encodings for
// the serialization of a hash value and its metadata.
// https://passlib.readthedocs.io/en/stable/modular_crypt_format.html
type MCF struct {
	id   id
	hash hash
}

func (mcf *MCF) Id() string {
	return string(mcf.id)
}
func (mcf *MCF) SetId(str string) error {
	e := id(str)
	if !e.validate() {
		return hashutils.ErrInvalidData
	}
	mcf.id = e
	return nil
}

func (mcf *MCF) Hash() string {
	return string(mcf.hash)
}
func (mcf *MCF) SetHash(str string) error {
	e := hash(str)
	if !e.validate() {
		return hashutils.ErrInvalidData
	}
	mcf.hash = e
	return nil
}

func (mcf *MCF) Serialize() string {
	return fmt.Sprintf("$%s$%s", mcf.id, mcf.hash)
}

func Parse(str string) (*MCF, error) {
	if !strings.HasPrefix(str, "$") {
		return nil, hashutils.ErrMissingPrefix
	}

	str = strings.TrimPrefix(str, "$")
	str = strings.TrimSuffix(str, "$")

	split := strings.Split(str, "$")
	if len(split) != 2 {
		return nil, hashutils.ErrElementCount
	}

	return &MCF{
		id:   id(split[0]),
		hash: hash(split[1]),
	}, nil
}
