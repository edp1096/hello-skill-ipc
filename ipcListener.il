procedure(proxyListener(cid data)
    ; println(concat("Received " data))
    ; println(concat("Received: " data))
    println(data)
)

; Prevent run twice
if((symeval('cid) == 'unbound)
then
    cid = ipcBeginProcess("proxyServer" "" "proxyListener")
)
