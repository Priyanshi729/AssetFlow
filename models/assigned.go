package models

type AssignAssetRequest struct {
	UserID string `json:"user_id" validate:"required"`
}
