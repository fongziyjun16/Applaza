package handler

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
)

func InitRouter() *mux.Router {
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(mySigningKye), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	router := mux.NewRouter()
	router.Handle("/upload", jwtMiddleware.Handler(http.HandlerFunc(uploadHandler))).Methods("POST")
	router.Handle("/checkout", jwtMiddleware.Handler(http.HandlerFunc(checkoutHandler))).Methods("POST")
	router.Handle("/search", jwtMiddleware.Handler(http.HandlerFunc(searchHandler))).Methods("GET")
	router.Handle("/checkout", jwtMiddleware.Handler(http.HandlerFunc(checkoutHandler))).Methods("POST")
	router.Handle("/app/{id}", jwtMiddleware.Handler(http.HandlerFunc(deleteHandler))).Methods("DELETE")
	router.Handle("/signup", http.HandlerFunc(signUpHandler)).Methods("POST")
	router.Handle("/signin", http.HandlerFunc(signInHandler)).Methods("POST")
	return router
}
