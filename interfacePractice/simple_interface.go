package interfacePractice

type Simpler interface {
	Get() int
	Set(i int)
}

type Simple struct {
	Salary int
}

func (s *Simple) Get() int {
	return s.Salary
}

func (s *Simple) Set(i int) {
	s.Salary = i
}

func InterfaceMain(s Simpler, sala int) int {
	s.Set(sala)
	return s.Get()

}
