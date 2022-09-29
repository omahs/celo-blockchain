package types

func (h *Header) IsIstanbulHeader() bool {
	if h.deserializedExtra == nil {
		// TODO: Should we also cache if deserialisation fails?
		extra, err := ExtractIstanbulExtra(h)
		if err != nil {
			return false
		}
		h.deserializedExtra = extra
	}

	// FIXME: Add proper signature check
	if h.deserializedExtra != nil && h.deserializedExtra.Seal != nil {
		return true
	}
	return false
}
