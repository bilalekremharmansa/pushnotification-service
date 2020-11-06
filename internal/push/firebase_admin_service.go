package push

import (
    "context"
    "log"
    "io/ioutil"

    "encoding/json"

    "google.golang.org/api/option"

    firebase "firebase.google.com/go/v4"
    "firebase.google.com/go/v4/messaging"
)

type FirebasePushNotificationService struct {
    firebaseClient *firebase.App
}

func NewWithServiceAccount(serviceAccountFilePath string) *FirebasePushNotificationService {
    log.Printf("Reading file: %s -- ", serviceAccountFilePath)
    serviceAccountAsBytes, err := ioutil.ReadFile(serviceAccountFilePath)

    var serviceAccount map[string]interface{}
    json.Unmarshal(serviceAccountAsBytes, &serviceAccount)

    projectId := serviceAccount["project_id"].(string)
    log.Printf("Initializing Firebase app instance for project [%s]", projectId)

    opt := option.WithCredentialsJSON(serviceAccountAsBytes)
    config := &firebase.Config{ProjectID: projectId}
    app, err := firebase.NewApp(context.Background(), config, opt)
    if err != nil {
        log.Fatalf("error initializing app: %v\n", err)
    }

    log.Printf("Firebase app instance created successfully")

    return &FirebasePushNotificationService{firebaseClient: app}
}

func (service FirebasePushNotificationService) Send(notification *Notification, token string) error {
    ctx := context.Background()
    client, err := service.firebaseClient.Messaging(ctx)
    if err != nil {
        log.Fatalf("error getting Messaging client: %v\n", err)
        return err
    }

    firebaseNotification := &messaging.Notification{Title: notification.Title, Body: notification.Body, ImageURL: notification.Image}

    message := &messaging.Message{
        Notification: firebaseNotification,
        Token: token,
    }

    log.Printf("Sending firebase message")
    response, err := client.Send(ctx, message)
    if err != nil {
        log.Printf("Sending message failed, [%v]", err)
        return err
    }
    // Response is a message ID string.
    log.Println("Successfully sent message:", response)

    return nil
}