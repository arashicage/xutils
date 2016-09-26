{{define "navbar"}}

<body>
    <!-- navbar -->
    <div class="navbar navbar-default navbar-fixed-top">
        <div class="container">
            <div class="navbar-header">
                <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                </button>
                <a class="navbar-brand hidden-sm" href="/">登陆</a>
            </div>
            <div class="navbar-collapse collapse">
                <ul class="nav navbar-nav">
                    <li><a href="/"><strong> X-SYSTEM</strong></a>
                    </li>
                    <!--
-->
                    <li><a href="/config"><span class="glyphicon glyphicon-cog"></span> 系统设置</a>
                    </li>

<!--
                    <li class="dropdown">
                        <a href="#" class="dropdown-toggle" data-toggle="dropdown"><span class="glyphicon glyphicon-briefcase"></span> 工具<b class="caret"></b></a>
                        <ul class="dropdown-menu">
                            <li class="">
                                <a href="/utils/expcsv" class="glyphicon glyphicon-download-alt"> 导出<FONT face=黑体>CSV</FONT></a>
                            </li>
                            <li class="">
                                <a href="/utils/expsql" class="glyphicon glyphicon-import"> 导出<FONT face=黑体>SQL</FONT></a>
                            </li>
                        </ul>
                    </li>
-->

                    

                    <li><a href="/utils/expcsv"><span class="glyphicon glyphicon-download-alt"></span> 导出<FONT face=黑体>CSV</FONT></a></li>
                    <li><a href="/utils/expsql"><span class="glyphicon glyphicon-import"></span> 导出<FONT face=黑体>SQL</FONT></a></li>
					 <li><a href="/redis/load"><span class="glyphicon glyphicon-log-in"></span> 加载Redis</a></li>
<!--
                    <li><a href="/crud"><span class="glyphicon glyphicon-list-alt"></span> 系统监控</a>
                    </li>
                    <li><a href="/routine"><span class="glyphicon glyphicon-check"></span> 日常检查</a>
                    </li>
                    <li><a href="/qrcode"><span class="glyphicon glyphicon-qrcode"></span> 二维码</a>
                    </li>
-->
                    <li><a href="/"><span class="glyphicon glyphicon-info-sign"></span> 关于</a>
                    </li>
                </ul>
            </div>
        </div>
    </div>


    <div id="wrap">

        <!-- 巨幕 -->
        <div class="jumbotron">

            <div class="container">
                <p>
                    人之初，性本善，性相近，习相远。苟不教，性乃迁，教之道，贵以专。昔孟母，择邻处，子不学，断机杼。 <br>
                    窦燕山，有义方，教五子，名俱扬。养不教，父之过，教不严，师之惰。... <br>                    
                </p>            
<!--          <img src="static/img/12122522291291.jpg" >-->
            </div>

        </div>

        {{end}}