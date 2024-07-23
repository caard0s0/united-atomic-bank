package api

import (
	"database/sql"
	"errors"
	"net/http"

	db "github.com/caard0s0/vanguard-server/internal/database/sqlc"
	"github.com/caard0s0/vanguard-server/pkg/token"
	"github.com/lib/pq"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
	_ "github.com/swaggo/swag/example/celler/model"
)

type createAccountRequest struct {
	Currency string `json:"currency" binding:"required,currency"`
}

// CreateAccount
//
//	@Summary		Create an account
//	@Description	Create an account. The client must create and log in a user before.
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			_	formData	api.createAccountRequest	true	"_"
//	@Success		201	{object}	db.Account
//	@Failure		400	"Account already exists!"
//	@Failure		401	"Unauthorized user!"
//	@Router			/accounts [POST]
func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreateAccountParams{
		Owner:    authPayload.Username,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	successfulCreatedAccounts.WithLabelValues("/accounts", "POST", "201").Inc()
	ctx.JSON(http.StatusCreated, account)
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// GetAccount
//
//	@Summary		Get an account
//	@Description	Get an account.
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		int	true	"Account ID"
//	@Success		200	{object}	db.Account
//	@Router			/accounts/{id} [GET]
func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if account.Owner != authPayload.Username {
		err := errors.New("account doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type listAccountRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// ListAccounts
//
//	@Summary		List accounts
//	@Description	List accounts.
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			page_id		query	int	true	"Page ID"	minimum(1)
//	@Param			page_size	query	int	true	"Page Size"	minimum(5)	maximum(10)
//	@Success		200			{array}	db.Account
//	@Router			/accounts [GET]
func (server *Server) listAccounts(ctx *gin.Context) {
	var req listAccountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.ListAccountsParams{
		Owner:  authPayload.Username,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
