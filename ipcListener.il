procedure(proxyListener(cid data)
    println(data)
)

cid = ipcBeginProcess("proxyServer" "" "proxyListener")
