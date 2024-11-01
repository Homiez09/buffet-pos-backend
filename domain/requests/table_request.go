package requests

type AddTableRequest struct {
	TableName string `json:"tableName" validate:"required"`
}
