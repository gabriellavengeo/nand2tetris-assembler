package instruction

// Label represents an label pseudo instruction
type Label interface {
	Label() string
}