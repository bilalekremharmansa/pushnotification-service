package push

type Notification struct {
    Title string `json:"title"`
    Body string `json:"body"`
    Image string `json:"image"`
}

func NewNotification(title string, body string) *Notification {
    return &Notification{Title: title, Body: body}
}

func NewNotificationWithImage(title string, body string, image string) *Notification {
    return &Notification{Title: title, Body: body, Image: image}
}