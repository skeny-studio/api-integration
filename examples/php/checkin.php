<?php

header('Content-Type: application/json');

require 'db.php';
require 'auth.php';

validateApiKey();

$data = json_decode(file_get_contents("php://input"), true);

if (!$data) {

    http_response_code(400);

    echo json_encode([
        "success" => false,
        "message" => "Invalid JSON"
    ]);

    exit;
}

$sessionKey = $data['session_key'] ?? '';

$stmt = $pdo->prepare(
    "SELECT id FROM attendance_records
     WHERE session_key = ?"
);

$stmt->execute([$sessionKey]);

if ($stmt->fetch()) {

    echo json_encode([
        "success" => true,
        "message" => "Already synced"
    ]);

    exit;
}

$stmt = $pdo->prepare("
    INSERT INTO attendance_records (
        session_key,
        person_id,
        person_name,
        shift_id,
        shift_name,
        check_in_time,
        check_out_time,
        late_minutes,
        overtime_minutes,
        status_text,
        device_id,
        idempotency_key
    ) VALUES (
        ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
    )
");

$stmt->execute([
    $data['session_key'] ?? null,
    $data['person_id'] ?? null,
    $data['person_name'] ?? null,
    $data['shift_id'] ?? null,
    $data['shift_name'] ?? null,
    $data['check_in_time'] ?? null,
    $data['check_out_time'] ?? null,
    $data['lateMinutes'] ?? 0,
    $data['overtimeMinutes'] ?? 0,
    $data['statusText'] ?? '',
    $data['deviceId'] ?? '',
    $data['idempotency_key'] ?? null
]);

echo json_encode([
    "success" => true,
    "message" => "Check-in saved"
]);