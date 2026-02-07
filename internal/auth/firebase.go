package auth

import (
	"context"
	"encoding/base64"
	"errors"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var FirebaseAuth *auth.Client

func InitFirebase(firebaseCredBase64 string) error {
	if firebaseCredBase64 == "" {
		return errors.New("FIREBASE_CREDENTIALS_BASE64 not set")
	}

	decodedCreds, err := base64.StdEncoding.DecodeString(firebaseCredBase64)
	if err != nil {
		return err
	}

	opt := option.WithCredentialsJSON(decodedCreds)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}

	FirebaseAuth, err = app.Auth(context.Background())
	if err != nil {
		return err
	}

	log.Println("Firebase initialized successfully")
	return nil
}

func VerifyIdToken(idToken string) (*auth.Token, error) {
	if FirebaseAuth == nil {
		return nil, errors.New("firebase not initialized")
	}
	return FirebaseAuth.VerifyIDToken(context.Background(), idToken)
}
