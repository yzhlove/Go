package main

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"time"
)

type Info struct {
	Cmd  string
	Msg  string
	Args []string
}

var (
	MasterID string
	SavelID  string
)

func main() {

	cmds := []*Info{
		&Info{Cmd: "docker", Args: []string{"run", "-d", "-p", "6380:6379", "--name", "saved", "redis"}},
	}
	if err := startDocker(cmds); err != nil {
		panic(err)
	}

	fmt.Println("cmd running successful .")

	fmt.Println(MasterID, " ", SavelID)

}

func startDocker(cmds []*Info) error {

	values := []*string{&MasterID, &SavelID}
	errMsg := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())

	go func(infos []*Info, values []*string) {
		defer cancel()
		for i := 0; i < len(infos); i++ {
			if result, err := run(infos[i]); err != nil {
				errMsg <- err
				return
			} else {
				*values[i] = string(result)
			}
		}
	}(cmds, values)

	select {
	case <-time.NewTimer(10 * time.Second).C:
		return fmt.Errorf("cmd running timeout")
	case <-ctx.Done():
		return nil
	case err := <-errMsg:
		if err != nil {
			return fmt.Errorf("cmd running err:%v ", err)
		}
	}
	return nil
}

func stopDocker() error {

	return nil
}

func run(info *Info) ([]byte, error) {
	var (
		out []byte
		err error
	)
	if out, err = exec.Command(info.Cmd, info.Args...).Output(); err != nil {
		return []byte{}, err
	}
	if info.Msg != "" && bytes.Index(out, []byte(info.Msg)) == -1 {
		return []byte{}, fmt.Errorf("Err:%v \n", info.Msg)
	}
	return out, nil
}
