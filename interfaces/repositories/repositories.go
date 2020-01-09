package repositories

// DbHandler is a high level interface that allows repository
// interrogation by hiding all the low level aspects.
type DbHandler interface {
	Execute(statement string) error
	Query(statement string) (Row, error)
}

// Row is a high level interface that allows data manipulation.
type Row interface {
	Scan(dest ...interface{}) error
	Next() bool
}

// DbRepo represents a general repository.
// The handlers map lets every repository use any other repository
// without giving up dependency injection.
type DbRepo struct {
	dbHandlers map[string]DbHandler
	dbHandler  DbHandler
}

// All is a list of all the repository names
var All = [...]string{"DbUserRepo", "DbAnimalRepo", "DbAdoptionRepo", "DbAdopterRepo"}
