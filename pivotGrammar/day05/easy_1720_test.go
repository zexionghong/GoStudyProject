package day05

/**
经编码后变为长度为 n - 1 的另一个整数数组 encoded ，其中 encoded[i] = arr[i] XOR arr[i + 1] 。例如，arr = [1,0,2,1] 经编码后得到 encoded = [1,2,3] 。

给你编码后的数组 encoded 和原数组 arr 的第一个元素 first（arr[0]）。

请解码返回原数组 arr 。可以证明答案存在并且是唯一的。

a^a=0
a^0=a
a^b^c =
*/

func decode(encoded []int, first int) []int {
	length := len(encoded)

	var tmp = make([]int, length+1)
	tmp[0] = first
	for i := 1; i < length+1; i++ {
		tmp[i] = tmp[i-1] ^ encoded[i-1]

	}
	return tmp
}
