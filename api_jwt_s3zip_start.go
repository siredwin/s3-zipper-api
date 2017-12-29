package main

import (
	"github.com/labstack/echo"
	"net/http"
	"net/url"
	"io/ioutil"
	"log"
	"strings"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
)

func s3JwtStart(c echo.Context) (err error){
	/**************** GET TOKEN FROM COOKIE ***************************/
	cookie, err := c.Cookie("newJwtToken")
	if err != nil {
		return err
	}

	/**************** ACCESS API WITH TOKEN  ***************************/
	var bearer = "Bearer " + cookie.Value
	client := &http.Client{}
	////////////////////////////////////////////////////////////////////
	form := url.Values{}
	form.Add("awsKey", "aws-key")
	form.Add("awsSecret", "aws-secret")
	form.Add("awsBucket", "bucket-name")
	form.Add("awsRegion", "us-east-1")
	form.Add("awsToken", "")
	form.Add("awsEndpoint", "")
	form.Add("filePaths", "path/to/file/or/folder")
	form.Add("filePaths", "path/to/file/or/folder")// You can add many of these // it is a list
	form.Add("zipTo", "")
	form.Add("resultsEmail", "email@mail.com") // email to send results to
	//////////////////////////////////////////////////////////////////////////////////////////////////
	req2 , err2:= http.NewRequest(echo.POST, "https://api.s3zipper.com/v1/zipstart",  strings.NewReader(form.Encode()))
	if err2 != nil {
		log.Fatal("NewRequest: ", err2)
		return err2
	}

	req2.PostForm = form
	req2.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req2.Header.Set("Authorization", bearer)
	///////////////////////
	resp3 ,err3 := client.Do(req2)
	if err3 != nil{
		log.Fatal("NewRequest: ", err3)
		return err3
	}
	defer resp3.Body.Close()
	//
	body2, err4:= ioutil.ReadAll(resp3.Body)
	if err4 != nil{
		return err4
	}

	/******* body contains json results // save it in session ************/
	// Was unable to save to cookies.
	//No need to parse body here // backend expects json string result for simplicity
	uuidBody := string(body2[:])
	// session starts here
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["allBodyUUIDs"] = uuidBody
	sess.Save(c.Request(), c.Response())

	/****************** Return json *************************************************/
	return c.String(http.StatusOK, string(body2[:]))

}
