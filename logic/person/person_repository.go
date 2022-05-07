package logic

import "sekareco_srv/domain"

type PersonRepository interface {
	Regist(domain.Person) (int, error)
	SelectOne(int) (domain.Person, error)
}
