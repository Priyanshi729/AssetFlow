package models

type AssignAssetRequest struct {
	UserID string `json:"user_id" validate:"required"`
}

type AssetAssignmentHistory struct {
	AssignmentID string  `db:"assignment_id" json:"assignment_id"`
	AssetID      string  `db:"asset_id" json:"asset_id"`
	Brand        string  `db:"brand" json:"brand"`
	Model        string  `db:"model" json:"model"`
	AssetType    string  `db:"asset_type" json:"asset_type"`
	UserID       string  `db:"user_id" json:"user_id"`
	Name         string  `db:"name" json:"name"`
	Status       string  `db:"status" json:"status"`
	AssignedOn   string  `db:"assigned_on" json:"assigned_on"`
	ReturnedAt   *string `db:"returned_at" json:"returned_at"`
}
