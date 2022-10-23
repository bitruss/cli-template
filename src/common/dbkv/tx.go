package dbkv

import (
	"fmt"
	"strconv"

	"github.com/coreservice-io/cli-template/plugin/redis_plugin"
	"github.com/coreservice-io/cli-template/src/common/smart_cache"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SetDBKV_Str(tx *gorm.DB, keystr string, value string, description string) error {
	err := tx.Table("dbkv").Clauses(clause.OnConflict{UpdateAll: true}).Create(&DBKVModel{Key: keystr, Value: value, Description: description}).Error
	if err != nil {
		return err
	}
	return nil
}

func SetDBKV_UInt64(tx *gorm.DB, keystr string, value uint64, description string) error {
	return SetDBKV_Str(tx, keystr, strconv.FormatUint(value, 10), description)
}

func SetDBKV_Int64(tx *gorm.DB, keystr string, value int64, description string) error {
	return SetDBKV_Str(tx, keystr, strconv.FormatInt(value, 10), description)
}

func SetDBKV_Int32(tx *gorm.DB, keystr string, value int32, description string) error {
	return SetDBKV_Str(tx, keystr, strconv.FormatInt(int64(value), 10), description)
}

func SetDBKV_Int(tx *gorm.DB, keystr string, value int, description string) error {
	return SetDBKV_Str(tx, keystr, strconv.Itoa(value), description)
}

func SetDBKV_Bool(tx *gorm.DB, keystr string, value bool, description string) error {
	if value {
		return SetDBKV_Str(tx, keystr, "true", description)
	} else {
		return SetDBKV_Str(tx, keystr, "false", description)
	}
}

func SetDBKV_Float32(tx *gorm.DB, keystr string, value float32, description string) error {
	return SetDBKV_Str(tx, keystr, fmt.Sprintf("%f", value), description)
}

func SetDBKV_Float64(tx *gorm.DB, keystr string, value float64, description string) error {
	return SetDBKV_Str(tx, keystr, fmt.Sprintf("%f", value), description)
}

func DeleteDBKV_Key(tx *gorm.DB, keystr string) error {
	if err := tx.Table("dbkv").Where(" `key` = ?", keystr).Delete(&DBKVModel{}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteDBKV_Id(tx *gorm.DB, id int64) error {
	if err := tx.Table("dbkv").Where(" `id` = ?", id).Delete(&DBKVModel{}).Error; err != nil {
		return err
	}
	return nil
}

func GetDBKV(tx *gorm.DB, id *int64, key *string, fromCache bool, updateCache bool) (*DBKVModel, error) {

	//gen_key
	ck := smart_cache.NewConnectKey("dbkv")
	ck.C_Int64_Ptr("id", id).C_Str_Ptr("key", key)

	_key := redis_plugin.GetInstance().GenKey(ck.String())

	/////
	resultHolderAlloc := func() interface{} {
		return &DBKVModel{}
	}

	/////
	query := func(resultHolder interface{}) error {
		queryResult := resultHolder.(*DBKVModel)

		query := tx.Table("dbkv")
		if id != nil {
			query.Where("id = ?", *id)
		}
		if key != nil {
			query.Where("dbkv.key =?", *key)
		}

		var total_count int64
		query.Count(&total_count)

		err := query.Find(queryResult).Error
		if err != nil {
			return err
		}

		if total_count == 0 {
			return smart_cache.QueryNilErr
		}

		return nil
	}

	/////
	sq_result, sq_err := smart_cache.SmartQuery(_key, resultHolderAlloc, fromCache, updateCache, 300, query, "DBKV Query")

	/////
	if sq_err != nil {
		return nil, sq_err
	} else {
		return sq_result.(*DBKVModel), nil
	}
}

type DBKVQueryResults struct {
	Kv         []*DBKVModel
	TotalCount int64
}

func QueryDBKV(tx *gorm.DB, id *int64, keys *[]string, fromCache bool, updateCache bool) (*DBKVQueryResults, error) {

	//gen_key
	ck := smart_cache.NewConnectKey("dbkv")
	ck.C_Int64_Ptr("id", id).C_Str_Array_Ptr("keys", keys)

	key := redis_plugin.GetInstance().GenKey(ck.String())

	/////
	resultHolderAlloc := func() interface{} {
		return &DBKVQueryResults{
			Kv:         []*DBKVModel{},
			TotalCount: 0,
		}
	}

	/////
	query := func(resultHolder interface{}) error {
		queryResults := resultHolder.(*DBKVQueryResults)

		query := tx.Table("dbkv")
		if id != nil {
			query.Where("id = ?", *id)
		}
		if keys != nil {
			query.Where("dbkv.key IN ?", *keys)
		}

		query.Count(&queryResults.TotalCount)

		err := query.Find(&queryResults.Kv).Error
		if err != nil {
			return err
		}

		return nil
	}

	/////
	sq_result, sq_err := smart_cache.SmartQuery(key, resultHolderAlloc, fromCache, updateCache, 300, query, "DBKV Query")

	/////
	if sq_err != nil {
		return nil, sq_err
	} else {
		return sq_result.(*DBKVQueryResults), nil
	}

}
