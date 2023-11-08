package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"os"

	"github.com/joho/godotenv"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func getYtComents(videoID string) []string {

	loadEnv()

	apiKey := os.Getenv("YOUTUBE_API_KEY")

	client := &http.Client{
		Transport: &transport.APIKey{Key: apiKey},
	}
	youtubeService, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	call := youtubeService.CommentThreads.List([]string{"snippet"}).
		VideoId(videoID).
		MaxResults(50)

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
func extractVideoIDFromLink(link string) (string, error) {
	regex := `^(?:https?:\/\/)?(?:www\.)?(?:youtube\.com\/(?:watch\?v=|embed\/|v\/)|youtu\.be\/)([a-zA-Z0-9_-]{11})`

	re := regexp.MustCompile(regex)

	matches := re.FindStringSubmatch(link)

	if len(matches) >= 2 {
		return matches[1], nil
	}

	return "", fmt.Errorf("Invalid YouTube link: %s", link)
}
