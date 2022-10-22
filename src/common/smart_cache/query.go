package smart_cache

import (
	"context"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/plugin/redis_plugin"
	"github.com/coreservice-io/cli-template/plugin/reference_plugin"
)

func SmartQuery(key string, resultHolderAlloc func() interface{}, fromCache bool, updateCache bool, cacheTTLSecs int64, DBQuery func(resultHolder interface{}) error, queryDescription string) (interface{}, error) {

	var resultHolder interface{}

	if fromCache {
		// try to get from reference
		result := Ref_Get(reference_plugin.GetInstance(), key)
		if result != nil {
			basic.Logger.Debugln(queryDescription + " hit from reference")
			return result, nil
		}

		resultHolder = resultHolderAlloc()

		err := Redis_Get(context.Background(), redis_plugin.GetInstance().ClusterClient, true, key, resultHolder)
		if err == nil {
			basic.Logger.Debugln(queryDescription, " hit from redis")
			Ref_Set(reference_plugin.GetInstance(), key, resultHolder)
			return resultHolder, nil
		} else if err == ErrNil {
			//continue to get from db part
		} else if err == ErrTempNil {
			//this happens when query db failed
			basic.Logger.Errorln(queryDescription, " smart_cache.TempNil")
			return nil, ErrTempNil
		} else {
			//redis may broken, just return to keep db safe
			return resultHolder, err
		}
	}

	//after cache miss ,try from remote database
	basic.Logger.Debugln(queryDescription, " try from database")

	err := DBQuery(resultHolder)

	if err != nil {
		basic.Logger.Errorln(queryDescription, " DBQuery err :", err)
		//set err_nil for db fast re-query safety
		RR_SetErrTempNil(context.Background(), redis_plugin.GetInstance().ClusterClient, key)
		return nil, err
	} else {
		if updateCache {
			RR_Set(context.Background(), redis_plugin.GetInstance().ClusterClient, reference_plugin.GetInstance(), true, key, resultHolder, cacheTTLSecs)
		}
		return resultHolder, nil
	}
}
