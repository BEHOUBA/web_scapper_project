let searchInput = document.getElementById("search-input")
let listing = document.getElementById("listing")


function searchRequest() {
    console.log(searchInput.value)

    axios.get(`/articles/?q=${searchInput.value}`)
        .then(res => {
            var data = JSON.parse(res.data)
            var items = ""
            data.forEach(element => {
                items += createItem(element)
            });
            listing.innerHTML = items;
        })
        .catch(err => {
            console.log(err)
        })
}

function createItem(data) {
    var html = `<a href="${data.link}">
                    <div class="card">
                    <div class="image">
                        <img src="${data.picture}">
                    </div>
                    <div class="info">
                        <p>${data.title}</p>
                        <p>${data.price}</p>
                    </div>
                    </div>
                </a>`
    return html
}