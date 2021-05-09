package accounts

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("/accounts").Handler(httptransport.NewServer(
		endpoints.CreateAccount,
		decodeAccountReq,
		encodeResponse,
	))

	r.Methods("GET").Path("/accounts/{id}").Handler(httptransport.NewServer(
		endpoints.GetAccount,
		decodeUsernameReq,
		encodeResponse,
	))

	r.Methods("GET").Path("/accounts").Handler(httptransport.NewServer(
		endpoints.GetAccounts,
		decodeAccountsReq,
		encodeResponse,
	))

	r.Methods("DELETE").Path("/accounts/{id}").Handler(httptransport.NewServer(
		endpoints.DeleteAccount,
		decodeDeleteReq,
		encodeResponse,
	))

	r.Methods("POST").Path("/payments").Handler(httptransport.NewServer(
		endpoints.Transfer,
		decodeTransferReq,
		encodeResponse,
	))

	r.Methods("GET").Path("/payments").Handler(httptransport.NewServer(
		endpoints.GetTransfers,
		decodeTransfersReq,
		encodeResponse,
	))

	return r

}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
