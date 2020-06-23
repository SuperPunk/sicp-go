package ch03

type Stream struct {
	value int
	next  Promise // 延时计算对象，允诺在将来某个时候计算下一个值
}

type Promise func() *Stream

type StreamMapProc func(int) int

type StreamWalkProc func(int)

type StreamPredication func(int) bool

func ConsStream(value int, promise Promise) *Stream {
	return &Stream{value, memoize(promise)}
}

func StreamEnumerateInterval(low, high int) *Stream {
	if low >= high {
		return nil
	}
	return ConsStream(low, memoize(func() *Stream { return StreamEnumerateInterval(low+1, high) }))
}

func (s Stream) StreamRef(n int) int {
	if n == 0 {
		return s.value
	}
	return s.next().StreamRef(n - 1) // 指针类型的方法集合，包含了所有值方法和指针方法
}

func (s *Stream) StreamFilter(p StreamPredication) *Stream {
	if s == nil {
		return nil
	}
	if p(s.value) {
		return ConsStream(s.value, func() *Stream { return s.next().StreamFilter(p) })
	} else {
		return s.next().StreamFilter(p)
	}
}

func (s *Stream) StreamMap(p StreamMapProc) *Stream {
	if s == nil {
		return nil
	}
	return ConsStream(p(s.value), func() *Stream { return s.next().StreamMap(p) })
}

func (s *Stream) StreamForEach(p StreamWalkProc) {
	if s == nil {
		return
	}
	p(s.value)
	s.next().StreamForEach(p)
}

func memoize(p Promise) Promise {
	alreadyRun := false
	var result *Stream
	return func() *Stream {
		if !alreadyRun {
			result = p()
			alreadyRun = true
		}
		return result
	}
}
