class PostCard extends HTMLElement {
    setPostData(data) {
        const shadow = this.attachShadow({mode: 'open'});
        const wrapper = this.getWrapper(data);
        shadow.appendChild(wrapper);
        const style = this.getStyle();
        shadow.appendChild(style);
    }

    getWrapper(data) {
        const wrapper = document.createElement('span');
        wrapper.setAttribute('class', 'wrapper');

        const title = document.createElement('h2');
        title.setAttribute('class', 'title');
        title.textContent = data.title;
        wrapper.appendChild(title);

        const meta = document.createElement('span');
        meta.setAttribute('class', 'meta');
        const formattedDate = (new Date(data.latestPublishedAt)).toLocaleString('nl-NL');
        meta.textContent = `Written by ${data.author} on ${formattedDate}`;
        wrapper.appendChild(meta);

        const nonTitleParagraphs = data.paragraphs.slice(1);
        for (const paragraph of nonTitleParagraphs) {
            const p = document.createElement('p');
            p.textContent = paragraph.text;
            wrapper.appendChild(p);
        }

        const a = document.createElement('a');
        a.setAttribute('href', '#');
        a.textContent = 'Read more';
        wrapper.appendChild(a);

        return wrapper;
    }

    getStyle() {
        const style = document.createElement('style');
        style.textContent = `
          .wrapper {
            background-color: aliceblue;
            border: 1px solid cornflowerblue;
            border-radius: 4px;
            box-shadow: 5px 5px 5px #ddd;
            box-sizing: border-box;
            display: inline-block;
            height: 100%;
            padding: 1.5em;
          }
          
          h2.title {
            margin-bottom: 0;
          }
          
          span.meta {
            color: #555;
            display: block;
            font-size: 80%;
            margin-bottom: 1em;
          }
        `;
        return style;
    }
}

customElements.define('post-card', PostCard);

function init() {
    fetch('/posts')
        .then(function(response) {
            return response.json();
        })
        .then(function(data) {
            data
                .map(item => {
                    const postCard = document.createElement('post-card');
                    postCard.setPostData(item);
                    return postCard;
                })
                .forEach(item => document.getElementById('main').appendChild(item));
        });
}

init();
