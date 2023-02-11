package phc

import (
	"fmt"
	"strings"
)

// Elements partially shared by the MCF and PHC format.

type id string

func (id *id) serialize() string {
	return "$" + string(*id)
}

func (id *id) validate() bool {
	if *id == "" {
		return false
	}
	// TODO
	// [a-z0-9-]
	return true
}

type version string

func (v *version) serialize() string {
	return "$v=" + string(*v)
}

func (v *version) validate() bool {
	if *v == "" {
		// optional
		return true
	}
	// TODO
	// [0-9]
	return true
}

type param struct {
	K string
	V string
}

func paramFromString(str string) param {
	pair := strings.SplitN(str, "=", 2)
	return param{
		K: pair[0],
		V: pair[1],
	}
}

type params []param

func (p params) serialize() string {
	var pairs []string
	for _, pa := range p {
		pairs = append(pairs, fmt.Sprintf("%s=%s", pa.K, pa.V))
	}
	return "$" + strings.Join(pairs, ",")
}

func (p params) validate() bool {
	if len(p) == 0 {
		// optional
		return true
	}
	for _, pa := range p {
		if pa.K == "" || pa.V == "" {
			return false
		}
	}
	// TODO
	// [a-z0-9-]=[a-zA-Z0-9/+.-]
	return true
}

type salt string

func (s *salt) serialize() string {
	return "$" + string(*s)
}

func (s *salt) validate() bool {
	if *s == "" {
		// optional
		return true
	}
	// TODO
	// [a-zA-Z0-9/+.-]
	return true
}

type hash string

func (h *hash) serialize() string {
	return "$" + string(*h)
}

func (h *hash) validate() bool {
	if *h == "" {
		// optional
		return true
	}
	// TODO
	// B64 format without padding (=) (requires salt)
	return true
}
