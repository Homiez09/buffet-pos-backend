package gorm

import (
	"context"
	"time"

	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TableGormRepository struct {
	DB *gorm.DB
}

func NewTableGormRepository(db *gorm.DB) *TableGormRepository {
	return &TableGormRepository{
		DB: db,
	}
}

func (t *TableGormRepository) Create(ctx context.Context, req *requests.AddTableRequest) error {
	// Generate UUID
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}

	table := &models.Table{
		ID:          id,
		TableName:   req.TableName,
		IsAvailable: true,
		QRCode:      nil,
		AccessCode:  nil,
		Invoices:    nil,
	}

	result := t.DB.Create(table)

	return result.Error
}

func (t *TableGormRepository) FindAll(ctx context.Context) ([]models.Table, error) {
	var tables []models.Table
	result := t.DB.Find(&tables)
	if result.Error != nil {
		return nil, result.Error
	}
	return tables, nil
}

func (t *TableGormRepository) FindByID(ctx context.Context, tableID string) (*models.Table, error) {
	var table models.Table
	result := t.DB.Where("id = ?", tableID).First(&table)

	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return &table, nil
}

func (t *TableGormRepository) FindByName(ctx context.Context, tableName string) (*models.Table, error) {
	var table models.Table
	result := t.DB.Where("table_name = ?", tableName).First(&table)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &table, nil
}

func (t *TableGormRepository) Edit(ctx context.Context, req *requests.EditTableRequest) error {
	table, err := t.FindByID(ctx, req.ID)
	if err != nil {
		return err
	}
	if table == nil {
		return nil
	}
	table.TableName = req.TableName
	result := t.DB.Save(table)
	return result.Error
}

func (t *TableGormRepository) Delete(ctx context.Context, tableID string) error {
	table, err := t.FindByID(ctx, tableID)
	if err != nil {
		return err
	}
	if table == nil {
		return nil
	}
	result := t.DB.Delete(table)
	return result.Error
}

func (t *TableGormRepository) FindByAccessCode(ctx context.Context, accessCode string) (*models.Table, error) {
	var table models.Table
	result := t.DB.Where("access_code = ?", accessCode).First(&table)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &table, nil
}

func (t *TableGormRepository) Assign(ctx context.Context, tableID string, accessCode string, qrCode string) error {
	table, err := t.FindByID(ctx, tableID)
	if err != nil {
		return err
	}
	if table == nil {
		return nil
	}

	now := time.Now()

	table.IsAvailable = false
	table.AccessCode = &accessCode
	table.QRCode = &qrCode
	table.EntryAt = &now
	result := t.DB.Save(table)
	return result.Error
}

func (t *TableGormRepository) SetAvailability(ctx context.Context, tableID string, availability bool) error {
	table, err := t.FindByID(ctx, tableID)
	if err != nil {
		return err
	}
	if table == nil {
		return nil
	}
	table.IsAvailable = availability
	if availability {
		table.AccessCode = nil
		table.QRCode = nil
		table.EntryAt = nil
	}
	result := t.DB.Save(table)
	return result.Error
}
