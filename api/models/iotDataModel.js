'use strict';
var mongoose = require('mongoose');
var Schema = mongoose.Schema;


var IotDataSchema = new Schema({
  name: {
    type: String,
    required: 'some name for IoT data'
  },
  Created_date: {
    type: Date,
    default: Date.now
  }/*,
  status: {
    type: [{
      type: String,
      enum: ['pending', 'ongoing', 'completed']
    }],
    default: ['pending']
  }*/
});

module.exports = mongoose.model('IotData', IotDataSchema);
