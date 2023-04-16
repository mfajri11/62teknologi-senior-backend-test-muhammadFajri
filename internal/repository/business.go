package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/mfajri11/62teknologi-senior-backend-test-muhammadFajri/config"
	"github.com/mfajri11/62teknologi-senior-backend-test-muhammadFajri/internal/model"
	"github.com/mfajri11/62teknologi-senior-backend-test-muhammadFajri/pkg/postgres"
)

type BusinessRepository interface {
	Create(ctx context.Context, business model.BusinessUpsertQuery) (string, error)
	Update(ctx context.Context, business model.BusinessUpsertQuery) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
	Search(ctx context.Context, business model.BusinessSearchRequest) ([]*model.BusinessJoinAll, error, error)
}

type businessRepository struct {
	sqlClient postgres.SQLer
}

func NewBusinessRepository(db postgres.SQLer) BusinessRepository {
	return &businessRepository{sqlClient: db}
}

func (repo *businessRepository) Delete(ctx context.Context, id string) (int64, error) {
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

func (repo *businessRepository) prepareArgsUpsert(business model.BusinessUpsertQuery) []interface{} {
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

func (repo *businessRepository) Update(ctx context.Context, business model.BusinessUpsertQuery) (int64, error) {
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

func (repo *businessRepository) Create(ctx context.Context, business model.BusinessUpsertQuery) (id string, err error) {
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

func (repo *businessRepository) Search(ctx context.Context, business model.BusinessSearchRequest) ([]*model.BusinessJoinAll, error, error) {

	query, args := repo.prepareBusinessSearchQueryArgs(business)
	rows, err := repo.sqlClient.Query(ctx, query, args...)
	businesses := make([]*model.BusinessJoinAll, 0)
	if err != nil {
		err = fmt.Errorf("repository.businessRepository.Search: error exec select query: %w", err)
		return nil, nil, err
	}
	defer rows.Close()
	for rows.Next() {
		row := new(model.BusinessJoinAll)
		err = rows.Scan(
			&row.ID,
			&row.Name,
			&row.Phone,
			&row.Price,
			&row.Categories,
			&row.OpenNow,
			&row.OpenAt,
			&row.Address,
			&row.District,
			&row.Province,
			&row.CountryCode,
			&row.ZipCode,
			&row.Latitude,
			&row.Longitude,
			&row.Rating,
			&row.Rating,
		)

		if err != nil {
			err = fmt.Errorf("repository.businessRepository.Search: error scan row: %w", err)
			return nil, nil, err
		}

		businesses = append(businesses, row)
	}

	if err != nil {
		err = fmt.Errorf("repository.businessRepository.Search: error next row: %w", err)
		return nil, nil, err
	}

	// err no rows
	if err == nil && !rows.Next() && len(businesses) == 0 {
		err = fmt.Errorf("repository.businessRepository.Search: error no row: %w", pgx.ErrNoRows)
		return nil, err, nil
	}

	return businesses, nil, nil
}

func (repo *businessRepository) prepareBusinessSearchQueryArgs(business model.BusinessSearchRequest) (string, []interface{}) {
	query := `SELECT
	bs.id, bs.name, bs.phone, bs.price, bs.categories, bs.open_now, bs.open_at,
	ba.address, ba.district, ba.province, ba.country_code, ba.zip_code, ba.latitude, ba.longitude,
	br.rating, br.rating_count
FROM business_basic_info bs
INNER JOIN business_address ba ON bs.id = ba.business_id
INNER JOIN business_rating br ON bs.id = br.business_id`
	nArgs := 0
	args := make([]interface{}, 0, 5)
	queriesConds := make([]string, 0, 5)
	if business.Location != "" {
		nArgs += 1
		queriesConds = append(queriesConds,
			fmt.Sprintf("(lower($%d) LIKE lower(ba.address) OR  lower($%d) LIKE lower(ba.district) OR lower($%d) LIKE lower(ba.province) OR lower($%d) LIKE  lower(ba.zip_code) OR lower($%d) LIKE lower(ba.country_code))", nArgs, nArgs, nArgs, nArgs, nArgs))
		args = append(args, business.Location)
	}
	if business.Location == "" {
		if business.Latitude != 0 && business.Longitude != 0 {
			queriesConds = append(queriesConds, fmt.Sprintf("(ba.latitude = $%d AND ba.longitude = $%d)", nArgs+1, nArgs+2))
			nArgs += 2
			args = append(args, business.Latitude, business.Longitude)
		}
	}

	if business.Term != "" {
		nArgs += 1
		queriesConds = append(queriesConds, fmt.Sprintf("bs.name ILIKE ('%%' || $%d || '%%')", nArgs))
		args = append(args, business.Term)
	}
	if business.OpenNow {
		queriesConds = append(queriesConds, "now()::date > bs.open_at")
	}

	if len(queriesConds) > 0 {
		query = fmt.Sprintf("%s WHERE %s", query, strings.Join(queriesConds, " AND "))
	}

	limitOffsetCond := ""
	if business.Limit != 0 {
		nArgs += 1
		limitOffsetCond = fmt.Sprintf("LIMIT $%d", nArgs)
		args = append(args, business.Limit)
		if business.Offset != 0 {
			nArgs += 1
			limitOffsetCond = fmt.Sprintf("%s OFFSET $%d", limitOffsetCond, nArgs)
			args = append(args, business.Offset)
		}
		query = fmt.Sprintf("%s %s", query, limitOffsetCond)
	}
	// if no query params just select with limit
	if limitOffsetCond == "" {
		query = fmt.Sprintf("%s LIMIT %d", query, config.Cfg().App.MaxLimitPagination)
	}
	fmt.Println(query)

	return query, args
}
