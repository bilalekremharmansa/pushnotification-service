package push

import "bilalekrem.com/pushnotification-service/internal/push"

type PushRequest struct {
    Notification push.Notification `json:"notification"`
    Token string `json:"token"`
}