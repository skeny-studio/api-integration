const express = require("express");
const router = express.Router();

const {
  checkIn,
  checkOut,
} = require("../controllers/AttendanceController");

const apiKeyMiddleware = require("../middlewares/ApiKeyMiddleware");

router.post("/checkin", apiKeyMiddleware, checkIn);
router.post("/checkout", apiKeyMiddleware, checkOut);

module.exports = router;