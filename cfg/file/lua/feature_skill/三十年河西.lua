local json = require("json")

--50岁之前大部分属性下降10% 50岁之后大部分属性增加10%
function inCalAttr(u)
    t = json.decode(u)




    str_json = json.encode(u)
    return str_json
end