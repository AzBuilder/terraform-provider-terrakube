package helpers

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func GetClaimFromToken(jwtToken string, claim string) (string, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(jwtToken, jwt.MapClaims{})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("failed to parse claims")
	}

	c, ok := claims[claim].(string)
	if !ok {
		return "", fmt.Errorf("claim %s is not a string", claim)
	}
	return c, nil
}

func GetIDFromToken(jwtToken string) (string, error) {
	return GetClaimFromToken(jwtToken, "jti")
}
