<?php

header('Content-Type: application/json');

require 'db.php';
require 'auth.php';

validateApiKey();

$data = json_decode(file_get_contents("php://input"), true);

$sessionKey = $data['session_key'] ?? '';
$checkOutTime = $data['check_out_time'] ?? null;

$stmt = $pdo->prepare("
    UPDATE attendance_records
    SET check_out_time = ?
    WHERE session_key = ?
");

$stmt->execute([
    $checkOutTime,
    $sessionKey
]);

if ($stmt->rowCount() === 0) {

    http_response_code(404);

    echo json_encode([
        "success" => false,
        "message" => "Session not found"
    ]);

    exit;
}

echo json_encode([
    "success" => true,
    "message" => "Check-out updated"
]);