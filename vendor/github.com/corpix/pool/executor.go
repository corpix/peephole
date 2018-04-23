package pool

import (
	"context"
)

type Executor func(context.Context)
type ExecutorWithResult func(context.Context) (v interface{}, err error)
