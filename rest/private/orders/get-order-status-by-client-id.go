package orders

import (
	"github.com/google/go-querystring/query"
	"net/http"
)

type RequestForOpenOrderByClientID struct {
	ClientID string `url:"client_id,omitempty"`
}

func (r *RequestForOpenOrderByClientID) Path() string {
	return "/orders/by_client_id/" + r.ClientID
}

func (r *RequestForOpenOrderByClientID) Method() string {
	return http.MethodGet
}

func (r *RequestForOpenOrderByClientID) Query() string {
	value, _ := query.Values(r)
	return value.Encode()
}

func (r *RequestForOpenOrderByClientID) Payload() []byte {
	return nil
}

type ResponseForOrderStatusByClientID struct {

}
