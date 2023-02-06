package mcf

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
	// [a-zA-Z0-9./] (required)
	return true
}

type hash string

func (h *hash) serialize() string {
	return "$" + string(*h)
}

func (h *hash) validate() bool {
	if *h == "" {
		return false
	}
	// TODO
	// [a-zA-Z0-9./]
	return true
}
