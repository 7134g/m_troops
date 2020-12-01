package processing

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

var Str string = `<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 22pt;"><span style="tab-stops
: center 207.65pt;"><span lang="EN-US" style="font-size: 8.5pt;"><span style="font-family: 宋体;"><font color="#000000">&nbsp;</font></span></span></span></span></p>

<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 22pt;"><b><span style="font-size: 14pt;"><span style="font-family: 宋体;"><font color="#000000">长沙绕城高速西南段及机场高速2020-2022</font><font color="#000000">年养护工程施工</font></span></span></b></span></p>

<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 22pt;"><b><span style="font-size: 14pt;"><span style="font-family: 宋体;"><font color="#000000">招标公告</font></span></span></b></span></p>

<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 22pt;"><font color="#000000">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; </font></span></p>

<p style="margin: 12pt 0cm;"><span style="page-break-after: avoid;"><a name="_Toc470760973"></a><a name="_Toc517787476"><b><span style="font-size: 14pt;"><span style="font-family: 宋体;"><font color="#000000">1. 招标条件</font></span></span></b></a></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">本招标项目</font><u><font color="#000000">长沙绕城高速西南段及机场高速2020-2022</font><font color="#000000">年养护工程施工</font></u><font color="#000000">已由</font><u><font color="#000000">湖南投资集团股份有限公司和长沙环路建设开发集团有限公司批准列入年度养护计划</font></u><font color="#000000">，项目业主为</font><u><font color="#000000">湖南投资集团股份有限公司</font></u><font color="#000000">、</font><u><font color="#000000">长沙环路建设开发集团有限公司 </font></u><font color="#000000">，招标执行机构为湖南投资集团股份有限公司绕城公路西南段分公司、长沙市环路建设开发有限公司机场路分公司，建设资金来自</font><u><font color="#000000">自筹资金</font></u><font color="#000000">，出资比例为</font><u><font color="#000000"> 100%</font></u><font color="#000000">，招标人为</font><u><font color="#000000">湖南投资集团股份有限公司</font></u><font color="#000000">、</font><u><font color="#000000">长沙环路建设开发集团有限公司</font></u><font color="#000000">。项目已具备招标条件，现对该项目的施工进行公开招标。 </font></span></span></span></p>

<p style="margin: 12pt 0cm;"><span style="page-break-after: avoid;"><a name="_Toc470760974"></a><a name="_Toc517787477"><b><span style="font-size: 14pt;"><span style="font-family: 宋体;"><font color="#000000">2. 项目概况与招标范围</font></span></span></b></a></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><a name="_Toc470760975"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">2.1 </font></span></span></a><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">建设地点：绕城高速西南段及机场高速。</font></span></span></span></p>

<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="line-height: 20pt;"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">2.2 </font></span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">项目建设规模及招标范围：湖南投资集团股份有限公司的长沙绕城高速西南段起于长沙黄花塘与绕城高速的西北段相接，终点为京港澳高速的长沙李家塘收费站，全长28.091km</font><font color="#000000">，全线采用沥青混凝土路面，原路面结构为：</font><font color="#000000">4cm</font><font color="#000000">多碎石混合料</font><font color="#000000">SAC-13</font><font color="#000000">上面层</font><font color="#000000">+5cmAC-20I</font><font color="#000000">中粒式沥青砼</font><font color="#000000">+6cmAC-25I</font><font color="#000000">粗粒式沥青砼</font><font color="#000000">+1cm</font><font color="#000000">沥青表处封层</font><font color="#000000">+20cm6%</font><font color="#000000">水泥稳定碎石上基层</font><font color="#000000">+18cm</font><font color="#000000">水泥稳定砂砾下基层</font><font color="#000000">+15cm</font><font color="#000000">水泥石灰稳定砂砾土底基层。长沙绕城高速西南段于</font><font color="#000000">2002</font><font color="#000000">年</font><font color="#000000">10</font><font color="#000000">月建成通车。</font></span></span></span></p>

<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">长沙环路建设开发集团有限公司的长沙机场高速公路（湖南高速公路编号S40</font><font color="#000000">）为长沙城区连接黄花机场的快速通道，西起自京港澳高速长潭段雨花互通，东终点为黄花机场新跑道西侧，途经雨花区黎托街道、跨浏阳河大桥、长沙县榔梨镇、干杉乡和黄花镇，全长</font><font color="#000000">17.338</font><font color="#000000">公里，其中城区雨花区段长</font><font color="#000000">6.58</font><font color="#000000">公里。全线设三个互通（雨花互通、榔梨互通、长株互通由长株高速进行保养），两个收费站（长沙东、榔梨）。于</font><font color="#000000">2001</font><font color="#000000">年</font><font color="#000000">2</font><font color="#000000">月立项批准建设，</font><font color="#000000">2001</font><font color="#000000">年</font><font color="#000000">9</font><font color="#000000">月开工建设，于</font><font color="#000000">2003</font><font color="#000000">年</font><font color="#000000">9</font><font color="#000000">月</font><font color="#000000">3</font><font color="#000000">日正式通车。</font><font color="#000000">&nbsp; </font></span></span></span></p>

<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">本次招标范围为长沙绕城高速西南段及机场高速2020-2022</font><font color="#000000">年养护工程施工，共</font><font color="#000000">1</font><font color="#000000">个标段。</font><b><font color="#000000">其资质要求、业绩要求等见附表。</font></b></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">2.3</font></span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">本项目计划养护工程总工期36</font><font color="#000000">个月，本次招标中、小修保养项目养护工期从</font><font color="#000000">2020</font><font color="#000000">年</font><font color="#000000">1</font><font color="#000000">月</font><font color="#000000">1</font><font color="#000000">日起至</font><font color="#000000">2022</font><font color="#000000">年</font><font color="#000000">12</font><font color="#000000">月</font><font color="#000000">31</font><font color="#000000">日 。</font><b><font color="#000000">各项具体工程的计划工期以招标人或监理人的指令为准。缺陷责任期：</font></b><font color="#000000">养护工程自实际交工日期起计算</font><font color="#000000">2</font><font color="#000000">年。</font></span></span></span></p>

<p style="margin: 12pt 0cm;"><span style="page-break-after: avoid;"><a name="_Toc517787478"><b><span style="font-size: 14pt;"><span style="font-family: 宋体;"><font color="#000000">3. 投标人资格要求</font></span></span></b></a></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">3.1 </font></span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">本次招标要求投标人须具备<u> 独立法人资格，持有有效的营业执照，安全生产许可证，具备湖南省交通运输主管部门颁发的公路养护工程综合二类甲级资质、</u></font><font color="#000000">具有附件</font><font color="#000000">1</font><font color="#000000">附录</font><font color="#000000">3</font><font color="#000000">业绩，并在人员、设备、资金等方面具有相应的施工能力。投标人应进入交通运输部&ldquo;全国公路建设市场信用信息管理系统（</font><font color="#000000">http</font><font color="#000000">：</font><font color="#000000">//glxy.mot.gov.cn</font><font color="#000000">）&rdquo;中的公路工程施工资质企业名录。</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 21pt;"><a name="_Hlk2725481"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">3.2 </font></span></span></a><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">本次招标<u>不接受</u></font><font color="#000000">联合体投标。</font></span></span><a name="_Hlk2710711"></a><a name="_Hlk2725461"></a></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">3.3 </span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">与招标人存在利害关系且可能影响招标公正性的单位，不得参加投标。单位负责人</span></span></font><a href="#_ftn1" name="_ftnref1" title=""><sup><sup><span lang="EN-US" style="font-size: 10.5pt;"><span style="font-family: &quot;Calibri&quot;,&quot;sans-serif&quot;;"><font color="#0066cc">[1]</font></span></span></sup></sup></a><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">为同一人或者存在控股</font></span></span><a href="#_ftn2" name="_ftnref2" title=""><sup><sup><span lang="EN-US" style="font-size: 10.5pt;"><span style="font-family: &quot;Calibri&quot;,&quot;sans-serif&quot;;"><font color="#0066cc">[2]</font></span></span></sup></sup></a><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">、管理</font></span></span><a href="#_ftn3" name="_ftnref3" title=""><sup><sup><span lang="EN-US" style="font-size: 10.5pt;"><span style="font-family: &quot;Calibri&quot;,&quot;sans-serif&quot;;"><font color="#0066cc">[3]</font></span></span></sup></sup></a><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">关系的不同单位，不得参加同一标段投标，否则，相关投标均无效。</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">3.4</span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">与招标人存在利害关系且可能影响招标公正性的单位，不得参加投标。单位负责人为同一人或者存在控股、管理关系的不同单位，参与同一招标工程类别投标的标段总数量（含奖励增加数量）不得超过该招标工程类别的标段总数，否则，相关投标均无效。</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">3.5</font></span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">招标人不接受存在《湖南省公路工程标准施工招标文件》（2019</font><font color="#000000">年版）&ldquo;投标人须知&rdquo;第</font><font color="#000000">1.4.3</font><font color="#000000">项和第</font><font color="#000000">1.4.4</font><font color="#000000">项情形之一的投标人或被湖南省交通运输厅评为最近第一年度（</font><font color="#000000">2018</font><font color="#000000">年，下同）</font><font color="#000000">D</font><font color="#000000">级、连续两年（最近第二年和最近第一年）评为</font><font color="#000000">C</font><font color="#000000">级、连续三年（最近第三年～最近第一年）评为</font><font color="#000000">B</font><font color="#000000">级及以下信用等级的投标人报名。</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">3.6 </span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">招标人不接受</span></span></font><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">在全国企业信用信息公示系统（http://www.gsxt.gov.cn</font><font color="#000000">）中被列入严重违法失信企业名单的或在</font><font color="#000000">&ldquo;</font><font color="#000000">信用中国</font><font color="#000000">&rdquo;</font><font color="#000000">网站（</font><font color="#000000">www.creditchina.gov.cn</font><font color="#000000">）中被列入失信被执行人名单的投标人投标。</font></span></span></span></p>

<p style="margin: 12pt 0cm;"><span style="page-break-after: avoid;"><a name="_Toc470760976"></a><a name="_Toc517787479"><b><span style="font-size: 14pt;"><span style="font-family: 宋体;"><font color="#000000">4. 招标文件的获取</font></span></span></b></a></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">4.1</font></span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">根据规定，本项目只接受网上报名。凡有意参加投标者，请在<u>湖南省</u></font><font color="#000000">公共资源交易中心窗口办理</font><font color="#000000">CA</font><font color="#000000">认证后，从</font><font color="#000000">2019</font><font color="#000000">年 </font><font color="#000000">12</font><font color="#000000">月</font><font color="#000000">6</font><font color="#000000">日</font><font color="#000000">16</font><font color="#000000">时至</font><font color="#000000">2019</font><font color="#000000">年 </font><font color="#000000">2019</font><font color="#000000">月 </font><font color="#000000">13 </font><font color="#000000">日</font><font color="#000000">16</font><font color="#000000">时（北京时间，下同）自行在湖南省公共资源交易网下载</font><font color="#000000">/</font><font color="#000000">获取招标文件、图纸等相关资料，网上下载的招标文件与书面招标文件具有同等法律效力。凡资料不全、投标人与其证照不符、超过招标文件下载购买时间将被拒绝。投标人应及时关注网上相关招标信息，如有遗漏招标人概不负责，所造成的投标失败或损失由投标人自行负责。</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">4.2 </font></span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">招标文件每份售价人民币1000</font><font color="#000000">元，技术资料每份售价人民币</font><font color="#000000">1000</font><font color="#000000">元，一律现金支付，在递交投标文件时收取。</font></span></span></span></p>

<p style="margin: 12pt 0cm;"><span style="page-break-after: avoid;"><a name="_Toc517787480"></a><a name="_Toc470760977"><b><span style="font-size: 14pt;"><span style="font-family: 宋体;"><font color="#000000">5. 投标文件的递交及相关事宜</font></span></span></b></a></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">5.1</span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">招标人不组织进行工程现场踏勘，不召开投标预备会，有意参加投标者可自行到现场进行踏勘</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">5.2 </font></span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">投标人递交投标文件截止时间为 2019 </font><font color="#000000">年 </font><font color="#000000">12</font><font color="#000000">月 </font><font color="#000000">27</font><font color="#000000">日 </font><font color="#000000">10 </font><font color="#000000">时 </font><font color="#000000">00 </font><font color="#000000">分。投标人应于当日</font><u><font color="#000000"> 8&nbsp; </font></u><font color="#000000">时</font><u><font color="#000000"> 30 </font></u><font color="#000000">分至</font><u><font color="#000000"> 10 </font></u><font color="#000000">时</font><u><font color="#000000"> 00 </font></u><font color="#000000">分（投标人递交投标文件截止时间以本公告时间为准）将投标文件递交至</font><u><font color="#000000"> 长沙市万家丽南路2</font><font color="#000000">段</font><font color="#000000">29</font><font color="#000000">号湖南省公共资源交易中心开标室（具体开标室详见交易中心大厅屏幕）交招标人签收，</font></u><font color="#000000">招标人在投标截止的同一时间，同一地点举行开标仪式。</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">5.3 </span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">逾期送达的、未送达指定地点的或不按照招标文件要求密封的投标文件，招标人将予以拒收。</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">5.4 </span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">投标保证金的递交：</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">投标保证金的金额：<u>捌拾</u></font><font color="#000000">万元</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">（1</font><font color="#000000">）</font><a name="_Hlk12419643"></a><a name="_Hlk11941435"><font color="#000000">投标人采用现金或者支票形式提交的投标保证金应当从其基本账户转出，投标人应在投标截止时间前以转账、电汇、网银方式从投标人基本账户一次性划款到以下指定投标保证金专用账号上（以到账时间为准）</font></a><font color="#000000">。</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">户&nbsp;&nbsp;&nbsp; </font><font color="#000000">名：湖南省公共资源交易中心</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">开户银行：长沙银行湘府路支行</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">账&nbsp;&nbsp;&nbsp; </font><font color="#000000">号：</font><font color="#000000">607015379</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">（2</font><font color="#000000">）采用银行保函时，应由投标人开立基本账户的银行出具保函，与银行查询授权书原件一并在投标截止时间前交招标人。</font></span></span></span></p>

<p style="margin: 12pt 0cm;"><span style="page-break-after: avoid;"><a name="_Toc516774586"></a><a name="_Toc517377022"><b><span style="font-size: 14pt;"><span style="font-family: 宋体;"><font color="#000000">6．评标办法</font></span></span></b></a></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="punctuation-trim: leading;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">本项目评标办法采用<u> 综合评分法 </u></font><font color="#000000">。</font></span></span><a name="_Toc517787481"></a><a name="_Toc470760978"></a></span></span></p>

<p style="margin: 12pt 0cm;"><span style="page-break-after: avoid;"><font color="#000000"><b><span lang="EN-US" style="font-size: 14pt;"><span style="font-family: 宋体;">7. </span></span></b><b><span style="font-size: 14pt;"><span style="font-family: 宋体;">发布公告的媒介</span></span></b></font><a name="_Toc516774588"></a></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">本次招标公告同时在湖南省招标投标监管网（http://www.bidding.hunan.gov.cn/</font><font color="#000000">）、湖南省交通运输厅网（</font><font color="#000000">http://www.hnjt.gov.cn/</font><font color="#000000">）、湖南省公共资源交易中心网（</font><font color="#000000">http://ggzy.hunan.gov.cn/</font><font color="#000000">）上发布。</font></span></span></span></p>

<p style="margin: 12pt 0cm;"><span style="page-break-after: avoid;"><font color="#000000"><b><span lang="EN-US" style="font-size: 14pt;"><span style="font-family: 宋体;">8</span></span></b><b><span style="font-size: 14pt;"><span style="font-family: 宋体;">. 附件</span></span></b></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="punctuation-trim: leading;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">附件1</font><font color="#000000">：资格审查条件要求（详见第二章 投标人须知之附录）</font></span></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="punctuation-trim: leading;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">附件2</font><font color="#000000">：评标办法：</font><font color="#000000">(</font><font color="#000000">详见第三章 评标办法</font><font color="#000000">)</font></span></span></span></span></p>

<p style="margin: 12pt 0cm;"><span style="page-break-after: avoid;"><a name="_Toc517787482"></a><a name="_Toc470760979"><b><span style="font-size: 14pt;"><span style="font-family: 宋体;"><font color="#000000">10. 联系方式</font></span></span></b></a></span></p>

<p style="margin: 7.5pt 0cm 0pt;"><span style="line-height: 16.5pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">招 标 人：湖南投资集团股份有限公司</font></span></span></span></p>

<p style="margin: 7.5pt 0cm 0pt;"><span style="line-height: 16.5pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">招标执行机构：湖南投资集团股份有限公司绕城公路西南段分公司</font></span></span></span></p>

<p style="margin: 7.5pt 0cm 0pt;"><span style="line-height: 16.5pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">地&nbsp;&nbsp;</font><font color="#000000">址：</font></span></span><font color="#000000"> <span style="font-size: 12pt;"><span style="font-family: 宋体;">湖南省长沙市岳麓区黄花塘镇黄花塘收费站旁</span></span></font></span></p>

<p style="margin: 7.5pt 0cm 0pt;"><span style="line-height: 16.5pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">联 系 人：袁先生</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="punctuation-trim: leading;"><font color="#000000"><span style="font-size: 12pt;"><span style="font-family: 宋体;">电</span></span>&nbsp; <span style="font-size: 12pt;"><span style="font-family: 宋体;">话：</span></span><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: &quot;Times New Roman&quot;,&quot;serif&quot;;">0731-88508952&nbsp;&nbsp;&nbsp;&nbsp; </span></span></font></span></span></p>

<p style="margin: 7.5pt 0cm 0pt;"><span style="line-height: 16.5pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">招 标 人：长沙环路建设开发集团有限公司</font></span></span></span></p>

<p style="margin: 7.5pt 0cm 0pt;"><span style="line-height: 16.5pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">招标执行机构：长沙市环路建设开发有限公司机场路分公司</font></span></span></span></p>

<p style="margin: 7.5pt 0cm 0pt;"><span style="line-height: 16.5pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">地&nbsp;&nbsp;</font><font color="#000000">址：长沙市黄花机场</font></span></span></span></p>

<p style="margin: 7.5pt 0cm 0pt;"><span style="line-height: 16.5pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">联 系 人：刘先生</font></span></span></span></p>

<p style="margin: 7.5pt 0cm 0pt;"><span style="line-height: 16.5pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">电&nbsp;&nbsp;</font><font color="#000000">话：</font><font color="#000000">0731-86398406</font></span></span></span></p>

<p style="margin: 7.5pt 0cm 0pt;"><span style="line-height: 16.5pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">招标代理：湖南中誉项目管理有限公司</font></span></span></span></p>

<p style="margin: 7.5pt 0cm 0pt;"><span style="line-height: 16.5pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">地&nbsp;&nbsp;&nbsp; </font><font color="#000000">址：湖南省长沙市岳麓区奥克斯广场环球中心</font><font color="#000000">A</font><font color="#000000">座</font><font color="#000000">24008</font><font color="#000000">号</font></span></span></span></p>

<p style="margin: 7.5pt 0cm 0pt;"><span style="line-height: 16.5pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">联 系 人：李先生&nbsp;</font></span></span></span></p>

<p style="margin: 7.5pt 0cm 0pt;"><span style="line-height: 16.5pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">电&nbsp;&nbsp;&nbsp; </font><font color="#000000">话：</font><font color="#000000">0731-85718371&nbsp;&nbsp;&nbsp;&nbsp; </font></span></span></span></p>

<p style="margin: 7.5pt 0cm 0pt;"><span style="line-height: 16.5pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">监督部门：湖南省交通运输厅</font></span></span></span></p>

<p style="margin: 7.5pt 0cm 0pt;"><span style="line-height: 16.5pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">地&nbsp;&nbsp;&nbsp; </font><font color="#000000">址：长沙市湘府西路</font><font color="#000000">199</font><font color="#000000">号</font></span></span></span></p>

<p style="margin: 7.5pt 0cm 0pt;"><span style="line-height: 16.5pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">电&nbsp;&nbsp;&nbsp; </font><font color="#000000">话：</font><font color="#000000">0731-88770091</font></span></span></span></p>

<p style="margin: 7.5pt 0cm 0pt;"><span style="line-height: 16.5pt;"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">0731-88770122</font></span></span></span></p>

<p style="margin: 7.5pt 0cm 0pt;"><span style="line-height: 16.5pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">传&nbsp;&nbsp;&nbsp; </font><font color="#000000">真：</font><font color="#000000">0731-88770094</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">邮&nbsp;&nbsp;&nbsp; </font></font><font color="#000000"><font size="3">编：</font></font><font size="3"><font color="#000000">410004&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; </font></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><b><span lang="EN" style="font-size: 16pt;"><font color="#000000"><font face="Calibri">&nbsp;</font></font></span></b></p>

<p style="margin: 0cm 0cm 0pt;"><font color="#000000"><b><span style="font-size: 16pt;"><span style="font-family: 宋体;">附件</span></span></b><b><span lang="EN-US" style="font-size: 16pt;"><span style="font-family: &quot;Times New Roman&quot;,&quot;serif&quot;;">1</span></span></b><b><span style="font-size: 16pt;"><span style="font-family: 宋体;">：资格审查条件要求</span></span></b></font></p>

<p align="center" style="margin: 12pt 0cm; text-align: center;"><span style="page-break-after: avoid;"><a name="_Toc517787496"></a><a name="_Toc234832863"><b><span style="font-size: 14pt;"><span style="font-family: 宋体;"><font color="#000000">附录1&nbsp; </font><font color="#000000">资格审查条件（资质最低要求）</font></span></span></b></a></span></p>

<table align="center" style="border: 1pt solid windowtext; border-image: none; border-collapse: collapse;" width="588">
	<tbody>
		<tr style="height: 48.75pt; page-break-inside: avoid;">
			<td style="padding: 0cm 5.4pt; border: 1pt solid windowtext; border-image: none; width: 440.95pt; height: 48.75pt;" width="588">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 20pt;"><span style="-ms-layout-grid-mode: char;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">施工企业资质等级要求</font></font></span></span></span></p>
			</td>
		</tr>
		<tr style="height: 68.95pt; page-break-inside: avoid;">
			<td style="border-width: 0px 1pt 1pt; border-style: none solid solid; border-color: rgb(0, 0, 0) windowtext windowtext; padding: 0cm 5.4pt; width: 440.95pt; height: 68.95pt;" width="588">
			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="-ms-layout-grid-mode: char;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">具备独立法人资格、持有有效的营业执照、安全生产许可证、湖南省交通运输主管部门颁发的公路养护工程综合二类甲级资质。投标人应进入交通运输部&ldquo;全国公路建设市场信用信息管理系统（http</font></font><font color="#000000"><font size="3">：</font></font><font color="#000000"><font size="3">//glxy.mot.gov.cn</font></font><font color="#000000"><font size="3">）&rdquo;中的公路工程施工资质企业名录。</font></font></span></span></span></p>
			</td>
		</tr>
	</tbody>
</table>

<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="line-height: 125%;"><span style="-ms-layout-grid-mode: char;"><span style="line-height: 125%;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">注：</font></font></span></span><b><span style="line-height: 125%;"><span style="font-family: 黑体;"><font size="3"><font color="#000000">投标人应根据招标文件第二章&ldquo;投标人须知&rdquo;第3.5.1</font></font><font color="#000000"><font size="3">项的要求在&ldquo;投标人基本情况表&rdquo;后附相关证明材料。</font></font></span></span></b></span></span></p>

<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="line-height: 22pt;"><span lang="EN-US" style="font-size: 10pt;"><span style="font-family: &quot;Times New Roman&quot;,&quot;serif&quot;;"><font color="#000000">&nbsp; </font></span></span></span></p>

<p align="center" style="margin: 12pt 0cm; text-align: center;"><span style="page-break-after: avoid;"><a name="_Toc234832864"></a><a name="_Toc517787497"><b><span style="font-size: 14pt;"><span style="font-family: 宋体;"><font color="#000000">附录2&nbsp; </font><font color="#000000">资格审查条件（财务最低要求）</font></span></span></b></a></span></p>

<table align="center" style="border: 1pt solid windowtext; border-image: none; border-collapse: collapse;" width="565">
	<tbody>
		<tr style="height: 12.1pt; page-break-inside: avoid;">
			<td style="padding: 0cm 5.4pt; border: 1pt solid windowtext; border-image: none; width: 424pt; height: 12.1pt;" width="565">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 20pt;"><span style="-ms-layout-grid-mode: char;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">财 务 要 求</font></font></span></span></span></p>
			</td>
		</tr>
		<tr style="height: 12.1pt; page-break-inside: avoid;">
			<td style="border-width: 0px 1pt 1pt; border-style: none solid solid; border-color: rgb(0, 0, 0) windowtext windowtext; padding: 0cm 5.4pt; width: 424pt; height: 12.1pt;" width="565">
			<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="line-height: 200%;"><font color="#000000"><b><span lang="EN-US" style="font-size: 12pt;"><span style="line-height: 200%;"><font face="Calibri">① </font></span></span></b><b><span lang="EN-US" style="font-size: 12pt;"><span style="line-height: 200%;"><span style="font-family: 宋体;">2018</span></span></span></b></font><b><span style="font-size: 12pt;"><span style="line-height: 200%;"><span style="font-family: 宋体;"><font color="#000000">年净资产不少于<u>3000</u></font><font color="#000000">万元；</font></span></span></span></b></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="-ms-layout-grid-mode: char;"><font color="#000000"><b><span lang="EN-US" style="font-size: 12pt;"><font face="Calibri">② </font></span></b><b><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">2018</span></span></b></font><b><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">年资产负债率不高于75%</font><font color="#000000">。</font></span></span></b></span></span></p>
			</td>
		</tr>
	</tbody>
</table>

<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="line-height: 22pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">注：</font></font></span><b><span style="font-family: 黑体;"><font size="3"><font color="#000000">投标人应根据招标文件第二章&ldquo;投标人须知&rdquo;第3.5.2</font></font><font color="#000000"><font size="3">项的要求在&ldquo;财务状况表&rdquo;后附相关证明材料。</font></font></span></b></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 15pt;"><span style="-ms-layout-grid-mode: char;"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">&nbsp;</font></span></span></span></span></p>

<p align="center" style="margin: 12pt 0cm; text-align: center;"><span style="page-break-after: avoid;"><a name="_Toc234832865"></a><a name="_Toc509554633"></a><a name="_Toc471671302"><b><span style="font-size: 14pt;"><span style="font-family: 宋体;"><font color="#000000">附录3 资格审查条件（业绩最低要求）</font></span></span></b></a></span></p>

<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">&nbsp;</font></font></span></p>

<table style="border-collapse: collapse;" width="612">
	<thead>
		<tr style="height: 30.35pt;">
			<td style="border-width: 1.5pt 1.5pt 1pt 1pt; border-style: solid; border-color: windowtext; padding: 0cm 5.4pt; width: 459pt; height: 30.35pt;" width="612">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 22pt;"><span style="punctuation-trim: leading;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">业绩要求</font></font></span></span></span></p>
			</td>
		</tr>
	</thead>
	<tbody>
		<tr style="height: 1pt;">
			<td style="border-width: 0px 1.5pt 1.5pt 1pt; border-style: none solid solid; border-color: rgb(0, 0, 0) windowtext windowtext; padding: 0cm 5.4pt; width: 459pt; height: 1pt;" width="612">
			<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="line-height: 200%;"><font color="#000000"><b><span style="font-size: 12pt;"><span style="line-height: 200%;"><span style="font-family: 宋体;">最</span></span></span></b><b><span style="font-size: 12pt;"><span style="line-height: 200%;"><span style="font-family: 宋体;">近</span></span></span></b><b><span lang="EN-US" style="font-size: 12pt;"><span style="line-height: 200%;"><font face="Calibri">5</font></span></span></b><b><span style="font-size: 12pt;"><span style="line-height: 200%;"><span style="font-family: 宋体;">年（递交投标文件截止之日前一日回溯</span></span></span></b><b><span lang="EN-US" style="font-size: 12pt;"><span style="line-height: 200%;"><font face="Calibri">5</font></span></span></b><b><span style="font-size: 12pt;"><span style="line-height: 200%;"><span style="font-family: 宋体;">年）完成过：</span></span></span></b></font></span></p>

			<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="line-height: 200%;"><font color="#000000"><b><span lang="EN-US" style="font-size: 12pt;"><span style="line-height: 200%;"><font face="Calibri">①1</font></span></span></b><b><span style="font-size: 12pt;"><span style="line-height: 200%;"><span style="font-family: 宋体;">个高速公路的养护中修工程或改扩建工程或高速公路路面专项维修病害处治工程，且单个合同价格</span></span></span></b><b><span lang="EN-US" style="font-size: 12pt;"><span style="line-height: 200%;"><font face="Calibri">5000</font></span></span></b><b><span style="font-size: 12pt;"><span style="line-height: 200%;"><span style="font-family: 宋体;">万元及以上；</span></span></span></b></font></span></p>

			<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="line-height: 200%;"><span style="punctuation-trim: leading;"><b><span style="font-size: 12pt;"><span style="line-height: 200%;"><span style="font-family: 宋体;"><font color="#000000">②任一单年度完成的高速公路小修(</font><font color="#000000">含连接线</font><font color="#000000">)</font><font color="#000000">保养项目累计养护里程</font></span></span></span></b><font color="#000000"><sup><span style="font-size: 12pt;"><span style="background: white;"><span style="line-height: 200%;"><span style="font-family: 宋体;">④</span></span></span></span></sup><b><span lang="EN-US" style="font-size: 12pt;"><span style="line-height: 200%;"><span style="font-family: 宋体;">75</span></span></span></b><b><span style="font-size: 12pt;"><span style="line-height: 200%;"><span style="font-family: 宋体;">公里及以上。</span></span></span></b></font></span></span></p>
			</td>
		</tr>
	</tbody>
</table>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 22pt;"><font color="#000000"><span style="font-size: 12pt;"><span style="background: white;"><span style="font-family: 宋体;">注：</span></span></span><span lang="EN-US" style="font-size: 12pt;"><span style="background: white;"><span style="font-family: &quot;Times New Roman&quot;,&quot;serif&quot;;">1.</span></span></span><span style="font-size: 12pt;"><span style="background: white;"><span style="font-family: 宋体;">投标人应根据招标文件第二章&ldquo;投标人须知&rdquo;第</span></span></span><span lang="EN-US" style="font-size: 12pt;"><span style="background: white;"><span style="font-family: &quot;Times New Roman&quot;,&quot;serif&quot;;">3.5.3</span></span></span><span style="font-size: 12pt;"><span style="background: white;"><span style="font-family: 宋体;">项的要求在&ldquo;近年完成的类似项目情况表&rdquo;后附相关证明材料。</span></span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 22pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="background: white;"><span style="font-family: &quot;Times New Roman&quot;,&quot;serif&quot;;">2.</span></span></span><span style="font-size: 12pt;"><span style="background: white;"><span style="font-family: 宋体;">如近年来，投标人法人机构发生合法变更或重组或法人名称变更时，应提供相关部门的合法批件或其他相关证明材料来证明其所附业绩的继承</span></span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;">&nbsp;</p>

<p style="margin: 0cm 0cm 0pt;">&nbsp;</p>

<table style="border: rgb(0, 0, 0); border-image: none;">
	<tbody>
		<tr>
			<td height="0" style="border: rgb(0, 0, 0); border-image: none;" width="28">&nbsp;</td>
		</tr>
		<tr>
			<td style="border: rgb(0, 0, 0); border-image: none;">&nbsp;</td>
			<td style="border: rgb(0, 0, 0); border-image: none;"><font color="#000000"><font face="Calibri"><font size="3"><img src="file:///C:/Users/ADMINI~1.PCO/AppData/Local/Temp/msohtmlclip1/01/clip_image002.png" style="width: 354px; height: 39px;" /></font></font></font></td>
		</tr>
	</tbody>
</table>

<p style="margin: 0cm 0cm 0pt;">&nbsp;</p>

<p style="margin: 0cm 0cm 0pt;"><br clear="ALL" />
<span style="line-height: 22pt;"><font color="#000000"><sup><span style="font-size: 12pt;"><span style="background: white;"><span style="font-family: 宋体;">④</span></span></span></sup><span style="font-size: 12pt;"><span style="background: white;"><span style="font-family: 宋体;">任一单年度累计养护里程计算公式如下：</span></span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 22pt;"><font color="#000000"><font face="宋体"><font size="3"><img align="left" src="file:///C:/Users/ADMINI~1.PCO/AppData/Local/Temp/msohtmlclip1/01/clip_image004.png" style="width: 589px; height: 24px;" /></font></font></font><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">n </span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">&mdash;&mdash;任一单年度履行的高速公路小修（含连接线）保养项目合同数</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 22pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">例如：某养护单位小修保养合同养护时间为2016</font><font color="#000000">年</font><font color="#000000">7</font><font color="#000000">月</font><font color="#000000">1</font><font color="#000000">日至</font><font color="#000000">2018</font><font color="#000000">年</font><font color="#000000">9</font><font color="#000000">月</font><font color="#000000">30</font><font color="#000000">日，合同养护路段里程为</font><font color="#000000">100</font><font color="#000000">公里，若该单位参加投标截止时间为</font><font color="#000000">2018</font><font color="#000000">年</font><font color="#000000">9</font><font color="#000000">月</font><font color="#000000">1</font><font color="#000000">日的某项目投标，则已履行合同养护里程为</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 22pt;"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">2018</font></span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">年折算合同养护里程=8</font><font color="#000000">&divide;</font><font color="#000000">12</font><font color="#000000">&times;</font><font color="#000000">100=66.67</font><font color="#000000">公里</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 22pt;"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">2017</font></span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">年折算合同养护里程=12</font><font color="#000000">&divide;</font><font color="#000000">12</font><font color="#000000">&times;</font><font color="#000000">100=100</font><font color="#000000">公里</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 22pt;"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">2016</font></span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">年折算合同养护里程=6</font><font color="#000000">&divide;</font><font color="#000000">12</font><font color="#000000">&times;</font><font color="#000000">100=50</font><font color="#000000">公里</font></span></span></span></p>

<p align="center" style="margin: 12pt 0cm; text-align: center;"><span style="page-break-after: avoid;"><b><span style="font-size: 14pt;"><span style="font-family: 宋体;"><font color="#000000">附录4&nbsp; </font><font color="#000000">资格审查条件（信誉最低要求）</font></span></span></b></span></p>

<table align="center" style="border: 1pt solid windowtext; border-image: none; border-collapse: collapse;" width="588">
	<tbody>
		<tr style="height: 48.75pt; page-break-inside: avoid;">
			<td style="padding: 0cm 5.4pt; border: 1pt solid windowtext; border-image: none; width: 440.95pt; height: 48.75pt;" width="588">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 20pt;"><span style="-ms-layout-grid-mode: char;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">信 誉 要 求</font></font></span></span></span></p>
			</td>
		</tr>
		<tr style="height: 186.15pt; page-break-inside: avoid;">
			<td style="border-width: 0px 1pt 1pt; border-style: none solid solid; border-color: rgb(0, 0, 0) windowtext windowtext; padding: 0cm 5.4pt; width: 440.95pt; height: 186.15pt;" width="588">
			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="-ms-layout-grid-mode: char;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">投标人不得存在以下情形：</font></font></span></span></span></p>

			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 20pt;"><span style="-ms-layout-grid-mode: char;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">&nbsp;</font></font></span></span></span></p>

			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 20pt;"><span style="-ms-layout-grid-mode: char;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">被湖南省交通运输厅评为最近第一年度D</font></font><font color="#000000"><font size="3">级、连续两年（最近第二年和最近第一年）评为</font></font><font color="#000000"><font size="3">C</font></font><font color="#000000"><font size="3">级、连续三年（最近第三年～最近第一年）评为</font></font><font color="#000000"><font size="3">B</font></font><font color="#000000"><font size="3">级及以下信用等级。</font></font></span></span></span></p>
			</td>
		</tr>
	</tbody>
</table>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 15pt;"><span style="-ms-layout-grid-mode: char;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">注：1.</font></font><font color="#000000"><font size="3">投标人应根据招标文件第二章&ldquo;投标人须知&rdquo;前附表附录</font></font><font color="#000000"><font size="3">4</font></font><font color="#000000"><font size="3">和&ldquo;投标人须知&rdquo;正文第</font></font><font color="#000000"><font size="3">1.4.4</font></font><font color="#000000"><font size="3">项规定，逐条说明其信誉情况。</font></font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 15pt;"><span style="-ms-layout-grid-mode: char;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">2.</font></font></span><span style="font-family: 宋体;"><font size="3"><font color="#000000">投标人应根据招标文件第二章&ldquo;投标人须知&rdquo;第3.5.4</font></font><font color="#000000"><font size="3">项的要求在&ldquo;投标人的信誉情况表&rdquo;后附相关证明材料。</font></font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 125%;"><span style="-ms-layout-grid-mode: char;"><span lang="EN-US" style="line-height: 125%;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">&nbsp;</font></font></span></span></span></span></p>

<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 125%;"><span style="-ms-layout-grid-mode: char;"><b><span lang="EN-US" style="font-size: 12pt;"><span style="line-height: 125%;"><span style="font-family: 宋体;"><font color="#000000">&nbsp;</font></span></span></span></b></span></span></p>

<p align="center" style="margin: 12pt 0cm; text-align: center;"><span style="page-break-after: avoid;"><a name="_Toc517787500"></a><a name="_Toc234832867"><b><span style="font-size: 14pt;"><span style="font-family: 宋体;"><font color="#000000">附录5&nbsp; </font><font color="#000000">资格审查条件（项目经理和项目总工最低要求）</font></span></span></b></a></span></p>

<table align="center" style="border: 1pt solid windowtext; border-image: none; border-collapse: collapse;" width="600">
	<tbody>
		<tr style="height: 34pt; page-break-inside: avoid;">
			<td style="padding: 0cm 5.4pt; border: 1pt solid windowtext; border-image: none; width: 83.5pt; height: 34pt;" width="111">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 125%;"><span style="-ms-layout-grid-mode: char;"><span style="line-height: 125%;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">人&nbsp; </font></font><font color="#000000"><font size="3">员</font></font></span></span></span></span></p>
			</td>
			<td style="border-width: 1pt 1pt 1pt 0px; border-style: solid solid solid none; border-color: windowtext windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 45.25pt; height: 34pt;" width="60">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 125%;"><span style="-ms-layout-grid-mode: char;"><span style="line-height: 125%;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">数&nbsp; </font></font><font color="#000000"><font size="3">量</font></font></span></span></span></span></p>
			</td>
			<td style="border-width: 1pt 1pt 1pt 0px; border-style: solid solid solid none; border-color: windowtext windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 5cm; height: 34pt;" width="189">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 125%;"><span style="-ms-layout-grid-mode: char;"><span style="line-height: 125%;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">资 格 要 求</font></font></span></span></span></span></p>
			</td>
			<td style="border-width: 1pt 1pt 1pt 0px; border-style: solid solid solid none; border-color: windowtext windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 76.05pt; height: 34pt;" width="101">
			<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">工作经验要求</font></font></span></p>
			</td>
			<td style="border-width: 1pt 1pt 1pt 0px; border-style: solid solid solid none; border-color: windowtext windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 103.65pt; height: 34pt;" width="138">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 125%;"><span style="-ms-layout-grid-mode: char;"><span style="line-height: 125%;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">在岗要求</font></font></span></span></span></span></p>
			</td>
		</tr>
		<tr style="height: 105.25pt; page-break-inside: avoid;">
			<td style="border-width: 0px 1pt 1pt; border-style: none solid solid; border-color: rgb(0, 0, 0) windowtext windowtext; padding: 0cm 5.4pt; width: 83.5pt; height: 105.25pt;" width="111">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">项目经理</font></font></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 45.25pt; height: 105.25pt;" width="60">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">1</font></font></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 5cm; height: 105.25pt;" width="189">
			<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">具有壹级注册建造师执业资格（公路工程专业）、中级职称、省级及以上交通运输主管部门颁发的B</font></font><font color="#000000"><font size="3">类安全生产考核合格证书（编号为</font></font><font color="#000000"><font size="3">B</font></font><font color="#000000"><font size="3">类）。</font></font></span></p>
			</td>
			<td rowspan="2" style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 76.05pt; height: 105.25pt;" width="101">
			<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">最近5</font></font><font color="#000000"><font size="3">年（递交投标文件截止之日前一日回溯</font></font><font color="#000000"><font size="3">5</font></font><font color="#000000"><font size="3">年），担任过一个高速公路养护工程施工项目的项目经理或总工。</font></font></span></p>
			</td>
			<td rowspan="2" style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 103.65pt; height: 105.25pt;" width="138">
			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 125%;"><span style="-ms-layout-grid-mode: char;"><font size="3"><font color="#000000"><span style="line-height: 125%;"><span style="font-family: 宋体;">无在岗项目（指目前未在其他项目上任职，或虽在其他项目</span></span><span style="line-height: 125%;"><span style="font-family: 宋体;">上</span></span><span style="line-height: 125%;"><span style="font-family: 宋体;">任职但本项目中标后能够从该项目撤离）</span></span></font></font></span></span></p>
			</td>
		</tr>
		<tr style="height: 84.15pt; page-break-inside: avoid;">
			<td style="border-width: 0px 1pt 1pt; border-style: none solid solid; border-color: rgb(0, 0, 0) windowtext windowtext; padding: 0cm 5.4pt; width: 83.5pt; height: 84.15pt;" width="111">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">项目总工</font></font></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 45.25pt; height: 84.15pt;" width="60">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">1</font></font></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 5cm; height: 84.15pt;" width="189">
			<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">具有高级职称（公路工程相关专业）；省级及以上交通运输部门核发的安全生产考核合格证书（编号为B</font></font><font color="#000000"><font size="3">类）。</font></font></span></p>
			</td>
		</tr>
	</tbody>
</table>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 12pt;"><span style="-ms-layout-grid-mode: char;"><b><span lang="EN-US" style="font-size: 16pt;"><span style="font-family: 宋体;"><font color="#000000">&nbsp;</font></span></span></b></span></span></p>

<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="line-height: 22pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">注：</font></font><a name="_Toc234832868"><font color="#000000"><font size="3">1. </font></font><font color="#000000"><font size="3">投标人应根据招标文件第二章&ldquo;投标人须知&rdquo;第</font></font><font color="#000000"><font size="3">3.5.5</font></font><font color="#000000"><font size="3">项的要求在&ldquo;拟委任的项目经理、项目总工资历表&rdquo;后附相关证明材料。</font></font></a></span></span></p>

<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="line-height: 22pt;"><font size="3"><font color="#000000"><span lang="EN-US" style="font-family: 宋体;">2.</span><span style="font-family: 宋体;">项目经理、项目总工不得处于暂停执业处罚期内，且未被取消资格。</span></font></font></span></p>

<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="line-height: 22pt;"><font size="3"><font color="#000000"><span lang="EN-US" style="font-family: 宋体;">3.</span><span style="font-family: 宋体;">未在投标文件中填报的人员不作为评审依据。</span></font></font></span></p>

<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="line-height: 22pt;"><font size="3"><font color="#000000"><span lang="EN-US" style="font-family: 宋体;">4.</span><span style="font-family: 宋体;">投标人必须按投标文件中填报的项目经理、项目总工进场，否则招标人将取消投标人的中标资格。</span></font></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 26.25pt;"><a name="_Toc517789236"></a><a name="_Toc234832931"></a><a name="_Toc517787566"></a><b><span style="font-size: 16pt;"><font color="#000000"><font face="宋体">附件</font></font></span></b><font color="#000000"><b><span lang="EN-US" style="font-size: 16pt;"><span style="font-family: &quot;Times New Roman&quot;,&quot;serif&quot;;">2</span></span></b><b><span style="font-size: 16pt;"><font face="宋体">：评标办法</font></span></b></font></span></p>

<p align="center" style="margin: 24pt 0cm 12pt; text-align: center;"><span style="page-break-after: avoid;"><span style="font-size: 22pt;"><span style="font-family: 宋体;"><font color="#000000">评标办法（综合评分法）</font></span></span><a href="http://218.76.24.174:8090/G2/gbp/jgw-notice!add.do#_ftn4" name="_ftnref4" title=""><b><sup><span style="font-size: 21pt;"><span style="font-family: 宋体;"><b><sup><span style="font-size: 21pt;"><span style="font-family: 宋体;"><font color="#0066cc">[4]</font></span></span></sup></b></span></span></sup></b></a></span></p>

<p style="margin: 24pt 0cm 12pt;"><span style="line-height: 19pt;"><span style="page-break-after: avoid;"><a name="_Toc234832943"></a><a name="_Toc501257157"></a><a name="_Toc517789268"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">评标办法前附表</font></span></span></a><a href="http://218.76.24.174:8090/G2/gbp/jgw-notice!add.do#_ftn5" name="_ftnref5" title=""><b><sup><span style="font-size: 14pt;"><span style="font-family: 宋体;"><b><sup><span style="font-size: 14pt;"><span style="font-family: 宋体;"><font color="#0066cc">[5]</font></span></span></sup></b></span></span></sup></b></a></span></span></p>

<table style="border: 1pt solid windowtext; border-image: none; border-collapse: collapse;" width="593">
	<thead>
		<tr>
			<td colspan="2" style="padding: 0cm 5.4pt; border: 1pt solid windowtext; border-image: none; width: 113.45pt;" width="151">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><b><span style="font-family: 宋体;"><font color="#000000"><font size="3">条款号</font></font></span></b></span></p>
			</td>
			<td style="border-width: 1pt 1pt 1pt 0px; border-style: solid solid solid none; border-color: windowtext windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 331.1pt;" width="441">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><b><span style="letter-spacing: 0.2pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">评审因素与评审标准</font></font></span></span></b></span></p>
			</td>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td style="border-width: 0px 1pt 1pt; border-style: none solid solid; border-color: rgb(0, 0, 0) windowtext windowtext; padding: 0cm 5.4pt; width: 49.45pt;" width="66">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 20pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">1</font></font></span></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 64pt;" width="85">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 20pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">评标方法</font></font></span></span></p>
			</td>
			<td style="border-width: 0px 1pt 0px 0px; border-style: none solid none none; border-color: rgb(0, 0, 0) windowtext rgb(0, 0, 0) rgb(0, 0, 0); padding: 0cm 5.4pt; width: 331.1pt;" width="441">
			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 17pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">综合评分相等时，评标委员会依次按照以下优先顺序推荐中标候选人或确定中标人：</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 17pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（1</font></font><font color="#000000"><font size="3">）评标价低的投标人优先；</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（2</font></font><font color="#000000"><font size="3">）被湖南省交通运输厅评为最近</font></font><font color="#000000"><font size="3">1</font></font><font color="#000000"><font size="3">年较高信用等级的投标人优先；</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（3</font></font><font color="#000000"><font size="3">）商务和技术得分较高的投标人优先；</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 17pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（4</font></font><font color="#000000"><font size="3">）随机摇号确定排序。</font></font></span></span></p>
			</td>
		</tr>
		<tr>
			<td style="border-width: 0px 1pt 1pt; border-style: none solid solid; border-color: rgb(0, 0, 0) windowtext windowtext; padding: 0cm 5.4pt; width: 49.45pt;" width="66">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">2.1.1</font></font></span></span></p>

			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">2.1.3</font></font></span></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 64pt;" width="85">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">形式评审与响应性评审标准</font></font></span></span></p>
			</td>
			<td style="border-width: 1pt 1pt 1pt 0px; border-style: solid solid solid none; border-color: windowtext windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 331.1pt;" width="441">
			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 17pt;"><b><span style="font-family: 宋体;"><font color="#000000"><font size="3">第一个信封（商务及技术文件）评审标准：</font></font></span></b></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 17pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（1</font></font><font color="#000000"><font size="3">）投标文件按照招标文件规定的格式、内容填写，字迹清晰可辨：</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 17pt;"><font size="3"><font color="#000000"><span lang="EN-US" style="font-family: 宋体;">a.</span><span style="font-family: 宋体;">投标函按招标文件规定填报了项目名称、标段号、补遗书编号（如有）、工期、工程质量要求及安全目标、环保目标；</span></font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 17pt;"><font size="3"><font color="#000000"><span lang="EN-US" style="font-family: 宋体;">b.</span><span style="font-family: 宋体;">投标函附录的所有数据均符合招标文件规定；</span></font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 17pt;"><font size="3"><font color="#000000"><span lang="EN-US" style="font-family: 宋体;">c.</span><span style="font-family: 宋体;">投标文件组成齐全完整，内容均按规定填写。</span></font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 17pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（2</font></font><font color="#000000"><font size="3">）投标文件上法定代表人或其委托代理人的签字、投标人的单位章盖章齐全，符合招标文件规定。</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 17pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（3</font></font><font color="#000000"><font size="3">）与申请资格预审时比较，投标人发生合并、分立、破产等重大变化的，仍具备资格预审文件规定的相应资格条件且其投标未影响招标公正性：</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 17pt;"><font size="3"><font color="#000000"><span lang="EN-US" style="font-family: 宋体;">a.</span><span style="font-family: 宋体;">投标人应提供相关部门的合法批件及企业法人营业执照和资质证书等证件的副本变更记录复印件；</span></font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 17pt;"><font size="3"><font color="#000000"><span lang="EN-US" style="font-family: 宋体;">b.</span><span style="font-family: 宋体;">投标人仍然满足资格预审文件中规定的资格预审条件最低要求（资质、业绩、人员、信誉、财务等）；</span></font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 17pt;"><font size="3"><font color="#000000"><span lang="EN-US" style="font-family: 宋体;">c.</span><span style="font-family: 宋体;">与所投标段的其他投标人不存在控股、管理关系或单位负责人为同一人的情况；与招标人也不存在利害关系并可能影响招标公正性。</span></font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 17pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（4</font></font><font color="#000000"><font size="3">）投标人按照招标文件的规定提供了投标保证金：</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 17pt;"><font size="3"><font color="#000000"><span lang="EN-US" style="font-family: 宋体;">a.</span><span style="font-family: 宋体;">投标保证金金额符合招标文件规定的金额，且投标保证金有效期应当与投标有效期一致；</span></font></font></span></p>
			</td>
		</tr>
	</tbody>
</table>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">&nbsp;</font></font></span></span></p>

<table style="border: 1pt solid windowtext; border-image: none; border-collapse: collapse;" width="593">
	<thead>
		<tr>
			<td colspan="2" style="padding: 0cm 5.4pt; border: 1pt solid windowtext; border-image: none; width: 113.45pt;" width="151">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><b><span style="font-family: 宋体;"><font color="#000000"><font size="3">条款号</font></font></span></b></span></p>
			</td>
			<td style="border-width: 1pt 1pt 1pt 0px; border-style: solid solid solid none; border-color: windowtext windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 331.1pt;" width="441">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><b><span style="letter-spacing: 0.2pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">评审因素与评审标准</font></font></span></span></b></span></p>
			</td>
		</tr>
	</thead>
	<thead>
		<tr>
			<td colspan="2" style="border-width: 0px 1pt 1pt; border-style: none solid solid; border-color: rgb(0, 0, 0) windowtext windowtext; padding: 0cm 5.4pt; width: 113.45pt;" width="151">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><b><span style="font-family: 宋体;"><font color="#000000"><font size="3">条款号</font></font></span></b></span></p>
			</td>
			<td style="border-width: 1pt 1pt 1pt 0px; border-style: solid solid solid none; border-color: windowtext windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 331.1pt;" width="441">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><b><span style="letter-spacing: 0.2pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">评审因素与评审标准</font></font></span></span></b></span></p>
			</td>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td style="border-width: 0px 1pt 1pt; border-style: none solid solid; border-color: rgb(0, 0, 0) windowtext windowtext; padding: 0cm 5.4pt; width: 49.45pt;" width="66">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">2.1.1</font></font></span></span></p>

			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">2.1.3</font></font></span></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 64pt;" width="85">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">形式评审与响应性评审标准</font></font></span></span></p>
			</td>
			<td style="border-width: 0px 1pt 0px 0px; border-style: none solid none none; border-color: rgb(0, 0, 0) windowtext rgb(0, 0, 0) rgb(0, 0, 0); padding: 0cm 5.4pt; width: 331.1pt;" width="441">
			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><font size="3"><font color="#000000"><span lang="EN-US" style="font-family: 宋体;">b.</span><span style="font-family: 宋体;">若投标保证金采用现金或支票形式提交，投标人应在递交投标文件截止时间之前，将投标保证金由投标人的基本账户转入招标人指定账户；</span></font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><font size="3"><font color="#000000"><span lang="EN-US" style="font-family: 宋体;">c.</span><span style="font-family: 宋体;">若投标保证金采用银行保函形式提交，银行保函的格式、开具保函的银行均满足招标文件要求，且在递交投标文件截止时间之前向招标人提交了银行保函原件。</span></font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（5</font></font><font color="#000000"><font size="3">）投标人法定代表人授权委托代理人签署投标文件的，须提交授权委托书，且授权人和被授权人均在授权委托书上签名，未使用印章、签名章或其他电子制版签名代替。</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（6</font></font><font color="#000000"><font size="3">）投标人法定代表人亲自签署投标文件的，提供了法定代表人身份证明，且法定代表人在法定代表人身份证明上签名，未使用印章、签名章或其他电子制版签名代替。</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（7</font></font><font color="#000000"><font size="3">）投标人以联合体形式投标时，联合体满足招标文件的要求：</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><font size="3"><font color="#000000"><span lang="EN-US" style="font-family: 宋体;">a.</span><span style="font-family: 宋体;">未进行资格预审的，投标人按照招标文件提供的格式签订了联合体协议书，</span><span style="font-family: 宋体;">明确各方承担连带责任，</span><span style="font-family: 宋体;">并明确了联合体牵头人；</span></font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><font size="3"><font color="#000000"><span lang="EN-US" style="font-family: 宋体;">b.</span><span style="font-family: 宋体;">已进行资格预审的，投标人提供了资格预审申请文件中所附的联合体协议书复印件，且通过资格预审后的联合体无成员增减或更换的情况。</span></font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（8</font></font><font color="#000000"><font size="3">）投标人如有分包计划，符合招标文件第二章</font></font><font color="#000000"><font size="3">&ldquo;</font></font><font color="#000000"><font size="3">投标人须知</font></font><font color="#000000"><font size="3">&rdquo;</font></font><font color="#000000"><font size="3">第</font></font><font color="#000000"><font size="3">1.11</font></font><font color="#000000"><font size="3">款规定，且按招标文件第九章</font></font><font color="#000000"><font size="3">&ldquo;</font></font><font color="#000000"><font size="3">投标文件格式</font></font><font color="#000000"><font size="3">&rdquo;</font></font><font color="#000000"><font size="3">的要求填写了</font></font><font color="#000000"><font size="3">&ldquo;</font></font><font color="#000000"><font size="3">拟分包项目情况表</font></font><font color="#000000"><font size="3">&rdquo;</font></font><font color="#000000"><font size="3">。</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（9</font></font><font color="#000000"><font size="3">）同一投标人未提交两个以上不同的投标文件，但招标文件要求提交备选投标的除外。</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（10</font></font><font color="#000000"><font size="3">）投标文件中未出现有关投标报价的内容。</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（11</font></font><font color="#000000"><font size="3">）投标文件载明的招标项目完成期限未超过招标文件规定的时限。</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（12</font></font><font color="#000000"><font size="3">）投标文件对招标文件的实质性要求和条件作出响应。</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（13</font></font><font color="#000000"><font size="3">）权利义务符合招标文件规定：</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><font size="3"><font color="#000000"><span lang="EN-US" style="font-family: 宋体;">a.</span><span style="font-family: 宋体;">投标人应接受招标文件规定的风险划分原则，未提出新的风险划分办法；</span></font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><font size="3"><font color="#000000"><span lang="EN-US" style="font-family: 宋体;">b.</span><span style="font-family: 宋体;">投标人未增加发包人的责任范围，或减少投标人义务；</span></font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><font size="3"><font color="#000000"><span lang="EN-US" style="font-family: 宋体;">c.</span><span style="font-family: 宋体;">投标人未提出不同的工程验收、计量、支付办法；</span></font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><font size="3"><font color="#000000"><span lang="EN-US" style="font-family: 宋体;">d.</span><span style="font-family: 宋体;">投标人对合同纠纷、事故处理办法未提出异议；</span></font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><font size="3"><font color="#000000"><span lang="EN-US" style="font-family: 宋体;">e.</span><span style="font-family: 宋体;">投标人在投标活动中无欺诈行为；</span></font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><font size="3"><font color="#000000"><span lang="EN-US" style="font-family: 宋体;">f.</span><span style="font-family: 宋体;">投标人未对合同条款有重要保留。</span></font></font></span></p>
			</td>
		</tr>
	</tbody>
	<tbody>
		<tr>
			<td style="border-width: 0px 1pt; border-style: none solid; border-color: rgb(0, 0, 0) windowtext; padding: 0cm 5.4pt; width: 49.45pt;" width="66">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">2.1.1</font></font></span></span></p>

			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">2.1.3</font></font></span></span></p>
			</td>
			<td style="border-width: 0px 1pt 0px 0px; border-style: none solid none none; border-color: rgb(0, 0, 0) windowtext rgb(0, 0, 0) rgb(0, 0, 0); padding: 0cm 5.4pt; width: 64pt;" width="85">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">形式评审与响应性评审标准</font></font></span></span></p>
			</td>
			<td style="border-width: 0px 1pt 0px 0px; border-style: none solid none none; border-color: rgb(0, 0, 0) windowtext rgb(0, 0, 0) rgb(0, 0, 0); padding: 0cm 5.4pt; width: 331.1pt;" width="441">
			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（14</font></font><font color="#000000"><font size="3">）投标文件正、副本份数符合招标文件第二章</font></font><font color="#000000"><font size="3">&ldquo;</font></font><font color="#000000"><font size="3">投标人须知</font></font><font color="#000000"><font size="3">&rdquo;</font></font><font color="#000000"><font size="3">第</font></font><font color="#000000"><font size="3">3.7.4</font></font><font color="#000000"><font size="3">项规定。</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">&nbsp;</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><b><span style="font-family: 宋体;"><font color="#000000"><font size="3">第二个信封（报价文件）评审标准：</font></font></span></b></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（1</font></font><font color="#000000"><font size="3">）投标文件按照招标文件规定的格式、内容填写，字迹清晰可辨：</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><font size="3"><font color="#000000"><span lang="EN-US" style="font-family: 宋体;">a.</span><span style="font-family: 宋体;">投标函按招标文件规定填报了项目名称、标段号、补遗书编号（如有）、投标价（包括大写金额和小写金额）；</span></font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><font size="3"><font color="#000000"><span lang="EN-US" style="font-family: 宋体;">b.</span><span style="font-family: 宋体;">已标价工程量清单说明文字与招标文件规定一致，未进行实质性修改和删减；</span></font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><font size="3"><font color="#000000"><span lang="EN-US" style="font-family: 宋体;">c.</span><span style="font-family: 宋体;">投标文件组成齐全完整，内容均按规定填写。</span></font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（2</font></font><font color="#000000"><font size="3">）投标文件上法定代表人或其委托代理人的签字、投标人的单位章盖章齐全，符合招标文件规定。</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（3</font></font><font color="#000000"><font size="3">）投标报价或调价函中的报价未超过招标文件设定的最高投标限价（如有）。</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（4</font></font><font color="#000000"><font size="3">）投标报价或调价函中报价的大写金额能够确定具体数值。</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（5</font></font><font color="#000000"><font size="3">）同一投标人未提交两个以上不同的投标报价，但招标文件要求提交备选投标的除外。</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（6</font></font><font color="#000000"><font size="3">）投标人若提交调价函，调价函符合招标文件第二章</font></font><font color="#000000"><font size="3">&ldquo;</font></font><font color="#000000"><font size="3">投标人须知</font></font><font color="#000000"><font size="3">&rdquo;</font></font><font color="#000000"><font size="3">第</font></font><font color="#000000"><font size="3">3.2.6</font></font><font color="#000000"><font size="3">项要求。</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（7</font></font><font color="#000000"><font size="3">）投标人若填写工程量固化清单，填写完毕的工程量固化清单未对工程量固化清单电子文件中的数据、格式和运算定义进行修改；工程量固化清单中的投标报价和投标函大写金额报价一致。</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（8</font></font><font color="#000000"><font size="3">）投标文件正、副本份数符合招标文件第二章</font></font><font color="#000000"><font size="3">&ldquo;</font></font><font color="#000000"><font size="3">投标人须知</font></font><font color="#000000"><font size="3">&rdquo;</font></font><font color="#000000"><font size="3">第</font></font><font color="#000000"><font size="3">3.7.4</font></font><font color="#000000"><font size="3">项规定。</font></font></span></span></p>
			</td>
		</tr>
		<tr>
			<td style="padding: 0cm 5.4pt; border: 1pt solid windowtext; border-image: none; width: 49.45pt;" width="66">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">2.1.2</font></font></span></span></p>
			</td>
			<td style="border-width: 1pt 1pt 1pt 0px; border-style: solid solid solid none; border-color: windowtext windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 64pt;" width="85">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">资格评审</font></font></span></span></p>

			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">标准</font></font></span></span></p>
			</td>
			<td style="border-width: 1pt 1pt 1pt 0px; border-style: solid solid solid none; border-color: windowtext windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 331.1pt;" width="441">
			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（1</font></font><font color="#000000"><font size="3">）投标人具备有效的营业执照、</font></font></span><font size="3"><font color="#000000"><span style="font-family: 宋体;">组织机构代码证、</span><span style="font-family: 宋体;">资质证书、安全生产许可证和基本账户开户许可证。 </span></font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（2</font></font><font color="#000000"><font size="3">）投标人的资质等级符合招标文件规定。</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（3</font></font><font color="#000000"><font size="3">）投标人的财务状况符合招标文件规定。</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（4</font></font><font color="#000000"><font size="3">）投标人的类似项目业绩符合招标文件规定。</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（5</font></font><font color="#000000"><font size="3">）投标人的信誉符合招标文件规定。</font></font></span></span></p>
			</td>
		</tr>
	</tbody>
</table>

<p>&nbsp;</p>

<p style="margin: 0cm 0cm 0pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">&nbsp;</font></font></span></p>

<table style="border: 1pt solid windowtext; border-image: none; border-collapse: collapse;" width="593">
	<thead>
		<tr>
			<td colspan="2" style="border-width: 0px 0px 1pt; border-style: none none solid; border-color: rgb(0, 0, 0) rgb(0, 0, 0) windowtext; padding: 0cm 5.4pt; width: 140.1pt;" width="187">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><b><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">&nbsp;</font></font></span></b></span></p>
			</td>
			<td style="border-width: 0px 0px 1pt; border-style: none none solid; border-color: rgb(0, 0, 0) rgb(0, 0, 0) windowtext; padding: 0cm 5.4pt; width: 304.45pt;" width="406">
			<p align="right" style="margin: 0cm 0cm 0pt; text-align: right;"><span style="line-height: 19pt;"><span style="letter-spacing: 0.2pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">续上表</font></font></span></span></span></p>
			</td>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td style="border-width: 0px 1pt; border-style: none solid; border-color: rgb(0, 0, 0) windowtext; padding: 0cm 5.4pt; width: 49.45pt;" width="66">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">2.1.2</font></font></span></span></p>
			</td>
			<td style="border-width: 0px 1pt 0px 0px; border-style: none solid none none; border-color: rgb(0, 0, 0) windowtext rgb(0, 0, 0) rgb(0, 0, 0); padding: 0cm 5.4pt; width: 90.65pt;" width="121">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">资格评审</font></font></span></span></p>

			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">标准</font></font></span></span></p>
			</td>
			<td style="border-width: 0px 1pt 0px 0px; border-style: none solid none none; border-color: rgb(0, 0, 0) windowtext rgb(0, 0, 0) rgb(0, 0, 0); padding: 0cm 5.4pt; width: 304.45pt;" width="406">
			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（6</font></font><font color="#000000"><font size="3">）投标人的项目经理和项目总工资格、在岗情况符合招标文件规定。</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（7</font></font><font color="#000000"><font size="3">）投标人的其他要求符合招标文件规定。</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（8</font></font><font color="#000000"><font size="3">）投标人不存在第二章</font></font><font color="#000000"><font size="3">&ldquo;</font></font><font color="#000000"><font size="3">投标人须知</font></font><font color="#000000"><font size="3">&rdquo;</font></font><font color="#000000"><font size="3">第</font></font><font color="#000000"><font size="3">1.4.3</font></font><font color="#000000"><font size="3">项或第</font></font><font color="#000000"><font size="3">1.4.4</font></font><font color="#000000"><font size="3">项规定的任何一种情形。</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（9</font></font><font color="#000000"><font size="3">）投标人符合第二章</font></font><font color="#000000"><font size="3">&ldquo;</font></font><font color="#000000"><font size="3">投标人须知</font></font><font color="#000000"><font size="3">&rdquo;</font></font><font color="#000000"><font size="3">第</font></font><font color="#000000"><font size="3">1.4.5</font></font><font color="#000000"><font size="3">项规定。</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（10</font></font><font color="#000000"><font size="3">）以联合体形式参与投标的，联合体各方均未再以自己名义单独或参加其他联合体在同一标段中投标；独立参与投标的，投标人未同时参加联合体在同一标段中投标。</font></font></span></span></p>
			</td>
		</tr>
		<tr>
			<td style="border-width: 1pt 1pt 0px; border-style: solid solid none; border-color: windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 49.45pt;" width="66">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><b><span style="font-family: 宋体;"><font color="#000000"><font size="3">条款号</font></font></span></b></span></p>
			</td>
			<td style="border-width: 1pt 1pt 0px 0px; border-style: solid solid none none; border-color: windowtext windowtext rgb(0, 0, 0) rgb(0, 0, 0); padding: 0cm 5.4pt; width: 90.65pt;" width="121">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><b><span style="font-family: 宋体;"><font color="#000000"><font size="3">条款内容</font></font></span></b></span></p>
			</td>
			<td style="border-width: 1pt 1pt 0px 0px; border-style: solid solid none none; border-color: windowtext windowtext rgb(0, 0, 0) rgb(0, 0, 0); padding: 0cm 5.4pt; width: 304.45pt;" width="406">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><b><span style="font-family: 宋体;"><font color="#000000"><font size="3">编列内容</font></font></span></b></span></p>
			</td>
		</tr>
		<tr>
			<td style="padding: 0cm 5.4pt; border: 1pt solid windowtext; border-image: none; width: 49.45pt;" width="66">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">2.2.1</font></font></span></span></p>
			</td>
			<td style="border-width: 1pt 1pt 1pt 0px; border-style: solid solid solid none; border-color: windowtext windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 90.65pt;" width="121">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">分值构成</font></font></span></span></p>

			<p align="center" style="margin: 0cm -5.25pt 0pt; text-align: center;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（总分100</font></font><font color="#000000"><font size="3">分）</font></font></span></span></p>
			</td>
			<td style="border-width: 1pt 1pt 1pt 0px; border-style: solid solid solid none; border-color: windowtext windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 304.45pt;" width="406">
			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><b><span style="font-family: 宋体;"><font color="#000000"><font size="3">第一个信封（商务及技术文件）评分分值构成：</font></font></span></b></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">施工组织设计：<u> 16 </u></font></font><font color="#000000"><font size="3">分（</font></font><font color="#000000"><font size="3">5~16</font></font><font color="#000000"><font size="3">分）</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">主要人员：<u> 10 </u></font></font><font color="#000000"><font size="3">分（</font></font><font color="#000000"><font size="3">5~10</font></font><font color="#000000"><font size="3">分）</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">技术能力：<u> 0 </u></font></font><font color="#000000"><font size="3">分（</font></font><font color="#000000"><font size="3">0~5</font></font><font color="#000000"><font size="3">分）</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">财务能力：</font></font><u><font color="#000000"><font size="3">&nbsp; </font></font><font color="#000000"><font size="3">6 </font></font></u><font color="#000000"><font size="3">分（</font></font><font color="#000000"><font size="3">3~6</font></font><font color="#000000"><font size="3">分）</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">业绩：</font></font><u><font color="#000000"><font size="3">&nbsp; </font></font><font color="#000000"><font size="3">8</font></font></u><font color="#000000"><font size="3">分（</font></font><font color="#000000"><font size="3">3~8</font></font><font color="#000000"><font size="3">分）</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">综合实力：<u> 5 </u></font></font><font color="#000000"><font size="3">分（</font></font><font color="#000000"><font size="3">3~5</font></font><font color="#000000"><font size="3">分）</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><b><span style="font-family: 宋体;"><font color="#000000"><font size="3">第二个信封（报价文件）评分分值构成：</font></font></span></b></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">评标价：<u> 55</u></font></font><font color="#000000"><font size="3">分</font></font></span></span></p>
			</td>
		</tr>
		<tr>
			<td style="border-width: 0px 1pt 1pt; border-style: none solid solid; border-color: rgb(0, 0, 0) windowtext windowtext; padding: 0cm 5.4pt; width: 49.45pt;" width="66">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">2.2.2</font></font></span></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 90.65pt;" width="121">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">评标基准价计算方法</font></font></span></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 304.45pt;" width="406">
			<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">评标基准价计算方法： </font></font></span></span></p>

			<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">在开标现场，招标人将当场计算并宣布评标基准价。</font></font></span></span></p>

			<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="line-height: 18pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">(1)</font></font></span><span style="font-family: 宋体;"><font size="3"><font color="#000000">评标价的确定：评标价=</font></font><font color="#000000"><font size="3">投标函文字报价</font></font></span></span></p>

			<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="line-height: 18pt;"><font size="3"><font color="#000000"><span lang="EN-US" style="font-family: 宋体;">(2) </span><span style="font-family: 宋体;">理论成本价的确定:</span></font></font></span></p>

			<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">理论成本价＝（最高投标限价&times;60</font></font><font color="#000000"><font size="3">％＋除按第二章&ldquo;投标人须知&rdquo;第</font></font><font color="#000000"><font size="3">5.2.4</font></font><font color="#000000"><font size="3">项规定开标现场被宣布为不进入理论成本价计算的投标报价之外的所有投标人的评标价（或去掉一个最高和一个最低值）的算术平均值&times;</font></font><font color="#000000"><font size="3">40</font></font><font color="#000000"><font size="3">％）&times;</font></font><font color="#000000"><font size="3">0.88</font></font><font color="#000000"><font size="3">（如果参与理论成本价平均值计算的有效投标人少于</font></font><font color="#000000"><font size="3">5</font></font><font color="#000000"><font size="3">家时，则计算理论成本价平均值时不去掉最高值和最低值）</font></font></span></span></p>

			<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="line-height: 18pt;"><font size="3"><font color="#000000"><span lang="EN-US" style="font-family: 宋体;">(3)</span><span style="font-family: 宋体;">评标价平均值的计算：</span></font></font></span></p>

			<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">除按第二章&ldquo;投标人须知&rdquo;第5.2.4</font></font><font color="#000000"><font size="3">项规定开标现场被宣布为不进入评标基准价计算的投标报价之外，所有投标人的评标价去掉一个最高值和一个最低值后的算术平均值即为评标价平均值（如果参与评标价平均值计算的有效投标人少于</font></font><font color="#000000"><font size="3">5</font></font><font color="#000000"><font size="3">家时，则计算评标价平均值时不去掉最高值和最低值）。</font></font></span></span></p>

			<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="line-height: 18pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">&nbsp;</font></font><font color="#000000"><font size="3">(4)</font></font></span><span style="font-family: 宋体;"><font size="3"><font color="#000000">评标基准价的确定:</font></font></span></span></p>

			<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">评标基准价=</font></font><font color="#000000"><font size="3">（最高上限价&times;</font></font><font color="#000000"><font size="3">0.6+</font></font><font color="#000000"><font size="3">评标价平均值&times;</font></font><font color="#000000"><font size="3">0.4</font></font><font color="#000000"><font size="3">）&times;（</font></font><font color="#000000"><font size="3">1-</font></font><font color="#000000"><font size="3">下浮系数），下浮系数将从</font></font><font color="#000000"><font size="3">1%</font></font><font color="#000000"><font size="3">、</font></font><font color="#000000"><font size="3">1.5%</font></font><font color="#000000"><font size="3">、</font></font><font color="#000000"><font size="3">2%</font></font><font color="#000000"><font size="3">、</font></font><font color="#000000"><font size="3">2.5%</font></font><font color="#000000"><font size="3">、</font></font><font color="#000000"><font size="3">3%</font></font><font color="#000000"><font size="3">中在第二信封开标现场随机抽取。</font></font></span></span></p>

			<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="line-height: 18pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">&nbsp;</font></font></span></span></p>

			<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">如果投标人认为评标基准价计算有误， 有权在开标现场提出， 经当场核实确认之后，可重新宣布评标基准价。</font></font></span></span></p>

			<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="line-height: 18pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">&nbsp;</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">在评标过程中，评标委员会应对招标人计算的评标基准价进行复核，存在计算错误的应予以修正并在评标报告中作出说明。除此之外，评标基准价在整个评标期间保持不变，不随任何因素发生变化。</font></font></span></span></p>
			</td>
		</tr>
		<tr>
			<td style="border-width: 0px 1pt 1pt; border-style: none solid solid; border-color: rgb(0, 0, 0) windowtext windowtext; padding: 0cm 5.4pt; width: 49.45pt;" width="66">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">2.2.3</font></font></span></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 90.65pt;" width="121">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">评标价的偏差率计算公式</font></font></span></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 304.45pt;" width="406">
			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">偏差率=100% &times;</font></font><font color="#000000"><font size="3">（投标人评标价－评标基准价）</font></font><font color="#000000"><font size="3">/</font></font><font color="#000000"><font size="3">评标基准价</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">偏差率保留4</font></font><font color="#000000"><font size="3">位小数</font></font></span></span></p>
			</td>
		</tr>
	</tbody>
</table>

<p align="right" style="margin: 0cm 0cm 0pt; text-align: right;"><span lang="EN-US" style="letter-spacing: 0.2pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">&nbsp;</font></font></span></span></p>

<p>&nbsp;</p>

<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span lang="EN-US" style="letter-spacing: 0.2pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">&nbsp;</font></font></span></span></p>

<p align="right" style="margin: 0cm 0cm 0pt; text-align: right;"><span style="letter-spacing: 0.2pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">续上表</font></font></span></span></p>

<table style="border: 0px rgb(0, 0, 0); border-image: none; border-collapse: collapse;" width="631">
	<thead>
		<tr>
			<td colspan="8" style="padding: 0cm 5.4pt; border: 1pt solid windowtext; border-image: none; width: 333.95pt;" valign="top" width="445">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><b><span style="font-family: 宋体;"><font color="#000000"><font size="3">评分因素与权重分值</font></font></span></b><a href="http://218.76.24.174:8090/G2/gbp/jgw-notice!add.do#_ftn6" name="_ftnref6" title=""><sup><sup><span lang="EN-US" style="font-size: 10.5pt;"><span style="font-family: &quot;Calibri&quot;,&quot;sans-serif&quot;;"><font color="#0066cc">[6]</font></span></span></sup></sup></a></span></p>
			</td>
			<td rowspan="2" style="border-width: 1pt 1pt 1pt 0px; border-style: solid solid solid none; border-color: windowtext windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 139.25pt;" width="186">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><b><span style="font-family: 宋体;"><font color="#000000"><font size="3">评分标准</font></font></span></b><a href="http://218.76.24.174:8090/G2/gbp/jgw-notice!add.do#_ftn7" name="_ftnref7" title=""><sup><sup><span lang="EN-US" style="font-size: 10.5pt;"><span style="font-family: &quot;Calibri&quot;,&quot;sans-serif&quot;;"><font color="#0066cc">[7]</font></span></span></sup></sup></a></span></p>
			</td>
		</tr>
		<tr>
			<td style="border-width: 0px 1pt 1pt; border-style: none solid solid; border-color: rgb(0, 0, 0) windowtext windowtext; padding: 0cm 5.4pt; width: 70.2pt;" width="94">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><b><span style="font-family: 宋体;"><font color="#000000"><font size="3">条款号</font></font></span></b></span></p>
			</td>
			<td colspan="2" style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 64.95pt;" width="87">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><b><span style="font-family: 宋体;"><font color="#000000"><font size="3">评分因素</font></font></span></b></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 49.85pt;" width="66">
			<p align="center" style="margin: 0cm -5.25pt 0pt; text-align: center;"><span style="line-height: 19pt;"><b><span style="font-family: 宋体;"><font color="#000000"><font size="3">评分因素权重分值</font></font></span></b></span></p>
			</td>
			<td colspan="3" style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 107.15pt;" width="143">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><b><span style="font-family: 宋体;"><font color="#000000"><font size="3">各评分因素细分项</font></font></span></b></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 41.8pt;" width="56">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><b><span style="font-family: 宋体;"><font color="#000000"><font size="3">分值</font></font></span></b></span></p>
			</td>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td rowspan="6" style="border-width: 0px 1pt 1pt; border-style: none solid solid; border-color: rgb(0, 0, 0) windowtext windowtext; padding: 0cm 5.4pt; width: 70.2pt;" width="94">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">2.2.4</font></font></span><span style="font-family: 宋体;"><font size="3"><font color="#000000">（1</font></font><font color="#000000"><font size="3">）</font></font></span></span></p>
			</td>
			<td colspan="2" rowspan="6" style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 64.95pt;" width="87">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">施工组织设计</font></font></span></span></p>
			</td>
			<td rowspan="6" style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 49.85pt;" width="66">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><font size="3"><font color="#000000"><u><span lang="EN-US" style="font-family: 宋体;">16</span></u><span style="font-family: 宋体;">分</span></font></font></span></p>
			</td>
			<td colspan="3" style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 107.15pt;" width="143">
			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">日常养护的巡查制度安排，包括内容、频率、方法等计划</font></font></span></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 41.8pt;" width="56">
			<p align="center" style="margin: 0cm 5.25pt 0pt 0cm; text-align: center;"><span style="line-height: 19pt;"><font size="3"><font color="#000000"><u><span lang="EN-US" style="font-family: 宋体;">3</span></u><span style="font-family: 宋体;">分</span></font></font></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 139.25pt;" width="186">
			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">好的计3</font></font><font color="#000000"><font size="3">分；</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">较好的计2.7</font></font><font color="#000000"><font size="3">分；</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">一般的计2.4</font></font><font color="#000000"><font size="3">分。</font></font></span></span></p>
			</td>
		</tr>
		<tr>
			<td colspan="3" style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 107.15pt;" width="143">
			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">对检查、专项调查和技术检测安排，及其结果的养护对策方案</font></font></span></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 41.8pt;" width="56">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><font size="3"><font color="#000000"><u><span lang="EN-US" style="font-family: 宋体;">3</span></u><span style="font-family: 宋体;">分</span></font></font></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 139.25pt;" width="186">
			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">好的计3</font></font><font color="#000000"><font size="3">分；</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">较好的计2.7</font></font><font color="#000000"><font size="3">分；</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">一般的计2.4</font></font><font color="#000000"><font size="3">分。</font></font></span></span></p>
			</td>
		</tr>
		<tr>
			<td colspan="3" style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 107.15pt;" width="143">
			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">施工组织、现场布置、劳动力、机械设备、材料供应、资金流量、保畅方案及进度计划等</font></font></span></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 41.8pt;" width="56">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><font size="3"><font color="#000000"><u><span lang="EN-US" style="font-family: 宋体;">3</span></u><span style="font-family: 宋体;">分</span></font></font></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 139.25pt;" width="186">
			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">好的计3</font></font><font color="#000000"><font size="3">分；</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">较好的计2.7</font></font><font color="#000000"><font size="3">分；</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">一般的计2.4</font></font><font color="#000000"><font size="3">分。</font></font></span></span></p>
			</td>
		</tr>
		<tr>
			<td colspan="3" style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 107.15pt;" width="143">
			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">质量、安全、环境保护措施保证体系</font></font></span></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 41.8pt;" width="56">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><font size="3"><font color="#000000"><u><span lang="EN-US" style="font-family: 宋体;">3</span></u><span style="font-family: 宋体;">分</span></font></font></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 139.25pt;" width="186">
			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">好的计3</font></font><font color="#000000"><font size="3">分；</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">较好的计2.7</font></font><font color="#000000"><font size="3">分；</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">一般的计2.4</font></font><font color="#000000"><font size="3">分。</font></font></span></span></p>
			</td>
		</tr>
		<tr>
			<td colspan="3" style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 107.15pt;" width="143">
			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">专项施工方案</font></font></span></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 41.8pt;" width="56">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><font size="3"><font color="#000000"><u><span lang="EN-US" style="font-family: 宋体;">2</span></u><span style="font-family: 宋体;">分</span></font></font></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 139.25pt;" width="186">
			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">好的计2</font></font><font color="#000000"><font size="3">分；</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">较好的计1.8</font></font><font color="#000000"><font size="3">分；</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">一般的计1.6</font></font><font color="#000000"><font size="3">分。</font></font></span></span></p>
			</td>
		</tr>
		<tr>
			<td colspan="3" style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 107.15pt;" width="143">
			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">应急方案及其它应说明的事项</font></font></span></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 41.8pt;" width="56">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><font size="3"><font color="#000000"><u><span lang="EN-US" style="font-family: 宋体;">2</span></u><span style="font-family: 宋体;">分</span></font></font></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 139.25pt;" width="186">
			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">好的计2</font></font><font color="#000000"><font size="3">分；</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">较好的计1.8</font></font><font color="#000000"><font size="3">分；</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">一般的计1.6</font></font><font color="#000000"><font size="3">分。</font></font></span></span></p>
			</td>
		</tr>
		<tr>
			<td rowspan="2" style="border-width: 0px 1pt 1pt; border-style: none solid solid; border-color: rgb(0, 0, 0) windowtext windowtext; padding: 0cm 5.4pt; width: 70.2pt;" width="94">
			<p align="center" style="margin: 0cm -4.4pt 0pt -4.6pt; text-align: center;"><span style="line-height: 18pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">2.2.4</font></font></span><span style="font-family: 宋体;"><font size="3"><font color="#000000">（2</font></font><font color="#000000"><font size="3">）</font></font></span></span></p>
			</td>
			<td colspan="2" rowspan="2" style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 64.95pt;" width="87">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 18pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">主要人员</font></font></span></span></p>
			</td>
			<td rowspan="2" style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 49.85pt;" width="66">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><font size="3"><font color="#000000"><u><span lang="EN-US" style="font-family: 宋体;">10</span></u><span style="font-family: 宋体;">分</span></font></font></span></p>
			</td>
			<td colspan="3" style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 107.15pt;" width="143">
			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">项目经理任职资格与业绩</font></font></span></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 41.8pt;" width="56">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><font size="3"><font color="#000000"><u><span lang="EN-US" style="font-family: 宋体;">5</span></u><span style="font-family: 宋体;">分</span></font></font></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 139.25pt;" width="186">
			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">1</font></font></span><span style="font-family: 宋体;"><font size="3"><font color="#000000">、满足资格审查条件（项目经理最低要求）得基本分4</font></font><font color="#000000"><font size="3">分；</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">2</font></font></span><span style="font-family: 宋体;"><font size="3"><font color="#000000">、最近5</font></font><font color="#000000"><font size="3">年（递交投标文件截止之日前一日回溯</font></font><font color="#000000"><font size="3">5</font></font><font color="#000000"><font size="3">年），每增加担任过一个高速公路养护工程施工项目的项目经理或总工的业绩，加</font></font><font color="#000000"><font size="3">1</font></font><font color="#000000"><font size="3">分，最多加</font></font><font color="#000000"><font size="3">1</font></font><font color="#000000"><font size="3">分。</font></font></span></span></p>
			</td>
		</tr>
		<tr style="height: 39pt;">
			<td colspan="3" style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 107.15pt; height: 39pt;" width="143">
			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">项目总工任职资格与业绩</font></font></span></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 41.8pt; height: 39pt;" width="56">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><font size="3"><font color="#000000"><u><span lang="EN-US" style="font-family: 宋体;">5</span></u><span style="font-family: 宋体;">分</span></font></font></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 139.25pt; height: 39pt;" width="186">
			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">1</font></font></span><span style="font-family: 宋体;"><font size="3"><font color="#000000">、满足资格审查条件（项目总工最低要求）得基本分4</font></font><font color="#000000"><font size="3">分；</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">2</font></font></span><span style="font-family: 宋体;"><font size="3"><font color="#000000">、最近5</font></font><font color="#000000"><font size="3">年（递交投标文件截止之日前一日回溯</font></font><font color="#000000"><font size="3">5</font></font><font color="#000000"><font size="3">年），每增加担任过一个高速公路养护工程施工项目的项目经理或总工的业绩，加</font></font><font color="#000000"><font size="3">1</font></font><font color="#000000"><font size="3">分，最多加</font></font><font color="#000000"><font size="3">1</font></font><font color="#000000"><font size="3">分。</font></font></span></span></p>
			</td>
		</tr>
		<tr>
			<td style="border-width: 0px 1pt 1pt; border-style: none solid solid; border-color: rgb(0, 0, 0) windowtext windowtext; padding: 0cm 5.4pt; width: 70.2pt;" width="94">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">2.2.4</font></font></span><span style="font-family: 宋体;"><font size="3"><font color="#000000">（3</font></font><font color="#000000"><font size="3">）</font></font></span></span></p>
			</td>
			<td colspan="2" style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 64.95pt;" width="87">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">评标价</font></font></span></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 49.85pt;" width="66">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><font size="3"><font color="#000000"><u><span lang="EN-US" style="font-family: 宋体;">55</span></u><span style="font-family: 宋体;">分</span></font></font></span></p>
			</td>
			<td colspan="5" style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 288.2pt;" valign="top" width="384">
			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">评标价得分计算公式示例：</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（1</font></font><font color="#000000"><font size="3">）如果投标人的评标价＞评标基准价，则评标价得分</font></font><font color="#000000"><font size="3">=55-</font></font><font color="#000000"><font size="3">偏差率&times;</font></font><font color="#000000"><font size="3">100</font></font><font color="#000000"><font size="3">&times;</font></font><font color="#000000"><font size="3">E1</font></font><font color="#000000"><font size="3">；</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（2</font></font><font color="#000000"><font size="3">）如果投标人的评标价&le;评标基准价，则评标价得分</font></font><font color="#000000"><font size="3">=55+</font></font><font color="#000000"><font size="3">偏差率&times;</font></font><font color="#000000"><font size="3">100</font></font><font color="#000000"><font size="3">&times;</font></font><font color="#000000"><font size="3">E2</font></font><font color="#000000"><font size="3">（低于理论成本价的投标报价计算评标价得分）；</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">其中：</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">E1</font></font></span><span style="font-family: 宋体;"><font size="3"><font color="#000000">是评标价每高于评标基准价一个百分点的扣分值，E1=1.0</font></font><font color="#000000"><font size="3">；</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">E2</font></font></span><span style="font-family: 宋体;"><font size="3"><font color="#000000">是评标价每低于评标基准价一个百分点的扣分值，E2=0.5</font></font><font color="#000000"><font size="3">。</font></font></span></span></p>
			</td>
		</tr>
		<tr style="height: 56.25pt;">
			<td rowspan="3" style="border-width: 0px 1pt 1pt; border-style: none solid solid; border-color: rgb(0, 0, 0) windowtext windowtext; padding: 0cm 5.4pt; width: 70.2pt; height: 56.25pt;" width="94">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">&nbsp;</font></font></span></span></p>

			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">2.2.4</font></font></span><span style="font-family: 宋体;"><font size="3"><font color="#000000">（4</font></font><font color="#000000"><font size="3">）</font></font></span></span></p>
			</td>
			<td rowspan="3" style="border-width: 0px 0px 1pt; border-style: none none solid; border-color: rgb(0, 0, 0) rgb(0, 0, 0) windowtext; padding: 0cm 5.4pt; width: 12pt; height: 56.25pt;" width="16">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">&nbsp;</font></font></span></span></p>
			</td>
			<td rowspan="3" style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 52.95pt; height: 56.25pt;" width="71">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">其他因素</font></font></span></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 49.85pt; height: 56.25pt;" width="66">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><font size="3"><font color="#000000"><span lang="EN-US" style="font-family: 宋体;">5</span><span style="font-family: 宋体;">分</span></font></font></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 30.8pt; height: 56.25pt;" width="41">
			<p style="margin: 0cm 6.2pt 0pt 0cm;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">综合实力</font></font></span></p>
			</td>
			<td colspan="4" style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 257.4pt; height: 56.25pt;" width="343">
			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">1</font></font></span><span style="font-family: 宋体;"><font size="3"><font color="#000000">、养护应急工作获奖（3</font></font><font color="#000000"><font size="3">分）</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">2015</font></font></span><span style="font-family: 宋体;"><font size="3"><font color="#000000">年1</font></font><font color="#000000"><font size="3">月</font></font><font color="#000000"><font size="3">1</font></font><font color="#000000"><font size="3">日至今，积极响应应急抢险工作，在高速公路养护应急工作中作出突出贡献的加</font></font><font color="#000000"><font size="3">3</font></font><font color="#000000"><font size="3">分。</font></font><font color="#000000"><font size="3">(</font></font><font color="#000000"><font size="3">附省级及以上公路管理机构出具的相关证明材料，联合体形式参与投标的，应急抢险突出贡献证明材料必须是联合体牵头方提供</font></font><font color="#000000"><font size="3">)</font></font><font color="#000000"><font size="3">；</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">2</font></font></span><span style="font-family: 宋体;"><font size="3"><font color="#000000">、交通安全生产标准化建设达标一级认证（2</font></font><font color="#000000"><font size="3">分）</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">具有交通安全生产标准化建设达标一级认证证书(</font></font><font color="#000000"><font size="3">限交通运输工程建设类</font></font><font color="#000000"><font size="3">)</font></font><font color="#000000"><font size="3">且证书处于有效期内投标人加</font></font><font color="#000000"><font size="3">2</font></font><font color="#000000"><font size="3">分。</font></font></span></span></p>
			</td>
		</tr>
		<tr style="height: 68.9pt;">
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 49.85pt; height: 68.9pt;" width="66">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><font size="3"><font color="#000000"><span lang="EN-US" style="font-family: 宋体;">6</span><span style="font-family: 宋体;">分</span></font></font></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 30.8pt; height: 68.9pt;" width="41">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">财务</font></font></span></span></p>

			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">能力</font></font></span></span></p>
			</td>
			<td colspan="4" style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 257.4pt; height: 68.9pt;" width="343">
			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">1.</font></font></span><span style="font-family: 宋体;"><font size="3"><font color="#000000">满足资格审查条件（财务最低要求），得基本分4.8</font></font><font color="#000000"><font size="3">分；</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">2</font></font></span><span style="font-family: 宋体;"><font size="3"><font color="#000000">．在满足资格审查条件（财务最低要求）的基础上，投标人连续三年（2016</font></font><font color="#000000"><font size="3">～</font></font><font color="#000000"><font size="3">2018</font></font><font color="#000000"><font size="3">年）资产负债率低于</font></font><font color="#000000"><font size="3">75%</font></font><font color="#000000"><font size="3">大于等于</font></font><font color="#000000"><font size="3">45%</font></font><font color="#000000"><font size="3">的加</font></font><font color="#000000"><font size="3">1.2</font></font><font color="#000000"><font size="3">分，最多加</font></font><font color="#000000"><font size="3">1.2</font></font><font color="#000000"><font size="3">分。</font></font></span></span></p>
			</td>
		</tr>
		<tr style="height: 23.85pt;">
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 49.85pt; height: 23.85pt;" width="66">
			<p align="center" style="margin: 0cm -5.25pt 0pt; text-align: center;"><span style="line-height: 19pt;"><font size="3"><font color="#000000"><span lang="EN-US" style="font-family: 宋体;">8</span><span style="font-family: 宋体;">分</span></font></font></span></p>
			</td>
			<td style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 30.8pt; height: 23.85pt;" width="41">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">业绩</font></font></span></span></p>
			</td>
			<td colspan="4" style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 257.4pt; height: 23.85pt;" width="343">
			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">1</font></font></span><span style="font-family: 宋体;"><font size="3"><font color="#000000">、满足资格审查条件（业绩最低要求）得 6.4</font></font><font color="#000000"><font size="3">分；</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">2</font></font></span><span style="font-family: 宋体;"><font size="3"><font color="#000000">、最近5</font></font><font color="#000000"><font size="3">年（递交投标文件截止之日前一日回溯</font></font><font color="#000000"><font size="3">5</font></font><font color="#000000"><font size="3">年），投标人提供的高速公路养护中修工程或改扩建工程或高速公路路面专项维修病害处治工程中，含特大桥施工工程，加</font></font><font color="#000000"><font size="3">0.4</font></font><font color="#000000"><font size="3">分，最多加</font></font><font color="#000000"><font size="3">0.4</font></font><font color="#000000"><font size="3">分；</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">3</font></font></span><span style="font-family: 宋体;"><font size="3"><font color="#000000">、近5</font></font><font color="#000000"><font size="3">年（递交投标文件截止之日前一日回溯</font></font><font color="#000000"><font size="3">5</font></font><font color="#000000"><font size="3">年），投标人提供的高速公路养护中修工程或改扩建工程或高速公路路面专项维修病害处治工程中，含特长隧道施工工程，加</font></font><font color="#000000"><font size="3">0.4</font></font><font color="#000000"><font size="3">分，最多加</font></font><font color="#000000"><font size="3">0.4</font></font><font color="#000000"><font size="3">分。</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">4</font></font></span><span style="font-family: 宋体;"><font size="3"><font color="#000000">、最近5</font></font><font color="#000000"><font size="3">年（递交投标文件截止之日前一日回溯</font></font><font color="#000000"><font size="3">5</font></font><font color="#000000"><font size="3">年）投标人在任一单年度完成的高速公路小修</font></font><font color="#000000"><font size="3">(</font></font><font color="#000000"><font size="3">含连接线</font></font><font color="#000000"><font size="3">)</font></font><font color="#000000"><font size="3">保养项目累计养护里程达到</font></font><font color="#000000"><font size="3">150</font></font><font color="#000000"><font size="3">公里计</font></font><font color="#000000"><font size="3">0.4</font></font><font color="#000000"><font size="3">分，累计养护里程达到</font></font><font color="#000000"><font size="3">300</font></font><font color="#000000"><font size="3">公里及以上加</font></font><font color="#000000"><font size="3">0.8</font></font><font color="#000000"><font size="3">分，最多加</font></font><font color="#000000"><font size="3">0.8</font></font><font color="#000000"><font size="3">分。</font></font></span></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><b><span style="font-family: 宋体;"><font color="#000000"><font size="3">投标人提供的业绩为联合体形式中标业绩的，应提供项目业主签订的合同协议书、联合体协议以及项目业主出具的能体现联合体各方已完业绩的证明材料（须包含业主联系人、联系电话），业绩认定只计算投标人的业绩，联合体其他成员业绩不予计算。</font></font></span></b></span></p>
			</td>
		</tr>
		<tr style="height: 25.05pt;">
			<td style="border-width: 0px 1pt 1pt; border-style: none solid solid; border-color: rgb(0, 0, 0) windowtext windowtext; padding: 0cm 5.4pt; width: 70.2pt; height: 25.05pt;" width="94">
			<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 19pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">3.6.1</font></font></span></span></p>
			</td>
			<td colspan="5" style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 168.75pt; height: 25.05pt;" width="225">
			<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">投标文件相关信息的核查</font></font></span></p>
			</td>
			<td colspan="3" style="border-width: 0px 1pt 1pt 0px; border-style: none solid solid none; border-color: rgb(0, 0, 0) windowtext windowtext rgb(0, 0, 0); padding: 0cm 5.4pt; width: 234.25pt; height: 25.05pt;" width="312">
			<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">不适用。</font></font></span></p>
			</td>
		</tr>
		<tr>
			<td colspan="9" style="border-width: 0px 1pt 1pt; border-style: none solid solid; border-color: rgb(0, 0, 0) windowtext windowtext; padding: 0cm 5.4pt; width: 473.2pt;" width="631">
			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 18pt;"><font size="3"><font color="#000000"><span lang="EN-US" style="font-family: 宋体;">3.4.3</span><span style="font-family: 宋体;">项修改为：</span></font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 19pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">评标委员会发现投标人的报价明显低于其他投标报价，且其评标价低于理论成本价时，应要求该投标人作出书面说明并提供相应的证明材料。投标人不能合理说明或不能提供相应证明材料的，由评标委员会认定该投标人以低于成本报价竞标，并否决其投标。理论成本价的确定按第二章&ldquo;投标人须知&rdquo;第5.2.4.1</font></font><font color="#000000"><font size="3">目规定计算。</font></font></span></span></p>
			</td>
		</tr>
		<tr>
			<td colspan="9" style="border-width: 0px 1pt 1pt; border-style: none solid solid; border-color: rgb(0, 0, 0) windowtext windowtext; padding: 0cm 5.4pt; width: 473.2pt;" width="631">
			<p style="margin: 0cm 0cm 0pt;"><b><span style="font-family: 宋体;"><font size="3"><font color="#000000">补充3.9.3</font></font><font color="#000000"><font size="3">：评标报告应当载明下列内容：</font></font></span></b></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">（一）招标项目基本情况；</font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">（二）评标委员会成员名单；</font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">（三）监督人员名单；</font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">（四）开标记录；</font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">（五）符合要求的投标人名单；</font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">（六）否决的投标人名单以及否决理由；</font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">（七）串通投标情形的评审情况说明；</font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">（八）评分情况；</font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">（九）经评审的投标人排序；</font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">（十）中标候选人名单；</font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">（十一）澄清、说明事项纪要；</font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">（十二）需要说明的其他事项；</font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">（十三）评标附表。</font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">对评标监督人员或者招标人代表干预正常评标活动，以及对招标投标活动的其他不正当言行，评标委员会应当在评标报告第（十二）项内容中如实记录。</font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">除第（一）、（三）、（四）项内容外，评标委员会所有成员应当在评标报告上逐页签字。对评标结果有不同意见的评标委员会成员应当以书面形式说明其不同意见和理由，评标报告应当注明该不同意见。评标委员会成员拒绝在评标报告上签字又不书面说明其不同意见和理由的，视为同意评标结果。</font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><b><span style="font-family: 宋体;"><font size="3"><font color="#000000">新增：3.9.4</font></font></span></b><span style="font-family: 宋体;"><font size="3"><font color="#000000">评标委员会对投标文件进行评审后，因有效投标不足3</font></font><font color="#000000"><font size="3">个使得投标明显缺乏竞争的</font></font><font color="#000000"><font size="3">,</font></font><font color="#000000"><font size="3">可以否决全部投标。未否决全部投标的，评标委员会应当在评标报告中阐明理由并推荐中标候选人。</font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">投标文件按照招标文件规定采用双信封形式密封的，通过第一信封商务文件和技术文件评审的投标人在3</font></font><font color="#000000"><font size="3">个以上的，招标人应当按照招标文件规定的程序进行第二信封报价文件开标；在对报价文件进行评审后，有效投标不足</font></font><font color="#000000"><font size="3">3</font></font><font color="#000000"><font size="3">个的，评标委员会应当按照本条第一款规定执行。</font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">通过第一信封商务文件和技术文件评审的投标人少于3</font></font><font color="#000000"><font size="3">个的，评标委员会可以否决全部投标；未否决全部投标的，评标委员会应当在评标报告中阐明理由，招标人应当按照招标文件规定的程序进行第二信封报价文件开标，但评标委员会在进行报价文件评审时仍有权否决全部投标；评标委员会未在报价文件评审时否决全部投标的，应当在评标报告中阐明理由并推荐中标候选人。</font></font></span></p>
			</td>
		</tr>
		<tr>
			<td colspan="9" style="border-width: 0px 1pt 1pt; border-style: none solid solid; border-color: rgb(0, 0, 0) windowtext windowtext; padding: 0cm 5.4pt; width: 473.2pt;" width="631">
			<p style="margin: 0cm 0cm 0pt;"><b><span style="font-family: 宋体;"><font size="3"><font color="#000000">增加4.0</font></font><font color="#000000"><font size="3">款：</font></font></span></b></p>

			<p style="margin: 0cm 0cm 0pt;"><font size="3"><font color="#000000"><b><span lang="EN-US" style="font-family: 宋体;">4.0 </span></b><b><span style="font-family: 宋体;">重新招标</span></b></font></font></p>

			<p style="margin: 0cm 0cm 0pt;"><font size="3"><font color="#000000"><span lang="EN-US" style="font-family: 宋体;">4.1.1</span><span style="font-family: 宋体;">依法必须进行招标的公路工程建设项目，有下列情形之一的，招标人在分析招标失败的原因并采取相应措施后，应当重新招标：</span></font></font></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（1</font></font><font color="#000000"><font size="3">）递交投标文件的投标人少于</font></font><font color="#000000"><font size="3">3</font></font><font color="#000000"><font size="3">个的；</font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（2</font></font><font color="#000000"><font size="3">）所有投标均被否决的；</font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">（3</font></font><font color="#000000"><font size="3">）中标候选人均未与招标人订立书面合同的。</font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">重新招标的，招标文件和招标投标情况的书面报告应当按规定重新报交通运输主管部门备案。</font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="font-family: 宋体;"><font size="3"><font color="#000000">重新招标后投标人仍少于3</font></font><font color="#000000"><font size="3">个的，属于按照国家有关规定需要履行项目审批、核准手续的依法必须进行招标的公路工程建设项目，报经项目审批、核准部门批准后可以不再进行招标；其他项目可由招标人自行决定不再进行招标。</font></font></span></p>

			<p style="margin: 0cm 0cm 0pt;"><span style="font-family: 宋体;"><font color="#000000"><font size="3">依照规定不再进行招标的，招标人可以邀请已提交投标文件的投标人进行谈判，确定项目承担单位，并将谈判报告报对该项目具有招标监督职责的交通运输主管部门备案。</font></font></span></p>
			</td>
		</tr>
		<tr height="0">
			<td style="border: 0px rgb(0, 0, 0); border-image: none;" width="94">&nbsp;</td>
			<td style="border: 0px rgb(0, 0, 0); border-image: none;" width="16">&nbsp;</td>
			<td style="border: 0px rgb(0, 0, 0); border-image: none;" width="71">&nbsp;</td>
			<td style="border: 0px rgb(0, 0, 0); border-image: none;" width="66">&nbsp;</td>
			<td style="border: 0px rgb(0, 0, 0); border-image: none;" width="41">&nbsp;</td>
			<td style="border: 0px rgb(0, 0, 0); border-image: none;" width="31">&nbsp;</td>
			<td style="border: 0px rgb(0, 0, 0); border-image: none;" width="71">&nbsp;</td>
			<td style="border: 0px rgb(0, 0, 0); border-image: none;" width="56">&nbsp;</td>
			<td style="border: 0px rgb(0, 0, 0); border-image: none;" width="186">&nbsp;</td>
		</tr>
	</tbody>
</table>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">&nbsp;</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">&nbsp;</font></span></span></p>

<p style="margin: 18pt 0cm 12pt;"><span style="page-break-after: avoid;"><a name="_Toc234832944"></a><a name="_Toc501257158"></a><a name="_Toc517789269"><span style="font-size: 14pt;"><span style="font-family: 宋体;"><font color="#000000">1. 评标方法</font></span></span></a></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">本次评标采用综合评分法。评标委员会对满足招标文件实质性要求的投标文件，按照本章第2.2</font><font color="#000000">款规定的评分标准进行打分，并按得分由高到低顺序推荐中标候选人，或根据招标人授权直接确定中标人，但投标报价低于其成本的除外。综合评分相等时，评标委员会应按照评标办法前附表规定的优先次序推荐中标候选人或确定中标人。</font></span></span></span></p>

<p style="margin: 18pt 0cm 12pt;"><span style="page-break-after: avoid;"><a name="_Toc517789270"></a><a name="_Toc501257159"><span style="font-size: 14pt;"><span style="font-family: 宋体;"><font color="#000000">2. 评审标准</font></span></span></a></span></p>

<p style="margin: 12pt 0cm;"><span style="line-height: 12pt;"><span style="page-break-after: avoid;"><a name="_Toc517789271"></a><a name="_Toc501257160"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">2.1 初步评审标准</font></span></span></a></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">2.1.1 </span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">形式评审标准：见评标办法前附表。</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">2.1.2 </span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">资格评审标准：见评标办法前附表（适用于未进行资格预审的）。</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">2.1.2 </font></span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">资格评审标准：见资格预审文件第三章&ldquo;</font><font color="#000000">资格审查办法</font><font color="#000000">&rdquo;</font><font color="#000000">详细审查标准（适用于已进行资格预审的）。</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">2.1.3 </span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">响应性评审标准：见评标办法前附表。</span></span></font></span></p>

<p style="margin: 12pt 0cm;"><span style="line-height: 12pt;"><span style="page-break-after: avoid;"><a name="_Toc501257161"></a><a name="_Toc517789272"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">2.2 分值构成与评分标准</font></span></span></a></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="tab-stops: 18.0pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">2.2.1 </span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">分值构成</span></span></font></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">（1</font><font color="#000000">）施工组织设计：见评标办法前附表；</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">（2</font><font color="#000000">）主要人员：见评标办法前附表；</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">（3</font><font color="#000000">）评标价：见评标办法前附表；</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">（4</font><font color="#000000">）其他评分因素：见评标办法前附表。</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">2.2.2 </span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">评标基准价计算</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">评标基准价计算方法：见评标办法前附表。</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">2.2.3 </span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">评标价的偏差率计算</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">评标价的偏差率计算公式：见评标办法前附表。</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">2.2.4 </span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">评分标准</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">（1</font><font color="#000000">）施工组织设计评分标准：见评标办法前附表；</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">（2</font><font color="#000000">）主要人员评分标准：见评标办法前附表；</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">（3</font><font color="#000000">）评标价评分标准：见评标办法前附表；</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">（4</font><font color="#000000">）其他因素评分标准：见评标办法前附表。</font></span></span></span></p>

<p style="margin: 18pt 0cm 12pt;"><span style="page-break-after: avoid;"><a name="_Toc501257162"></a><a name="_Toc517789273"><span style="font-size: 14pt;"><span style="font-family: 宋体;"><font color="#000000">3. 评标程序</font></span></span></a></span></p>

<p style="margin: 12pt 0cm;"><span style="line-height: 12pt;"><span style="page-break-after: avoid;"><a name="_Toc501257163"></a><a name="_Toc517789274"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">3.1 第一个信封初步评审</font></span></span></a></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">3.1.1 </font></span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">评标委员会可以要求投标人提交第二章&ldquo;</font><font color="#000000">投标人须知</font><font color="#000000">&rdquo;</font><font color="#000000">第</font><font color="#000000">3.5.1</font><font color="#000000">项至第</font><font color="#000000">3.5.6</font><font color="#000000">项规定的有关证明和证件的原件，以便核验。评标委员会依据本章第</font><font color="#000000">2.1</font><font color="#000000">款规定的标准对投标文件第一个信封（商务及技术文件）进行初步评审。有一项不符合评审标准的，评标委员会应否决其投标。（适用于未进行资格预审的）</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">3.1.1 </font></span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">评标委员会依据本章第2.1.1</font><font color="#000000">项、第</font><font color="#000000">2.1.3</font><font color="#000000">项规定的评审标准对投标文件第一个信封（商务及技术文件）进行初步评审。有一项不符合评审标准的，评标委员会应否决其投标。当投标人资格预审申请文件的内容发生重大变化时，评标委员会依据本章第</font><font color="#000000">2.1.2</font><font color="#000000">项规定的标准对其更新资料进行评审。（适用于已进行资格预审的）</font></span></span></span></p>

<p style="margin: 12pt 0cm;"><span style="line-height: 12pt;"><span style="page-break-after: avoid;"><a name="_Toc501257164"></a><a name="_Toc517789275"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">3.2 第一个信封详细评审</font></span></span></a></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">3.2.1 </font></span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">评标委员会按本章第2.2</font><font color="#000000">款规定的量化因素和分值进行打分，并计算出各投标人的商务和技术得分。</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">（1</font><font color="#000000">）按本章第</font><font color="#000000">2.2.4</font><font color="#000000">项（</font><font color="#000000">1</font><font color="#000000">）目规定的评审因素和分值对施工组织设计部分计算出得分</font><font color="#000000">A</font><font color="#000000">；</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">（2</font><font color="#000000">）按本章第</font><font color="#000000">2.2.4</font><font color="#000000">项（</font><font color="#000000">2</font><font color="#000000">）目规定的评审因素和分值对主要人员部分计算出得分</font><font color="#000000">B</font><font color="#000000">；</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">（3</font><font color="#000000">）按本章第</font><font color="#000000">2.2.4</font><font color="#000000">项（</font><font color="#000000">4</font><font color="#000000">）目规定的评审因素和分值对其他部分计算出得分</font><font color="#000000">D</font><font color="#000000">。</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">3.2.2 </font></span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">投标人的商务和技术得分分值计算保留小数点后两位，小数点后第三位&ldquo;</font><font color="#000000">四舍五入</font><font color="#000000">&rdquo;</font><font color="#000000">。</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">3.2.3 </font></span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">投标人的商务和技术得分=A+B+D</font><font color="#000000">。 </font></span></span></span></p>

<p style="margin: 12pt 0cm;"><span style="line-height: 12pt;"><span style="page-break-after: avoid;"><a name="_Toc517789276"></a><a name="_Toc501257165"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">3.3 第二个信封开标</font></span></span></a></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">第一个信封（商务及技术文件）评审结束后，招标人将按照第二章&ldquo;</font><font color="#000000">投标人须知</font><font color="#000000">&rdquo;</font><font color="#000000">第</font><font color="#000000">5.1</font><font color="#000000">款规定的时间和地点对通过投标文件第一个信封（商务及技术文件）评审的投标文件第二个信封（报价文件）进行开标。</font></span></span></span></p>

<p style="margin: 12pt 0cm;"><span style="line-height: 12pt;"><span style="page-break-after: avoid;"><a name="_Toc501257166"></a><a name="_Toc517789277"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">3.4第二个信封初步评审</font></span></span></a></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">3.4.1 </font></span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">评标委员会依据本章第2.1.1</font><font color="#000000">项、第</font><font color="#000000">2.1.3</font><font color="#000000">项规定的评审标准对投标文件第二个信封（报价文件）进行初步评审。有一项不符合评审标准的，评标委员会应否决其投标。</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">3.4.2 </span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">投标报价有算术错误的，评标委员会按以下原则对投标报价进行修正，修正的价格经投标人书面确认后具有约束力。投标人不接受修正价格的，评标委员会应否决其投标。</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">（1</font><font color="#000000">）投标文件中的大写金额与小写金额不一致的，以大写金额为准；</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">（2</font><font color="#000000">）总价金额与依据单价计算出的结果不一致的，以单价金额为准修正总价，但单价金额小数点有明显错误的除外</font><font color="#000000">;</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">（3</font><font color="#000000">）当单价与数量相乘不等于合价时，以单价计算为准，如果单价有明显的小数点位置差错，应以标出的合价为准，同时对单价予以修正；</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">（4</font><font color="#000000">）当各子目的合价累计不等于总价时，应以各子目合价累计数为准，修正总价。</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">3.4.3 </span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">工程量清单中的投标报价有其他错误的，评标委员会按以下原则对投标报价进行修正，修正的价格经投标人书面确认后具有约束力。投标人不接受修正价格的，评标委员会应否决其投标。</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">（1</font><font color="#000000">）在招标人给定的工程量清单中漏报了某个工程子目的单价、合价或总额价，或所报单价、合价或总额价减少了报价范围，则漏报的工程子目单价、合价和总额价或单价、合价和总额价中减少的报价内容视为已含入其他工程子目的单价、合价和总额价之中。</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">（2</font><font color="#000000">）在招标人给定的工程量清单中多报了某个工程子目的单价、合价或总额价，或所报单价、合价或总额价增加了报价范围，则从投标报价中扣除多报的工程子目报价或工程子目报价中增加了报价范围的部分报价。</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">（3</font><font color="#000000">）当单价与数量的乘积与合价（金额）虽然一致，但投标人修改了该子目的工程数量，则其合价按招标人给定的工程数量乘以投标人所报单价予以修正。</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">3.4.4 </span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">修正后的最终投标报价若超过最高投标限价（如有），评标委员会应否决其投标。</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">3.4.5 </span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">修正后的最终投标报价仅作为签订合同的一个依据，不参与评标价得分的计算。</span></span></font></span></p>

<p style="margin: 12pt 0cm;"><span style="line-height: 12pt;"><span style="page-break-after: avoid;"><a name="_Toc501257167"></a><a name="_Toc517789278"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">3.5 第二个信封详细评审</font></span></span></a></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">3.5.1 </font></span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">评标委员会按本章第2.2.4</font><font color="#000000">（</font><font color="#000000">3</font><font color="#000000">）目规定的评审因素和分值对评标价计算出得分</font><font color="#000000">C</font><font color="#000000">。评标价得分分值计算保留小数点后两位，小数点后第三位</font><font color="#000000">&ldquo;</font><font color="#000000">四舍五入</font><font color="#000000">&rdquo;</font><font color="#000000">。</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">3.5.2 </font></span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">投标人综合得分=</font><font color="#000000">投标人的商务和技术得分</font><font color="#000000">+C</font><font color="#000000">。</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">3.5.3 </span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">评标委员会发现投标人的报价明显低于其他投标报价，使得其投标报价可能低于其个别成本的，应要求该投标人作出书面说明并提供相应的证明材料。投标人不能合理说明或不能提供相应证明材料的，评标委员会应认定该投标人以低于成本报价竞标，并否决其投标。</span></span></font></span></p>

<p style="margin: 12pt 0cm;"><span style="line-height: 12pt;"><span style="page-break-after: avoid;"><a name="_Toc517789279"></a><a name="_Toc501257168"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">3.6 投标文件相关信息的核查</font></span></span></a></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">3.6.1 </font></span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">在评标过程中，评标委员会应查询交通运输主管部门&ldquo;</font><font color="#000000">公路建设市场信用信息管理系统</font><font color="#000000">&rdquo;</font><font color="#000000">，对投标人的资质、业绩、主要人员资历和目前在岗情况、信用等级等信息进行核实。若投标文件载明的信息与交通运输主管部门</font><font color="#000000">&ldquo;</font><font color="#000000">公路建设市场信用信息管理系统</font><font color="#000000">&rdquo;</font><font color="#000000">发布的信息不符，使得投标人的资格条件不符合招标文件规定的，评标委员会应否决其投标。</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">3.6.2 </span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">评标委员会应对在评标过程中发现的投标人与投标人之间、投标人与招标人之间存在的串通投标的情形进行评审和认定。投标人存在串通投标、弄虚作假、行贿等违法行为的，评标委员会应否决其投标。</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">（1</font><font color="#000000">）有下列情形之一的，属于投标人相互串通投标：</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">a.</span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">投标人之间协商投标报价等投标文件的实质性内容；</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">b.</span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">投标人之间约定中标人；</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">c.</span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">投标人之间约定部分投标人放弃投标或中标；</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">d.</span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">属于同一集团、协会、商会等组织成员的投标人按照该组织要求协同投标；</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">e.</span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">投标人之间为谋取中标或排斥特定投标人而采取的其他联合行动。</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">（2</font><font color="#000000">）有下列情形之一的，视为投标人相互串通投标：</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">a.</span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">不同投标人的投标文件由同一单位或个人编制；</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">b.</span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">不同投标人委托同一单位或个人办理投标事宜；</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">c.</span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">不同投标人的投标文件载明的项目管理成员为同一人；</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">d.</span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">不同投标人的投标文件异常一致或投标报价呈规律性差异；</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">e.</span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">不同投标人的投标文件相互混装；</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">f.</span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">不同投标人的投标保证金从同一单位或个人的账户转出。</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">（3</font><font color="#000000">）有下列情形之一的，属于招标人与投标人串通投标：</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">a.</span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">招标人在开标前开启投标文件并将有关信息泄露给其他投标人;</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">b.</span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">招标人直接或间接向投标人泄露标底、评标委员会成员等信息；</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">c.</span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">招标人明示或暗示投标人压低或抬高投标报价；</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">d.</span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">招标人授意投标人撤换、修改投标文件；</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">e.</span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">招标人明示或暗示投标人为特定投标人中标提供方便；</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">f.</span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">招标人与投标人为谋求特定投标人中标而采取的其他串通行为。</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">（4</font><font color="#000000">）投标人有下列情形之一的，属于弄虚作假的行为：</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">a.</span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">使用通过受让或租借等方式获取的资格、资质证书投标；</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">b.</span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">使用伪造、变造的许可证件；</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">c.</span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">提供虚假的财务状况或业绩；</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">d.</span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">提供虚假的项目负责人或主要技术人员简历、劳动关系证明；</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">e.</span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">提供虚假的信用状况；</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">f.</span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">其他弄虚作假的行为。</span></span></font></span></p>

<p style="margin: 12pt 0cm;"><span style="line-height: 20pt;"><span style="page-break-after: avoid;"><a name="_Toc517789280"></a><a name="_Toc501257169"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">3.7 投标文件的澄清和说明</font></span></span></a></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">3.7.1 </span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">在评标过程中，评标委员会可以书面形式要求投标人对投标文件中含义不明确的内容、明显文字或计算错误进行书面澄清或说明。评标委员会不接受投标人主动提出的澄清、说明。投标人不按评标委员会要求澄清或说明的，评标委员会应否决其投标。</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">3.7.2 </span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">澄清和说明不得超出投标文件的范围或改变投标文件的实质性内容（算术性错误的修正除外）。投标人的书面澄清、说明属于投标文件的组成部分。</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">3.7.3 </span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">评标委员会不得暗示或诱导投标人作出澄清、说明，对投标人提交的澄清、说明有疑问的，可以要求投标人进一步澄清或说明，直至满足评标委员会的要求。</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">3.7.4 </span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">凡超出招标文件规定的或给发包人带来未曾要求的利益的变化、偏差或其他因素在评标时不予考虑。</span></span></font></span></p>

<p style="margin: 12pt 0cm;"><span style="line-height: 20pt;"><span style="page-break-after: avoid;"><a name="_Toc517789281"></a><a name="_Toc501257170"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">3.8 不得否决投标的情形</font></span></span></a></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">投标文件存在第二章&ldquo;</font><font color="#000000">投标人须知</font><font color="#000000">&rdquo;</font><font color="#000000">第</font><font color="#000000">1.12.3</font><font color="#000000">项所列情形的，均视为细微偏差，评标委员会不得否决投标人的投标，应按照第二章</font><font color="#000000">&ldquo;</font><font color="#000000">投标人须知</font><font color="#000000">&rdquo;</font><font color="#000000">第</font><font color="#000000">1.12.4</font><font color="#000000">项规定的原则处理。</font></span></span></span></p>

<p style="margin: 12pt 0cm;"><span style="line-height: 12pt;"><span style="page-break-after: avoid;"><a name="_Toc517789282"></a><a name="_Toc501257171"><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">3.9 评标结果</font></span></span></a></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">3.9.1 </font></span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;"><font color="#000000">除第二章&ldquo;</font><font color="#000000">投标人须知</font><font color="#000000">&rdquo;</font><font color="#000000">前附表授权直接确定中标人外，评标委员会按照得分由高到低的顺序推荐中标候选人，并标明排序。</font></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 20pt;"><font color="#000000"><span lang="EN-US" style="font-size: 12pt;"><span style="font-family: 宋体;">3.9.2 </span></span><span style="font-size: 12pt;"><span style="font-family: 宋体;">评标委员会完成评标后，应向招标人提交书面评标报告。</span></span></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 26.25pt;"><span lang="EN-US" style="font-family: &quot;微软雅黑&quot;,&quot;sans-serif&quot;;"><font color="#000000"><font size="3">&nbsp;</font></font></span></span></p>

<p align="left" style="margin: 0cm 0cm 0pt; text-align: left;"><span style="line-height: 22pt;"><span lang="EN-US" style="font-family: 宋体;"><font color="#000000"><font size="3">&nbsp;</font></font></span></span></p>

<p align="center" style="margin: 0cm 0cm 0pt; text-align: center;"><span style="line-height: 30pt;"><span style="-ms-layout-grid-mode: char;"><span lang="EN-US" style="font-size: 18pt;"><span style="font-family: 黑体;"><font color="#000000">&nbsp;</font></span></span></span></span></p>

<p style="margin: 0cm 0cm 0pt;"><font color="#000000"><font face="Calibri"><font size="3">&nbsp;</font></font></font></p>

<div>&nbsp;
<hr align="left" size="1" width="33%" />
<div id="ftn1">
<p style="margin: 0cm 0cm 0pt;"><a href="http://218.76.24.174:8090/G2/gbp/jgw-notice!add.do#_ftnref1" name="_ftn1" title=""><span lang="EN-US" style="font-size: 9pt;"><span style="font-family: &quot;Calibri&quot;,&quot;sans-serif&quot;;"><font color="#0066cc">[1]</font></span></span></a><span style="font-family: 宋体;"><font color="#000000"><font size="2">单位负责人，是指单位法定代表人或者法律、行政法规规定代表单位行使职权的主要负责人。</font></font></span></p>
</div>

<div id="ftn2">
<p style="margin: 0cm 0cm 0pt;"><a href="http://218.76.24.174:8090/G2/gbp/jgw-notice!add.do#_ftnref2" name="_ftn2" title=""><span lang="EN-US" style="font-size: 9pt;"><span style="font-family: &quot;Calibri&quot;,&quot;sans-serif&quot;;"><font color="#0066cc">[2]</font></span></span></a><font size="2"><font color="#000000"><span style="font-family: 宋体;">控股，是指出资额（持股）占股本总额</span><font face="Calibri">50%</font><span style="font-family: 宋体;">以上或虽不足</span><font face="Calibri">50%</font><span style="font-family: 宋体;">，但依出资额或所持股份所享有的表决权已足以对股东会、股东大会的决议产生重大影响的，或者国有企事业单位通过投资关系、协议或者其他安排，能够实际支配公司行为的。</span></font></font></p>
</div>

<div id="ftn3">
<p style="margin: 0cm 0cm 0pt;"><a href="http://218.76.24.174:8090/G2/gbp/jgw-notice!add.do#_ftnref3" name="_ftn3" title=""><span lang="EN-US" style="font-size: 9pt;"><span style="font-family: &quot;Calibri&quot;,&quot;sans-serif&quot;;"><font color="#0066cc">[3]</font></span></span></a><span style="font-family: 宋体;"><font color="#000000"><font size="2">管理，是指不具有出资持股关系的其他单位之间存在的管理与被管理关系。</font></font></span></p>

<p style="margin: 0cm 0cm 0pt;"><font color="#000000"><font face="Calibri"><font size="2">&nbsp;</font></font></font></p>
</div>

<div id="ftn4">
<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 15pt;"><font color="#000000"><font face="Calibri"><font size="2">&nbsp;</font></font></font></span></p>
</div>

<div id="ftn5">
<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 15pt;"><font color="#000000"><font face="Calibri"><font size="2">&nbsp;</font></font></font></span></p>
</div>

<div id="ftn6">
<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 15pt;"><font color="#000000"><font face="Calibri"><font size="2">&nbsp;</font></font></font></span></p>
</div>

<div id="ftn7">
<p style="margin: 0cm 0cm 0pt;"><span style="line-height: 15pt;"><font color="#000000"><font face="Calibri"><font size="2">&nbsp;</font></font></font></span></p>
</div>
</div>
`

func TestReCleanHtml(t *testing.T) {
	r, _ := regexp.Compile(`<.*?>|&nbsp;|\<[\S\s]+?\>`)
	clean, _ := regexp.Compile(`\n{3,}`)

	result := r.ReplaceAllString(Str, "")
	result = clean.ReplaceAllString(result, "\n")
	result = strings.TrimRight(result, "\n")
	result = strings.TrimLeft(result, "\n")
	fmt.Println(result)
}

func TestReGetOneString(t *testing.T) {
	fmt.Println(ReGetOneString(`queryContent_(\d+)-jygk`, "/queryContent_11-jygk.jspx"))
	//fmt.Println(ReGetOneString(`/(\d+)页`, "共25419条记录 2/aaaa页"))
}
