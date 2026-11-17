package isbnverifier
import  "strings"

func IsValidISBN(isbn string) bool {
  var isbn1=strings.ReplaceAll(isbn,"-","")
    var s int
 if len(isbn1)!=10{
     return false
 }
	for i:=0;i<len(isbn1);i++{
        str:=isbn1[i]
    	var num int
     
            if (str=='x'||str=='X')&&i==9{
                num=10
            }else if str>='0'&&str<='9'{    
                num= int(str-'0')
               }else{
                	return false
                               }
         
            
            s+=num*(10-i)
          
     
        
    }
    if s%11==0{
        return true
    }
	return false
}
