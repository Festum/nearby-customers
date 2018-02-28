# Assignment: Customer Records

We have some customer records in a text file `customers.json` -- one customer per line, JSON-encoded. We want to invite any customer within 100km of our Dublin office for some food and drinks on us. Write a program that will
- read the full list of customers and
- output the names and user ids of matching customers (within 100km),
- sorted by User ID (ascending).
You can use the first formula from this Wikipedia article to calculate distance. Don't forget, you'll need to convert degrees to radians.
The GPS coordinates for our Dublin office are 53.339428, -6.257664.
You can find the Customer list here.
⭑ Please don’t forget, your code should be production ready, clean and tested!

## Usage

Direct pull this project and run will initial a web server for this API:
```
git clone nearby-customers
cd nearby-customers/
go get ./...
go run main.go
```

Hit GET http://localhost:8080/nearcustomers

## Test

```
go test
```
