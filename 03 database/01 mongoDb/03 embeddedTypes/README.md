# Find Commands

## Embedded 

embedded...

```
db.flightData.updateMany({}, {$set: {status: {description: "on-time", lastUpdated: "1 hour ago"}}})
db.flightData.find().pretty()
{
        "_id" : ObjectId("5d504e3ef533571837f5e32a"),
        "departureAirport" : "MUC",
        "arrivalAirport" : "SFO",
        "aircraft" : "Airbus A380",
        "distance" : 12000,
        "intercontinental" : true,
        "status" : {
                "description" : "on-time",
                "lastUpdated" : "1 hour ago"
        }
}
{
        "_id" : ObjectId("5d504ef6f533571837f5e32b"),
        "departureAirport" : "LHR",
        "arrivalAirport" : "TXL",
        "aircraft" : "Airbus A320",
        "distance" : 950,
        "intercontinental" : false,
        "status" : {
                "description" : "on-time",
                "lastUpdated" : "1 hour ago"
        }
}
```

embedded in embedded...

```
db.flightData.updateMany({}, {$set: {status: {description: "on-time", lastUpdated: "1 hour ago", details: {responsible: "Dagobert Duck"}}}})
db.flightData.find().pretty()
{
        "_id" : ObjectId("5d504e3ef533571837f5e32a"),
        "departureAirport" : "MUC",
        "arrivalAirport" : "SFO",
        "aircraft" : "Airbus A380",
        "distance" : 12000,
        "intercontinental" : true,
        "status" : {
                "description" : "on-time",
                "lastUpdated" : "1 hour ago",
                "details" : {
                        "responsible" : "Dagobert Duck"
                }
        }
}
{
        "_id" : ObjectId("5d504ef6f533571837f5e32b"),
        "departureAirport" : "LHR",
        "arrivalAirport" : "TXL",
        "aircraft" : "Airbus A320",
        "distance" : 950,
        "intercontinental" : false,
        "status" : {
                "description" : "on-time",
                "lastUpdated" : "1 hour ago",
                "details" : {
                        "responsible" : "Dagobert Duck"
                }
        }
}
```

### Find for Embedded

...

```
db.flightData.find({"status.description": "on-time"}).pretty()
{
        "_id" : ObjectId("5d504e3ef533571837f5e32a"),
        "departureAirport" : "MUC",
        "arrivalAirport" : "SFO",
        "aircraft" : "Airbus A380",
        "distance" : 12000,
        "intercontinental" : true,
        "status" : {
                "description" : "on-time",
                "lastUpdated" : "1 hour ago",
                "details" : {
                        "responsible" : "Dagobert Duck"
                }
        }
}
{
        "_id" : ObjectId("5d504ef6f533571837f5e32b"),
        "departureAirport" : "LHR",
        "arrivalAirport" : "TXL",
        "aircraft" : "Airbus A320",
        "distance" : 950,
        "intercontinental" : false,
        "status" : {
                "description" : "on-time",
                "lastUpdated" : "1 hour ago",
                "details" : {
                        "responsible" : "Dagobert Duck"
                }
        }
}
```

```
db.flightData.find({"status.details.responsible": "Dagobert Duck"}).pretty()
{
        "_id" : ObjectId("5d504e3ef533571837f5e32a"),
        "departureAirport" : "MUC",
        "arrivalAirport" : "SFO",
        "aircraft" : "Airbus A380",
        "distance" : 12000,
        "intercontinental" : true,
        "status" : {
                "description" : "on-time",
                "lastUpdated" : "1 hour ago",
                "details" : {
                        "responsible" : "Dagobert Duck"
                }
        }
}
{
        "_id" : ObjectId("5d504ef6f533571837f5e32b"),
        "departureAirport" : "LHR",
        "arrivalAirport" : "TXL",
        "aircraft" : "Airbus A320",
        "distance" : 950,
        "intercontinental" : false,
        "status" : {
                "description" : "on-time",
                "lastUpdated" : "1 hour ago",
                "details" : {
                        "responsible" : "Dagobert Duck"
                }
        }
}
```

## Arrays

lists...

```
db.passengers.updateOne({name: "Albert Twostone"}, {$set: {hobbies: ["sports", "cooking"]}})
db.passengers.find().pretty()
...
{
        "_id" : ObjectId("5d505797f533571837f5e33f"),
        "name" : "Albert Twostone",
        "age" : 68,
        "hobbies" : [
                "sports",
                "cooking"
        ]
}
Type "it" for more
```

### Find for Arrays

...

```
db.passengers.findOne({name: "Albert Twostone"}).hobbies
[ "sports", "cooking" ]
```

```
db.passengers.find({hobbies: "sports"}).pretty()
{
        "_id" : ObjectId("5d505797f533571837f5e33f"),
        "name" : "Albert Twostone",
        "age" : 68,
        "hobbies" : [
                "sports",
                "cooking"
        ]
}
```
