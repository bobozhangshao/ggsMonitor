function _interopRequireDefault(obj) {
    return obj && obj.__esModule ? obj : {
        default: obj
    };
}

function t(e, t) {
    var o = Object.keys(e);
    if (Object.getOwnPropertySymbols) {
        var n = Object.getOwnPropertySymbols(e);
        t && (n = n.filter(function (t) {
            return Object.getOwnPropertyDescriptor(e, t).enumerable;
        }))
        o.push.apply(o, n);
    }
    return o;
}

function o(o) {
    for (var n = 1; n < arguments.length; n++) {
        var a = null != arguments[n] ? arguments[n] : {};
        if (n % 2) {
            t(Object(a), !0).forEach(function (t) {
                (0, _interopRequireDefault)(o, t, a[t]);
            })
        } else {
            if (Object.getOwnPropertyDescriptors) {
                Object.defineProperties(o, Object.getOwnPropertyDescriptors(a))
            } else {
                t(Object(a)).forEach(function (e) {
                    Object.defineProperty(o, e, Object.getOwnPropertyDescriptor(a, e));
                });
            }
        }
    }
    return o;
}


function s(e) {
    var t = "";
    for (var o in e) t += o + "=" + e[o] + "&";
    t = t.length > 0 ? "?" + t.substring(0, t.length - 1) : ""
    return encodeURI(t);
}

function getSign(t, n) {
    var o = [];
    for (var r in t) o.push("".concat(r, "=").concat(t[r]));
    var a = "".concat(o.join("&")).concat("e348db70-2e67-4a72-9578-8b40ad809cbb");
    return n && (a = "".concat(a).concat(JSON.stringify(n))), a = a.replace(/a/gm, "c").replace(/e/gm, "g").replace(new RegExp(" ", "gm"), ""),
        (0, _interopRequireDefault)(a);
}

function post(params) {
    var a = {
        os: "APPLET",
        osVersion: "1.0.0",
        userId: "",
        userToken: ""
    }
    var query = o({}, a, {
        userId: "123456",
        userToken: "0123456789"
    })
    query.sign = getSign(o({}, query), t || {});
    // console.log(query)
    return s(query)
}


module.exports = {
    post: post,
};