const jwtService = require('../services/jwtService');

const { privateRoute } = require('../helpers/constant');

module.exports = {
  validateToken: async (req, _, next) => {
    const {
      headers,
      method,
      path,
    } = req;
    try {
      // TODO: ADD validation to backend golang
      return next();
    } catch (err) {
      return next(err);
    }
  },
};
