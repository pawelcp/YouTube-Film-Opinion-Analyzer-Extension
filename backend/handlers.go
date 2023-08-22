package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOpinionHandler(c *gin.Context) {
	link := c.Query("link")

	if link == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'link' parameter"})
		return
	}

	videoID, err := extractVideoIDFromLink(link)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error extracting video ID"})
		return
	}

	commentsFormatted := getYtComents(videoID)
	opinion, err := getGptOpinion(commentsFormatted)
	if err != nil {
		log.Fatalf("Failed to send the request: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{"opinion": opinion})
}
