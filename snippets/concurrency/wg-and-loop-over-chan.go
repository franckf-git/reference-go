func closest(particles []Particle) Particle {
    // snip...
	for _, particle := range particles {
		go move(particle, 10000, pch, &wg)
	}

	wg.Wait()
	close(pch)

    for p := range pch {
        findcp(cp)
    }

    return cp
}
