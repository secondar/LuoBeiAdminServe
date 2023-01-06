{{template "home/common/header.tpl" .}}

<body>
  {{template "home/common/nav.tpl" .}}

  <div class="article-main">
    <div class="box-item banner">
      <div class="layui-carousel" id="banner">
        <div carousel-item>
          <div><img src="./static/home/img/20211223100157.png" alt=""></div>
          <div><img src="./static/home/img/20211223100157.png" alt=""></div>
        </div>
      </div>
    </div>
    <div class="box-item">
      <div class="sort">
        <div class="item {{if eq 0 .Sort}}active{{end}}"><a href="/article/0-1.html">全部</a></div>
        {{range $index, $elem := .SortList}}
        <span class="separate"></span>
        <div class="item  {{if eq $elem.Id $.Sort}}active{{end}}"><a
            href="/article/{{$elem.Id}}-1.html">{{$elem.Title}}</a></div>
        {{end}}
      </div>
      <div class="list">
        {{range $index, $elem := .ArticleList}}
        <div class="item">
          <div class="avatar">
            <img src="https://www.bugquit.com/wp-content/uploads/2019/12/logo-1.png" alt="">
          </div>
          <div class="info">
            <div class="top">
              <span>
                {{range $indexSortList, $elemSortList := $.SortList}}
                {{if eq $elem.Sort $elemSortList.Id}}{{$elemSortList.Title}}{{end}}
                {{end}}
              </span>
              <a href="/article/details/{{$elem.Id}}.html" target="_blank">{{$elem.Title}}</a>
            </div>
            <div class="bottom">
              <div class="left">
                <div class="author">secondar</div>
                <div class="time">{{$elem.Addtime.GetStr}}</div>
              </div>
              <div class="right"><i class="icon iconfont">&#xf01f0;</i>{{$elem.Hot}}</div>
            </div>
          </div>
        </div>
        {{end}}
      </div>
      <div class="page" id="page">
        <a href="/article/{{.Sort}}-{{.NextPage}}.html">下一页</a>
      </div>
    </div>
  </div>

  {{template "home/common/footer.tpl" .}}
</body>
<script>
  layui.use(['carousel', 'laypage'], function () {
    var carousel = layui.carousel;
    var laypage = layui.laypage;
    carousel.render({
      elem: '#banner',
      width: '100%',
      arrow: 'always'
    });
    laypage.render({
      elem: 'page',
      count: '{{.Page.Count}}',
      limit: '{{.Page.Limit}}',
      curr: '{{.Page.Curr}}',
      layout: ['count', 'prev', 'page', 'next', 'skip'],
      jump: function (obj, first) {
        //首次不执行
        if (!first) {
          location.href = "/article/{{.Sort}}-" + obj.curr + ".html"
        }
      }
    });
  });
</script>

</html>