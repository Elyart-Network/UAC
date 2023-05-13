package model

type UATokenReqBase struct {
	UAMode string     `json:"ua_mode"` // ds/pv
	Client ClientBase `json:"client"`  // client info
	Secret string     `json:"secret"`  // RS256 encrypted AES key
}
