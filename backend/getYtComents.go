package main

import (
	"fmt"
	"log"
	"net/http"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

func getYtComents() []string {
	apiKey := "AIzaSyAl_OcorffLEvm6Itoz8kvBmjjd4qQhISY"
	videoID := "DmEx49-Pc0Y"

	client := &http.Client{
		Transport: &transport.APIKey{Key: apiKey},
	}
	youtubeService, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	call := youtubeService.CommentThreads.List([]string{"snippet"}).
		VideoId(videoID).
		MaxResults(35)

	comments, err := call.Do()
	if err != nil {
		log.Fatalf("Error fetching comments: %v", err)
	}

	commentsFormatted := make([]string, 0)

	for i, comment := range comments.Items {
		commentFormatted := fmt.Sprintf("Comment %d: %s", i+1, comment.Snippet.TopLevelComment.Snippet.TextOriginal)
		commentsFormatted = append(commentsFormatted, commentFormatted)
	}

	return commentsFormatted
}
