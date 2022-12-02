This is my first project in Golang . 
It includes making a basic server in Golang.
Important takeaways : <br/>

1. To go to main page use http.FileServer() and http.Dir()

2. To render a static page where no function is called use http.Handle() , if function is called use 
http.HadleFunc().

3. To set the port you can use - <br/>
	<i>srv := &http.Server{<br/>
		Addr: "localhost:5000",<br/>
	}</i>

4. To any route there are 2 things 
    1-> w http.ResposeWriter [ Response ]
    2-> r *httpRequest [ Request ]

5. To get FormValues we can use r.FormValue()

<br/>
<br/>
TO RUN THE CODE PLEASE CLONE THE FILE AND COMPLETE THE GOLANG SETUP. AFTER THAT OPEN TERMINAL AND ENTER<br/>
-> go run main.go<br/><br/>
Following are the routes<br/>
-localhost:5000<br/>
-localhost:5000/hello<br/>
-localhost:5000/form.html<br/>

Enter them manually on-by-one in your browser to check upon all of them.

<b>CONGRATULATIONS ON DOING THE FIRST GOLANG PROJECT!!!</b>