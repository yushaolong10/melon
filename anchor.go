package melon

//anchor point the metric length & threshold

func OptionAnchor(total, threshold int64) Option {
	return func(me *melon) {
		if threshold <= total && total <= me.size {
			me.anchors = append(me.anchors, newAnchor(total, threshold))
		}
	}
}

func newAnchor(total, threshold int64) *anchor {
	return &anchor{
		total:     total,
		threshold: threshold,
	}
}

type anchor struct {
	total     int64
	threshold int64
}

func (o *anchor) bitter(size int64, points []point, index int64) bool {
	start := (index - o.total + size) % size
	end := index % size
	if start > end {
		end = end + size
	}
	var count int64
	for i := start; i < end; i++ {
		if points[i%size].val {
			count++
		}
	}
	if count >= o.threshold {
		return true
	}
	return false
}
