package testscommon

import (
	"github.com/TerraDharitri/drt-go-chain-core/data/outport"
	"github.com/TerraDharitri/drt-go-chain-sovereign-notifier/process"
)

// SovereignNotifierStub -
type SovereignNotifierStub struct {
	NotifyCalled          func(finalizedBlock *outport.OutportBlock) error
	RegisterHandlerCalled func(handler process.IncomingHeaderSubscriber) error
}

// Notify -
func (sn *SovereignNotifierStub) Notify(finalizedBlock *outport.OutportBlock) error {
	if sn.NotifyCalled != nil {
		return sn.NotifyCalled(finalizedBlock)
	}

	return nil
}

// RegisterHandler -
func (sn *SovereignNotifierStub) RegisterHandler(handler process.IncomingHeaderSubscriber) error {
	if sn.RegisterHandlerCalled != nil {
		return sn.RegisterHandlerCalled(handler)
	}

	return nil
}

// IsInterfaceNil -
func (sn *SovereignNotifierStub) IsInterfaceNil() bool {
	return sn == nil
}
