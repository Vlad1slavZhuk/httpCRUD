# httpCRUD - Task 2

This web server handles requests (GET, POST, PUT, PATCH, DELETE) to the routing API.

### How to start HTTP server:  
The usual way to start:
``` 
go run cmd/main.go
```

#### `http://localhost:[port]/`  
Load page `add.html` that is in the folder `template`.

#### `http://localhost:[port]/car` (GET, POST)
Loads the page of form `add.html`. After filling out the form, add to the list of cars.
##### Request `GET`:
Redirect to `http://localhost:[port]/`  
##### Request `POST`:
Receives data from forms and from JSON and creates a new car sale ad.  
FORM:
Filing form and press button "Create and add a new car"

JSON:
In app `Postman` in tab Body, press raw and write data in `JSON` format. After filing send request POST.
```json5
{
    "model": "Mustang Shelby GT500",
    "color": "Black",
    "price": 120000
}
```

#### `http://localhost:[port]/cars` (GET)  
Return in `JSON` format list cars. In my example:

`http://localhost:8081/cars`
```json5
{
   "1": {
      "model": "Mazda CX-5",
      "color": "Aqua",
      "price": 25000
   },
   "2": {
      "model": "Aston Martin One 77",
      "color": "Space Grey",
      "price": 80000.5
   }
}
```
#### `http://localhost:[port]/cars/{id:[0-9]+}` (GET)
Return in `JSON` format found car. In my example:  

`http://localhost:8081/cars/1`
``` json5
{
   "model": "Mazda CX-5",
   "color": "Aqua",
   "price": 25000
}
```

#### `http://localhost:[port]/cars/{id:[0-9]+}` (PUT)
Updates vehicle data.

#### `http://localhost:[port]/cars/{id:[0-9]+}` (DELETE)
Removes a car from the list by `ID`.

If Success - write text `Delete a car!`.  
Otherwise `Not found or already deleted.`
