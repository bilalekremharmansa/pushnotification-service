package mgmt

import (
    "net/http"
//     "log"

    "bilalekrem.com/pushnotification-service/internal/httputil"
    "bilalekrem.com/pushnotification-service/internal/response"
)

func getStatus(w http.ResponseWriter, r *http.Request) {
    httputil.EncodeJSONBody(w, r, response.New("OK"))
}