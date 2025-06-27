package util

import "maps"

type Set[T comparable] struct {
	elements map[T]struct{}
}

func NewSet[T comparable]() Set[T] {
	set := Set[T]{
		elements: make(map[T]struct{}),
	}

	return set
}

func MakeSetFromSlice[T comparable](slice []T) Set[T] {
	set := NewSet[T]()
	for _, s := range slice {
		set.Add(s)
	}
	return set
}

func (s Set[T]) Add(p T) {
	s.elements[p] = struct{}{}
}

func (s Set[T]) Delete(p T) {
	delete(s.elements, p)
}

func (s Set[T]) IsThere(p T) bool {
	_, it := s.elements[p]
	return it
}

func (s Set[T]) Elements() []T {
	var elems []T
	for e := range s.elements {
		elems = append(elems, e)
	}

	return elems
}

func (s Set[T]) Count() int {
	return len(s.elements)
}

func (s Set[T]) DeepCopy() Set[T] {
	new := NewSet[T]()
	maps.Copy(new.elements, s.elements)
	return new
}

// use: for subset := range set.ForEachSubset() { process(subset) }
func (s Set[T]) ForEachSubset() <-chan Set[T] {
	ch := make(chan Set[T])
	go func() {
		var i int
		var subset Set[T]
		var elements = s.Elements()
		numOfElems := len(elements)
		conti := true
		choice := make([]bool, numOfElems)
		i = 0
		for i < numOfElems {
			choice[i] = true
			i++
		}

		for conti {
			i = 0
			conti = false
			subset = NewSet[T]()
			for i < numOfElems {
				if choice[i] {
					subset.Add(elements[i])
					if !conti {
						choice[i] = false
					}
					conti = true
				} else {
					if !conti {
						choice[i] = true
					}
				}
				i++
			}

			ch <- subset
		}

		close(ch)
	}()
	return ch
}
