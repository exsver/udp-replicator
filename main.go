package main

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

	srcConn, err := NewSourceListener(config.Source)
	if err != nil {
		Log.Error.Fatal(err)
	}

	dstConn, err := MakeDestinationConnections(config)

	var payload = make([]byte, 9000)

	for {
		size, err := srcConn.Read(payload)
		if err != nil {
			Log.Error.Println(err)
		}
		data := make([]byte, size)
		copy(data, payload[0:size])

		Log.Debug.Printf("Send data: %s", data)

		for i, _ := range dstConn {
			_, _ = dstConn[i].Write(data)
		}
	}
}
