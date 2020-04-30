package main

import (
	"log"
	"github.com/joeshaw/envdecode"
)

var conn net.Conn
var reader io.ReadCloser
var (
	authClient *oauth.authClient
	creds *oauth.Credentials
)

/* The strings inside the backtick alongside each field in struct are called tags and are available through 
reflection interface, which is how encdecode knows which variables to look for,
We added a required argument to this package, which indicates that it is an error for any env variables to be missing
*/

func setupTwitterAuth() {
	//inline struct, different from normal type <name> struct 
	var ts struct {
		ConsumerKey 	string `env:"SP_TWITTER_KEY, required"`
		ConsumerSecret	string `env:"SP_TWITTER_SECRET, required"`
		AccessToken 	string `env:"SP_TWITTER_ACCESSTOKEN, required"`
		AccessSecret	string `env:"SP_TWITTER_ACCESSSECRET, required"`
	}

	if err:=envdecode.Decode(&ts); err!=nil{
		log.Fatalln(err)
	}

	creds=&oauth.Credentials{
		Token: ts.AccessToken,
		Secret: ts.AccessSecret,
	}

	authClient= &oauth.Client{
		Credentials oauth.Credentials{
			Token: ts.ConsumerKey,
			Secret: ts.ConsumerSecret,
		},
	}
}

//function to dial and connect to our twitter livestream API
func dial(netw, addr string) (net.Conn, error){

	//we first ensure that the connection is closed before 
	//opening a new connection
	if conn!=nil{
		conn.Close()
		conn=nil
	}

	netc,err:=net.DialTimeout(netw, addr, 5*time.Second)
	if err!=nil {
		return nil,err
	}

	conn=netc
	return netc, nil
}


//function to periodically close the connection and 
//Reader which reads from the body os response
func closeConn() {
	if conn!=nil{
		con.Close()
	}

	if reader!=nil{
		reader.Close()
	}
}

var (
	authSetupOnce sync.Once 
	httpClient *http.Client
)

func make Request(req *http.Request, params url.Values) (*http.Response, error) {
	authSetupOnce.Do(func() {
		setupTwitterAuth()
		httpClient =&http.Client{
			Transport: &http.Transport {
				Dial: dial,
			},
		}
	})

	formEnc:=params.Encode()
	req.Header.Set("Content-Type", "application/x-www-form- urlencoded")
	req.Header.Set("Content-Length", strconv.Itoa(len(formEnc)))
	req.Header.Set("Authorization", authClient.AuthorizationHeader(creds,
		"POST",
		req.URL, params))
		return httpClient.Do(req)
}

