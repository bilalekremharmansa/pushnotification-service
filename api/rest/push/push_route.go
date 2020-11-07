package push

import (
    "bilalekrem.com/pushnotification-service/api/rest/router"
)

func Routes() []router.Route {
	return []router.Route {
		router.NewPostRoute("/push", SendPushNotification),
    }
}

