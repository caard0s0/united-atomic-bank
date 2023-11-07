package api

import (
	"os"
	"testing"
	"time"

	"github.com/caard0s0/united-atomic-bank-server/configs"
	db "github.com/caard0s0/united-atomic-bank-server/internal/database/sqlc"
	"github.com/caard0s0/united-atomic-bank-server/internal/util"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := configs.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
