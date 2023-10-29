package utils

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func ConnectStorage() *storage.BucketHandle {
	config := &firebase.Config{
		StorageBucket: "suravent-81e1c.appspot.com",
	}
	opt := option.WithCredentialsFile(os.Getenv("firebase_path_credentials"))
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}

	return bucket
}

func Upload(file *multipart.FileHeader, typeFile string) (string, error) {
	storage := ConnectStorage()

	f, err := file.Open()
	defer f.Close()

	str := fmt.Sprintf("%s-%s", typeFile, RandomString(20))
	o := storage.Object(str)

	wc := o.NewWriter(context.Background())
	if _, err = io.Copy(wc, f); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}

	return str, nil
}
