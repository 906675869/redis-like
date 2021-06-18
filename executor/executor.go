package executor

import "context"

type Executor interface {
	Execute(ctx context.Context, invocation InvocationInter) (interface{}, error)
}

type SimpleExecutor struct {
	invoker InvokerInter
}

func (s *SimpleExecutor) Execute(ctx context.Context, invocation InvocationInter) (interface{}, error) {
	invoker := s.invoker
	var result ResultInter
	for {
		result = invoker.Invoke(ctx, invocation)
		if result.Success() || !invoker.hasNext() {
			break
		}
		invocation.AddCallbacks(invoker.Callback())
		invoker = invoker.next()
	}
	invocation.OnFinished(ctx, result)
	return result.Result(), result.Error()
}
