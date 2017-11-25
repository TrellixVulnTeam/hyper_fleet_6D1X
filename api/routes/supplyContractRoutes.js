'use strict';
module.exports = function(app) {
  var controller = require('../controllers/supplyContractController');

  app.route('/supply_contract')
      .get(controller.show);
};
