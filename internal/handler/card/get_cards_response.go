package card

import "pg/internal/model"

func dataToResponseArray(cards []model.Card) []ACardResponse {
	if len(cards) == 0 {
		return nil
	}
	resp := make([]ACardResponse, len(cards))
	for i, s := range cards {
		resp[i].ID = s.ID
		resp[i].ExpiredDate = s.ExpiredDate
		resp[i].CVV = s.CVV
		resp[i].UserID = s.UserID
		resp[i].CreatedAt = s.CreatedAt
		resp[i].UpdatedAt = s.UpdatedAt
		resp[i].Balance = s.Balance
	}
	return resp
}

type getCardsResponse struct {
	Cards  []ACardResponse `json:"cards"`
	Cursor int64           `json:"cursor"`
}

func toGetCardsResponse(resp []ACardResponse) getCardsResponse {
	if len(resp) == 0 {
		return getCardsResponse{}
	}
	return getCardsResponse{
		Cards:  resp,
		Cursor: resp[len(resp)-1].ID,
	}
}
