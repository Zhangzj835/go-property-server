package enums

type JsonResultCode int

const (
	JRCodeSucc JsonResultCode = iota
	JRCodeFailed
	JRCode302   = 302   //跳转至地址
	JRCode401   = 401   //未授权访问
	JRCode10001 = 10001 //token已过期
)

const (
	Deleted = iota - 1
	Disabled
	Enabled
)
