CREATE TABLE business_basic_info(
	id TEXT PRIMARY KEY,
	name TEXT NOT NULL,
	phone VARCHAR(20) NOT NULL,
	price FLOAT4 NOT NULL,
	categories TEXT,
	open_at TIME
);

CREATE INDEX IF NOT EXISTS idx_business_basic_info_id ON business_basic_info(id);

CREATE INDEX IF NOT EXISTS idx_business_basic_info_name ON business_basic_info(name);

CREATE TABLE business_rating(
    id serial PRIMARY KEY,
	business_id TEXT NOT NULL REFERENCES business_basic_info(id) ON DELETE CASCADE,
	rating float4,
	rating_count BIGINT NOT NULL DEFAULT 0
);

CREATE INDEX IF NOT EXISTS idx_business_rating_business_id ON business_rating(business_id);

CREATE TABLE business_address(
    id serial PRIMARY KEY,
	business_id TEXT REFERENCES business_basic_info(id) ON DELETE CASCADE,
	address TEXT NOT NULL,
	district TEXT NOT NULL,
	province TEXT NOT NULL,
	country_code TEXT NOT NULL,
	zip_code TEXT,
	latitude FLOAT8,
	longitude FLOAT8
);

CREATE INDEX IF NOT EXISTS idx_business_address_business_id ON business_address(business_id);

CREATE INDEX IF NOT EXISTS idx_business_address_address ON business_address(address);

CREATE INDEX IF NOT EXISTS idx_business_address_district ON business_address(district);