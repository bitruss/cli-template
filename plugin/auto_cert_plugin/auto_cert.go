package auto_cert_plugin

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/coreservice-io/job"
)

var instanceMap = map[string]*Cert{}

type Cert struct {
	Download_url   string
	Local_crt_path string
	Local_key_path string
	Auto_updating  bool
}

type Config struct {
	Download_url   string
	Local_crt_path string
	Local_key_path string
}

func GetInstance() *Cert {
	return instanceMap["default"]
}

func GetInstance_(name string) *Cert {
	return instanceMap[name]
}

func (cert *Cert) AutoUpdate() {
	if !cert.Auto_updating {
		cert.Auto_updating = true
		job.Start(
			"cert_auto_update_job",
			// job process
			func() {
				cert.update_()
			},
			// onPanic callback, run if panic happened
			func(err interface{}) {

			},
			// job interval in seconds
			30,
			job.TYPE_PANIC_REDO,
			// check continue callback, the job will stop running if return false
			// the job will keep running if this callback is nil
			func(job *job.Job) bool {
				return true
			},
			// onFinish callback
			func(inst *job.Job) {

			},
		)
	}
}

type RemoteRespCert struct {
	Crt_content string `json:"crt_content"`
	Key_content string `json:"key_content"`
}

type RemoteResp struct {
	Cert         RemoteRespCert `json:"cert"`
	Meta_status  int64          `json:"meta_status"`
	Meta_message string         `json:"meta_message"`
}

//update from remote url
func (cert *Cert) update_() error {

	downloadClient := http.Client{
		Timeout: time.Second * 15, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, cert.Download_url, nil)
	if err != nil {
		return err
	}

	res, getErr := downloadClient.Do(req)
	if getErr != nil {
		return getErr
	}

	if res.Body == nil {
		return errors.New("no body response")
	}

	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return readErr
	}

	rp := RemoteResp{}
	jsonErr := json.Unmarshal(body, &rp)
	if jsonErr != nil {
		return jsonErr
	}

	if rp.Meta_status <= 0 {
		return errors.New(rp.Meta_message)
	}

	//////save .crt and .key/////
	crt_file_err := file_overwrite(cert.Local_crt_path, rp.Cert.Crt_content)
	if crt_file_err != nil {
		return crt_file_err
	}

	key_file_err := file_overwrite(cert.Local_key_path, rp.Cert.Key_content)
	if key_file_err != nil {
		return key_file_err
	}

	return nil
}

func file_overwrite(path string, content string) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	_, werr := f.WriteString(content)
	if werr != nil {
		return werr
	}
	return nil
}

func Init(conf Config) error {
	return Init_("default", conf)
}

func Init_(name string, conf Config) error {
	if name == "" {
		name = "default"
	}

	_, exist := instanceMap[name]
	if exist {
		return fmt.Errorf("cert instance <%s> has already been initialized", name)
	}

	if conf.Download_url == "" || conf.Local_crt_path == "" || conf.Local_key_path == "" {
		return errors.New("params Download_url|Local_crt_path|Local_key_path must all be set ")
	}

	cert := &Cert{
		conf.Download_url,
		conf.Local_crt_path,
		conf.Local_key_path,
		false,
	}

	first_update_err := cert.update_()
	if first_update_err != nil {
		return errors.New("cert init failed," + first_update_err.Error())
	}

	instanceMap[name] = cert
	return nil
}
