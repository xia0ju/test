package luhn
import "strings"

func Valid(id string) bool {
    id=strings.ReplaceAll(id," ","")
	    if len(id)<=1{
            return false
        }    
    var s=0
    for i:=0;i<len(id);i++{
       var num=id[i]
        if num<'0'||num>'9'{
            return false
        }
        var n =int (num-'0')
        if (len(id)%2==0&&i%2==0 )||(len(id)%2!=0&&i%2!=0){
            n=n*2
            if n>9{
                n=n-9
            }
        }
        s+=n
   
    
}
    if s%10==0{
        return true
    }
    return false
}
