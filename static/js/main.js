window.onload = function(){
    var data
    var url = "/datazh"
    fetch(url).then(data => data.JSON)
    JSON.parse(data)
    data = data.append(data[5])
    data = data.append(data[5])
    data = data.append(data[5])
    var arr=new Array('Pokemon', 'Ability', 'TeraType','Nature','Item', 'Move_1', 'Move_2', 'Move_3', 'Move_4')
    function addopt(){
                for(var i=0; i < 9; i++) {                    
                    for(var j=0; j < data[i].length; j++) {
                        var opt=document.createElement('option')
                        opt.text = data[i][j]
                        opt.value = data[i][j]
                        document.getElementById(arr[i]).add(opt)
                    }                   
                }
    }
    addopt()
}
