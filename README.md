# flash
Flash is highly optimised typeahead(search recommendation) server. We are building it with an aim to learn various tradeoffs/practices.
While building it, we will use different data strtcures/storage devices/protocols/deployment strategies and compare them from different aspects.

### To build server:
 `go run src/server.go`

### Test latency:
`curl -w "@curl-format.txt" -s "localhost:8080/isAlive"`

> for healthcheck route in localhost, I am getting response ~1-2ms.

## Handy commands
- `go run ./...` - to run server
- `lsof -i:[PORT]` - to know PID of proccess running on PORT.
- `curl -w "@curl-format.txt" -s [URL]` - to know response time of your request with curl
- `sudo pmap [PID]` - check total memory used by PID

## Design perspective
1. ### Expection with the server?
   keeping in mind the fact that we are designing for Humans. Average time at which humans respond is 200ms( interval b/w consecutive key stokes), So we need to keep this in mind and keep our latency < 200ms always and..... Amazon search recommendation api calls responds in around ~100ms.
      
2. ### Where our we storing our dataset (strings which server will recommend)?

    Lookup time for Memory, SSD, and HDD is 100ns, 150us, 10 ms respective.
    It makes sense for us to keep all data in memory itself for low latency. We can do sharding on range bases, if data grows more than the size of 1 node. further we can use disk to take backup by taking periodic snapshot of data, for fault tolerance on poweroff. Later, we can also try to use SSD for storage and see how trade off b/w cost vs latency works.


2. ### Which Data structure to choose for hold strings in memory?

    Our usecase is to fetch popular strings with same prefix. Trie looks intuitive here, but I feel we could perform better with hashmap as with hashmap it will be easy for us to shard, Maps are native for memory, and for every prefix as key, we can use a priority queue as value which contains all popular strings with same prefix as key. Doing this with trie will be little harder. But it's going to be interesting to compare both DS. Making an algorithm to scale them.   


3. ### Are we taking care of protocol too?
    yes, Idea is to make api call to server with every key stroke in search box and return all strings with same prefix as response, means server will be in heavy read load. 
    We can use websockets to faster resonse from server and not setting up TCP connection for every keystroke but I feel using ws is going to be overkill as we don't need duplex connection. We can start with HTTP and do some optimisations like adding "Connection: keep-alive" header to tell server to not close connection for subsequent connections.

4. ### Deployment Strategies.
    As latency is most important factor, We can try to deploy in various regions and using a load balancer to route to closes geolocation server. We need low latency, with high avaibility and less or eventual consistency that means it's fine if users in different region see different suggestions for their prefix. I feel deployment is dominant part in keeping latency low, We can do benchmarking after adding every new component(load balancer, sharding) to understand their affects.
    
5. ### Why Go?
   Keeping perfomance in mind, I chose GO over any imprative language. Rust might be even better choice, but as I don't have enough knowledge with Rust rigth now so may be We an migrate it later.

6. ### Lot to add.

