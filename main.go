package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"

	coinbasepro "github.com/preichenberger/go-coinbasepro"
)

func main() {

	flag.Parse()

	cli := coinbasepro.NewClient()

	//	pr, _ := cli.GetProducts()
	//	fmt.Println(pr)
	if parameters.Filename == "" {
		panic("wrong filename")
	}

	file, err := os.Create(parameters.Filename)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// RFC3339 = 2006-01-02T15:04:05Z
	start, err := time.Parse(time.RFC3339, parameters.Start)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Printf("start: %s\n", start.String())
	stop := start.Add(time.Duration(parameters.Granularity*300) * time.Second)
	fmt.Printf("Stop: %s\n", stop.String())

	wr := csv.NewWriter(file)

	wr.Write([]string{"Time", "Volume", "Open", "Close", "High", "Low"})
	wr.Flush()

	for i := 1; i <= parameters.Iteration; i++ {
		fmt.Println(start, stop)
		params := coinbasepro.GetHistoricRatesParams{
			Start:       start,
			End:         stop,
			Granularity: int(parameters.Granularity),
		}
		ret, err := cli.GetHistoricRates(parameters.Product, params)
		start = stop
		stop = start.Add(time.Duration(parameters.Granularity*300) * time.Second)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if stop.After(time.Now()) {
			stop = time.Now()
			fmt.Println("stop: ", stop)
		}

		if start.After(time.Now()) {
			fmt.Println("break: ", start)
			break
		}

		for x := range ret {
			str := []string{
				ret[len(ret)-1-x].Time.String(),
				fmt.Sprintf("%f", ret[len(ret)-1-x].Volume),
				fmt.Sprintf("%f", ret[len(ret)-1-x].Open),
				fmt.Sprintf("%f", ret[len(ret)-1-x].Close),
				fmt.Sprintf("%f", ret[len(ret)-1-x].High),
				fmt.Sprintf("%f", ret[len(ret)-1-x].Low),
			}
			wr.Write(str)
			wr.Flush()
		}
		time.Sleep(time.Second * 1)
	}

}
