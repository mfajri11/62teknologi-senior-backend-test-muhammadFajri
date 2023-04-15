package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mfajri11/62teknologi-senior-backend-test-muhammadFajri/internal/model"
	"github.com/mfajri11/62teknologi-senior-backend-test-muhammadFajri/internal/repository"
	"github.com/mfajri11/62teknologi-senior-backend-test-muhammadFajri/pkg/apperror"
)

type businessService struct {
	businessRepo repository.BusinesserRepository
}

type BusinessService interface {
	Delete(ctx context.Context, id string) error
	Create(ctx context.Context, req model.BusinessCreateRequest) (resp *model.BusinessCreateResponse, err error)
	Update(ctx context.Context, id string, req model.BusinessUpdateRequest) (resp *model.BusinessUpdateResponse, err error)
}

func NewBusinessService(businessRepo repository.BusinesserRepository) BusinessService {
	return &businessService{businessRepo: businessRepo}
}

func (svc *businessService) Delete(ctx context.Context, id string) error {
	// TODO: add token authentication & authorization (for admin delete & create & update)
	nAffected, err := svc.businessRepo.Delete(ctx, id)

	if err != nil {
		err = fmt.Errorf("service.businessService.Delete: error Delete business: %w", err)
		err = apperror.WrapError(err, apperror.ErrInternalError)
		return err
	}

	if err == nil && nAffected == 0 {
		err = fmt.Errorf("service.businessService.Delete: error Delete business: %w", err)
		err = apperror.WrapError(err, apperror.ErrNotFound)
		return err
	}

	return nil
}

func (svc *businessService) Create(ctx context.Context, req model.BusinessCreateRequest) (resp *model.BusinessCreateResponse, err error) {
	// TODO: add token authentication & authorization (for admin delete & create & update)
	id, err := uuid.NewUUID()
	if err != nil {
		// internal server
		err = fmt.Errorf("service.businessService.Create: error Create business: %w", err)
		err = apperror.WrapError(err, apperror.ErrInternalError)
		return nil, err
	}

	business := requestToQueryStruct(id.String(), req)
	businessID, err := svc.businessRepo.Create(ctx, business)

	if err != nil {
		err = fmt.Errorf("service.businessService.Create: error create business: %w", err)
		return nil, err
	}

	resp = &model.BusinessCreateResponse{ID: businessID, BusinessCreateRequest: req}
	resp.DisplayAddress = []string{resp.Address, resp.District, resp.City, resp.Province, resp.ZipCode, resp.CountryCode}
	// compute price range $, $$ or $$$
	// https://www.cmsmax.com/faqs/misc/price-ranges
	resp.PriceRange = priceRange(resp.Price)

	return resp, nil
}

func (svc *businessService) Update(ctx context.Context, id string, req model.BusinessUpdateRequest) (resp *model.BusinessUpdateResponse, err error) {
	// TODO: add token authentication & authorization (for admin delete & create & update)
	business := requestToQueryStruct(id, req)
	nAffected, err := svc.businessRepo.Update(ctx, business)
	if err != nil {
		err = fmt.Errorf("service.businessService.update: error update business: %w", err)
		err = apperror.WrapError(err, apperror.ErrInternalError)
		return nil, err
	}

	if err == nil && nAffected == 0 {
		err = fmt.Errorf("service.businessService.Delete: error update business: %w", err)
		err = apperror.WrapError(err, apperror.ErrNotFound)
		return nil, err
	}

	resp = &model.BusinessUpdateResponse{ID: id, BusinessCreateRequest: req}
	resp.DisplayAddress = []string{resp.Address, resp.District, resp.City, resp.Province, resp.ZipCode, resp.CountryCode}
	resp.PriceRange = priceRange(resp.Price)

	return resp, nil
}

func priceRange(price float32) string {
	switch {
	case price < 0:
		return ""
	case price < 10:
		return "$"
	case price < 25:
		return "$$"
	case price < 45:
		return "$$$"
	case price > 45:
		return "$$$$"
	default:
		return ""
	}
}

func requestToQueryStruct(id string, req model.BusinessCreateRequest) model.BusinessUpsertQuery {
	business := model.BusinessUpsertQuery{
		ID:          id,
		Name:        req.Name,
		Phone:       req.Phone,
		OpenNow:     req.OpenNow,
		OpenAt:      req.OpenAt,
		Address:     req.Address,
		District:    req.District,
		City:        req.City,
		Province:    req.Province,
		ZipCode:     req.ZipCode,
		CountryCode: req.CountryCode,
		Latitude:    req.Latitude,
		Longitude:   req.Longitude,
		Rating:      req.Rating,
		RatingCount: req.RatingCount,
	}

	if len(req.Categories) > 0 {
		categories := ""
		separator := ","
		for i, cat := range req.Categories {
			if i == len(req.Categories)-1 {
				separator = ""
			}
			categories += fmt.Sprintf("%s:%s%s", cat.Alias, cat.Title, separator)

		}

		business.Categories = categories
	}

	return business

}
