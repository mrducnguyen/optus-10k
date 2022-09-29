attachEvents();

function attachEvents() {
    const form = document.getElementById('checkForm');
    form.addEventListener('submit', async event => {
        event.preventDefault();
        const res = await fetch('/check', {
            method: 'post',
            body: new FormData(form)
        });
        const result = await res.text();
        if (result === 'true') {
            highlightFound();
        } else {
            highlightNotFound();
        }
    });
}

function highlightFound() {
    const result = document.getElementById('result');
    result.classList.remove('not-found');
    result.classList.add('found');
    result.innerText = "You mobile or phone was FOUND in the public 10.2k records!";
}

function highlightNotFound() {
    const result = document.getElementById('result');
    result.classList.remove('found');
    result.classList.add('not-found');
    result.innerText = "You mobile or phone was NOT FOUND in the public 10.2k records. This doesn't mean you're safe";
}