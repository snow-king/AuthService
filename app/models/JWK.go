package models

type JWK struct {
	Keys []struct {
		Kty string `json:"kty,omitempty"`
		Kid string `json:"kid,omitempty"`
		Alg string `json:"alg,omitempty"`
	} `json:"keys,omitempty"`
}
