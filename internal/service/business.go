package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	external "github.com/mfajri11/62teknologi-senior-backend-test-muhammadFajri/external/auth"
	"github.com/mfajri11/62teknologi-senior-backend-test-muhammadFajri/internal/model"
	"github.com/mfajri11/62teknologi-senior-backend-test-muhammadFajri/internal/repository"
	"github.com/mfajri11/62teknologi-senior-backend-test-muhammadFajri/pkg/apperror"
)

type businessService struct {
	businessRepo repository.BusinessRepository
}

type BusinessService interface {
	Delete(ctx context.Context, id string) error
	Create(ctx context.Context, req model.BusinessCreateRequest) (resp *model.BusinessCreateResponse, err error)
	Update(ctx context.Context, id string, req model.BusinessUpdateRequest) (resp *model.BusinessUpdateResponse, err error)
	Search(ctx context.Context, req model.BusinessSearchRequest) (resp []*model.BusinessSearchResponse, err error)
}

func NewBusinessService(businessRepo repository.BusinessRepository) BusinessService {
	return &businessService{businessRepo: businessRepo}
}

func (svc *businessService) Delete(ctx context.Context, id string) (err error) {
	tokenString, ok := ctx.Value("Authorization").(string)
	if !ok {
		err = fmt.Errorf("service.businessService.Delete: invalid token type want string got %T", tokenString)
		err = apperror.WrapError(err, apperror.ErrUnauthorized)
		return err
	}
	payload, err := external.ValidateJWT(tokenString)
	if err != nil {
		err = fmt.Errorf("service.businessService.Delete: error validate api key: %w", err)
		err = apperror.WrapError(err, apperror.ErrAuthorize)
		return err
	}
	claim, ok := payload.(*external.JwtBusinessClaim)
	if !ok {
		err = fmt.Errorf("service.businessService.Delete: invalid claim type %T", payload)
		err = apperror.WrapError(err, apperror.ErrUnauthorized)
		return err
	}
	if claim.Role != external.Admin {
		err = fmt.Errorf("service.businessService.Delete: error validate api key: %w", err)
		err = apperror.WrapError(err, apperror.ErrAuthorize)
		return err
	}
	nAffected, err := svc.businessRepo.Delete(ctx, id)

	if err != nil {
		err = fmt.Errorf("service.businessService.Delete: error delete business: %w", err)
		err = apperror.WrapError(err, apperror.ErrInternalError)
		return err
	}

	if err == nil && nAffected == 0 {
		err = fmt.Errorf("service.businessService.Delete: error no rows: %w", pgx.ErrNoRows)
		err = apperror.WrapError(err, apperror.ErrNotFound)
		return err
	}

	return nil
}

func (svc *businessService) Create(ctx context.Context, req model.BusinessCreateRequest) (resp *model.BusinessCreateResponse, err error) {
	tokenString, ok := ctx.Value("Authorization").(string)
	if !ok {
		err = fmt.Errorf("service.businessService.Create: invalid token type want string got %T", tokenString)
		err = apperror.WrapError(err, apperror.ErrUnauthorized)
		return nil, err
	}
	payload, err := external.ValidateJWT(tokenString)
	if err != nil {
		err = fmt.Errorf("service.businessService.Create: error validate api key: %w", err)
		err = apperror.WrapError(err, apperror.ErrAuthorize)
		return nil, err
	}
	claim, ok := payload.(*external.JwtBusinessClaim)
	if !ok {
		err = fmt.Errorf("service.businessService.Create: invalid claim type %T", payload)
		err = apperror.WrapError(err, apperror.ErrUnauthorized)
		return nil, err
	}
	if claim.Role != external.Admin {
		err = fmt.Errorf("service.businessService.Create: error validate api key: %w", err)
		err = apperror.WrapError(err, apperror.ErrAuthorize)
		return nil, err
	}
	id, err := uuid.NewUUID()
	if err != nil {
		err = fmt.Errorf("service.businessService.Create: error create business: %w", err)
		err = apperror.WrapError(err, apperror.ErrInternalError)
		return nil, err
	}

	business := upsertRequestToQueryStruct(id.String(), req)
	businessID, err := svc.businessRepo.Create(ctx, business)

	if err != nil {
		err = fmt.Errorf("service.businessService.Create: error create business: %w", err)
		err = apperror.WrapError(err, apperror.ErrInternalError)
		return nil, err
	}

	resp = &model.BusinessCreateResponse{ID: businessID, BusinessCreateRequest: req}
	resp.DisplayAddress = []string{resp.Address, resp.District, resp.Province, resp.ZipCode, resp.CountryCode}
	// compute price range $, or $$ and more.
	// https://www.cmsmax.com/faqs/misc/price-ranges
	resp.PriceRange = priceRange(resp.Price)

	return resp, nil
}

