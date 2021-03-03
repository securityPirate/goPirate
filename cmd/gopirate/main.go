package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"gopirate.com/connector"
	"gopirate.com/crypter"
	"gopirate.com/instances"
	"gopirate.com/local"
	"gopirate.com/providers"
)

func main() {

	configDirExist, path := local.Check()
	var logFile, pirateFile *os.File
	logger := log.New(logFile, "", log.LstdFlags|log.Lshortfile)
	p := providers.Provider{}
	machines := instances.Fleet{}
	clients := []connector.Client{}
	var sym crypter.Symmetric
	var config []byte

	pin := flag.String("pin", "", "Encryption Secret Key")                       //defualt conf
	pn := flag.String("n", "AWS", "Default Cloud Provider")                      //provide name
	pr := flag.String("r", "us-east-2", "Default Region for the Cloud Provider") //provide region
	pa := flag.String("access", "", "Access Code")                               //provide access token
	ps := flag.String("secret", "", "Secret Code")                               //provide secret token
	ic := flag.Int64("i", 3, "Number of instances")                              //instance counter
	//rp := flag.Bool("rand", false, "Select a rondom providers")                  //Random porovider
	//pc := flag.Int("p", 1, "Number of providers")                                //provide counter
	flag.Parse()

	if configDirExist {
		//.goPirate dir exist
		//ask user to load last fleet configuration or start from fresh
		//load the old conf file and append to it
		//create a new conf file and archive the other logs for a fresh start
		fmt.Println("Do you wanna to load the last session?")
		if local.Answer() {
			if len(*pin) < 6 || *pin == "" {
				fmt.Print("Please Enter the PIN:")
				*pin = local.Readline()
				//decrypt the config file
				sym.Generate(*pin)
				config = local.ReadEncryptedFile(path+"/pirate.conf", sym)
			} else {
				//decrypt the config file
				sym.Generate(*pin)
				config = local.ReadEncryptedFile(path+"/pirate.conf", sym)
				fmt.Println("loading.........................")
			}
		} else {
			//create a new session
			local.DeleteFile(path + "/pirate.conf")
			local.MoveFile(path+"/pirate.log", path+"/log/"+time.Now().String()+".log")
			fmt.Println("Creating new session ...........")
			logFile, _ = local.OpenFile(path + "/pirate.log")
			pirateFile, _ = local.OpenFile(path + "/pirate.conf")
		}
	} else {
		//create .goPirate dir
		//create a log file
		os.MkdirAll(path+"/log", 0755)
		logFile, _ = local.OpenFile(path + "/pirate.log")
		pirateFile, _ = local.OpenFile(path + "/pirate.conf")
	}
	defer pirateFile.Close()
	defer logFile.Close()

	fmt.Printf("Hunting start using %s in %s with %d instances \n", *pr, *pn, *ic)
	logger.Println("initializing the fleet")

	p.ConnectAndLunch(*pn, *pa, *ps, *pr, *ic, logger , config)

	machines.IP = p.ListIPs(*ic)
	auth := connector.Auth{
		Keys: []string{path + "/" + "keyname"},
	}

	for _, v := range machines.IP {

		client, err := connector.NewNativeClient("ec2-user", string(v), "", 22, &auth, nil)
		if err != nil {
			logger.Printf("Failed to create new client - %s", err)
		}

		clients = append(clients, client)
	}

	for {
		var cmd string
		fmt.Print("$")
		fmt.Scanln(&cmd)
		if cmd == "" || cmd == "exit" {
			break
		}

		for i := range clients {
			clients[i].Shell(cmd)
		}

	}

	fmt.Println("Do you wanna to flush? press y")
	if local.Answer() {
		p.Flush()
		os.Exit(1)
	} else {
		os.Exit(1)
	}

}
