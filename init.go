package gache

func (g *Gache[T]) init() error {
	if g.initialized {
		return nil
	}

	err := g.load()
	if err != nil {
		return err
	}

	g.initialized = true
	return nil
}
