package middleware

import (
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

const (
	mainWalletID       = "main_wallet_id"
	mainWalletIDEnvKEY = "MAIN_WALLET_ID"
	internalErrorCode  = 500
)

func SetMainWallet() gin.HandlerFunc {
	return func(c *gin.Context) {
		walletID := os.Getenv(mainWalletIDEnvKEY)
		if len(walletID) == 0 {
			c.AbortWithStatus(internalErrorCode)
			return
		}
		numWalletID, err := strconv.Atoi(walletID)
		if err != nil {
			c.AbortWithStatus(internalErrorCode)
			return
		}
		c.Set(mainWalletID, numWalletID)
		c.Next()
	}
}
