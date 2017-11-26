'use strict';


var mongoose = require('mongoose'),
  SupplyContract = mongoose.model('SupplyContract');

exports.show = function(req, res) {
  var invokerInstance = require('../../invoker');
  var response = invokerInstance.get_cargo_state();

  SupplyContract.find({}, function(err, contracts) {
    if (err) {
      res.send(err);
    }

    var last_element = contracts[contracts.length - 1];
    res.json(last_element);
  });
};

exports.reset_to_normal = function(req, res) {
  var resetterInstance = require('../../resetter');
  var response = resetterInstance.reset_cargo_state_to_normal();

  SupplyContract.find({}, function(err, contracts) {
    if (err) {
      res.send(err);
    }

    var last_element = contracts[contracts.length - 1];
    res.json(last_element);
  });
};
