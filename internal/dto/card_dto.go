package dto

import "wallet/internal/model"

type CardResponse struct {
	CardNumber string `json:"card_number"`
}

func mapCardsModelToCardResponse(cards []model.Card) []CardResponse {
	res := make([]CardResponse, len(cards))
	for i, c := range cards {
		var cartNumber = c.CardNumber
		if len(c.CardNumber) > 16 {
			cartNumber = c.CardNumber[:15]
		}
		res[i] = CardResponse{
			CardNumber: cartNumber,
		}
	}
	return res
}
