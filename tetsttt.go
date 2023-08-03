package main

import (
	"fmt"
	"github.com/scrapli/scrapligo/driver/options"
	"github.com/scrapli/scrapligo/platform"
)

func main() {
	p, err := platform.NewPlatform(
		// cisco_iosxe refers to the included cisco iosxe platform definition
		"",
		"10.80.10.59",
		options.WithAuthNoStrictKey(),
		options.WithAuthUsername("admin"),
		options.WithAuthPassword("Linklogis!@#$0123"),
	)
	if err != nil {
		fmt.Printf("failed to create platform; error: %+v\n", err)

		return
	}

	d, err := p.GetNetworkDriver()
	if err != nil {
		fmt.Printf("failed to fetch network driver from the platform; error: %+v\n", err)

		return
	}

	err = d.Open()
	if err != nil {
		fmt.Printf("failed to open driver; error: %+v\n", err)

		return
	}

	defer d.Close()

	r, err := d.SendCommand("show version")
	if err != nil {
		fmt.Printf("failed to send command; error: %+v\n", err)
		return
	}

	fmt.Printf(
		"sent command '%s', output received (SendCommand):\n %s\n\n\n",
		r.Input,
		r.Result,
	)
}
