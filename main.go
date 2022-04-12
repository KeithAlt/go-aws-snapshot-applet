package main

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

// Initialize our env variables
func initEnv() error {
	log.Println("Initializing environment variables...")

	loadErr := godotenv.Load("./config/.env")
	if loadErr != nil {
		return loadErr
	}

	if os.Getenv("AWS_DB_ID") == "<INSERT_HERE>" || os.Getenv("AWS_CONFIG_REGION") == "<INSERT_HERE>" {
		return errors.New("configurationFailure: your /config/.env file isn't configured; do so then restart this app")
	}

	log.Println("Successfully initialized environment variables")
	return nil
}

// Create our AWS session
func createSession() *session.Session {
	log.Println("Establishing AWS session...")

	Session := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_CONFIG_REGION")),
	}))

	log.Println("Successfully established AWS session")
	return Session
}

// Create our database snapshot
func createSnapshot(session *session.Session) (*rds.CreateDBSnapshotOutput, error) {
	curTime := time.Now()
	dbID := "snapshot-backup-" + curTime.Format("2006-01-02-at-15-04")

	log.Println("Attempting to create new database backup: ", dbID, "...")

	svc := rds.New(session)
	DBInstanceIdentifier := os.Getenv("AWS_DB_ID")
	DBSnapshotIdentifier := dbID
	dbSnapshotReq := rds.CreateDBSnapshotInput{
		DBInstanceIdentifier: &DBInstanceIdentifier,
		DBSnapshotIdentifier: &DBSnapshotIdentifier,
	}

	validateErr := dbSnapshotReq.Validate()
	if validateErr != nil {
		return nil, validateErr
	}

	res, resErr := svc.CreateDBSnapshot(&dbSnapshotReq)
	if resErr != nil {
		return nil, resErr
	}

	log.Println("Successfully created database snapshot ", dbID)
	return res, nil
}

// Our driver
func main() {
	log.Println("Starting remote database backup...")

	initErr := initEnv()
	if initErr != nil {
		log.Fatal(initErr)
	}

	Session := createSession()
	_, resErr := createSnapshot(Session)
	if resErr != nil {
		log.Fatal(resErr)
	}

	log.Println("Completed remote database backup")
}
