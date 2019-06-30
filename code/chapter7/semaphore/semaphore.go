// This sample program demonstrates how to implement a semaphore using
// channels that can allow multiple reads but a single write.
//
// It uses the generator pattern to create channels and goroutines.
//
// Multiple reader/writers can be created and run concurrently. Then after
// a timeout period, the program shutdowns cleanly.
//
// http://www.golangpatterns.info/concurrency/semaphores

// 这个demo演示了如何使用channel实现信号量,改信号量允许多个读操作但是只允许一个写操作
// It uses the generator pattern to create channels and goroutines.
package main

type(
	semaphore chan struct{}
)

type (
	// readerWriter provides a structure for safely reading and writing a shared resource.
	// It supports multiple readers and a single writer goroutine using a semaphore construct.
	readerWriter struct {
		// The name of this object.
		name string

		// write forces reading to stop to allow the write to occur safely.
		write sync.WaitGroup

		// readerControl is a semaphore that allows a fixed number of reader goroutines
		// to read at the same time. This is our semaphore.
		readerControl semaphore

		// shutdown is used to signal to running goroutines to shutdown.
		shutdown chan struct{}

		// reportShutdown is used by the goroutines to report they are shutdown.
		reportShutdown sync.WaitGroup

		// maxReads defined the maximum number of reads that can occur at a time.
		maxReads int

		// maxReaders defines the number of goroutines launched to perform read operations.
		maxReaders int

		// currentReads keeps a safe count of the current number of reads occurring
		// at any give time.
		currentReads int32
	}
)

