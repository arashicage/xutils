{{define "footer"}}

    </div> 
    <!-- Footer================================================== -->
    <footer class="bs-docs-footer" role="contentinfo">
      <div class="container" style="padding-left: 0px; padding-top: 10px; padding-bottom: 10px;">
        <div class="col-xs-8 col-md-8" style="width: 600px;">
          <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>

          <p>版权所有 © 2013-2014 长城计算机软件与系统有限公司</p>
          给我们发送 <span class="glyphicon glyphicon-envelope"></span> <a href="mailto:admin#gwssi.com.cn">反馈</a> 或者提交 <i class="icon-tasks"></i> <a target="_issue" href="/issues">网站问题</a>. 
        </div>

      </div>
    </footer>

    <!-- Bootstrap core JavaScript
    ================================================== -->
    <!-- Placed at the end of the document so the pages load faster 
    <script src="http://cdn.bootcss.com/jquery/1.11.1/jquery.min.js"></script>
    <script src="http://cdn.bootcss.com/bootstrap/3.1.1/js/bootstrap.min.js"></script>
    <script src="http://cdn.bootcss.com/unveil/1.3.0/jquery.unveil.min.js"></script>
    <link href="http://cdn.bootcss.com/highlight.js/8.0/styles/monokai.min.css" rel="stylesheet">
    <script src="http://cdn.bootcss.com/highlight.js/8.0/highlight.min.js"></script>
    <script>hljs.initHighlightingOnLoad();</script>
    <script src="http://static.bootcss.com/www/assets/js/jquery.scrollUp.min.js"></script>

	-->

	<script src="/static/js/jquery/1.11.1/jquery.min.js"></script>
   <!-- <script src="/static/js/jQueryFormTooltip/jquery.formtooltip.min.js"></script>-->
    <script src="/static/js/bootstrap/3.1.1/js/bootstrap.min.js"></script>
    <script src="/static/js/unveil/1.3.0/jquery.unveil.min.js"></script>
    <link href="/static/css/monokai.min.css" rel="stylesheet">
    <script src="/static/js/highlight.js/8.0/highlight.min.js"></script>
    <script>hljs.initHighlightingOnLoad();</script>
    <script src="/static/js/jquery.scrollUp.min.js"></script>
    <script src="/static/js/qrcode/jquery.qrcode-0.11.0.js"></script>

    <script>
      $(document).ready(function(){
        $("img.lazy").unveil();

        $.scrollUp({
              scrollName: 'scrollUp', // Element ID
              topDistance: '300', // Distance from top before showing element (px)
              topSpeed: 300, // Speed back to top (ms)
              animation: 'fade', // Fade, slide, none
              animationInSpeed: 200, // Animation in speed (ms)
              animationOutSpeed: 200, // Animation out speed (ms)
              scrollText: '', // Text for element
              activeOverlay: false  // Set CSS color to display scrollUp active point, e.g '#00FFFF'
        });
      });
    </script>

  </body>
</html>

{{end}}