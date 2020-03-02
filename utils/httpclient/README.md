# httpclient
> 是对resty的包装,为了好用. https://github.com/go-resty/resty
Installation
# Go Modules
require github.com/go-resty/resty/v2 v2.2.0
Usage
The following samples will assist you to become as comfortable as possible with resty library.

// Import resty into your code and refer it as `resty`
`import resty  "github.com/go-resty/resty/v2"`


Usage
The following samples will assist you to become as comfortable as possible with resty library.

// Import resty into your code and refer it as `resty`.
import "github.com/go-resty/resty/v2"
Simple GET
// Create a Resty Client
client := resty.New()

resp, err := client.R().
		EnableTrace().
		Get("https://httpbin.org/get")

// Explore response object
fmt.Println("Response Info:")
fmt.Println("Error      :", err)
fmt.Println("Status Code:", resp.StatusCode())
fmt.Println("Status     :", resp.Status())
fmt.Println("Time       :", resp.Time())
fmt.Println("Received At:", resp.ReceivedAt())
fmt.Println("Body       :\n", resp)
fmt.Println()

// Explore trace info
fmt.Println("Request Trace Info:")
ti := resp.Request.TraceInfo()
fmt.Println("DNSLookup    :", ti.DNSLookup)
fmt.Println("ConnTime     :", ti.ConnTime)
fmt.Println("TLSHandshake :", ti.TLSHandshake)
fmt.Println("ServerTime   :", ti.ServerTime)
fmt.Println("ResponseTime :", ti.ResponseTime)
fmt.Println("TotalTime    :", ti.TotalTime)
fmt.Println("IsConnReused :", ti.IsConnReused)
fmt.Println("IsConnWasIdle:", ti.IsConnWasIdle)
fmt.Println("ConnIdleTime :", ti.ConnIdleTime)

/* Output
Response Info:
Error      : <nil>
Status Code: 200
Status     : 200 OK
Time       : 465.301137ms
Received At: 2019-06-16 01:52:33.772456 -0800 PST m=+0.466672260
Body       :
 {
  "args": {},
  "headers": {
    "Accept-Encoding": "gzip",
    "Host": "httpbin.org",
    "User-Agent": "go-resty/2.0.0 (https://github.com/go-resty/resty)"
  },
  "origin": "0.0.0.0",
  "url": "https://httpbin.org/get"
}

Request Trace Info:
DNSLookup    : 2.21124ms
ConnTime     : 393.875795ms
TLSHandshake : 319.313546ms
ServerTime   : 71.109256ms
ResponseTime : 94.466µs
TotalTime    : 465.301137ms
IsConnReused : false
IsConnWasIdle: false
ConnIdleTime : 0s
*/
Enhanced GET
// Create a Resty Client
client := resty.New()

resp, err := client.R().
      SetQueryParams(map[string]string{
          "page_no": "1",
          "limit": "20",
          "sort":"name",
          "order": "asc",
          "random":strconv.FormatInt(time.Now().Unix(), 10),
      }).
      SetHeader("Accept", "application/json").
      SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F").
      Get("/search_result")


// Sample of using Request.SetQueryString method
resp, err := client.R().
      SetQueryString("productId=232&template=fresh-sample&cat=resty&source=google&kw=buy a lot more").
      SetHeader("Accept", "application/json").
      SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F").
      Get("/show_product")
Various POST method combinations
// Create a Resty Client
client := resty.New()

// POST JSON string
// No need to set content type, if you have client level setting
resp, err := client.R().
      SetHeader("Content-Type", "application/json").
      SetBody(`{"username":"testuser", "password":"testpass"}`).
      SetResult(&AuthSuccess{}).    // or SetResult(AuthSuccess{}).
      Post("https://myapp.com/login")

// POST []byte array
// No need to set content type, if you have client level setting
resp, err := client.R().
      SetHeader("Content-Type", "application/json").
      SetBody([]byte(`{"username":"testuser", "password":"testpass"}`)).
      SetResult(&AuthSuccess{}).    // or SetResult(AuthSuccess{}).
      Post("https://myapp.com/login")

