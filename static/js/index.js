let searchInput = document.getElementById("search-input")
let listing = document.getElementById("listing")
let searchForm = document.getElementById("search-form")


searchForm.addEventListener("submit", e => {
    e.preventDefault()
    if (searchInput.value === "") {
        return
    }
    searchRequest()
})


function searchRequest() {
    axios.get(`/articles/?q=${searchInput.value}`)
        .then(res => {
            listing.innerHTML = ""
            if (res.data) {
                var data = JSON.parse(res.data)
                var items = ""
                data.forEach(element => {
                    items += createItem(element)
                });
                listing.innerHTML = items;
            }
        })
        .catch(err => {
            console.log(err)
        })
}

function createItem(data) {
    var html = `<li>
                    <a href="${data.link}" target="blank">
                    <img src="${data.picture}" alt="">
                    <p>${data.title}</p>
                    <p>${data.price}</p>
                    <p>${data.origin}</p>
                    </a>
                </li>`
    return html
}