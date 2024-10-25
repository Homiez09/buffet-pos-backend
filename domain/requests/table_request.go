package requests

type AddTableRequest struct {
	TableName string `json:"tableName" validate:"required"`
}

type FindTableByIDRequest struct {
	TableID string `json:"tableID" validate:"required"`
}
