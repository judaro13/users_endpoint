### App description
This app is like a messages producer for rabbitmq and expose two paths:

GET /       -> this route is for check if is active the app
POST /users -> this route is for send the user data to a rabbitqm queue

The post endpoint receive a form-urlencoded with the user attributes, this values are validates, if the validation is success, is created a string with the user data in json structure, this string is send to a specified channel in rabbitqm.
When exist a missing field in the user data, is send a Status Bad Request response with the missing data values.
If there is no errors, the application response with Ok status and a json with the user data sent.

In this project was used  gorilla/mux and gorilla/schema as router and streadway/amqp to handle rabbitmq queue.


### Install
No special instructions for install,  this project if configured to run in port 8000 and are needed some env vars for rabbitqm:

RABBIT_PATH="amqp://guest:guest@localhost:5672/"
RABBIT_CHANNEL="postUser"
