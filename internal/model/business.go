package model

// entities model
type BasicBusinessInfo struct {
	ID    string
	Name  string
	Phone string
	Price float32
	// categories value example: alias1:title1, alias2:title2, ... aliasN:titleN, N: total categories.
	Categories string
	OpenNow    string
	OpenAt     string
}

type BusinessRating struct {
	// ID          string
	BusinessID  string
	Rating      float32
	RatingCount int64
}

type BusinessAddress struct {
	BusinessID  string  `json:"business_id"`
	Address     string  `json:"address"`
	District    string  `json:"district"`
	Province    string  `json:"province"`
	CountryCode string  `json:"country_code"`
	ZipCode     string  `json:"zipcode"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	// DisplayAddress will dynamically compute
}
type BusinessUpsertQuery struct {
	ID          string
	Name        string
	Phone       string
	OpenNow     string
	OpenAt      string
	Price       float32
	Categories  string
	Address     string
	City        string
	District    string
	Province    string
	CountryCode string
	ZipCode     string
	Latitude    float64
	Longitude   float64
	Rating      float32
	RatingCount int64
}

// requests model
type BusinessSearchRequest struct {
	Location  string  `json:"location"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Term      string  `json:"term"`
	// radius
	OpenNow string `json:"open_now"`
	OpenAt  string `json:"open_at"`
	Limit   int64  `json:"limit"`
	Offset  int64  `json:"offset"`
}

type coordinate struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type BusinessCategory struct {
	Title string `json:"title"`
	Alias string `json:"alias"`
}
type BusinessCreateRequest struct {
	Name    string  `json:"name"`
	Phone   string  `json:"phone"`
	OpenNow string  `json:"open_now"`
	OpenAt  string  `json:"open_at"`
	Price   float32 `json:"Price"`
	// in put request any existing categories will be replaced ensure existing categories it included if the're won't be replaced
	Categories  []BusinessCategory `json:"categories"`
	Address     string             `json:"address"`
	City        string             `json:"city"`
	District    string             `json:"district"`
	Province    string             `json:"province"`
	CountryCode string             `json:"country_code"`
	ZipCode     string             `json:"zip_code"`
	Latitude    float64            `json:"latitude"`
	Longitude   float64            `json:"longitude"`
	Rating      float32            `json:"rating"`
	RatingCount int64              `json:"rating_count"`
}

type BusinessUpdateRequest = BusinessCreateRequest

// responses model
type BusinessCreateResponse struct {
	ID string `json:"id"`
	BusinessCreateRequest
	DisplayAddress []string `json:"display_address"`
	PriceRange     string
}

type BusinessUpdateResponse = BusinessCreateResponse

type BusinessSearchResponse struct {
	Alias       string             `json:"alias"`
	Categories  []BusinessCategory `json:"categories"`
	Latitude    float64            `json:"latitude"`
	Longitude   float64            `json:"longitude"`
	Coordinates coordinate         `json:"coordinates"`
	Location    BusinessAddress    `json:"location"`
	Name        string             `json:"name"`
	Phone       string             `json:"phone"`
	Price       float32            `json:"price"`
	PriceRange  string             `json:"price_range"`
	Rating      float32            `json:"rating"`
	ReviewCount int64              `json:"review_count"`
}

type BusinessResponse struct {
	Businesses []BusinessSearchResponse `json:"businesses"`
}
