package response

type Response struct {
    Data interface{} `json:"data"`
    Success bool `json:"success"`
}

func New(data interface{}) *Response {
    return &Response{
        Data: data,
        Success: true,
    }
}

func NewWithSuccess() *Response {
    return &Response{
        Data: nil,
        Success: true,
    }
}

func NewWithFailure() *Response {
    return &Response{
        Data: nil,
        Success: false,
    }
}

func NewWithFailureMessage(data interface{}) *Response {
    return &Response{
        Data: data,
        Success: false,
    }
}