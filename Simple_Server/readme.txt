This is my first project in Golang . 
It includes making a basic server in Golang.

Important takeaways : 

1. To go to main page use http.FileServer() and http.Dir()

2. To render a static page where no function is called use http.Handle() , if function is called use 
http.HadleFunc().

3. To set the port you can use - 
	srv := &http.Server{
		Addr: "localhost:5000",
	}

4. To any route there are 2 things 
    1-> w http.ResposeWriter [ Response ]
    2-> r *httpRequest [ Request ]

5. To get FormValues we can use r.FormValue()