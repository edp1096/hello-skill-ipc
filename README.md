# Allegro PCB SKILL IPC practice

## proxyServer - External interface tool
```powershell
cd proxyServer
go build -o ..
```


## In Allegro telskill
```lisp
; in skill console
ipcWriteProcess(cid "data_to_send\n")
telskill=> "Received: chan-ipc-0\n"
```


## Using REST API Client
* See `request.http`
```sh
GET http://localhost:1323 HTTP/1.1
Browser or Rest API client=> Hello, world!!
telskill=> "Received :chan-rest-1\n"
```
