package configuration

/*IConfiguration - интерфейс для парсинга параметров*/
type IConfiguration interface {
	Parse(pathname string) (*interface{}, error)
}
