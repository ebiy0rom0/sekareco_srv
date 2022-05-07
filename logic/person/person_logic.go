package logic

import "sekareco_srv/domain"

type PersonLogic struct {
	Repository PersonRepository
}

func (logic *PersonLogic) RegistPerson(p domain.Person) (personId int, err error) {
	personId, err = logic.Repository.Regist(p)
	return
}

func (logic *PersonLogic) SelectPerson(personId int) (person domain.Person, err error) {
	person, err = logic.Repository.SelectOne(personId)
	return
}
