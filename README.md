# Attendance API Integration

Dokumentasi integrasi API Attendance untuk menerima data absensi karyawan dari aplikasi Attendance.

## Overview

Attendance API digunakan untuk mengirim data absensi dari aplikasi Attendance ke sistem backend pihak ketiga.

## Endpoint

### POST Attendance

**URL**

```text
https://your.api.com/attendance
```

> Ganti URL di atas dengan endpoint API Anda.

## Authentication

Gunakan API Key pada header request.

| Header       | Value            |
| ------------ | ---------------- |
| X-API-Key    | YOUR_API_KEY     |
| Content-Type | application/json |

## cURL Example

```bash
curl -X POST "https://your.api.com/attendance" \
-H "X-API-Key: YOUR_API_KEY" \
-H "Content-Type: application/json"
```

## Request Body

```json
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
```

## Field Description

| Field          | Type    | Description                   |
| -------------- | ------- | ----------------------------- |
| personId       | integer | Employee ID                   |
| personName     | string  | Employee name                 |
| status         | integer | Attendance status             |
| timestamp      | long    | Unix timestamp (milliseconds) |
| tanggal        | string  | Attendance date               |
| jamMasuk       | string  | Check-in time                 |
| jamPulang      | string  | Check-out time                |
| shiftId        | integer | Shift ID                      |
| shiftName      | string  | Shift name                    |
| shiftSchedule  | string  | Shift schedule                |
| terlambatMenit | integer | Late minutes                  |
| lemburMenit    | integer | Overtime minutes              |
| statusText     | string  | Attendance status text        |

## Success Response

```json
{
  "success": true,
  "message": "OK"
}
```

## Error Response

```json
{
  "success": false,
  "message": "Invalid API Key"
}
```

## Example Integrations

| Language | File                       |
| -------- | -------------------------- |
| Golang   | examples/golang/main.go    |
| Node.js  | examples/nodejs/index.js   |
| PHP      | examples/php/index.php     |
| Python   | examples/python/main.py    |
| C#       | examples/csharp/Program.cs |