// POST Struct, default is JSON content type. No need to set one
resp, err := client.R().
      SetBody(User{Username: "testuser", Password: "testpass"}).
      SetResult(&AuthSuccess{}).    // or SetResult(AuthSuccess{}).
      SetError(&AuthError{}).       // or SetError(AuthError{}).
      Post("https://myapp.com/login")

// POST Map, default is JSON content type. No need to set one
resp, err := client.R().
      SetBody(map[string]interface{}{"username": "testuser", "password": "testpass"}).
      SetResult(&AuthSuccess{}).    // or SetResult(AuthSuccess{}).
      SetError(&AuthError{}).       // or SetError(AuthError{}).
      Post("https://myapp.com/login")

// POST of raw bytes for file upload. For example: upload file to Dropbox
fileBytes, _ := ioutil.ReadFile("/Users/jeeva/mydocument.pdf")

// See we are not setting content-type header, since go-resty automatically detects Content-Type for you
resp, err := client.R().
      SetBody(fileBytes).
      SetContentLength(true).          // Dropbox expects this value
      SetAuthToken("<your-auth-token>").
      SetError(&DropboxError{}).       // or SetError(DropboxError{}).
      Post("https://content.dropboxapi.com/1/files_put/auto/resty/mydocument.pdf") // for upload Dropbox supports PUT too

// Note: resty detects Content-Type for request body/payload if content type header is not set.
//   * For struct and map data type defaults to 'application/json'
//   * Fallback is plain text content type
Sample PUT
You can use various combinations of PUT method call like demonstrated for POST.

// Note: This is one sample of PUT method usage, refer POST for more combination

// Create a Resty Client
client := resty.New()

// Request goes as JSON content type
// No need to set auth token, error, if you have client level settings
resp, err := client.R().
      SetBody(Article{
        Title: "go-resty",
        Content: "This is my article content, oh ya!",
        Author: "Jeevanandam M",
        Tags: []string{"article", "sample", "resty"},
      }).
      SetAuthToken("C6A79608-782F-4ED0-A11D-BD82FAD829CD").
      SetError(&Error{}).       // or SetError(Error{}).
      Put("https://myapp.com/article/1234")
Sample PATCH
You can use various combinations of PATCH method call like demonstrated for POST.

// Note: This is one sample of PUT method usage, refer POST for more combination

// Create a Resty Client
client := resty.New()

// Request goes as JSON content type
// No need to set auth token, error, if you have client level settings
resp, err := client.R().
      SetBody(Article{
        Tags: []string{"new tag1", "new tag2"},
      }).
      SetAuthToken("C6A79608-782F-4ED0-A11D-BD82FAD829CD").
      SetError(&Error{}).       // or SetError(Error{}).
      Patch("https://myapp.com/articles/1234")
Sample DELETE, HEAD, OPTIONS
// Create a Resty Client
client := resty.New()

// DELETE a article
// No need to set auth token, error, if you have client level settings
resp, err := client.R().
      SetAuthToken("C6A79608-782F-4ED0-A11D-BD82FAD829CD").
      SetError(&Error{}).       // or SetError(Error{}).
      Delete("https://myapp.com/articles/1234")

// DELETE a articles with payload/body as a JSON string
// No need to set auth token, error, if you have client level settings
resp, err := client.R().
      SetAuthToken("C6A79608-782F-4ED0-A11D-BD82FAD829CD").
      SetError(&Error{}).       // or SetError(Error{}).
      SetHeader("Content-Type", "application/json").
      SetBody(`{article_ids: [1002, 1006, 1007, 87683, 45432] }`).
      Delete("https://myapp.com/articles")

// HEAD of resource
// No need to set auth token, if you have client level settings
resp, err := client.R().
      SetAuthToken("C6A79608-782F-4ED0-A11D-BD82FAD829CD").
      Head("https://myapp.com/videos/hi-res-video")

// OPTIONS of resource
// No need to set auth token, if you have client level settings
resp, err := client.R().
      SetAuthToken("C6A79608-782F-4ED0-A11D-BD82FAD829CD").
      Options("https://myapp.com/servers/nyc-dc-01")
