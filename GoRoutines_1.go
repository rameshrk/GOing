package main // a channel to tell it to stop
stopchan := make(chan struct{})

// a channel to signal that it's stopped
stoppedchan := make(chan struct{})

go func () { // work in background

	// close the stoppedchan when this func
	// exits
	defer close(stoppedchan)

	// TODO: do setup work
	defer func() {
		// TODO: do teardown work
	}()

	for {
		select {
		default:

			// TODO: do a bit of the work

		case <-stopchan:
			// stop
			return
		}
	}

}()
