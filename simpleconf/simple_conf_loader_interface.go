package simpleconf

type Config interface {
	ValueOf(key string) (value string, err error)
	ValuesFor(key string) (values []string)

	ExtractValueOf(key string, from []string) (value string, err error)
}
