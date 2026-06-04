<?php

function validateApiKey()
{
    $headers = getallheaders();

    $apiKey = $headers['X-API-Key'] ?? '';

    if ($apiKey !== getenv('ATTENDANCE_API_KEY')) {

        http_response_code(401);

        echo json_encode([
            "success" => false,
            "message" => "Invalid API Key"
        ]);

        exit;
    }
}