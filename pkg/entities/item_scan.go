package entities

import "time"

type ItemScan struct {
	Time      time.Time    `json:"time,omitempty" db:"time"`
	MachineID *SettingType `json:"machine_id,omitempty" db:"machine_id"`
	QrCodeID  *Item        `json:"qr_code_id,omitempty" db:"qr_code_id"`
}
