package cowww
type Header map[string]string

func (h Header) Set(key, val string) {
	if h == nil {
		return
	}

	h[key] = val
}

func (h Header) Get(key string) string {
	if h == nil {
		return ""
	}

	return h[key]
}

const DefaultProto = "HTTP/1.1"
