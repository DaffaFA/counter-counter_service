package entities

type Setting struct {
	ID               int          `json:"id,omitempty" db:"id"`
	SettingTypeAlias string       `json:"setting_type_alias,omitempty" db:"setting_type_alias"`
	SettingType      *SettingType `json:"setting_type,omitempty"`
	Value            *string      `json:"value,omitempty" db:"value"` // Nullable value
	ParentID         int          `json:"parent_id,omitempty" db:"parent_id"`
}

type SettingPagination struct {
	Total    int       `json:"total"`
	Settings []Setting `json:"settings"`
}

type MachineDetail struct {
	Factory string `json:"factory"`
	Machine string `json:"machine"`
}
