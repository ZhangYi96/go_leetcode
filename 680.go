package main
import "fmt"
func validPalindrome(s string) bool {
	return validFunc(s, 1)
}

func validFunc(s string, matchError int) bool{
	if matchError<0{
		return false
	}
	if len(s) <2{
		return true
	}
	fmt.Println(s)
	if s[0] == s[len(s)-1]{
		return validFunc(s[1:len(s)-1], matchError)
	}else{
		return validFunc(s[0:len(s)-1], matchError-1) || validFunc(s[1:len(s)], matchError-1)
	}
}
// func main(){
// 	fmt.Println(validPalindrome("abc"))
	
// }
