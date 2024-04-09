package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vivekbnwork/bz-backend/bz-main/service"
)

type AWSController struct {
	awsService *service.AWSService
}
type FileNamesPayload struct {
	FileNames []string `json:"fileNames"`
}
type DownloadRequest struct {
	FileKeys []string `json:"fileKeys"`
}

// Constructor function for BeerController
func NewAWSController(s *service.AWSService) *AWSController {
	return &AWSController{awsService: s}
}

func (uc *AWSController) GetPresignedURLsForDownload(c *gin.Context) {
	bucket := os.Getenv("AWS_S3_SECURE_BUCKET_NAME")

	var downloadRequest DownloadRequest
	if err := c.ShouldBindJSON(&downloadRequest); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request format"})
		return
	}

	urls, err := uc.awsService.GetPresignedURLsForDownload(bucket, downloadRequest.FileKeys)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"urls": urls})
}

func (uc *AWSController) GetPresignedURLsForUpload(c *gin.Context) {
	var payload FileNamesPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	fmt.Println("fileNames == ", payload.FileNames)

	urls, err := uc.awsService.GetPresignedURLsForUpload(payload.FileNames)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"urls": urls})
}

func (uc *AWSController) DeleteFilesFromS3(c *gin.Context) {
	var payload FileNamesPayload

	// Bind the JSON payload to the payload struct.
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Delete the files from S3.
	err := uc.awsService.DeleteObjectsFromS3(payload.FileNames)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If everything went well, return a success status.
	c.JSON(http.StatusOK, gin.H{"message": "Files deleted successfully"})
}
