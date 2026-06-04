# Attendance API Integration

Documentation for integrating the Attendance API to receive employee attendance records from the Attendance application.

## Overview

The Attendance API allows your system to receive employee attendance data from the Attendance application through secure HTTP endpoints.

## Base URL

```text
http://your.api.com
```

> Replace the URL above with your actual API endpoint.

---

## Authentication

All requests must include an API Key in the HTTP headers.

| Header       | Value            |
| ------------ | ---------------- |
| X-API-Key    | YOUR_API_KEY     |
| Content-Type | application/json |

---

# Check-In Endpoint

## Request

**POST**

```text
/check-in
```

### Example

```bash
curl -X POST "http://your.api.com/check-in" \
-H "X-API-Key: YOUR_API_KEY" \
-H "Content-Type: application/json" \
-d '{
  "sessionKey": "SESSION-20260510-001",
  "personId": 1,
  "personName": "John Doe",

  "shiftId": 2,
  "shiftName": "Morning Shift",

  "checkInTime": 1778450400000,
  "checkOutTime": null,

  "lateMinutes": 0,
  "overtimeMinutes": 0,

  "statusText": "Present",

  "deviceId": "ANDROID-KIOSK-01",
  "idempotencyKey": "ATT-1-1778450400000"
}'
```

---

# Check-Out Endpoint

## Request

**PUT**

```text
/check-out
```

### Example

```bash
curl -X PUT "http://your.api.com/check-out" \
-H "X-API-Key: YOUR_API_KEY" \
-H "Content-Type: application/json" \
-d '{
  "sessionKey": "SESSION-20260510-001",
  "personId": 1,
  "personName": "John Doe",

  "shiftId": 2,
  "shiftName": "Morning Shift",

  "checkInTime": 1778450400000,
  "checkOutTime": 1778483100000,

  "lateMinutes": 1,
  "overtimeMinutes": 15,

  "statusText": "Present",

  "deviceId": "ANDROID-KIOSK-01",
  "idempotencyKey": "ATT-1-1778450400000"
}'
```

---

## Field Description
| Field           | Type        | Description                                      |
| --------------- | ----------- | ------------------------------------------------ |
| sessionKey      | string      | Unique attendance session identifier             |
| personId        | integer     | Employee identifier                              |
| personName      | string      | Employee name                                    |
| shiftId         | integer     | Shift identifier                                 |
| shiftName       | string      | Shift name                                       |
| checkInTime     | long        | Check-in timestamp in milliseconds               |
| checkOutTime    | long / null | Check-out timestamp in milliseconds              |
| lateMinutes     | integer     | Late arrival duration in minutes                 |
| overtimeMinutes | integer     | Overtime duration in minutes                     |
| statusText      | string      | Attendance status description                    |
| deviceId        | string      | Device identifier                                |
| idempotencyKey  | string      | Unique key used to prevent duplicate submissions |


---

## Success Response

```json
{
  "success": true,
  "message": "OK"
}
```

---

## Error Response

```json
{
  "success": false,
  "message": "Invalid API Key"
}
```

---

## Example Integrations

| Language | Example                    |
| -------- | -------------------------- |
| Golang   | examples/golang/main.go    |
| Node.js  | examples/nodejs/index.js   |
| PHP      | examples/php/index.php     |
| Python   | examples/python/main.py    |
| C#       | examples/csharp/Program.cs |

---

## License

This repository is provided as an example implementation for Attendance API integration.
