package model

type AttendanceRecord struct {
	ID int `gorm:"primaryKey"`
	SessionKey string `gorm:"uniqueIndex"`

	PersonID   int64
	PersonName string

	ShiftID   int64
	ShiftName string

	CheckInTime  *int64
	CheckOutTime *int64

	LateMinutes     int    `json:"lateMinutes"`
	OvertimeMinutes int    `json:"overtimeMinutes"`
	StatusText     string `json:"statusText"`
	DeviceID       string `json:"deviceId"` 
	IdempotencyKey string `gorm:"uniqueIndex"`
}