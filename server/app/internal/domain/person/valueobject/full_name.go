package person_vo

import "fmt"

type FullName struct {
	value string
}

func NewFullName(name Name, surname Name, patronymic *Patronymic) FullName {
	if patronymic == nil {
		return FullName{fmt.Sprintf("%s %s", name.Value(), surname.Value())}
	}
	return FullName{fmt.Sprintf("%s %s %s", name.Value(), surname.Value(), patronymic.Value())}
}

func (fn FullName) Value() string {
	return fn.value
}

func (fn FullName) Equals(other FullName) bool {
	return fn.Value() == other.Value()
}
