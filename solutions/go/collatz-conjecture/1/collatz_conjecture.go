package collatzconjecture
import "errors"

func CollatzConjecture(n int) (int, error) {
	var t=0
       if n<=0{
        return 0,errors.New("error")
    }
	for n!=1{
        if n%2==0{
            n=n/2
            t++
        }else {
            n=n*3+1
            t++
        }
        
    }
 
    return t,nil
}
