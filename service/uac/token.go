package uac

import (
	"context"
	"encoding/json"
	"github.com/Elyart-Network/UAC/config"
	"github.com/Elyart-Network/UAC/data/actions"
	"github.com/Elyart-Network/UAC/internal/encrypt"
	"github.com/Elyart-Network/UAC/model"
	"github.com/Elyart-Network/UAC/pkg/ualib"
	ULibMod "github.com/Elyart-Network/UAC/pkg/ualib/model"
	"github.com/Elyart-Network/UAC/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func GenToken(private string, userUUID string, clientUUID string) (string, error) {
	// Generate Token
	timeNow := time.Now().Unix()
	ExpTime := time.Hour * 2
	RefTime := time.Hour * 24 * 7
	JwtId, err := encrypt.UUIDv2P()
	if err != nil {
		return "", err
	}
	payload := ULibMod.UATokenPayload{
		Issuer:     "UAC",
		Subject:    userUUID,
		Audience:   clientUUID,
		NotBefore:  timeNow,
		ExpireAt:   timeNow + int64(ExpTime.Seconds()),
		IssuedAt:   timeNow,
		RefreshExp: timeNow + int64(RefTime.Seconds()),
		JwtId:      JwtId.String(),
	}
	uat, err := ualib.GenUAT(payload, []byte(private))
	if err != nil {
		return "", err
	}

	// Cache Session
	sact := actions.NewSession(context.Background())
	err = sact.New("JwtId:"+JwtId.String(), payload)
	if err != nil {
		return "", err
	}
	return uat, nil
}

// TokenService godoc
// @Summary Token Endpoint
// @Description Generate and Maintain Tokens
// @Tags UAC
// @Accept json
// @Produce json
// @Success 200 {object} ULibMod.UAToken
// @Router /uac/token [post]
func TokenService(ctx *gin.Context) {
	// Prepare Values
	private := config.Get("encrypt.rsa_private").(string)
	codeSign := config.Get("encrypt.code_sign").(string)
	uact := actions.NewUser(ctx)
	cact := actions.NewClient()

	// Bind JSON
	req := ULibMod.UATokenReqBase{}
	err := ctx.BindJSON(&req)
	if err != nil {
		utils.CommonResponse(ctx, 400)
		return
	}

	// Decrypt AES Key
	aesKey, err := encrypt.RsaDecrypt([]byte(req.Secret), []byte(private))
	if err != nil {
		utils.CommonResponse(ctx, 500)
		return
	}

	// Check Client
	client, err := cact.Get(req.Client.UUID)
	if err != nil {
		utils.CommonResponse(ctx, 400)
		return
	}
	secret := encrypt.AesDecryptCBC([]byte(req.Client.Secret), aesKey)
	if client.Secret != string(secret) {
		utils.CommonResponse(ctx, 400)
		return
	}

	// UAC under Data Source Mode
	if req.UAMode == "ds" {
		// Bind JSON
		data := ULibMod.UATokenDSm{}
		err := ctx.BindJSON(&data)
		if err != nil {
			utils.CommonResponse(ctx, 400)
			return
		}

		// Decrypt Data
		raw := encrypt.AesDecryptCBC([]byte(data.Data), aesKey)
		var auth struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Identity string `json:"identity"`
			Addition string `json:"addition"`
			Register bool   `json:"register"`
		}
		err = json.Unmarshal(raw, &auth)
		if err != nil {
			utils.CommonResponse(ctx, 400)
			return
		}

		// If Register
		if auth.Register {
			pwdEnc, err := encrypt.BcryptE(auth.Password)
			uuid, err := encrypt.UUIDv5([]byte(auth.Username))
			if err != nil {
				utils.CommonResponse(ctx, 500)
				return
			}
			userNew := model.Users{
				UUID:     uuid.String(),
				Username: auth.Username,
				Password: pwdEnc,
				Identity: auth.Identity,
				IsAdmin:  false,
				Addition: auth.Addition,
			}
			_, err = uact.Insert(userNew)
			if err != nil {
				utils.CommonResponse(ctx, 500)
				return
			}
		}

		// Check Auth
		uuid, err := uact.ID(auth.Username)
		user, err := uact.Get(uuid)
		if err != nil {
			utils.CommonResponse(ctx, 500)
			return
		}
		if user.UUID == "" {
			utils.CommonResponse(ctx, 401)
			return
		}
		vfPwd, err := encrypt.BcryptV(auth.Password, user.Password)
		if !vfPwd {
			utils.CommonResponse(ctx, 401)
			return
		}

		// Generate Token
		uat, err := GenToken(private, user.UUID, client.UUID)

		// Response
		resp := ULibMod.UAToken{UAToken: uat}
		utils.BaseResponse(ctx, 200, "ok", resp)
	}

	// UAC under Provider Mode
	if req.UAMode == "pv" {
		// Bind JSON
		data := ULibMod.UATokenPVm{}
		err := ctx.BindJSON(&data)
		if err != nil {
			utils.CommonResponse(ctx, 400)
			return
		}

		// Check Code
		codeRaw := encrypt.AesDecryptCBC([]byte(data.Code), []byte(codeSign))
		var code struct {
			UUID      string `json:"uuid"`
			SessionId string `json:"session_id"`
		}
		err = json.Unmarshal(codeRaw, &code)
		if err != nil {
			utils.CommonResponse(ctx, 400)
			return
		}
		user, err := uact.Get(code.UUID)
		if err != nil {
			utils.CommonResponse(ctx, 500)
			return
		}
		if user.UUID == "" {
			utils.CommonResponse(ctx, 401)
			return
		}

		// TODO:Check Session

		// Generate Token
		uat, err := GenToken(private, user.UUID, client.UUID)

		// Response
		resp := ULibMod.UAToken{UAToken: uat}
		utils.BaseResponse(ctx, 200, "ok", resp)
	}
}

