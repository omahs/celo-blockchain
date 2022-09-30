package types

func (h *Header) IstanbulExtra() (*IstanbulExtra, error) {
	h.extraLock.RLock()
	extraData := h.deserializedExtra
	h.extraLock.RUnlock()

	if extraData == nil {
		extra, err := extractIstanbulExtra(h)
		if err != nil {
			return nil, err
		}
		h.extraLock.Lock()
		h.deserializedExtra = extra
		h.extraLock.Unlock()
	}

	h.extraLock.RLock()
	defer h.extraLock.RUnlock()
	return h.deserializedExtra, nil
}
