package model

// entities model
type BasicBusinessInfo struct {
	ID    string
	Name  string
	Phone string
	Price float32
	// Transactions string
	OpenNow string
	OpenAt  string
}

type BusinessRating struct {
	ID          string
	BusinessID  string
	Rating      float32
	ReviewCount int64
}

type BusinessCategory struct {
	ID         string
	BusinessID string
	Title      string
	Alias      string // i don't have more information whether alias and title always the same words or not (ignoring the case) better have 2 field for safety
}

// I think junction table will help if many business entities has share the same categories
// e.g. BusinessA, BusinessB, BusinessC  has the same categories (Categories1)
type BusinessCategoryJunction struct {
	ID         string
	BusinessID string
	CategoryID string
}

type BusinessAddress struct {
	Address     string  `json:"address"`
	District    string  `json:"district"`
	Province    string  `json:"province"`
	CountryCode string  `json:"country_code"`
	ZipCode     string  `json:"zip_code"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	// DisplayAddress string // useful for searching even though can dynamically create & every time address updated displayAddress must be updated too
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

// type category struct {
// 	Alias string
// 	Title string
// }

type coordinate struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

// type location struct {
// 	Address        string
// 	District       string
// 	Province       string
// 	Country        string
// 	ZipCode        string
// 	DisplayAddress string
// }

type BusinessSearchCreateRequest struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	// Transactions
	OpenNow string `json:"open_now"`
	OpenAt  string
	// in put request any existing categories will be replaced ensure existing categories it included if the're won't be replaced
	Categories   []BusinessCategory `json:"categories"`
	Address      string             `json:"address"`
	City         string             `json:"city"`
	District     string             `json:"district"`
	Province     string             `json:"province"`
	Country_code string             `json:"country_code"`
	ZipCode      string             `json:"zip_code"`
}

type BusinessSearchPutRequest = BusinessSearchCreateRequest

// responses model
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
	// transactions
}

type BusinessResponse struct {
	Businesses []BusinessSearchResponse `json:"businesses"`
}
