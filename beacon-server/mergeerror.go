package server

// mergeerror merges error channels into one channel, remember to
// close the channel when done reading from it
func mergeerror(errcs ...<-chan error) chan error {
	out := make(chan error)
	for _, errc := range errcs {
		go func(errc <-chan error) {
			defer recover()
			for err := range errc {
				if err != nil {
					out <- err
				}
			}
		}(errc)
	}
	return out
}
