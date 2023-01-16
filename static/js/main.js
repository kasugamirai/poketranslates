window.onload = function(){
    var data
    var url = "/datazh"
    fetch(url).then(data => data.JSON)
    JSON.parse(data)
    Pokemon_zh, Abilities_zh, Types_zh, Natures_zh, Items_zh, Moves_zh
    var arr=new Array('Pokemon', 'Ability', 'TeraType','Nature','Item', 'Move_1', 'Move_2', 'Move_3', 'Move_4')
    function addopt(){
                for(var i=0; i < 9; i++) {                    
                    for(var j=0; j < 6; j++) {
                        var opt=document.createElement('option')
                        opt.text = data[i][1]
                        opt.value = data[1][1]
                        document.getElementById(arr[i]).add(opt)
                    }                   
                }
    }
    addopt()
}
