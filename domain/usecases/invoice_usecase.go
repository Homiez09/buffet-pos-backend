package usecases

import (
	"context"

	"github.com/cs471-buffetpos/buffet-pos-backend/configs"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/repositories"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/responses"
)

type InvoiceUseCase interface {
	GetAllUnpaidInvoices(ctx context.Context) ([]responses.InvoiceDetail, error)
	GetAllPaidInvoices(ctx context.Context) ([]responses.InvoiceDetail, error)
	DeleteInvoice(ctx context.Context, invoiceID string) error
	SetPaidInvoice(ctx context.Context, invoiceID string) error
	CustomerGetInvoice(ctx context.Context, tableID string) (responses.InvoiceDetail, error)
}

type InvoiceService struct {
	invoiceRepo repositories.InvoiceRepository
	tableRepo   repositories.TableRepository
	orderRepo   repositories.OrderRepository
	config      *configs.Config
}

func NewInvoiceService(invoiceRepo repositories.InvoiceRepository, tableRepo repositories.TableRepository, orderRepo repositories.OrderRepository, config *configs.Config) InvoiceUseCase {
	return &InvoiceService{
		invoiceRepo: invoiceRepo,
		tableRepo:   tableRepo,
		orderRepo:   orderRepo,
		config:      config,
	}
}

func (i *InvoiceService) GetAllUnpaidInvoices(ctx context.Context) ([]responses.InvoiceDetail, error) {
	invoices, err := i.invoiceRepo.GetAllUnpaid(ctx)
	if err != nil {
		return nil, err
	}
	invoiceDetails := make([]responses.InvoiceDetail, 0)

	for _, invoice := range invoices {
		invoiceDetails = append(invoiceDetails, responses.InvoiceDetail{
			BaseInvoice: responses.BaseInvoice{
				ID:           invoice.ID,
				PeopleAmount: invoice.PeopleAmount,
				TotalPrice:   invoice.TotalPrice,
				IsPaid:       invoice.IsPaid,
				TableID:      invoice.TableID,
				CreatedAt:    invoice.CreatedAt,
				UpdatedAt:    invoice.UpdatedAt,
			},
		})
	}
	return invoiceDetails, nil
}

func (i *InvoiceService) GetAllPaidInvoices(ctx context.Context) ([]responses.InvoiceDetail, error) {
	invoices, err := i.invoiceRepo.GetAllPaid(ctx)
	if err != nil {
		return nil, err
	}
	invoiceDetails := make([]responses.InvoiceDetail, 0)
	for _, invoice := range invoices {
		invoiceDetails = append(invoiceDetails, responses.InvoiceDetail{
			BaseInvoice: responses.BaseInvoice{
				ID:           invoice.ID,
				PeopleAmount: invoice.PeopleAmount,
				TotalPrice:   invoice.TotalPrice,
				IsPaid:       invoice.IsPaid,
				TableID:      invoice.TableID,
				CreatedAt:    invoice.CreatedAt,
				UpdatedAt:    invoice.UpdatedAt,
			},
		})
	}
	return invoiceDetails, nil
}

func (i *InvoiceService) DeleteInvoice(ctx context.Context, invoiceID string) error {
	invoice, err := i.invoiceRepo.FindByID(ctx, invoiceID)
	if err != nil {
		return err
	}

	cancel := i.invoiceRepo.Cancel(ctx, invoiceID)
	if cancel != nil {
		return cancel
	}

	err = i.tableRepo.SetAvailability(ctx, invoice.TableID.String(), true)
	if err != nil {
		return err
	}

	err = i.orderRepo.SetAllPreparingToServed(ctx, invoice.TableID.String())
	if err != nil {
		return err
	}

	return nil
}

func (i *InvoiceService) SetPaidInvoice(ctx context.Context, invoiceID string) error {
	invoice, err := i.invoiceRepo.FindByID(ctx, invoiceID)
	if err != nil {
		return err
	}

	paid := i.invoiceRepo.SetPaid(ctx, invoiceID)
	if paid != nil {
		return paid
	}

	err = i.tableRepo.SetAvailability(ctx, invoice.TableID.String(), true)
	if err != nil {
		return err
	}

	return nil
}

func (i *InvoiceService) CustomerGetInvoice(ctx context.Context, tableID string) (responses.InvoiceDetail, error) {
	invoice, err := i.invoiceRepo.GetByTableID(ctx, tableID)
	if err != nil {
		return responses.InvoiceDetail{}, err
	}
	return responses.InvoiceDetail{
		BaseInvoice: responses.BaseInvoice{
			ID:           invoice.ID,
			PeopleAmount: invoice.PeopleAmount,
			TotalPrice:   invoice.TotalPrice,
			IsPaid:       invoice.IsPaid,
			TableID:      invoice.TableID,
			CreatedAt:    invoice.CreatedAt,
			UpdatedAt:    invoice.UpdatedAt,
		},
	}, nil
}
