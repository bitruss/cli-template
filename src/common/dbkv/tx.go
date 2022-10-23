package dbkv

import (
	"github.com/coreservice-io/cli-template/plugin/redis_plugin"
	"github.com/coreservice-io/cli-template/src/common/smart_cache"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SetDBKV(tx *gorm.DB, keystr string, value string) error {
	err := tx.Table("dbkv").Clauses(clause.OnConflict{UpdateAll: true}).Create(&DBKVModel{Key: keystr, Value: value}).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteDBKV(tx *gorm.DB, keystr string) error {
	if err := tx.Table("dbkv").Where(" `key` = ?", keystr).Delete(&DBKVModel{}).Error; err != nil {
		return err
	}
	return nil
}

func GetDBKV(tx *gorm.DB, keyStr string, fromCache bool, updateCache bool) (*DBKVModel, error) {

	//gen_key
	ck := smart_cache.NewConnectKey("dbkv")
	ck.C_Str(keyStr)

	key := redis_plugin.GetInstance().GenKey(ck.String())

	/////
	resultHolderAlloc := func() interface{} {
		return &DBKVModel{}
	}

	/////
	query := func(resultHolder interface{}) error {
		queryResult := resultHolder.(*DBKVModel)

		queryResults := []*DBKVModel{}

		err := tx.Table("dbkv").Where("`key` = ?", keyStr).Find(&queryResults).Error
		if err != nil {
			return err
		}

		if len(queryResults) == 0 {
			return smart_cache.QueryNilErr
		}

		*queryResult = *queryResults[0]
		return nil
	}

	/////
	sq_result, sq_err := smart_cache.SmartQuery(key, resultHolderAlloc, fromCache, updateCache, 300, query, "DBKV Query")

	/////
	if sq_err != nil {
		return nil, sq_err
	} else {
		return sq_result.(*DBKVModel), nil
	}

}
