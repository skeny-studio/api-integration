require("dotenv").config();

const express = require("express");
const sequelize = require("./database");

const attendanceRoutes = require("./routes/AttendanceRoutes");

const app = express();

app.use(express.json());

app.use("/attendance", attendanceRoutes);

sequelize
  .authenticate()
  .then(() => {
    console.log("Database connected");
    return sequelize.sync();
  })
  .then(() => {
    app.listen(3000, () => {
      console.log("Server running on port 3000");
    });
  })
  .catch(console.error);