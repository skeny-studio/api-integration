<?php

$uri = parse_url($_SERVER['REQUEST_URI'], PHP_URL_PATH);

switch ($uri) {

    case '/check-in':
        require 'checkin.php';
        break;

    case '/check-out':
        require 'checkout.php';
        break;

    default:
        http_response_code(404);

        echo json_encode([
            'success' => false,
            'message' => 'Route not found'
        ]);
}