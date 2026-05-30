import requests

url = "https://your.api.com/attendance"

headers = {
    "X-API-Key": "YOUR_API_KEY",
    "Content-Type": "application/json"
}

try:
    response = requests.post(url, headers=headers)
    response.raise_for_status()

    result = response.json()

    print("Status:", result.get("status"))

    for item in result.get("data", []):
        print("Name:", item.get("personName"))
        print("Date:", item.get("tanggal"))
        print("Check In:", item.get("jamMasuk"))
        print("Check Out:", item.get("jamPulang"))
        print("Status:", item.get("statusText"))
        print("---------------------")

except requests.exceptions.RequestException as e:
    print("Request failed:", e)