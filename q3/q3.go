package q3


func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}
	degrauAnterior := 1
	degrauAnterior2 := 2
	degrauAtual := 0

	for i := 3; i <= n; i++ {
		degrauAtual = degrauAnterior + degrauAnterior2
		degrauAnterior = degrauAnterior2
		degrauAnterior2 = degrauAtual
	}
	return degrauAtual
}

