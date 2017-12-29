package main

import (
	"github.com/labstack/echo"
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
	"log"
	"time"
)

func getJWTToken(c echo.Context) (err error) {
	resp, err := http.PostForm("https://api.s3zipper.com/gentoken",
		url.Values{
			"userKey": {"user-key"},
			"userSecret": {"user-secret"},
		})

	if err != nil {
		log.Println("errorination happened getting the response", err)
		return err
	}
	defer resp.Body.Close()
	/*******************************************************************/
	body, err := ioutil.ReadAll(resp.Body)
	log.Printf("\nbody -- %v\n", string(body))

	if nil != err {
		log.Println("Error happened reading the body", err)
		return err
	}
	/*******************************************************************/
	var p GetToken
	err = json.Unmarshal(body, &p)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return err
	}
	token := p.Token

	//save token in cookie
	if len(token) > 1{
		//save token
		cookie := new(http.Cookie)
		cookie.Name = "newJwtToken"
		cookie.Value = token
		cookie.Path = "/"
		cookie.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(cookie)
	}else{
		return c.JSON(http.StatusOK, "No token was saved in cookie")
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})

}
