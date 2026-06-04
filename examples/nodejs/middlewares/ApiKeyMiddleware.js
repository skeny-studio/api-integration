require("dotenv").config();

const apiKeyMiddleware = (req, res, next) => {
  const apiKey = req.header("X-API-Key");

  if (apiKey !== process.env.ATTENDANCE_API_KEY) {
    return res.status(401).json({
      success: false,
      message: "Invalid API Key",
    });
  }

  next();
};

module.exports = apiKeyMiddleware;