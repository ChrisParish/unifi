package unidev

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"net/http/cookiejar"
	"strconv"

	influx "github.com/influxdata/influxdb/client/v2"
	"github.com/pkg/errors"
)

// LoginPath is Unifi Controller Login API Path
const LoginPath = "/api/login"

// Asset provides a common interface to retreive metrics from a device or client.
// It currently only supports InfluxDB, but could be amended to support other
// libraries that have a similar interface.
// This app only uses the .AddPoint/s() methods with the Asset type.
type Asset interface {
	// Point() means this is useful to influxdb..
	Points() ([]*influx.Point, error)
	// Add more methods to achieve more usefulness from this library.
}

// AuthedReq is what you get in return for providing a password!
type AuthedReq struct {
	*http.Client
	baseURL string
}

// FlexInt provides a container and unmarshalling for fields that may be
// numbers or strings in the Unifi API
type FlexInt struct {
	Number float64
	String string
}

// UnmarshalJSON converts a string or number to an integer.
func (f *FlexInt) UnmarshalJSON(b []byte) error {
	var unk interface{}
	if err := json.Unmarshal(b, &unk); err != nil {
		return err
	}
	switch i := unk.(type) {
	case float64:
		f.Number = i
		f.String = strconv.FormatFloat(i, 'f', -1, 64)
		return nil
	case string:
		f.String = i
		f.Number, _ = strconv.ParseFloat(i, 64)
		return nil
	default:
		return errors.New("Cannot unmarshal to FlexInt")
	}
}

// AuthController creates a http.Client with authenticated cookies.
// Used to make additional, authenticated requests to the APIs.
func AuthController(user, pass, url string, verifySSL bool) (*AuthedReq, error) {
	json := `{"username": "` + user + `","password": "` + pass + `"}`
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, errors.Wrap(err, "cookiejar.New(nil)")
	}
	a := &AuthedReq{&http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: !verifySSL}},
		Jar:       jar,
	}, url}
	req, err := a.UniReq(LoginPath, json)
	if err != nil {
		return a, errors.Wrap(err, "UniReq(LoginPath, json)")
	}
	resp, err := a.Do(req)
	if err != nil {
		return a, errors.Wrap(err, "authReq.Do(req)")
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println("resp.Body.Close():", err) // Not fatal. Just log it.
		}
	}()
	if resp.StatusCode != http.StatusOK {
		return a, errors.Errorf("authentication failed (%v): %v (status: %v/%v)",
			user, url+LoginPath, resp.StatusCode, resp.Status)
	}
	return a, nil
}

// UniReq is a small helper function that adds an Accept header.
func (a AuthedReq) UniReq(apiPath string, params string) (req *http.Request, err error) {
	if params != "" {
		req, err = http.NewRequest("POST", a.baseURL+apiPath, bytes.NewBufferString(params))
	} else {
		req, err = http.NewRequest("GET", a.baseURL+apiPath, nil)
	}
	if err == nil {
		req.Header.Add("Accept", "application/json")
	}
	return
}
