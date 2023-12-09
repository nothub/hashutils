package phc

import (
	"fmt"
	"hub.lol/hashutils"
	"hub.lol/hashutils/mcf"
	"strings"
)

// PHC string format specifies encodings for
// the serialization of a hash value and its metadata.
// https://github.com/P-H-C/phc-string-format/blob/master/phc-sf-spec.md
type PHC struct {
	id      id
	version version
	params  params
	salt    salt
	hash    hash
}

func (phc *PHC) Id() string {
	return string(phc.id)
}

func (phc *PHC) Version() string {
	return string(phc.version)
}

func (phc *PHC) Params() params {
	return phc.params
}

func (phc *PHC) Salt() string {
	return string(phc.salt)
}

func (phc *PHC) Hash() string {
	return string(phc.hash)
}

func (phc *PHC) Serialize() string {
	return fmt.Sprintf("%s%s%s%s%s",
		phc.id.serialize(),
		phc.version.serialize(),
		phc.params.serialize(),
		phc.salt.serialize(),
		phc.hash.serialize(),
	)
}

func (phc *PHC) Validate() bool {
	if phc == nil ||
		!phc.id.validate() ||
		!phc.version.validate() ||
		!phc.params.validate() ||
		!phc.salt.validate() ||
		!phc.hash.validate() {
		return false
	}

	if phc.hash != "" && phc.salt == "" {
		return false
	}

	return true
}

func (phc *PHC) ToMCF() (*mcf.MCF, error) {
	var format mcf.MCF

	err := format.SetId(string(phc.id))
	if err != nil {
		return nil, err
	}

	err = format.SetHash(string(phc.hash))
	if err != nil {
		return nil, err
	}

	return &format, nil
}

func Parse(str string) (*PHC, error) {
	if !strings.HasPrefix(str, "$") {
		return nil, hashutils.ErrMissingPrefix
	}

	str = strings.TrimPrefix(str, "$")
	str = strings.TrimSuffix(str, "$")

	split := strings.Split(str, "$")
	if len(split) < 1 {
		return nil, hashutils.ErrElementCount
	}
	if len(split) > 5 {
		return nil, hashutils.ErrElementCount
	}

	var phc PHC

	// hash algo identifier
	phc.id = id(split[0])
	split = split[1:]

	for len(split) > 0 {
		t := split[0]
		split = split[1:]

		if strings.Contains(t, "=") {
			if strings.Contains(t, ",") {
				// list of kv pairs
				for _, kv := range strings.Split(t, ",") {
					phc.params = append(phc.params, paramFromString(kv))
				}

			} else {
				// single kv pair
				pair := strings.SplitN(t, "=", 2)
				if len(pair) < 2 {
					return nil, hashutils.ErrParseFail
				}
				if pair[0] == "v" && phc.version == "" {
					phc.version = version(pair[1])
				} else {
					phc.params = append(phc.params, param{
						K: pair[0],
						V: pair[1],
					})
				}
			}
			continue
		}

		// only non-kv-pair elements left (salt and hash)
		if len(split) == 1 {
			phc.salt = salt(t)
			phc.hash = hash(split[0])
		} else {
			phc.hash = hash(t)
		}
	}

	if !phc.Validate() {
		return nil, hashutils.ErrInvalidData
	}

	return &phc, nil
}
