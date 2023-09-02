import { pinghandler } from './showgoctl';

const request = {
    req: 123 // 你需要传递的参数值
};

pinghandler(request)
    .then(response => {
        // 处理响应结果
        console.log(response.re);
    })
    .catch(error => {
        // 处理错误
        console.error(error);
    });
