const middleware = require('../middlewares');
const authCtrl = require('../controllers/auth');

module.exports = (router) => {
  // TODO: add auth
  router.get('/token/introspect', middleware.auth.validateToken({ adminOnly: false }), authCtrl.tokenIntrospect);
};
