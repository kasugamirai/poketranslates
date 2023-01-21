const url = "/data_zh"
var dic = new Array()
var arr=new Array('Pokemon', 'Ability', 'TeraType','Nature','Item', 'Move')
var Pokemons=new Array('Nickname', 'Pokemon', 'Gender', 'Item', 'Ability', 'Level', 'TeraType', 'IVs', 'EVs','Nature', 'Move_1', 'Move_2', 'Move_3', 'Move_4')
var ivs = ['HP', 'Atk', 'Def', 'SpA', 'SpD', 'Spe']
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
        const data = await getJSON(url);
        addopt(data)
}
  
window.onload = function(){
    logFetch(url)
    for (var i = 0; i < 6; i++) {
        let input = document.createElement('input')
        input.id = ivs[i]
        input.value = "31"
        input.max = "31"
        input.min = "0"
        input.type = "number"
        input.step = "1"
        document.getElementById('IVs').append(ivs[i], input)
    }
    for (var i = 0; i < 6; i++) {
        let input = document.createElement('input')
        input.id = 'stat-' + ivs[i]
        input.value = "0"
        input.max = "252"
        input.min = "0"
        input.type = "number"
        input.step = "1"
        document.getElementById('EVs').append(ivs[i], input)
    }
}

function add(){
    var EVs
    var IVs
    var arr = {}
    for(var i=0; i<14; i++) {
        if (document.querySelector("." + Pokemons[i]) != undefined) {
            arr[Pokemons[i]] = document.querySelector("." + Pokemons[i]).value
        } else {
            arr[Pokemons[i]] = ""
        }
    }
    for(var j=0; j<6; j++) {
        if (document.getElementById('stat-' + ivs[j]).value != 0) {
            if (EVs == undefined) {
                EVs = document.getElementById('stat-' + ivs[j]).value + ' ' + ivs[j]
            } else {
                EVs += ' / ' + document.getElementById('stat-' + ivs[j]).value + ' ' + ivs[j]
            }
        }
        if (document.getElementById(ivs[j]).value != 31){
            if (IVs == undefined) {
                IVs = document.getElementById(ivs[j]).value + ' ' + ivs[j]
            } else {
                IVs += ' / ' + document.getElementById(ivs[j]).value + ' ' + ivs[j]
            }
        }
    }
    arr['EVs'] = EVs
    arr['IVs'] = IVs 
    dic.push(arr)
    console.log(dic)
    let span = document.createElement("span")
    document.getElementById('team').append(arr.Pokemon, span)
}



async function submisson() {
    res = await postJSON("/upload", dic)
    console.log(res.link)
    alert("查询码: " + res.link)
}

async function postJSON(url, data) {
    try {
        let response = await fetch(url, {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify(data),
        }).then(data => data.json())
        return response
        } catch (error) {
        console.log('Request Failed', error);
      }
}

async function getJSON(url) {
    try {
      let response = await fetch(url);
      return await response.json();
    } catch (error) {
      console.log('Request Failed', error);
    }
}

