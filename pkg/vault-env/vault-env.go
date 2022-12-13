package vault_env

import (
	"fmt"
	vault "github.com/hashicorp/vault/api"
)

type VaultEnv struct{}

func (v *VaultEnv) Eject(secret *vault.KVSecret) error {
	for k := range secret.Data {
		fmt.Println("unset", k)
	}
	return nil
}

func (v *VaultEnv) Inject(secret *vault.KVSecret) error {
	for k, v := range secret.Data {
		fmt.Println("export", k+"="+v.(string))
	}
	return nil
}
