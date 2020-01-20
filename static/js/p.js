var makePageLists = function (q) {
    var pageList = {}
    var firstPage = $.extend(true, {}, q)
    firstPage.pn = 0;
    pageList.firstPage = '/p' + '?' + $.param(firstPage)
    var nextPage = $.extend(true, {}, q)
    nextPage.pn = q.pn + 50;
    pageList.nextPage = '/p' + '?' + $.param(nextPage)
    var lastPage = $.extend(true, {}, q)
    lastPage.pn = q.pn - 50;
    if (lastPage.pn <0){
        lastPage.pn = 0
    }
    pageList.lastPage = '/p' + '?' + $.param(lastPage)
    var sidePages = []
    for (i = -5; i <= 5; i++) {
        var page = $.extend(true, {}, q)
        pn = ~~(q.pn/50) + i * 50;
        if(pn >= 0){
            page.pn = pn
            sidePages.push({
                url:'/p' + '?' + $.param(page),
                page: ~~(page.pn/50) +1
            })
        }
    }
    pageList.sidePages = sidePages
    return pageList
};


var router = new VueRouter({
    mode: 'history',
    routes: []
});


var post_list = new Vue({
    router,
    el: 'div#app',
    renderError(h, err) {
        return h('pre', { style: { color: 'red' } }, err.stack)
    },
    data: {
        message: 'Hello Vue!',
        posts: {},
        q: {},
        pageList: {}
    },
    mounted: function () {
        // this.threads = testdata.threads
        this.q = this.$route.query
        if ($.isEmptyObject(this.q.pn)) {
            this.q.pn = 0
        } else {
            this.q.pn = parseInt(this.q.pn)
        }
        if (this.q < 0) {
            this.q.pn = 0
        }

        this.pageList = makePageLists(this.q)

        console.log(this.q)
        fetch('/api/p/' + this.q.kz + '?' + $.param(this.q))
            .then(
                function (response) {
                    if (response.status !== 200) {
                        console.log('Looks like there was a problem. Status Code: ' +
                            response.status);
                        return;
                    }

                    // Examine the text in the response
                    response.json().then(function (data) {
                        this.post_list.posts = data.posts
                    });
                }
            )
            .catch(function (err) {
                console.log('Fetch Error :-S', err);
            });
    },
});

