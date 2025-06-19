package main

func pushDominoes(dominoes string) string {
	b := []byte(dominoes)
	type Op struct {
		idx int
		op  byte
	}
	lastOp := Op{idx: -1}
	for i := 0; i < len(dominoes); i++ {
		if dominoes[i] == '.' {
			continue
		}
		op := Op{idx: i, op: dominoes[i]}
		if lastOp.idx == -1 && op.op == 'L' {
			for j := 0; j < op.idx; j++ {
				b[j] = 'L'
			}
		}
		if lastOp.idx != -1 {
			switch {
			case op.op == lastOp.op:
				for j := lastOp.idx; j < op.idx; j++ {
					b[j] = op.op
				}
			case lastOp.op == 'R' && op.op == 'L':
				for j, k := lastOp.idx, op.idx; j < k; j, k = j+1, k-1 {
					b[j] = 'R'
					b[k] = 'L'
				}
			}
		}
		lastOp = op
	}
	if lastOp.op == 'R' {
		for j := lastOp.idx; j < len(b); j++ {
			b[j] = 'R'
		}
	}
	return string(b)
}
