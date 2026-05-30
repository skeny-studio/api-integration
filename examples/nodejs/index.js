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

    console.log(response.data);
  } catch (err) {
    console.error(err.message);
  }
}

getAttendance();