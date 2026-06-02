package model

type AttendanceRecord struct {
	PersonID       int    `json:"personId"`
	PersonName     string `json:"personName"`
	Status         int    `json:"status"`
	Timestamp      int64  `json:"timestamp"`
	Tanggal        string `json:"tanggal"`
	JamMasuk       string `json:"jamMasuk"`
	JamPulang      string `json:"jamPulang"`
	ShiftID        int    `json:"shiftId"`
	ShiftName      string `json:"shiftName"`
	ShiftSchedule  string `json:"shiftSchedule"`
	TerlambatMenit int    `json:"terlambatMenit"`
	LemburMenit    int    `json:"lemburMenit"`
	StatusText     string `json:"statusText"`
}