package gormodel

type Person struct {
	Id      int64
	Name    string
	Age     int
	Address string
}

func CreaTeable(per Person) {
	orm.Create(&per)
}

func Insertperson(per Person) {
	orm.Save(per)
}

func Query(pers []*Person) {
	orm.Find(&pers)
}
