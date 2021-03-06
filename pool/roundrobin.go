package pool

import (
	"fmt"
	"sync"

	"github.com/codescot/admock/logger"
	"github.com/codescot/go-common/env"
)

// RoundRobin a pool representing a round robin list
type RoundRobin struct {
	mux       sync.Mutex
	upstreams []string
	size      int
	index     int
}

const cloudflareDNS1 = "1.1.1.1"
const cloudflareDNS2 = "1.0.0.1"

// NewRoundRobin pool with dns servers
func NewRoundRobin() Pool {
	dns1 := env.Optional("ADMOCK_DNS1", cloudflareDNS1)
	dns2 := env.Optional("ADMOCK_DNS2", cloudflareDNS2)

	logger.PrintWithTimeStamp(fmt.Sprintf("DNS1: %s", dns1))
	logger.PrintWithTimeStamp(fmt.Sprintf("DNS2: %s", dns2))

	list := []string{}
	if len(dns1) > 0 {
		list = append(list, dns1)
	}

	if len(dns2) > 0 {
		list = append(list, dns2)
	}

	return &RoundRobin{
		mux:       sync.Mutex{},
		upstreams: list,
		size:      len(list),
		index:     0,
	}
}

// Next get next in pool
func (rr *RoundRobin) Next() string {
	rr.mux.Lock()
	defer rr.mux.Unlock()

	rr.increment()
	return rr.current()
}

func (rr *RoundRobin) current() string {
	return rr.upstreams[rr.index]
}

func (rr *RoundRobin) increment() int {
	nextIndex := rr.index

	rr.index++
	if rr.index == rr.size {
		rr.index = 0
	}

	return nextIndex
}
