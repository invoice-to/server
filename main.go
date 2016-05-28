package main

import (
	"fmt"
	"net/http"
	anaconda "github.com/ChimeraCoder/anaconda"
)

var TwitterAPI = createTwitterAPI()

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/auth", authHandler)
}

func createTwitterAPI() anaconda.TwitterAPI {
	anaconda.SetConsumerKey(config.TwitterToken.ConsumerKey)
	anaconda.SetConsumerSecret(config.TwitterToken.ConsumerSecret)

	return anaconda.NewTwitterApi(config.TwitterToken.Token, config.TwitterToken.TokenSecret)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "👋 \ncheck out /auth (it's 🔥)")
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "yo")
}
