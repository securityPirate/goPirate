package main

import (
	"flag"
	"fmt"
	"gopirate.com/connector"
	"gopirate.com/instances"
	"gopirate.com/providers"
	"os"
	"time"
)

func main() {
	var input string
	p := providers.Provider{}
	machines := instances.Fleet{}
	clients := []connector.Client{}

	d := flag.Bool("d", false, "load the defualt configuration")                 //defualt conf
	pn := flag.String("n", "AWS", "Default Cloud Provider")                      //provide name
	pr := flag.String("r", "us-east-2", "Default Region for the Cloud Provider") //provide region
	pa := flag.String("access", "xxxxxx", "Access Code")                         //provide access token
	ps := flag.String("secret", "......", "Secret Code")                         //provide secret token
	ic := flag.Int64("i", 3, "Number of instances")                              //instance counter
	rp := flag.Bool("rand", false, "Select a rondom providers")                  //Random porovider
	pc := flag.Int("p", 1, "Number of providers")                                //provide counter
	flag.Parse()

	if *d == true {
		fmt.Println("initiating with the default Values")
		flag.PrintDefaults()
	} else {
		fmt.Printf("Hunting start using %s in %s with %d instances \n", *pr, *pn, *ic)
		for i := 0; i < *pc; i++ {
			fmt.Println(time.Now())
			fmt.Println("initializing ...")
			p.ConnectAndLunch(*pn, *pa, *ps, *pr, *ic)
			fmt.Println(time.Now(), "you spend 3 minuties staring to the screen")
			//p.InstanceStaus()
			machines.IP = p.ListIPs(*ic)
		}

		auth := connector.Auth{
			Keys: []string{"/home/ahmed/.goPirate/gonhuntKey.pem"},
		}

		for _, v := range machines.IP {

			client, err := connector.NewNativeClient("ec2-user", string(v), "", 22, &auth, nil)
			if err != nil {
				fmt.Printf("Failed to create new client - %s", err)
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

		//Randomly select the provider
		if *rp == true {
		}

		fmt.Println("Do you wanna to flush? press y")
		fmt.Scan(&input)
		if input == "y" {
			p.Flush()
			os.Exit(1)
		}
	}
}
