package auth

import (
	"log"

	jwt "github.com/dgrijalva/jwt-go"
)

type UserJwt struct {
	Id   uint
	Mail string
	jwt.StandardClaims
}

func CreateTokenString(id uint, mail string) string {

	// User情報をtokenに込める
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &UserJwt{
		Id:   id,
		Mail: mail,
	})
	// Secretで文字列にする. このSecretはサーバだけが知っている
	tokenstring, err := token.SignedString([]byte("foobar"))
	if err != nil {
		log.Fatalln(err)
	}
	return tokenstring
}

func ValidateTokenString(tokenstring string) UserJwt {
	log.Println(tokenstring)

	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return []byte("foobar"), nil
	})

	// Parseメソッドを使うと、Claimsはmapとして得られる
	log.Println(token.Claims, err)
	userJwt := UserJwt{}
	token, err = jwt.ParseWithClaims(tokenstring, &userJwt, func(token *jwt.Token) (interface{}, error) {
		return []byte("foobar"), nil
	})
	return userJwt
	//log.Println(token.Valid, user, err)

}
