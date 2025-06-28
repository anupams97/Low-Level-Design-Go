```json
[
  {
    "asset_id": "TRAILER-001",
    "timestamp": "2024-06-01T10:00:00Z",
    "lat": 37.7749,
    "lng": -122.4194,
    "battery_level": 85,
    "is_moving": false
  },
  {
    "asset_id": "TRAILER-002",
    "timestamp": "2024-06-01T10:01:00Z",
    "lat": 37.7788,
    "lng": -122.4150,
    "battery_level": 76,
    "is_moving": true
  }
]
```
1. setup a http server with the endpoints
2. create the handlers for those endpoints
3. write code to store in inDB or in memory 
4. write code to fetch from inDB
5. write middleware
6. validate request
7. connection pooling 
8. singleton db connection
9. repo interface 
10. adapter to convert the format to ping data
11. middleware patter by go middleware pattern or classic decorator patter  n
12. observer pattern to pupblish data to ui dashboard or kafka
13. strategy pattern for alert 

