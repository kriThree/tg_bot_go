package statemanager

import "context"

type CreatingParams struct {
	Name string
	Tag  string
	Mean string
}
type BotContext struct {
	context.Context
}

type State struct {
	usersState map[int]UserState
}

type UserState struct {
	TgID      int
	Operation string
	Creatng   CreatingParams
	DbId      string
}

func New() *State {
	return &State{
		usersState: make(map[int]UserState),
	}
}

func (s *State) GetUser(id int) UserState {
	return s.usersState[id]
}

func (s *State) SetUser(id int, state UserState) {
	s.usersState[id] = state
}
