var resp;

fetch("/api/product", { method: "GET" })
.then(res => { return res.json(); })
.then(res => {
    let products_body = document.getElementsByClassName("products_body")[0];
    
    for (let prod of res.products) {
        let tr = `<tr>`;
        for (let key in prod) {
            tr += `<td>${prod[key]}</td>`
        }
        tr += `</tr>`;
        products_body.innerHTML += tr;
    }
});