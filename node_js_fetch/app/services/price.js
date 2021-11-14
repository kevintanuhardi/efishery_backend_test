/* eslint-disable new-cap */
/* global Helpers */
const moment = require('moment');
const efisherySrv = require('./eFishery');
const currConvSrv = require('./currconv');
const constant = require('../helpers/constant');

module.exports = {
  getPrice: async () => {
    const priceList = await efisherySrv.fetchPriceList();

    const currentUsdToIdr = await currConvSrv.fetchConvertCur('IDR', 'USD');

    return priceList.map((priceDatum) => {
      const parsedPrice = {
        ...priceDatum,
        price: Number(priceDatum.price),
        priceInUSD: priceDatum.price * currentUsdToIdr,
      };

      if (parsedPrice[constant.provinceField]) {
        parsedPrice[constant.provinceField] = parsedPrice[constant.provinceField].toUpperCase();
      }

      return parsedPrice;
    });
  },
  getPriceAggregate: async (groupBy) => {
    const priceList = await module.exports.getPrice();
    const priceMapping = {};
    const aggregateData = {};

    if (groupBy === 'province') {
      priceList.forEach((priceDatum) => {
        // continue loop if province is undefined or null
        if (!priceDatum[constant.provinceField]) return;

        if (priceMapping[priceDatum[constant.provinceField]]) {
          priceMapping[priceDatum[constant.provinceField]].push(priceDatum);
          aggregateData[priceDatum[constant.provinceField]].sum += Number(priceDatum.price);
        } else {
          priceMapping[priceDatum[constant.provinceField]] = [priceDatum];

          aggregateData[priceDatum[constant.provinceField]] = {
            sum: Number(priceDatum.price),
          };
        }
      });
    } else if (groupBy === 'week') {
      priceList.forEach((priceDatum) => {
        // continue loop if province is undefined or null
        if (!priceDatum.timestamp) return;

        const weeknumber = new moment.unix(priceDatum.timestamp).week();

        if (priceMapping[weeknumber]) {
          priceMapping[weeknumber].push(priceDatum);
          aggregateData[weeknumber].sum += Number(priceDatum.price);
        } else {
          priceMapping[weeknumber] = [priceDatum];

          aggregateData[weeknumber] = {
            sum: Number(priceDatum.price),

          };
        }
      });
    }
    // return priceMapping;
    Object.keys(priceMapping).forEach(key => {
      priceMapping[key] = priceMapping[key].sort((a, b) => {
        if (a.price < b.price) {
          return -1;
        }
        if (a.price > b.price) {
          return 1;
        }
        return 0;
      });
      const medianIndex = (
        (priceMapping[key].length + 1) / 2
      ) - 1;
      let medianValue;

      if (Number.isInteger(medianIndex)) {
        medianValue = Number(
          priceMapping[key][medianIndex].price,
        );
      } else {
        medianValue = (
          Number(priceMapping[key][medianIndex - 0.5].price)
					+ Number(priceMapping[key][medianIndex + 0.5].price)
        ) / 2;
      }

      aggregateData[key].median = medianValue;
      aggregateData[key].average = aggregateData[key].sum / priceMapping[key].length;
      // eslint-disable-next-line prefer-destructuring
      aggregateData[key].min = priceMapping[key][0].price;
      aggregateData[key].max = priceMapping[key][priceMapping[key].length - 1].price;

      // // TODO: improvement to insertion sort
      // // let startIdx;
      // // let endIdx;

      // // if (priceDatum.price > medianValue) {
      // //   startIdx = medianIndex;
      // //   endIdx = priceMapping[priceDatum[constant.provinceField]].length - 1;
      // // } else {
      // //   startIdx = 0;
      // //   endIdx = medianIndex;
      // // }

      // for (let i = 0; i < priceMapping[priceDatum[constant.provinceField]].length; i++) {
      //   if (
      //     Number(priceDatum.price)
      // 		<= Number(priceMapping[priceDatum[constant.provinceField]][i])
      //   ) {
      //     priceMapping[priceDatum[constant.provinceField]].splice(i, 0, priceDatum);
      //     break;
      //   } else if (i === priceMapping[priceDatum[constant.provinceField]].length - 1) {
      //     priceMapping[priceDatum[constant.provinceField]].push(priceDatum);
      //   }
      // }
    });
    return aggregateData;
  },
};
