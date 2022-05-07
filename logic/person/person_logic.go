package logic

import "sekareco_srv/domain"

type PersonLogic struct {
	Repo PersonRepo
}

func (logic *PersonLogic) RegistPerson(p domain.Person) (personId int, err error) {
	personId, err = logic.Repo.Regist(p)
	return
}

func (logic *PersonLogic) SelectPerson(personId int) (person domain.Person, err error) {
	person, err = logic.Repo.SelectOne(personId)
	return
}
