/* eslint-disable no-mixed-spaces-and-tabs */
const axios = require('axios');

const config = require('../../config');

const apiKey = config.get('currConv_api_key');

const currconvBaseUrl = `https://free.currconv.com/api/v7/convert?&compact=ultra&apiKey=${apiKey}`;

const fetchConvertCur = async (currCodeFrom, currCodeTo) => {
	 const queryCurrency = `&q=${currCodeFrom}_${currCodeTo}`;
	 const { data } = await axios.get(currconvBaseUrl + queryCurrency);

	 return data[`${currCodeFrom}_${currCodeTo}`];
};

module.exports = { fetchConvertCur };
