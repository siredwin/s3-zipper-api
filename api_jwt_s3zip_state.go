package main

import (
	"github.com/labstack/echo"
	"net/http"
	"io/ioutil"
	"log"
	"strings"
	"github.com/labstack/echo-contrib/session"
)

func s3JwtState(c echo.Context) (err error){
	/**************** GET TOKEN FROM COOKIE ***************************/
	cookie, err := c.Cookie("newJwtToken")
	if err != nil {
		return err
	}
	var bearer = "Bearer " + cookie.Value

	/**************** GET UUIDS FROM SESSION ***************************/
	sess, _ := session.Get("session", c)
	allBodyUUIDs := sess.Values["allBodyUUIDs"]
	/**************** ACCESS API WITH TOKEN  ***************************/
	client := &http.Client{}
	// create new request with allbody set
	// allbody contains uuids

	req2 , err2:= http.NewRequest("POST", "https://api.s3zipper.com/v1/zipstate",  strings.NewReader(allBodyUUIDs.(string)))
	if err2 != nil {
		log.Fatal("err2: ", err2)
		return err2
	}
	req2.Header.Set("Authorization", bearer)
	req2.Header.Set("Content-Type", "application/json; charset=utf-8")

	// get response
	resp3 ,err3 := client.Do(req2)
	if err3 != nil{
		log.Fatal("err3: ", err3)
		return err3
	}
	defer resp3.Body.Close()
	// read body
	body2, err4:= ioutil.ReadAll(resp3.Body)
	if err4 != nil{
		log.Fatal("err4: ", err3)
		return err4
	}
	return c.String(http.StatusOK, string(body2[:]))

}

