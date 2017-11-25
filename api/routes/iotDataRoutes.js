'use strict';
module.exports = function(app) {
  var controller = require('../controllers/iotDataController');

  app.route('/iot_data')
      .get(controller.list)
      .post(controller.create_a_iotdata);
};
