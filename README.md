Berikut catatan dokumentasi backend Go untuk integrasi API attendance.

# Attendance API Integration (Go)

API ini digunakan untuk mengambil data absensi karyawan.

## Endpoint

```bash
POST 
Headers
Header	Value
X-API-Key	YOUR_API_KEY
Content-Type	application/json

Request
curl -X POST "https://your.api.com/attendance" \
-H "X-API-Key: YOUR_API_KEY" \
-H "Content-Type: application/json"

</> JSON:
{
  "personId": 1,
  "personName": "Budi Santoso",
  "status": 1,
  "timestamp": 1778450400000,
  "tanggal": "10 May 2026",
  "jamMasuk": "08:01",
  "jamPulang": "17:05",
  "shiftId": 2,
  "shiftName": "Morning",
  "shiftSchedule": "08:00 - 17:00",
  "terlambatMenit": 1,
  "lemburMenit": 15,
  "statusText": "Present"
}

Response
</> JSON :
{
  "success": true,
  "message": "OK"
}


Golang Example

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AttendanceResponse struct {
	Status string       `json:"status"`
	Data   []Attendance `json:"data"`
}

type Attendance struct {
	PersonID        int    `json:"personId"`
	PersonName      string `json:"personName"`
	Status          int    `json:"status"`
	Timestamp       int64  `json:"timestamp"`
	Tanggal         string `json:"tanggal"`
	JamMasuk        string `json:"jamMasuk"`
	JamPulang       string `json:"jamPulang"`
	ShiftID         int    `json:"shiftId"`
	ShiftName       string `json:"shiftName"`
	ShiftSchedule   string `json:"shiftSchedule"`
	TerlambatMenit  int    `json:"terlambatMenit"`
	LemburMenit     int    `json:"lemburMenit"`
	StatusText      string `json:"statusText"`
}

func main() {

	url := "https://your.api.com"

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("X-API-Key", "YOUR_API_KEY")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result AttendanceResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}

	fmt.Println("Status:", result.Status)

	for _, item := range result.Data {
		fmt.Println("Nama:", item.PersonName)
		fmt.Println("Tanggal:", item.Tanggal)
		fmt.Println("Jam Masuk:", item.JamMasuk)
		fmt.Println("Jam Pulang:", item.JamPulang)
		fmt.Println("Status:", item.StatusText)
		fmt.Println("---------------------")
	}
}
