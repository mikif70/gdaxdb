package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	coinbasepro "github.com/preichenberger/go-coinbasepro"
)

func main() {
	cli := coinbasepro.NewClient()

	//	pr, _ := cli.GetProducts()
	//	fmt.Println(pr)

	file, err := os.Create(os.Args[1])
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	start := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)
	stop := start.Add(time.Hour * 24 * 10)
	wr := csv.NewWriter(file)

	wr.Write([]string{"Time", "Volume", "Open", "Close", "High", "Low"})
	wr.Flush()

	for i := 1; i < 90; i++ {
		fmt.Println(start, stop)
		params := coinbasepro.GetHistoricRatesParams{
			Start:       start,
			End:         stop,
			Granularity: 3600,
		}
		ret, err := cli.GetHistoricRates("ETC-EUR", params)
		start = stop
		stop = stop.Add(time.Hour * 24 * 10)
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
		time.Sleep(time.Second * 2)
	}

}
