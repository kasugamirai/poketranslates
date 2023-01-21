async function search() {
    res = await postJSON("/search", document.getElementById('link').value)
    ans = res.res.replace(/(\n)/g, '<br>')
    document.getElementById('ret').innerHTML = ans
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
