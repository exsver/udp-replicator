package main

import (
	"context"
	"time"
)

func main() {
	// Parse flags
	flags := parseFlags()

	// Set log-level
	setLogger(*flags.LogLevel)

	// Read config
	config, err := getConfiguration(flags)
	if err != nil {
		Log.Error.Fatal(err)
	}

	// Init counters
	counters := NewCounters()

	srcConn, err := NewSourceListener(config.Source)
	if err != nil {
		Log.Error.Fatal(err)
	}

	dstConn, err := MakeDestinationConnections(config)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		ticker := time.NewTicker(time.Duration(config.CountersInterval) * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				Log.Info.Printf("%s", counters.GetStringAndReset())
			case <-ctx.Done():
				return
			}
		}
	}()

	var payload = make([]byte, 9000)

	for {
		size, err := srcConn.Read(payload)
		if err != nil {
			Log.Error.Println(err)
		}
		data := make([]byte, size)
		copy(data, payload[0:size])

		Log.Debug.Printf("Send data: %s", data)

		counters.Datagrams.Add(1)

		for i, _ := range dstConn {
			_, _ = dstConn[i].Write(data)
		}
	}
}
