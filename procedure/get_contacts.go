package procedure

const (
	GetContactsPath = "getContacts"

	SexMale   = 1
	SexFemale = 2
)

type GetContactsInput struct {
	IdGroup *int64  `json:"idGroup"`
	Phone   *string `json:"phone"`
}

type GetContactsOutput struct {
	Code     int64  `json:"code"`
	Descr    string `json:"descr"`
	Contacts string `json:"contacts"`
}

type ContactList struct {
	Data   []Contact `json:"data"`
	AllCol int64     `json:"allCol"`
}

type Contact struct {
	Id    string `json:"id"`
	Phone string `json:"phone"`
	Fio   string `json:"fio"`
	BDay  string `json:"bday"`
	Sex   int64  `json:"sex"`
}
