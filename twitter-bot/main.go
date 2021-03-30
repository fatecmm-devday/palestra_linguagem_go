package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/coreos/pkg/flagutil"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/fatih/color"
)

func main() {
	flags := flag.NewFlagSet("user-auth", flag.ExitOnError)
	consumerKey := flags.String("consumer-key", "", "Twitter Consumer Key")
	consumerSecret := flags.String("consumer-secret", "", "Twitter Consumer Secret")
	accessToken := flags.String("access-token", "", "Twitter Access Token")
	accessSecret := flags.String("access-secret", "", "Twitter Access Secret")
	flags.Parse(os.Args[1:])
	flagutil.SetFlagsFromEnv(flags, "TWITTER")

	if *consumerKey == "" || *consumerSecret == "" || *accessToken == "" || *accessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	config := oauth1.NewConfig(*consumerKey, *consumerSecret)
	token := oauth1.NewToken(*accessToken, *accessSecret)

	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Search Tweets
	sinceIDstr := ""

	userMap := make(map[string]int)

	fmt.Printf("Procurando por tweets...\n")
	for i := 0; i <= 10; i++ {

		search, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
			Query:      "#fatecdevday21",
			Count:      100,
			Since:      sinceIDstr,
			ResultType: "recent",
		})

		if err != nil {
			fmt.Printf("Error: %v \n", err)
		}

		lastResult := len(search.Statuses)
		for _, v := range search.Statuses {
			userMap[v.User.Name] = userMap[v.User.Name] + 1
		}

		sinceIDstr = search.Statuses[lastResult-1].IDStr
		time.Sleep(6 * time.Second)
	}

	fmt.Printf("Total de usuarios participando: %v \n", len(userMap))

	luckyNumber := sorteio(0, len(userMap))
	contador := 0
	for k, v := range userMap {
		if luckyNumber == contador {
			fmt.Println()
			color.Red("E o vencedor foi...")
			color.Yellow(fmt.Sprintf("%v", k))
			fmt.Println()
		} else {
			fmt.Printf("Usuario: %v - total tweets: %d \n", k, v)
		}

		contador++
	}

}

func sorteio(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return (rand.Intn(max-min+1) + min)
}
