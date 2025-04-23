package pb

import (
	context "context"
	"fmt"

	"github.com/lwmacct/250300-go-mod-mlog/pkg/mlog"
)

type ImplExamples struct {
	UnimplementedRpcServer
}

func (s *ImplExamples) Int64AskString(ctx context.Context, in *KvInt64) (*KvString, error) {
	out := &KvString{
		Key: "res",
		Val: fmt.Sprintf("%d", in.Val),
	}
	mlog.Info(mlog.H{
		"in":  in,
		"out": out,
	})
	return out, nil
}