Multipart File(s) upload
Using io.Reader
profileImgBytes, _ := ioutil.ReadFile("/Users/jeeva/test-img.png")
notesBytes, _ := ioutil.ReadFile("/Users/jeeva/text-file.txt")

// Create a Resty Client
client := resty.New()

resp, err := client.R().
      SetFileReader("profile_img", "test-img.png", bytes.NewReader(profileImgBytes)).
      SetFileReader("notes", "text-file.txt", bytes.NewReader(notesBytes)).
      SetFormData(map[string]string{
          "first_name": "Jeevanandam",
          "last_name": "M",
      }).
      Post("http://myapp.com/upload")
Using File directly from Path
// Create a Resty Client
client := resty.New()

// Single file scenario
resp, err := client.R().
      SetFile("profile_img", "/Users/jeeva/test-img.png").
      Post("http://myapp.com/upload")

// Multiple files scenario
resp, err := client.R().
      SetFiles(map[string]string{
        "profile_img": "/Users/jeeva/test-img.png",
        "notes": "/Users/jeeva/text-file.txt",
      }).
      Post("http://myapp.com/upload")

// Multipart of form fields and files
resp, err := client.R().
      SetFiles(map[string]string{
        "profile_img": "/Users/jeeva/test-img.png",
        "notes": "/Users/jeeva/text-file.txt",
      }).
      SetFormData(map[string]string{
        "first_name": "Jeevanandam",
        "last_name": "M",
        "zip_code": "00001",
        "city": "my city",
        "access_token": "C6A79608-782F-4ED0-A11D-BD82FAD829CD",
      }).
      Post("http://myapp.com/profile")
Sample Form submission
// Create a Resty Client
client := resty.New()

// just mentioning about POST as an example with simple flow
// User Login
resp, err := client.R().
      SetFormData(map[string]string{
        "username": "jeeva",
        "password": "mypass",
      }).
      Post("http://myapp.com/login")

// Followed by profile update
resp, err := client.R().
      SetFormData(map[string]string{
        "first_name": "Jeevanandam",
        "last_name": "M",
        "zip_code": "00001",
        "city": "new city update",
      }).
      Post("http://myapp.com/profile")

// Multi value form data
criteria := url.Values{
  "search_criteria": []string{"book", "glass", "pencil"},
}
resp, err := client.R().
      SetFormDataFromValues(criteria).
      Post("http://myapp.com/search")
Save HTTP Response into File
// Create a Resty Client
client := resty.New()

// Setting output directory path, If directory not exists then resty creates one!
// This is optional one, if you're planning using absoule path in
// `Request.SetOutput` and can used together.
client.SetOutputDirectory("/Users/jeeva/Downloads")

// HTTP response gets saved into file, similar to curl -o flag
_, err := client.R().
          SetOutput("plugin/ReplyWithHeader-v5.1-beta.zip").
          Get("http://bit.ly/1LouEKr")

// OR using absolute path
// Note: output directory path is not used for absolute path
_, err := client.R().
          SetOutput("/MyDownloads/plugin/ReplyWithHeader-v5.1-beta.zip").
          Get("http://bit.ly/1LouEKr")
Request URL Path Params
Resty provides easy to use dynamic request URL path params. Params can be set at client and request level. Client level params value can be overridden at request level.

// Create a Resty Client
client := resty.New()

client.R().SetPathParams(map[string]string{
   "userId": "sample@sample.com",
   "subAccountId": "100002",
}).
Get("/v1/users/{userId}/{subAccountId}/details")

// Result:
//   Composed URL - /v1/users/sample@sample.com/100002/details
Request and Response Middleware
Resty provides middleware ability to manipulate for Request and Response. It is more flexible than callback approach.

// Create a Resty Client
client := resty.New()

// Registering Request Middleware
client.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
    // Now you have access to Client and current Request object
    // manipulate it as per your need

    return nil  // if its success otherwise return error
  })

// Registering Response Middleware
client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
    // Now you have access to Client and current Response object
    // manipulate it as per your need

    return nil  // if its success otherwise return error
  })
Redirect Policy
Resty provides few ready to use redirect policy(s) also it supports multiple policies together.

// Create a Resty Client
client := resty.New()

// Assign Client Redirect Policy. Create one as per you need
client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(15))

