package service

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AWSService struct {
	s3Client *s3.S3
}

func NewAWSService(s3Client *s3.S3) *AWSService {
	return &AWSService{
		s3Client: s3Client,
	}
}

func (s *AWSService) GetPresignedURLsForUpload(fileNames []string) ([]string, error) {

	var urls []string
	for _, fileName := range fileNames {
		key := fmt.Sprintf("uploads/%s", fileName)
		req, _ := s.s3Client.PutObjectRequest(&s3.PutObjectInput{
			Bucket: aws.String("brewszilla-secure"),
			Key:    aws.String(key),
		})

		urlStr, err := req.Presign(15 * time.Minute) // URL valid for 15 minutes
		if err != nil {
			return nil, err
		}

		urls = append(urls, urlStr)
	}

	return urls, nil
}

func (s *AWSService) GetPresignedURLsForDownload(bucket string, keys []string) ([]string, error) {

	var urls []string
	for _, key := range keys {
		req, _ := s.s3Client.GetObjectRequest(&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		})

		urlStr, err := req.Presign(15 * time.Minute) // URL valid for 15 minutes
		if err != nil {
			return nil, fmt.Errorf("failed to sign request for key %s: %v", key, err)
		}

		urls = append(urls, urlStr)
	}

	return urls, nil
}

func (s *AWSService) DeleteObjectsFromS3(fileNames []string) error {

	// Prepare the slice of objects to delete
	var objectsToDelete []*s3.ObjectIdentifier
	for _, fileName := range fileNames {
		objectsToDelete = append(objectsToDelete, &s3.ObjectIdentifier{Key: aws.String(fmt.Sprintf("uploads/%s", fileName))})
	}

	input := &s3.DeleteObjectsInput{
		Bucket: aws.String("brewszilla-secure"),
		Delete: &s3.Delete{
			Objects: objectsToDelete,
			Quiet:   aws.Bool(true), // Set to true if you don't want the service to return information about the delete markers
		},
	}

	_, err := s.s3Client.DeleteObjects(input)
	if err != nil {
		return err
	}

	return nil
}