func (svc *businessService) Update(ctx context.Context, id string, req model.BusinessUpdateRequest) (resp *model.BusinessUpdateResponse, err error) {
	tokenString, ok := ctx.Value("Authorization").(string)
	if !ok {
		err = fmt.Errorf("service.businessService.Update: invalid token type want string got %T", tokenString)
		err = apperror.WrapError(err, apperror.ErrUnauthorized)
		return nil, err
	}
	payload, err := external.ValidateJWT(tokenString)
	if err != nil {
		err = fmt.Errorf("service.businessService.Update: error validate api key: %w", err)
		err = apperror.WrapError(err, apperror.ErrAuthorize)
		return nil, err
	}
	claim, ok := payload.(*external.JwtBusinessClaim)
	if !ok {
		err = fmt.Errorf("service.businessService.Update: invalid claim type %T", payload)
		err = apperror.WrapError(err, apperror.ErrUnauthorized)
		return nil, err
	}
	if claim.Role != external.Admin {
		err = fmt.Errorf("service.businessService.Update: error validate api key: %w", err)
		err = apperror.WrapError(err, apperror.ErrAuthorize)
		return nil, err
	}
	business := upsertRequestToQueryStruct(id, req)
	nAffected, err := svc.businessRepo.Update(ctx, business)
	if err != nil {
		err = fmt.Errorf("service.businessService.update: error update business: %w", err)
		err = apperror.WrapError(err, apperror.ErrInternalError)
		return nil, err
	}

	if err == nil && nAffected == 0 {
		err = fmt.Errorf("service.businessService.Delete: error update business: %w", pgx.ErrNoRows)
		err = apperror.WrapError(err, apperror.ErrNotFound)
		return nil, err
	}

	resp = &model.BusinessUpdateResponse{ID: id, BusinessUpdateRequest: req}
	resp.DisplayAddress = []string{resp.Address, resp.District, resp.Province, resp.ZipCode, resp.CountryCode}
	resp.PriceRange = priceRange(resp.Price)

	return resp, nil
}

func (svc *businessService) Search(ctx context.Context, req model.BusinessSearchRequest) (resp []*model.BusinessSearchResponse, err error) {

	tokenString, ok := ctx.Value("Authorization").(string)
	if !ok {
		err = fmt.Errorf("service.businessService.Update: invalid token type want string got %T", tokenString)
		err = apperror.WrapError(err, apperror.ErrUnauthorized)
		return nil, err
	}
	// authenticate only, everyone has token can access search services
	_, err = external.ValidateJWT(tokenString)
	if err != nil {
		err = fmt.Errorf("service.businessService.Update: error validate api key: %w", err)
		err = apperror.WrapError(err, apperror.ErrAuthorize)
		return nil, err
	}

	businesses, errNoRow, err := svc.businessRepo.Search(ctx, req)
	if errNoRow != nil {
		return []*model.BusinessSearchResponse{}, nil
	}

	if err != nil {
		err = fmt.Errorf("service.businessService.Search: error search: %w", err)
		err = apperror.WrapError(err, apperror.ErrInternalError)
		return nil, err
	}
	resps := make([]*model.BusinessSearchResponse, 0)
	for _, b := range businesses {
		resp := new(model.BusinessSearchResponse)
		resp.ID = b.ID
		resp.Name = b.Name
		resp.Phone = b.Phone
		resp.Price = b.Price
		resp.PriceRange = priceRange(b.Price)
		resp.Location.Address = b.Address
		resp.Location.District = b.District
		resp.Location.Province = b.Province
		resp.Location.ZipCode = b.ZipCode
		resp.Latitude = b.Latitude
		resp.Longitude = b.Longitude
		resp.Rating = b.Rating
		resp.RatingCount = b.RatingCount
		resp.Location.DisplayAddress = []string{b.Address, b.District, b.Province, b.ZipCode}
		categories := strings.Split(b.Categories, ",")
		catResp := make([]model.BusinessCategory, len(categories), len(categories))
		if b.Categories != "" {
			for i, c := range categories {
				cats := strings.Split(c, ":")
				catResp[i] = model.BusinessCategory{
					Alias: cats[0],
					Title: cats[1],
				}
			}
			resp.Categories = catResp
		}

		resps = append(resps, resp)
	}

	return resps, nil
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

func upsertRequestToQueryStruct(id string, r interface{}) model.BusinessUpsertQuery {
	switch req := r.(type) {
	case model.BusinessCreateRequest:
		business := model.BusinessUpsertQuery{
			ID:          id,
			Name:        req.Name,
			Price:       req.Price,
			Phone:       req.Phone,
			OpenAt:      req.OpenAt,
			Address:     req.Address,
			District:    req.District,
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

	case model.BusinessUpdateRequest:
		business := model.BusinessUpsertQuery{
			ID:          id,
			Name:        req.Name,
			Price:       req.Price,
			Phone:       req.Phone,
			OpenAt:      req.OpenAt,
			Address:     req.Address,
			District:    req.District,
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

	default:
		return model.BusinessUpsertQuery{}
	}
}
