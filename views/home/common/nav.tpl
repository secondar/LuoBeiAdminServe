<div class="nav">
  <div class="left">
    <a href="/">
      <img src="https://www.bugquit.com/wp-content/uploads/2019/12/logo.png" alt="home">
    </a>
  </div>
  <div class="right">
    <div class="mobile-menu" id="mobile-menu">
      <i class="layui-icon layui-icon-more-vertical"></i>
    </div>
    <ul class="layui-nav" id="layui-nav">
      <li class="layui-nav-item {{if ne "" (or .IsHome "")}}layui-this{{end}}"><a href="/">首页</a></li>
      <li class="layui-nav-item"><a href="https://github.com/secondar/LuoBeiAdminForBeegoAndVueElementUiAdmin/issues"
          target="_blank">讨论</a></li>
      <li class="layui-nav-item {{if eq "" (or .IsHome "")}}layui-this{{end}}"><a href="/article/0-1.html">动态</a></li>
      <li class="layui-nav-item"><a href="https://doc.imoecg.com" target="_blank">文档</a></li>
      <li class="layui-nav-item"><a href="https://www.bugquit.com" target="_blank">博客</a></li>
      <li class="layui-nav-item">
        <a href="javascript:;">项目</a>
        <dl class="layui-nav-child">
          <dd><a href="https://gitee.com/secondar/luo-bei-admin-for-beego-and-vue-element-ui-admin.git"
              target="_blank">GITEE</a></dd>
          <dd><a href="https://github.com/secondar/LuoBeiAdminForBeegoAndVueElementUiAdmin.git"
              target="_blank">GITHUB</a></dd>
        </dl>
      </li>
      <li class="layui-nav-item">
        <a href="javascript:;">仓库</a>
        <dl class="layui-nav-child">
          <dd><a href="https://gitee.com/secondar/luo-bei-admin-for-beego-and-vue-element-ui-admin.git"
              target="_blank">GITEE</a></dd>
          <dd><a href="https://github.com/secondar/LuoBeiAdminForBeegoAndVueElementUiAdmin.git"
              target="_blank">GITHUB</a></dd>
        </dl>
      </li>
    </ul>
  </div>
</div>
<script>
  layui.use(['element', 'jquery'], function () {
    var element = layui.element;
    let $ = layui.jquery
    let mobile_menu_show = false;
    $('#mobile-menu').click(() => {
      if (mobile_menu_show) {
        mobile_menu_show = false
        $("#layui-nav").hide();
      } else {
        mobile_menu_show = true
        $("#layui-nav").show();
      }
    })
  });
</script>