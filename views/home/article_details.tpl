{{template "home/common/header.tpl" .}}

<body>
  {{template "home/common/nav.tpl" .}}
  <div class="article-details">
    <div class="box-item banner">
      <div class="layui-carousel" id="banner">
        <div carousel-item>
          <div><img src="./static/home/img/20211223100157.png" alt=""></div>
          <div><img src="./static/home/img/20211223100157.png" alt=""></div>
        </div>
      </div>
    </div>
    <div class="box-item details">
      <div class="head">
        <div class="title">{{.Info.Title}}</div>
        <div class="info">
          <div class="author">
            <div>作者：</div>
            <div>secondar</div>
          </div>
          <div class="time">
            <div>发布时间：</div>
            <div>{{.InfoAddtime}}</div>
          </div>
          <div class="hot">
            <div>阅读量：</div>
            <div>{{.Info.Hot}}</div>
          </div>
        </div>
      </div>
      <div class="content">
        {{.Info.Content | str2html}}
      </div>
    </div>
  </div>
  {{template "home/common/footer.tpl" .}}
</body>

<script>
  layui.use('carousel', function () {
    var carousel = layui.carousel;
    carousel.render({
      elem: '#banner',
      width: '100%',
      arrow: 'always'
    });
  });
</script>

</html>