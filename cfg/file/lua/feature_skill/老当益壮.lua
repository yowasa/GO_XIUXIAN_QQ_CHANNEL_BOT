local json = require("json")

--属性随着年龄增长而增长
function inCalAttr(u)
    t = json.decode(u)
    t["BattleInfo"]["HP"]=t["BattleInfo"]["HP"]+t["Age"]
    t["BattleInfo"]["MP"]=t["BattleInfo"]["MP"]+t["Age"]
    t["BattleInfo"]["ATK"]=t["BattleInfo"]["ATK"]+math.ceil(t["Age"]/10)
    t["BattleInfo"]["DEF"]=t["BattleInfo"]["DEF"]+math.ceil(t["Age"]/10)
    t["BattleInfo"]["SPD"]=t["BattleInfo"]["SPD"]+math.ceil(t["Age"]/10)
    str_json = json.encode(t)
    return str_json
end