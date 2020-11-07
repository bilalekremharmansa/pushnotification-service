package mgmt

import (
    "bilalekrem.com/pushnotification-service/api/rest/router"
)

func Routes() []router.Route {
	return []router.Route {
	    router.NewGetRoute("/", getStatus),
		router.NewGetRoute("/mgmt/status", getStatus),
    }
}
