package firebaseadminsdk

import (
    "testing"

    push "bilalekrem.com/pushnotification-service/internal/push/firebaseadminsdk"
)

func Test_init(t *testing.T) {
    serviceAccount := "/tmp/service.json"
    _ = push.NewWithServiceAccount(serviceAccount)
}