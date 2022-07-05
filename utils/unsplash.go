package utils

import (
	"context"
	"os"

	"github.com/hbagdi/go-unsplash/unsplash"
	"golang.org/x/oauth2"
)

var unsplashClient *unsplash.Unsplash

func SetUnsplash() {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "Client-ID " + os.Getenv("ACCESS_TOKEN")},
	)
	client := oauth2.NewClient(context.Background(), ts)

	unsplashClient = unsplash.New(client)
}

func Unsplash() *unsplash.Unsplash {
	return unsplashClient
}
