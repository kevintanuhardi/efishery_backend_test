const jwtService = require('../services/jwtService');

module.exports = {
  validateToken: ({adminOnly = false}) => async (req, _, next) => {
    const {
      headers,
    } = req;
    try {
      const header = headers.Authorization || headers.authorization;
      if (!header) throw ({ status: 401, message: 'authorization is not exist' });

      const bearer = header.split(' ');

      const response = await jwtService.getDataFromToken(bearer[1]);
      if (adminOnly && response.role !== 'admin') {
        throw ({ status: 401, message: 'This endpoint is for admin only' });
      }
      // req.userId = userId;
      return next();
    } catch (err) {
      return next(err);
    }
  },
};