// Wanna multiple policies such as redirect count, domain name check, etc
client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(20),
                        resty.DomainCheckRedirectPolicy("host1.com", "host2.org", "host3.net"))
Custom Redirect Policy
Implement RedirectPolicy interface and register it with resty client. Have a look redirect.go for more information.

// Create a Resty Client
client := resty.New()

// Using raw func into resty.SetRedirectPolicy
client.SetRedirectPolicy(resty.RedirectPolicyFunc(func(req *http.Request, via []*http.Request) error {
  // Implement your logic here

  // return nil for continue redirect otherwise return error to stop/prevent redirect
  return nil
}))

//---------------------------------------------------

// Using struct create more flexible redirect policy
type CustomRedirectPolicy struct {
  // variables goes here
}

func (c *CustomRedirectPolicy) Apply(req *http.Request, via []*http.Request) error {
  // Implement your logic here

  // return nil for continue redirect otherwise return error to stop/prevent redirect
  return nil
}

// Registering in resty
client.SetRedirectPolicy(CustomRedirectPolicy{/* initialize variables */})
Custom Root Certificates and Client Certificates
// Create a Resty Client
client := resty.New()

// Custom Root certificates, just supply .pem file.
// you can add one or more root certificates, its get appended
client.SetRootCertificate("/path/to/root/pemFile1.pem")
client.SetRootCertificate("/path/to/root/pemFile2.pem")
// ... and so on!

// Adding Client Certificates, you add one or more certificates
// Sample for creating certificate object
// Parsing public/private key pair from a pair of files. The files must contain PEM encoded data.
cert1, err := tls.LoadX509KeyPair("certs/client.pem", "certs/client.key")
if err != nil {
  log.Fatalf("ERROR client certificate: %s", err)
}
// ...

// You add one or more certificates
client.SetCertificates(cert1, cert2, cert3)
Custom Root Certificates and Client Certificates from string
// Custom Root certificates from string
// You can pass you certificates throught env variables as strings
// you can add one or more root certificates, its get appended
client.SetRootCertificateFromString("-----BEGIN CERTIFICATE-----content-----END CERTIFICATE-----")
client.SetRootCertificateFromString("-----BEGIN CERTIFICATE-----content-----END CERTIFICATE-----")
// ... and so on!

// Adding Client Certificates, you add one or more certificates
// Sample for creating certificate object
// Parsing public/private key pair from a pair of files. The files must contain PEM encoded data.
cert1, err := tls.X509KeyPair([]byte("-----BEGIN CERTIFICATE-----content-----END CERTIFICATE-----"), []byte("-----BEGIN CERTIFICATE-----content-----END CERTIFICATE-----"))
if err != nil {
  log.Fatalf("ERROR client certificate: %s", err)
}
// ...

// You add one or more certificates
client.SetCertificates(cert1, cert2, cert3)
Proxy Settings - Client as well as at Request Level
Default Go supports Proxy via environment variable HTTP_PROXY. Resty provides support via SetProxy & RemoveProxy. Choose as per your need.

Client Level Proxy settings applied to all the request

// Create a Resty Client
client := resty.New()

// Setting a Proxy URL and Port
client.SetProxy("http://proxyserver:8888")

// Want to remove proxy setting
client.RemoveProxy()
Retries
Resty uses backoff to increase retry intervals after each attempt.

Usage example:

// Create a Resty Client
client := resty.New()

// Retries are configured per client
client.
    // Set retry count to non zero to enable retries
    SetRetryCount(3).
    // You can override initial retry wait time.
    // Default is 100 milliseconds.
    SetRetryWaitTime(5 * time.Second).
    // MaxWaitTime can be overridden as well.
    // Default is 2 seconds.
    SetRetryMaxWaitTime(20 * time.Second).
    // SetRetryAfter sets callback to calculate wait time between retries.
    // Default (nil) implies exponential backoff with jitter
    SetRetryAfter(func(client *Client, resp *Response) (time.Duration, error) {
        return 0, errors.New("quota exceeded")
    })
Above setup will result in resty retrying requests returned non nil error up to 3 times with delay increased after each attempt.

