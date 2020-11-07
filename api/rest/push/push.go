package push

import (
    "net/http"

    "bilalekrem.com/pushnotification-service/internal/push/firebaseadminsdk"
    "bilalekrem.com/pushnotification-service/internal/httputils"
    "bilalekrem.com/pushnotification-service/internal/response"
)

var (
    service = firebaseadminsdk.NewWithServiceAccount("/tmp/service.json")
)

func SendPushNotification(w http.ResponseWriter, r *http.Request) {
    var request PushRequest
    err := httputils.DecodeJSONBody(w, r, &request)

    // todo send push
    notification := &request.Notification
    token := request.Token

    err = service.Send(notification, token)
    if err != nil {
        httputils.EncodeJSONBody(w, r, response.NewWithFailureMessage("Sending PushNotification failed"))
        return
    }

    httputils.EncodeJSONBody(w, r, response.NewWithSuccess())
}