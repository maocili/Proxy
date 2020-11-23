package proxy


func (p *ProxyPool) FilterIP(filterFunc func(info IPInfo) bool) (info IPInfo) {
	p.m.Lock()
	defer p.m.Unlock()

	for _, info := range p.ips {
		if filterFunc != nil && filterFunc(info) {
			return info
		}
	}
	return info

}

