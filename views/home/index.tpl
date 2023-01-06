{{template "home/common/header.tpl" .}}

<body>
  {{template "home/common/nav.tpl" .}}
  <div class="home-main">
    <div class="banner">
      <div class="left">
        <div class="title">
          LuoBei <span>Admin</span> 社区
        </div>
        <div class="describe">
          致 力 于 让 " Web " 开 发 变 得 简 单 优 雅
        </div>
        <div class="btn">
          <a href="https://demo.imoecg.com" target="_blank" class="layui-btn demo"><i
              class="layui-icon layui-icon-template-1"></i>
            演示站点
          </a>
          <button type="button" class="layui-btn warehouse"><i class="layui-icon layui-icon-code-circle"></i>
            代码仓库
          </button>
        </div>
      </div>
      <div class="right">
        <img src="/static/home/img/home.png" alt="">
      </div>
    </div>
    <div class="row"></div>
    <div class="merit">
      <div class="layui-row">
        <div class="item layui-col-md6">
          <div class="head">
            <img src="/static/home/img/github.jpeg" alt="">
          </div>
          <div class="content">
            <div class="title">
              源码开放
            </div>
            <p>源码注释详细，便于阅读</p>
            <p>代码无后门，可放心使用</p>
          </div>
        </div>
        <div class="item layui-col-md6">
          <div class="head">
            <img src="/static/home/img/un2.webp" alt="">
          </div>
          <div class="content">
            <div class="title">
              易于开发
            </div>
            <p>基于常见的框架，文档齐全</p>
            <p>采用beego，element ui，mysql</p>
          </div>
        </div>
      </div>
      <div class="layui-row">
        <div class="item layui-col-md6">
          <div class="head">
            <img src="/static/home/img/un1.webp" alt="">
          </div>
          <div class="content">
            <div class="title">
              多端支持
            </div>
            <p>支持电脑端、智能手机</p>
            <p>windwos、linux、mac部署</p>
          </div>
        </div>
        <div class="item layui-col-md6">
          <div class="head">
            <img src="/static/home/img/un3.webp" alt="">
          </div>
          <div class="content">
            <div class="title">
              轻量框架
            </div>
            <p>没有过多的代码量,不限系统</p>
            <p>仅有权限管理,没有过多开发,易于定制</p>
          </div>
        </div>
      </div>
    </div>
  </div>
  {{template "home/common/footer.tpl" .}}
  <div class="show-messbox" style="padding: 5rem;display: none;">
    <a href="https://gitee.com/secondar/luo-bei-admin-for-beego-and-vue-element-ui-admin.git" target="_blank"
      class="layui-btn demo"><i class="layui-icon layui-icon-template-1"></i>
      GITEE
    </a>
    <a href="https://github.com/secondar/LuoBeiAdminForBeegoAndVueElementUiAdmin.git" target="_blank"
      class="layui-btn demo"><i class="layui-icon layui-icon-template-1"></i>
      GITHUB
    </a>
  </div>
</body>
<script>
  layui.use(['layer', 'jquery'], function () {
    var layer = layui.layer,
      $ = layui.jquery
    $(".warehouse").click(() => {

      layer.open({
        type: 1,
        title: '代码仓库',
        content: $(".show-messbox")
      });
    })
  });
</script>


</html>