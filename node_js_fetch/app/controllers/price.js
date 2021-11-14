/* global Helpers */
const priceSrv = require('../services/price');

module.exports = {
  getPrice: async (_, res) => {
    try {
      const priceList = await priceSrv.getPrice();

      return Helpers.successResponse(
        res,
        200,
        { data: priceList },
      );
    } catch (err) {
      return Helpers.errorResponse(res, null, err);
    }
  },
  getPriceAggregate: async (req, res) => {
    try {
      const {
        groupBy,
      } = req.query;

      const priceListAgg = await priceSrv.getPriceAggregate(groupBy);

      return Helpers.successResponse(
        res,
        200,
        { data: priceListAgg },
      );
    } catch (err) {
      return Helpers.errorResponse(res, null, err);
    }
  },
};
