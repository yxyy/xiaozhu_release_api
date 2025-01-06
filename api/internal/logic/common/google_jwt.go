package common

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
	"xiaozhu/internal/model/key"
)

var GoogleJWT *GoogleJWTResponse

const googleTokenUrl = "https://oauth2.googleapis.com/token"

// GoogleJWTResponse https://developers.google.com/identity/protocols/oauth2/service-account?hl=zh-cn#httprest
type GoogleJWTResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	ExpireDate  int64
}

func NewJWT() (string, error) {
	claims := jwt.MapClaims{
		"iss":   viper.GetString("pay.4.iss "),
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(key.GoogleServerAPITokenExpress).Unix(),
		"aud":   "https://oauth2.googleapis.com/token",
		"scope": viper.GetString("pay.4.bundleId"),
	}
	Token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	Token.Header["kid"] = viper.GetString("pay.4.keyId") // App Store Connect 中的密钥 ID
	return Token.SignedString([]byte(viper.GetString("pay.4.key")))
}

func GetGoogleServerAPIToken(ctx context.Context) error {
	newJWT, err := NewJWT()
	if err != nil {
		return err
	}

	values := url.Values{}
	values.Add("grant_type", "urn:ietf:params:oauth:grant-type:jwt-bearer")
	values.Add("assertion", newJWT)

	request, err := http.NewRequestWithContext(ctx, "POST", googleTokenUrl, strings.NewReader(values.Encode()))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch token, status: %d, response: %s", resp.StatusCode, body)
	}

	var jwtResponse GoogleJWTResponse
	if err = json.Unmarshal(body, &jwtResponse); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	jwtResponse.ExpireDate = time.Now().Add(time.Duration(jwtResponse.ExpiresIn) * time.Second).Unix()
	GoogleJWT = &jwtResponse
	return nil

}

func GetGoogleToken(ctx context.Context) (string, error) {
	if GoogleJWT == nil || GoogleJWT.ExpireDate <= time.Now().Unix() {
		if err := GetGoogleServerAPIToken(ctx); err != nil {
			return "", err
		}
	}

	return GoogleJWT.AccessToken, nil

}