// RefreshService godoc
// @Summary Refresh Endpoint
// @Description Refresh Tokens
// @Tags UAC
// @Accept json
// @Produce json
// @Success 200 {object} ULibMod.UAToken
// @Router /uac/token [put]
func RefreshService(ctx *gin.Context) {
	// Prepare Values
	public := config.Get("encrypt.rsa_public").(string)
	private := config.Get("encrypt.rsa_private").(string)
	cact := actions.NewClient()
	sact := actions.NewSession(ctx)

	// Bind JSON
	req := ULibMod.UATokenRef{}
	err := ctx.BindJSON(&req)
	if err != nil {
		utils.CommonResponse(ctx, 400)
		return
	}

	// Check Client
	client, err := cact.Get(req.Client.UUID)
	if err != nil {
		utils.CommonResponse(ctx, 400)
		return
	}
	if client.Secret != req.Client.Secret {
		utils.CommonResponse(ctx, 400)
		return
	}

	// Check Token
	payload, err := ualib.VfUAT(req.UAToken, []byte(public))
	if err != nil {
		utils.CommonResponse(ctx, 400)
		return
	}
	timeNow := time.Now().Unix()
	if timeNow < payload.NotBefore {
		utils.CommonResponse(ctx, 400)
		return
	}
	if timeNow-payload.IssuedAt < payload.ExpireAt {
		utils.CommonResponse(ctx, 400)
		return
	}
	if timeNow-payload.IssuedAt > payload.RefreshExp {
		utils.CommonResponse(ctx, 401)
		return
	}

	// Check If Token(JwtId) is in Cache
	session, err := sact.Get("JwtId:" + payload.JwtId)
	if err != nil || session == "[]" {
		utils.CommonResponse(ctx, 401)
		return
	}

	// Generate Token
	uat, err := GenToken(private, payload.Subject, client.UUID)

	// Remove Old Token from Cache
	err = sact.Delete("JwtId:" + payload.JwtId)
	if err != nil {
		utils.CommonResponse(ctx, 500)
		return
	}

	// Response
	resp := ULibMod.UAToken{UAToken: uat}
	utils.BaseResponse(ctx, 200, "ok", resp)
}
