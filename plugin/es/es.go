package es

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"time"

	elasticSearch "github.com/olivere/elastic/v7"
	"github.com/universe-30/CliAppTemplate/basic"
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
var once sync.Once

func Init() {
	//only run once
	once.Do(func() {
		var err error = nil
		es, err = newElasticSearch()
		if err != nil {
			basic.Logger.Fatalln(err)
		}
	})
}

func GetSingleInstance() *elasticSearch.Client {
	Init()
	return es
}

func newElasticSearch() (*elasticSearch.Client, error) {

	elasticSearchAddr, elasticSearchAddr_Err := configuration.Config.GetString("elasticsearch_addr", "")
	if elasticSearchAddr_Err != nil {
		return nil, errors.New("elasticsearch_addr [string] in config error," + elasticSearchAddr_Err.Error())
	}

	elasticSearchUserName, elasticSearchUserName_Err := configuration.Config.GetString("elasticsearch_username", "")
	if elasticSearchUserName_Err != nil {
		return nil, errors.New("elasticsearch_username_err [string] in config error," + elasticSearchUserName_Err.Error())
	}

	elasticSearchPassword, elasticSearchPassword_Err := configuration.Config.GetString("elasticsearch_password", "")
	if elasticSearchPassword_Err != nil {
		return nil, errors.New("elasticsearch_password [string] in config error," + elasticSearchPassword_Err.Error())
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
