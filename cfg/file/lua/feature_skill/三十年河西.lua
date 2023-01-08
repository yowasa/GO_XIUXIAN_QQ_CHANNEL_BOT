local json = require("json")

--50岁之前大部分属性下降10% 50岁之后大部分属性增加10%
function inCalAttr(u)
    t = json.decode(u)
    radio=1.1
    if t["Age"]<50
    then
        radio=0.9
    end
    t["BattleInfo"]["HP"]=math.ceil(t["BattleInfo"]["HP"]*radio)
    t["BattleInfo"]["MP"]=math.ceil(t["BattleInfo"]["MP"]*radio)
    t["BattleInfo"]["ATK"]=math.ceil(t["BattleInfo"]["ATK"]*radio)
    t["BattleInfo"]["DEF"]=math.ceil(t["BattleInfo"]["DEF"]*radio)
    t["BattleInfo"]["SPD"]=math.ceil(t["BattleInfo"]["SPD"]*radio)
    str_json = json.encode(t)
    return str_json
end