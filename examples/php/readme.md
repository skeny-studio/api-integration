POST /checkin.php
POST /checkout.php

Header:
X-API-Key: your-secret-key

Body JSON:
{
  "session_key":"abc123",
  "person_id":1,
  "person_name":"John Doe"
}