package cos418_hw1_1

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
)

// Sum numbers from channel `nums` and output sum to `out`.
// You should only output to `out` once.
// Do NOT modify function signature.

func sumWorker(nums chan int, out chan int) {
	// TODO: implement me
	// HINT: use for loop over `nums`
	s := 0
	for i := range nums {
		s += i
	}
	out <- s
}

// Read integers from the file `fileName` and return sum of all values.
// This function must launch `num` go routines running
// `sumWorker` to find the sum of the values concurrently.
// You should use `checkError` to handle potential errors.
// Do NOT modify function signature.
func sum(num int, fileName string) int {
	// TODO: implement me
	// HINT: use `readInts` and `sumWorkers`
	// HINT: used buffered channels for splitting numbers between workers
	f, err := os.Open(fileName)
	if err != nil {
		log.Panic("we are panicking boys")
	}
	ints, err := readInts(f)
	nums := make(chan int, len(ints))
	out := make(chan int, num)

	var wg sync.WaitGroup
	for w := 0; w < num; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sumWorker(nums, out)
		}()
	}
	for _, v := range ints {
		nums <- v
	}
	close(nums)
	wg.Wait()
	close(out)

	//return <-out
	//return 0
	sumr := 0
	for s := range out {
		sumr += s
	}
	return sumr

}

// Read a list of integers separated by whitespace from `r`.
// Return the integers successfully read with no error, or
// an empty slice of integers and the error that occurred.
// Do NOT modify this function.
func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var elems []int
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return elems, err
		}
		elems = append(elems, val)
	}
	return elems, nil
}
