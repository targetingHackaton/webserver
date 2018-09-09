Targeting Hackathon api
============

# Short information
- This project is made in GoLang and is a webserver for 2 cases
    - for the interface (GUI) made in Symfony, which gets recommendations from this webserver (product ids)
    - for a RPI (Raspberry PI) that sends data in. ex:
        - each time a customer enters a showroom, send it's data (age, gender)         
        - each time a customer leaves a showroom, send it's data (age, gender)         
        - each time a customer approaches a TV+Camera in a showroom, send it's data (age, gender)
- The project uses neo4j for the Graph Database
    - The database contains data and relations between customers, products, orders, etc         

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
    - return recommendations based on entered email (customer)
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
