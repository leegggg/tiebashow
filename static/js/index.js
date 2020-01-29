var router = new VueRouter({
    mode: 'history',
    routes: []
});


var indexPage = new Vue({
    router,
    el: 'div#app',
    renderError(h, err) {
        return h('pre', { style: { color: 'red' } }, err.stack)
    },
    data: {
        message: 'Hello Vue!',
        status: {
            status: 1000,
            msg: 'Unknown',
            config: {
                img_dir: 'unknown',
                db_path: 'unknown'
            },
            kws: []
        }
    },
    mounted: function () {
        fetch('/api/status')
            .then(
                function (response) {
                    if (response.status !== 200) {
                        console.log('Looks like there was a problem. Status Code: ' +
                            response.status);
                        return;
                    }

                    // Examine the text in the response
                    response.json().then(function (data) {
                        this.indexPage.status = data
                    });
                }
            )
            .catch(function (err) {
                console.log('Fetch Error :-S', err);
            });
    },
});

