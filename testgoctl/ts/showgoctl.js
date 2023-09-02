"use strict";
exports.__esModule = true;
var gocliRequest_1 = require("./gocliRequest");
/**
 * @description ping server
 * @param req
 */
function pinghandler(req) {
    return gocliRequest_1["default"].post("/ping", req);
}
exports.pinghandler = pinghandler;
