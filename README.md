# Allegro PCB SKILL IPC practice

## Tools
```powershell
# Response msg to Allegro
cd proxyServer
go build ../proxyServer.exe

cd ..

# Send msg to proxyServer (from proxySender) then response msg to Allegro (from proxyServer)
cd proxySender
go build ../proxySender.exe
```

Because Allegro often recognize child process which spawned by `ipcBeginProcess` as dead so, need `proxySender` when send to `proxtServer`

## In Allegro telskill
```sh
runSender("hello_no_space")
telskill=> "hello_no_space"
telskill=> "hello^_^- 0\n"
```

## Using REST API Client
* See `request.http`
```http
GET http://localhost:1323 HTTP/1.1
Browser or Rest API client=> Hello, world!!
telskill=> "hello^_^- 1\n"
```
