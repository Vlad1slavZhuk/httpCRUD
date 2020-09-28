# httpCRUD - Task 2

`http://localhost:[port]/` - Show text - "Hello World!".  
`http://localhost:[port]/car` - Loads the page of form `add.html`. After filling out the form, add to the list of cars.

### GET
`http://localhost:[port]/car/add` - Decode `JSON` format and add new car in lists of cars.  
`http://localhost:[port]/cars` - Show in format `JSON` list of cars.  
`http://localhost:[port]/cars/{id:[0-9]+}` - Show in format `JSON` information about the car by `id`. 
### POST
`http://localhost:[port]/car/add` - Take data from the `add.html` form and add it to the list of cars.  

### PUT / PATCH
`http://localhost:[port]cars/{id:[0-9]+}` - Updates vehicle data.

### DELETE
`http://localhost:[port]/cars/{id:[0-9]+}` - Deleted car for `id` car.

---
### Postman - How to send request 
Send [**Any Protocol HTTP**] with Body - raw `JSON` format.  
Example:

``` json5
{
    model: "Volvo",
    color: "red",
    price: 50000.12
}
```
---