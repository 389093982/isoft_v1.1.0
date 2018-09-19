package heartbeat

import "isoft/isoft_storage/lib"

func StartHeartbeat() {
	proxy := &lib.LocateAndHeartbeatProxy{}
	proxy.SendHeartbeat()
}
