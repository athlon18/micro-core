package casbin

import (
	"errors"
	"github.com/athlon18/micro-core/casbin/structs"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
	"sync"
)

type CasbinService struct {
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
	db             *gorm.DB
	modelPath      string
}

func NewCasbinRepository(db *gorm.DB, modelPath string) *CasbinService {
	return &CasbinService{db: db}
}

func (c *CasbinService) Casbin() *casbin.SyncedEnforcer {
	c.once.Do(func() {
		a, err := gormadapter.NewAdapterByDB(c.db)
		if err != nil {
			panic(err.Error())
		}
		c.syncedEnforcer, err = casbin.NewSyncedEnforcer(c.modelPath, a)
		if err != nil {
			panic(err.Error())
		}

	})
	if err := c.syncedEnforcer.LoadPolicy(); err != nil {
		panic(err.Error())
	}
	return c.syncedEnforcer
}

// GetPolicyPathByAuthorityId 获取权限列表
func (c *CasbinService) GetPolicyPathByAuthorityId(authorityId string) (pathMaps []structs.CasbinInfo) {
	list := c.Casbin().GetFilteredPolicy(0, authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, structs.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

// ClearCasbin 清除匹配的权限
func (c *CasbinService) ClearCasbin(v int, p ...string) bool {
	success, _ := c.Casbin().RemoveFilteredPolicy(v, p...)
	return success

}

// UpdateCasbinApi API更新
func (c *CasbinService) UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	return c.db.Model(&structs.CasbinRule{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
}

// UpdateCasbin 更新casbin权限
func (c *CasbinService) UpdateCasbin(authorityId string, casbinInfos []structs.CasbinInfo) error {
	c.ClearCasbin(0, authorityId)
	var rules [][]string
	for _, v := range casbinInfos {
		cm := structs.CasbinRule{
			Ptype:       "p",
			AuthorityId: authorityId,
			Path:        v.Path,
			Method:      v.Method,
		}
		rules = append(rules, []string{cm.AuthorityId, cm.Path, cm.Method})
	}
	success, err := c.Casbin().AddPolicies(rules)
	if err != nil {
		return err
	}
	if !success {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	return nil
}
