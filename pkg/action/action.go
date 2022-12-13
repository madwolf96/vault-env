package action

import (
	"fmt"
	vault "github.com/hashicorp/vault/api"
)

type Action struct{}

func (v *Action) Eject(secret *vault.KVSecret) error {
	for k := range secret.Data {
		fmt.Println("unset", k)
	}
	return nil
}

func (v *Action) Inject(secret *vault.KVSecret) error {
	for k, v := range secret.Data {
		fmt.Println("export", k+"="+v.(string))
	}
	return nil
}
