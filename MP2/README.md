# Spring 2021 Distributed Systems: MP2 #

A simple chat room application that supports only private message.

**Author:** [Yicheng Shen](https://github.com/YichengShen)

## To Run ##

### Run Server ###
```bash
go run server.go --port=[port number]
```   
Fill in `[port number]`. Or use default port number, `8080`, if that argument is left blank.

To terminate the server, type `EXIT` in command-line.


### Run Client ###
```bash
go run client.go --ip=[server ip] --port=[server port number] --username=[your username]
```  
The default ip, port, and username are `127.0.0.1`, `8080`, and `u1` respectively.    

You can ask the client to do the following things in command-line. 

1. Type `[receiver username] [message]` to send a chat message

2. Type `EXIT` to terminate

## Go Structs ##
1. Message
```go
type Message struct {
	To string
	From string
	Content string
}
```
The Message struct is defined for transferring messages between TCP server and clients. `To` contains the username of the receiver. `From` contains the username of the sender. 

2. Server
```go
type Server struct {
	Ln   net.Listener
	Quit chan interface{}
	Addr string
	Conns map[string]net.Conn
}
```
The Server struct is used in `tcp/server.go`. `Quit` is a channel used for terminating the server. `Conns` contains a map from the usernames of clients to the corresponding connections. 

## Architecture ##
### Server
1. After the server runs, it constantly accepts new connections with clients in a goroutine.
2. It creates a separate goroutine to handle a connection with a client.
    1. It receives the initial message from a client which always contains the client's username. It stores that connection using the username as the key.
    2. It then enters a loop to receive the following messages.
        * On receiving the client's EXIT message, delete the client from the connection map and close the connection.
        * On receiving the client's send message request, check the receiver's status.
            * If receiver is connected, send the message to the receiver.
            * If receiver is not connected, send error message to the sender.
3. In the main thread, wait for user's command-line input. On 'EXIT', close the `Quit` channel to notify the termination of the server.

### Client
1. The client connects to the server.
2. It first sends an initial message to the server containing its username.
3. It makes `quitChan` and `msgChan` channels.   
4. It launches a goroutine to handle messages from the server.
    * On receiving error message about the receiver not connected, print send failed.
    * On receiving the server's EXIT message, close the `quitChan` to notify termination.
    * On receiving a normal chat message, print the message and where it's from.
5. In the main thread, it enters a loop that deals with sending messages and exiting.
    1. It launches a goroutine to handle command-line inputs. (Use goroutine here to prevent blocking.)
        * On user inputting 'EXIT', close the `quitChan`.
        * On user inputting a normal chat message, put the message into `msgChan`.
    2. It blocks and waits for `quitChan` or `msgChan` to be triggered.
        * If `quitChan` is closed, exit the program.
        * If new message from `msgChan`, send the new message.
## References ##

1. [Signal channel](https://medium.com/@matryer/golang-advent-calendar-day-two-starting-and-stopping-things-with-a-signal-channel-f5048161018)

2. [Go Concurrency Patterns: Pipelines and cancellation](https://blog.golang.org/pipelines)
