package storage

import (
	"context"
	"sync"
)

// Result represents the Size function result
type Result struct {
	Size  int64 // Total Size of File objects
	Count int64 // Count is a count of File objects processed
}

type DirSizer interface {
	// calculates the size of a given Dir
	Size(ctx context.Context, d Dir) (Result, error)
}

// sizer implement the DirSizer interface
type sizer struct {
	maxWorkersCount int // maximum number of goroutines
}

// NewSizer returns new DirSizer instance
func NewSizer() DirSizer {
	return &sizer{maxWorkersCount: 5}
}

func (a *sizer) Size(ctx context.Context, d Dir) (Result, error) {
	result := Result{} // to hold final file count and size
	var mu sync.Mutex
	var wg sync.WaitGroup

	fileCh := make(chan File, a.maxWorkersCount) // chan for sending files to workers
	errrosCh := make(chan error, 1)              // error chan to show the 1st encountered error
	doneCh := make(chan struct{})                // to signal workers that traversal is complete

	// start worker pool with maxWorkersCount
	for i := 0; i < a.maxWorkersCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done() // mark worker done when it exits
			for {
				select {
				case <-ctx.Done(): // exit when ctx is canceled
					return
				case <-doneCh: // exit if traversal is complete
					return
				case file, ok := <-fileCh: // get file from fileCh
					if !ok { // exit if chan is closed
						return
					}
					size, err := file.Stat(ctx) // get file size
					if err != nil {
						select {
						case errrosCh <- err: // send error to errorsCh
						default: // ignore further errors after the first one
						}
						return
					}
					mu.Lock()           // lock result
					result.Size += size // update total size
					result.Count++      // update file count
					mu.Unlock()         // unlock result
				}
			}
		}()
	}

	// traverse directories and enqueue files
	traverseWg := sync.WaitGroup{} // wg to track directory
	traverseWg.Add(1)
	go func() {
		defer close(fileCh)                        // close fileCh when traverse complete
		a.traverseDir(ctx, d, fileCh, &traverseWg) // start directory traversal
		traverseWg.Wait()                          // wait for all goroutines
	}()

	// wait for all workers
	go func() {
		wg.Wait()
		close(doneCh)
	}()

	select {
	case err := <-errrosCh: // return the 1st error
		return Result{}, err
	case <-doneCh: // return result if there's no errors
		return result, nil
	}
}

// traverseDir recursively adds files from the directory and its subdirectories to filesCh
func (a *sizer) traverseDir(ctx context.Context, d Dir, filesCh chan<- File, wg *sync.WaitGroup) {
	defer wg.Done()

	dirs, files, err := d.Ls(ctx) // list directories and files in current dir
	if err != nil {
		return
	}

	// send file to the fileCh
	for _, file := range files {
		select {
		case <-ctx.Done(): // exit if ctx is canceled
			return
		case filesCh <- file: // send file to the channel for worker processing
		}
	}

	for _, file := range dirs {
		wg.Add(1)
		go a.traverseDir(ctx, file, filesCh, wg)
	}
}
