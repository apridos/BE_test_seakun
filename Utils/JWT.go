package Utils

import (
	"os"
	Model "seakun/Model"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte(os.Getenv("SECRET_KEY"))

func GenerateJWTToken(admin Model.Admin) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = admin.Username
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateJWTRefreshToken(admin Model.Admin) (string, error) {
	refresh_token := jwt.New(jwt.SigningMethodHS256)

	claims := refresh_token.Claims.(jwt.MapClaims)

	claims["username"] = admin.Username
	claims["exp"] = time.Now().Add(time.Hour * 24 * 14).Unix()

	tokenString, err := refresh_token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWTToken(tokenString string, username *string) int {

	var code = 200

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			code = 400
		}
		checkExpired := token.Claims.(jwt.MapClaims).VerifyExpiresAt(time.Now().Unix(), false)
		if !checkExpired {
			code = 410
		}

		return mySigningKey, nil
	})

	if err != nil {
		if code != 410 {
			code = 400
		}
	}

	*username, _ = token.Claims.(jwt.MapClaims)["username"].(string)

	if token.Valid {
		return code
	}

	return code
}
