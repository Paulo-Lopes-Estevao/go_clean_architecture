package presenter

type ParkRestPresenterContext interface {
	JSON(code int, i interface{}) error
	Bind(i interface{}) error
	Param(name string) string
	ResponseData(map[string]interface{}) interface{}
}
