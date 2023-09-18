package storage

import (
	"context"
	"encoding/csv"
	"matching/src/dataframe"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// https://developers.google.com/sheets/api/quickstart/go

type Creds struct {
}

func (c Creds) Retrieve(ctx context.Context) (aws.Credentials, error) {
	return aws.Credentials{
		AccessKeyID:     os.Getenv("AWS_ACCESS_KEY"),
		SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}, nil
}

type IStorage interface {
	GetQuestionsFromStorage(key string) (*dataframe.Dataframe, error)
}

type Storage struct {
	client *s3.Client
}

func NewStorage() *Storage {
	client := s3.New(s3.Options{
		Region:      "us-east-2",
		Credentials: Creds{},
	})

	return &Storage{
		client: client,
	}
}

func (s *Storage) GetQuestionsFromStorage(key string) (*dataframe.Dataframe, error) {
	rawObject, err := s.client.GetObject(
		context.Background(),
		&s3.GetObjectInput{
			Bucket: aws.String(os.Getenv("AWS_BUCKET")),
			Key:    aws.String(key),
		})
	if err != nil {
		return nil, err
	}
	csvReader := csv.NewReader(rawObject.Body)
	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	df, err := dataframe.NewDataframe(records)
	if err != nil {
		return nil, err
	}
	return df, nil
}
