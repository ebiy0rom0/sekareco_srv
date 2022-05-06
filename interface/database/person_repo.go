package database

import "sekareco_srv/domain"

type PersonRepo struct {
	SqlHandler
}

func (repo *PersonRepo) Regist(p *domain.Person) (id int, err error) {
	// TODO: wip
	result, err := repo.Execute("", p)
	if err != nil {
		return
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return
	}

	id = int(id64)
	return
}

func (repo *PersonRepo) SelectOne(personId int)
