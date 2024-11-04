package requests

type EditPricePerPerson struct {
	Price float64 `json:"price" validate:"required"`
}
