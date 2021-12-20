package presenter

type ResponseData (map[string]interface{})

type ParkRestPresenterContext interface {
	JSON(code int, i interface{}) error
	Bind(i interface{}) error
	Param(name string) string
}
