package differenceofsquares

func SquareOfSum(n int) int {
	var s=0
    var ss=0

    for  i:=1;i<=n;i++{
        s+=i
    }
    ss=s*s
    return ss
}

func SumOfSquares(n int) int {
	var s=0
 
    for  i:=1;i<=n;i++{
        s+=i*i
    }
    return s
}

func Difference(n int) int {
	return SquareOfSum(n)-SumOfSquares(n)
}
