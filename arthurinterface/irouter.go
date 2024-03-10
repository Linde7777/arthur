package arthurinterface

// IRouter 用于给框架使用者自定义业务方法
type IRouter interface {
	PreHandle(req IRequest) error
	Handle(req IRequest) error
	PostHandle(req IRequest) error
}
