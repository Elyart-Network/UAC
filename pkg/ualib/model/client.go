package model

type ClientBase struct {
	UUID     string `json:"uuid"`
	Secret   string `json:"secret"`   // AES encrypted secret
	Type     string `json:"type"`     // native/oidc
	Callback string `json:"callback"` // UAMode = "pv"
}
