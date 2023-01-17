const url = "/data_zh"
var dic = new Array()
var arr=new Array('Pokemon', 'Ability', 'TeraType','Nature','Item', 'Move')
var Pokemons=new Array('Nickname', 'Pokemon', 'Gender', 'Item', 'Ability', 'Level', 'TeraType', 'EVs', 'Nature', 'Move_1', 'Move_2', 'Move_3', 'Move_4')
function addopt(data){
    for(var i=0; i < 9; i++) {    
        if (data[i] != undefined){
            for(var j=0; j < data[i].length; j++) {
                if (data[i][j] != undefined) {
                    var opt=document.createElement('option')
                    opt.value = data[i][j]
                    document.getElementById(arr[i]).appendChild(opt)
                }
            }            
        }      
    }
}

async function logFetch(url) {
    try {
        const data = await fetch(url).then(data => data.json());        
        addopt(data)
    } catch (err) {
        console.log('fetch failed', err);
    }
}
  
window.onload = function(){
    logFetch(url) 
}

function add(){
    var arr = new Array()
    for(var i=0; i<13; i++) {
        if (document.querySelector("." + Pokemons[i]) != undefined) {
            arr[Pokemons[i]] = document.querySelector("." + Pokemons[i]).value
        } else {
            arr[Pokemons[i]] = ""
        }
    }
    dic.push(arr)
}

function submisson(){
    const options = {
        method: 'POST',
        headers: {
        'Content-Type': 'application/json',
        },
        body: JSON.stringify(dic),
        };        
    const url = "/upload"
    fetch(url, options).then(data => console.log(data.json()));
}

