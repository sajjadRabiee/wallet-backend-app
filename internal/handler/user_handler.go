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
	wallet, cards, err := h.walletService.GetCardWalletByUserId(input)
	if err != nil {
		response := utils.ErrorResponse("show profile failed", http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formattedUser := dto.FormatUserDetail(user, wallet, cards)
	response := utils.SuccessResponse("show profile success", http.StatusOK, formattedUser)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) Cards(c *gin.Context) {
	user := c.MustGet("user").(*model.User)
	input := &dto.CardsBody{}
	err := c.ShouldBindJSON(input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		response := utils.ErrorResponse("transfer failed", http.StatusUnprocessableEntity, errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	card, err := h.userService.SaveCard(input.CardNumber, user.ID)
	if err != nil {
		return
	}
	response := utils.SuccessResponse("show new card", http.StatusOK, card)
	c.JSON(http.StatusOK, response)
}
