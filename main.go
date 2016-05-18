package main

import (
	"fmt"
	"net/http"
	"github.com/ChimeraCoder/anaconda"
	config "./config"
)

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/auth", authHandler)
	setupTwitter(config.TwitterTokens.key, config.TwitterTokens.secret)
}

func setupTwitter(key string, secret string) {
	fmt.Printf("setting up")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "👋 \ncheck out /auth (it's 🔥)")
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "yo")
}
