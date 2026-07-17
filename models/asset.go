package models

import "time"

type CreateAssetRequest struct {
	Brand         string     `json:"brand" validate:"required"`
	Model         string     `json:"model" validate:"required"`
	SerialNumber  string     `json:"serial_number" validate:"required"`
	AssetType     string     `json:"asset_type" validate:"required"`
	Status        string     `json:"status"`
	OwnerType     string     `json:"owner_type"`
	WarrantyStart *time.Time `json:"warranty_start" validate:"required"`
	WarrantyEnd   *time.Time `json:"warranty_end" validate:"required"`

	Laptop   LaptopRequestSpecific   `json:"laptop"`
	Mobile   MobileRequestSpecific   `json:"mobile"`
	Keyboard KeyboardRequestSpecific `json:"keyboard"`
	Mouse    MouseRequestSpecific    `json:"mouse"`
}

type LaptopRequestSpecific struct {
	Processor       string `db:"processor" json:"processor"`
	RAM             string `db:"ram" json:"ram"`
	Storage         string `db:"storage" json:"storage"`
	OperatingSystem string `db:"operating_system" json:"operating_system"`
	Charger         bool   `db:"charger" json:"charger"`
	DevicePassword  string `db:"device_password" json:"device_password"`
}

type MobileRequestSpecific struct {
	RAM             string `db:"ram" json:"ram"`
	Storage         string `db:"storage" json:"storage"`
	OperatingSystem string `db:"operating_system" json:"operating_system"`
	Charger         bool   `db:"charger" json:"charger"`
	DevicePassword  string `db:"device_password" json:"device_password"`
}

type MouseRequestSpecific struct {
	DPI          int    `db:"dpi" json:"dpi"`
	Connectivity string `db:"connectivity" json:"connectivity"`
}

type KeyboardRequestSpecific struct {
	Layout       string `db:"layout" json:"layout"`
	Connectivity string `db:"connectivity" json:"connectivity"`
}

type Asset struct {
	AssetID       string     `db:"asset_id" json:"asset_id"`
	Brand         string     `db:"brand" json:"brand"`
	Model         string     `db:"model" json:"model"`
	SerialNumber  string     `db:"serial_number" json:"serial_number"`
	AssetType     string     `db:"asset_type" json:"asset_type"`
	Status        string     `db:"status" json:"status"`
	OwnerType     string     `db:"owner_type" json:"owner_type"`
	WarrantyStart *time.Time `db:"warranty_start" json:"warranty_start"`
	WarrantyEnd   *time.Time `db:"warranty_end" json:"warranty_end"`

	Laptop   *LaptopRequestSpecific   `json:"laptop,omitempty"`
	Mobile   *MobileRequestSpecific   `json:"mobile,omitempty"`
	Keyboard *KeyboardRequestSpecific `json:"keyboard,omitempty"`
	Mouse    *MouseRequestSpecific    `json:"mouse,omitempty"`
}

type UpdateAssetRequest struct {
	Brand         *string `json:"brand"`
	Model         *string `json:"model"`
	SerialNumber  *string `json:"serial_number"`
	Status        *string `json:"status"`
	OwnerType     *string `json:"owner_type"`
	WarrantyStart *string `json:"warranty_start"`
	WarrantyEnd   *string `json:"warranty_end"`

	Laptop   *UpdateLaptopRequest   `json:"laptop"`
	Mobile   *UpdateMobileRequest   `json:"mobile"`
	Keyboard *UpdateKeyboardRequest `json:"keyboard"`
	Mouse    *UpdateMouseRequest    `json:"mouse"`
}
type UpdateLaptopRequest struct {
	Processor       *string `json:"processor"`
	RAM             *string `json:"ram"`
	Storage         *string `json:"storage"`
	OperatingSystem *string `json:"operating_system"`
	Charger         *bool   `json:"charger"`
	DevicePassword  *string `json:"device_password"`
}

type UpdateMobileRequest struct {
	RAM             *string `db:"ram" json:"ram"`
	Storage         *string `db:"storage" json:"storage"`
	OperatingSystem *string `db:"operating_system" json:"operating_system"`
	Charger         *bool   `db:"charger" json:"charger"`
	DevicePassword  *string `db:"device_password" json:"device_password"`
}

type UpdateMouseRequest struct {
	DPI          *int    `db:"dpi" json:"dpi"`
	Connectivity *string `db:"connectivity" json:"connectivity"`
}

type UpdateKeyboardRequest struct {
	Layout       *string `db:"layout" json:"layout"`
	Connectivity *string `db:"connectivity" json:"connectivity"`
}
