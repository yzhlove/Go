package st

import "errors"

var ErrStateNotFound = errors.New("state not found")
var ErrForbidSameStateTransit = errors.New("forbid same state transit")
var ErrCannotTransitToState = errors.New("cannot transit to state")

type StateManager struct {
	stateByName map[string]State
	OnChange    func(from, on State)
	curr        State
}

func (sm *StateManager) Add(s State) {
	name := GetStateName(s)
	s.(interface{ setName(string) }).setName(name)
	if sm.Get(name) != nil {
		panic("duplicate state:" + name)
	}
	sm.stateByName[name] = s
}

func (sm *StateManager) Get(name string) State {
	if v, ok := sm.stateByName[name]; ok {
		return v
	}
	return nil
}

func (sm *StateManager) CurrState() State {
	return sm.curr
}

func (sm *StateManager) CanCurrTransitTo(name string) bool {
	if sm.curr == nil {
		return true
	}
	if sm.curr.Name() == name && !sm.curr.EnableSameTransit() {
		return false
	}
	return sm.curr.CanTransitTo(name)
}

func (sm *StateManager) Transit(name string) error {
	next := sm.Get(name)
	if next == nil {
		return ErrStateNotFound
	}
	pre := sm.curr
	if sm.curr != nil {
		if sm.curr.Name() == name && !sm.curr.EnableSameTransit() {
			return ErrForbidSameStateTransit
		}
		if !sm.curr.CanTransitTo(name) {
			return ErrCannotTransitToState
		}
		sm.curr.OnEnd()
	}
	sm.curr = next
	sm.curr.OnBegin()
	if sm.OnChange != nil {
		sm.OnChange(pre, sm.curr)
	}
	return nil
}

func NewStateManager() *StateManager {
	return &StateManager{
		stateByName: make(map[string]State),
	}
}
