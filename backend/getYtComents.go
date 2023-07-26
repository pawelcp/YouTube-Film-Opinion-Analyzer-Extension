package main

import (
	"fmt"
	"log"
	"net/http"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

func main() {
	apiKey := "AIzaSyAl_OcorffLEvm6Itoz8kvBmjjd4qQhISY"
	videoID := "0CW6cxApOBg"

	client := &http.Client{
		Transport: &transport.APIKey{Key: apiKey},
	}
	youtubeService, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	call := youtubeService.CommentThreads.List([]string{"snippet"}).
		VideoId(videoID).
		MaxResults(100)

	comments, err := call.Do()
	if err != nil {
		log.Fatalf("Error fetching comments: %v", err)
	}

	for i, comment := range comments.Items {
		fmt.Printf("Comment %d: %s\n", i+1, comment.Snippet.TopLevelComment.Snippet.TextOriginal)
	}

}
