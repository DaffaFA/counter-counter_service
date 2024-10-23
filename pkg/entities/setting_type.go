package entities

type SettingType struct {
	Alias string `json:"alias" db:"alias"`
	Name  int    `json:"name" db:"name"`
}
