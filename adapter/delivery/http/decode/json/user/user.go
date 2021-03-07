package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/valonekowd/clean-architecture/infrastructure/router"
	userReq "github.com/valonekowd/clean-architecture/usecase/request/user"
	"github.com/valonekowd/clean-architecture/util/helper"
)

func GetTransactions(_ context.Context, req *http.Request) (interface{}, error) {
	userID, err := helper.RequestParam(router.URLParam(req, "userID")).Int64()
	if err != nil {
		return nil, err
	}

	accountID, err := helper.RequestParam(router.QueryParam(req, "account_id")).Int64()
	if err != nil {
		return nil, err
	}

	return userReq.GetTransactions{
		UserID:    userID,
		AccountID: accountID,
	}, nil
}

func CreateTransaction(_ context.Context, req *http.Request) (interface{}, error) {
	r := userReq.CreateTransaction{}

	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		return nil, err
	}

	r.UserID, err = helper.RequestParam(router.URLParam(req, "userID")).Int64()
	if err != nil {
		return nil, err
	}

	return r, nil
}
