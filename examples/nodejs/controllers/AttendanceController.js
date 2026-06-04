const AttendanceRecord = require("../models/AttendanceRecord");

const checkIn = async (req, res) => {
  try {
    const body = req.body;

    const existing = await AttendanceRecord.findOne({
      where: {
        sessionKey: body.sessionKey,
      },
    });

    if (existing) {
      return res.status(200).json({
        success: true,
        message: "Already synced",
      });
    }

    await AttendanceRecord.create(body);

    return res.status(200).json({
      success: true,
      message: "Check-in saved",
    });
  } catch (err) {
    return res.status(500).json({
      success: false,
      message: err.message,
    });
  }
};

const checkOut = async (req, res) => {
  try {
    const { sessionKey, checkOutTime } = req.body;

    const [updatedRows] = await AttendanceRecord.update(
      {
        checkOutTime,
      },
      {
        where: {
          sessionKey,
        },
      }
    );

    if (updatedRows === 0) {
      return res.status(404).json({
        success: false,
        message: "Session not found",
      });
    }

    return res.status(200).json({
      success: true,
      message: "Check-out updated",
    });
  } catch (err) {
    return res.status(500).json({
      success: false,
      message: err.message,
    });
  }
};

module.exports = {
  checkIn,
  checkOut,
};