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
	FindAllTables(ctx context.Context) ([]responses.TableDetail, error)
	FindTableByID(ctx context.Context, tableID string) (*responses.TableDetail, error)
	EditTable(ctx context.Context, req *requests.EditTableRequest) error
	DeleteTable(ctx context.Context, tableID string) error
	FindByAccessCode(ctx context.Context, accessCode string) (*responses.TableDetail, error)
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

func (t *tableService) FindAllTables(ctx context.Context) ([]responses.TableDetail, error) {
	tables, err := t.tableRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	tableDetails := make([]responses.TableDetail, 0)
	for _, table := range tables {
		tableDetails = append(tableDetails, responses.TableDetail{
			BaseTable: responses.BaseTable{
				ID:          table.ID,
				TableName:   table.TableName,
				IsAvailable: table.IsAvailable,
				QRCode:      table.QRCode,
				AccessCode:  table.AccessCode,
				CreatedAt:   table.CreatedAt,
				UpdatedAt:   table.UpdatedAt,
			},
		})
	}
	return tableDetails, nil
}

func (t *tableService) FindTableByID(ctx context.Context, tableID string) (*responses.TableDetail, error) {
	table, err := t.tableRepo.FindByID(ctx, tableID)
	if err != nil {
		return nil, err
	}

	if table == nil {
		return nil, exceptions.ErrTableNotFound
	}

	return &responses.TableDetail{
		BaseTable: responses.BaseTable{
			ID:          table.ID,
			TableName:   table.TableName,
			IsAvailable: table.IsAvailable,
			QRCode:      table.QRCode,
			AccessCode:  table.AccessCode,
			CreatedAt:   table.CreatedAt,
			UpdatedAt:   table.UpdatedAt,
		},
	}, nil
}

func (t *tableService) EditTable(ctx context.Context, req *requests.EditTableRequest) error {
	table, err := t.tableRepo.FindByID(ctx, req.ID)
	if err != nil {
		return err
	}
	if table == nil {
		return exceptions.ErrTableNotFound
	}
	return t.tableRepo.Edit(ctx, req)
}

func (t *tableService) DeleteTable(ctx context.Context, tableID string) error {
	table, err := t.tableRepo.FindByID(ctx, tableID)
	if err != nil {
		return err
	}
	if table == nil {
		return exceptions.ErrTableNotFound
	}
	return t.tableRepo.Delete(ctx, tableID)
}

func (t *tableService) FindByAccessCode(ctx context.Context, accessCode string) (*responses.TableDetail, error) {
	table, err := t.tableRepo.FindByAccessCode(ctx, accessCode)
	if err != nil {
		return nil, err
	}
	if table == nil {
		return nil, exceptions.ErrTableNotFound
	}
	return &responses.TableDetail{
		BaseTable: responses.BaseTable{
			ID:          table.ID,
			TableName:   table.TableName,
			IsAvailable: table.IsAvailable,
			QRCode:      table.QRCode,
			AccessCode:  table.AccessCode,
			CreatedAt:   table.CreatedAt,
			UpdatedAt:   table.UpdatedAt,
		},
	}, nil
}
