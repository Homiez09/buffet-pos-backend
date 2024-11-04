package usecases

import (
	"context"
	"strconv"

	"github.com/cs471-buffetpos/buffet-pos-backend/configs"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/exceptions"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/repositories"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/responses"
	"github.com/google/uuid"
)

type TableUseCase interface {
	AddTable(ctx context.Context, req *requests.AddTableRequest) error
	FindAllTables(ctx context.Context) ([]responses.TableDetail, error)
	FindTableByID(ctx context.Context, tableID string) (*responses.TableDetail, error)
	EditTable(ctx context.Context, req *requests.EditTableRequest) error
	DeleteTable(ctx context.Context, tableID string) error
	FindByAccessCode(ctx context.Context, accessCode string) (*responses.TableDetail, error)
	AssignTable(ctx context.Context, req *requests.AssignTableRequest) (*responses.TableDetail, error)
}

type tableService struct {
	tableRepo   repositories.TableRepository
	invoiceRepo repositories.InvoiceRepository
	settingRepo repositories.SettingRepository
	config      *configs.Config
}

func NewTableService(tableRepo repositories.TableRepository, invoiceRepo repositories.InvoiceRepository, settingRepo repositories.SettingRepository, config *configs.Config) TableUseCase {
	return &tableService{
		tableRepo:   tableRepo,
		invoiceRepo: invoiceRepo,
		settingRepo: settingRepo,
		config:      config,
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
				EntryAt:     table.EntryAt,
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
			EntryAt:     table.EntryAt,
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
			EntryAt:     table.EntryAt,
			CreatedAt:   table.CreatedAt,
			UpdatedAt:   table.UpdatedAt,
		},
	}, nil
}

func (t *tableService) AssignTable(ctx context.Context, req *requests.AssignTableRequest) (*responses.TableDetail, error) {
	table, err := t.tableRepo.FindByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	if table == nil {
		return nil, exceptions.ErrTableNotFound
	}
	if !table.IsAvailable {
		return nil, exceptions.ErrTableAlreadyAssigned
	}
	tableID := table.ID.String()

	accessCode, _ := uuid.NewV7()
	accessCodeStr := accessCode.String()

	qrcode := t.config.FrontendURL + "/user/" + accessCodeStr

	if err := t.tableRepo.Assign(ctx, tableID, accessCode.String(), qrcode); err != nil {
		return nil, err
	}

	pricePerPerson, err := t.settingRepo.GetSetting(ctx, "pricePerPerson")
	if err != nil {
		return nil, err
	}

	pricePerPersonFloat, _ := strconv.ParseFloat(pricePerPerson.Value, 64)

	totalPrice := float64(req.PeopleAmount) * pricePerPersonFloat
	if err := t.invoiceRepo.Create(ctx, tableID, totalPrice, req.PeopleAmount); err != nil {
		return nil, err
	}

	table, err = t.tableRepo.FindByID(ctx, tableID)

	return &responses.TableDetail{
		BaseTable: responses.BaseTable{
			ID:          table.ID,
			TableName:   table.TableName,
			IsAvailable: table.IsAvailable,
			QRCode:      table.QRCode,
			AccessCode:  table.AccessCode,
			EntryAt:     table.EntryAt,
			CreatedAt:   table.CreatedAt,
			UpdatedAt:   table.UpdatedAt,
		},
	}, nil
}
