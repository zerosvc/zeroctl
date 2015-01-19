# discovery

## discovery.node

Persistent nodes should send a message describing service every $heartbeat interval; it should be send to `discovery.node.$nodename` ; body should contain:

* `node-name`
* `node-uuid`
* `ts` -  in unixtime, optinally with ms/us part after dot (1421701103.134 for ms accuracy)
* `hb-interval` - configured interval between heartbeats. Max 120s.
* `ttl` - time after node should be considered inactive if it didn't send hb. shoudl be 3x hb-interval
* `services` - hash of services node is providing and their relevant info; same as in discovery.service



## discovery.service

Each node providing service should listen on `discovery.service.$servicename` and if it 