package format

import (
	"fmt"
	"github.com/nothub/hashutils"
	"strings"
)

// PHC string format specifies encodings for the serialization of a hash value and its metadata.
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

func (phc *PHC) Params() map[string]string {
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

func (phc *PHC) ToMCF() MCF {
	return MCF{
		id:   phc.id,
		hash: phc.hash,
	}
}

func ParsePHC(str string) (*PHC, error) {
	if !strings.HasPrefix(str, "$") {
		return nil, fmt.Errorf("%s, missing $ prefix", hashutils.ErrParseFail.Error())
	}

	str = strings.TrimPrefix(str, "$")
	str = strings.TrimSuffix(str, "$")

	split := strings.Split(str, "$")
	if len(split) < 1 {
		return nil, fmt.Errorf("%s, missing elements", hashutils.ErrParseFail.Error())
	}
	if len(split) > 5 {
		return nil, fmt.Errorf("%s, too many elements", hashutils.ErrParseFail.Error())
	}

	var phc PHC
	phc.params = make(params)

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
					pair := strings.SplitN(kv, "=", 2)
					if len(pair) < 2 {
						return nil, hashutils.ErrParseFail
					}
					phc.params[pair[0]] = pair[1]
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
					phc.params[pair[0]] = pair[1]
				}
			}
			continue
		}

		// only non-kv-pair elements left are salt and hash
		if len(split) == 1 {
			phc.salt = salt(t)
			phc.hash = hash(split[0])
		} else {
			phc.hash = hash(t)
		}
	}

	if !phc.Validate() {
		return nil, fmt.Errorf("%s, data is not valid", hashutils.ErrParseFail.Error())
	}

	return &phc, nil
}
