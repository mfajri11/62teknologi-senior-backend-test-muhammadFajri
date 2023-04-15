package repository

import (
	"context"
	"fmt"

	"github.com/mfajri11/62teknologi-senior-backend-test-muhammadFajri/internal/model"
	"github.com/mfajri11/62teknologi-senior-backend-test-muhammadFajri/pkg/postgres"
)

type BusinesserRepository interface {
	Create(ctx context.Context, business model.BusinessUpsertQuery) (string, error)
	Update(ctx context.Context, business model.BusinessUpsertQuery) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
	// Search()
}

type BusinessRepository struct {
	sqlClient postgres.SQLer
}

func NewBusinessRepository(db postgres.SQLer) *BusinessRepository {
	return &BusinessRepository{sqlClient: db}
}

func (repo *BusinessRepository) Delete(ctx context.Context, id string) (int64, error) {
	query := `
	WITH 
		basic AS 
			(DELETE FROM business_basic_info WHERE id = $1),
		address AS
			(DELETE FROM business_address WHERE business_id = $1)
		DELETE FROM business_rating WHERE business_id = $1`

	result, err := repo.sqlClient.Exec(ctx, query, id)

	if err != nil {
		err = fmt.Errorf("repository.BusinessRepository.DeleteByID: error exec delete query: %w", err)
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (repo *BusinessRepository) prepareArgsUpsert(business model.BusinessUpsertQuery) []interface{} {
	args := make([]interface{}, 16)
	args[0] = business.ID
	args[1] = business.Name
	args[2] = business.Phone
	args[3] = business.Price
	args[4] = business.Categories
	args[5] = business.OpenNow
	args[6] = business.OpenAt
	args[7] = business.Address
	args[8] = business.District
	args[9] = business.Province
	args[10] = business.CountryCode
	args[11] = business.ZipCode
	args[12] = business.Latitude
	args[13] = business.Longitude
	args[14] = business.Rating
	args[15] = business.RatingCount

	return args
}

func (repo *BusinessRepository) Update(ctx context.Context, business model.BusinessUpsertQuery) (int64, error) {
	query := `
	WITH basic AS (
		UPDATE business_basic_info
		SET 
			name = COALESCE(NULLIF($2, ''), name),
			phone = COALESCE(NULLIF($3, ''), phone),
			price = COALESCE(NULLIF($4, 0), price),
			categories = COALESCE(NULLIF($5, ''), categories),
			open_now = COALESCE(NULLIF($6, ''), open_now),
			open_at = COALESCE(NULLIF($7, ''), open_at)
		WHERE id = $1
	),
	address AS (
		UPDATE business_address
		SET
			address = COALESCE(NULLIF($8, ''), address),
			district = COALESCE(NULLIF($9, ''), district),
			province = COALESCE(NULLIF($10, ''), province),
			country_code = COALESCE(NULLIF($11, ''), country_code),
			zip_code = COALESCE(NULLIF($12, ''), zip_code),
			latitude = COALESCE(NULLIF($13, 0), latitude),
			longitude = COALESCE(NULLIF($14, 0), longitude)
		WHERE business_id = $1
	) UPDATE business_rating
	SET
		rating = COALESCE(NULLIF($15, 0), rating),
		rating_count = COALESCE(NULLIF($16, 0), rating_count)
	WHERE business_id = $1;`

	args := repo.prepareArgsUpsert(business)
	res, err := repo.sqlClient.Exec(ctx, query, args...)

	if err != nil {
		err := fmt.Errorf("repository.BusinessRepository.UpdateByID: error exec update query: %w", err)
		return 0, err
	}

	return res.RowsAffected(), nil
}

func (repo *BusinessRepository) Create(ctx context.Context, business model.BusinessUpsertQuery) (id string, err error) {
	query := `
	WITH basic AS (
		INSERT INTO business_basic_info
			(id, name, phone, price, categories, open_now, open_at)
		VALUES
			($1, $2, $3, $4, $5, $6, $7)
	),
	address AS(
		INSERT INTO business_address
			(business_id, address, district, province, country_code, zip_code, latitude, longitude)
		VALUES
			($1, $8, $9, $10, $11, $12, $13, $14)
	) INSERT INTO business_rating
	(business_id, rating, rating_count)
	VALUES
	($1, $15, $16);`

	args := repo.prepareArgsUpsert(business)
	_, err = repo.sqlClient.Exec(ctx, query, args...)

	if err != nil {
		err := fmt.Errorf("repository.BusinessRepository.Create: error exec insert query: %w", err)
		return "", err
	}
	return business.ID, nil
}
