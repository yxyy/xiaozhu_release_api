package common

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"time"
	"xiaozhu/internal/config/cache"
	"xiaozhu/internal/model/key"
)

func GetAppStoreServerAPIToken() (string, error) {
	claims := jwt.MapClaims{
		"iss": viper.GetString("pay.5.iss "),
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(key.AppStoreServerAPITokenExpress).Unix(),
		"aud": "appstoreconnect-v1",
		"bid": viper.GetString("pay.5.bundleId"),
	}

	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 创建 JWT 的 Header
	header := map[string]interface{}{
		"alg": "HS256",                        // 加密算法
		"kid": viper.GetString("pay.5.keyId"), // App Store Connect 中的密钥 ID
		"typ": "JWT",                          // 类型
	}

	Token.Header = header

	return Token.SignedString([]byte(viper.GetString("pay.5.key")))

}

func GetAppleToken(ctx context.Context) (string, error) {
	result, err := cache.RedisDB00.Get(ctx, key.AppStoreServerAPIToken).Result()
	if err == nil {
		return result, nil
	}

	if !errors.Is(err, redis.Nil) {
		return "", err
	}

	token, err := GetAppStoreServerAPIToken()
	if err != nil {
		return "", err
	}

	err = cache.RedisDB00.SetNX(ctx, key.AppStoreServerAPIToken, token, key.AppStoreServerAPITokenExpress-time.Minute).Err()
	if err != nil {
		return "", err
	}

	return token, nil

}

// JWSTransactionDecodedPayload https://developer.apple.com/documentation/appstoreserverapi/jwstransactiondecodedpayload
type JWSTransactionDecodedPayload struct {
	AppAccountToken             string `json:"appAccountToken"`
	BundleId                    string `json:"bundleId"`
	Currency                    string `json:"currency"`
	Environment                 string `json:"environment"`
	ExpiresDate                 int64  `json:"expiresDate"`
	InAppOwnershipType          string `json:"inAppOwnershipType"`
	IsUpgraded                  bool   `json:"isUpgraded"`
	OfferDiscountType           string `json:"offerDiscountType"`
	OfferIdentifier             string `json:"offerIdentifier"`
	OfferType                   int32  `json:"offerType"`
	OriginalPurchaseDate        int64  `json:"originalPurchaseDate"`
	OriginalTransactionId       string `json:"originalTransactionId"`
	Price                       int64  `json:"price"`
	ProductId                   string `json:"productId"`
	PurchaseDate                int64  `json:"purchaseDate"`
	Quantity                    int32  `json:"quantity"`
	RevocationDate              int64  `json:"revocationDate"`
	RevocationReason            int32  `json:"revocationReason"`
	SignedDate                  int64  `json:"signedDate"`
	Storefront                  string `json:"storefront"`
	StorefrontId                string `json:"storefrontId"`
	SubscriptionGroupIdentifier string `json:"subscriptionGroupIdentifier"`
	TransactionId               string `json:"transactionId"`
	TransactionReason           string `json:"transactionReason"`
	PType                       string `json:"type"`
	WebOrderLineItemId          string `json:"webOrderLineItemId"`
	jwt.RegisteredClaims
}

func ParseAppStoreServerAPIToken(tokenString string) (*JWSTransactionDecodedPayload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWSTransactionDecodedPayload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("Auth.RefreshSecret")), nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	// 确保 token 是有效的并且是我们期望的类型
	claims, ok := token.Claims.(*JWSTransactionDecodedPayload)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token or claims")
	}

	return claims, nil
}
