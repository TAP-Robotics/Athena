package shared

type SendTempalte struct {
	Message string      `json:"message"`
	Content interface{} `json:"content"`
}
