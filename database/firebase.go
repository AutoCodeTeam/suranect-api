package database

import (
	"context"
	"log"
	"os"
	"sync"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type Model []map[string]interface{}

type FirebaseResponse struct {
	sync.Mutex
	Client     *firestore.Client
	Collection *firestore.CollectionRef
	Ctx        context.Context
}

func ConnectFirebase(collection string) FirebaseResponse {
	var responseFirebase FirebaseResponse

	ctx := context.Background()
	sa := option.WithCredentialsFile(os.Getenv("firebase_path_credentials"))
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	responseFirebase.Client = client
	responseFirebase.Collection = client.Collection(collection)
	responseFirebase.Ctx = ctx

	return responseFirebase
}

func (response FirebaseResponse) Get() Model {
	docs, err := response.Collection.Documents(response.Ctx).GetAll()

	if err != nil {
		log.Fatalln(err)
	}

	var data Model
	for _, ds := range docs {
		data = append(data, ds.Data())
	}

	defer response.Client.Close()
	return data
}

func (response FirebaseResponse) Update(updatedData interface{}) *firestore.WriteResult {
	docs, err := response.Collection.Doc("RYCwtQRP2b8DX9M6uHtR").Update(response.Ctx, []firestore.Update{
		{
			Path:  "currentTicket",
			Value: 0,
		},
	})

	if err != nil {
		log.Fatalln(err)
	}

	defer response.Client.Close()
	return docs
}
