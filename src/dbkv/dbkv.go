package dbkv

import (
	"context"

	"github.com/coreservice-io/CliAppTemplate/basic"
	"github.com/coreservice-io/CliAppTemplate/plugin/redis_plugin"
	"github.com/coreservice-io/CliAppTemplate/plugin/reference_plugin"
	"github.com/coreservice-io/CliAppTemplate/tools/smartCache"
	"github.com/go-redis/redis/v8"
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
	ck := smartCache.NewConnectKey("dbkv")
	ck.C_Str(keyStr)

	key := redis_plugin.GetInstance().GenKey(ck.String())

	if fromCache {
		// try to get from reference
		result := smartCache.Ref_Get(reference_plugin.GetInstance(), key)
		if result != nil {
			basic.Logger.Debugln("GetDBKV hit from reference")
			return result.(*DBKVModel), nil
		}

		redis_result := &DBKVModel{}
		// try to get from redis
		err := smartCache.Redis_Get(context.Background(), redis_plugin.GetInstance().ClusterClient, true, key, redis_result)
		if err == nil {
			basic.Logger.Debugln("GetDBKV hit from redis")
			smartCache.Ref_Set(reference_plugin.GetInstance(), key, redis_result)
			return redis_result, nil
		} else if err == redis.Nil {
			//continue to get from db part
		} else {
			//redis may broken, just return to keep db safe
			return redis_result, err
		}
	}

	//after cache miss ,try from remote database
	basic.Logger.Debugln("GetDBKV try from database")

	queryResult := &DBKVModel{}

	query := tx.Table("dbkv")
	query.Where("`key` = ?", keyStr)

	err := query.First(queryResult).Error
	if err != nil {
		basic.Logger.Errorln("GetDBKV err :", err)
		return nil, err
	} else {
		if updateCache {
			smartCache.RR_Set(context.Background(), redis_plugin.GetInstance().ClusterClient, reference_plugin.GetInstance(), true, key, queryResult, 300)
		}
		return queryResult, nil
	}
}
