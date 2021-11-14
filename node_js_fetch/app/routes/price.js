const priceCtr = require('../controllers/price');
const middleware = require('../middlewares');

module.exports = (router) => {
  // TODO: add auth
  router.get('/', middleware.auth.validateToken({ adminOnly: false }), priceCtr.getPrice);
  router.get('/aggregate', middleware.schemaValidations.getPriceAggregate, middleware.auth.validateToken({ adminOnly: true }), priceCtr.getPriceAggregate);
};
