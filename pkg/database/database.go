package database

type DataBaseProvider string

const (
	MYSQL_DB      DataBaseProvider = "MYSQL"
	MEMORY_DB     DataBaseProvider = "MEMORY"
	UNKNOWN_CACHE DataBaseProvider = "UNKNOWN"
)

type IDatabase interface {
	GetProvider() string
}
