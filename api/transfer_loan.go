package api

import (
	db "atomic-bank/db/sqlc"
	"atomic-bank/token"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TransferLoanRequest struct {
	AccountID  int64 `json:"account_id" binding:"required,min=1"`
	LoanAmount int64 `json:"loan_amount" binding:"required,gt=0"`
}

func (server *Server) createLoan(ctx *gin.Context) {
	var req TransferLoanRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	toAccount, valid := server.validAccountLoan(ctx, req.AccountID)
	if !valid {
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if toAccount.Owner != authPayload.Username {
		err := errors.New("account doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	arg := db.TransferLoanParams{
		AccountID:    req.AccountID,
		LoanAmount:   req.LoanAmount,
		InterestRate: 1,
		Status:       "Active",
		EndDate:      time.Now().Add(time.Minute).Truncate(time.Second),
	}

	loan, err := server.store.TransferLoan(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, loan)
}

func (server *Server) validAccountLoan(ctx *gin.Context, accountID int64) (db.Account, bool) {
	account, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return account, false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return account, false
	}

	return account, true
}
