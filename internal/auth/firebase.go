package auth

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)


var FirebaseAuth *auth.Client

func InitFirebase() error {
    opt := option.WithCredentialsFile("protect/firebase-service-account.json")


	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil{
		log.Fatalf("Failed to initaialize firebase: %v", err)
	}

	FirebaseAuth, err = app.Auth(context.Background())
	if err != nil{
		log.Fatalf("Failed to initialize Firebase Auth: %v", err)
	}

	return nil
}

func VerifyIdToken(idToken string) (*auth.Token, error){
	return FirebaseAuth.VerifyIDToken(context.Background(), idToken)
}

 