package main

func EffectiveMem(n int) uint64 {

	f := make([]uint64, 10000) // make huge allocation

	f[0] = 0
	f[1] = 1

	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}

	return f[n]
}
