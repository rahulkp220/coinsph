package accounts

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	CreateAccountRequest struct {
		Username string `json:"username"`
		Currency string `json:"currency"`
		Balance  string `json:"balance"`
	}
	CreateAccountResponse struct {
		Ok string `json:"ok"`
	}
	GetAccountRequest struct {
		Id string `json:"id"`
	}
	GetAccountResponse struct {
		Username string `json:"username"`
	}
	GetAccountsRequest  struct{}
	GetAccountsResponse struct {
		Result []Account `json:"result"`
	}
	DeleteAccountRequest struct {
		Id string `json:"id"`
	}
	DeleteAccountResponse struct {
		Ok string `json:"ok"`
	}
	TransferRequest struct {
		Sender   string  `json:"sender"`
		Reciever string  `json:"reciever"`
		Amount   float64 `json:"amount"`
	}
	TransferResponse struct {
		Ok string `json:"ok"`
	}
	GetTransfersRequest  struct{}
	GetTransfersResponse struct {
		Result []Payment `json:"result"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeAccountReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateAccountRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeUsernameReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetAccountRequest
	vars := mux.Vars(r)

	req = GetAccountRequest{
		Id: vars["id"],
	}
	return req, nil
}

func decodeAccountsReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetAccountsRequest
	var err error
	return req, err
}

func decodeDeleteReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req DeleteAccountRequest
	vars := mux.Vars(r)

	req = DeleteAccountRequest{
		Id: vars["id"],
	}
	return req, nil
}

func decodeTransferReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req TransferRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeTransfersReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetTransfersRequest
	var err error
	return req, err
}
