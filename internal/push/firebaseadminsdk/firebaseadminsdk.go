package firebaseadminsdk

import (
    "context"
    "log"
    "io/ioutil"

    "encoding/json"

    "google.golang.org/api/option"

    firebase "firebase.google.com/go/v4"
    "firebase.google.com/go/v4/messaging"

    "bilalekrem.com/pushnotification-service/internal/push"
    "bilalekrem.com/pushnotification-service/internal/config"
)

type FirebasePushNotificationService struct {
    firebaseClient *firebase.App
}

var (
    // it is designed to use as a single service account per process
    instance *FirebasePushNotificationService
)

// this function returns a *FirebasePushNotificationService instance and uses service account in AppConfig. This instance
// will be served as a singleton instance for current feature set
func GetInstance() *FirebasePushNotificationService {
    if instance != nil {
        return instance
    }

    serviceAccountFile := config.GetAppConfig().GetFirebaseConfig().GetServiceAccountFile()
    return NewWithServiceAccount(serviceAccountFile)
}

func NewWithServiceAccount(serviceAccountFilePath string) *FirebasePushNotificationService {
    log.Printf("Reading file: [%s]", serviceAccountFilePath)
    serviceAccountAsBytes, err := ioutil.ReadFile(serviceAccountFilePath)
    if err != nil {
        log.Fatalf("Reading service account failed: [%v]", err)
    }

    var serviceAccount map[string]interface{}
    err = json.Unmarshal(serviceAccountAsBytes, &serviceAccount)
    if err != nil {
        log.Fatalf("Service account file could not be parsed: [%v]", err)
    }

    projectId := serviceAccount["project_id"].(string)
    log.Printf("Initializing Firebase app instance for project [%s]", projectId)

    opt := option.WithCredentialsJSON(serviceAccountAsBytes)
    config := &firebase.Config{ProjectID: projectId}
    app, err := firebase.NewApp(context.Background(), config, opt)
    if err != nil {
        log.Fatalf("Initializing Firebase app instance failed: [%v]", err)
    }

    log.Printf("Firebase app instance created successfully")
    instance = &FirebasePushNotificationService{firebaseClient: app}

    return instance
}

func (service FirebasePushNotificationService) Send(notification *push.Notification, token string) error {
    ctx := context.Background()
    client, err := service.firebaseClient.Messaging(ctx)
    if err != nil {
        log.Printf("Getting messaging client failed: [%v]", err)
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