package main

import "fmt"

func main() {
	link := "https://www.youtube.com/watch?v=EmZtTd1YRmA&ab_channel=FORMULA1"
	videoID, err := extractVideoIDFromLink(link)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("YouTube Video ID:", videoID)
	commentsFormatted := getYtComents(videoID)
	getGptOpinion(commentsFormatted)
}
