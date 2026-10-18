package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"bufio"
	"flag"
	"math/rand/v2"
	cb "github.com/tiagomelo/go-clipboard/clipboard"
	"github.com/harsiz/MessyClipboard/internal/randutil"
	
)

type msgs struct {
	m []string
	verb bool // verbose (show stdout)
	devm bool
}

func main() {
	// VAR
	var xm msgs

	// FLAG - verbose
	verbose := flag.Bool(
		"verbose",
		false,
		"Display print messages (if stdout set to terminal)",
	)
	dev := flag.Bool(
		"dev",
		false,
		"Faster Clipboard Time",
	)
	
	flag.Parse()

	// Clipfile 
	
	if *verbose {
		xm.verb = true
		fmt.Println("Verbose On ✅")
	}
	if *dev {
		xm.devm = true
		if *verbose {
			fmt.Println("Developer Mode On ✅")
		}
	}

	clipfile, err := os.Open("clipboard-messages.txt")
	
	if err != nil {
		log.Fatal(err)
	}
	defer clipfile.Close()

	scanner := bufio.NewScanner(clipfile)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		xm.m = append(xm.m, scanner.Text())
		
		if xm.verb == true {
			fmt.Println(scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Main Loop
	if len(xm.m) < 1 {
		log.Fatal("Nothing in clipboard-messages.txt! Add something then re-rerun!")
	}
	go func() {
		c := cb.New()
		for {
			if xm.devm {
				time.Sleep(15 * time.Second)
				if xm.verb {
					fmt.Printf("Next Cycle: ✅\n")
				}
			} else {
				time.Sleep(time.Duration(rand.IntN(20) + 1) * time.Minute)
			}
			ran, t := randutil.Choice(xm.m)
			if t == false {
				log.Fatal("clipboard-messages.txt has nothing in it!")
			}
			if err := c.CopyText(ran); err != nil {
				log.Fatal(err)
			}
			
		}
	}()

	go func() {
		for {
			time.Sleep(30 * time.Second)
			if xm.verb {
				fmt.Printf("Heartbeat: Clipboard Still Working ✅\n")
			}
		}
	}()
	
	select {}
}