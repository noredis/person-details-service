package person_vo

import "strings"

type Patronymic struct {
	value string
}

func NewPatronymic(patronymic string) *Patronymic {
	patronymic = strings.Trim(patronymic, " ")
	if patronymic == "" {
		return nil
	}

	return &Patronymic{patronymic}
}

func (p Patronymic) Value() string {
	return p.value
}

func (p Patronymic) Equals(other Patronymic) bool {
	return p.Value() == other.Value()
}
