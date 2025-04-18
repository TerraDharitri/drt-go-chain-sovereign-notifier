package indexer

import "github.com/TerraDharitri/drt-go-chain-core/data/outport"

// OutportBlockCache defines a simple cache able to store *outport.OutportBlock
type OutportBlockCache interface {
	Add(outportBlock *outport.OutportBlock) error
	Extract(headerHash []byte) (*outport.OutportBlock, error)
	IsInterfaceNil() bool
}

// DataProcessor defines a payload processor for incoming ws data
type DataProcessor interface {
	ProcessPayload(payload []byte, topic string, version uint32) error
	Close() error
	IsInterfaceNil() bool
}
