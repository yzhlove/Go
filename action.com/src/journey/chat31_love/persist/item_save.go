package persist

import "log"

func ItemSave() chan interface{} {
	out := make(chan interface{}, 128)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item:%d %v\n", itemCount, item)
			itemCount++
		}
	}()
	return out
}
