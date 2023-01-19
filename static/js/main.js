const url = "/data_zh"
var dic = new Array()
var arr=new Array('Pokemon', 'Ability', 'TeraType','Nature','Item', 'Move')
var Pokemons=new Array('Nickname', 'Pokemon', 'Gender', 'Item', 'Ability', 'Level', 'TeraType', 'IVs', 'EVs','Nature', 'Move_1', 'Move_2', 'Move_3', 'Move_4')
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
    var EVs
    var IVs
    var arr = {}
    var ivs = ['HP', 'Atk', 'Def', 'SpA', 'SpD', 'Spe']
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
}

async function submissons(){
    const options = {
        method: 'POST',
        headers: {
        'Content-Type': 'application/json',
        },
        body: JSON.stringify(dic),
        };        
    const url = "/upload"
    const data = await fetch(url, options).then(data => console.log(data.json()));
    console.log(data)
}


async function submisson() {
    res = await postJSON("/upload", dic)
    console.log(res.link)
    alert(res.link)
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

function tests() {
    const data = { username: 'example' };
    var b = JSON.stringify(data)
    console.log(b)
    console.log(typeof b)

    fetch('/test', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    })
    .then(response => response.json())
    .then(data => {
        alert(JSON.parse(data))
      console.log('Success:', JSON.parse(data));
      
    })
    .catch((error) => {
      console.error('Error:', error);
    });
}
async function search() {
    res = await postJSON("/search", document.getElementById('link').value)
    ans = res.res.replace(/(\n)/g, '<br>')
    document.getElementById('ret').innerHTML = ans
}
