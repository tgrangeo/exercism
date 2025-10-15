package clock

import "fmt"

type Clock struct {
	hours   int
	minutes int
}

func (c *Clock) normalize() {
	c.hours += c.minutes / 60
	c.minutes %= 60
	if c.minutes < 0 {
		c.minutes += 60
		c.hours--
	}
	c.hours %= 24
	if c.hours < 0 {
		c.hours += 24
	}
}

func New(h, m int) Clock {
	clock := Clock{h, m}
	clock.normalize()
	return clock
}

func (c Clock) Add(m int) Clock {
	c.minutes += m
	c.normalize()
	return c
}

func (c Clock) Subtract(m int) Clock {
	c.minutes -= m
	c.normalize()
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}
