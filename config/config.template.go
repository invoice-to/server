package config

type OAuthToken struct {
	ConsumerKey, ConsumerSecret, Token, TokenSecret string
}

var TwitterTokens = OAuthToken{
	ConsumerKey: "",
	ConsumerSecret: "",
	Token: "",
	TokenSecret: "",
}
