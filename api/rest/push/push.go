package push

import (
    "net/http"

    "bilalekrem.com/pushnotification-service/internal/httputils"
    "bilalekrem.com/pushnotification-service/internal/response"
)

func (router *pushRouter) SendPushNotification(w http.ResponseWriter, r *http.Request) {
    var request PushRequest
    err := httputils.DecodeJSONBody(w, r, &request)
    if err != nil {
        httputils.EncodeJSONBody(w, r, response.NewWithError(err))
        return
    }

    // todo send push
    notification := &request.Notification
    token := request.Token

    err = router.service.Send(notification, token)
    if err != nil {
        httputils.EncodeJSONBody(w, r, response.NewWithError(err))
        return
    }

    httputils.EncodeJSONBody(w, r, response.NewWithSuccess())
}

func (router *pushRouter) SendPushNotificationAsync(w http.ResponseWriter, r *http.Request) {
    var request PushRequest
    err := httputils.DecodeJSONBody(w, r, &request)
    if err != nil {
        httputils.EncodeJSONBody(w, r, response.NewWithError(err))
        return
    }

    // todo send push
    notification := &request.Notification
    token := request.Token

    go router.service.Send(notification, token)
    httputils.EncodeJSONBody(w, r, response.NewWithSuccess())
}