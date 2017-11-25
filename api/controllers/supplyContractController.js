'use strict';


var mongoose = require('mongoose'),
  SupplyContract = mongoose.model('SupplyContract');

exports.show = function(req, res) {
  SupplyContract.find({}, function(err, task) {
    if (err)
      res.send(err);
    res.json(task);
  });

  var invokerInstance = require('../../invoker');
  var response = invokerInstance.get_cargo_state();

};
