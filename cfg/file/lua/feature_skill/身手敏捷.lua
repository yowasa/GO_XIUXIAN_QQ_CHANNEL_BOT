local json = require("json")
--创建角色时敏捷低于60则重新roll一个60-100的随机数
function inCreate(u)
    t = json.decode(u)
    print(t["User"])
    str_json = json.encode(u)
    return str_json
end

--计算属性时 速度*1.1倍
function inCalAttr(u)
    t = json.decode(u)
    print(t["User"])
    str_json = json.encode(u)
    return str_json
end