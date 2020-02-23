package registry

import (
	"fmt"

	"github.com/ic2hrmk/azck/app"
)

type RegistryContainer map[string]FactoryMethod

func NewRegistryContainer() RegistryContainer {
	return make(RegistryContainer)
}

type FactoryMethod func() (app.MicroService, error)

func (c *RegistryContainer) Add(serviceName string, fabric FactoryMethod) {
	(*c)[serviceName] = fabric
}

func (c *RegistryContainer) Get(serviceName string) (FactoryMethod, error) {
	entry, ok := (*c)[serviceName]
	if !ok {
		return nil, fmt.Errorf("service [%s] isn't registered", serviceName)
	}

	return entry, nil
}
