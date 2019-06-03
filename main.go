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

	var file *os.File
	var err error
	//	var rd *csv.Reader

	if parameters.Append {
		file, err = os.OpenFile(parameters.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		//		rd = csv.NewReader(file)
	} else {
		file, err = os.Create(parameters.Filename)
	}
	defer file.Close()

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

	start = start.UTC()

	fmt.Printf("start: %s\n", start.String())
	stop := start.Add(time.Duration(parameters.Granularity*300) * time.Second)
	//	fmt.Printf("Stop: %s\n", stop.String())

	wr := csv.NewWriter(file)

	wr.Write([]string{"Time", "Volume", "Open", "Close", "High", "Low"})
	wr.Flush()

	now := time.Now().UTC()

	for i := 1; i <= parameters.Iteration; i++ {
		fmt.Printf("%d) Time: %s -> %s", i, start.String(), stop.String())
		if start.After(now) {
			fmt.Printf(" - break -> start: %s\n", start.String())
			break
		}

		if stop.After(now) {
			stop = time.Now()
			fmt.Printf(" - new stop: %s", stop.String())
		}

		params := coinbasepro.GetHistoricRatesParams{
			Start:       start,
			End:         stop,
			Granularity: int(parameters.Granularity),
		}
		ret, err := cli.GetHistoricRates(parameters.Product, params)
		if err != nil {
			fmt.Printf(" - get error: %+v", err.Error())
		}
		fmt.Printf(" - retval = %d", len(ret))
		start = stop
		stop = start.Add(time.Duration(parameters.Granularity*300) * time.Second)
		if err != nil {
			fmt.Printf(" - history error: %s -> %s - %s\n", err.Error(), start.String(), stop.String())
			continue
		}

		for x := range ret {
			var tm string
			if parameters.Timestamp {
				tm = fmt.Sprintf("%d", ret[len(ret)-1-x].Time.Unix())
			} else {
				tm = ret[len(ret)-1-x].Time.Format("2006-01-02 15:04:05")
			}
			str := []string{
				tm,
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
		fmt.Print("\n")
	}

}
