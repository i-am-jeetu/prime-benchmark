<!-- TODO : Compile And Produce GitHub Repo -->
<!-- Prepare LinkedIn Post -->

![MIT License][license-shield]


<center>
<img src="https://jonathanmh.com/wp-content/uploads/2018/01/jonathan-gopher.png" alt="drawing" width="150"/>
<!-- <h2 style="color:#00acd7;">GoFast</h3> -->
</center>


A Go and python/flask server to serve primality test (prime or not) of integers ~~(useless, I know right)~~.   

## Results
<!-- - A, B : Denotes the number of requests handles by server and maximum concurrent requests made at any particular time respectively. -->
100000 requests were made with atmost 100 at a given time.

| Language/Framework | Time (in Seconds) |
|-----------------|---------------|
|Python/Flask|91.515|
|Node/Express|30.144|
|Go|5.850 |


## FAQs 
* Let me explain all my work through FAQs.

### Why Primality Test?
Three days into Go language and having heard/read multiple times about Go's speed and how it reduced X's server node count, I wanted to give it a try. So, to test, I thought many tasks such as reading a file to serve Key-Value, word auto-complete suggestion from a file, etc. But decided to go with ~~useless~~ primality test of 64-bit integers to simulate some amount of calculations involved.

### Test Environments
I tested on my Intel i7- 7700 HQ processor, with 8 GB of ram. At the time of testing, a fair amount of applications such as browser, Nautilus, and VSCode was running. 



### How to run test?
> You can test by running one server at a time, along with Apache Benchmark.

- Run Python Server by running the following command in Project's ``Python`` directory
```sh
python3 PrimalityServer.py
```

- Alternatively, Run Go Server by running the following command in Project's Go directory 
```sh
chmod +x server && ./server
```

- Once the server are running, you are ready to run Apache Benchmark.
```sh
ab -n 1000 -c 10 -T "application/json" -p body.txt http://127.0.0.1:8001/prime 
```
Here, 1000 represents total number of requests and 10 represents number of concurrent request at any time. Change them accordingly. 

You may need to install ab by running 
```sh
sudo apt install apache2-utils
```



### Build Go Binary from Source
Binary can be build by command 
```sh
build server.go isPrime.go 
```

### Limitation
By Nature of implementation, the Primality Test is only valid for 64bit integers. Integers larger than that are not supported.


<br/>



I was writing a long article on this, but now have just reduced to tldrs, hope you like the article summarie - 
1) Inclined to GO, because of Nutanix seniors' gopher profile photo. Didnt persue it but went to Rust (coz Mozilla )
3) Got few small PRs merged at Mozilla Neqo. But didnt enjoyed writing in rust.
2) Rust was too restrictive, ownership rules kept bugging me here and there. 
4) Read many times Go solved companies server latency issues, wanted to give it a try
5) Created primality test as a fake task to test.
6) GO is now my GOTO language, will keep bugging my future colleagues to adopt it in projects (if they havent already)
TODO : Learn concurrency patterns in GO (SUGGEST ME SOURCES), include nodejs in the becnhmarking test.



[license-shield]: https://img.shields.io/github/license/othneildrew/Best-README-Template.svg?style=flat-square
