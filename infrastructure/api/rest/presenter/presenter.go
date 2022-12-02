package presenter

type ResponseData (map[string]interface{})

type PresenterContext interface {
	JSON(code int, i interface{}) error
	Bind(i interface{}) error
	Param(name string) string
}
