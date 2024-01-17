package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"wallet/internal/dto"
	"wallet/internal/model"
	"wallet/pkg/utils"
)

func (h *Handler) Profile(c *gin.Context) {
	user := c.MustGet("user").(*model.User)

	input := &dto.WalletRequestBody{}
	input.UserID = int(user.ID)
	wallet, err := h.walletService.GetWalletByUserId(input)
	if err != nil {
		response := utils.ErrorResponse("show profile failed", http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formattedUser := dto.FormatUserDetail(user, wallet)
	response := utils.SuccessResponse("show profile success", http.StatusOK, formattedUser)
	c.JSON(http.StatusOK, response)
}
