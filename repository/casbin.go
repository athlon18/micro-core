package repository

import (
	"github.com/athlon18/micro-core/casbin/structs"
	"github.com/casbin/casbin/v2"
)

type CasbinRepository interface {
	Casbin() *casbin.SyncedEnforcer
	GetPolicyPathByAuthorityId(string) []structs.CasbinInfo
	ClearCasbin(int, ...string) bool
	UpdateCasbinApi(string, string, string, string) error
	UpdateCasbin(string, []structs.CasbinInfo) error
}