You can optionally provide client with custom retry conditions:

// Create a Resty Client
client := resty.New()

client.AddRetryCondition(
    // RetryConditionFunc type is for retry condition function
	  // input: non-nil Response OR request execution error
    func(r *resty.Response, err error) bool {
        return r.StatusCode() == http.StatusTooManyRequests
    },
)
Above example will make resty retry requests ended with 429 Too Many Requests status code.

Multiple retry conditions can be added.

It is also possible to use resty.Backoff(...) to get arbitrary retry scenarios implemented. Reference.

Allow GET request with Payload
// Create a Resty Client
client := resty.New()

// Allow GET request with Payload. This is disabled by default.
client.SetAllowGetMethodPayload(true)
Wanna Multiple Clients
// Here you go!
// Client 1
client1 := resty.New()
client1.R().Get("http://httpbin.org")
// ...

// Client 2
client2 := resty.New()
client2.R().Head("http://httpbin.org")
// ...

// Bend it as per your need!!!
Remaining Client Settings & its Options
// Create a Resty Client
client := resty.New()

// Unique settings at Client level
//--------------------------------
// Enable debug mode
client.SetDebug(true)

// Assign Client TLSClientConfig
// One can set custom root-certificate. Refer: http://golang.org/pkg/crypto/tls/#example_Dial
client.SetTLSClientConfig(&tls.Config{ RootCAs: roots })

// or One can disable security check (https)
client.SetTLSClientConfig(&tls.Config{ InsecureSkipVerify: true })

// Set client timeout as per your need
client.SetTimeout(1 * time.Minute)


// You can override all below settings and options at request level if you want to
//--------------------------------------------------------------------------------
// Host URL for all request. So you can use relative URL in the request
client.SetHostURL("http://httpbin.org")

// Headers for all request
client.SetHeader("Accept", "application/json")
client.SetHeaders(map[string]string{
        "Content-Type": "application/json",
        "User-Agent": "My custom User Agent String",
      })

// Cookies for all request
client.SetCookie(&http.Cookie{
      Name:"go-resty",
      Value:"This is cookie value",
      Path: "/",
      Domain: "sample.com",
      MaxAge: 36000,
      HttpOnly: true,
      Secure: false,
    })
client.SetCookies(cookies)

// URL query parameters for all request
client.SetQueryParam("user_id", "00001")
client.SetQueryParams(map[string]string{ // sample of those who use this manner
      "api_key": "api-key-here",
      "api_secert": "api-secert",
    })
client.R().SetQueryString("productId=232&template=fresh-sample&cat=resty&source=google&kw=buy a lot more")

// Form data for all request. Typically used with POST and PUT
client.SetFormData(map[string]string{
    "access_token": "BC594900-518B-4F7E-AC75-BD37F019E08F",
  })

// Basic Auth for all request
client.SetBasicAuth("myuser", "mypass")

// Bearer Auth Token for all request
client.SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F")

// Enabling Content length value for all request
client.SetContentLength(true)

// Registering global Error object structure for JSON/XML request
client.SetError(&Error{})    // or resty.SetError(Error{})
Unix Socket
unixSocket := "/var/run/my_socket.sock"

// Create a Go's http.Transport so we can set it in resty.
transport := http.Transport{
	Dial: func(_, _ string) (net.Conn, error) {
		return net.Dial("unix", unixSocket)
	},
}

// Create a Resty Client
client := resty.New()

// Set the previous transport that we created, set the scheme of the communication to the
// socket and set the unixSocket as the HostURL.
client.SetTransport(&transport).SetScheme("http").SetHostURL(unixSocket)

// No need to write the host's URL on the request, just the path.
client.R().Get("/index.html")
Bazel support
Resty can be built, tested and depended upon via Bazel. For example, to run all tests:

bazel test :go_default_test
Mocking http requests using httpmock library
In order to mock the http requests when testing your application you could use the httpmock library.

When using the default resty client, you should pass the client to the library as follow:

// Create a Resty Client
client := resty.New()

// Get the underlying HTTP Client and set it to Mock
httpmock.ActivateNonDefault(client.GetClient())
More detailed example of mocking resty http requests using ginko could be found here