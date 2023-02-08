package bkmkitap

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type LoginResponse struct {
	Status     int    `json:"status"`
	Field      string `json:"field"`
	Counter    int    `json:"counter"`
	StatusText string `json:"statusText"`
}

type User struct {
	c       http.Client
	Cookies []*http.Cookie
}

func Login(mail string, password string) (*User, error) {
	reqBody := fmt.Sprintf("password=%s&vendor=0&security=", password)
	req, err := http.NewRequest("POST", loginEndpoint+mail, bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		return nil, err
	}

	// add headers
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	req.Header.Add("Host", "www.bkmkitap.com")
	req.Header.Add("Origin", "https://www.bkmkitap.com")
	req.Header.Add("Referer", "https://www.bkmkitap.com/uye-giris-sayfasi")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("DNT", "1")
	req.Header.Add("TE", "trailers")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/109.0")
	req.Header.Add("Content-Length", fmt.Sprint(len(reqBody)))

	// do the actual request
	c := http.Client{}
	m := &LoginResponse{}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	// sometimes body doesn't have a valid json at all, delightful!
	_ = json.NewDecoder(resp.Body).Decode(&m)
	return &User{c: c, Cookies: resp.Cookies()}, nil
}

func (user *User) GetFavorites() (*Favorites, error) {
	req, err := http.NewRequest("GET", favoritesEndpoint, nil)
	if err != nil {
		return nil, err
	}

	for _, cookie := range user.Cookies {
		req.AddCookie(cookie)
	}

	resp, err := user.c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	favorites := new(Favorites)
	err = json.NewDecoder(resp.Body).Decode(favorites)
	if err != nil {
		return nil, err
	}

	return favorites, nil
}

func (user *User) GetBasket() (*Basket, error) {
	req, err := http.NewRequest("GET", getBasketEndpoint, nil)
	if err != nil {
		return nil, err
	}

	for _, cookie := range user.Cookies {
		req.AddCookie(cookie)
	}

	resp, err := user.c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	basket := new(Basket)
	err = json.NewDecoder(resp.Body).Decode(basket)
	if err != nil {
		return nil, err
	}

	return basket, nil
}
