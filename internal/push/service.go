package push

type NotificationService interface {

    Send(notification *Notification, token string) error

}