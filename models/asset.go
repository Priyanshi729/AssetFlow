package models

type CreateAssetRequest struct {
	Brand         string `json:"brand" validate:"required"`
	Model         string `json:"model" validate:"required"`
	SerialNumber  string `json:"serial_number" validate:"required"`
	AssetType     string `json:"asset_type" validate:"required"`
	Status        string `json:"status"`
	OwnerType     string `json:"owner_type"`
	WarrantyStart string `json:"warranty_start" validate:"required"`
	WarrantyEnd   string `json:"warranty_end" validate:"required"`

	OperatingSystem string `json:"operating_system"`
	RAM             string `json:"ram"`
	Storage         string `json:"storage"`
	Charger         bool   `json:"charger"`
	DevicePassword  string `json:"device_password"`

	Processor string `json:"processor"`

	Layout string `json:"layout"`

	Connectivity string `json:"connectivity"`

	DPI int `json:"dpi"`
}

type Asset struct {
	AssetID       string `db:"asset_id" json:"asset_id"`
	Brand         string `db:"brand" json:"brand"`
	Model         string `db:"model" json:"model"`
	SerialNumber  string `db:"serial_number" json:"serial_number"`
	AssetType     string `db:"asset_type" json:"asset_type"`
	Status        string `db:"status" json:"status"`
	OwnerType     string `db:"owner_type" json:"owner_type"`
	WarrantyStart string `db:"warranty_start" json:"warranty_start"`
	WarrantyEnd   string `db:"warranty_end" json:"warranty_end"`
	CreatedAt     string `db:"created_at" json:"created_at"`
}

type AssetDetails struct {
	Asset

	Processor       *string `db:"processor" json:"processor,omitempty"`
	RAM             *string `db:"ram" json:"ram,omitempty"`
	Storage         *string `db:"storage" json:"storage,omitempty"`
	OperatingSystem *string `db:"operating_system" json:"operating_system,omitempty"`
	Charger         *bool   `db:"charger" json:"charger,omitempty"`
	DevicePassword  *string `db:"device_password" json:"device_password,omitempty"`

	Layout       *string `db:"layout" json:"layout,omitempty"`
	Connectivity *string `db:"connectivity" json:"connectivity,omitempty"`

	DPI *int `db:"dpi" json:"dpi,omitempty"`
}
