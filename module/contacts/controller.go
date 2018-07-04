package contacts

import (
	"github.com/pkg/errors"
	"github.com/golang/glog"
)

type Controller struct{
	ContactService *Service
}

func New(service *Service) (*Controller,error){

	if service == nil {
		glog.Error("Contact Service is nil, cannot instantiate Contact Controller")
		return nil, errors.New("Contact Service is nil, cannot instantiate Contact Controller")
	}

	return &Controller{service}, nil
}

