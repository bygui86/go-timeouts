
# Go sample project to handle timeouts

## Run

### JSON Server
```
cd go-timeout/
npm install -g json-server
json-server --watch db.json
curl http://localhost:3000/todos/1
	> {
	>   "userId": 1,
	>   "id": 1,
	>   "title": "delectus aut autem",
	>   "completed": false
	> }
```

### Example
```
go run main.go
```

---

## Links

* https://medium.com/swlh/the-simplest-way-to-handle-timeouts-in-golang-11e371dc6188
* https://jsonplaceholder.typicode.com/todos/1
* https://github.com/typicode/json-server
