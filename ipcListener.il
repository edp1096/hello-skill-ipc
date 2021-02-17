procedure(proxyListener(cid data)
    println(data)
)

procedure(runSender(param)
let((proc csid result isKilled)
    sprintf(proc "proxySender -cmd=%s" param)
    csid = ipcBeginProcess(proc "" "proxyListener")

    ipcWaitForProcess(csid)
    ipcReadProcess(cid 1)
	
    ; ipcKillProcess(csid) ; I don't want to see nil printing until proxsyServer response
))

if((symeval('cid) == 'unbound)
then
    cid = ipcBeginProcess("proxyServer" "" "proxyListener")
)
