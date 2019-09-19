package server

import "sync"

func mergeerror(errcs ...<-chan error) <-chan error {
	out := make(chan error, len(errcs))
	var wg sync.WaitGroup
	wg.Add(len(errcs))
	for _, errc := range errcs {
		go func(errc <-chan error) {
			defer wg.Done()
			for err := range errc {
				if err != nil {
					out <- err
					return
				}
			}
		}(errc)
	}
	go func() {
		close(out)
		wg.Wait()
	}()
	return out
}

func waitPipeline(errs ...<-chan error) error {
	errc := mergeerror(errs...)
	for err := range errc {
		if err != nil {
			return err
		}
	}
	return nil
}
