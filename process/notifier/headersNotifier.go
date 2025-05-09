package notifier

import (
	"encoding/hex"
	"sync"

	"github.com/TerraDharitri/drt-go-chain-core/core/check"
	"github.com/TerraDharitri/drt-go-chain-core/data/sovereign"
	"github.com/TerraDharitri/drt-go-chain-sovereign-notifier/process"
)

type headersNotifier struct {
	mutSubscribers sync.RWMutex
	subscribers    []process.IncomingHeaderSubscriber
}

func newHeadersNotifier() *headersNotifier {
	return &headersNotifier{
		mutSubscribers: sync.RWMutex{},
		subscribers:    make([]process.IncomingHeaderSubscriber, 0),
	}
}

func (hn *headersNotifier) registerSubscriber(handler process.IncomingHeaderSubscriber) error {
	if check.IfNil(handler) {
		return errNilHeaderSubscriber
	}

	hn.mutSubscribers.Lock()
	hn.subscribers = append(hn.subscribers, handler)
	hn.mutSubscribers.Unlock()

	return nil
}

func (hn *headersNotifier) notifyHeaderSubscribers(header sovereign.IncomingHeaderHandler, headerHash []byte) error {
	log.Debug("notifying incoming header", "hash", hex.EncodeToString(headerHash))

	hn.mutSubscribers.RLock()
	defer hn.mutSubscribers.RUnlock()

	for _, handler := range hn.subscribers {
		err := handler.AddHeader(headerHash, header)
		if err != nil {
			return err
		}
	}

	return nil
}
