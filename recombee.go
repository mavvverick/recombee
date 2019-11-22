package recombee

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	libraryVersion = "1.0.0"
	defaultBaseURL = "http://rapi.recombee.com/"
	userAgent      = "reco/" + libraryVersion
	mediaType      = "application/json"
)

var db = ""
var secret = ""

//os.Getenv("RECOMBEE_SECRET")

var logics = map[string]string{
	"recombee:default":               "recombee:default",
	"recombee:homepage":              "recombee:homepage",
	"recombee:personal":              "recombee:personal",
	"recombee:similar":               "recombee:similar",
	"recombee:popular":               "recombee:popular",
	"recombee:recently-viewed":       "recombee:recently-viewed",
	"ecommerce:homepage":             "ecommerce:homepage",
	"ecommerce:cross-sell":           "ecommerce:cross-sell",
	"ecommerce:bestseller":           "ecommerce:bestseller",
	"ecommerce:similarly-purchasing": "ecommerce:similarly-purchasing",
	"classifieds:homepage":           "classifieds:homepage",
	"classifieds:personal":           "classifieds:personal",
	"classifieds:similar-ads":        "classifieds:similar-ads",
}

// Client manages communication with Recombee API.
type Client struct {
	client *http.Client

	BaseURL *url.URL

	UserAgent string
	DB        string
	Secret    string
	//Rate Rate
	Item  ItemService
	Reco  RecoService
	User  UserService
	Admin AdminService
	Batch BatchService
	// Optional function called after every successful request made to the DO APIs
	onRequestCompleted RequestCompletionCallback
}

type Response struct {
	*http.Response

	// Links that were returned with the response. These are parsed from
	// request body and not the header.
	//Links *Links

	// Monitoring URI
	Monitor string

	//Rate
}

type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response

	// Error message
	StatusCode string `json:"statusCode"`
	Message    string `json:"message"`
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

// NewClient returns a new Recombee API client.
func NewClient(httpClient *http.Client, dbKey string, secretKey string) *Client {
	if httpClient == nil {
		//httpClient = http.DefaultClient
		httpClient = &http.Client{
			Timeout: time.Second * 300,
		}
	}

	baseURL, _ := url.Parse(defaultBaseURL)
	db = dbKey
	secret = secretKey
	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent, DB: db, Secret: secret}
	c.Item = &ItemServiceOp{client: c}
	c.Reco = &RecoServiceOp{client: c}
	c.User = &UserServiceOp{client: c}
	c.Admin = &AdminServiceOp{client: c}
	c.Batch = &BatchServiceOp{client: c}
	return c
}

func (c *Client) NewRequest(ctx context.Context, method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)

	if body != nil {
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", mediaType)
	req.Header.Add("Accept", mediaType)
	req.Header.Add("User-Agent", c.UserAgent)

	return req, err
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && len(data) > 0 {
		err := json.Unmarshal(data, errorResponse)
		if err != nil {
			errorResponse.Message = string(data)
		}
	}

	return errorResponse
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	resp, err := DoRequestWithClient(ctx, c.client, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if rerr := resp.Body.Close(); err != nil {
			err = rerr
		}
	}()

	response := newResponse(resp)

	err = CheckResponse(resp)

	if err != nil {
		return response, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
			if err != nil {
				return nil, err
			}
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err != nil {
				return nil, err
			}
		}
	}

	return response, err

}

func newResponse(r *http.Response) *Response {
	response := Response{Response: r}
	//TODO check rate limiting
	return &response
}

func (r *ErrorResponse) Error() string {
	if r.StatusCode != "" {
		return fmt.Sprintf("%v %v: %d (request %q) %v",
			r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.StatusCode, r.Message)
	}
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message)
}

func DoRequestWithClient(
	ctx context.Context,
	client *http.Client,
	req *http.Request) (*http.Response, error) {
	req = req.WithContext(ctx)
	return client.Do(req)
}

func GenURL(url string) string {
	urlWithEpoch := url + "hmac_timestamp=" + strconv.FormatInt(time.Now().Unix(), 10)
	hash := GenHasH(urlWithEpoch)
	urlWithEpoch += "&hmac_sign=" + hash
	return urlWithEpoch
}

func GenHasH(message string) string {
	hash := hmac.New(sha1.New, []byte(secret))
	hash.Write([]byte(message))
	str := hash.Sum(nil)
	return hex.EncodeToString(str)
}
