const methodsTemplate = {
    "get": {
        method: "get"
    },
    "post": {
        method: "post"
    },
    "put": {
        method: "put"
    },
    "patch": {
        method: "patch"
    }
};
const urlsTemplate = {
    "user": {
        "get": {
            "info": "user/me",
            "logout": "user/logout",
        },
        "post": {
            "register": "user/register",
            "login": "user/login",
            "list": "users",
            "updateMe": "user/update/me",
            "update": "user/update",
            "dealer": "user/dealer",
        }
    },
    "drug": {
        "post": {
            "add": "drug/add",
            "list": "drug/list",
            "sale": "drug/sale",
            "dealer": "drug/dealer",
            "accept": "drug/accept",
            "buy": "drug/buy",
            "buyList": "drug/buy/list",
        }
    }

};
const urls = {};
Object.keys(urlsTemplate).forEach(function (group) {
    Object.keys(urlsTemplate[group]).forEach(function (methodName) {
        Object.keys(urlsTemplate[group][methodName]).forEach(function (name) {
            if (typeof urlsTemplate[group][methodName][name] === 'string') {
                urlsTemplate[group][methodName][name] = {
                    url: urlsTemplate[group][methodName][name]
                };
            }
            if (urls[group] === undefined) {
                urls[group] = {};
            }
            if (urls[group][methodName] === undefined) {
                urls[group][methodName] = {};
            }
            urls[group][methodName][name] = Object.assign({}, methodsTemplate[methodName], urlsTemplate[group][methodName][name]);
        });
    });
});
console.log(urls)
export default urls