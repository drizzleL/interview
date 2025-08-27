package main

func numOfSubsequences(s string) int64 {
	var tCnt int
	for _, c := range s {
		if c == 'T' {
			tCnt += 1
		}
	}
	lcnt1, tcnt1 := 1, tCnt
	lcnt2, tcnt2 := 0, tCnt+1
	var cnt1, cnt2 int
	for _, c := range s {
		if c == 'C' {
			cnt1 += lcnt1 * tcnt1
			cnt2 += lcnt2 * tcnt2
			continue
		}
		if c == 'L' {
			lcnt1 += 1
			lcnt2 += 1
		} else if c == 'T' {
			tcnt1 -= 1
			tcnt2 -= 1
		}
	}
	ret := max(cnt1, cnt2)
	cnt, lcnt, rcnt, maxInc := 0, 0, tCnt, 0
	for _, c := range s {
		maxInc = max(maxInc, lcnt*rcnt)
		if c == 'C' {
			cnt += lcnt * rcnt
			continue
		}
		if c == 'L' {
			lcnt += 1
		} else if c == 'T' {
			rcnt -= 1
		}
	}
	ret = max(ret, cnt+maxInc)
	return int64(ret)
}
