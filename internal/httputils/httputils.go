// ref: https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body
package httputils

import (
    "encoding/json"
    "errors"
    "fmt"
    "net/http"
    "strings"
    "io"
    "log"

    "github.com/golang/gddo/httputil/header"
    "bilalekrem.com/pushnotification-service/internal/response"
)

type MalformedRequest struct {
    Status int
    Msg    string
}

func (mr *MalformedRequest) Error() string {
    return mr.Msg
}

func EncodeJSONBody(w http.ResponseWriter, r *http.Request, resp *response.Response) {
    var status int
    if resp.Success {
        status = http.StatusOK
    } else {
        status = http.StatusBadRequest
    }

    encodeJSONBody(w, r, resp, status)
}

func EncodeJSONBodyWithStatus(w http.ResponseWriter, r *http.Request, resp *response.Response, status int) {
    encodeJSONBody(w, r, resp, status)
}

func encodeJSONBody(w http.ResponseWriter, r *http.Request, resp *response.Response, status int) {
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(*resp); err != nil {
       // should not happen
       w.WriteHeader(http.StatusInternalServerError)
       log.Println("Internal server error occurred while encoding json body", err)

       json.NewEncoder(w).Encode(response.NewWithFailure())
   }
}

func DecodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
    if r.Header.Get("Content-Type") != "" {
        value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
        if value != "application/json" {
            msg := "Content-Type header is not application/json"
            return &MalformedRequest{Status: http.StatusUnsupportedMediaType, Msg: msg}
        }
    }

    r.Body = http.MaxBytesReader(w, r.Body, 1048576)

    dec := json.NewDecoder(r.Body)
    dec.DisallowUnknownFields()

    err := dec.Decode(&dst)
    if err != nil {
        var syntaxError *json.SyntaxError
        var unmarshalTypeError *json.UnmarshalTypeError

        switch {
        case errors.As(err, &syntaxError):
            msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
            return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

        case errors.Is(err, io.ErrUnexpectedEOF):
            msg := fmt.Sprintf("Request body contains badly-formed JSON")
            return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

        case errors.As(err, &unmarshalTypeError):
            msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
            return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

        case strings.HasPrefix(err.Error(), "json: unknown field "):
            fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
            msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
            return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

        case errors.Is(err, io.EOF):
            msg := "Request body must not be empty"
            return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

        case err.Error() == "http: request body too large":
            msg := "Request body must not be larger than 1MB"
            return &MalformedRequest{Status: http.StatusRequestEntityTooLarge, Msg: msg}

        default:
            return err
        }
    }

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
        msg := "Request body must only contain a single JSON object"
        return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}
    }

    return nil
}
