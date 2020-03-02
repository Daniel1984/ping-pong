### Tests
- in order to run tests, from root dir run `go test ./...`

### Code structure
- programs/executables live in `/cmd` directory. Currently we have `server` and `client` in there
- all your building blocks/shared packages live in `/pkg` directory

### Instructions
- running server:
  - from root dir run `go run cmd/server/main.go` or `go run -race cmd/client/main.go` to ensure there's no race conditions

- running client(s):
  - from root dir run `go run cmd/client/main.go` or `go run -race cmd/client/main.go` and follow on screen instructions.
  - client assumes that user has at least 1 friend in order to test message passing when joining/leaving server
  - server will not let user join when user with given ID is already online

### Q&A
- What changes if we switch TCP to UDP?
  - It would get faster, but we would lose guaranteed delivery of data. Datagram packet may become corrupt or lost in transit. It has no flow control, packets arrive in a continuous stream or they are dropped. There would be no guarantee that packets sent from a server would be delivered to the client in the same order. 
- How would you detect the user (connection) is still online?
  - in current design we constantly try to read/write to the connection so if client, for any reason, is lost, that will end up in server capturing error for that particullar connection and we remove it from list of clients. Any further references of connection by ID(how it's currenlty stored) will fail allocating it, meaning connection is offline.
- What happens if the user has a lot of friends?
  - We would need to make sure that instance, server runs on, has enough resources or have multiple instances of the server running. Meaning we would need to scale vertically or horizontally.
- How design of your application will change if there will be more than one instance of the server
  - Would need to introduce a message broker / pub sub. Could be something like rabbitmq, redis, kafka etc. Each instance would be subscribed to centralized broker and would be writing and reading messages from in order to communicate between the clients that exist on different instances.

### TODO
- profiling
