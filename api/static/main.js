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
        meta.textContent = `Written by ${data.author} on ${data.latestPublishedAt}`;
        wrapper.appendChild(meta);

        for (const paragraph of data.paragraphs) {
            const p = document.createElement('p');
            p.textContent = paragraph.text;
            wrapper.appendChild(p);
        }

        return wrapper;
    }

    getStyle() {
        const style = document.createElement('style');
        style.textContent = `
          .wrapper {
            border: 1px solid darkblue;
            border-radius: 4px;
            box-shadow: 5px 5px 5px #ddd;
            display: inline-block;
            padding: 1.5em;
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
                .forEach(item => document.body.appendChild(item));
        });
}

init();
