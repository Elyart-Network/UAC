package model

// UATokenDSm UAToken under Data Source mode model [request]
type UATokenDSm struct {
	UATokenReqBase
	Data string `json:"data"` // AES encrypted data
}

// UATokenPVm UAToken under Provider mode model [request]
type UATokenPVm struct {
	UATokenReqBase
	Code string `json:"code"` // code from auth flow
}

// UAToken response model [response]
type UAToken struct {
	UAToken string `json:"ua_token"`
}

// UATokenRef request model [request]
type UATokenRef struct {
	UATokenReqBase
	UAToken string `json:"ua_token"` // Old UAToken
}

// UATokenHeader UAToken JWT Header Section
type UATokenHeader struct {
	Type string `json:"typ"` // JWT
	Alg  string `json:"alg"` // RS256
}

// UATokenPayload UAToken JWT Payload Section
type UATokenPayload struct {
	Issuer     string `json:"iss,omitempty"` // UAC
	Subject    string `json:"sub,omitempty"` // User UUID
	Audience   string `json:"aud,omitempty"` // Client UUID
	NotBefore  int64  `json:"nbf,omitempty"` // Not before
	ExpireAt   int64  `json:"exp,omitempty"` // Expiration time
	IssuedAt   int64  `json:"iat,omitempty"` // Issued at
	RefreshExp int64  `json:"ref,omitempty"` // Refresh expiration
	JwtId      string `json:"jti,omitempty"` // JWT ID
}
