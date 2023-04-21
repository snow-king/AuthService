package models

type JWK struct {
	Keys []struct {
		Kty string `json:"kty,omitempty"`
		Kid string `json:"kid,omitempty"`
		K   string `json:"k,omitempty"`
		Alg string `json:"alg,omitempty"`
		Use string `json:"use,omitempty"`
	} `json:"keys,omitempty"`
}
