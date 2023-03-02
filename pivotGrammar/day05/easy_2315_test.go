package day05

/**
给你一个字符串s，每两个连续竖线'|'为 一对。换言之，第一个和第二个'|'为一对，第三个和第四个'|'为一对，以此类推。

请你返回 不在 竖线对之间，s中'*'的数目。

注意，每个竖线'|'都会 恰好属于一个对。

*/

func countAsterisks(s string) int {
	var flag int
	var num int

	for _, str := range s {
		if string(str) == "|" {
			flag++
		}
		if flag%2 == 0 {
			if string(str) == "*" {
				num++
			}
		}

	}
	return num
}
