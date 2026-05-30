const axios = require("axios");

async function getAttendance() {
  try {
    const response = await axios.post(
      "https://your.api.com/attendance",
      {},
      {
        headers: {
          "X-API-Key": "YOUR_API_KEY",
          "Content-Type": "application/json"
        }
      }
    );

    const result = response.data;

    console.log("Status:", result.status);

    if (Array.isArray(result.data)) {
      result.data.forEach((item) => {
        console.log("Name:", item.personName);
        console.log("Date:", item.tanggal);
        console.log("Check In:", item.jamMasuk);
        console.log("Check Out:", item.jamPulang);
        console.log("Status:", item.statusText);
        console.log("---------------------");
      });
    }
  } catch (error) {
    console.error("Request failed:", error.message);
  }
}

getAttendance();