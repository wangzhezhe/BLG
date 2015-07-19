package main

import (
	"fmt"
	"net"
	"strings"
)

func localAddresses() {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Print(fmt.Errorf("localAddresses: %+v\n", err.Error()))
		return
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		name := i.Name

		if err != nil {
			fmt.Print(fmt.Errorf("localAddresses: %+v\n", err.Error()))
			continue
		}
		fmt.Println(addrs, name)
		for _, a := range addrs {
			switch v := a.(type) {
			case *net.IPAddr:
				fmt.Printf("%v : %s (%s)\n", i.Name, v, v.IP.DefaultMask())
			}

		}
	}
}

func checkMaster(ifac string) (string, error) {
	ifacobj, err := net.InterfaceByName(ifac)
	if err != nil {
		//fmt.Println(err.Error())
		return "", err
	}

	addrarry, err := ifacobj.Addrs()
	if err != nil {
		//fmt.Println(err.Error())
		return "", err
	}

	fmt.Println(addrarry)
	var masterip = ""
	for _, ip := range addrarry {
		IP := ip.String()
		if strings.Contains(IP, "/24") {
			masterip = strings.TrimSuffix(IP, "/24")
			fmt.Printf("the master ip is : %v \n", masterip)

		}
	}

	return masterip, nil

}

func main() {
	//localAddresses()
	masterip, err := checkMaster("eth0")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(masterip)

}
