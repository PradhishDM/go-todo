package auth

import (
	"context"
	"errors"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var FirebaseAuth *auth.Client

func InitFirebase() (*firebase.App, error) {

	// STEP 1: Read JSON from environment variable
	firebaseJSON := os.Getenv("FIREBASE_SERVICE_ACCOUNT_JSON")
	if firebaseJSON == "" {
		log.Fatal("FIREBASE_SERVICE_ACCOUNT_JSON is not set")
	}

	// STEP 2: Pass JSON directly to Firebase SDK
	opt := option.WithCredentialsJSON([]byte(firebaseJSON))

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	FirebaseAuth, err = app.Auth(context.Background())
	if err != nil {
		return nil, err
	}

	return app, nil
}

func VerifyIdToken(idToken string) (*auth.Token, error) {
	if FirebaseAuth == nil {
		return nil, errors.New("firebase not initialized")
	}
	return FirebaseAuth.VerifyIDToken(context.Background(), idToken)
}
