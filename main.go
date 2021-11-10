package main

import (
	"flag"
	"log"
	"time"

	playwright "github.com/mxschmitt/playwright-go"
)

func main() {
	// process browser preferences via flags
	// This script uses Firefox by default, i.e. no flags means use firefox
	// -w, --w, -webkit, --webkit all switch the browser to WebKit
	// -c, --c, -chromium, --chromium all switch the browser to Chromium
	var (
		browser playwright.Browser
		err     error
		url     string = "https://www.google.com"
	)
	pw, err := playwright.Run()
	if err != nil {
		handleErr(err, "Failed to start PlayWright!")
	}
	useChromiumC, useChromium := flag.Bool("c", false, "Use Chromium to test frontend"), flag.Bool("chromium", false, "Use Chromium to test frontend")
	useWebKitW, useWebKit := flag.Bool("w", false, "Use WebKit to test frontend"), flag.Bool("webkit", false, "Use WebKit to test frontend")
	flag.Parse()
	if *useChromiumC || *useChromium {
		browser, err = pw.Chromium.Launch()
		if err != nil {
			handleErr(err, "Failed to launch Chromium!")
		}
		log.Println("Launched Chromium successfully!")
	} else if *useWebKitW || *useWebKit {
		browser, err = pw.WebKit.Launch()
		if err != nil {
			handleErr(err, "Failed to launch WebKit!")
		}
		log.Println("Launched WebKit successfully!")
	} else {
		browser, err = pw.Firefox.Launch()
		if err != nil {
			handleErr(err, "Failed to launch Firefox!")
		}
		log.Println("Launched Firefox successfully!")
	}
	// by now browser should be selected
	page, err := browser.NewPage()
	if err != nil {
		handleErr(err, "Failed to create page!")
	}
	log.Println("Created New Page!")
	if _, err := page.Goto(url); err != nil {
		handleErr(err, "Failed to go to "+url+"!")
	}
	log.Println("Reached " + url + "!")
	time.Sleep(time.Second * 5)
	if err := browser.Close(); err != nil {
		handleErr(err, "Failed to close browser!")
	}
	log.Println("Browser closed successfully!")
	if err := pw.Stop(); err != nil {
		handleErr(err, "Failed to stop Playwright!")
	}
	log.Println("Closed Playwright successfully!")
}

func handleErr(err error, msg string) {
	log.Fatalf(msg+" Error: %v", err)
}
