'use strict';
module.exports = function(app) {
  var controller = require('../controllers/supplyContractController');

  app.route('/supply_contract')
      .get(controller.show);

  app.route('/reset_state_to_normal')
      .get(controller.reset_to_normal);
};
