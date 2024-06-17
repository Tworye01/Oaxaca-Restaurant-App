package app

type Method int

const (
	Get     Method = iota // 0
	Post                  // 1
	Patch                 // 2
	Put                   // 3
	Delete                // 4
	Options               // 5
	Head                  // 6
	Static                // 7
)
