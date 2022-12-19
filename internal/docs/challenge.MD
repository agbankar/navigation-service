##### 1. Implement unique users visits counter service per user id and url.
##### 2. Implement HTTP server with two endpoints:
   -- POST /visit - accept json record from examples provided.
   -- GET /info?url=string - display total uniq visitors for requested url.



Example POST requests:
````
{“uid”: “user1", “url”: “http://localost/page-1”}
{“uid”: “user1", “url”: “http://localost/page-2”}
{“uid”: “user2", “url”: “http://localost/page-1”}
{“uid”: “user3", “url”: “http://localost/page-3”}
{“uid”: “user4", “url”: “http://localost/page-5”}
{“uid”: “user5", “url”: “http://localost/page-5”}
etc...
````
