package handlers

// @Description Ответ autocomplete эндпоинта воркера.
type AutocompleteResponse struct {
	Q string
	Response []string
}