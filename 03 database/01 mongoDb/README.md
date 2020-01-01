# Working with MongoDB

## Install MongoDB

Download the MongoDB Community Server:

```
https://www.mongodb.com/download-center/community
```

Start the service by specifying a folder to store the data:

```
mongod --dbpath $HOME/Development/Mongo/
```

## Allgemeine Erkundungen

Übersicht über alle Datenbanken und deren Speichernutzung

```
> show dbs
admin   0.000GB
config  0.000GB
local   0.000GB
```

Wechsel zu einer Datenbank / Neuanlage wenn die Db nicht existiert

```
> use shop
switched to db shop
```

Insert in Tabelle (return wert mit id / lasse ich in zukunft weg!)
Neuanlage falls die Collection nicht existiert
Ausgabe aller Datensätze der Collection / ...und formatiert

```
> db.products.insertOne({name: "Nintendo Switch", price: 309.00, available: true})
{
        "acknowledged" : true,
        "insertedId" : ObjectId("5d4eb018262fc6a58aaaa1fe")
}
> db.products.find()
{ "_id" : ObjectId("5d4eb018262fc6a58aaaa1fe"), "name" : "Nintendo Switch", "price" : 309, "available" : true }
> db.products.find().pretty()
{
        "_id" : ObjectId("5d4eb018262fc6a58aaaa1fe"),
        "name" : "Nintendo Switch",
        "price" : 309,
        "available" : true
}
```

Insert zweiter Datensatz mit zusätzlichem Attribut

```
> db.products.insertOne({name: "Mario Kart 8 Deluxe", price: 59.99, available: true, multiplayer: true})
> db.products.find().pretty()
{
        "_id" : ObjectId("5d4eb018262fc6a58aaaa1fe"),
        "name" : "Nintendo Switch",
        "price" : 309,
        "available" : true
}
{
        "_id" : ObjectId("5d4eb295262fc6a58aaaa1ff"),
        "name" : "Mario Kart 8 Deluxe",
        "price" : 59.99,
        "available" : true,
        "multiplayer" : true
}
```

Insert mit embedded Document

```
> db.products.insertOne({name: "Fender Mustang Bass", price: 555.00, available: true, details: {material: "Maple", scale: 762}})
> db.products.find().pretty()
...
{
        "_id" : ObjectId("5d4eb7fb262fc6a58aaaa200"),
        "name" : "Fender Mustang Bass",
        "price" : 555,
        "available" : true,
        "details" : {
                "material" : "Maple",
                "scale" : 762
        }
}
``` 
