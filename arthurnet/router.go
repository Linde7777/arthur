package arthurnet

import (
	"errors"
	"main/arthurinterface"
)

type BaseRouter struct {
}

func (br *BaseRouter) PreHandle(req arthurinterface.IRequest) error {
	return errors.New("implement me")
}

func (br *BaseRouter) Handle(req arthurinterface.IRequest) error {
	return errors.New("implement me")
}

func (br *BaseRouter) PostHandle(req arthurinterface.IRequest) error {
	return errors.New("implement me")
}
