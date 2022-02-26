# fight-irl
JSON API for giving directions to you and the person youre arguing with online to a meeting place where you can fight in person, based on your ip addresses

## Endpoints
In the following, substitute `ip_address` with a valid ipv4/ipv6 address

`GET /`
---
returns location information about the ip address making the request:
```
{
  "your_address_information": {
    "ip": string,
    "coordinates": {
      "latitude": number,
      "longitude": number
    },
    "address": string
  }
}
```

`GET /ip/ip_address`
---
returns location information about `ip_address`:
```
{
  "their_address_information": {
    "ip": string,
    "coordinates": {
      "latitude": number,
      "longitude": number
    },
    "address": string
  }
}
```

`GET /meet/ip_address`
---
returns an intermediate meeting location (the geographic midpoint between the location of `ip_address` and the address making the request) as well as directions
from both places to this location:
```
{
  "meeting": {
    "meeting_location": {
      "latitude": number,
      "longitude": number
    },
    "your_start_address": string,
    "their_start_address": string,
    "your_directions": {
      "steps": [
        {
          "html_instructions": string,
          "distance": string"
        },
        {
          "html_instructions": string,
          "distance": string"
        },
        ...
        {
          "html_instructions": string,
          "distance": string
        }
      ]
    },
    "their_directions": {
      "steps": [
        {
          "html_instructions": string,
          "distance": string"
        },
        ...
      ]
    }
  }
}
```

### Error Responses
Any of the above endpoints may return the following:
```
{
  "error": string
}
```
