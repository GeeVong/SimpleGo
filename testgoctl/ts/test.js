"use strict";
exports.__esModule = true;
var showgoctl_1 = require("./showgoctl");
var request = {
    req: 123 // 你需要传递的参数值
};
showgoctl_1.pinghandler(request)
    .then(function (response) {
    // 处理响应结果
    console.log(response.re);
})["catch"](function (error) {
    // 处理错误
    console.error(error);
});
