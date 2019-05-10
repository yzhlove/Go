package st

import "fmt"

type IdleState struct {
	StateInfo
}

func (i *IdleState) OnBegin() {
	fmt.Println("IdleState Begin")
}

func (i *IdleState) OnEnd() {
	fmt.Println("IdleState End")
}

type MoveState struct {
	StateInfo
}

func (m *MoveState) OnBegin() {
	fmt.Println("MoveState Begin")
}

func (m *MoveState) EnableSameTransit() bool {
	return true
}

type JumpState struct {
	StateInfo
}

func (jump *JumpState) OnBegin() {
	fmt.Println("JumpState Begin")
}

func (jump *JumpState) CanCurrTransitTo(name string) bool {
	return name != "MoveState"
}

func TransitAndReport(sm *StateManager, target string) {
	if err := sm.Transit(target); err != nil {
		fmt.Printf("FAILED! %s ---> %s,%s\n\n", sm.CurrState().Name(), target, err.Error())
	}
}
