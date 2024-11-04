package gorm

import (
	"context"

	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InvoiceGormRepository struct {
	DB *gorm.DB
}

func NewInvoiceGormRepository(db *gorm.DB) *InvoiceGormRepository {
	return &InvoiceGormRepository{
		DB: db,
	}
}

func (i *InvoiceGormRepository) Create(ctx context.Context, tableID string, totalPrice float64, peopleAmount int) error {
	id, _ := uuid.NewV7()

	tableIDParse, err := uuid.Parse(tableID)
	if err != nil {
		return err
	}

	invoice := &models.Invoice{
		ID:           id,
		TableID:      &tableIDParse,
		PeopleAmount: peopleAmount,
		TotalPrice:   totalPrice,
		IsPaid:       false,
	}
	result := i.DB.Create(invoice)
	return result.Error
}

func (i *InvoiceGormRepository) SetPaid(ctx context.Context, invoiceID string) error {
	invoiceIDParse, err := uuid.Parse(invoiceID)
	if err != nil {
		return err
	}
	result := i.DB.Model(&models.Invoice{}).Where("id = ?", invoiceIDParse).Update("is_paid", true)
	return result.Error
}
