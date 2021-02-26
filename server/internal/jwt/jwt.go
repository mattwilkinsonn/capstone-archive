package jwt

import (
	"errors"
	"os"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

func getJWK() (jwk.Key, error) {
	key, err := jwk.New([]byte(os.Getenv("JWT_SECRET_KEY")))

	return key, err
}

func CreateToken(username string) (string, error) {
	token := jwt.New()

	err := token.Set("username", username)
	if err != nil {
		return "", errors.New("Err setting username")
	}

	key, err := getJWK()
	if err != nil {
		return "", err
	}

	// TODO: key doesnt match algo atm
	signed, err := jwt.Sign(token, jwa., key)
	if err != nil {
		return "", err
	}

	return string(signed), nil
}

func ParseToken(encoded string) (username string, err error) {

	key, err := getJWK()
	if err != nil {
		return "", err
	}

	token, err := jwt.ParseString(encoded, jwt.WithVerify(jwa.RS256, key))
	if err != nil {
		return "", err
	}

	returned, b := token.Get("username")
	if !b {
		return "", errors.New("Error getting username")
	}
	username = returned.(string)

	return

}
