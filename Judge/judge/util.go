package judge

import (
    "./unix"
    "encoding/json"
    "net/http"
    "io/ioutil"
    "fmt"
    "strings"
    "errors"
    "bytes"
)

func containerExist(containerName string) bool {
    cli := unix.NewClient("")
    var client *http.Client         //For editor VSCode recognize
    client = cli
    re, err := client.Get("http://unix.socket/1.0/containers/" + containerName)
    defer re.Body.Close()
    if err != nil {
        return false
    }
    var response map[string]interface{}
    bytes, err := ioutil.ReadAll(re.Body)
    if err != nil {
        return false
    }
    err =  json.Unmarshal(bytes, &response)
    if err != nil {
        return false
    }
    code, ok := response["error_code"]
    if ok && code.(float64) == 404 {
        return false
    }
    return true
}

func containerLaunch(containerName, remoteServer, remoteAlias string) error {
    cli := unix.NewClient("")
    var client *http.Client         //For editor VSCode recognize
    client = cli
    str := `{"name":"%s","architecture":2,"profiles":["default"],"ephemeral":true,"config":{"limits.cpu":"1"},"source":{"type":"image","mode":"pull","server":"%s","certificate":"PEM certificate","alias":"%s"}}`
    post := fmt.Sprintf(str, containerName, remoteServer, remoteAlias)
    re, err := client.Post("http://unix.socket/1.0/containers", "application/json", strings.NewReader(post))
    if err != nil {
        return err
    }
    defer re.Body.Close()
    bytes, err := ioutil.ReadAll(re.Body)
    if err != nil {
        return err
    }
    var response map[string]interface{}
    err = json.Unmarshal(bytes, &response)
    if err != nil {
        return err
    }
    code, ok := response["status_code"]
    if !ok || code.(float64) != 100 {
        return errors.New(response["error"].(string))
    }
    return nil
}

func waitForLaunch(containerName string) {
    for !containerExist(containerName) {
        //Waiting...
    }
}

func containerLaunchAndWait(containerName, remoteServer, remoteAlias string) error {
    if err := containerLaunch(containerName, remoteServer, remoteAlias); err != nil {
        return err
    }
    waitForLaunch(containerName)
    return nil
}

func containerStart(containerName string) error{
    if !containerExist(containerName) {
        return errors.New("Container not exist!")
    }
    cli := unix.NewClient("")
    var client *http.Client         //For editor VSCode recognize
    client = cli
    str := `{
                "action": "start",
                "timeout": 30,          
                "force": true  
            }`
    req,err := http.NewRequest("PUT", "http://unix.socket/1.0/containers/" + containerName + "/state", strings.NewReader(str))
    if err != nil {
        return err
    }
    re, err := client.Do(req)
    if err != nil {
        return err
    }
    defer re.Body.Close()
    bytes, err := ioutil.ReadAll(re.Body)
    if err != nil {
        return err
    }
    var response map[string]interface{}
    err = json.Unmarshal(bytes, &response)
    if err != nil {
        return err
    }
    code, ok := response["status_code"]
    if !ok || code.(float64) != 100 {
        return errors.New(response["error"].(string))
    }
    return nil
}

func containerState(containerName string) (state string, ips []string, err error) {
    if !containerExist(containerName) {
        err = errors.New("Container not exist!")
        return
    }
    cli := unix.NewClient("")
    var client *http.Client         //For editor VSCode recognize
    client = cli
    re, err := client.Get("http://unix.socket/1.0/containers/" + containerName + "/state")
    if err != nil {
        return
    }
    defer re.Body.Close()
    bytes, err := ioutil.ReadAll(re.Body)
    if err != nil {
        return
    }
    var response map[string]interface{}
    err = json.Unmarshal(bytes, &response)
    std_status, ok := response["status_code"]
    if !ok || std_status.(float64) != 200 {
        err = errors.New("HTTP Request Fail")
        return
    }
    metadata := response["metadata"].(map[string]interface{})
    status, ok := metadata["status"]
    if !ok {
        err = errors.New("status missing")
        return
    }
    state, ok = status.(string)
    if !ok {
        err = errors.New("status error")
        return
    }
    ips_raw, ok := metadata["ips"].([]interface{})
    if !ok {
        err = errors.New("ips error")
        return
    }
    for _, ip_raw := range ips_raw {
        ip := ip_raw.(map[string]interface{})
        if ip["protocol"] == "IPV4" {
            if ip["address"].(string) == "127.0.0.1" {
                continue
            }
            ips = append(ips, ip["address"].(string))
        }
    }
    return
}

func waitForStart(containerName string) {
    state, _, _ := containerState(containerName)
    for state != "Running" {
        state, _, _ = containerState(containerName)
    }
}

//Sync
func containerPush(containerName string, filePath string, fileBytes []byte) error {
    cli := unix.NewClient("")
    var client *http.Client
    client = cli
    re, err := client.Post("http://unix.socket/1.0/containers/" + containerName + "/files?path=" + filePath, "multipart/form-data", bytes.NewReader(fileBytes))
    if err != nil {
        return err
    }
    defer re.Body.Close()
    bytes,err := ioutil.ReadAll(re.Body)
    var response map[string]interface{}
    err = json.Unmarshal(bytes, &response)
    std_status, ok := response["status_code"]
    if !ok || std_status.(float64) != 200 {
        err = errors.New("HTTP Request Fail")
        return err
    }
    return nil
}