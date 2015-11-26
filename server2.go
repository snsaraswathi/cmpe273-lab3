package main

import (
"fmt"
"github.com/julienschmidt/httprouter"
"net/http"
"encoding/json"
"strconv"
)

var keyValMap map[int]string

type Response struct
{
    Key int `json: key`
    Value string `json: value`
}

type AllKeysResponse struct 
{
         Values []Response
}

func main(){
                    keyValMap = make(map[int]string)
                    mux := httprouter.New()
                    mux.GET("/keys", handleGetAll)
                    mux.GET("/keys/:key_id", handleGet)
                    mux.PUT("/keys/:key_id/:value", handlePut)

                    server := http.Server{
                        Addr: "0.0.0.0:3001",
                        Handler: mux,
                    }
                    server.ListenAndServe()
            }

func handlePut(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
        keyString:=p.ByName("key_id")
        key, err := strconv.Atoi(keyString)
        if err!=nil{
			fmt.Println("Error occured when converting to string")
        }
        value:=p.ByName("value")
       // fmt.Println("the obtained key , value pair:",key,value)
        keyValMap[key]=value
	rw.WriteHeader(http.StatusOK)
}


func handleGet(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	keyString:=p.ByName("key_id")
	key, err := strconv.Atoi(keyString)
        if err!=nil{
			fmt.Println("Error occured when converting to string")
        }

        if _,err1  := keyValMap[key]; !err1 {
                    fmt.Println("Id does not exist")
               }
               value:=keyValMap[key]

	resp:= Response{}
	resp.Key=key
	resp.Value=value

           outputJson, err2 := json.Marshal(resp)
           if err2!=nil{
            fmt.Print("Marshal error")
        }
        rw.Header().Set("Content-Type","application/json")
        rw.WriteHeader(http.StatusOK)
        fmt.Fprintf(rw, "%s", outputJson)
	}


func handleGetAll(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	var Values []Response
	for key,value := range keyValMap {
		var temp Response
		temp.Key=key
		temp.Value=value
		Values=append(Values,temp)
	}
	AllKeys:=AllKeysResponse{Values}
           outputJson, err := json.Marshal(AllKeys)
           if err!=nil{
            fmt.Print("Marshal error")
        }
        rw.Header().Set("Content-Type","application/json")
        rw.WriteHeader(http.StatusOK)
        fmt.Fprintf(rw, "%s", outputJson)
}
