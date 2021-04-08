## Steps to package and run the server in port 8080
* Install the latest golang
* Check out this repository
* Clear the "bin" folder
* Set the current path to the "GOPATH" environment variable
```bash
export GOPATH=<PWD>
```
* Run the following command to package it
```bash
go install firstgolang
```
* check the "bin" folder, an executable file called "firstgolang" will be created
* Move to "bin" folder and run the file
```bash
cd bin
./firstgolang
```
* If you don't see the cursor, means server started and listening in port 8080

## Testing with CURL

```bash
curl -H "Content-Type: text/plain" http://localhost:8080/uppercase -d incredible


INCREDIBLE
```

```bash
curl -H "Content-Type: text/plain" http://localhost:8080/exclaimer -d incredible

{"T1":"wow incredible !!!!.","T2":20}
```

```bash
curl -H "Content-Type: text/plain" http://localhost:8080/uppercase,exclaimer -d incredible

{"T1":"wow INCREDIBLE !!!!.","T2":20}
```

```bash
curl -H "Content-Type: text/plain" http://localhost:8080/exclaimer,uppercase -d incredible

{"T1":"WOW INCREDIBLE !!!!.","T2":20}
```




