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
