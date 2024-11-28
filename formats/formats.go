package formats

type Format interface {
	String(args ...interface{}) string
	Byte(args ...interface{}) []byte
}
