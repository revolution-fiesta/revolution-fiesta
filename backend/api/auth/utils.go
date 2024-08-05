package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"main/backend/config"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type AudienceType string

const (
	audienceTypeAccessTokenUser AudienceType = "access_token_user"
)

func GenerateAccessToken(userId int, privateKey *rsa.PrivateKey) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.RegisteredClaims{
		Subject:   strconv.Itoa(userId),
		Issuer:    config.ProjectName,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(config.AccessTokenExpiration)),
		Audience:  jwt.ClaimStrings{string(audienceTypeAccessTokenUser)},
		// actually this is only used for key rotation.
		ID: uuid.NewString(),
	})
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", errors.Wrapf(err, "failed to sign token")
	}
	return signedToken, nil
}

// validate access token and return user id.
func ValidateAccessToken(token string, publicKey *rsa.PublicKey) (int, error) {
	claims := &jwt.RegisteredClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (any, error) {
		if alg := t.Method.Alg(); alg != jwt.SigningMethodRS256.Name {
			return nil, errors.Errorf("sighing method not supported: %q", alg)
		}
		return publicKey, nil
	})
	if err != nil {
		return 0, errors.Wrapf(err, "failed to parse claims")
	}

	if len(claims.Audience) != 1 || claims.Audience[0] != string(audienceTypeAccessTokenUser) {
		return 0, errors.New(fmt.Sprintf("audience type not supported: %q", claims.Audience[0]))
	}

	if claims.Issuer != config.ProjectName {
		return 0, errors.New(fmt.Sprintf("wrong issuer: %q", claims.Issuer))
	}

	userId, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

// hash salt and password with SHA-256.
func Sha256(text string, salt string) string {
	hasher := sha256.New()
	_, _ = hasher.Write([]byte(salt))
	_, _ = hasher.Write([]byte(text))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func GeneratePrivateKey() (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}
