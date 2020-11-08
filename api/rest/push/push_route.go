package push

import (
    "bilalekrem.com/pushnotification-service/internal/config"
    "bilalekrem.com/pushnotification-service/internal/push/firebaseadminsdk"
    "bilalekrem.com/pushnotification-service/api/rest/router"
)

type pushRouter struct {
    service *firebaseadminsdk.FirebasePushNotificationService
}

func NewRouter() pushRouter {
    serviceAccountFile := config.GetConfig().FirebaseConfig.ServiceAccountFile
    service := firebaseadminsdk.NewWithServiceAccount(serviceAccountFile)

    return pushRouter{service: service}
}

func (r *pushRouter) Routes() []router.Route {
	return []router.Route {
		router.NewPostRoute("/push", r.SendPushNotification),
		router.NewPostRoute("/pushAsync", r.SendPushNotificationAsync),
    }
}

