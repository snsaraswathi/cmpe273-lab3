package main

import (
"fmt"
"hash/fnv"
"container/ring"
"strconv"
"net/http"
"io/ioutil"
"encoding/json"
)

type Response struct
{
    Key int `json: key`
    Value string `json: value`
}

type AllKeysResponse struct
{
        Values []Response
}
	var server1 string
        var server2 string
        var server3 string

func main(){
		
	var serverArr [13]string		//13 is the total number of hash slots	
	server1="http://localhost:3000"
	server2="http://localhost:3001"
	server3="http://localhost:3002"

	//Hash the server
	serverIndex1:=serverHash(server1)	
	serverIndex2:=serverHash(server2)
	serverIndex3:=serverHash(server3)

	serverArr[serverIndex1]=server1
	serverArr[serverIndex2]=server2
	serverArr[serverIndex3]=server3

	key:=[]int{1,2,3,4,5,6,7,8,9,10}
	valueInput:=[]string{"a","b","c","d","e","f","g","h","i","j"}

	//create a ring array 
	ringArr := ring.New(len(serverArr))
	for i := 0; i < ringArr.Len(); i++ {
		ringArr.Value = serverArr[i]
		ringArr = ringArr.Next()
	}

	p:=ringArr
	for i:=0;i<len(key);i++{
	var hashVal uint32
	hashVal = hash(strconv.Itoa(key[i]))

	a := int(hashVal)
	index:=a%(len(serverArr))

	for j:=0;j<index;j++{
		ringArr = ringArr.Next()
	}
	for ringArr.Value=="" {
		ringArr = ringArr.Next()
	}
	serverName,_:=(ringArr.Value).(string) 
	url:=serverName+"/keys/"+strconv.Itoa(key[i])+"/"+valueInput[i]
	//Put request 
	req1, errReq := http.NewRequest("PUT", url, nil)
    if errReq!=nil{
        fmt.Println("Put request error")
        return
    }

    client := &http.Client{}
    resp, err5 := client.Do(req1)
    if err5 != nil {
        fmt.Println("Put request error")
        return
    }
    defer resp.Body.Close()
	ringArr=p
}
//GETTING the values
	for i:=0;i<len(key);i++{
	var hVal uint32
	hVal = hash(strconv.Itoa(key[i]))

	a := int(hVal)
	index:=a%(len(serverArr))

	for j:=0;j<index;j++{
		ringArr = ringArr.Next()
	}

	for ringArr.Value=="" {
				ringArr = ringArr.Next()
			}

	serverName,_:=(ringArr.Value).(string) 

	url:=serverName+"/keys/"+strconv.Itoa(key[i])
	resp, err := http.Get(url);
                    if err != nil {
                        fmt.Println("Get error")
                    }

                    defer resp.Body.Close()
                    body, err1 := ioutil.ReadAll(resp.Body)
                    if err1 != nil {
                        fmt.Println("Get error")

                    }
                    var result Response 
                    err2:=json.Unmarshal(body,&result)
                    if err2 != nil {
                       fmt.Println("Get error")
                    }

		fmt.Println(result.Key,"=>",result.Value)
			ringArr=p
	}

	GetAllServer1()
	GetAllServer2()
	GetAllServer3()
}

func hash(s string) uint32 {
        h := fnv.New32a()
        h.Write([]byte(s))
        return h.Sum32()
}

func serverHash(url string) int {
 hashVal:=hash(url)
 index:=(hashVal*12345)%13
 return int(index)

}

//get all values
func GetAllServer1(){
		url1:= server1+"/keys/"
                        resp, err := http.Get(url1);
                    if err != nil {
                        fmt.Println("Get all error")
                    }   

                    defer resp.Body.Close()
                    body, err1 := ioutil.ReadAll(resp.Body)
                    if err1 != nil {
                        fmt.Println("Get error")

                    } 

                    var result1 AllKeysResponse 
                    err2:=json.Unmarshal(body,&result1)
                    if err2 != nil {
                       fmt.Println("Get error")
                    }   
			
			fmt.Println("All Key values stored in server1") 
                    for l:=0;l<len(result1.Values);l++{
                        temp:=result1.Values[l]
                                fmt.Println(temp.Key,"=>",temp.Value)
                        }
}  

func GetAllServer2(){
                server2Url:= "http://localhost:3001"+"/keys/"
                    resp, err := http.Get(server2Url);
                    if err != nil {
                        fmt.Println("Get all error")
                    }   

                    defer resp.Body.Close()
                    body, err1 := ioutil.ReadAll(resp.Body)
                    if err1 != nil {
                        fmt.Println("Get error")

                    }   
                    var getAllResult AllKeysResponse 
                    err2:=json.Unmarshal(body,&getAllResult)
                    if err2 != nil {
                       fmt.Println("Get error")
                    }   

                        fmt.Println("All Key values stored in server2")
                    for l:=0;l<len(getAllResult.Values);l++{
                        temp:=getAllResult.Values[l]
                                fmt.Println(temp.Key,"=>",temp.Value)
                        }   
}

func GetAllServer3(){
                    server3Url:= "http://localhost:3002"+"/keys/"
                    resp, err := http.Get(server3Url);
                    if err != nil {
                        fmt.Println("Get all error")
                    }   

                    defer resp.Body.Close()
                    body, err1 := ioutil.ReadAll(resp.Body)
                    if err1 != nil {
                        fmt.Println("Get error")

                    }   
                    var getAllResult AllKeysResponse 
                    err2:=json.Unmarshal(body,&getAllResult)
                    if err2 != nil {
                       fmt.Println("Get error")
                    }   

                        fmt.Println("All Key values stored in server3")
                    for l:=0;l<len(getAllResult.Values);l++{
                        temp:=getAllResult.Values[l]
                                fmt.Println(temp.Key,"=>",temp.Value)
                        }   
} 
