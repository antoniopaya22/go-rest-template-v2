package crypto

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

// Generates Hash Password
//
// parameters:
//		pwd []byte
// returns
//		string hashed password
func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// Compares two passwords
//
// parameters:
//		hashedPwd string
//		plainPwd []byte
//
//	returns:
//		boolean True o False
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}
	return true
}

// Creates and returns JWT token
//
// parameters:
//		username string User name
//		role string User Roles
//
// returns:
//		token string
func CreateToken(username string, role string) (string, error) {
	secret := "jdnfksdmfksda"
	if os.Getenv("jwtTokenSecret") != "" {
		secret = os.Getenv("jwtTokenSecret")
	}

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["username"] = username
	atClaims["role"] = role
	atClaims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, atClaims)
	token, err := at.SignedString([]byte(secret)) // SECRET
	if err != nil {
		return "token creation error", err
	}
	return token, nil
}


// Validates token
//
// parameters:
//		tokenString string JWT Token
// returns:
//		boolean Valid token
func ValidateToken(tokenString string) bool {
	secret := "jdnfksdmfksda"
	if os.Getenv("jwtTokenSecret") != "" {
		secret = os.Getenv("jwtTokenSecret")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false
	}
	return token.Valid
}
