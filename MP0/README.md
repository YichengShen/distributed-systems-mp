# MP0 for Distributed Systems

This MP aims to send emails from the TCP client to the TCP server.

## To run
1. To run the TCP server:  

    ```cd cmd/server```

    ```go run main.go```

2. To run the TCP client:  

    ```cd cmd/client```

    ```go run main.go```
    
## GO Files

* **messages.go** contains the Msg struct which is used to distinguish different types of messages.
* **email.go** contains the Email struct.
* **tcp_server.go** contains the server process and its helper functions.
* **tcp_client.go** contains the client process and its helper functions.
* **utils.go** contains the functions about ACK shared by both of the server and the client.

## Structs

* Msg
    * Type (the type of a message)
    * Content (the Email struct)
    
    If Msg.Type is email, then Msg.Content will be the Email struct. If Msg.Type is acknowledgement, then Msg.Content will be left empty.
    
* Email
    * To
    * From
    * Date
    * Title
    * Content
    
## Major Packages Used
* [gob](https://golang.org/pkg/encoding/gob/) is used to send and receive streams of data through connections.
* [net](https://golang.org/pkg/net/) is used for TCP.

## References
1. Sending struct [using encoding/gob](https://stackoverflow.com/questions/11202058/unable-to-send-gob-data-over-tcp-in-go-programming) 
2. [TCP tutorial](https://www.linode.com/docs/guides/developing-udp-and-tcp-clients-and-servers-in-go/)
3. [Golang standard project layout](https://github.com/golang-standards/project-layout/tree/master/cmd) for organizing files
4. [Graceful shutdown](https://eli.thegreenplace.net/2020/graceful-shutdown-of-a-tcp-server-in-go/) of the TCP server
