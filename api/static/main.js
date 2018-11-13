fetch('/posts')
    .then(function(response) {
        return response.json();
    })
    .then(function(data) {
        // console.log(data);
        const titles = data.map(item => item.title);
        const str = titles.join(', ');
        document.getElementById('main').innerText = str;
    });
