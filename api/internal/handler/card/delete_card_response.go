package card

type deleteOrderResponse struct {
	Message string `json:"message"`
}

func toSuccessDelete() deleteOrderResponse {
	return deleteOrderResponse{Message: "Deleted Card"}
}
