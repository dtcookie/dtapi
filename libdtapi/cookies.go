package dtapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
)

var cookieJar = newStickyCookies()

type cookieconfig struct {
	Cookies map[string]*http.Cookie
}

func newCookieConfig() *cookieconfig {
	c := cookieconfig{}
	c.Cookies = make(map[string]*http.Cookie)
	return &c
}

func (config *cookieconfig) toJSON() ([]byte, error) {
	return json.MarshalIndent(config, "", "  ")
}

func (config *cookieconfig) fromJSON(bytes []byte) {
	err := json.Unmarshal(bytes, config)
	if err != nil {
		fmt.Println("[WARNING] " + err.Error())
	}
}

type stickycookies struct {
	http.CookieJar
	Config *cookieconfig
}

func newStickyCookies() *stickycookies {
	c := &stickycookies{}
	c.Config = newCookieConfig()
	return c
}

func (c *stickycookies) load() {
	var cookieFile string
	var err error
	var jsonFile *os.File
	var bytes []byte

	if cookieFile, err = c.getCookieFile(); err != nil {
		fmt.Println("[WARNING] [getCookieFile]" + err.Error())
		return
	}
	if _, err := os.Stat(cookieFile); err == nil {
		if jsonFile, err = os.Open(cookieFile); err != nil {
			fmt.Println("[WARNING] [os.Open]" + err.Error())
			return
		}
		defer jsonFile.Close()
		if bytes, err = ioutil.ReadAll(jsonFile); err != nil {
			fmt.Println("[WARNING] [ioutil.ReadAll]" + err.Error())
			return
		}
		c.Config.fromJSON(bytes)
	}
}

func (c *stickycookies) save() {
	var cookieFile string
	var err error
	var jsonFile *os.File
	var bytes []byte

	if cookieFile, err = c.getCookieFile(); err != nil {
		fmt.Println("[WARNING] " + err.Error())
		return
	}
	if jsonFile, err = os.OpenFile(cookieFile, os.O_TRUNC|os.O_CREATE, os.ModeAppend); err != nil {
		fmt.Println("[WARNING] [os.Open] " + err.Error())
		return
	}
	defer jsonFile.Close()
	if bytes, err = c.Config.toJSON(); err != nil {
		fmt.Println("[WARNING] [toJSON] " + err.Error())
		return
	}
	if _, err := jsonFile.WriteAt(bytes, 0); err != nil {
		fmt.Println("[WARNING] [WriteAt] " + err.Error())
		return
	}
}

func (c *stickycookies) getCookieFile() (string, error) {
	var homedir string
	var err error
	if homedir, err = c.getHomeDir(); err != nil {
		return homedir, err
	}
	return filepath.Join(homedir, ".dtcfcli.cookies.json"), nil
}

func (c *stickycookies) getHomeDir() (string, error) {
	user, err := user.Current()
	return user.HomeDir, err
}

func (c *stickycookies) SetCookies(u *url.URL, cookies []*http.Cookie) {
	for _, cookie := range cookies {
		// fmt.Println("... Set-Cookie: " + cookie.Name + ": " + cookie.Value)
		c.Config.Cookies[cookie.Name] = cookie
	}
	c.save()
}

func (c *stickycookies) Cookies(u *url.URL) []*http.Cookie {
	c.load()
	result := make([]*http.Cookie, 0)
	for _, cookie := range c.Config.Cookies {
		// fmt.Println("... Cookie: " + cookie.Name + ": " + cookie.Value)
		result = append(result, cookie)
	}
	return result
}
