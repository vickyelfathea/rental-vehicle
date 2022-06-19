package helpers

import ("github.com/golang-jwt/jwt/v4"
		"time"
	// "fmt"
)

var mySecretKey = []byte("rahasia")

type claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewToken(username string) *claims {
	return &claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims {
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
		},
	}
}

func (c *claims) Create() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString(mySecretKey)

}

func CheckToken(token string) (bool, error) {

	tokens, err := jwt.ParseWithClaims(token, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(mySecretKey), nil
	})

	if err != nil {
		return false, err
	}

	return tokens.Valid, nil
}


