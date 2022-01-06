package components

import (
	"context"
	"errors"
	"net/http"
	"time"

	elasticSearch "github.com/olivere/elastic/v7"
	"github.com/universe-30/CliAppTemplate/basic"
)

type ElasticSRetrier struct {
}

func (r *ElasticSRetrier) Retry(ctx context.Context, retry int, req *http.Request, resp *http.Response, err error) (time.Duration, bool, error) {
	return 120 * time.Second, true, nil //retry after 2mins
}

/*
elasticSearchAddr
elasticSearchUserName
*/

func NewElasticSearch() (*elasticSearch.Client, error) {

	elasticSearchAddr, elasticSearchAddr_Err := basic.Config.GetString("elasticsearch_addr", "")
	if elasticSearchAddr_Err != nil {
		return nil, errors.New("elasticsearch_addr [string] in config.json not defined," + elasticSearchAddr_Err.Error())
	}

	elasticSearchUserName, elasticSearchUserName_Err := basic.Config.GetString("elasticsearch_username", "")
	if elasticSearchUserName_Err != nil {
		return nil, errors.New("elasticsearch_username_err [string] in config.json not defined," + elasticSearchUserName_Err.Error())
	}

	elasticSearchPassword, elasticSearchPassword_Err := basic.Config.GetString("elasticsearch_password", "")
	if elasticSearchPassword_Err != nil {
		return nil, errors.New("elasticsearch_password [string] in config.json not defined," + elasticSearchPassword_Err.Error())
	}

	ElasticSClient, err := elasticSearch.NewClient(
		elasticSearch.SetURL(elasticSearchAddr),
		elasticSearch.SetBasicAuth(elasticSearchUserName, elasticSearchPassword),
		elasticSearch.SetSniff(false),
		elasticSearch.SetHealthcheckInterval(30*time.Second),
		elasticSearch.SetRetrier(&ElasticSRetrier{}),
		elasticSearch.SetGzip(true),
	)
	if err != nil {
		return nil, err
	}

	return ElasticSClient, nil
}
