package main

type DefaultSet map[int]map[int]struct{}

func (d *DefaultSet) Get(key int) []int {

	if *d == nil {
		*d = make(map[int]map[int]struct{})
	}
	if _, ok := (*d)[key]; !ok {
		(*d)[key] = make(map[int]struct{})
	}

	values := make([]int, 0)

	for k := range (*d)[key] {
		values = append(values, k)
	}

	return values
}

func (d *DefaultSet) Add(key, value int) bool {
	if *d == nil {
		*d = make(map[int]map[int]struct{})
	}

	if _, ok := (*d)[key]; !ok {
		(*d)[key] = make(map[int]struct{})
	}

	(*d)[key][value] = struct{}{}

	return true
}

func (d *DefaultSet) Exists(key, value int) bool {
	values := d.Get(key)
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}
