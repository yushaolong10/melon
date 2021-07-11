package melon

//anchor point the metric total & threshold

func OptionAnchor(total, threshold int64) Option {
	return func(r *Ring) {
		if threshold <= total && total <= r.length {
			r.anchors = append(r.anchors, newAnchor(total, threshold))
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

func (a *anchor) bitter(size int64, points []point, index int64) bool {
	start := (index - a.total + size) % size
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
	if count >= a.threshold {
		return true
	}
	return false
}
