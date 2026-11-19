package phonenumber
import  "errors"
//import  "strings"

import  "fmt"

func Number(phoneNumber string) (string, error) {
    	var str string
    for _,v :=range phoneNumber{
        if v>='0'&&v<='9'{
         str+=string(v)
        }
        
      
    }
      if len(str)==11&&str[0]=='1'{
            str=str[1:]
        }
    if len(str)!=10{
        return "",errors.New("error")
    }
    if str[0]=='0'||str[0]=='1'{
        return "",errors.New("error")
    }
       if str[3]=='0'||str[3]=='1'{
        return "",errors.New("error")
    }
      return str,nil
}

func AreaCode(phoneNumber string) (string, error) {
	newNumber,error:=Number(phoneNumber)
    if error!=nil{
        return "",error
    }
    return newNumber[0:3],nil
}

func Format(phoneNumber string) (string, error) {
	areacode,error1:=AreaCode(phoneNumber)
    if error1!=nil{
        return "",error1
    }
    num,error2:=Number(phoneNumber)
    if error2!=nil{
        return "",error2
    }
    return fmt.Sprintf("(%v) %v-%v",areacode,num[3:6],num[6:]),nil	
}
