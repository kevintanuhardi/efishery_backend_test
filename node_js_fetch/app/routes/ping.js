/* global Helpers */

module.exports = (router) => {
  router.get('/', (req, res) => Helpers.successResponse(res, 200, 'pong'));
};
