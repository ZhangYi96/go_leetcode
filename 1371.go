package main
import "fmt"

func findTheLongestSubstring(s string) int {

	ctn := make(map[byte]int)
	ctn['a']=0
	ctn['e']=0
	ctn['i']=0
	ctn['o']=0
	ctn['u']=0

	return find(s, ctn, 0,0)

}

func valid(ctn map[byte]int) bool{
	for _,v := range ctn{
		if v % 2 != 0{
			return false
		}
	}
	return true
}

func find(s string, ctn map[byte]int, left, right int) int {
	if right >= len(s){
		if valid(ctn){
			return right -left
		}
		return 0
		
	}

	maxResult := 0

	strLen := len(s)
	for ;right < strLen; right ++{
		fmt.Println(ctn)
		if _, ok :=ctn[s[right]]; ok{
			ctn[s[right]] = ctn[s[right]] + 1
		}
		flag := valid(ctn)

		if flag == true || left ==right{	
			maxResult = right - left + 1
		}else {
			r := find(s, ctn, left, right+1)

			if _,ok := ctn[s[left]]; ok{
				ctn[s[left]] = ctn[s[left]] - 1
			}
			if _, ok :=ctn[s[right]]; ok{

				ctn[s[right]] = ctn[s[right]] - 1
			}
			l := find(s, ctn, left+1, right)

			if r>l{
				
				return r
			}
			return l
		}
		
	}

	return maxResult
}

func main()  {
	fmt.Println(findTheLongestSubstring("eleetminicoworoep"))
}