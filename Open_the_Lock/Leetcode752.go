package main

func openLock(deadends []string, target string) int {

	type lock struct {
		slots string
		step  int
	}
	set := make(map[string]struct{})
	for _, d := range deadends {
		set[d] = struct{}{}
	}
	if _, ok := set["0000"]; ok {
		return -1
	}
	seen := make(map[string]struct{})
	seen["0000"] = struct{}{}
	queue := make([]lock, 0)
	queue = append(queue, lock{"0000", 0})
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		slots, step := front.slots, front.step
		if slots == target {
			return step
		}
		for i := 0; i < len(slots); i++ {
			c := slots[i]
			nc := c + 1
			if nc > '9' {
				nc = '0'
			}
			ns := slots[:i] + string(nc) + slots[i+1:]
			if _, ok := seen[ns]; !ok {
				if _, ok2 := set[ns]; !ok2 {
					queue = append(queue, lock{ns, step + 1})
					seen[ns] = struct{}{}
				}
			}
			nc = c - 1
			if nc < '0' {
				nc = '9'
			}
			ns = slots[:i] + string(nc) + slots[i+1:]
			if _, ok := seen[ns]; !ok {
				if _, ok2 := set[ns]; !ok2 {
					queue = append(queue, lock{ns, step + 1})
					seen[ns] = struct{}{}
				}
			}

		}
	}
	return -1
}
