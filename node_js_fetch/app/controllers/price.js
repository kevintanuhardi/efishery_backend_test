/* global Helpers */
const efisherySrv = require('../services/eFishery');
const currConvSrv = require('../services/currconv');

module.exports = {
  getPrice: async (req, res) => {
    try {
      let priceList = await efisherySrv.fetchPriceList();

      const currentUsdToIdr = await currConvSrv.fetchConvertCur('IDR', 'USD');

      priceList = priceList.map((priceDatum) => ({
        ...priceDatum,
        priceInUSD: priceDatum.price * currentUsdToIdr,
      }));

      return Helpers.successResponse(
        res,
        200,
        { data: priceList },
      );
    } catch (err) {
      return Helpers.errorResponse(res, null, err);
    }
  },
};
