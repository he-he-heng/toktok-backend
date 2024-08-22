package provider

import (
	"testing"
)

func TestProvider(t *testing.T) {
	provider := Get()
	defer Get().TerminateContainer(MyContext)
	if provider == nil {
		t.Errorf("provider is nil")
	}

	if provider.Client == nil {
		t.Error("provider.client is nil")
	}

}
