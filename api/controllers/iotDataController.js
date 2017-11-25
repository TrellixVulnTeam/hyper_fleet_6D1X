'use strict';


var mongoose = require('mongoose'),
  IotData = mongoose.model('IotData');

exports.create_a_iotdata = function(req, res) {
  var new_iot_data = new IotData(req.body);

  console.info('received object: ' + new_iot_data.name);
  res.json('data received! ' + new_iot_data.name);

  /*
  new_iot_data.save(function(err, data) {
    if (err)
      res.send(err);
    res.json(data);
  });
  */
};

exports.list = function(req, res) {
  IotData.find({}, function(err, task) {
    if (err)
      res.send(err);
    res.json(task);
  });
};
