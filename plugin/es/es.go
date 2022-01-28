package es

import (
	"context"
	"errors"
	"net/http"
	"time"

	elasticSearch "github.com/olivere/elastic/v7"
	"github.com/universe-30/CliAppTemplate/configuration"
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

var es *elasticSearch.Client

func GetSingleInstance() *elasticSearch.Client {
	return es
}

func Init() error {
	elasticSearchAddr, err := configuration.Config.GetString("elasticsearch_addr", "")
	if err != nil {
		return errors.New("elasticsearch_addr [string] in config error," + err.Error())
	}

	elasticSearchUserName, err := configuration.Config.GetString("elasticsearch_username", "")
	if err != nil {
		return errors.New("elasticsearch_username_err [string] in config error," + err.Error())
	}

	elasticSearchPassword, err := configuration.Config.GetString("elasticsearch_password", "")
	if err != nil {
		return errors.New("elasticsearch_password [string] in config error," + err.Error())
	}

	es, err = elasticSearch.NewClient(
		elasticSearch.SetURL(elasticSearchAddr),
		elasticSearch.SetBasicAuth(elasticSearchUserName, elasticSearchPassword),
		elasticSearch.SetSniff(false),
		elasticSearch.SetHealthcheckInterval(30*time.Second),
		elasticSearch.SetRetrier(&ElasticSRetrier{}),
		elasticSearch.SetGzip(true),
	)
	if err != nil {
		return err
	}
	return nil
}
