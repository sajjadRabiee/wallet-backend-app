package dto

import "wallet/internal/model"

type CardResponse struct {
	CardNumber string `json:"card_number"`
}

func mapCardsModelToCardResponse(cards []model.Card) []CardResponse {
	res := make([]CardResponse, len(cards))
	for i, c := range cards {
		res[i] = CardResponse{
			CardNumber: c.CardNumber,
		}
	}
	return res
}
