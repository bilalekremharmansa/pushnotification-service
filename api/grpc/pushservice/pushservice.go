package pushservice

import (
	"context"
	"log"

    "bilalekrem.com/pushnotification-service/internal/push"
    "bilalekrem.com/pushnotification-service/internal/push/firebaseadminsdk"

    pb "bilalekrem.com/pushnotification-service/proto/pushservice"
)

type pushService struct {
    pb.UnimplementedPushNotificationServiceServer

    service *firebaseadminsdk.FirebasePushNotificationService
}

func NewService() *pushService {
    return &pushService{service: firebaseadminsdk.GetInstance()}
}

func (ps *pushService) Send(ctx context.Context, in *pb.PushRequest) (*pb.PushResponse, error) {
    log.Printf("Received gRPC send request")

    notification := extractLocalNotificationFromProto(in)
    token := in.Token

    err := ps.service.Send(notification, token)
    if err != nil {
        return nil, err
    }

    return &pb.PushResponse{Success:true}, nil
}

func extractLocalNotificationFromProto(r *pb.PushRequest) *push.Notification {
    notification := r.Notification

    return push.NewNotificationWithImage(notification.Title, notification.Body, notification.Image)
}
