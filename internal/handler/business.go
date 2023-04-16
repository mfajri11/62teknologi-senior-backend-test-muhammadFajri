package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/mfajri11/62teknologi-senior-backend-test-muhammadFajri/internal/model"
	"github.com/mfajri11/62teknologi-senior-backend-test-muhammadFajri/internal/service"
	"github.com/mfajri11/62teknologi-senior-backend-test-muhammadFajri/pkg/apperror"
	"github.com/mfajri11/62teknologi-senior-backend-test-muhammadFajri/pkg/web"
)

func LocationOrLatitudeLongitude(sl validator.StructLevel) {
	searchRequest := sl.Current().Interface().(model.BusinessSearchRequest)
	// either location or (latitude & longitude) required
	if searchRequest.Location == "" && (searchRequest.Longitude == 0 || searchRequest.Latitude == 0) {
		sl.ReportError(searchRequest.Location, "location", "Location", "optionallyrequired", "")
		sl.ReportError(searchRequest.Latitude, "latitude", "Latitude", "optionallyrequired", "")
		sl.ReportError(searchRequest.Longitude, "longitude", "Longitude", "optionallyrequired", "")
	}

}

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterStructValidation(LocationOrLatitudeLongitude, model.BusinessSearchRequest{})
	}
}

type BusinnessHandler interface {
	Delete() gin.HandlerFunc
	Create() gin.HandlerFunc
	Update() gin.HandlerFunc
	Search() gin.HandlerFunc
}

type businessHandler struct {
	businessSvc service.BusinessService
}

func NewBusinessHandler(businessSvc service.BusinessService) BusinnessHandler {
	return &businessHandler{businessSvc: businessSvc}
}

func (handler *businessHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		id, err := web.RequiredParam(c, "id")
		if err != nil {
			web.WriteError(c, err)
			return
		}

		token, err := web.AuthorizationRequired(c)
		if err != nil {
			web.WriteError(c, err)
			return
		}
		c.Set("Authorization", token)
		err = handler.businessSvc.Delete(c, id)
		if err != nil {
			web.WriteError(c, err)
			return
		}

		c.Status(http.StatusOK)
	}
}

func (handler *businessHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			err error
			req model.BusinessUpdateRequest
		)
		id, err := web.RequiredParam(c, "id")
		if err != nil {
			web.WriteError(c, err)
			return
		}

		token, err := web.AuthorizationRequired(c)
		if err != nil {
			web.WriteError(c, err)
			return
		}
		c.Set("Authorization", token)
		err = c.ShouldBindJSON(&req)
		if err != nil {
			err := fmt.Errorf("handler.businessHandler.Update: error validate request")
			err = apperror.WrapError(err, apperror.ErrInvalidRequest)
			web.WriteError(c, err)
			return

		}
		resp, err := handler.businessSvc.Update(c, id, req)
		if err != nil {
			web.WriteError(c, err)
			return
		}
		c.JSON(http.StatusOK, resp)
	}
}

func (handler *businessHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			err error
			req model.BusinessCreateRequest
		)

		token, err := web.AuthorizationRequired(c)
		if err != nil {
			web.WriteError(c, err)
			return
		}
		c.Set("Authorization", token)

		err = c.ShouldBindJSON(&req)
		if err != nil {
			err := fmt.Errorf("handler.businessHandler.Create: error validate request: %w", err)
			err = apperror.WrapError(err, apperror.ErrInvalidRequest)
			web.WriteError(c, err)
			return

		}
		resp, err := handler.businessSvc.Create(c, req)
		if err != nil {
			web.WriteError(c, err)
			return
		}
		c.JSON(http.StatusOK, resp)
	}
}

func (handler *businessHandler) Search() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			err  error
			req  model.BusinessSearchRequest
			resp model.BusinessResponse
		)
		token, err := web.AuthorizationRequired(c)
		if err != nil {
			web.WriteError(c, err)
			return
		}

		err = c.ShouldBindQuery(&req)
		if err != nil {
			err := fmt.Errorf("handler.businessHandler.Search: error validate request: %w", err)
			err = apperror.WrapError(err, apperror.ErrInvalidRequest)
			web.WriteError(c, err)
			return

		}
		c.Set("Authorization", token)

		businesses, err := handler.businessSvc.Search(c, req)
		if err != nil {
			web.WriteError(c, err)
			return
		}
		resp = model.BusinessResponse{Businesses: businesses}
		c.JSON(http.StatusOK, resp)

	}
}
