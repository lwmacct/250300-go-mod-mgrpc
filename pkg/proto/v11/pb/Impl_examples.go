package pb

import (
	context "context"
	"fmt"

	"github.com/lwmacct/250300-go-mod-mlog/pkg/mlog"
)

type ImplExamples struct {
	UnimplementedRpcServer
}

func (s *ImplExamples) Int64AskString(ctx context.Context, in *KvOneof) (*KvOneof, error) {
	out := &KvOneof{
		Key: "res",
		Val: &KvOneof_Str{
			Str: fmt.Sprintf("%d", in.Val),
		},
	}
	mlog.Info(mlog.H{
		"in":  in,
		"out": out,
	})
	return out, nil
}
