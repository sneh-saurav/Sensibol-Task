package main

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/oauth2/clientcredentials"

	"github.com/zmb3/spotify"
)

const clientID = "9613716112a346febac9260df8882781"
const clientSecret = "0cf561a2e5a845779b81e796acbca349"

var link = "https://open.spotify.com/track/6DDux0xQXq8DcEgdaWWtC7?si="
var TrackURL []string

func main() {
	config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     spotify.TokenURL,
	}
	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	client := spotify.Authenticator{}.NewClient(token)

	results, err := client.Search("Singh Vs Kaur", spotify.SearchTypeTrack)
	if err != nil {
		log.Fatal(err)
	}

	// handle track results
	if results.Tracks != nil {
		fmt.Println("Track Ids:")
		for _, track := range results.Tracks.Tracks {
			fmt.Println("   ", track.ID.String())
			link = link + track.ID.String()
			TrackURL = append(TrackURL, link)
		}
	}
	fmt.Println("Spotify Link for the song", TrackURL)
}
