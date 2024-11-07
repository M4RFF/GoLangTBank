package executor

import (
	"context"
)

type (
	In  <-chan any
	Out = In
)

type Stage func(in In) (out Out)

func ExecutePipeline(ctx context.Context, in In, stages ...Stage) Out {
	for _, stage := range stages { // go over each stage
		in = runStage(ctx, in, stage) // for each stage passing in as input
	}
	return in // return the final output for the last stage
}

func runStage(ctx context.Context, in In, stage Stage) Out {
	out := make(chan any, 3) // an output chan for passing data to the next pipeline

	go func() {
		defer close(out)

		var stageOut Out = stage(in) // passing an input chan to the stage func

		for v := range stageOut {
			select {
			case <-ctx.Done():
				return // if ctx is canceled then exit
			case out <- v: // passing data the out chan for the next pipeline
			}
		}
	}()

	return out // return out chan for current stage, it's going to be an input for the next stage
}
