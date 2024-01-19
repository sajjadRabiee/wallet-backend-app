package dto

import "wallet/internal/model"

type CardResponse struct {
	CardNumber string `json:"card_number"`
}

func mapCardsModelToCardResponse(cards []model.Card) []CardResponse {
	res := make([]CardResponse, len(cards))
	for i, c := range cards {
		var cardNumber = c.CardNumber
		if len(c.CardNumber) > 16 {
			cardNumber = c.CardNumber[:16]
		}
		res[i] = CardResponse{
			CardNumber: cardNumber,
		}
	}
	return res
}
