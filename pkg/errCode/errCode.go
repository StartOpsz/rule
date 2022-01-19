package errCode

const (
	NormalCode = 0
	NormalMsg = "ok"
	
	NotFoundCode = 14040
	NotFoundMsg = "资源未发现"
	
	UserPermissionDenyCode = 14030
	UserPermissionDenyMsg = "用户权限拒绝"
	
	ParameterNotFoundCode = 14041
	ParameterNotFoundMsg = "参数缺失"
	
	DatabaseOpErrorCode = 17001
	DatabaseOpErrorMsg = "内部系统异常，请核对code。"
	
	BizOpErrorCode = 17002
	BizOpErrorMsg = "业务代码异常"
	
	
)
