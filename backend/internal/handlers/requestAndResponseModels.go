package handlers

// @Description Ответ autocomplete эндпоинта воркера.
type AutocompleteResponse struct {
	Q string
	Response []string
}

type RegisterData struct {
	TgId      string `json:"tg_id"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
}