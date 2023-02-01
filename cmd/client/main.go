package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"test-service-grpc/client"
	pb "test-service-grpc/pb"
)

// We allow ourselves to hardcode the port since it's for testing purpose
const (
	address = ":8080"
)

func main() {
	listFlag := flag.Bool("list", false, "list will retrieve top stories")
	whoIsFlag := flag.String("whois", "", "user will retrieve its related informations")
	flag.Parse()
	ctx := context.Background()
	grpcMethods := client.Connect(address)

	// who has the priority between list and whois command if both are triggered :
	// `go run main.go -list -whois fra` should be valid, is it ok ?
	if *listFlag {
		req := &pb.TopStoriesRequest{}
		res, err := grpcMethods.GetTopStories(ctx, req)
		if err != nil {
			log.Println("Couldn't retrieve top stories, Reason: ", err)
		}

		// handmade pretty printer typically
		// -> one blank line after command line
		// -> one blank before the prompt at the end
		for _, v := range res.Stories {
			fmt.Println("\n- " + v.Title)
			fmt.Println("  " + v.Url)
		}
		fmt.Println()
		return
	}
	if *whoIsFlag != "" {
		req := &pb.WhoisRequest{User: *whoIsFlag}
		res, err := grpcMethods.Whois(ctx, req)
		if err != nil {
			log.Println("Couldn't retrieve user, Reason: ", err)
		}
		fmt.Println("User: \t", res.Nick)
		fmt.Println("Karma: \t", res.Karma)
		fmt.Println("About: \t", res.About)
		fmt.Println("Joined: \t", res.JoinedAt)
		return
	}
}
