package requests

type AddTableRequest struct {
	TableName string `json:"tableName" validate:"required"`
}

type EditTableRequest struct {
	ID        string `json:"id" validate:"required"`
	TableName string `json:"tableName" validate:"required"`
}

type AssignTableRequest struct {
	ID string `json:"id" validate:"required"`
}
