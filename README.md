# httpCRUD - Task 2

This web server handles requests (GET, POST, PUT, DELETE) to the routing API.

## How to start HTTP server:  
### Makefile:  
`make help` - Show list commands.  
`make build` - Create `server.exe`.  
`make start` - Create and Run `server.exe`. (dependence `build`)  
`make clean` - Remove `server.exe`.  

**Docker**  
`make docker` -  Build and Run image `server` (dependence `docker-build`, `docker-up`)  
`make docker-build` - Build image with name `server`.  
`make docker-up` - Start image `server`.  
`make docker-clean` - Clean image. Run two commands `docker system prune` and `docker rmi -f server`.  

---

#### `http://localhost:[port]/`  
Load page `add.html` that is in the folder `template`.

---
#### `http://localhost:[port]/car` (GET, POST)
Loads the page of form `add.html`. After filling out the form, add to the list of cars.
##### Request `GET`:
Redirect to `http://localhost:[port]/`  
##### Request `POST`:
Receives data from forms and from JSON and creates a new car sale ad.  
`FORM`:  
Filing form and press button "Create and add a new car".  
`JSON`:  
In app `Postman` in tab `Body`, press `raw` and write data in `JSON` format. After filing send request `POST`.
```json5
{
    "brand": "Mustang Shelby",
    "model": "GT500",
    "color": "Black",
    "price": 120000
}
```
---
#### `http://localhost:[port]/cars` (GET)  
Return in `JSON` format list cars. In my example:

`http://localhost:8081/cars`
```json5
{
   "1": {
      "brand": "Mazda",
      "model": "CX-5",
      "color": "Aqua",
      "price": 25000
   },
   "2": {
      "brand": "Aston Martin",
      "model": "One 77",
      "color": "Space Grey",
      "price": 80000.5
   }
}
```
---
#### `http://localhost:[port]/cars/{id:[0-9]+}` (GET)
Return in `JSON` format found car. In my example:  

`http://localhost:8081/cars/1`
```json5
{
   "brand": "Mazda",
   "model": "CX-5",
   "color": "Aqua",
   "price": 25000
}
```
---
#### `http://localhost:[port]/cars/{id:[0-9]+}` (PUT)
Updates vehicle data.  

`JSON`:  
In app `Postman` in tab `Body`, press `raw` and write data in `JSON` format. After filing send request `PUT`.

`localhost:8081/cars/1`
```json5
{
    "brand": "Subaru",
    "model": "Forester",
    "color": "Blue",
    "price": 45000.82
}
```
We updated data:
```json5
{
    "1": {
        "brand": "Subaru",
        "model": "Forester",
        "color": "Blue",
        "price": 45000.82
    },
    "2": {
        "brand": "Aston Martin",
        "model": "One 77",
        "color": "Space Grey",
        "price": 80000.5
    }
}
```
---
#### `http://localhost:[port]/cars/{id:[0-9]+}` (DELETE)
Removes a car from the list by `ID`. In my Example.

`http://localhost:8081/cars/1`  
We delete `ID`= 1. After request `DELETE` we see such a list: 
```json5
{
    "2": {
        "brand": "Aston Martin",
        "model": "One 77",
        "color": "Space Grey",
        "price": 80000.5
    }
}
```