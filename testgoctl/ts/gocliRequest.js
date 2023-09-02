"use strict";
var __assign = (this && this.__assign) || function () {
    __assign = Object.assign || function(t) {
        for (var s, i = 1, n = arguments.length; i < n; i++) {
            s = arguments[i];
            for (var p in s) if (Object.prototype.hasOwnProperty.call(s, p))
                t[p] = s[p];
        }
        return t;
    };
    return __assign.apply(this, arguments);
};
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (_) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
exports.__esModule = true;
/**
 * Parse route parameters for responseType
 */
var reg = /:[a-z|A-Z]+/g;
function parseParams(url) {
    var ps = url.match(reg);
    if (!ps) {
        return [];
    }
    return ps.map(function (k) { return k.replace(/:/, ''); });
}
exports.parseParams = parseParams;
/**
 * Generate url and parameters
 * @param url
 * @param params
 */
function genUrl(url, params) {
    if (!params) {
        return url;
    }
    var ps = parseParams(url);
    ps.forEach(function (k) {
        var reg = new RegExp(":" + k);
        url = url.replace(reg, params[k]);
    });
    var path = [];
    var _loop_1 = function (key) {
        if (!ps.find(function (k) { return k === key; })) {
            path.push(key + "=" + params[key]);
        }
    };
    for (var _i = 0, _a = Object.keys(params); _i < _a.length; _i++) {
        var key = _a[_i];
        _loop_1(key);
    }
    return url + (path.length > 0 ? "?" + path.join('&') : '');
}
exports.genUrl = genUrl;
function request(_a) {
    var method = _a.method, url = _a.url, data = _a.data, _b = _a.config, config = _b === void 0 ? {} : _b;
    return __awaiter(this, void 0, void 0, function () {
        var response;
        return __generator(this, function (_c) {
            switch (_c.label) {
                case 0: return [4 /*yield*/, fetch(url, __assign({ method: method.toLocaleUpperCase(), credentials: 'include', headers: {
                            'Content-Type': 'application/json'
                        }, body: data ? JSON.stringify(data) : undefined }, config))];
                case 1:
                    response = _c.sent();
                    return [2 /*return*/, response.json()];
            }
        });
    });
}
exports.request = request;
function api(method, url, req, config) {
    if (method === void 0) { method = 'get'; }
    if (url.match(/:/) || method.match(/get|delete/i)) {
        url = genUrl(url, req.params || req.forms);
    }
    method = method.toLocaleLowerCase();
    switch (method) {
        case 'get':
            return request({ method: 'get', url: url, data: req, config: config });
        case 'delete':
            return request({ method: 'delete', url: url, data: req, config: config });
        case 'put':
            return request({ method: 'put', url: url, data: req, config: config });
        case 'post':
            return request({ method: 'post', url: url, data: req, config: config });
        case 'patch':
            return request({ method: 'patch', url: url, data: req, config: config });
        default:
            return request({ method: 'post', url: url, data: req, config: config });
    }
}
exports.webapi = {
    get: function (url, req, config) {
        return api('get', url, req, config);
    },
    "delete": function (url, req, config) {
        return api('delete', url, req, config);
    },
    put: function (url, req, config) {
        return api('get', url, req, config);
    },
    post: function (url, req, config) {
        return api('post', url, req, config);
    },
    patch: function (url, req, config) {
        return api('patch', url, req, config);
    }
};
exports["default"] = exports.webapi;
