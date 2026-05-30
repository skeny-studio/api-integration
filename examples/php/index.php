<?php

$url = "https://your.api.com/attendance";

$headers = [
    "X-API-Key: YOUR_API_KEY",
    "Content-Type: application/json"
];

$ch = curl_init($url);

curl_setopt_array($ch, [
    CURLOPT_POST => true,
    CURLOPT_RETURNTRANSFER => true,
    CURLOPT_HTTPHEADER => $headers
]);

$response = curl_exec($ch);

if (curl_errno($ch)) {
    die("Request failed: " . curl_error($ch));
}

curl_close($ch);

$result = json_decode($response, true);

echo "Status: " . ($result['status'] ?? 'N/A') . PHP_EOL;

if (!empty($result['data'])) {
    foreach ($result['data'] as $item) {
        echo "Name: " . $item['personName'] . PHP_EOL;
        echo "Date: " . $item['tanggal'] . PHP_EOL;
        echo "Check In: " . $item['jamMasuk'] . PHP_EOL;
        echo "Check Out: " . $item['jamPulang'] . PHP_EOL;
        echo "Status: " . $item['statusText'] . PHP_EOL;
        echo "---------------------" . PHP_EOL;
    }
}