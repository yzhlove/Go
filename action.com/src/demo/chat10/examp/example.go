package examp

type accountNumber int

//New 返回alertCounter 类型的值
func New(value int) accountNumber {
	return accountNumber(value)
}
