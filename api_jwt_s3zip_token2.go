package main

import (
	"github.com/labstack/echo"
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
	"log"
	"strings"
	"time"
)

func getJWTToken2(c echo.Context) (err error) {
	// initiate the client
	client := &http.Client{}
	// set the form
	form := url.Values{}
	form.Add("userKey", "user-key")
	form.Add("userSecret", "user-secret")
	// create a  new request
	req2 , err2:= http.NewRequest("POST", "https://api.s3zipper.com/gentoken",  strings.NewReader(form.Encode()))
	if err2 != nil {
		log.Fatal("NewRequest: ", err2)
		return err2
	}
	// set form in request
	req2.PostForm = form
	req2.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// get response
	resp3 ,err3 := client.Do(req2)
	if err3 != nil{
		log.Fatal("NewRequest: ", err3)
		return err3
	}
	defer resp3.Body.Close()
	// read the body
	body, err4:= ioutil.ReadAll(resp3.Body)
	if err4 != nil{
		return err4
	}

	/*******************************************************************/
	// Get the token from the body
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

