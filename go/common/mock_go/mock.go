package mock_go

type DBMock interface {
	Get(key string) (int, error)
}

func GetFromDB(db DBMock, key string) int {
	if value, err := db.Get(key); err == nil {
		return value
	}

	return -1
}

//type Stream interface {
//	Get(key string) string
//	Add(key string) string
//}
//
//type MySteam struct {
//}
//
//func (m *MySteam) Get(key string) string {
//	return "123"
//}
//
//func (m MySteam) Add(key string) error {
//	return errors.New("add error")
//}
