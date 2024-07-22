package main

//text，哈哈哈哈哈哈哈哈哈哈哈哈哈
import (
	"fmt"
	"math"
)

const (
	sml float64 = 1e-7
	N   int     = 110
)

var mat [N][N]float64

func swaq(i int, j int, n int) {
	for k := 1; k <= n+1; k++ {
		var tmp float64 = mat[j][k]
		mat[j][k] = mat[i][k]
		mat[i][k] = tmp
	}
}
func gauss(n int) {
	for i := 1; i <= n; i++ {
		var max int = i
		for j := 1; j <= n; j++ {
			if j < i && mat[j][j] >= sml {
				continue
			}
			if math.Abs(mat[j][i]) > math.Abs(mat[max][i]) {
				max = j
			}
		}
		swaq(i, max, n)
		if math.Abs(mat[i][i]) >= sml {
			//把首元变成1
			for j := n + 1; j >= i; j-- {
				mat[i][j] /= mat[i][i]
			}
			//把除了首元所在行之外的列变成0，j当行，k当列
			for j := n; j >= 1; j-- {
				if j != i {
					for k := n + 1; k >= 1; k-- {
						mat[j][k] -= mat[i][k] * mat[j][i]
					}
				}
			}
		}
	}
}
func PrintfArray(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n+1; j++ {
			fmt.Printf("%.2f ", mat[i][j])
		}
		fmt.Println()
	}
}
func init_arr(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n+1; j++ {
			mat[i][j] = 0
		}
	}
}
func main() {
	//读取数据
	var fk int = 0
	var max int = 1
	var maxmax int = 1
	var n int
	var flag int = 1
	fmt.Scan(&n)
	arr := make([][]float64, n+2)
	for i := range arr {
		arr[i] = make([]float64, n+3)
	}
	for i := 1; i <= n+1; i++ {
		var m float64
		fmt.Scan(&m)
		arr[i][1] = m
		for j := 2; j <= int(m)+1; j++ {
			fmt.Scan(&arr[i][j])
		}
		fmt.Scan(&arr[i][int(m)+2])
	}
	//进行n + 1次高斯消元
	for i := 1; i <= n+1; i++ {
		init_arr(n)
		//对mat数组进行赋值
		t, j := 1, 1
		for j, t = 1, 1; j <= n+1; j++ {
			if i == j {
				continue
			} else {
				m := arr[j][1]
				k := 2
				for k = 2; k <= int(m)+1; k++ {
					mat[t][int(arr[j][k])] = 1
				}
				mat[t][n+1] = arr[j][k]
				t += 1

			}
		}
		//	PrintfArray(n)
		//	fmt.Println()
		gauss(n)
		//	PrintfArray(n)
		//	fmt.Println()
		for i := 1; i <= n; i++ {
			if math.Abs(mat[i][i]) <= sml {
				flag = 2
				break
			} else if mat[i][n+1] < sml {
				flag = 2
				break
			} else {
				flag = 1
			}
		}
		//fmt.Println(flag)

		if flag == 1 {
			max = 1
			for i := 2; i <= n; i++ {
				if mat[i][n+1] != float64(int(mat[i][n+1])) {
					flag = 2
					break
				}
				if mat[i][n+1] == mat[max][n+1] {
					flag = 2

					break
				}
				if mat[i][n+1] > mat[max][n+1] {
					max = i
				}
				//fmt.Println(mat[i][n + 1])
			}

		}
		//	fmt.Println(flag)
		if flag == 1 {
			if fk == 1 {
				fk = 2

				flag = 2
				break
			}
			maxmax = max
			fk = 1

		}
	}
	if fk == 0 || fk == 2 {
		fmt.Println("illegal")
	} else if fk == 1 {
		fmt.Println(maxmax)
	}

}
