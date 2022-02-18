package es

import (
	"context"
	"fmt"
	"net/http"
	"time"

	elasticSearch "github.com/olivere/elastic/v7"
)

type ElasticSRetrier struct {
}

func (r *ElasticSRetrier) Retry(ctx context.Context, retry int, req *http.Request, resp *http.Response, err error) (time.Duration, bool, error) {
	return 120 * time.Second, true, nil //retry after 2mins
}

var instanceMap = map[string]*elasticSearch.Client{}

func GetDefaultInstance() *elasticSearch.Client {
	return instanceMap["default"]
}

func GetInstance(name string) *elasticSearch.Client {
	return instanceMap[name]
}

/*
elasticSearchAddr
elasticSearchUserName
elasticSearchPassword
*/
type Config struct {
	Address  string
	UserName string
	Password string
}

// Init a new instance.
//  If only need one instance, use empty name "". Use GetDefaultInstance() to get.
//  If you need several instance, run Init() with different <name>. Use GetInstance(<name>) to get.
func Init(name string, esConfig Config) error {
	if name == "" {
		name = "default"
	}

	_, exist := instanceMap[name]
	if exist {
		return fmt.Errorf("elasticSearch instance <%s> has already initialized", name)
	}

	es, err := elasticSearch.NewClient(
		elasticSearch.SetURL(esConfig.Address),
		elasticSearch.SetBasicAuth(esConfig.UserName, esConfig.Password),
		elasticSearch.SetSniff(false),
		elasticSearch.SetHealthcheckInterval(30*time.Second),
		elasticSearch.SetRetrier(&ElasticSRetrier{}),
		elasticSearch.SetGzip(true),
	)
	if err != nil {
		return err
	}
	instanceMap[name] = es
	return nil
}