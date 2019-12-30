# Find Commands

## Preparation 

```
db.passengers.insertMany([
    {"name": "Max Schwarzmueller", "age": 29},
    {"name": "Manu Lorenz", "age": 30},
    {"name": "Chris Hayton", "age": 35},
    {"name": "Sandeep Kumar", "age": 28},
    {"name": "Maria Jones", "age": 30},
    {"name": "Alexandra Maier", "age": 27},
    {"name": "Dr. Phil Evans", "age": 47},
    {"name": "Sandra Brugge", "age": 33},
    {"name": "Elisabeth Mayr", "age": 29},
    {"name": "Frank Cube", "age": 41},
    {"name": "Karandeep Alun", "age": 48},
    {"name": "Michaela Drayer", "age": 39},
    {"name": "Bernd Hoftstadt", "age": 22},
    {"name": "Scott Tolib", "age": 44},
    {"name": "Freddy Melver", "age": 41},
    {"name": "Alexis Bohed", "age": 35},
    {"name": "Melanie Palace", "age": 27},
    {"name": "Armin Glutch", "age": 35},
    {"name": "Klaus Arber", "age": 53},
    {"name": "Albert Twostone", "age": 68},
    {"name": "Gordon Black", "age": 38}
])
```

## Find

Cursor...

```
db.passengers.find().pretty()
```

...

```
Type "it" for more
> it
```

```
db.passengers.find().toArray()
```

```
db.passengers.find().forEach((passengerData) => {printjson(passengerData)})
```

## FindOne

Daten
Deswegen kein pretty!

## Projection

```
db.passengers.find({}, {name: 1})
{ "_id" : ObjectId("5d505797f533571837f5e32c"), "name" : "Max Schwarzmueller" }
{ "_id" : ObjectId("5d505797f533571837f5e32d"), "name" : "Manu Lorenz" }
{ "_id" : ObjectId("5d505797f533571837f5e32e"), "name" : "Chris Hayton" }
...
```

```
db.passengers.find({}, {name: 1, _id: 0})
{ "name" : "Max Schwarzmueller" }
{ "name" : "Manu Lorenz" }
{ "name" : "Chris Hayton" }
```
