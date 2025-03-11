package chronus

import (
	"sync"
	"time"
)

var offsetEnable = true

var instMap = sync.Map{}

// SetOffsetEnable set global flag allowed time offsets
func SetOffsetEnable(enable bool) {
	offsetEnable = enable
}

type Chronus struct {
	id        string
	offset    time.Duration
	callbacks []SetOffset
}

type SetOffset interface {
	SetOffset(offset time.Duration)
}

// GetByID Get Chronus by id
func GetByID(id string) *Chronus {
	value, ok := instMap.Load(id)
	if !ok {
		return nil
	}
	return value.(*Chronus)
}

func NewChronus(id string, offset time.Duration) *Chronus {
	c := &Chronus{id: id, offset: offset}
	instMap.Store(id, c)
	return c
}

func (c *Chronus) Now() time.Time {
	if !offsetEnable {
		return time.Now()
	}
	return time.Now().Add(c.offset)
}

func (c *Chronus) SetOffset(offset time.Duration) {
	c.offset = offset
	for _, cb := range c.callbacks {
		cb.SetOffset(offset)
	}
}

func (c *Chronus) AddOffsetCallback(cb SetOffset) {
	c.callbacks = append(c.callbacks, cb)
}

func (c *Chronus) NewCron(opts ...Option) *Cron {
	cron := New(opts...)
	cron.chronus = c
	c.AddOffsetCallback(cron)
	return cron
}
