package myset

var exists = true

//Set simpel set
type Set struct {
	m map[int]bool
}

//NewSet return set
func NewSet() *Set {
	s := &Set{}
	s.m = make(map[int]bool)
	return s
}

//GetInnerMap return inner map
func (s *Set) GetInnerMap() map[int]bool {
	return s.m
}

//Add add item
func (s *Set) Add(value int) {
	s.m[value] = exists
}

//Remove remove item
func (s *Set) Remove(value int) {
	delete(s.m, value)
}

//Contains check item
func (s *Set) Contains(value int) bool {
	_, c := s.m[value]
	return c
}

// func main() {
// 	s := NewSet()

// 	s.Add("Peter")
// 	s.Add("David")

// 	fmt.Println(s.Contains("Peter"))  // True
// 	fmt.Println(s.Contains("George")) // False

// 	s.Remove("David")
// 	fmt.Println(s.Contains("David")) // False
// }
