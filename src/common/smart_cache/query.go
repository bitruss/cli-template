package smart_cache

import (
	"context"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/plugin/redis_plugin"
	"github.com/coreservice-io/cli-template/plugin/reference_plugin"
)

func SmartQuery(key string, fromCache bool, updateCache bool, resultHolder interface{}, ttl_secs int64, DBQuery func(resultHolder interface{}) error, query_description string) (interface{}, error) {

	if fromCache {
		// try to get from reference
		result := Ref_Get(reference_plugin.GetInstance(), key)
		if result != nil {
			basic.Logger.Debugln(query_description + " hit from reference")
			return result, nil
		}

		err := Redis_Get(context.Background(), redis_plugin.GetInstance().ClusterClient, true, key, resultHolder)
		if err == nil {
			basic.Logger.Debugln(query_description, " hit from redis")
			Ref_Set(reference_plugin.GetInstance(), key, resultHolder)
			return resultHolder, nil
		} else if err == ErrNil {
			//continue to get from db part
		} else if err == ErrTempNil {
			//this happens when query db failed
			basic.Logger.Errorln(query_description, " smart_cache.TempNil")
			return nil, ErrTempNil
		} else {
			//redis may broken, just return to keep db safe
			return resultHolder, err
		}
	}

	//after cache miss ,try from remote database
	basic.Logger.Debugln(query_description, " try from database")

	err := DBQuery(resultHolder)

	if err != nil {
		basic.Logger.Errorln(query_description, " DBQuery err :", err)
		//set err_nil for db fast re-query safety
		RR_SetErrTempNil(context.Background(), redis_plugin.GetInstance().ClusterClient, key)
		return nil, err
	} else {
		if updateCache {
			RR_Set(context.Background(), redis_plugin.GetInstance().ClusterClient, reference_plugin.GetInstance(), true, key, resultHolder, ttl_secs)
		}
		return resultHolder, nil
	}
}
