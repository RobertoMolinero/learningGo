# CRUD

## Create

```
use flights
db.flightData.insertOne({"departureAirport": "MUC", "arrivalAirport": "SFO", "aircraft": "Airbus A380", "distance": 12000, "intercontinental": true})
db.flightData.insertOne({departureAirport: "TXL", arrivalAirport: "LHR"})
db.flightData.insertOne({departureAirport: "TXL", arrivalAirport: "LHR", _id: "txl-lhr-1"})
```

## Read

```
db.flightData.find().pretty()
{
        "_id" : ObjectId("5d503da6f533571837f5e328"),
        "departureAirport" : "MUC",
        "arrivalAirport" : "SFO",
        "aircraft" : "Airbus A380",
        "distance" : 12000,
        "intercontinental" : true
}
{
        "_id" : ObjectId("5d503efef533571837f5e329"),
        "departureAirport" : "TXL",
        "arrivalAirport" : "LHR"
}
{
        "_id" : "txl-lhr-1",
        "departureAirport" : "TXL",
        "arrivalAirport" : "LHR"
}
```

```
db.flightData.find({distance: {$gt: 10000}})
{ "_id" : ObjectId("5d504e3ef533571837f5e32a"), "departureAirport" : "MUC", "arrivalAirport" : "SFO", "aircraft" : "Airbus A380", "distance" : 12000, "intercontinental" : true }
```

## Update

```
db.flightData.updateOne({distance: 12000}, {$set: {marker: "delete"}})
db.flightData.updateOne({}, {$set: {marker: "toDelete"}})
```

```
db.flightData.updateOne({_id: ObjectId("5d504e3ef533571837f5e32a")}, {$set: {delayed: true}})
```

```
db.flightData.update({_id: ObjectId("5d504e3ef533571837f5e32a")}, {delayed: false})
WriteResult({ "nMatched" : 1, "nUpserted" : 0, "nModified" : 1 })
db.flightData.find().pretty()
{ "_id" : ObjectId("5d504e3ef533571837f5e32a"), "delayed" : false }
{
        "_id" : ObjectId("5d504ef6f533571837f5e32b"),
        "departureAirport" : "LHR",
        "arrivalAirport" : "TXL",
        "aircraft" : "Airbus A320",
        "distance" : 950,
        "intercontinental" : false
}
```

```
db.flightData.replaceOne({_id: ObjectId("5d504e3ef533571837f5e32a")}, {"departureAirport": "MUC", "arrivalAirport": "SFO", "aircraft": "Airbus A380", "distance": 12000, "intercontinental": true})
```

## Delete

```
db.flightData.deleteOne({departureAirport: "TXL"})
{ "acknowledged" : true, "deletedCount" : 1 }

db.flightData.deleteOne({departureAirport: "TXL"})
{ "acknowledged" : true, "deletedCount" : 1 }

db.flightData.deleteMany({})
```
