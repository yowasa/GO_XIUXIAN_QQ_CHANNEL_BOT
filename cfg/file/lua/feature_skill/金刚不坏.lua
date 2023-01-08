local json = require("json")

--计算防御时 (防御+10)*1.2
function inCalAttr(u)
    t = json.decode(u)
    t["BattleInfo"]["DEF"]=math.ceil((t["BattleInfo"]["DEF"]+10)*1.2)
    str_json = json.encode(t)
    return str_json
end