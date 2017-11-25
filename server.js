var express = require('express'),
  app = express(),
  port = process.env.PORT || 3000,
  mongoose = require('mongoose'),
  IotData = require('./api/models/iotDataModel'), //created model loading here
  SupplyContract = require('./api/models/supplyContractModel')
  bodyParser = require('body-parser');

// mongoose instance connection url connection
mongoose.Promise = global.Promise;
mongoose.connect('mongodb://localhost/HyperFleetdb');


app.use(bodyParser.urlencoded({ extended: true }));
app.use(bodyParser.json());


var iot_routes = require('./api/routes/iotDataRoutes'); //importing route
iot_routes(app); //register the route

var contract_routes = require('./api/routes/supplyContractRoutes'); //importing route
contract_routes(app); //register the route


app.listen(port);


console.log('server started on: ' + port);
