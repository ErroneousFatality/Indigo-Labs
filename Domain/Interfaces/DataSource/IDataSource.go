package DataSource

type IDataSource interface {
	ReadStream() <-chan StreamResult
}
