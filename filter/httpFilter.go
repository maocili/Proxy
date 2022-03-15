package filter

import "proxy/internal/proxy"

func HttpFilter(info proxy.IPInfo) bool {
	if info.Rating >= 60 && info.IPType == proxy.HTTP {
		return true
	}
	return false
}

func HttpsFilter(info proxy.IPInfo) bool {
	if info.Rating >= 60 && info.IPType == proxy.HTTPS {
		return true
	}
	return false
}
