package responseModels

type MessageResponse struct {
	Data struct {
		Message string `json:"data.message"`
	} `json:"data"`
}
