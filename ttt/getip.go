package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/op/go-logging"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

var (
	log    = logging.MustGetLogger("iptables")
	format = logging.MustStringFormatter(
		"%{color}%{time:15:04:05.000} %{shortfunc} >> %{level:.4s} %{id:03x}%{color:reset} %{message}")
)

func init() {
	backend := logging.NewLogBackend(os.Stdout, "", 0)
	backendFormat := logging.NewBackendFormatter(backend, format)
	backendLevel := logging.AddModuleLevel(backend)
	backendLevel.SetLevel(logging.DEBUG, "")
	logging.SetBackend(backendFormat, backendLevel)
}

func addWhiteList(ip string) string {
	args := []string{"/sbin/iptables", "-I", "INPUT", "-s", ip, "-p", "tcp", "--dport", "3879", "-j", "ACCEPT"}
	cmd := exec.Command("sudo", args...)
	err := cmd.Run()
	if err != nil {
		panic(err.Error())
	}
	return strings.Join(args, " ")
}

func getDeleteList() []int {
	var ret []int
	args := []string{
		"/sbin/iptables", "-nL", "--line-num",
	}
	ipt := exec.Command("sudo", args...)
	filters := exec.Command("grep", "3879")
	iptStdout, _ := ipt.StdoutPipe()
	if err := ipt.Start(); err != nil {
		log.Errorf("Error: Can't start ipt %s\n", err)
		return ret
	}
	iptOutputBuf := bufio.NewReader(iptStdout)
	filterStdin, _ := filters.StdinPipe()
	iptOutputBuf.WriteTo(filterStdin)
	var outPutbuf bytes.Buffer
	filters.Stdout = &outPutbuf
	if err := filters.Start(); err != nil {
		log.Errorf("Error: filters Can't start %s\n", err)
	}
	if err := filterStdin.Close(); err != nil {
		log.Errorf("Error: filters stdin Can't close %s\n", err)
		return ret
	}
	if err := filters.Wait(); err != nil {
		log.Errorf("Error: Can't wait filter %s\n", err)
		return ret
	}

	bs := outPutbuf.Bytes()
	//fmt.Printf("pipe: %s\n", bs)
	s := strings.Split(string(bs), "\n")
	for _, v := range s {
		if v != "" {
			s2 := strings.Fields(v)
			if ruleID, err := strconv.Atoi(s2[0]); err != nil {
				log.Errorf("convert string to int error!! %s", err)
				continue
			} else {
				ret = append(ret, ruleID)
			}
		}
	}
	return ret
}

func manager(s []int) {
	//sort slice
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	var DeleteCmd [][]string
	for i := 0; i < len(s); i++ {
		cmd := []string{"/sbin/iptables", "-D", "INPUT", strconv.Itoa(s[i])}
		DeleteCmd = append(DeleteCmd, cmd)
	}
	log.Debugf("%v", DeleteCmd)
	// delete rule
	for _, item := range DeleteCmd {
		log.Debugf("delete: %v", item)
		if _, err := exec.Command("sudo", item...).CombinedOutput(); err != nil {
			log.Errorf("delete Error: %s", err)
		}
	}
	// default Drop rule
	log.Info("Append Default Drop rule...")
	DefautlRule := []string{
		"/sbin/iptables", "-A", "INPUT", "-p", "tcp", "--dport", "3879", "-j", "DROP"}
	if _, err := exec.Command("sudo", DefautlRule...).CombinedOutput(); err != nil {
		log.Errorf("%s", err.Error())
	}
	// Default whitelist
	DefautlWhiteList := []string{"192.168.8.35", "192.168.8.91"}
	//bufio.NewScanner()
	for _, item := range DefautlWhiteList {
		log.Infof("Insert %s to WhiteList!", item)
		addWhiteList(item)
	}
}

func main() {
	//args := WhiteList("192.168.3.22")
	//cmd := exec.Command("sudo" ,args...)
	//b, _ := cmd.Output()
	//fmt.Println(string(b))
	s := getDeleteList()
	manager(s)
	fmt.Println(s)
	//cmd1 :=addWhiteList("192.168.8.90")
	//cmd2 := addWhiteList("192.168.8.91")
	//fmt.Println(cmd1,cmd2)
}
