package model

// entities model
type BusinessBasicInfo struct {
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
	Latitude    float64 `json:"latitude,omitempty"`
	Longitude   float64 `json:"longitude,omitempty"`
	// DisplayAddress will dynamically compute
}

type BusinessJoinAll struct {
	BusinessBasicInfo
	BusinessAddress
	BusinessRating
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

// type BusinessSearchQuery struct {
// 	Location  string
// 	Latitude  float64
// 	Longitude float64
// 	Term      string
// 	OpenNow   string
// 	OpenAt    string
// 	Limit     int64
// 	Offset    int64
// }

// requests model
type BusinessSearchRequest struct {
	Location  string  `form:"location"`
	Latitude  float64 `form:"latitude"`
	Longitude float64 `form:"longitude"`
	Term      string  `form:"term"`
	OpenNow   bool    `form:"open_now"`
	Limit     int64   `form:"limit"`
	Offset    int64   `form:"offset"`
}

type Coordinate struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type BusinessCategory struct {
	Title string `json:"title"`
	Alias string `json:"alias"`
}
type BusinessCreateRequest struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	// TODO: open_now must be boolean
	OpenNow string  `json:"open_now"`
	OpenAt  string  `json:"open_at"`
	Price   float32 `json:"price"`
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
type BusinessLocation struct {
	BusinessAddress
	DisplayAddress []string
}

type BusinessSearchResponse struct {
	ID          string             `form:"id"`
	Categories  []BusinessCategory `form:"categories"`
	Latitude    float64            `form:"latitude"`
	Longitude   float64            `form:"longitude"`
	Coordinates Coordinate         `form:"coordinates"`
	Location    BusinessLocation   `form:"location"`
	Name        string             `form:"name"`
	Phone       string             `form:"phone"`
	Price       float32            `form:"price"`
	PriceRange  string             `form:"price_range"`
	Rating      float32            `form:"rating"`
	RatingCount int64              `form:"rating_count"`
}

type BusinessResponse struct {
	Businesses []*BusinessSearchResponse `json:"businesses"`
}
