# Allegro PCB SKILL IPC practice

## External interface tool
```powershell
# Response msg to Allegro
cd proxyServer
go build ../proxyServer.exe
```

## In Allegro telskill
```sh
ipcWriteProcess(cid "data_to_send")
telskill=> "Received-chan-ipc-0\n"
```


## Using REST API Client
* See `request.http`
```http
GET http://localhost:1323 HTTP/1.1
Browser or Rest API client=> Hello, world!!
telskill=> "Received-chan-rest-1\n"
```
