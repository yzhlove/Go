package main

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type Info struct {
	Cmd  string
	Msg  string
	Args []string
}

var (
	MasterID string
	SaveID   string
)

func main() {

	cmds := []*Info{
		&Info{Cmd: "docker", Args: []string{"run", "-d", "-p", "6380:6379", "--name", "saved", "redis"}},
		&Info{Cmd: "docker", Args: []string{"run", "-d", "-p", "6379:6379", "--name", "masted", "--link", "saved:saved", "-v", "/Develop/Docker/Volumes/Redis:/data", "redis"}},
	}
	if err := startDocker(cmds); err != nil {
		panic(err)
	}

	fmt.Println("cmd running successful .")

	fmt.Printf("%q %q \n", MasterID, SaveID)

	time.Sleep(time.Second * 15)

	infos := []*Info{
		&Info{Cmd: "docker", Args: []string{"stop", MasterID}},
		&Info{Cmd: "docker", Args: []string{"stop", SaveID}},
		&Info{Cmd: "docker", Args: []string{"rm", "-f", "saved"}},
		&Info{Cmd: "docker", Args: []string{"rm", "-f", "masted"}},
	}

	if err := stopDocker(infos); err != nil {
		fmt.Println("stop docker err: ", err)
	}

	fmt.Println("Done .")
}

func startDocker(cmds []*Info) error {

	values := []*string{&MasterID, &SaveID}
	errMsg := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())

	go func(infos []*Info, values []*string) {
		defer cancel()
		for i := 0; i < len(infos); i++ {
			if result, err := run(infos[i]); err != nil {
				errMsg <- err
				return
			} else {
				*values[i] = strings.TrimRight(string(result), "\n")
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

func stopDocker(cmds []*Info) error {
	errMsg := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())
	go func(infos []*Info) {
		defer cancel()
		for i := 0; i < len(infos); i++ {
			if _, err := run(infos[i]); err != nil {
				errMsg <- err
			}
		}
	}(cmds)
	for {
		select {
		case <-time.NewTimer(10 * time.Second).C:
			return fmt.Errorf("stop docker timeout  ...")
		case <-ctx.Done():
			return nil
		case err := <-errMsg:
			fmt.Println("[ERROR] stop docker err:", err)
		}
	}
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
