package dto

type Tweet struct {
	Content string `json:"content" validate:"required,max=255"`
}

func (t *Tweet) Validate() error {
	return validate.Struct(t)
}
