package fake_public_component_1

import "fmt"

const (
	COMPONENT = "FakePublicComponent_1"
	VERSION   = "0.0.1"
)

type FakePublicComponent1 interface {
	DoSomething() string
}

type FakePublicComponentStruct struct {
}

func (f FakePublicComponentStruct) DoSomething() string {
	return fmt.Sprintf("Component: %s  Version: %s", COMPONENT, VERSION)
}
