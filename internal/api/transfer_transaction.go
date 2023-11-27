package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	db "github.com/caard0s0/united-atomic-bank-server/internal/database/sqlc"
	"github.com/caard0s0/united-atomic-bank-server/internal/email"
	"github.com/caard0s0/united-atomic-bank-server/pkg/token"

	"github.com/gin-gonic/gin"
)

type TransferTransactionRequest struct {
	FromAccountID    int64  `json:"from_account_id" binding:"required,min=1"`
	FromAccountOwner string `json:"from_account_owner" binding:"required"`
	ToAccountID      int64  `json:"to_account_id" binding:"required,min=1"`
	ToAccountOwner   string `json:"to_account_owner" binding:"required"`
	Amount           int64  `json:"amount" binding:"required,gt=0"`
	Currency         string `json:"currency" binding:"required,currency"`
}

// CreateTransfer
//
//	@Summary		Create a transfer
//	@Description	Create a transfer.
//	@Tags			transfers
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			_	formData	api.TransferTransactionRequest	true	"_"
//	@Success		201	{object}	db.TransferTransactionResult
//	@Router			/transfers [POST]
func (server *Server) createTransfer(ctx *gin.Context) {
	var req TransferTransactionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fromAccount, valid := server.validAccount(ctx, req.FromAccountID, req.Currency)
	if !valid {
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if fromAccount.Owner != authPayload.Username {
		err := errors.New("from account doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	_, valid = server.validAccount(ctx, req.ToAccountID, req.Currency)
	if !valid {
		return
	}

	arg := db.CreateTransferParams{
		FromAccountID:    req.FromAccountID,
		FromAccountOwner: req.FromAccountOwner,
		ToAccountID:      req.ToAccountID,
		ToAccountOwner:   req.ToAccountOwner,
		Amount:           req.Amount,
	}

	result, err := server.store.TransferTransaction(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	user, _ := server.getUser(ctx, req.FromAccountOwner)
	email.SendEmailWithSuccessfulTransfer(result, user.Email)

	ctx.JSON(http.StatusCreated, result)
}

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string) (db.Account, bool) {
	account, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return account, false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return account, false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch %s vs %s", accountID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return account, false
	}

	return account, true
}

func (server *Server) getUser(ctx *gin.Context, username string) (db.User, error) {
	user, err := server.store.GetUser(ctx, username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return user, err
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return user, err
	}

	return user, nil
}

type listTransferRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=50"`
}

// ListTransfers
//
//	@Summary		List transfers
//	@Description	List transfers.
//	@Tags			transfers
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			page_id		query	int	true	"Page ID"	minimum(1)
//	@Param			page_size	query	int	true	"Page Size"	minimum(5)	maximum(10)
//	@Success		200			{array}	db.Transfer
//	@Router			/transfers [GET]
func (server *Server) listTransfers(ctx *gin.Context) {
	var req listTransferRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.ListTransfersParams{
		FromAccountOwner: authPayload.Username,
		ToAccountOwner:   authPayload.Username,
		Limit:            req.PageSize,
		Offset:           (req.PageID - 1) * req.PageSize,
	}

	transfers, err := server.store.ListTransfers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, transfers)
}
