const { DataTypes } = require("sequelize");
const sequelize = require("../database");

const AttendanceRecord = sequelize.define(
  "AttendanceRecord",
  {
    id: {
      type: DataTypes.INTEGER,
      primaryKey: true,
      autoIncrement: true,
    },
    sessionKey: {
      type: DataTypes.STRING,
      unique: true,
      allowNull: false,
    },

    personID: DataTypes.BIGINT,
    personName: DataTypes.STRING,

    shiftID: DataTypes.BIGINT,
    shiftName: DataTypes.STRING,

    checkInTime: {
      type: DataTypes.BIGINT,
      allowNull: true,
    },

    checkOutTime: {
      type: DataTypes.BIGINT,
      allowNull: true,
    },

    lateMinutes: DataTypes.INTEGER,
    overtimeMinutes: DataTypes.INTEGER,

    statusText: DataTypes.STRING,
    deviceId: DataTypes.STRING,

    idempotencyKey: {
      type: DataTypes.STRING,
      unique: true,
    },
  },
  {
    tableName: "attendance_records",
    timestamps: false,
  }
);

module.exports = AttendanceRecord;