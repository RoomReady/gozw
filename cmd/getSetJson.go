package main

/*
	These routines don't get/set/load json directly from file.  They get them
	from rrcJsonServer which keeps a copy of the most current version of the
	json file.
*/
import (
	"errors"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/Jeffail/gabs"
)

const userFiles = "/data/"

func loadJsonFile(fName string) (*gabs.Container, bool) {
	// loads a json file and returns the gabs container
	// returns empty json and false if it fails
	j := gabs.New()
	err := getJson(userFiles+fName, j)
	if err != nil {
		fmt.Println("Warning  : could not load json file (" + fName + "): " + err.Error())
		j, _ = gabs.ParseJSON([]byte("{}"))
		return j, false
	}
	return j, true
}

func getJsonBool(filePath, jsonPath string) (bool, bool) {
	j := gabs.New()
	err := getJson(filePath, j)
	if err == nil {
		value, ok := j.Path(jsonPath).Data().(bool)
		if ok {
			return value, ok
		} else {
			return false, ok
		}
	}
	return false, false
}

func getJsonString(filePath, jsonPath string) (string, bool) {
	j := gabs.New()
	err := getJson(filePath, j)
	if err == nil {
		value, ok := j.Path(jsonPath).Data().(string)
		if ok {
			return value, ok
		} else {
			return "", ok
		}
	}
	return "", false
}

func setJson(filepath string, jsn *gabs.Container) error {
	// reqJson will be our json request to the json server
	reqJson := gabs.New()
	reqJson.Set("push", "operation")
	reqJson.Set(filepath, "filepath")
	reqJson.Set(jsn.Data(), "data")
	//fmt.Println(reqJson.StringIndent("", "    "))
	//return nil

	servAddr := "localhost:61000"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		fmt.Println("Fatal    : setJson(): could not resolve jsonServer")
		time.Sleep(time.Second * 1)
		os.Exit(2)
		//println("ResolveTCPAddr failed:", err.Error())
		return err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("Fatal    : setJson(): could not connect to jsonServer")
		time.Sleep(time.Second * 1)
		os.Exit(2)
		//println("Dial failed:", err.Error())
		return err
	}
	defer conn.Close()

	conn.Write([]byte(reqJson.String()))

	reply := make([]byte, 32768)
	for {
		n, err := conn.Read(reply)
		if err != nil {
			//println("Write to server failed:", err.Error())
			return err
		}

		strReply := strings.Trim(string(reply[:n]), "\" \t\r\n")
		j, err := gabs.ParseJSON([]byte(strReply))
		if err != nil {
			//fmt.Println("Return: " + j.StringIndent("", "    "))
			fmt.Println("Warning  : setJson(): " + j.S("error").Data().(string))
		}
		return nil
	}
	return errors.New("Unknown error in setJson")
}
func getJson(filepath string, jsn *gabs.Container) error {
	// err := getJson(filepath string, *gabs.Container())

	// clear out the original container
	children, err := jsn.ChildrenMap()
	if err == nil {
		for key, _ := range children {
			jsn.Delete(key)
		}
	}
	//jsn.Set(false, "success")

	// reqJson will be our json request to the json server
	reqJson := gabs.New()
	reqJson.Set("pull", "operation")
	reqJson.Set(filepath, "filepath")

	servAddr := "localhost:61000"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		fmt.Println("Fatal    : getJson(): could not resolve jsonServer")
		time.Sleep(time.Second * 1)
		os.Exit(2)
		//println("ResolveTCPAddr failed:", err.Error())
		return err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("Fatal    : getJson(): could not connect to jsonServer")
		time.Sleep(time.Second * 1)
		os.Exit(2)
		//println("Dial failed:", err.Error())
		return err
	}
	defer conn.Close()

	conn.Write([]byte(reqJson.String()))

	reply := make([]byte, 32768)
	for {
		n, err := conn.Read(reply)
		if err != nil {
			//println("Write to server failed:", err.Error())
			return err
		}

		strReply := strings.Trim(string(reply[:n]), "\" \t\r\n")
		j, err := gabs.ParseJSON([]byte(strReply))
		if err != nil {
			return errors.New("there was an error parsing return json: " + err.Error())
		}

		if j.Exists("success") {
			if j.S("success").Data().(bool) {
				if j.Exists("data") {
					children, _ := j.S("data").ChildrenMap()
					for key, child := range children {
						jsn.Set(child.Data(), key)
					}
					return nil
				} else {
					return errors.New(`No "data" key found in returned json`)
				}
			} else {
				return errors.New(`success == false`)
			}
		} else {
			return errors.New(`No "success" key found in returned json`)
		}

		return nil
	}
}
