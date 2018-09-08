Targeting Hackathon api
============

# Routes
## RPI (Raspberry PI) endpoints
#### These endpoints are called by RPI and sends age interval + gender
#### All endpoints should have GET param `showroomId`
- /in
    - called by RPI and sends age interval + gender (when person entered)
    - GET params: age, gender
- /out
    - called by RPI and sends age interval + gender (when person exists)
    - GET params: age, gender
- /front
    - called by RPI and sends age interval + gender (when person sits in front of a camera)
    - GET params: age, gender, cameraId
## Scenario routes (each will display recommendations based on something)
- /all
    - return recommendations based on people in the room (age and gender)
- /person
    - todo: return recommendations based on entered email (customer)
    - GET params: email
- /camera
    - return recommendations based on who is in front of camera (age and gender)
    - GET params: cameraId

## How to build webserver
- `go build src/webserver/webserver.go`
## How to run webserver
- `./webserver bolt://IP:PORT`
    - where IP:PORT is neo4j server
    - localhost:8080 is the server url after starts
