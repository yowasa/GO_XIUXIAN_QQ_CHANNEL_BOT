local json = require("json")

--计算防御时 (防御+10)*1.2
function inCalAttr(u)
    t = json.decode(u)
    print(t["User"])
    str_json = json.encode(u)
    return str_json
end