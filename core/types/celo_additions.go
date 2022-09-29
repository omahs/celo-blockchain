package types

func (h *Header) IstanbulExtra() (*IstanbulExtra, error) {
	if h.deserializedExtra == nil {
		extra, err := ExtractIstanbulExtra(h)
		if err != nil {
			return nil, err
		}
		h.deserializedExtra = extra
	}

	return h.deserializedExtra, nil
}
