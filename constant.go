package manupay_client

const (
	//地址
	UNIFIEDORDER_PATH = "/pay/order/place"
)

const (
	MchNo         = "MC1702569748987908096"                                                                                  //分配的商户号
	PrivateSecret = "bDwJP1Ug17mKwq9xqlULT98IQ90QKNgqu9jx2jeF0LCntbwLTlupNPnKtLwAXnY6ZwIgfNGwQC1qDPg2qjhy8nySKZmCylcKjt1YTI" //分配的私钥
)

// 支付方式
const (
	SailCashier    string = "SAIL_CASHIER"     //测试收单
	SailCashierBrl string = "SAIL_CASHIER_BRL" //巴西收单
	SailCashierInr string = "SAIL_CASHIER_INR" //印度收单
)

// 货币编码
const (
	Inr string = "inr" //印度卢比
	Brl string = "brl" //巴西
)

// 分账模式
const (
	NoDivision     int = 0 //- 该笔订单不允许分账[默认],
	AutoDivision   int = 1 //- 支付成功按配置自动完成分账,
	ManualDivision int = 2 //- 商户手动分账(解冻商户金额)
)

// 订单状态
const (
	StateCreate     int = 0 //-订单生成
	StatePaying     int = 1 //-支付中
	StatePaySucceed int = 2 //-支付成功
	StatePayFailed  int = 3 //-支付失败
	StateCanceled   int = 4 //-已撤销
	StateChargeBack int = 5 //-已退款
	StateClosed     int = 6 //-订单关闭
)

// 货币编码
const (
	PayUrl     string = "payUrl"     //-跳转链接的方式
	Form       string = "form"       //-表单方式
	CodeUrl    string = "codeUrl"    //-二维码地址
	CodeImgUrl string = "codeImgUrl" //-二维码图片地址
	None       string = "none"       //-空支付参数
)
