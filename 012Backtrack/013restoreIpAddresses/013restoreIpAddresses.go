package main

// 93. 复原 IP 地址

//func main() {
//	s := "25525511135"
//	fmt.Println(restoreIpAddresses(s))
//}
//
//var result []string

//func restoreIpAddresses(s string) []string {
//	result = make([]string, 0)
//	backtrack(s, 0, 0, []int{})
//	return result
//}
//
//func backtrack(s string, k, step int, path []int) {
//	if k == len(s) && step == 4 {
//		sb := strings.Builder{}
//		for i := 0; i < 3; i++ {
//			sb.Write([]byte(fmt.Sprintf("%d", path[i])))
//			sb.WriteByte(byte('.'))
//		}
//		sb.Write([]byte(fmt.Sprintf("%d", path[3])))
//		result = append(result, sb.String())
//	}
//	if step > 4 {
//		return
//	}
//	if k == len(s) {
//		return
//	}
//	val := 0
//	// 1位数
//
//}
