package movement

type Caterpillar struct {
	position float64
}

func (c *Caterpillar) Crawl(distanceMeters float64) error {
	c.position += distanceMeters

	return nil
}

type Eagle struct {
	position float64
}

func (e *Eagle) Fly(distanceMeters float64) error {
	e.position += distanceMeters

	return nil
}
