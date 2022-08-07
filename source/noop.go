package source

type noopWatcher struct {
	exit chan struct{}
}

func (w *noopWatcher) Next() (*ChangeSet, error) {
	<-w.exit

	return nil, ErrWatcherStopped
}

func (w *noopWatcher) Stop() error {
	close(w.exit)
	return nil
}

// NewNoopWatcher returns a watcher that blocks on Next() until Stop() is called.
func NewNoopWatcher() (Watcher, error) {
	return &noopWatcher{exit: make(chan struct{})}, nil
}
