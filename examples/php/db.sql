CREATE TABLE attendance_records (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,

    session_key VARCHAR(255) UNIQUE,

    person_id BIGINT,
    person_name VARCHAR(255),

    shift_id BIGINT,
    shift_name VARCHAR(255),

    check_in_time BIGINT NULL,
    check_out_time BIGINT NULL,

    late_minutes INT DEFAULT 0,
    overtime_minutes INT DEFAULT 0,

    status_text VARCHAR(255),
    device_id VARCHAR(255),

    idempotency_key VARCHAR(255) UNIQUE,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        ON UPDATE CURRENT_TIMESTAMP
);