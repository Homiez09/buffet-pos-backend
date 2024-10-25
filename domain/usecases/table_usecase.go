package usecases

import (
	"context"

	"github.com/cs471-buffetpos/buffet-pos-backend/configs"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/exceptions"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/repositories"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/responses"
)

type TableUseCase interface {
	AddTable(ctx context.Context, req *requests.AddTableRequest) error
	FindTableByID(ctx context.Context, tableID string) (*responses.FindTableResponse, error)
}

type tableService struct {
	tableRepo repositories.TableRepository
	config    *configs.Config
}

func NewTableService(tableRepo repositories.TableRepository, config *configs.Config) TableUseCase {
	return &tableService{
		tableRepo: tableRepo,
		config:    config,
	}
}

func (t *tableService) AddTable(ctx context.Context, req *requests.AddTableRequest) error {
	table, err := t.tableRepo.FindByName(ctx, req.TableName)
	if err != nil {
		return err
	}

	if table != nil {
		return exceptions.ErrDuplicatedTableName
	}

	return t.tableRepo.Create(ctx, req)
}

func (t *tableService) FindTableByID(ctx context.Context, tableID string) (*responses.FindTableResponse, error) {
	table, err := t.tableRepo.FindByID(ctx, tableID)
	if err != nil {
		return nil, err
	}

	if table == nil {
		return nil, exceptions.ErrTableNotFound
	}

	return &responses.FindTableResponse{
		ID:          table.ID,
		TableName:   table.TableName,
		IsAvailable: table.IsAvailable,
		QRCode:      table.QRCode,
		AccessCode:  table.AccessCode,
		CreatedAt:   table.CreatedAt,
		UpdatedAt:   table.UpdatedAt,
	}, nil
}
