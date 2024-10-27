package entities

import "time"

type ItemScan struct {
	Time      time.Time    `json:"time,omitempty" db:"time"`
	MachineID int          `json:"machine_id,omitempty" db:"machine_id"`
	Machine   *SettingType `json:"machine,omitempty" db:"machine_id"`
	QrCodeID  int          `json:"qr_code_id,omitempty" db:"qr_code_id"`
	QrCode    *Item        `json:"qr_code,omitempty" db:"qr_code_id"`
}

type LatestScan struct {
	Time    time.Time `json:"time,omitempty" db:"time"`
	Machine string    `json:"machine,omitempty" db:"machine"`
	QrCode  string    `json:"qr_code,omitempty" db:"qr_code"`
	Buyer   string    `json:"buyer,omitempty" db:"buyer"`
	Style   string    `json:"style,omitempty" db:"style"`
	Size    string    `json:"size,omitempty" db:"size"`
	Color   string    `json:"color,omitempty" db:"color"`
}

type ScannedItem struct {
	QrCode string `json:"qr_code,omitempty" db:"qr_code"`
	Buyer  string `json:"buyer,omitempty" db:"buyer"`
	Style  string `json:"style,omitempty" db:"style"`
	Size   string `json:"size,omitempty" db:"size"`
	Color  string `json:"color,omitempty" db:"color"`
	Count  int    `json:"count,omitempty" db:"count"`
}

type ScanItemParam struct {
	Code string `json:"code"`
}
