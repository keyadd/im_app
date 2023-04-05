package wsmanage

type IRouter interface {
	PreHandle(request Request) (err error) //在处理conn业务之前的钩子方法
	Handle(request Request)                //处理conn业务的方法
	PostHandle(request Request)

	//处理conn业务之后的钩子方法
}

// 实现router时，先嵌入这个基类，然后根据需要对这个基类的方法进行重写
type BaseRouter struct {
	// This mutex protects Keys map.

}

// 这里之所以BaseRouter的方法都为空，
// 是因为有的Router不希望有PreHandle或PostHandle
// 所以Router全部继承BaseRouter的好处是，不需要实现PreHandle和PostHandle也可以实例化
func (br *BaseRouter) PreHandle(req Request) (err error) { return }
func (br *BaseRouter) Handle(req Request)                {}
func (br *BaseRouter) PostHandle(req Request)            {}
