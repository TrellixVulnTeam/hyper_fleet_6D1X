'use strict';
var mongoose = require('mongoose');
var Schema = mongoose.Schema;


var SupplyContractSchema = new Schema({
  name: {
    type: String,
    required: 'contract#1'
  },
  status: {
    type: String,
    required: 'status_ok'
  }
});

module.exports = mongoose.model('SupplyContract', SupplyContractSchema);
