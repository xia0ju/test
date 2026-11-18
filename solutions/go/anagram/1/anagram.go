package anagram
import  "strings"

import  "sort"

func Detect(subject string, candidates []string) []string {
    var sum =[]string{}
    
    for _,candidate :=range candidates{
        if strings.EqualFold(subject,candidate){
             continue
        }
	s1:=sortString(strings.ToLower(subject))
    s2:=sortString(strings.ToLower(candidate))
    if s1==s2{
        sum=append(sum,candidate)
        
    }
    }
    return sum
}
func sortString(str string)string{
    bytes:=[]rune (str)
    sort.Slice(bytes,func(i,j int)bool {
        return bytes[i]<bytes[j]
    })
    return string(bytes)
}
