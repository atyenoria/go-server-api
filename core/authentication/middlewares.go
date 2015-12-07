package authentication

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func RequireTokenAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {

		//		//x504 auth
		//		authBackend := InitJWTAuthenticationBackend()
		//		fmt.Println(c.Request.Header.Get("Authorization"))
		//		token, err := jwt.Parse(c.Request.Header.Get("Authorization"), func(token *jwt.Token) (interface{}, error) {
		//			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		//				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		//			} else {
		//				return authBackend.PublicKey, nil
		//			}
		//		})
		//		if err == nil && token.Valid && !authBackend.IsInBlacklist(c.Request.Header.Get("Authorization")) {
		//			c.Next()
		//		} else {
		//			c.Writer.WriteHeader(http.StatusUnauthorized)
		//		}
		//
		//

		//Secret string auth
		const myToken = "test"
		dtoken, err := jwt.Parse(c.Request.Header.Get("Authorization"), func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			//		return myLookupKey(token.Header["kid"])
			return []byte(myToken), nil
		})

		if err == nil && dtoken.Valid {
			fmt.Println("ok")
			fmt.Println(dtoken)
			c.Next()
		} else {
			c.Redirect(301, "/aaaaaaa")

		}

	}
}
