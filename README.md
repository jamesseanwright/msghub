# msghub

msghub is a TCP server written in Go for sending byte data to various recipients.


## Messages

All messages must be sent to the server as JSON, corresponding to the following contract:

* `type`: the type of request being made

The below two fields only apply to the `"sendMessage"` type:

* `userIDs`: an array of unsigned 64-bit integers (i.e. user IDs) to which the message should be sent
* `message`: an array of bytes to send to the recipients


### Message Types

* `getUserID` - returns the user ID for the requester
* `getAllUsers` - returns all of the other users who are connected to the server
* `sendMessage` - sends a message to the target users based upon the above JSON contract
* `logout` - removes a connection's user metadata. Send this before closing a TCP connection

You can use a program such as [nc](http://linux.die.net/man/1/nc) to send messages to the server:

```
echo "{ \"type\": \"getAllUsers\" }" | nc localhost 9001
```


## Local Setup

* Ensure that this zip archive is extracted to `$GOPATH/src/msghub`
* `cd $GOPATH/src/msghub`

### Go Version

msghub was developed against Go 1.6, but should be compatible with earlier versions

### Running The Tests

msghub is covered by both functional and unit tests. Run these with `go test -v ./...`.


## Running The Server

You can also run the standalone server:

* `go build`
* `./msghub`