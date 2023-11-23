package api

import (
	"database/sql"
	"errors"
	"net/http"

	db "github.com/caard0s0/united-atomic-bank-server/internal/database/sqlc"
	"github.com/caard0s0/united-atomic-bank-server/pkg/token"

	"github.com/gin-gonic/gin"
)

type LoanTransferTransactionRequest struct {
	AccountID int64 `json:"account_id" binding:"required,min=1"`
	Amount    int64 `json:"amount" binding:"required,gt=0"`
}

// CreateLoanTransfer
//
//	@Summary		Create a loan
//	@Description	Create a loan.
//	@Tags			loans
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			_	formData	api.LoanTransferTransactionRequest	true	"_"
//	@Success		201	{object}	db.LoanTransferTransactionResult
//	@Router			/loans [POST]
func (server *Server) createLoanTransfer(ctx *gin.Context) {
	var req LoanTransferTransactionRequest
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

	arg := db.CreateLoanTransferParams{
		AccountID: req.AccountID,
		Amount:    req.Amount,
	}

	loan, err := server.store.LoanTransferTransaction(ctx, arg)
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
