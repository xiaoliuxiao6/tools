package tools

import "strconv"

// uint64 解码为 0x 开头的16进制字符串
// 参考：https://github.com/ethereum/go-ethereum/blob/master/common/hexutil/hexutil.go
func EncodeUint64(i uint64) string {
	enc := make([]byte, 2, 10)
	copy(enc, "0x")
	return string(strconv.AppendUint(enc, i, 16))
}
