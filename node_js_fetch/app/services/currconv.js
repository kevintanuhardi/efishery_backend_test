/* eslint-disable no-mixed-spaces-and-tabs */
const axios = require('axios');
const fs = require('fs');
const moment = require('moment');
const url = require('url');

const config = require('../../config');

const apiKey = config.get('currConv_api_key');

const cacheFileLocation = 'app/data/cacheCurrency.json';

const currconvBaseUrl = `https://free.currconv.com/api/v7/convert?&compact=ultra&apiKey=${apiKey}`;

const fetchConvertCur = async (currCodeFrom, currCodeTo) => {
  const currentDate = new moment().format('YYYY-MM-DD');
  const rawData = fs.readFileSync(url.pathToFileURL(cacheFileLocation));
  const jsonCache = JSON.parse(rawData);

  if (jsonCache[currentDate] && jsonCache[currentDate][`${currCodeFrom}_${currCodeTo}`]) {
    return jsonCache[currentDate][`${currCodeFrom}_${currCodeTo}`];
  }

	 const queryCurrency = `&q=${currCodeFrom}_${currCodeTo}`;
	 const { data } = await axios.get(currconvBaseUrl + queryCurrency);

	 if (!jsonCache[currentDate]) jsonCache[currentDate] = {};

	 jsonCache[currentDate][`${currCodeFrom}_${currCodeTo}`] = data[`${currCodeFrom}_${currCodeTo}`];
	 fs.writeFileSync(url.pathToFileURL(cacheFileLocation), JSON.stringify(jsonCache, null, 2));

	 return data[`${currCodeFrom}_${currCodeTo}`];
};

module.exports = { fetchConvertCur };
