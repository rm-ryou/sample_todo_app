package request

type Room struct {
	Name string `json:"name" validate:"required,max=50"`
}
