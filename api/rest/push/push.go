package push

import (
    "net/http"
    "log"

    "bilalekrem.com/pushnotification-service/internal/push/firebaseadminsdk"
    "bilalekrem.com/pushnotification-service/internal/httputil"
    "bilalekrem.com/pushnotification-service/internal/response"
)

var (
    service = firebaseadminsdk.NewWithServiceAccount("/tmp/service.json")
)

func SendPushNotification(w http.ResponseWriter, r *http.Request) {
    var request PushRequest
    err := httputil.DecodeJSONBody(w, r, &request)

    // todo send push
    notification := &request.Notification
    token := request.Token

    log.Println(notification)
    log.Println(token)
    err = service.Send(notification, token)
    if err != nil {
        httputil.EncodeJSONBody(w, r, response.NewWithFailureMessage("Sending PushNotification failed"))
        return
    }

    httputil.EncodeJSONBody(w, r, response.NewWithSuccess())
}