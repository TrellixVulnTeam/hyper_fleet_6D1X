'use strict';
var mongoose = require('mongoose');
var Schema = mongoose.Schema;


var SupplyContractSchema = new Schema({
  state: {
    type: String,
    required: 'state_ok'
  }
});

module.exports = mongoose.model('SupplyContract', SupplyContractSchema);
