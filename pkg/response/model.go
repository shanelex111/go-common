package response

type response struct {
	Status  int    `json:"-"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
