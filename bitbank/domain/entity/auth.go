package entity

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
	"time"

	"github.com/ganganbiz1/auto-trade-batch-go/bitbank/config"
)

type Auth struct {
	APIConfig *config.APIConfig
	Nonce     string
	APIInfo   *APIInfo
}

func NewAuth(
	c *config.APIConfig,
	apiInfo *APIInfo,
) *Auth {
	return &Auth{
		APIConfig: c,
		Nonce:     strconv.FormatInt(time.Now().UnixNano(), 10),
		APIInfo:   apiInfo,
	}
}

func (a *Auth) CreateRequestHeader() map[string]string {
	header := map[string]string{
		"ACCESS-KEY":       a.APIConfig.APIKey,
		"ACCESS-NONCE":     a.Nonce,
		"ACCESS-SIGNATURE": a.Signature(),
	}

	if a.APIInfo.Method == "POST" {
		header["Content-Type"] = "application/json"
	}
	return header
}

func (a *Auth) Signature() string {
	var message string
	if a.APIInfo.Method == "GET" {
		message = a.signaturePayloadForGet()
	} else {
		message = a.signaturePayloadForPost(a.APIInfo.Body)
	}

	mac := hmac.New(sha256.New, []byte(a.APIConfig.APISecret))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}

func (a *Auth) signaturePayloadForGet() string {
	queries := []string{}
	params := ""
	if len(a.APIInfo.Query) > 0 {
		for k, v := range a.APIInfo.Query {
			queryTmp := k + "=" + v
			queries = append(queries, queryTmp)
		}
		params = "?" + strings.Join(queries, "&")
	}
	return a.Nonce + a.APIInfo.Path + params
}

func (a *Auth) signaturePayloadForPost(body []byte) string {
	return strconv.FormatInt(time.Now().UnixNano(), 10) + string(body)
}
