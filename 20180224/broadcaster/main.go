
	// Initialize the server and listen all clients
	sv := broadcaster.NewServer(context.Background(), handlers)
	go sv.Listen()

	http.Handle(*endpoint, websocket.Handler(func(ws *websocket.Conn) {
        // Add new client to server and listen ws connection
		sv.NewClient(&Conn{ws: ws}).Listen()
	}))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		panic(err)
	}
