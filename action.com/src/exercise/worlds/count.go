package worlds

import "strings"

func CountWorlds(text string) (count int) {

	count = len(strings.Fields(text))

	return
}
