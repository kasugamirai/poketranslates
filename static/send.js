window.onload = function(){
    var data
    var url = "localhost:8080/data"
    fetch(url).then(response => data)
    var arr=new Array('Nickname', 'Pokemon', 'Gender', 'Item', 'Ability', 'Level', 'TeraType', 'EVs', 'Nature', 'Move_1', 'Move_2', 'Move_3', 'Move_4')
    function addopt(){
               for(var i=0;i<counts;i++){
                    var opt=document.createElement('option')
                    document.getElementById(arr[i]).add(opt, data[i])
                }
    }
}
