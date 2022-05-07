package logic

import "sekareco_srv/domain"

type PersonRepo interface {
	Regist(domain.Person) (int, error)
	SelectOne(int) (domain.Person, error)
}
