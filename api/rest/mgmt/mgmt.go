package mgmt

import (
    "net/http"
//     "log"

    "bilalekrem.com/pushnotification-service/internal/httputils"
    "bilalekrem.com/pushnotification-service/internal/response"
)

func getStatus(w http.ResponseWriter, r *http.Request) {
    httputils.EncodeJSONBody(w, r, response.New("OK"))
}