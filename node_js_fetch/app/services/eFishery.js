const axios = require('axios');

const efisheryBaseUrl = 'https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list';

const fetchPriceList = async () => {
	 const { data } = await axios.get(efisheryBaseUrl);

	 return data.filter((priceDatum) => priceDatum.uuid && priceDatum.price);
};

module.exports = { fetchPriceList };
