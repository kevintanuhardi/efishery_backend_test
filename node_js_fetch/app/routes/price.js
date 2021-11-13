const priceCtr = require('../controllers/price');

module.exports = (router) => {
  // TODO: add auth
  router.get('/', priceCtr.getPrice);
};
