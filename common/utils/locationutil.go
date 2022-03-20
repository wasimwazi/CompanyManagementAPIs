package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

func VerifyLocation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userIP := ReadUserIP(w, r)
		//Send location API request
		requestURL := fmt.Sprintf("https://ipapi.co/%s/country/", userIP)
		resp, err := http.Get(requestURL)
		if err != nil {
			log.Println("Error : Location API error(VerifyLocation) -", err.Error())
			Fail(w, InternalServerErrorCode, err.Error())
			return
		}
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error : Reading responses error(VerifyLocation) -", err.Error())
			Fail(w, InternalServerErrorCode, err.Error())
			return
		}
		if userIP != "127.0.0.1" {
			if string(bytes) != "CY" && string(bytes) != "IN" {
				Fail(w, BadRequestCode, LocationNotAllowedError)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func ReadUserIP(w http.ResponseWriter, r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Println("Error : request IP is not IP:port format -", err.Error())
		return ""
	}
	userIP := net.ParseIP(ip)
	if userIP == nil {
		log.Println("Error : request IP is not IP:port format")
		return ""
	}
	return userIP.String()
}
