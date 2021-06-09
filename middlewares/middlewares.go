package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func SetCors(next http.Handler) http.Handler{

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods","*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, token")
		if r.Method == "OPTIONS"{
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
		
	})
}


func CheckAuth(next http.Handler) http.Handler{

			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				tokenHeader := r.Header.Get("token")
				if len(tokenHeader) < 2{
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("Token was not correctly formatted"))
				} else{
					token, err := jwt.Parse(tokenHeader, func(token *jwt.Token) (interface{}, error){
						if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
							return nil, fmt.Errorf("Unexpected signing method")
						}
						return []byte(os.Getenv("SECRET")), nil
					})
					if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid{
						ctx := context.WithValue(r.Context(), "id", claims)

						next.ServeHTTP(w, r.WithContext(ctx))
					} else {
						fmt.Println(err)
						w.WriteHeader(http.StatusUnauthorized)
						w.Write([]byte("Unauthorized"))
					}
				}
			})
}