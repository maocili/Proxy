package proxy

type Queue struct {
	ch       chan IPInfo //用于接收新ipd的chan
	Consumer func(ch <-chan IPInfo, pool *ProxyPool)
}

func NewInsertQueue(buffer int,conFunc func(i <-chan IPInfo, pool *ProxyPool) ) *Queue {
	return &Queue{
		ch:       make(chan IPInfo, buffer),
		Consumer: conFunc,
	}
}


