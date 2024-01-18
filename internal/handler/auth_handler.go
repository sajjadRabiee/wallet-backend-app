package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"wallet/internal/dto"
	"wallet/pkg/utils"
)

const loginFailedMessage = "login failed"
const registerFailedMessage = "register failed"
const forgotPasswordFailedMessage = "forgot password failed"
const resetPasswordFailedMessage = "reset password failed"

func (h *Handler) Register(c *gin.Context) {
	input := &dto.RegisterRequestBody{}

	err := c.ShouldBindJSON(input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		response := utils.ErrorResponse(registerFailedMessage, http.StatusUnprocessableEntity, errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.CreateUser(input)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse(registerFailedMessage, statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	newWallet, err := h.walletService.CreateWallet(newUser)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse(registerFailedMessage, statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	token, err := h.jwtService.GenerateToken(int(newUser.ID))
	if err != nil {
		response := utils.ErrorResponse(registerFailedMessage, http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formattedLogin := dto.FormatLogin(newUser, newWallet, token)
	response := utils.SuccessResponse("register success", http.StatusOK, formattedLogin)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) Login(c *gin.Context) {
	input := &dto.LoginRequestBody{}

	err := c.ShouldBindJSON(input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		response := utils.ErrorResponse(loginFailedMessage, http.StatusUnprocessableEntity, errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.authService.Attempt(input)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse(loginFailedMessage, statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	inputWallet := &dto.WalletRequestBody{}
	inputWallet.UserID = int(loggedinUser.ID)
	wallet, err := h.walletService.GetWalletByUserId(inputWallet)
	if err != nil {
		response := utils.ErrorResponse(loginFailedMessage, http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	token, err := h.jwtService.GenerateToken(int(loggedinUser.ID))
	if err != nil {
		response := utils.ErrorResponse(loginFailedMessage, http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formattedLogin := dto.FormatLogin(loggedinUser, wallet, token)
	response := utils.SuccessResponse("login success", http.StatusOK, formattedLogin)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) ForgotPassword(c *gin.Context) {
	input := &dto.ForgotPasswordRequestBody{}

	err := c.ShouldBindJSON(input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		response := utils.ErrorResponse(forgotPasswordFailedMessage, http.StatusUnprocessableEntity, errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	passwordReset, err := h.authService.ForgotPass(input)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse(forgotPasswordFailedMessage, statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	formattedPasswordReset := dto.FormatForgotPassword(passwordReset)
	response := utils.SuccessResponse("forgot password success", http.StatusOK, formattedPasswordReset)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) ResetPassword(c *gin.Context) {
	input := &dto.ResetPasswordRequestBody{}

	err := c.ShouldBindJSON(input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		response := utils.ErrorResponse(resetPasswordFailedMessage, http.StatusUnprocessableEntity, errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	passwordReset, err := h.authService.ResetPass(input)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse(resetPasswordFailedMessage, statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	formattedUser := dto.FormatUser(&passwordReset.User)
	response := utils.SuccessResponse("reset password success", http.StatusOK, formattedUser)
	c.JSON(http.StatusOK, response)
}
