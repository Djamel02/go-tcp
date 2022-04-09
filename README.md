# GO-TCP
## Description
Short chat app built with go using tcp server
## Developement
run `go build .`  on your command line<br/>
run also `./tcp-server` to run the server  (on `localhost:8888`) <br/>
to make new client install `telnet` package then run `telnet localhost 8888`
### Commands
those below are our app commands: <br/>
- `/nickname <your-name>` to define the client nickname else it will give him 'anonymous' by default
- `/rooms` to display the available rooms
- `/join <room-name>` to join a specific room
- `/quit` to left the current room
