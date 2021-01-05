package main

import (
	"flag"
	"fmt"
	"gopirate.com/instances"
	"gopirate.com/providers"
	"time"
)

func main() {
	var input string
	p := providers.Provider{}
	machines := instances.Instance{}

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
			p.ConnectAndLunch(*pn, *pa, *ps, *pr, *ic)
			fmt.Println("initializing ...")
			fmt.Println(time.Now())
			//time.Sleep(3 * time.Minute)
			fmt.Println(time.Now(), "you spend 3 minuties staring to the screen")
			p.InstanceStaus()
			machines.IP = p.ListIPs(*ic)
		}

		instances.CreateTunnel(machines)

		//Randomly select the provider
		if *rp == true {
		}

		fmt.Println("Do you wanna to flush? press y")
		fmt.Scan(&input)
		if input == "y" {
			p.Flush()
		}
	}
}
