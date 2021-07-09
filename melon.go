package melon

import (
	"sync/atomic"
)

const (
	defaultPropFactor = 3
)

//melon use ring to set it's metrics
//sweet  bitter

type melon struct {
	//input set length
	length int64

	//real data info
	size    int64
	anchors []*anchor
	//factor means the buf ring length, to compute more accuracy
	//big is better, but will lead more memory usage.
	factor int64

	//current index
	index int64

	points []point
}

type point struct {
	val   bool
	__pad [63]byte //padding to prevent false sharing
}

type Option func(*melon)

func New(length int64, options ...Option) *melon {
	//default buff length
	size := length * defaultPropFactor
	//melon
	me := &melon{
		size:   size,
		points: make([]point, size),
		factor: defaultPropFactor,
	}
	for _, opt := range options {
		opt(me)
	}
	return me
}

func (r *melon) Reset() {
	r.points = make([]point, r.size)
	atomic.StoreInt64(&r.index, 0)
}

func (r *melon) Feed(sweet bool) {
	index := atomic.AddInt64(&r.index, 1)
	//default false means success
	//if sweet == true; store false
	r.points[index%r.size].val = !sweet
}

func (r *melon) Good() bool {
	index := atomic.LoadInt64(&r.index)
	for _, opt := range r.anchors {
		if opt.bitter(r.size, r.points, index) {
			return false
		}
	}
	return true
}

func (r *melon) Stats() []uint8 {
	var i int64
	stat := make([]uint8, r.size)
	for i = 0; i < r.size; i++ {
		if r.points[i].val {
			stat[i] = 1
		}
	}
	return stat
}
