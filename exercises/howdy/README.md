# Location API

This Rails API provides user location and language detection functionality.

## Endpoint

### GET /whereami

Retrieves user location and language information based on their IP address and browser preferences.

#### Response

Returns a JSON object with the following structure:

```json
{
  "ip": "172.217.2.46",
  "country": "US",
  "language": "en-us"
}
```

#### Fields

- **ip**: The user's IP address (string or null if not detectable)
- **country**: ISO country code derived from the IP address (string or null if not detected)
- **language**: Preferred language from the Accept-Language header (string, defaults to "en")

#### Example Requests

```bash
# Using curl
curl -X GET http://localhost:3000/whereami

# Using curl with specific language preference
curl -X GET http://localhost:3000/whereami \
  -H "Accept-Language: es-ES,es;q=0.9,en;q=0.8"
```

#### Example Response

```json
{
  "ip": "203.0.113.42",
  "country": "US",
  "language": "en-us"
}
```

#### Notes

- The endpoint uses ip-api.com for IP geolocation (free tier)
- Private/localhost IPs will return null for country detection
- Language detection uses the HTTP Accept-Language header
- Returns HTTP 200 status code on success

## Setup

1. Ensure you have the required dependencies in your Rails application
2. Add the routes to your `config/routes.rb`
3. The controller handles IP detection, country lookup, and language preference parsing
4. No additional gems required (uses built-in Net::HTTP and JSON libraries)

## Error Handling

- Invalid IPs or API failures will return null for the country field
- Missing Accept-Language headers default to "en"
- The endpoint is designed to always return a 200 status with partial data rather than failing
