const jwt = require('jsonwebtoken');

const config = require('../../config');

const jwtSecret = config.get('jwt_secret');

const getDataFromToken = async (bearerToken) => jwt.verify(
  bearerToken, jwtSecret, (err, decoded) => {
    if (err) {
      throw (err);
    }
    // if (decoded && decoded.exp >= new Date().valueOf() / 1000) {
    return {
      active: true,
      phoneNumber: decoded.phone,
      name: decoded.name,
      role: decoded.role,
    };
    // }
    // throw ({
    //   message: 'Token unauthorized',
    //   status: 401,
    // });
  },
);

module.exports = { getDataFromToken };
