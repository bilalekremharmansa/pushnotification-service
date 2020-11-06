package push_test

import (
    "testing"

    push "bilalekrem.com/pushnotification-service/internal/push"
)

func Test_init(t *testing.T) {
    serviceAccount := "/tmp/service.json"
    _ = push.NewWithServiceAccount(serviceAccount)
}