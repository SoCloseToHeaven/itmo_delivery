package state

type UserState struct {
	UserID    int
	UserState State
}

type State uint32

const (
	Default State = iota
)
