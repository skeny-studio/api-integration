# Attendance API Integration

Documentation for integrating the Attendance API to receive employee attendance records from the Attendance application.

## Overview

The Attendance API allows your system to receive employee attendance data from the Attendance application through a secure HTTP endpoint.

## Endpoint

### POST Attendance

**URL**

```text
https://your.api.com/attendance
```

> Replace the URL above with your actual API endpoint.

## Authentication

Requests must include an API Key in the HTTP headers.

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

| Field          | Type    | Description                      |
| -------------- | ------- | -------------------------------- |
| personId       | integer | Employee identifier              |
| personName     | string  | Employee name                    |
| status         | integer | Attendance status code           |
| timestamp      | long    | Unix timestamp in milliseconds   |
| tanggal        | string  | Attendance date                  |
| jamMasuk       | string  | Check-in time                    |
| jamPulang      | string  | Check-out time                   |
| shiftId        | integer | Shift identifier                 |
| shiftName      | string  | Shift name                       |
| shiftSchedule  | string  | Scheduled shift time             |
| terlambatMenit | integer | Late arrival in minutes          |
| lemburMenit    | integer | Overtime duration in minutes     |
| statusText     | string  | Human-readable attendance status |

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

| Language | Example                    |
| -------- | -------------------------- |
| Golang   | examples/golang/main.go    |
| Node.js  | examples/nodejs/index.js   |
| PHP      | examples/php/index.php     |
| Python   | examples/python/main.py    |
| C#       | examples/csharp/Program.cs |

## License

This repository is provided as an example implementation for Attendance API integration.
