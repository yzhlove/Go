package main

import (
	"fmt"
	"teaching/chat05/st"
)

func main() {

	sm := st.NewStateManager()
	sm.OnChange = func(from, on st.State) {
		fmt.Printf("%s ---> %s\n\n", st.GetStateName(from), st.GetStateName(on))
	}

	sm.Add(new(st.IdleState))
	sm.Add(new(st.MoveState))
	sm.Add(new(st.JumpState))

	st.TransitAndReport(sm, "IdleState")
	st.TransitAndReport(sm, "MoveState")
	st.TransitAndReport(sm, "MoveState")
	st.TransitAndReport(sm, "JumpState")
	st.TransitAndReport(sm, "JumpState")
	st.TransitAndReport(sm, "IdleState")

}
