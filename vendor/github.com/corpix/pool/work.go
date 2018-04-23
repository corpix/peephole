package pool

import (
	"context"
)

type Work struct {
	Context  context.Context
	Executor Executor
}

func executorWithResultToExecutor(result chan<- *Result, executor ExecutorWithResult) Executor {
	return func(context context.Context) {
		v, err := executor(context)
		if err != nil {
			result <- NewResult(nil, err)
			return
		}
		result <- NewResult(v, nil)
	}
}

func NewWork(context context.Context, executor Executor) *Work {
	return &Work{
		Context:  context,
		Executor: executor,
	}
}

func NewWorkWithResult(context context.Context, result chan<- *Result, executor ExecutorWithResult) *Work {
	return &Work{
		Context:  context,
		Executor: executorWithResultToExecutor(result, executor),
	}
}
