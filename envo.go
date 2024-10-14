package envo

import "fmt"

type Env string

const (
	Local Env = "local"
	Dev   Env = "dev"
	Prod  Env = "prod"
)

func (e Env) Valid() bool {
	return e.IsLocal() || e.IsDev() || e.IsProd()
}

func (e Env) IsProd() bool {
	return e == Prod
}

func (e Env) IsDev() bool {
	return e == Dev
}

func (e Env) IsLocal() bool {
	return e == Local
}

func (e Env) NotProd() bool {
	return e != Prod
}

func (e Env) NotDev() bool {
	return e != Dev
}

func (e Env) NotLocal() bool {
	return e != Local
}

func (e Env) String() string {
	return string(e)
}

func New(value string) (*Env, error) {
	env := Env(value)

	if !env.Valid() {
		return nil, fmt.Errorf("некорректное название: %s", value)
	}

	return &env, nil
}

func MustNew(value string) *Env {
	env, err := New(value)

	if err != nil {
		panic(err)
	}

	return env
}
