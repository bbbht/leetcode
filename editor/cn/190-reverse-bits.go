package main
// 190 reverse-bits 2022-11-29 11:02:31

//leetcode submit region begin(Prohibit modification and deletion)
func reverseBits(num uint32) uint32 {
	revBits := uint32(0)
	for i := 0; i < 32; i++ {
		revBits <<= 1
		revBits |= num >> i & 1
	}

	return revBits
}
//leetcode submit region end(Prohibit modification and deletion)


func main() {

}