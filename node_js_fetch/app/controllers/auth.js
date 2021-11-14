/* global Helpers */

module.exports = {
  tokenIntrospect: async (req, res) => {
    try {
      return Helpers.successResponse(
        res,
        200,
        { data: req.privateClaims },
      );
    } catch (err) {
      return Helpers.errorResponse(res, null, err);
    }
  },
};
