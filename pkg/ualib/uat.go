package ualib

import (
	"encoding/base64"
	"encoding/json"
	"github.com/Elyart-Network/UAC/internal/encrypt"
	"github.com/Elyart-Network/UAC/pkg/ualib/model"
	"strings"
)

func GenUAT(payload model.UATokenPayload, private []byte) (uat string, err error) {
	// Prepare header and payload [JSON]
	header := model.UATokenHeader{
		Type: "JWT",
		Alg:  "RS256",
	}
	headerJSON, err := json.Marshal(header)
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	// Encode header and payload
	b64Header := base64.StdEncoding.EncodeToString(headerJSON)
	b64Payload := base64.StdEncoding.EncodeToString(payloadJSON)
	encodedStr := b64Header + "." + b64Payload

	// Sign [RS256]
	sign, err := encrypt.SignRS256([]byte(encodedStr), private)
	if err != nil {
		return "", err
	}
	b64Sign := base64.StdEncoding.EncodeToString(sign)
	uat = encodedStr + "." + b64Sign
	return uat, nil
}

func VfUAT(uat string, public []byte) (payload model.UATokenPayload, err error) {
	// Split UAT
	uatSplit := strings.Split(uat, ".")

	// Validate signature
	sign, err := base64.StdEncoding.DecodeString(uatSplit[2])
	if err != nil {
		return payload, err
	}
	encodedStr := uatSplit[0] + "." + uatSplit[1]
	err = encrypt.VerifyRS256([]byte(encodedStr), sign, public)
	if err != nil {
		return payload, err
	}

	// Decode payload
	payloadJSON, err := base64.StdEncoding.DecodeString(uatSplit[1])
	if err != nil {
		return payload, err
	}
	err = json.Unmarshal(payloadJSON, &payload)
	if err != nil {
		return payload, err
	}
	return payload, nil
}
