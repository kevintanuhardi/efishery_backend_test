/* global Helpers */
const Joi = require('joi');

module.exports = {
  getPriceAggregate: (req, res, next) => {
    const { query } = req;

    const validateRes = Joi.object({
      groupBy: Joi.string().valid('week', 'province').required(),
    }).validate(query);

    if (validateRes.error) {
      return Helpers.errorResponse(res, null, validateRes.error.details[0]);
    }
    return next();
  },
};
