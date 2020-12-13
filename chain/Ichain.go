package chain

import (
	"context"

	"github.com/golang/protobuf/proto"
)

type Module interface {
	ModuleName() string
	BeginBlock(ctx context.Context, message proto.Message) error
	EndBlock(ctx context.Context) (proto.Message, error)
	LoadHeight(height uint64) error
}
