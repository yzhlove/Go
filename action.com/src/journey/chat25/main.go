package main

import (
	"fmt"
	"sort"
)

//结构体排序

/*
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
*/

type HeroKind int

const (
	None HeroKind = iota
	Tank
	Assassin
	Mage
)

type Hero struct {
	Element string
	kind    HeroKind
}

type Heros []*Hero

func (h Heros) Len() int {
	return len(h)
}

func (h Heros) Less(i, j int) bool {
	if h[i].kind != h[j].kind {
		return h[i].kind > h[j].kind
	}
	return h[i].Element > h[j].Element
}

func (h Heros) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func main() {
	heros := Heros{
		&Hero{"吕布", Tank},
		&Hero{"貂蝉", Assassin},
		&Hero{"董卓", Tank},
		&Hero{"曹操", Mage},
		&Hero{"刘备", Mage},
		&Hero{"关羽", Tank},
		&Hero{"张飞", Tank},
		&Hero{"赵云", Assassin},
		&Hero{"诸葛亮", Mage},
		&Hero{"司马懿", Mage},
	}

	sort.Sort(heros)

	tmp := Heros{}
	for _, v := range heros {
		tmp = append(tmp, v)
	}

	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].kind != tmp[j].kind {
			return tmp[i].kind < tmp[j].kind
		}
		return tmp[i].Element < tmp[j].Element
	})

	for _, v := range tmp {
		fmt.Printf("%v \n", v)
	}

	fmt.Println()

	for _, v := range heros {
		fmt.Printf("%v \n", v)
	}
}
