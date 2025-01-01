## Server

### 1. ListenAndServe

- [] What the heck
    It means all in its name, this function listens on the TCP network address, then call 
    Serve with "handler" to handle requests incoming connections. 
    ListenAndServe always return a non nil error
- [] Param
    - [] address: host ad port
    - [] handler: how the services handle incomming requests
- [] ListenAndServeTLS: 
    

### 2. ServerMux

- [] pattern
- [] precedence
- [] Hanlde()

### 3. Handler

- [] Handler
- [] HandlerFunc

### 4. Middleware
- [] Example 
    - [] logger
    - [] requestracking_id

### 5. Error handling
- [] Status, Message
- [] Add to middleware
