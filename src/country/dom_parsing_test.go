package country

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
	"testing"
)

// countryListFromDom

func TestCountryList(t *testing.T) {
	for _, testCase := range countryListHtmlVectors {
		htmlReader := strings.NewReader(testCase.html)
		htmlDoc, err := goquery.NewDocumentFromReader(htmlReader)
		if err != nil {
			t.Error("Country List htmlReader error", testCase.name, err)
		}
		list, err := countryListFromDom(htmlDoc)
		if len(list) != len(testCase.expectedList) {
			t.Error("Country List length mismatch", testCase.name, list, testCase.expectedList)
		}
		for i, _ := range list {
			if list[i] != testCase.expectedList[i] {
				t.Error("Country List value mismatch at index", i, testCase.name, list, testCase.expectedList)
				break
			}
		}
		if err != testCase.expectedError {
			t.Error("Country List expected error", testCase.name, err, testCase.expectedError)
		}
	}
}

// textForFieldKey

func TestTextForFieldKey(t *testing.T) {
	for _, testCase := range textForFieldKeyVectors {
		htmlReader := strings.NewReader(testCase.html)
		htmlDoc, err := goquery.NewDocumentFromReader(htmlReader)
		if err != nil {
			t.Error("Field Key htmlReader error", testCase.name, err)
		}
		s, err := textForFieldKey(htmlDoc, testCase.fieldkey)
		if s != testCase.expectedString {
			t.Error("Field Key value mismatch", testCase.name, s, testCase.expectedString)
			break
		}
		if err != testCase.expectedError {
			t.Error("Field Key expected error", testCase.name, err, testCase.expectedError)
		}
	}
}

// test vectors

type countryListTestCase struct {
	name          string
	html          string
	expectedList  []string
	expectedError error
}

type textForFieldKeyTestCase struct {
	name           string
	html           string
	fieldkey       string
	expectedString string
	expectedError  error
}

var countryListHtmlVectors = []countryListTestCase{
	countryListTestCase{
		name: "single option",
		html: `
			<select>
				<option value="../geos/xx.html">World</option>
			</select>
		`,
		expectedList: []string{
			"xx.html",
		},
		expectedError: nil,
	},
	countryListTestCase{
		name: "multiple option",
		html: `
			<select>
				<option value="../geos/xx.html">World</option>
				<option value="../geos/af.html">Afghanistan</option>
			</select>
		`,
		expectedList: []string{
			"xx.html",
			"af.html",
		},
		expectedError: nil,
	},
	countryListTestCase{
		name: "no selects",
		html: `
			<div>
				No selects here
			</div>
		`,
		expectedList:  []string{},
		expectedError: IncorrectNumberOfSelects,
	},
	countryListTestCase{
		name: "too many selects",
		html: `
			<select>
				<option value="../geos/xx.html">World</option>
			</select>
			<select><option value="error"></option></select>
		`,
		expectedList: []string{
			"xx.html",
		},
		expectedError: nil,
	},
	countryListTestCase{
		name:          "real data from aa.html",
		html:          aaHtml20140902,
		expectedList:  aaCountryList20140902,
		expectedError: nil,
	},
}

var textForFieldKeyVectors = []textForFieldKeyTestCase{
	textForFieldKeyTestCase{
		name:           "Single line string extraction old format",
		html:           aaHtml20140902,
		fieldkey:       "2011",
		expectedString: "12 30 N, 69 58 W",
		expectedError:  nil,
	},
	textForFieldKeyTestCase{
		name:           "Single line string extraction new format",
		html:           asHtml20170320,
		fieldkey:       "2011",
		expectedString: "27 00 S, 133 00 E",
		expectedError:  nil,
	},
	textForFieldKeyTestCase{
		name:           "Invalid fieldkey value",
		html:           asHtml20170320,
		fieldkey:       "invalid_field_key",
		expectedString: "",
		expectedError:  IncorrectNumberOfFieldKeyLinks,
	},
	textForFieldKeyTestCase{
		name:           "Multiline fieldkey value",
		html:           asHtml20170320,
		fieldkey:       "2021",
		expectedString: "cyclones along the coast; severe droughts; forest fires\nvolcanism: volcanic activity on Heard and McDonald Islands",
		expectedError:  nil,
	},
}

// From here on is long rambling data from real data sources

const aaHtml20140902 = `

<!doctype html>
<!--[if lt IE 7]> <html class="no-js lt-ie9 lt-ie8 lt-ie7" lang="en"> <![endif]-->
<!--[if IE 7]>    <html class="no-js lt-ie9 lt-ie8" lang="en"> <![endif]-->
<!--[if IE 8]>    <html class="no-js lt-ie9" lang="en"> <![endif]-->
<!--[if gt IE 8]><!--> 
 <!--<![endif]-->
<html class="no-js" lang="en"><!-- InstanceBegin template="/Templates/wfbext_template.dwt.cfm" codeOutsideHTMLIsLocked="false" -->
<head>


<script type="text/javascript" src="/static/js/analytics.js"></script>
<script type="text/javascript">archive_analytics.values.server_name="wwwb-app10.us.archive.org";archive_analytics.values.server_ms=199;</script>
<link type="text/css" rel="stylesheet" href="/static/css/banner-styles.css"/>




  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
  <!-- InstanceBeginEditable name="doctitle" -->
  <title>The World Factbook</title>
  <!-- InstanceEndEditable -->
  <meta name="description" content="">
  <meta name="viewport" content="width=device-width">
  
  
  
  <link href="/web/20140902034709cs_/https://www.cia.gov/library/publications/the-world-factbook/css/fullscreen-external.css" rel="stylesheet" type="text/css">
  <script src="/web/20140902034709js_/https://www.cia.gov/library/publications/the-world-factbook/js/modernizr-latest.js"></script><!--developers version - switch to specific production /web/20140902034709/http://modernizr.com/download/--> 
  <script type="text/javascript" src="/web/20140902034709js_/https://www.cia.gov/library/publications/the-world-factbook/js/jquery-1.8.3.min.js"></script>
  <script type="text/javascript" src="/web/20140902034709js_/https://www.cia.gov/library/publications/the-world-factbook/js/jquery.main.js"></script>
  
  
  <link rel="stylesheet" type="text/css" href="/web/20140902034709cs_/https://www.cia.gov/library/publications/the-world-factbook/styles/jquery.ui.core.css">
  <link rel="stylesheet" type="text/css" href="/web/20140902034709cs_/https://www.cia.gov/library/publications/the-world-factbook/styles/jquery.qtip.css">
  <link rel="stylesheet" type="text/css" href="/web/20140902034709cs_/https://www.cia.gov/library/publications/the-world-factbook/styles/listnav.css"/> 
  <link rel="stylesheet" type="text/css" href="/web/20140902034709cs_/https://www.cia.gov/library/publications/the-world-factbook/styles/navigation.css">
  <link rel="stylesheet" type="text/css" href="/web/20140902034709cs_/https://www.cia.gov/library/publications/the-world-factbook/styles/wfb_styles.css">
  <link rel="stylesheet" type="text/css" href="/web/20140902034709cs_/https://www.cia.gov/library/publications/the-world-factbook/scripts/galleria/themes/classic/galleria.classic.css">
  
  
  <script type="text/javascript" src="/web/20140902034709js_/https://www.cia.gov/library/publications/the-world-factbook/scripts/jquery.idTabs.min.js" charset="utf-8"></script>
  <script type="text/javascript" src="/web/20140902034709js_/https://www.cia.gov/library/publications/the-world-factbook/scripts/jquery.listmenu.js"></script>
  <script type="text/javascript" src="/web/20140902034709js_/https://www.cia.gov/library/publications/the-world-factbook/scripts/jquery.listnav.pack.2.1.js" charset="utf-8"></script>

  <script src="/web/20140902034709js_/https://www.cia.gov/library/publications/the-world-factbook/scripts/jquery-ui-1.9.2.custom.js" type="text/javascript"></script>
  <script type="text/javascript" src="/web/20140902034709js_/https://www.cia.gov/library/publications/the-world-factbook/scripts/jquery.qtip-2.0.js"></script>
  <script type="text/javascript" src="/web/20140902034709js_/https://www.cia.gov/library/publications/the-world-factbook/scripts/jquery.qtip.min.js"></script>
  <script src="/web/20140902034709js_/https://www.cia.gov/library/publications/the-world-factbook/scripts/imgscale.js"></script>
  <script type="text/javascript" src="/web/20140902034709js_/https://www.cia.gov/library/publications/the-world-factbook/scripts/wfb_scripts.js" charset="utf-8"></script> 
  
  <!--[if IE]><script type="text/javascript" src="../js/ie.js"></script><![endif]-->
<!-- InstanceBeginEditable name="head" -->
<!-- load Galleria -->
<link rel="stylesheet" type="text/css" href="/web/20140902034709cs_/https://www.cia.gov/library/publications/the-world-factbook/scripts/galleria/themes/classic/galleria.classic.css">
<script src="/web/20140902034709js_/https://www.cia.gov/library/publications/the-world-factbook/scripts/galleria/galleria-1.2.9.min.js"></script>
<!-- InstanceEndEditable -->
</head>
<body>


<!-- BEGIN WAYBACK TOOLBAR INSERT -->
<script type="text/javascript" src="/static/js/disclaim-element.js" ></script>
<script type="text/javascript" src="/static/js/graph-calc.js" ></script>
<script type="text/javascript">//<![CDATA[
var __wm = (function(imgWidth,imgHeight,yearImgWidth,monthImgWidth){
var wbPrefix = "/web/";
var wbCurrentUrl = "https://www.cia.gov/library/publications/the-world-factbook/geos/aa.html";

var firstYear = 1996;
var displayDay = "2";
var displayMonth = "Sep";
var displayYear = "2014";
var prettyMonths = ["Jan","Feb","Mar","Apr","May","Jun","Jul","Aug","Sep","Oct","Nov","Dec"];
var $D=document,$=function(n){return document.getElementById(n)};
var trackerVal,curYear = -1,curMonth = -1;
var yearTracker,monthTracker;
function showTrackers(val) {
  if (val===trackerVal) return;
  var $ipp=$("wm-ipp");
  var $y=$("displayYearEl"),$m=$("displayMonthEl"),$d=$("displayDayEl");
  if (val) {
    $ipp.className="hi";
  } else {
    $ipp.className="";
    $y.innerHTML=displayYear;$m.innerHTML=displayMonth;$d.innerHTML=displayDay;
  }
  yearTracker.style.display=val?"inline":"none";
  monthTracker.style.display=val?"inline":"none";
  trackerVal = val;
}
function trackMouseMove(event,element) {
  var eventX = getEventX(event);
  var elementX = getElementX(element);
  var xOff = Math.min(Math.max(0, eventX - elementX),imgWidth);
  var monthOff = xOff % yearImgWidth;

  var year = Math.floor(xOff / yearImgWidth);
  var monthOfYear = Math.min(11,Math.floor(monthOff / monthImgWidth));
  // 1 extra border pixel at the left edge of the year:
  var month = (year * 12) + monthOfYear;
  var day = monthOff % 2==1?15:1;
  var dateString = zeroPad(year + firstYear) + zeroPad(monthOfYear+1,2) +
    zeroPad(day,2) + "000000";

  $("displayYearEl").innerHTML=year+firstYear;
  $("displayMonthEl").innerHTML=prettyMonths[monthOfYear];
  // looks too jarring when it changes..
  //$("displayDayEl").innerHTML=zeroPad(day,2);
  var url = wbPrefix + dateString + '/' +  wbCurrentUrl;
  $("wm-graph-anchor").href=url;

  if(curYear != year) {
    var yrOff = year * yearImgWidth;
    yearTracker.style.left = yrOff + "px";
    curYear = year;
  }
  if(curMonth != month) {
    var mtOff = year + (month * monthImgWidth) + 1;
    monthTracker.style.left = mtOff + "px";
    curMonth = month;
  }
}
function hideToolbar() {
  $("wm-ipp").style.display="none";
}
function bootstrap() {
  var $spk=$("wm-ipp-sparkline");
  yearTracker=$D.createElement('div');
  yearTracker.className='yt';
  with(yearTracker.style){
    display='none';width=yearImgWidth+"px";height=imgHeight+"px";
  }
  monthTracker=$D.createElement('div');
  monthTracker.className='mt';
  with(monthTracker.style){
    display='none';width=monthImgWidth+"px";height=imgHeight+"px";
  }
  $spk.appendChild(yearTracker);
  $spk.appendChild(monthTracker);

  var $ipp=$("wm-ipp");
  $ipp&&disclaimElement($ipp);
}
return{st:showTrackers,mv:trackMouseMove,h:hideToolbar,bt:bootstrap};
})(550, 27, 25, 2);//]]>
</script>
<style type="text/css">
body {
  margin-top:0 !important;
  padding-top:0 !important;
  min-width:800px !important;
}
</style>
<div id="wm-ipp" lang="en" style="display:none;">

<div style="position:fixed;left:0;top:0;width:100%!important">
<div id="wm-ipp-inside">
   <table style="width:100%;"><tbody><tr>
   <td id="wm-logo">
       <a href="/web/" title="Wayback Machine home page"><img src="/static/images/toolbar/wayback-toolbar-logo.png" alt="Wayback Machine" width="110" height="39" border="0" /></a>
   </td>
   <td class="c">
       <table style="margin:0 auto;"><tbody><tr>
       <td class="u" colspan="2">
       <form target="_top" method="get" action="/web/form-submit.jsp" name="wmtb" id="wmtb"><input type="text" name="url" id="wmtbURL" value="https://www.cia.gov/library/publications/the-world-factbook/geos/aa.html" style="width:400px;" onfocus="this.focus();this.select();" /><input type="hidden" name="type" value="replay" /><input type="hidden" name="date" value="20140902034709" /><input type="submit" value="Go" /><span id="wm_tb_options" style="display:block;"></span></form>
       </td>
       <td class="n" rowspan="2">
           <table><tbody>
           <!-- NEXT/PREV MONTH NAV AND MONTH INDICATOR -->
           <tr class="m">
           	<td class="b" nowrap="nowrap">
		
		    <a href="/web/20140707160603/http://www.cia.gov/library/publications/the-world-factbook/geos/aa.html" title="7 Jul 2014">JUL</a>
		
		</td>
		<td class="c" id="displayMonthEl" title="You are here: 3:47:09 Sep 2, 2014">SEP</td>
		<td class="f" nowrap="nowrap">
		
		    <a href="/web/20141007042512/https://www.cia.gov/library/publications/the-world-factbook/geos/aa.html" title="7 Oct 2014"><strong>OCT</strong></a>
		
                </td>
	    </tr>
           <!-- NEXT/PREV CAPTURE NAV AND DAY OF MONTH INDICATOR -->
           <tr class="d">
               <td class="b" nowrap="nowrap">
               
                   <a href="/web/20140707160603/http://www.cia.gov/library/publications/the-world-factbook/geos/aa.html" title="16:06:03 Jul 7, 2014"><img src="/static/images/toolbar/wm_tb_prv_on.png" alt="Previous capture" width="14" height="16" border="0" /></a>
               
               </td>
               <td class="c" id="displayDayEl" style="width:34px;font-size:24px;" title="You are here: 3:47:09 Sep 2, 2014">2</td>
	       <td class="f" nowrap="nowrap">
               
		   <a href="/web/20140924095441/https://www.cia.gov/library/publications/the-world-factbook/geos/aa.html" title="9:54:41 Sep 24, 2014"><img src="/static/images/toolbar/wm_tb_nxt_on.png" alt="Next capture" width="14" height="16" border="0" /></a>
	       
	       </td>
           </tr>
           <!-- NEXT/PREV YEAR NAV AND YEAR INDICATOR -->
           <tr class="y">
	       <td class="b" nowrap="nowrap">
               
                   <a href="/web/20130820070442/https://www.cia.gov/library/publications/the-world-factbook/geos/aa.html" title="20 Aug 2013"><strong>2013</strong></a>
               
               </td>
               <td class="c" id="displayYearEl" title="You are here: 3:47:09 Sep 2, 2014">2014</td>
	       <td class="f" nowrap="nowrap">
               
	           <a href="/web/20150905102908/https://www.cia.gov/library/publications/the-world-factbook/geos/aa.html" title="5 Sep 2015"><strong>2015</strong></a>
	       
	       </td>
           </tr>
           </tbody></table>
       </td>
       </tr>
       <tr>
       <td class="s">
           <a class="t" href="/web/20140902034709*/https://www.cia.gov/library/publications/the-world-factbook/geos/aa.html" title="See a list of every capture for this URL">232 captures</a>
           <div class="r" title="Timespan for captures of this URL">12 Jun 07 - 18 Jan 17</div>
       </td>
       <td class="k">
       <a href="" id="wm-graph-anchor">
       <div id="wm-ipp-sparkline" title="Explore captures for this URL">
	 <img id="sparklineImgId" alt="sparklines"
		 onmouseover="__wm.st(1)" onmouseout="__wm.st(0)"
		 onmousemove="__wm.mv(event,this)"
		 width="550"
		 height="27"
		 border="0"
		 src="/web/jsp/graph.jsp?graphdata=550_27_1996:-1:000000000000_1997:-1:000000000000_1998:-1:000000000000_1999:-1:000000000000_2000:-1:000000000000_2001:-1:000000000000_2002:-1:000000000000_2003:-1:000000000000_2004:-1:000000000000_2005:-1:000000000000_2006:-1:000000000000_2007:-1:000001212111_2008:-1:114611113121_2009:-1:211541565744_2010:-1:645564300002_2011:-1:000121061200_2012:-1:15001010f431_2013:-1:451210022420_2014:8:112000202201_2015:-1:011401123011_2016:-1:636230000110_2017:-1:200000000000" />
       </div>
       </a>
       </td>
       </tr></tbody></table>
   </td>
   <td class="r">
       <a href="#close" onclick="__wm.h();return false;" style="background-image:url(/static/images/toolbar/wm_tb_close.png);top:5px;" title="Close the toolbar">Close</a>
       <a href="http://faq.web.archive.org/" style="background-image:url(/static/images/toolbar/wm_tb_help.png);bottom:5px;" title="Get some help using the Wayback Machine">Help</a>
   </td>
   </tr></tbody></table>
</div>
</div>
</div>
<script type="text/javascript">__wm.bt();</script>
<!-- END WAYBACK TOOLBAR INSERT -->
 
  <noscript>Javascript must be enabled for the correct page display</noscript>
  <div id="wrapper">
    <header id="header">
      <div class="header-holder">
       	<a class="skip" accesskey="S" href="#main-content">skip to content</a>
        <span class="bg-globe"></span>
        <div class="header-panel">
          <hgroup>
              <h1 class="logo"><a href="/web/20140902034709/https://www.cia.gov/"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/images/logo.png" alt="Central Intelligence Agency"><span>Central Intelligence Agency</span></a></h1>
              <h2 class="work-text">The Work Of A Nation. The Center of Intelligence.</h2>
          </hgroup>
          <div class="search-form">
            <div class="row">
              <div class="add-nav">
                <ul>
                  <li><a class="active" href="/web/20140902034709/https://www.cia.gov/contact-cia/report-threats.html">Report Threats</a></li>
                  <li><a href="/web/20140902034709/https://www.cia.gov/about-cia/iraqi-rewards-program.html">رعربيعربي</a></li>
                  <li><a href="/web/20140902034709/https://www.cia.gov/contact-cia/index.html" title="A single point of contact for all CIA inquiries.">Contact</a></li>
                </ul>
              </div>
            </div>
            <div class="row">
              <form id="ciaSearchForm" method="get" action="/web/20140902034709/https://www.cia.gov/search">
                <fieldset>
                  <legend class="visuallyhidden">Search CIA.gov</legend>
                  <label class="visuallyhidden" for="q">Search</label>
                  <input name="q" type="text" class="text" id="q" maxlength="2047" placeholder="Search CIA.gov..."/>
				  <input type="hidden" name="site" value="CIA" />
				  <input type="hidden" name="output" value="xml_no_dtd" />
				  <input type="hidden" name="client" value="CIA" />
				  <input type="hidden" name="myAction" value="/search" />
				  <input type="hidden" name="proxystylesheet" value="CIA" />
				  <input type="hidden" name="submitMethod" value="get" />
                  <input type="submit" value="Search" class="submit" />
                </fieldset>
              </form>
            </div>
            <div class="row">              
              <ul class="lang-list">
                <li lang="ar" xml:lang="ar"><a href="/web/20140902034709/https://www.cia.gov/ar/index.html">عربي</a></li>
                <li lang="zh-cn" xml:lang="zh-cn"><a href="/web/20140902034709/https://www.cia.gov/zh/index.html">中文</a></li>
                <li lang="en" xml:lang="en"><a href="/web/20140902034709/https://www.cia.gov/index.html">English</a></li>
                <li lang="fr" xml:lang="fr"><a href="/web/20140902034709/https://www.cia.gov/fr/index.html">Français</a></li>
                <li lang="ru" xml:lang="ru"><a href="/web/20140902034709/https://www.cia.gov/ru/index.html">Русский</a></li>
                <li lang="es" xml:lang="es"><a href="/web/20140902034709/https://www.cia.gov/es/index.html">Español</a></li>
                <li lang="en" xml:lang="en"><a title="additional-info" class="more" href="/web/20140902034709/https://www.cia.gov/foreign-languages/index.html">More<span class="visuallyhidden"> Languages</span></a></li>
              </ul>
            </div>
          </div>   
        </div>
        <nav id="nav">
          <h3 class="visuallyhidden">Navigation</h3>
          <ul>
            <li>
              <a href="/web/20140902034709/https://www.cia.gov/">Home</a>
              <span class="arrow"></span>
            </li>
            <li>
              <a href="/web/20140902034709/https://www.cia.gov/about-cia/">About CIA</a>
              <span class="arrow"></span>
              <div class="drop">
                <ul>
                    <li><a href="/web/20140902034709/https://www.cia.gov/about-cia/todays-cia/index.html">Today's CIA</a></li>
                    <li><a href="/web/20140902034709/https://www.cia.gov/about-cia/leadership/index.html">Leadership</a></li>
                    <li><a href="/web/20140902034709/https://www.cia.gov/about-cia/cia-vision-mission-values/index.html">CIA Vision, Mission &amp; Values</a></li>
                    <li><a href="/web/20140902034709/https://www.cia.gov/about-cia/headquarters-tour/index.html">Tour Headquarters</a></li>
                    <li><a href="/web/20140902034709/https://www.cia.gov/about-cia/cia-museum/index.html">CIA Museum</a></li>
                    <li><a href="/web/20140902034709/https://www.cia.gov/about-cia/history-of-the-cia/index.html">History of the CIA</a></li>
                    <li><a href="/web/20140902034709/https://www.cia.gov/about-cia/faqs/index.html">FAQs</a></li>
                    <li><a href="/web/20140902034709/https://www.cia.gov/about-cia/no-fear-act/index.html">NoFEAR Act</a></li>
                    <li><a href="/web/20140902034709/https://www.cia.gov/about-cia/site-policies/index.html">Site Policies</a></li>
                </ul>
                <div class="info-box">
                  <img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/images/navthumb-about.jpg" alt="About CIA" />
                  <h4>About CIA</h4>
                  <p>Discover the CIA <a href="/web/20140902034709/https://www.cia.gov/about-cia/">history, mission, vision and values</a>. </p>
                </div>
              </div>
            </li>
            <li>
              <a href="/web/20140902034709/https://www.cia.gov/careers/">Careers &amp; Internships</a>
              <span class="arrow"></span>
              <div class="drop">
                <ul>
                    <li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/careers/opportunities" title="This is an overview of all career opportunities at the CIA. "> <span>Career Opportunities </span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/careers/student-opportunities" title="This is the student profile page for candidates looking for jobs/ job listings at the CIA. Student Opportunities - Student Profiles"> <span>Student Opportunities</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/careers/application-process" title="How to apply to the CIA."> <span>Application Process</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/careers/life-at-cia" title="This is the about CIA section of the Careers Site"> <span>Life at CIA</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/careers/diversity" title="This is the diversity information for the Careers Site"> <span>Diversity</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/careers/military-transition" title="Your prior military service could qualify you to continue to serve your nation at the Central Intelligence Agency. Opportunities for qualified applicants are available in the U.S. and abroad."> <span>Military Transition</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/careers/games-information" title=""> <span>Diversions &amp; Information</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/careers/faq" title="Frequently Asked Questions/ FAQ for a Career at the CIA in the Careers Section"> <span>FAQs</span> </a> </li>
                </ul>
                <div class="info-box">
                  <img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/images/navthumb-careers.jpg" alt="Careers &amp; Internships" />
                  <h4>Careers &amp; Internships</h4>
                  <p>Your talent. Your diverse skills. Our mission. Learn more about <a href="/web/20140902034709/https://www.cia.gov/careers/">Careers Opportunities at CIA</a>.</p>
                </div>
              </div>
            </li>
            <li>
              <a href="/web/20140902034709/https://www.cia.gov/offices-of-cia/">Offices of CIA</a>
              <span class="arrow"></span>
              <div class="drop">
                <ul>
                    <li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/offices-of-cia/intelligence-analysis" title=""> <span>Intelligence &amp; Analysis</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/offices-of-cia/clandestine-service" title=""> <span>Clandestine Service</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/offices-of-cia/science-technology" title=""> <span>Science &amp; Technology</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/offices-of-cia/mission-support" title=""> <span>Support to Mission</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/offices-of-cia/human-resources" title=""> <span>Human Resources</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/offices-of-cia/public-affairs" title="Public Affairs"> <span>Public Affairs</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/offices-of-cia/general-counsel" title=""> <span>General Counsel</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/offices-of-cia/equal-employment-opportunity" title=""> <span>Equal Employment Opportunity</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/offices-of-cia/congressional-affairs" title="Office of Congressional Affairs"> <span>Congressional Affairs</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/offices-of-cia/inspector-general" title="Inspector General"> <span>Inspector General</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/offices-of-cia/military-affairs" title="Military Affairs"> <span>Military Affairs</span> </a> </li>
                </ul>
                <div class="info-box">
                  <img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/images/navthumb-offices.jpg" alt="Offices of CIA" />
                  <h4>Offices of CIA</h4>
                  <p><a href="/web/20140902034709/https://www.cia.gov/offices-of-cia/">Learn how the CIA is organized</a> into directorates and key offices, responsible for securing our nation.</p>
                </div>
              </div>
            </li>
            <li>
              <a href="/web/20140902034709/https://www.cia.gov/news-information/">News &amp; Information</a>
              <span class="arrow"></span>
              <div class="drop">
                <ul>
                    <li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/news-information/press-releases-statements" title=""> <span>Press Releases &amp; Statements</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/news-information/speeches-testimony" title=""> <span>Speeches &amp; Testimony</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/news-information/cia-the-war-on-terrorism" title=""> <span>CIA &amp; the War on Terrorism</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/news-information/featured-story-archive" title="index for featured story"> <span>Featured Story Archive</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/news-information/Whats-New-on-CIAgov" title=""> <span>What&#8217;s New Archive</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/news-information/your-news" title=""> <span>Your News</span> </a> </li>
                </ul>
                <div class="info-box">
                  <img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/images/navthumb-news.jpg" alt="News &amp; Information" />
                  <h4>News & Information</h4>
                  <p>The most up-to-date CIA <a href="/web/20140902034709/https://www.cia.gov/news-information/">news, press releases, information and more</a>.</p>
                </div>
              </div>
            </li>
            <li class="active">
              <a href="/web/20140902034709/https://www.cia.gov/library/">Library</a>
              <span class="arrow"></span>
              <div class="drop right">
                <ul>
                    <li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/library/publications" title=""> <span>Publications</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/library/center-for-the-study-of-intelligence" title="CSI section"> <span>Center for the Study of Intelligence</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/library/foia" title=""> <span>Freedom of Information Act Electronic Reading Room</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/library/kent-center-occasional-papers" title=""> <span>Kent Center Occasional Papers</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/library/intelligence-literature" title=""> <span>Intelligence Literature</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/library/reports" title="Reports"> <span>Reports</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/library/related-links.html" title="Related Links"> <span>Related Links</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/library/video-center" title="Repository of CIA videos"> <span>Video Center</span> </a> </li>
                </ul>
                <div class="info-box">
                  <img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/images/navthumb-library.jpg" alt="Library" />
                  <h4>Library</h4>
                  <p>Our <a href="/web/20140902034709/https://www.cia.gov/library/">open-source library</a> houses the thousands of documents, periodicals, maps and reports released to the public.</p>
                </div>
              </div>
            </li>
            <li class="last">
              <a href="/web/20140902034709/https://www.cia.gov/kids-page/">Kids' Zone</a>
              <span class="arrow"></span>
              <div class="drop right">
                <ul>
                   <li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/kids-page/k-5th-grade" title="K-5th Grade"> <span>K-5th Grade</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/kids-page/6-12th-grade" title=""> <span>6-12th Grade</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/kids-page/parents-teachers" title=""> <span>Parents &amp; Teachers</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/kids-page/games" title=""> <span>Games</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/kids-page/related-links" title=""> <span>Related Links</span> </a> </li>
										<li class="plain "> <a class="" href="/web/20140902034709/https://www.cia.gov/kids-page/privacy-statement" title=""> <span>Privacy Statement</span> </a> </li>
                </ul>                
                <div class="info-box">
                  <img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/images/navthumb-kids.jpg" alt="Kids' Zone" />
                  <h4>Kids' Zone</h4>
                  <p><a href="/web/20140902034709/https://www.cia.gov/kids-page/">Learn more about the Agency</a> – and find some top secret things you won't see anywhere else.</p>
                </div>
              </div>
            </li>
          </ul>
        </nav>
      </div>
    </header>
    <div class="main-block">
     	<section id="main">
           <div class="heading-panel">
             <h1>Library</h1>
           </div>           
           <div class="main-holder">
             <div id="sidebar">
               <nav class="sidebar-nav">
                 <h2 class="visuallyhidden">Secondary Navigation</h2>
                 <ul>
                   <li><a class="active" href="/web/20140902034709/https://www.cia.gov/library/">Library</a></li>
                   <li>
                     <a href="/web/20140902034709/https://www.cia.gov/library/publications/">Publications</a>
                     <ul>
                       <li class="mark"><a class="active" href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/">The World Factbook</a></li>
                       <li><a href="/web/20140902034709/https://www.cia.gov/library/publications/world-leaders-1/">World Leaders</a></li>
                       <li><a href="/web/20140902034709/https://www.cia.gov/library/publications/cia-maps-publications/">CIA Maps</a></li>
                       <li><a href="/web/20140902034709/https://www.cia.gov/library/publications/historical-collection-publications/">Historical Collection Publications</a></li>
                       <li><a href="/web/20140902034709/https://www.cia.gov/library/publications/additional-publications/">Additional Publications</a></li>
                     </ul>
                   </li>
                    <li><a href="/web/20140902034709/https://www.cia.gov/library/center-for-the-study-of-intelligence/">Center for the Study of Intelligence</a></li>
                    <li><a href="/web/20140902034709/http://www.foia.cia.gov/">Freedom of Information Act Electronic Reading Room</a></li>
                    <li><a href="/web/20140902034709/https://www.cia.gov/library/kent-center-occasional-papers/">Kent Center Occasional Papers</a></li>
                    <li><a href="/web/20140902034709/https://www.cia.gov/library/intelligence-literature/">Intelligence Literature: Suggested Reading List</a></li>
                    <li><a href="/web/20140902034709/https://www.cia.gov/library/reports/">Reports</a></li>
                    <li><a href="/web/20140902034709/https://www.cia.gov/library/video-center/">Video Center</a></li>
                    <li><a href="/web/20140902034709/https://www.cia.gov/library/related-links.html">Related Links</a></li>
                 </ul>
               </nav>
             </div>
             <div id="content">
               <ul class="breadcrumbs">
                 <li><a href="/web/20140902034709/https://www.cia.gov/">Home</a></li>
                 <li><a href="/web/20140902034709/https://www.cia.gov/library/">Library </a></li>
                 <li><a href="/web/20140902034709/https://www.cia.gov/library/publications/">Publications</a></li>
                 <li>The World Factbook</li>
               </ul>
               <article class="description-box">
               	 <a id="main-content" tabindex="-1">&nbsp;</a>
                 <div class="text-holder-full">
			  <a name="wfbtop"></a>
				<div class="text-box" style="width: 770px; float: left;" id="wfb_data">
				
					<table width="100%" border="0" cellpadding="0" cellspacing="0" >
						<tr>
							<td> 
<table width="100%" border="0" cellspacing="0" cellpadding="0" style="padding-top: 10px;height: 50px;">
		<tr>
			
				<td valign="top"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/banner_ext2.png" border="0" title="World Factbook Title" width="770"  usemap="#Map" style="padding-bottom: 5px;"/></td>
				
		</tr>
		<tr style="height:2px;">
			<td></td>
		</tr>
	</table>
	

<map name="Map" id="Map">
	<area shape="poly" coords="478,17,624,17,615,0,490,0" href="/web/20140902034709/https://www.cia.gov/library/publications/cia-maps-publications/index.html" target="_blank" />
</map>
</td> 
						</tr>
						<tr>
							<td align="right" style="padding-top: 5px; padding-bottom: 10px; background-image:url(/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/country_gradiant_back.jpg); background-position: top; background-repeat: repeat-x;"> 
<meta http-equiv="Content-type" content="text/html; charset=utf-8" />
<meta name="viewport" content="initial-scale=1, maximum-scale=1, user-scalable=no">
<meta http-equiv="X-UA-Compatible" content="IE=Edge" />



<script>
  	$(document).ready(function() {

//		$(".selecter_links").selecter({
//        	defaultLabel: "Please Select a Country to View",
//			links: true
//		});
		$( ".selecter_links" ).change(function(e) {
			if (this.form.selecter_links.selectedIndex > 0)	{
				window.location = this.form.selecter_links.options[this.form.selecter_links.selectedIndex].value;
			}
		});

	});
  
</script>
<div class="option_table_wrapper">
	<form action="#" method="GET">
		<select name="selecter_links" class="selecter_links">
		<option value="">Please select a country to view</option>
		
				<option value="../geos/xx.html"> World </option>
			
				<option value="../geos/af.html"> Afghanistan </option>
			
				<option value="../geos/ax.html"> Akrotiri </option>
			
				<option value="../geos/al.html"> Albania </option>
			
				<option value="../geos/ag.html"> Algeria </option>
			
				<option value="../geos/aq.html"> American Samoa </option>
			
				<option value="../geos/an.html"> Andorra </option>
			
				<option value="../geos/ao.html"> Angola </option>
			
				<option value="../geos/av.html"> Anguilla </option>
			
				<option value="../geos/ay.html"> Antarctica </option>
			
				<option value="../geos/ac.html"> Antigua and Barbuda </option>
			
				<option value="../geos/xq.html"> Arctic Ocean </option>
			
				<option value="../geos/ar.html"> Argentina </option>
			
				<option value="../geos/am.html"> Armenia </option>
			
				<option value="../geos/aa.html"> Aruba </option>
			
				<option value="../geos/at.html"> Ashmore and Cartier Islands </option>
			
				<option value="../geos/zh.html"> Atlantic Ocean </option>
			
				<option value="../geos/as.html"> Australia </option>
			
				<option value="../geos/au.html"> Austria </option>
			
				<option value="../geos/aj.html"> Azerbaijan </option>
			
				<option value="../geos/bf.html"> Bahamas, The </option>
			
				<option value="../geos/ba.html"> Bahrain </option>
			
				<option value="../geos/um.html"> Baker Island </option>
			
				<option value="../geos/bg.html"> Bangladesh </option>
			
				<option value="../geos/bb.html"> Barbados </option>
			
				<option value="../geos/bo.html"> Belarus </option>
			
				<option value="../geos/be.html"> Belgium </option>
			
				<option value="../geos/bh.html"> Belize </option>
			
				<option value="../geos/bn.html"> Benin </option>
			
				<option value="../geos/bd.html"> Bermuda </option>
			
				<option value="../geos/bt.html"> Bhutan </option>
			
				<option value="../geos/bl.html"> Bolivia </option>
			
				<option value="../geos/bk.html"> Bosnia and Herzegovina </option>
			
				<option value="../geos/bc.html"> Botswana </option>
			
				<option value="../geos/bv.html"> Bouvet Island </option>
			
				<option value="../geos/br.html"> Brazil </option>
			
				<option value="../geos/io.html"> British Indian Ocean Territory </option>
			
				<option value="../geos/vi.html"> British Virgin Islands </option>
			
				<option value="../geos/bx.html"> Brunei </option>
			
				<option value="../geos/bu.html"> Bulgaria </option>
			
				<option value="../geos/uv.html"> Burkina Faso </option>
			
				<option value="../geos/bm.html"> Burma </option>
			
				<option value="../geos/by.html"> Burundi </option>
			
				<option value="../geos/cv.html"> Cabo Verde </option>
			
				<option value="../geos/cb.html"> Cambodia </option>
			
				<option value="../geos/cm.html"> Cameroon </option>
			
				<option value="../geos/ca.html"> Canada </option>
			
				<option value="../geos/cj.html"> Cayman Islands </option>
			
				<option value="../geos/ct.html"> Central African Republic </option>
			
				<option value="../geos/cd.html"> Chad </option>
			
				<option value="../geos/ci.html"> Chile </option>
			
				<option value="../geos/ch.html"> China </option>
			
				<option value="../geos/kt.html"> Christmas Island </option>
			
				<option value="../geos/ip.html"> Clipperton Island </option>
			
				<option value="../geos/ck.html"> Cocos (Keeling) Islands </option>
			
				<option value="../geos/co.html"> Colombia </option>
			
				<option value="../geos/cn.html"> Comoros </option>
			
				<option value="../geos/cg.html"> Congo, Democratic Republic of the </option>
			
				<option value="../geos/cf.html"> Congo, Republic of the </option>
			
				<option value="../geos/cw.html"> Cook Islands </option>
			
				<option value="../geos/cr.html"> Coral Sea Islands </option>
			
				<option value="../geos/cs.html"> Costa Rica </option>
			
				<option value="../geos/iv.html"> Cote d'Ivoire </option>
			
				<option value="../geos/hr.html"> Croatia </option>
			
				<option value="../geos/cu.html"> Cuba </option>
			
				<option value="../geos/cc.html"> Curacao </option>
			
				<option value="../geos/cy.html"> Cyprus </option>
			
				<option value="../geos/ez.html"> Czech Republic </option>
			
				<option value="../geos/da.html"> Denmark </option>
			
				<option value="../geos/dx.html"> Dhekelia </option>
			
				<option value="../geos/dj.html"> Djibouti </option>
			
				<option value="../geos/do.html"> Dominica </option>
			
				<option value="../geos/dr.html"> Dominican Republic </option>
			
				<option value="../geos/ec.html"> Ecuador </option>
			
				<option value="../geos/eg.html"> Egypt </option>
			
				<option value="../geos/es.html"> El Salvador </option>
			
				<option value="../geos/ek.html"> Equatorial Guinea </option>
			
				<option value="../geos/er.html"> Eritrea </option>
			
				<option value="../geos/en.html"> Estonia </option>
			
				<option value="../geos/et.html"> Ethiopia </option>
			
				<option value="../geos/fk.html"> Falkland Islands (Islas Malvinas) </option>
			
				<option value="../geos/fo.html"> Faroe Islands </option>
			
				<option value="../geos/fj.html"> Fiji </option>
			
				<option value="../geos/fi.html"> Finland </option>
			
				<option value="../geos/fr.html"> France </option>
			
				<option value="../geos/fp.html"> French Polynesia </option>
			
				<option value="../geos/fs.html"> French Southern and Antarctic Lands </option>
			
				<option value="../geos/gb.html"> Gabon </option>
			
				<option value="../geos/ga.html"> Gambia, The </option>
			
				<option value="../geos/gz.html"> Gaza Strip </option>
			
				<option value="../geos/gg.html"> Georgia </option>
			
				<option value="../geos/gm.html"> Germany </option>
			
				<option value="../geos/gh.html"> Ghana </option>
			
				<option value="../geos/gi.html"> Gibraltar </option>
			
				<option value="../geos/gr.html"> Greece </option>
			
				<option value="../geos/gl.html"> Greenland </option>
			
				<option value="../geos/gj.html"> Grenada </option>
			
				<option value="../geos/gq.html"> Guam </option>
			
				<option value="../geos/gt.html"> Guatemala </option>
			
				<option value="../geos/gk.html"> Guernsey </option>
			
				<option value="../geos/gv.html"> Guinea </option>
			
				<option value="../geos/pu.html"> Guinea-Bissau </option>
			
				<option value="../geos/gy.html"> Guyana </option>
			
				<option value="../geos/ha.html"> Haiti </option>
			
				<option value="../geos/hm.html"> Heard Island and McDonald Islands </option>
			
				<option value="../geos/vt.html"> Holy See (Vatican City) </option>
			
				<option value="../geos/ho.html"> Honduras </option>
			
				<option value="../geos/hk.html"> Hong Kong </option>
			
				<option value="../geos/um.html"> Howland Island </option>
			
				<option value="../geos/hu.html"> Hungary </option>
			
				<option value="../geos/ic.html"> Iceland </option>
			
				<option value="../geos/in.html"> India </option>
			
				<option value="../geos/xo.html"> Indian Ocean </option>
			
				<option value="../geos/id.html"> Indonesia </option>
			
				<option value="../geos/ir.html"> Iran </option>
			
				<option value="../geos/iz.html"> Iraq </option>
			
				<option value="../geos/ei.html"> Ireland </option>
			
				<option value="../geos/im.html"> Isle of Man </option>
			
				<option value="../geos/is.html"> Israel </option>
			
				<option value="../geos/it.html"> Italy </option>
			
				<option value="../geos/jm.html"> Jamaica </option>
			
				<option value="../geos/jn.html"> Jan Mayen </option>
			
				<option value="../geos/ja.html"> Japan </option>
			
				<option value="../geos/um.html"> Jarvis Island </option>
			
				<option value="../geos/je.html"> Jersey </option>
			
				<option value="../geos/um.html"> Johnston Atoll </option>
			
				<option value="../geos/jo.html"> Jordan </option>
			
				<option value="../geos/kz.html"> Kazakhstan </option>
			
				<option value="../geos/ke.html"> Kenya </option>
			
				<option value="../geos/um.html"> Kingman Reef </option>
			
				<option value="../geos/kr.html"> Kiribati </option>
			
				<option value="../geos/kn.html"> Korea, North </option>
			
				<option value="../geos/ks.html"> Korea, South </option>
			
				<option value="../geos/kv.html"> Kosovo </option>
			
				<option value="../geos/ku.html"> Kuwait </option>
			
				<option value="../geos/kg.html"> Kyrgyzstan </option>
			
				<option value="../geos/la.html"> Laos </option>
			
				<option value="../geos/lg.html"> Latvia </option>
			
				<option value="../geos/le.html"> Lebanon </option>
			
				<option value="../geos/lt.html"> Lesotho </option>
			
				<option value="../geos/li.html"> Liberia </option>
			
				<option value="../geos/ly.html"> Libya </option>
			
				<option value="../geos/ls.html"> Liechtenstein </option>
			
				<option value="../geos/lh.html"> Lithuania </option>
			
				<option value="../geos/lu.html"> Luxembourg </option>
			
				<option value="../geos/mc.html"> Macau </option>
			
				<option value="../geos/mk.html"> Macedonia </option>
			
				<option value="../geos/ma.html"> Madagascar </option>
			
				<option value="../geos/mi.html"> Malawi </option>
			
				<option value="../geos/my.html"> Malaysia </option>
			
				<option value="../geos/mv.html"> Maldives </option>
			
				<option value="../geos/ml.html"> Mali </option>
			
				<option value="../geos/mt.html"> Malta </option>
			
				<option value="../geos/rm.html"> Marshall Islands </option>
			
				<option value="../geos/mr.html"> Mauritania </option>
			
				<option value="../geos/mp.html"> Mauritius </option>
			
				<option value="../geos/mx.html"> Mexico </option>
			
				<option value="../geos/fm.html"> Micronesia, Federated States of </option>
			
				<option value="../geos/um.html"> Midway Islands </option>
			
				<option value="../geos/md.html"> Moldova </option>
			
				<option value="../geos/mn.html"> Monaco </option>
			
				<option value="../geos/mg.html"> Mongolia </option>
			
				<option value="../geos/mj.html"> Montenegro </option>
			
				<option value="../geos/mh.html"> Montserrat </option>
			
				<option value="../geos/mo.html"> Morocco </option>
			
				<option value="../geos/mz.html"> Mozambique </option>
			
				<option value="../geos/wa.html"> Namibia </option>
			
				<option value="../geos/nr.html"> Nauru </option>
			
				<option value="../geos/bq.html"> Navassa Island </option>
			
				<option value="../geos/np.html"> Nepal </option>
			
				<option value="../geos/nl.html"> Netherlands </option>
			
				<option value="../geos/nc.html"> New Caledonia </option>
			
				<option value="../geos/nz.html"> New Zealand </option>
			
				<option value="../geos/nu.html"> Nicaragua </option>
			
				<option value="../geos/ng.html"> Niger </option>
			
				<option value="../geos/ni.html"> Nigeria </option>
			
				<option value="../geos/ne.html"> Niue </option>
			
				<option value="../geos/nf.html"> Norfolk Island </option>
			
				<option value="../geos/cq.html"> Northern Mariana Islands </option>
			
				<option value="../geos/no.html"> Norway </option>
			
				<option value="../geos/mu.html"> Oman </option>
			
				<option value="../geos/zn.html"> Pacific Ocean </option>
			
				<option value="../geos/pk.html"> Pakistan </option>
			
				<option value="../geos/ps.html"> Palau </option>
			
				<option value="../geos/um.html"> Palmyra Atoll </option>
			
				<option value="../geos/pm.html"> Panama </option>
			
				<option value="../geos/pp.html"> Papua New Guinea </option>
			
				<option value="../geos/pf.html"> Paracel Islands </option>
			
				<option value="../geos/pa.html"> Paraguay </option>
			
				<option value="../geos/pe.html"> Peru </option>
			
				<option value="../geos/rp.html"> Philippines </option>
			
				<option value="../geos/pc.html"> Pitcairn Islands </option>
			
				<option value="../geos/pl.html"> Poland </option>
			
				<option value="../geos/po.html"> Portugal </option>
			
				<option value="../geos/rq.html"> Puerto Rico </option>
			
				<option value="../geos/qa.html"> Qatar </option>
			
				<option value="../geos/ro.html"> Romania </option>
			
				<option value="../geos/rs.html"> Russia </option>
			
				<option value="../geos/rw.html"> Rwanda </option>
			
				<option value="../geos/tb.html"> Saint Barthelemy </option>
			
				<option value="../geos/sh.html"> Saint Helena, Ascension, and Tristan da Cunha </option>
			
				<option value="../geos/sc.html"> Saint Kitts and Nevis </option>
			
				<option value="../geos/st.html"> Saint Lucia </option>
			
				<option value="../geos/rn.html"> Saint Martin </option>
			
				<option value="../geos/sb.html"> Saint Pierre and Miquelon </option>
			
				<option value="../geos/vc.html"> Saint Vincent and the Grenadines </option>
			
				<option value="../geos/ws.html"> Samoa </option>
			
				<option value="../geos/sm.html"> San Marino </option>
			
				<option value="../geos/tp.html"> Sao Tome and Principe </option>
			
				<option value="../geos/sa.html"> Saudi Arabia </option>
			
				<option value="../geos/sg.html"> Senegal </option>
			
				<option value="../geos/ri.html"> Serbia </option>
			
				<option value="../geos/se.html"> Seychelles </option>
			
				<option value="../geos/sl.html"> Sierra Leone </option>
			
				<option value="../geos/sn.html"> Singapore </option>
			
				<option value="../geos/sk.html"> Sint Maarten </option>
			
				<option value="../geos/lo.html"> Slovakia </option>
			
				<option value="../geos/si.html"> Slovenia </option>
			
				<option value="../geos/bp.html"> Solomon Islands </option>
			
				<option value="../geos/so.html"> Somalia </option>
			
				<option value="../geos/sf.html"> South Africa </option>
			
				<option value="../geos/oo.html"> Southern Ocean </option>
			
				<option value="../geos/sx.html"> South Georgia and South Sandwich Islands </option>
			
				<option value="../geos/od.html"> South Sudan </option>
			
				<option value="../geos/sp.html"> Spain </option>
			
				<option value="../geos/pg.html"> Spratly Islands </option>
			
				<option value="../geos/ce.html"> Sri Lanka </option>
			
				<option value="../geos/su.html"> Sudan </option>
			
				<option value="../geos/ns.html"> Suriname </option>
			
				<option value="../geos/sv.html"> Svalbard </option>
			
				<option value="../geos/wz.html"> Swaziland </option>
			
				<option value="../geos/sw.html"> Sweden </option>
			
				<option value="../geos/sz.html"> Switzerland </option>
			
				<option value="../geos/sy.html"> Syria </option>
			
				<option value="../geos/tw.html"> Taiwan </option>
			
				<option value="../geos/ti.html"> Tajikistan </option>
			
				<option value="../geos/tz.html"> Tanzania </option>
			
				<option value="../geos/th.html"> Thailand </option>
			
				<option value="../geos/tt.html"> Timor-Leste </option>
			
				<option value="../geos/to.html"> Togo </option>
			
				<option value="../geos/tl.html"> Tokelau </option>
			
				<option value="../geos/tn.html"> Tonga </option>
			
				<option value="../geos/td.html"> Trinidad and Tobago </option>
			
				<option value="../geos/ts.html"> Tunisia </option>
			
				<option value="../geos/tu.html"> Turkey </option>
			
				<option value="../geos/tx.html"> Turkmenistan </option>
			
				<option value="../geos/tk.html"> Turks and Caicos Islands </option>
			
				<option value="../geos/tv.html"> Tuvalu </option>
			
				<option value="../geos/ug.html"> Uganda </option>
			
				<option value="../geos/up.html"> Ukraine </option>
			
				<option value="../geos/ae.html"> United Arab Emirates </option>
			
				<option value="../geos/uk.html"> United Kingdom </option>
			
				<option value="../geos/us.html"> United States </option>
			
				<option value="../geos/um.html"> United States Pacific Island Wildlife Refuges </option>
			
				<option value="../geos/uy.html"> Uruguay </option>
			
				<option value="../geos/uz.html"> Uzbekistan </option>
			
				<option value="../geos/nh.html"> Vanuatu </option>
			
				<option value="../geos/ve.html"> Venezuela </option>
			
				<option value="../geos/vm.html"> Vietnam </option>
			
				<option value="../geos/vq.html"> Virgin Islands </option>
			
				<option value="../geos/wq.html"> Wake Island </option>
			
				<option value="../geos/wf.html"> Wallis and Futuna </option>
			
				<option value="../geos/we.html"> West Bank </option>
			
				<option value="../geos/wi.html"> Western Sahara </option>
			
				<option value="../geos/ym.html"> Yemen </option>
			
				<option value="../geos/za.html"> Zambia </option>
			
				<option value="../geos/zi.html"> Zimbabwe </option>
			
				<option value="../geos/ee.html"> European Union </option>
			
		</select>
	</form>
</div>

							</td>
						</tr>
						<tr>
							<td> 
<style>
.description-box .text-holder-full .text-box { line-height: 12px; }
</style>
<link rel="stylesheet" type="text/css" href="/web/20140902034709cs_/https://www.cia.gov/library/publications/the-world-factbook/styles/navigation.css">
<script type="text/javascript">
var timeout         = 500;
var closetimer		= 0;
var ddmenuitem      = 0;

function wfbNav_open()
{	wfbNav_canceltimer();
	wfbNav_close();
	ddmenuitem = $(this).find('ul').eq(0).css('visibility', 'visible');}

function wfbNav_close()
{	if(ddmenuitem) ddmenuitem.css('visibility', 'hidden');}

function wfbNav_timer()
{	closetimer = window.setTimeout(wfbNav_close, timeout);}

function wfbNav_canceltimer()
{	if(closetimer)
	{	window.clearTimeout(closetimer);
		closetimer = null;}}

$(document).ready(function()
{	$('#wfbNav > li').bind('mouseover', wfbNav_open);
	$('#wfbNav > li').bind('mouseout',  wfbNav_timer);});

document.onclick = wfbNav_close;
</script>

	<div>
		<ul id="wfbNav" style="z-index: 9999;">
			<li style="border-bottom: 2px solid #CCCCCC; "><a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/index.html" style="width:20px; height: 12px;" title="The World Factbook Home"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/home_on.png" border="0"></a></li>
			<li style="border-bottom: 2px solid #CCCCCC; "><a href="javascript:void(0);" style="width:65px;" title="About">ABOUT</a>
				<ul class="sub_menu">
					<li><a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/history.html">&nbsp;&nbsp;History</a></li>
					<li><a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/contributor_copyright.html">&nbsp;&nbsp;Copyright and Contributors</a></li>
					<li><a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/purchase_info.html">&nbsp;&nbsp;Purchasing</a></li>
					<li><a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/didyouknow.html">&nbsp;&nbsp;Did You Know?</a></li>
					
				</ul>
			</li>
			<li style="border-bottom: 2px solid #CCCCCC; "><a href="javascript:void(0);"  style="width:95px;" title="References">REFERENCES</a>
				<ul class="sub_menu">
					<li><a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/refmaps.html">&nbsp;&nbsp;Regional and World Maps</a></li>
					<li><a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/flagsoftheworld.html">&nbsp;&nbsp;Flags of the World</a></li>
					<li><a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/gallery.html">&nbsp;&nbsp;Gallery of Covers</a></li>
					<li><a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html">&nbsp;&nbsp;Definitions and Notes</a></li>
					<li><a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/profileguide.html">&nbsp;&nbsp;Guide to Country Profiles</a></li>
					<li><a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/rankorderguide.html">&nbsp;&nbsp;Guide to Country Comparisons</a></li>
					
					<li><a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/guidetowfbook.html">&nbsp;&nbsp;The World Factbook Users Guide</a></li>
				</ul>
			</li>
			<li style="border-bottom: 2px solid #CCCCCC; "><a href="javascript:void(0);" title="Appendices">APPENDICES</a>
				<ul class="sub_menu">
					<li><a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/appendix/appendix-a.html">&nbsp;&nbsp;A: abbreviations</a></li>
					<li><a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/appendix/appendix-b.html">&nbsp;&nbsp;B: international organizations and groups</a></li>
					<li><a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/appendix/appendix-c.html">&nbsp;&nbsp;C: selected international environmental agreements</a></li>
					<li><a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/appendix/appendix-d.html">&nbsp;&nbsp;D: cross-reference list of country data codes</a></li>
					<li><a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/appendix/appendix-e.html">&nbsp;&nbsp;E: cross-reference list of hydrographic data codes</a></li>
					<li><a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/appendix/appendix-f.html">&nbsp;&nbsp;F: cross-reference list of geographic names</a></li>
					<li><a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/appendix/appendix-g.html">&nbsp;&nbsp;G: weights and measures</a></li>
				</ul>
			</li>
			<li id="faqs" style="border-bottom: 2px solid #CCCCCC; "><a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/faqs.html" style="cursor:pointer;width:50px;">FAQ<span style="text-transform:lowercase;">s</span></a></li>
			
			
				<li id="contact" style="border-bottom: 2px solid #CCCCCC; "> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/contact.html" title="Contact" style="cursor:pointer;width:73px;"> CONTACT </a> </li>
			
			
			
			
			
		</ul>
	</div>
	<div class="smalltext_nav" align="right" valign="bottom" style="border-bottom: 2px solid #CCCCCC; height: 22px;">
		
			<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/print/textversion.html"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/bandwidth_btn_off.gif" alt="View Text Low Bandwidth Version" title="View Text Low Bandwidth Version" width="173" height="8" border="0" /></a><br>
			
			<a href="/web/20140902034709/https://www.cia.gov/library/publications/download/"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/download_btn_off.gif" alt="Download Publication" title="Download Publication"  border="0" id="Download Publication" /></a>
		
	</div>
	<div class="clear"> </div>
	</div>
</td>
						</tr>
						<tr> 
							<td>
								<!-- InstanceBeginEditable name="mainContent" -->
								
<script src="/web/20140902034709js_/https://www.cia.gov/library/publications/the-world-factbook/scripts/imgscale.js"></script>

<script language="javascript" type="text/javascript">
<!--
function MM_swapImgRestore() { //v3.0
  var i,x,a=document.MM_sr; for(i=0;a&&i<a.length&&(x=a[i])&&x.oSrc;i++) x.src=x.oSrc;
}
function MM_preloadImages() { //v3.0
  var d=document; if(d.images){ if(!d.MM_p) d.MM_p=new Array();
    var i,j=d.MM_p.length,a=MM_preloadImages.arguments; for(i=0; i<a.length; i++)
    if (a[i].indexOf("#")!=0){ d.MM_p[j]=new Image; d.MM_p[j++].src=a[i];}}
}

function MM_findObj(n, d) { //v4.01
  var p,i,x;  if(!d) d=document; if((p=n.indexOf("?"))>0&&parent.frames.length) {
    d=parent.frames[n.substring(p+1)].document; n=n.substring(0,p);}
  if(!(x=d[n])&&d.all) x=d.all[n]; for (i=0;!x&&i<d.forms.length;i++) x=d.forms[i][n];
  for(i=0;!x&&d.layers&&i<d.layers.length;i++) x=MM_findObj(n,d.layers[i].document);
  if(!x && d.getElementById) x=d.getElementById(n); return x;
}

function MM_swapImage() { //v3.0
  var i,j=0,x,a=MM_swapImage.arguments; document.MM_sr=new Array; for(i=0;i<(a.length-2);i+=3)
   if ((x=MM_findObj(a[i]))!=null){document.MM_sr[j++]=x; if(!x.oSrc) x.oSrc=x.src; x.src=a[i+2];}
}
//-->
</script>



			<div id="print">
				
					<a href="print_countrydata_holder.html" target="_blank"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/print.gif" alt="Print Page" width="25" height="18" border="0" title="Print Page" style="color:#999999" /></a>
				
			</div>
		
	<table width="100%" border="0" cellpadding="0" cellspacing="0">

				<tr class="cam_dark">
					<td  valign="middle">
						<table border="0" cellpadding="0" cellspacing="0">
							<tr>
								<td height="30" valign="middle">
									<div class="region1">
										<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/wfbExt/region_cam.html" style="color: #FFFFFF;">Central America and Caribbean</a> <strong>:: </strong><span class="region_name1">Aruba</span> 
									</div>
									
										<div class="affiliation"><em>(part of the Kingdom of the Netherlands)</em></div>
									
								</td>
							</tr>
						</table>
					</td>
				<td width="20" align="right" valign="middle" class="cam_dark"><a href="print/country/countrypdf_aa.pdf"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/print.gif" style="padding: 3px;"></a></td>
				</tr>
		</table>
	
				<table width="100%" border="0" cellspacing="0" cellpadding="0" style="background-image:url(/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/cam_lgmap_bkgrnd.jpg); background-repeat: repeat-x; background-position: top ;">
			<tr>
				<td width="323" align="center" valign="top" style="  "><table width="100%" align="center" style="border: 1px solid #ccc; height: 195px;" >
						<tr>
							<td colspan="2" align="left" valign="middle" class="smalltext_nav" style="height: 12px;" >Page last updated on June 23, 2014 </td>
						</tr>
						<tr>
							<td height="230" align="center" valign="middle" class="area" style="width: 50%; height: 120px;">
							
							
								
									<a href="javascript:void(0);" title="Click flag for description"> <img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/flags/large/aa-lgflag.gif" 
												border="0" 
												style="cursor:pointer; border: 1px solid #CCC; "  
												id="flagDialog2_aa" 
												name="aa" 
												regioncode="cam" 
												countrycode="aa"  
												countryname="Aruba" 
												flagsubfield="" 
												countryaffiliation="(part of the Kingdom of the Netherlands)"
												flagdescription="blue, with two narrow, horizontal, yellow stripes across the lower portion and a red, four-pointed star outlined in white in the upper hoist-side corner; the star represents Aruba and its red soil and white beaches, its four points the four major languages (Papiamento, Dutch, Spanish, English) as well as the four points of a compass, to indicate that its inhabitants come from all over the world; the blue symbolizes Caribbean waters and skies; the stripes represent the island's two main &quot;industries&quot;: the flow of tourists to the sun-drenched beaches and the flow of minerals from the earth " 
												flagdescriptionnote="" 
												region="Central America and Caribbean" 
												class="flagFit cam_lgflagborder"
												typeimage = "flag"></a>
								
							</td>
							<td align="center" valign="middle" class="area" style="width: 50%; height: 120px;"><a href="javascript:void(0);" title="Click locator to enlarge"> <img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/locator/cam/aa_large_locator.gif" 
											border="0" 
											style="cursor:pointer; border: 1px solid #CCC;" 
											id="locatorDialog2_aa" 
											name="aa" 
											regioncode="cam" 
											countrycode="aa"  
											countryname="Aruba" 
											flagsubfield="" 
											countryaffiliation=""
											flagdescription="" 
											flagdescriptionnote="" 
											region="Central America and Caribbean" 
											class="locatorFit cam_lgflagborder"
											typeimage = "locator"></a></td>
						</tr>
					</table></td>
				<td width="1%" rowspan="2" align="center" valign="middle" bgcolor="#FFFFFF" style="border: 1px solid #fff;">&nbsp;</td>
				<td rowspan="2" align="center" valign="middle" style="border: 1px solid #E4D4D4;">
					<div align="center" valign="middle" > <a href="javascript:void(0);" title="Click map to enlarge"> <img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/maps/aa-map.gif" 
					border="0" 
					style="cursor:pointer; border: 1px solid #CCC; display: block; "  
					id="mapDialog2_aa" 
					name="aa"  
					regioncode="cam" 
					countrycode="aa"  
					countryname="Aruba" 
					flagsubfield="" 
					countryaffiliation=""
					flagdescription="" 
					flagdescriptionnote="" 
					region="Central America and Caribbean" 
					class="mapFit cam_lgflagborder"
					typeimage = "map"></a>
					</div></td>
			</tr>
			<tr>
			
				<td height="140" align="center" valign="top" class="photo_bkgrnd_static" bgcolor="#FFFFFF">
					<table width="100%" border="0" align="left" cellpadding="0" cellspacing="0">
						<tr>
							<td height="10" colspan="3"></td>
						</tr>
						<tr>
							
							<td width="100%" rowspan="3" align="center" valign="middle" class="smalltext_nav" >				
								
										<a href= "javascript:void(0);" title="Photos of Aruba" > 
										<img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/photo_on.gif" 
											name="aa" 
											regioncode="cam" 
											countrycode="aa"  
											countryname="Aruba" 
											region="Central America and Caribbean" 
											width="123" 
											height="81"  
											border="0"
											id="photoDialog"
											 style="padding-top:10px;"
											/></a>
										<div class="smalltext_nav"> view <span style="color: #007d7d;; letter-spacing:1px;">
											<a href="javascript:void(0);" name="aa" 
											regioncode="cam" 
											countrycode="aa"  
											countryname="Aruba" 
											region="Central America and Caribbean" 
											width="123" 
											height="81"  
											border="0"
											class="photoDialog" ><strong>8 
													photos
													
											</strong></a>
											</span>
											of <br>Aruba
										</div>
										
							</td>
						
						</tr>
						<tr>
							<td height="50%" align="center" valign="top" ></td>
						</tr>
					</table>
					
				</td>
			</tr>
		</table>
	
	<div id="countryInfo" style="display: none;">
		
				<div class="wrapper">
					<div style="float:right" class="expand_all">
						<a href="javascript:void(0)" class="expand">EXPAND ALL</a><a href="javascript:void(0)" class="collapse" style="display: none;">COLLAPSE ALL</a> 
					</div>
				</div>
		
<script src="jClocksGMT-master/js/jClocksGMT.js"></script>
<script src="jClocksGMT-master/js/jquery.rotate.js"></script>
<link rel="stylesheet" href="jClocksGMT-master/css/jClocksGMT.css">



		
<script>
	
		$(document).ready(function() { 
			
				$('[id^="CollapsiblePanel1"] h2').css({'background-color':'#cce5e5',"border-bottom":"2px solid white","cursor":"pointer"}); // cam 			
			
		});
	
   </script>



<div id="CollapsiblePanel1_Intro" class="CollapsiblePanel" style="width:100%; ">
<div class="wrapper">
<h2 class="question question-back" ccode="aa" sectiontitle="Introduction"><a href="javascript:void();">Introduction</span> ::</span><span class="region">Aruba</span></a></h2>
<div class="answer" align="left">
	<div class="box" style="padding: 0px; margin: 0px;">
		<ul style="text-align: left;padding: 0px;margin: 0px;width: 100%;">
			<table border="0" cellspacing="0" cellpadding="0"  style="width: 100%;">
				 
				
					<tr class="cam_light" >
					
					<td width="450" height="20"><div class="category" style="padding-left:5px;" id="field"> 
										<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2028&amp;alphaletter=B&amp;term=Background" title="Notes and Definitions: Background"> Background</a>:
									 </div></td>
					
						<td align="right">
						
						
						
								<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2028.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"> <img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0"  style="text-decoration:none;"> </a>
								
						</tr>
						
					
					<tr>
					
					<td id="data" colspan="2" style="vertical-align:middle;">
					
					
					
					
						<div class="category_data">Discovered and claimed for Spain in 1499, Aruba was acquired by the Dutch in 1636. The island's economy has been dominated by three main industries. A 19th century gold rush was followed by prosperity brought on by the opening in 1924 of an oil refinery. The last decades of the 20th century saw a boom in the tourism industry. Aruba seceded from the Netherlands Antilles in 1986 and became a separate, autonomous member of the Kingdom of the Netherlands. Movement toward full independence was halted at Aruba's request in 1990.</div>
					
				<tr>
					<td class="category_data" style="padding-bottom: 5px;"></td>
				</tr>
			</table>
		</ul>
	</div>
	
</div>

<script src="jClocksGMT-master/js/jClocksGMT.js"></script>
<script src="jClocksGMT-master/js/jquery.rotate.js"></script>
<link rel="stylesheet" href="jClocksGMT-master/css/jClocksGMT.css">



		
<script>
	
		$(document).ready(function() { 
			
				$('[id^="CollapsiblePanel1"] h2').css({'background-color':'#cce5e5',"border-bottom":"2px solid white","cursor":"pointer"}); // cam 			
			
		});
	
   </script>



<div id="CollapsiblePanel1_Geo" class="CollapsiblePanel" style="width:100%; ">
<div class="wrapper">
<h2 class="question question-back" ccode="aa" sectiontitle="Geography"><a href="javascript:void();">Geography</span> ::</span><span class="region">Aruba</span></a></h2>
<div class="answer" align="left">
	<div class="box" style="padding: 0px; margin: 0px;">
		<ul style="text-align: left;padding: 0px;margin: 0px;width: 100%;">
			<table border="0" cellspacing="0" cellpadding="0"  style="width: 100%;">
				 
				
					<tr class="cam_light" >
					
					<td width="450" height="20"><div class="category" style="padding-left:5px;" id="field"> 
										<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2144&amp;alphaletter=L&amp;term=Location" title="Notes and Definitions: Location"> Location</a>:
									 </div></td>
					
						<td align="right">
						
						
						
								<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2144.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"> <img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0"  style="text-decoration:none;"> </a>
								
						</tr>
						
					
					<tr>
					
					<td id="data" colspan="2" style="vertical-align:middle;">
					
					
					
					
						<div class="category_data">Caribbean, island in the Caribbean Sea, north of Venezuela</div>
					
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2011&amp;alphaletter=G&amp;term=Geographic%20coordinates" title="Notes and Definitions: Geographic coordinates"> Geographic coordinates:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2011.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">12 30 N, 69 58 W</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2145&amp;alphaletter=M&amp;term=Map%20references" title="Notes and Definitions: Map references"> Map references:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2145.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">
										
  <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/graphics/ref_maps/physical/pdf/central_america.pdf" target="_blank" class="category_data">Central America and the Caribbean</a>
  
									</div>
									
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2147&amp;alphaletter=A&amp;term=Area" title="Notes and Definitions: Area"> Area:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2147.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
							<div class="category" style="padding-top: 2px;">
								total:
									
								
								<span class="category_data" style="font-weight:normal; vertical-align:bottom;">180 sq km</span></div>
								
								
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2147rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=218#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 218 </a> </span>
											
								<div class="category" style="padding-top: 2px;">
									land:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">180 sq km </span></div>
								
								<div class="category" style="padding-top: 2px;">
									water:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">0 sq km </span></div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2023&amp;alphaletter=A&amp;term=Area%20-%20comparative" title="Notes and Definitions: Area - comparative"> Area - comparative:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2023.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
											<div class="category_data">slightly larger than Washington, DC</div>
										
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2096&amp;alphaletter=L&amp;term=Land%20boundaries" title="Notes and Definitions: Land boundaries"> Land boundaries:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2096.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">0 km</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2060&amp;alphaletter=C&amp;term=Coastline" title="Notes and Definitions: Coastline"> Coastline:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2060.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">68.5 km</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2106&amp;alphaletter=M&amp;term=Maritime%20claims" title="Notes and Definitions: Maritime claims"> Maritime claims:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2106.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
							<div class="category" style="padding-top: 2px;">
								territorial sea:
									
								
								<span class="category_data" style="font-weight:normal; vertical-align:bottom;">12 nm</span></div>
								
								
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2059&amp;alphaletter=C&amp;term=Climate" title="Notes and Definitions: Climate"> Climate:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2059.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">tropical marine; little seasonal temperature variation</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2125&amp;alphaletter=T&amp;term=Terrain" title="Notes and Definitions: Terrain"> Terrain:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2125.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">flat with a few hills; scant vegetation</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2020&amp;alphaletter=E&amp;term=Elevation%20extremes" title="Notes and Definitions: Elevation extremes"> Elevation extremes:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2020.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
							<div class="category" style="padding-top: 2px;">
								lowest point:
									
								
								<span class="category_data" style="font-weight:normal; vertical-align:bottom;">Caribbean Sea 0 m</span></div>
								
								
								
								<div class="category" style="padding-top: 2px;">
									highest point:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">Ceru Jamanota 188 m </span></div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2111&amp;alphaletter=N&amp;term=Natural%20resources" title="Notes and Definitions: Natural resources"> Natural resources:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2111.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">NEGL; white sandy beaches</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2097&amp;alphaletter=L&amp;term=Land%20use" title="Notes and Definitions: Land use"> Land use:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2097.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
							<div class="category" style="padding-top: 2px;">
								arable land:
									
								
								<span class="category_data" style="font-weight:normal; vertical-align:bottom;">11.11%</span></div>
								
								
								
								<div class="category" style="padding-top: 2px;">
									permanent crops:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">0% </span></div>
								
								<div class="category" style="padding-top: 2px;">
									other:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">88.89% (2005) </span></div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2146&amp;alphaletter=I&amp;term=Irrigated%20land" title="Notes and Definitions: Irrigated land"> Irrigated land:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2146.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">NA</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2021&amp;alphaletter=N&amp;term=Natural%20hazards" title="Notes and Definitions: Natural hazards"> Natural hazards:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2021.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">hurricanes; lies outside the Caribbean hurricane belt and is rarely threatened</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2032&amp;alphaletter=E&amp;term=Environment%20-%20current%20issues" title="Notes and Definitions: Environment - current issues"> Environment - current issues:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2032.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">NA</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2113&amp;alphaletter=G&amp;term=Geography%20-%20note" title="Notes and Definitions: Geography - note"> Geography - note:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2113.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">a flat, riverless island renowned for its white sand beaches; its tropical climate is moderated by constant trade winds from the Atlantic Ocean; the temperature is almost constant at about 27 degrees Celsius (81 degrees Fahrenheit)</div>
								
				<tr>
					<td class="category_data" style="padding-bottom: 5px;"></td>
				</tr>
			</table>
		</ul>
	</div>
	
</div>

<script src="jClocksGMT-master/js/jClocksGMT.js"></script>
<script src="jClocksGMT-master/js/jquery.rotate.js"></script>
<link rel="stylesheet" href="jClocksGMT-master/css/jClocksGMT.css">



		
<script>
	
		$(document).ready(function() { 
			
				$('[id^="CollapsiblePanel1"] h2').css({'background-color':'#cce5e5',"border-bottom":"2px solid white","cursor":"pointer"}); // cam 			
			
		});
	
   </script>



<div id="CollapsiblePanel1_People" class="CollapsiblePanel" style="width:100%; ">
<div class="wrapper">
<h2 class="question question-back" ccode="aa" sectiontitle="People and Society"><a href="javascript:void();">People and Society</span> ::</span><span class="region">Aruba</span></a></h2>
<div class="answer" align="left">
	<div class="box" style="padding: 0px; margin: 0px;">
		<ul style="text-align: left;padding: 0px;margin: 0px;width: 100%;">
			<table border="0" cellspacing="0" cellpadding="0"  style="width: 100%;">
				 
				
					<tr class="cam_light" >
					
					<td width="450" height="20"><div class="category" style="padding-left:5px;" id="field"> 
										<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2110&amp;alphaletter=N&amp;term=Nationality" title="Notes and Definitions: Nationality"> Nationality</a>:
									 </div></td>
					
						<td align="right">
						
						
						
								<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2110.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"> <img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0"  style="text-decoration:none;"> </a>
								
						</tr>
						
					
					<tr>
					
					<td id="data" colspan="2" style="vertical-align:middle;">
					
					
					
					
						<div class="category">noun: <span class="category_data" style="font-weight:normal;">Aruban(s)</span> </div>
						
								<div class="category" style="padding-top: 2px;">
									adjective:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">Aruban; Dutch </span></div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2075&amp;alphaletter=E&amp;term=Ethnic%20groups" title="Notes and Definitions: Ethnic groups"> Ethnic groups:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2075.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">Dutch 82.1%, Colombian 6.6%, Venezuelan 2.2%, Dominican 2.2%, Haitian 1.2%, other 5.5%, unspecified 0.1% (2010 est.)</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2098&amp;alphaletter=L&amp;term=Languages" title="Notes and Definitions: Languages"> Languages:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2098.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">Papiamento (a Spanish-Portuguese-Dutch-English dialect) 69.4%, Spanish 13.7%, English (widely spoken) 7.1%, Dutch (official) 6.1%, Chinese 1.5%, other 1.7%, unspecified 0.4% (2010 est.)</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2122&amp;alphaletter=R&amp;term=Religions" title="Notes and Definitions: Religions"> Religions:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2122.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">Roman Catholic 75.3%, Protestant 4.9% (includes Methodist .9%, Adventist .9%, Anglican .4%, other Protestant 2.7%), Jehovah's Witness 1.7%, other 12%, none 5.5%, unspecified 0.5% (2010 est.)</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2119&amp;alphaletter=P&amp;term=Population" title="Notes and Definitions: Population"> Population:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2119.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">110,663 (July 2014 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2119rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=190#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 190 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2010&amp;alphaletter=A&amp;term=Age%20structure" title="Notes and Definitions: Age structure"> Age structure:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2010.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
							<div class="category" style="padding-top: 2px;">
								0-14 years:
									
								
								<span class="category_data" style="font-weight:normal; vertical-align:bottom;">17.8% (male 9,852/female 9,797)</span></div>
								
								
								
								<div class="category" style="padding-top: 2px;">
									15-24 years:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">13.5% (male 7,469/female 7,427) </span></div>
								
								<div class="category" style="padding-top: 2px;">
									25-54 years:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">43% (male 22,981/female 24,615) </span></div>
								
								<div class="category" style="padding-top: 2px;">
									55-64 years:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">13.5% (male 6,804/female 8,093) </span></div>
								
								<div class="category" style="padding-top: 2px;">
									65 years and over:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">12.3% (male 5,346/female 8,279) (2014 est.) </span></div>
								
										<div class="category">
										<span style="margin-bottom:0px; vertical-align:bottom;">population pyramid:</span> <a href="javascript:void();" title="<img src = '../graphics/populationpyramid_tooltip.gif'>"> <img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/poppyramid_icon.jpg" 
											border="0" 
											style="cursor:pointer; border: 0px solid #CCC;" 
											id="flagDialog2_aa" 
											name="aa" 
											regioncode="cam" 
											countrycode="aa"  
											countryname="Aruba" 
											flagsubfield="" 
											countryaffiliation="(part of the Kingdom of the Netherlands)"
											flagdescription="blue, with two narrow, horizontal, yellow stripes across the lower portion and a red, four-pointed star outlined in white in the upper hoist-side corner; the star represents Aruba and its red soil and white beaches, its four points the four major languages (Papiamento, Dutch, Spanish, English) as well as the four points of a compass, to indicate that its inhabitants come from all over the world; the blue symbolizes Caribbean waters and skies; the stripes represent the island's two main &quot;industries&quot;: the flow of tourists to the sun-drenched beaches and the flow of minerals from the earth" 
											flagdescriptionnote="" 
											region="Central America and Caribbean" 
											typeimage="population"
											> </a>
									
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2261&amp;alphaletter=D&amp;term=Dependency%20ratios" title="Notes and Definitions: Dependency ratios"> Dependency ratios:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2261.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
							<div class="category" style="padding-top: 2px;">
								total dependency ratio:
									
								
								<span class="category_data" style="font-weight:normal; vertical-align:bottom;">44.1 %</span></div>
								
								
								
								<div class="category" style="padding-top: 2px;">
									youth dependency ratio:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">27.1 % </span></div>
								
								<div class="category" style="padding-top: 2px;">
									elderly dependency ratio:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">16.9 % </span></div>
								
								<div class="category" style="padding-top: 2px;">
									potential support ratio:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">5.9 (2014 est.) </span></div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2177&amp;alphaletter=M&amp;term=Median%20age" title="Notes and Definitions: Median age"> Median age:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2177.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
							<div class="category" style="padding-top: 2px;">
								total:
									
								
								<span class="category_data" style="font-weight:normal; vertical-align:bottom;">38.8 years</span></div>
								
								
								
								<div class="category" style="padding-top: 2px;">
									male:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">36.9 years </span></div>
								
								<div class="category" style="padding-top: 2px;">
									female:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">40.6 years (2014 est.) </span></div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2002&amp;alphaletter=P&amp;term=Population%20growth%20rate" title="Notes and Definitions: Population growth rate"> Population growth rate:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2002.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">1.36% (2014 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2002rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=90#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 90 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2054&amp;alphaletter=B&amp;term=Birth%20rate" title="Notes and Definitions: Birth rate"> Birth rate:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2054.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">12.65 births/1,000 population (2014 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2054rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=158#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 158 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2066&amp;alphaletter=D&amp;term=Death%20rate" title="Notes and Definitions: Death rate"> Death rate:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2066.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">8.09 deaths/1,000 population (2014 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2066rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=96#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 96 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2112&amp;alphaletter=N&amp;term=Net%20migration%20rate" title="Notes and Definitions: Net migration rate"> Net migration rate:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2112.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">9.04 migrant(s)/1,000 population (2014 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2112rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=15#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 15 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2212&amp;alphaletter=U&amp;term=Urbanization" title="Notes and Definitions: Urbanization"> Urbanization:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2212.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
							<div class="category" style="padding-top: 2px;">
								urban population:
									
								
								<span class="category_data" style="font-weight:normal; vertical-align:bottom;">46.8% of total population (2011)</span></div>
								
								
								
								<div class="category" style="padding-top: 2px;">
									rate of urbanization:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">0.54% annual rate of change (2010-15 est.) </span></div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2219&amp;alphaletter=M&amp;term=Major%20urban%20areas%20-%20population" title="Notes and Definitions: Major urban areas - population"> Major urban areas - population:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2219.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">ORANJESTAD (capital) 37,000 (2011)</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2018&amp;alphaletter=S&amp;term=Sex%20ratio" title="Notes and Definitions: Sex ratio"> Sex ratio:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2018.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
							<div class="category" style="padding-top: 2px;">
								at birth:
									
								
								<span class="category_data" style="font-weight:normal; vertical-align:bottom;">1.02 male(s)/female</span></div>
								
								
								
								<div class="category" style="padding-top: 2px;">
									0-14 years:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">1.01 male(s)/female </span></div>
								
								<div class="category" style="padding-top: 2px;">
									15-24 years:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">1.01 male(s)/female </span></div>
								
								<div class="category" style="padding-top: 2px;">
									25-54 years:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">0.93 male(s)/female </span></div>
								
								<div class="category" style="padding-top: 2px;">
									55-64 years:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">0.9 male(s)/female </span></div>
								
								<div class="category" style="padding-top: 2px;">
									65 years and over:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">0.65 male(s)/female </span></div>
								
								<div class="category" style="padding-top: 2px;">
									total population:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">0.9 male(s)/female (2014 est.) </span></div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2091&amp;alphaletter=I&amp;term=Infant%20mortality%20rate" title="Notes and Definitions: Infant mortality rate"> Infant mortality rate:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2091.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
							<div class="category" style="padding-top: 2px;">
								total:
									
								
								<span class="category_data" style="font-weight:normal; vertical-align:bottom;">11.74 deaths/1,000 live births</span></div>
								
								
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2091rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=128#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 128 </a> </span>
											
								<div class="category" style="padding-top: 2px;">
									male:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">15.44 deaths/1,000 live births </span></div>
								
								<div class="category" style="padding-top: 2px;">
									female:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">7.97 deaths/1,000 live births (2014 est.) </span></div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2102&amp;alphaletter=L&amp;term=Life%20expectancy%20at%20birth" title="Notes and Definitions: Life expectancy at birth"> Life expectancy at birth:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2102.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
							<div class="category" style="padding-top: 2px;">
								total population:
									
								
								<span class="category_data" style="font-weight:normal; vertical-align:bottom;">76.35 years</span></div>
								
								
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2102rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=82#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 82 </a> </span>
											
								<div class="category" style="padding-top: 2px;">
									male:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">73.3 years </span></div>
								
								<div class="category" style="padding-top: 2px;">
									female:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">79.47 years (2014 est.) </span></div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2127&amp;alphaletter=T&amp;term=Total%20fertility%20rate" title="Notes and Definitions: Total fertility rate"> Total fertility rate:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2127.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">1.84 children born/woman (2014 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2127rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=150#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 150 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2216&amp;alphaletter=D&amp;term=Drinking%20water%20source" title="Notes and Definitions: Drinking water source"> Drinking water source:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2216.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
							<div class="category" style="padding-top: 2px;">
								improved:
									
								
								<span class="category_data" style="font-weight:normal; vertical-align:bottom;"></span></div>
								
								
								
								<div class="category_data" style="padding-top: 3px;">urban: 97.8% of population </div>
							
								<div class="category_data" style="padding-top: 3px;">rural: 97.8% of population </div>
							
								<div class="category_data" style="padding-top: 3px;">total: 97.8% of population </div>
							
								<div class="category" style="padding-top: 2px;">
									unimproved:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;"> </span></div>
								
								<div class="category_data" style="padding-top: 3px;">urban: 2.2% of population </div>
							
								<div class="category_data" style="padding-top: 3px;">rural: 2.2% of population </div>
							
								<div class="category_data" style="padding-top: 3px;">total: 2.2% of population (2012 est.) </div>
							
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2217&amp;alphaletter=S&amp;term=Sanitation%20facility%20access" title="Notes and Definitions: Sanitation facility access"> Sanitation facility access:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2217.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
							<div class="category" style="padding-top: 2px;">
								improved:
									
								
								<span class="category_data" style="font-weight:normal; vertical-align:bottom;"></span></div>
								
								
								
								<div class="category_data" style="padding-top: 3px;">urban: 97.7% of population </div>
							
								<div class="category_data" style="padding-top: 3px;">rural: 97.7% of population </div>
							
								<div class="category_data" style="padding-top: 3px;">total: 97.7% of population </div>
							
								<div class="category" style="padding-top: 2px;">
									unimproved:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;"> </span></div>
								
								<div class="category_data" style="padding-top: 3px;">urban: 2.3% of population </div>
							
								<div class="category_data" style="padding-top: 3px;">rural: 2.3% of population </div>
							
								<div class="category_data" style="padding-top: 3px;">total: 2.3% of population (2012 est.) </div>
							
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2155&amp;alphaletter=H&amp;term=HIV/AIDS%20-%20adult%20prevalence%20rate" title="Notes and Definitions: HIV/AIDS - adult prevalence rate"> HIV/AIDS - adult prevalence rate:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2155.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">NA</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2156&amp;alphaletter=H&amp;term=HIV/AIDS%20-%20people%20living%20with%20HIV/AIDS" title="Notes and Definitions: HIV/AIDS - people living with HIV/AIDS"> HIV/AIDS - people living with HIV/AIDS:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2156.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">NA</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2157&amp;alphaletter=H&amp;term=HIV/AIDS%20-%20deaths" title="Notes and Definitions: HIV/AIDS - deaths"> HIV/AIDS - deaths:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2157.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">NA</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2206&amp;alphaletter=E&amp;term=Education%20expenditures" title="Notes and Definitions: Education expenditures"> Education expenditures:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2206.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">6% of GDP (2011)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2206rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=41#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 41 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2103&amp;alphaletter=L&amp;term=Literacy" title="Notes and Definitions: Literacy"> Literacy:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2103.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
							<div class="category" style="padding-top: 2px;">
								definition:
									
								
								<span class="category_data" style="font-weight:normal; vertical-align:bottom;">age 15 and over can read and write</span></div>
								
								
								
								<div class="category" style="padding-top: 2px;">
									total population:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">96.8% </span></div>
								
								<div class="category" style="padding-top: 2px;">
									male:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">96.9% </span></div>
								
								<div class="category" style="padding-top: 2px;">
									female:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">96.7% (2010 est.) </span></div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2205&amp;alphaletter=S&amp;term=School%20life%20expectancy%20(primary%20to%20tertiary%20education)" title="Notes and Definitions: School life expectancy (primary to tertiary education)"> School life expectancy (primary to tertiary education):</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2205.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
							<div class="category" style="padding-top: 2px;">
								total:
									
								
								<span class="category_data" style="font-weight:normal; vertical-align:bottom;">13 years</span></div>
								
								
								
								<div class="category" style="padding-top: 2px;">
									male:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">13 years </span></div>
								
								<div class="category" style="padding-top: 2px;">
									female:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">14 years (2011) </span></div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2229&amp;alphaletter=U&amp;term=Unemployment,%20youth%20ages%2015-24" title="Notes and Definitions: Unemployment, youth ages 15-24"> Unemployment, youth ages 15-24:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2229.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
							<div class="category" style="padding-top: 2px;">
								total:
									
								
								<span class="category_data" style="font-weight:normal; vertical-align:bottom;">23.2%</span></div>
								
								
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2229rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=47#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 47 </a> </span>
											
								<div class="category" style="padding-top: 2px;">
									male:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">24.1% </span></div>
								
								<div class="category" style="padding-top: 2px;">
									female:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">22.9% (2007) </span></div>
								
				<tr>
					<td class="category_data" style="padding-bottom: 5px;"></td>
				</tr>
			</table>
		</ul>
	</div>
	
</div>

<script src="jClocksGMT-master/js/jClocksGMT.js"></script>
<script src="jClocksGMT-master/js/jquery.rotate.js"></script>
<link rel="stylesheet" href="jClocksGMT-master/css/jClocksGMT.css">



		
<script>
	
		$(document).ready(function() { 
			
				$('[id^="CollapsiblePanel1"] h2').css({'background-color':'#cce5e5',"border-bottom":"2px solid white","cursor":"pointer"}); // cam 			
			
		});
	
   </script>



<div id="CollapsiblePanel1_Govt" class="CollapsiblePanel" style="width:100%; ">
<div class="wrapper">
<h2 class="question question-back" ccode="aa" sectiontitle="Government"><a href="javascript:void();">Government</span> ::</span><span class="region">Aruba</span></a></h2>
<div class="answer" align="left">
	<div class="box" style="padding: 0px; margin: 0px;">
		<ul style="text-align: left;padding: 0px;margin: 0px;width: 100%;">
			<table border="0" cellspacing="0" cellpadding="0"  style="width: 100%;">
				 
				
					<tr class="cam_light" >
					
					<td width="450" height="20"><div class="category" style="padding-left:5px;" id="field"> 
										<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2142&amp;alphaletter=C&amp;term=Country%20name" title="Notes and Definitions: Country name"> Country name</a>:
									 </div></td>
					
						<td align="right">
						
						
						
								<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2142.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"> <img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0"  style="text-decoration:none;"> </a>
								
						</tr>
						
					
					<tr>
					
					<td id="data" colspan="2" style="vertical-align:middle;">
					
					
					
					
						<div class="category">conventional long form: <span class="category_data" style="font-weight:normal;">none</span> </div>
						
								<div class="category" style="padding-top: 2px;">
									conventional short form:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">Aruba </span></div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2006&amp;alphaletter=D&amp;term=Dependency%20status" title="Notes and Definitions: Dependency status"> Dependency status:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2006.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">constituent country of the Kingdom of the Netherlands; full autonomy in internal affairs obtained in 1986 upon separation from the Netherlands Antilles; Dutch Government responsible for defense and foreign affairs</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2128&amp;alphaletter=G&amp;term=Government%20type" title="Notes and Definitions: Government type"> Government type:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2128.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">parliamentary democracy</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2057&amp;alphaletter=C&amp;term=Capital" title="Notes and Definitions: Capital"> Capital:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2057.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
							<div class="category" style="padding-top: 2px;">
								name:
									
								
								<span class="category_data" style="font-weight:normal; vertical-align:bottom;">Oranjestad</span></div>
								
								
								
								<div class="category" style="padding-top: 2px;">
									geographic coordinates:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">12 31 N, 70 02 W </span></div>
								
								<div class="category" style="padding-top: 2px;">
									time difference:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">UTC-4 (1 hour ahead of Washington, DC, during Standard Time) </span></div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2051&amp;alphaletter=A&amp;term=Administrative%20divisions" title="Notes and Definitions: Administrative divisions"> Administrative divisions:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2051.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">none (part of the Kingdom of the Netherlands)</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2088&amp;alphaletter=I&amp;term=Independence" title="Notes and Definitions: Independence"> Independence:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2088.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">none (part of the Kingdom of the Netherlands)</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2109&amp;alphaletter=N&amp;term=National%20holiday" title="Notes and Definitions: National holiday"> National holiday:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2109.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">Flag Day, 18 March (1976)</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2063&amp;alphaletter=C&amp;term=Constitution" title="Notes and Definitions: Constitution"> Constitution:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2063.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">previous 1947, 1955; latest drafted and approved August 1985, enacted 1 January 1986 (regulates governance of Aruba, but is subordinate to the Charter for the Kingdom of the Netherlands); note - in October 2010, following dissolution of the Netherlands Antilles, Aruba became a constituent country within the Kingdom of the Netherlands (2013)</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2100&amp;alphaletter=L&amp;term=Legal%20system" title="Notes and Definitions: Legal system"> Legal system:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2100.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">civil law system based on the Dutch civil code</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2123&amp;alphaletter=S&amp;term=Suffrage" title="Notes and Definitions: Suffrage"> Suffrage:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2123.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">18 years of age; universal</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2077&amp;alphaletter=E&amp;term=Executive%20branch" title="Notes and Definitions: Executive branch"> Executive branch:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2077.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
							<div class="category" style="padding-top: 2px;">
								chief of state:
									
								
								<span class="category_data" style="font-weight:normal; vertical-align:bottom;">King WILLEM-ALEXANDER of the Netherlands (since 30 April 2013); represented by Governor General Fredis REFUNJOL (since 11 May 2004)</span></div>
								
								
								
								<div class="category" style="padding-top: 2px;">
									head of government:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">Prime Minister Michiel "Mike" Godfried EMAN (since 30 October 2009) </span></div>
								
								<div class="category" style="padding-top: 2px;">
									cabinet:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">Council of Ministers elected by the Staten </span></div>
								
									<span class="category" style="padding-left:7px;font-weight: normal;">(For more information visit the <a href="/web/20140902034709/https://www.cia.gov/library/publications/world-leaders-1/AA.html" target="_blank">World Leaders website</a>&nbsp;<img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/cam_newwindow.gif" alt="Opens in New Window" title="Opens in New Window" border="0"/>)</span>
									
								<div class="category" style="padding-top: 2px;">
									elections:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">the monarchy is hereditary; governor general appointed for a six-year term by the monarch; prime minister and deputy prime minister elected by the Staten for four-year terms; election last held on 25 September 2009 (next to be held by September 2013) </span></div>
								
								<div class="category" style="padding-top: 2px;">
									election results:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">Michiel "Mike" Godfried EMAN elected prime minister; percent of legislative vote - NA </span></div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2101&amp;alphaletter=L&amp;term=Legislative%20branch" title="Notes and Definitions: Legislative branch"> Legislative branch:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2101.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">unicameral Legislature or Staten (21 seats; members elected by direct popular vote to serve four-year terms)</div>
								
								<div class="category" style="padding-top: 2px;">
									elections:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">last held on 27 September 2013 (next to be held in 2017) </span></div>
								
								<div class="category" style="padding-top: 2px;">
									election results:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">percent of vote by party - NA; seats by party - AVP 13, MEP 8 </span></div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2094&amp;alphaletter=J&amp;term=Judicial%20branch" title="Notes and Definitions: Judicial branch"> Judicial branch:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2094.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
							<div class="category" style="padding-top: 2px;">
								highest court(s):
									
								
								<span class="category_data" style="font-weight:normal; vertical-align:bottom;">Joint Court of Justice of Aruba, Curacao, Sint Maarten, and of Bonaire, Sint Eustatitus and Saba or  "Joint Court of Justice" (consists of the presiding judge, NA members, and NA substitutes); final appeals heard by the Supreme Court, in The Hague, Netherlands</span></div>
								
								
								
								<div class="category_data" style="padding-top: 3px;">note - prior to 2010, the Joint Court of Justice was the Common Court of Justice of the  Netherlands Antilles and Aruba </div>
							
								<div class="category" style="padding-top: 2px;">
									judge selection and term of office:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">Joint Court judges appointed by the monarch for life </span></div>
								
								<div class="category" style="padding-top: 2px;">
									subordinate courts:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">Courts in First Instance </span></div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2118&amp;alphaletter=P&amp;term=Political%20parties%20and%20leaders" title="Notes and Definitions: Political parties and leaders"> Political parties and leaders:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2118.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">Aliansa/Aruban Social Movement or MSA [Robert WEVER]</div>
								
								<div class="category_data" style="padding-top: 3px;">Aruban Liberal Organization or OLA [Glenbert CROES] </div>
							
								<div class="category_data" style="padding-top: 3px;">Aruban Patriotic Movement or MPA [Monica ARENDS-KOCK] </div>
							
								<div class="category_data" style="padding-top: 3px;">Aruban Patriotic Party or PPA [Benny NISBET] </div>
							
								<div class="category_data" style="padding-top: 3px;">Aruban People's Party or AVP [Michiel "Mike" EMAN] </div>
							
								<div class="category_data" style="padding-top: 3px;">People's Electoral Movement Party or MEP [Nelson O. ODUBER] </div>
							
								<div class="category_data" style="padding-top: 3px;">Real Democracy or PDR [Andin BIKKER] </div>
							
								<div class="category_data" style="padding-top: 3px;">RED [Rudy LAMPE] </div>
							
								<div class="category_data" style="padding-top: 3px;">Workers Political Platform or PTT [Gregorio WOLFF] </div>
							
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2115&amp;alphaletter=P&amp;term=Political%20pressure%20groups%20and%20leaders" title="Notes and Definitions: Political pressure groups and leaders"> Political pressure groups and leaders:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2115.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
							<div class="category" style="padding-top: 2px;">
								other:
									
								
								<span class="category_data" style="font-weight:normal; vertical-align:bottom;">environmental groups</span></div>
								
								
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2107&amp;alphaletter=I&amp;term=International%20organization%20participation" title="Notes and Definitions: International organization participation"> International organization participation:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2107.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">Caricom (observer), FATF, ILO, IMF, Interpol, IOC, ITUC (NGOs), UNESCO (associate), UNWTO (associate), UPU</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2149&amp;alphaletter=D&amp;term=Diplomatic%20representation%20in%20the%20US" title="Notes and Definitions: Diplomatic representation in the US"> Diplomatic representation in the US:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2149.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">none (represented by the Kingdom of the Netherlands); note - Mr. Henry BAARH, Minister Plenipotentiary for Aruba at the Embassy of the Kingdom of the Netherlands</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2007&amp;alphaletter=D&amp;term=Diplomatic%20representation%20from%20the%20US" title="Notes and Definitions: Diplomatic representation from the US"> Diplomatic representation from the US:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2007.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">the US does not have an embassy in Aruba; the Consul General to Curacao, currently Consul General Valerie BELON, is accredited to Aruba</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2081&amp;alphaletter=F&amp;term=Flag%20description" title="Notes and Definitions: Flag description"> Flag description:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2081.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">blue, with two narrow, horizontal, yellow stripes across the lower portion and a red, four-pointed star outlined in white in the upper hoist-side corner; the star represents Aruba and its red soil and white beaches, its four points the four major languages (Papiamento, Dutch, Spanish, English) as well as the four points of a compass, to indicate that its inhabitants come from all over the world; the blue symbolizes Caribbean waters and skies; the stripes represent the island's two main "industries": the flow of tourists to the sun-drenched beaches and the flow of minerals from the earth</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2218&amp;alphaletter=N&amp;term=National%20anthem" title="Notes and Definitions: National anthem"> National anthem:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2218.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
							<div class="category" style="padding-top: 2px;">
								name:
									
								
								<span class="category_data" style="font-weight:normal; vertical-align:bottom;">"Aruba Deshi Tera" (Aruba Precious Country)</span></div>
								
								
								
										<link rel="stylesheet" type="text/css" href="/web/20140902034709cs_/https://www.cia.gov/library/publications/the-world-factbook/styles/jquery.ui.core.css"/>
										<p><a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/anthems/AA.mp3" class="playAnthem" name="Aruba" target="_new"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/play_anthem.png"></a></p>
										
										
								<div class="category" style="padding-top: 2px;">
									lyrics/music:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">Juan Chabaya 'Padu' LAMPE/Rufo Inocencio WEVER </span></div>
								
								<div class="category" style="padding-top: 2px;">
									
										<em>note:</em>
										
									<span class="category_data" style="font-weight:normal; vertical-align:top;">local anthem adopted 1986; as part of the Kingdom of the Netherlands, "Het Wilhelmus" is official (see Netherlands) </span></div>
								
				<tr>
					<td class="category_data" style="padding-bottom: 5px;"></td>
				</tr>
			</table>
		</ul>
	</div>
	
</div>

<script src="jClocksGMT-master/js/jClocksGMT.js"></script>
<script src="jClocksGMT-master/js/jquery.rotate.js"></script>
<link rel="stylesheet" href="jClocksGMT-master/css/jClocksGMT.css">



		
<script>
	
		$(document).ready(function() { 
			
				$('[id^="CollapsiblePanel1"] h2').css({'background-color':'#cce5e5',"border-bottom":"2px solid white","cursor":"pointer"}); // cam 			
			
		});
	
   </script>



<div id="CollapsiblePanel1_Econ" class="CollapsiblePanel" style="width:100%; ">
<div class="wrapper">
<h2 class="question question-back" ccode="aa" sectiontitle="Economy"><a href="javascript:void();">Economy</span> ::</span><span class="region">Aruba</span></a></h2>
<div class="answer" align="left">
	<div class="box" style="padding: 0px; margin: 0px;">
		<ul style="text-align: left;padding: 0px;margin: 0px;width: 100%;">
			<table border="0" cellspacing="0" cellpadding="0"  style="width: 100%;">
				 
				
					<tr class="cam_light" >
					
					<td width="450" height="20"><div class="category" style="padding-left:5px;" id="field"> 
										<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2116&amp;alphaletter=E&amp;term=Economy%20-%20overview" title="Notes and Definitions: Economy - overview"> Economy - overview</a>:
									 </div></td>
					
						<td align="right">
						
						
						
								<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2116.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"> <img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0"  style="text-decoration:none;"> </a>
								
						</tr>
						
					
					<tr>
					
					<td id="data" colspan="2" style="vertical-align:middle;">
					
					
					
					
						<div class="category_data">Tourism and offshore banking are the mainstays of the small open Aruban economy. Tourist arrivals have rebounded strongly following a dip after the 11 September 2001 attacks. Tourism now accounts for over 80 % of economic activity. Over 1.5 million tourists per year visit Aruba, with 75% of those from the US. The rapid growth of the tourism sector has resulted in a substantial expansion of other activities. Construction continues to boom with hotel capacity five times the 1985 level. Aruba is heavily dependent on imports and is making efforts to expand exports to achieve a more desirable trade balance. Aruba weathered two major shocks in recent years: fallout from the global financial crisis, which had its largest impact on tourism, and the closure of its oil refinery in 2009. Economic recovery is progressing gradually, but output is still 12% below its pre-crisis level. Aruba’s banking sector withstood the recession well, and unemployment has significantly decreased.</div>
					
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2001&amp;alphaletter=G&amp;term=GDP%20(purchasing%20power%20parity)" title="Notes and Definitions: GDP (purchasing power parity)"> GDP (purchasing power parity):</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2001.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">$2.516 billion (2009 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2001rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=186#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 186 </a> </span>
											
								<div class="category_data" style="padding-top: 3px;">$2.258 billion (2005 est.) </div>
							
								<div class="category_data" style="padding-top: 3px;">$2.205 billion (2004 est.) </div>
							
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2195&amp;alphaletter=G&amp;term=GDP%20(official%20exchange%20rate)" title="Notes and Definitions: GDP (official exchange rate)"> GDP (official exchange rate):</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2195.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">$2.516 billion (2009 est.)</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2003&amp;alphaletter=G&amp;term=GDP%20-%20real%20growth%20rate" title="Notes and Definitions: GDP - real growth rate"> GDP - real growth rate:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2003.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">2.4% (2005 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2003rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=133#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 133 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2004&amp;alphaletter=G&amp;term=GDP%20-%20per%20capita%20(PPP)" title="Notes and Definitions: GDP - per capita (PPP)"> GDP - per capita (PPP):</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2004.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">$25,300 (2011 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2004rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=59#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 59 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2012&amp;alphaletter=G&amp;term=GDP%20-%20composition,%20by%20sector%20of%20origin" title="Notes and Definitions: GDP - composition, by sector of origin"> GDP - composition, by sector of origin:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2012.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
							<div class="category" style="padding-top: 2px;">
								agriculture:
									
								
								<span class="category_data" style="font-weight:normal; vertical-align:bottom;">0.4%</span></div>
								
								
								
								<div class="category" style="padding-top: 2px;">
									industry:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">33.3% </span></div>
								
								<div class="category" style="padding-top: 2px;">
									services:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">66.3% (2002 est.) </span></div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2052&amp;alphaletter=A&amp;term=Agriculture%20-%20products" title="Notes and Definitions: Agriculture - products"> Agriculture - products:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2052.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">aloes; livestock; fish</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2090&amp;alphaletter=I&amp;term=Industries" title="Notes and Definitions: Industries"> Industries:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2090.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">tourism, transshipment facilities, banking</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2089&amp;alphaletter=I&amp;term=Industrial%20production%20growth%20rate" title="Notes and Definitions: Industrial production growth rate"> Industrial production growth rate:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2089.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">NA%</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2095&amp;alphaletter=L&amp;term=Labor%20force" title="Notes and Definitions: Labor force"> Labor force:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2095.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">51,610</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2095rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=192#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 192 </a> </span>
											
								<div class="category" style="padding-top: 2px;">
									
										<em>note:</em>
										
									<span class="category_data" style="font-weight:normal; vertical-align:top;">of the 51,610 workers aged 15 and over in the labor force, 32,252 were born in Aruba and 19,353 came from abroad; foreign workers are 38% of the employed population (2007 est.) </span></div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2048&amp;alphaletter=L&amp;term=Labor%20force%20-%20by%20occupation" title="Notes and Definitions: Labor force - by occupation"> Labor force - by occupation:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2048.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
							<div class="category" style="padding-top: 2px;">
								agriculture:
									
								
								<span class="category_data" style="font-weight:normal; vertical-align:bottom;">NA%</span></div>
								
								
								
								<div class="category" style="padding-top: 2px;">
									industry:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">NA% </span></div>
								
								<div class="category" style="padding-top: 2px;">
									services:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">NA% </span></div>
								
								<div class="category" style="padding-top: 2px;">
									
										<em>note:</em>
										
									<span class="category_data" style="font-weight:normal; vertical-align:top;">most employment is in wholesale and retail trade, followed by hotels and restaurants </span></div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2129&amp;alphaletter=U&amp;term=Unemployment%20rate" title="Notes and Definitions: Unemployment rate"> Unemployment rate:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2129.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">6.9% (2005 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2129rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=72#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 72 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2046&amp;alphaletter=P&amp;term=Population%20below%20poverty%20line" title="Notes and Definitions: Population below poverty line"> Population below poverty line:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2046.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">NA%</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2047&amp;alphaletter=H&amp;term=Household%20income%20or%20consumption%20by%20percentage%20share" title="Notes and Definitions: Household income or consumption by percentage share"> Household income or consumption by percentage share:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2047.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
							<div class="category" style="padding-top: 2px;">
								lowest 10%:
									
								
								<span class="category_data" style="font-weight:normal; vertical-align:bottom;">NA%</span></div>
								
								
								
								<div class="category" style="padding-top: 2px;">
									highest 10%:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">NA% </span></div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2056&amp;alphaletter=B&amp;term=Budget" title="Notes and Definitions: Budget"> Budget:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2056.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
							<div class="category" style="padding-top: 2px;">
								revenues:
									
								
								<span class="category_data" style="font-weight:normal; vertical-align:bottom;">$625.1 million</span></div>
								
								
								
								<div class="category" style="padding-top: 2px;">
									expenditures:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">$813.9 million (2013 est.) </span></div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2221&amp;alphaletter=T&amp;term=Taxes%20and%20other%20revenues" title="Notes and Definitions: Taxes and other revenues"> Taxes and other revenues:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2221.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">24.8% of GDP (2013 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2221rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=132#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 132 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2222&amp;alphaletter=B&amp;term=Budget%20surplus%20(+)%20or%20deficit%20(-)" title="Notes and Definitions: Budget surplus (+) or deficit (-)"> Budget surplus (+) or deficit (-):</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2222.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">-7.5% of GDP (2013 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2222rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=189#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 189 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2186&amp;alphaletter=P&amp;term=Public%20debt" title="Notes and Definitions: Public debt"> Public debt:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2186.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">67% of GDP (2013)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2186rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=42#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 42 </a> </span>
											
								<div class="category_data" style="padding-top: 3px;">55% of GDP (2012) </div>
							
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2080&amp;alphaletter=F&amp;term=Fiscal%20year" title="Notes and Definitions: Fiscal year"> Fiscal year:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2080.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">calendar year</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2092&amp;alphaletter=I&amp;term=Inflation%20rate%20(consumer%20prices)" title="Notes and Definitions: Inflation rate (consumer prices)"> Inflation rate (consumer prices):</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2092.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">-2% (2013 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2092rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=2#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 2 </a> </span>
											
								<div class="category_data" style="padding-top: 3px;">0.6% (2012 est.) </div>
							
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2207&amp;alphaletter=C&amp;term=Central%20bank%20discount%20rate" title="Notes and Definitions: Central bank discount rate"> Central bank discount rate:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2207.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">1% (31 December 2010 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2207rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=103#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 103 </a> </span>
											
								<div class="category_data" style="padding-top: 3px;">3% (31 December 2009 est.) </div>
							
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2208&amp;alphaletter=C&amp;term=Commercial%20bank%20prime%20lending%20rate" title="Notes and Definitions: Commercial bank prime lending rate"> Commercial bank prime lending rate:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2208.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">10.5% (31 December 2013 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2208rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=107#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 107 </a> </span>
											
								<div class="category_data" style="padding-top: 3px;">8.4% (31 December 2012 est.) </div>
							
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2214&amp;alphaletter=S&amp;term=Stock%20of%20narrow%20money" title="Notes and Definitions: Stock of narrow money"> Stock of narrow money:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2214.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">$1.022 billion (31 December 2012 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2214rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=150#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 150 </a> </span>
											
								<div class="category_data" style="padding-top: 3px;">$868.5 million (31 December 2011 est.) </div>
							
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2215&amp;alphaletter=S&amp;term=Stock%20of%20broad%20money" title="Notes and Definitions: Stock of broad money"> Stock of broad money:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2215.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">$1.91 billion (31 December 2012 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2215rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=153#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 153 </a> </span>
											
								<div class="category_data" style="padding-top: 3px;">$1.765 billion (31 December 2011 est.) </div>
							
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2211&amp;alphaletter=S&amp;term=Stock%20of%20domestic%20credit" title="Notes and Definitions: Stock of domestic credit"> Stock of domestic credit:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2211.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">$1.594 billion (31 December 2012 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2211rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=140#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 140 </a> </span>
											
								<div class="category_data" style="padding-top: 3px;">$1.448 billion (31 December 2011 est.) </div>
							
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2078&amp;alphaletter=E&amp;term=Exports" title="Notes and Definitions: Exports"> Exports:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2078.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">$2.222 billion (2013 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2078rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=142#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 142 </a> </span>
											
								<div class="category_data" style="padding-top: 3px;">$1.389 billion (2012 est.) </div>
							
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2049&amp;alphaletter=E&amp;term=Exports%20-%20commodities" title="Notes and Definitions: Exports - commodities"> Exports - commodities:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2049.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">live animals and animal products, art and collectibles, machinery and electrical equipment, transport equipment</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2050&amp;alphaletter=E&amp;term=Exports%20-%20partners" title="Notes and Definitions: Exports - partners"> Exports - partners:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2050.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">Colombia 39.4%, Venezuela 29.3%, US 13%, Netherlands Antilles 4.1% (2012)</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2087&amp;alphaletter=I&amp;term=Imports" title="Notes and Definitions: Imports"> Imports:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2087.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">$3.162 billion (2013 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2087rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=146#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 146 </a> </span>
											
								<div class="category_data" style="padding-top: 3px;">$2.039 billion (2012 est.) </div>
							
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2058&amp;alphaletter=I&amp;term=Imports%20-%20commodities" title="Notes and Definitions: Imports - commodities"> Imports - commodities:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2058.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">machinery and electrical equipment, crude oil for refining and reexport, chemicals; foodstuffs</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2061&amp;alphaletter=I&amp;term=Imports%20-%20partners" title="Notes and Definitions: Imports - partners"> Imports - partners:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2061.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">US 46.4%, Netherlands 11.5%, UK 5.4% (2012)</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2079&amp;alphaletter=D&amp;term=Debt%20-%20external" title="Notes and Definitions: Debt - external"> Debt - external:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2079.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">$533.4 million (2005 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2079rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=173#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 173 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2076&amp;alphaletter=E&amp;term=Exchange%20rates" title="Notes and Definitions: Exchange rates"> Exchange rates:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2076.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">Aruban guilders/florins per US dollar -</div>
								
								<div class="category_data" style="padding-top: 3px;">1.79 (2013 est.) </div>
							
								<div class="category_data" style="padding-top: 3px;">1.79 (2012 est.) </div>
							
				<tr>
					<td class="category_data" style="padding-bottom: 5px;"></td>
				</tr>
			</table>
		</ul>
	</div>
	
</div>

<script src="jClocksGMT-master/js/jClocksGMT.js"></script>
<script src="jClocksGMT-master/js/jquery.rotate.js"></script>
<link rel="stylesheet" href="jClocksGMT-master/css/jClocksGMT.css">



		
<script>
	
		$(document).ready(function() { 
			
				$('[id^="CollapsiblePanel1"] h2').css({'background-color':'#cce5e5',"border-bottom":"2px solid white","cursor":"pointer"}); // cam 			
			
		});
	
   </script>



<div id="CollapsiblePanel1_Energy" class="CollapsiblePanel" style="width:100%; ">
<div class="wrapper">
<h2 class="question question-back" ccode="aa" sectiontitle="Energy"><a href="javascript:void();">Energy</span> ::</span><span class="region">Aruba</span></a></h2>
<div class="answer" align="left">
	<div class="box" style="padding: 0px; margin: 0px;">
		<ul style="text-align: left;padding: 0px;margin: 0px;width: 100%;">
			<table border="0" cellspacing="0" cellpadding="0"  style="width: 100%;">
				 
				
					<tr class="cam_light" >
					
					<td width="450" height="20"><div class="category" style="padding-left:5px;" id="field"> 
										<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2232&amp;alphaletter=E&amp;term=Electricity%20-%20production" title="Notes and Definitions: Electricity - production"> Electricity - production</a>:
									 </div></td>
					
						<td align="right">
						
						
						
								<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2232.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"> <img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0"  style="text-decoration:none;"> </a>
								
						</tr>
						
					
					<tr>
					
					<td id="data" colspan="2" style="vertical-align:middle;">
					
					
					
					
						<div class="category_data">980 million kWh (2010 est.)</div>
					
									<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2232rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=147#aa147" onMouseDown="" title="Country comparison to the world" alt="Country comparison to the world">147</a> </span>
									
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2233&amp;alphaletter=E&amp;term=Electricity%20-%20consumption" title="Notes and Definitions: Electricity - consumption"> Electricity - consumption:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2233.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">911.4 million kWh (2010 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2233rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=153#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 153 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2234&amp;alphaletter=E&amp;term=Electricity%20-%20exports" title="Notes and Definitions: Electricity - exports"> Electricity - exports:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2234.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">0 kWh (2012 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2234rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=92#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 92 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2235&amp;alphaletter=E&amp;term=Electricity%20-%20imports" title="Notes and Definitions: Electricity - imports"> Electricity - imports:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2235.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">0 kWh (2012 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2235rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=110#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 110 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2236&amp;alphaletter=E&amp;term=Electricity%20-%20installed%20generating%20capacity" title="Notes and Definitions: Electricity - installed generating capacity"> Electricity - installed generating capacity:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2236.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">266,000 kW (2010 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2236rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=153#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 153 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2237&amp;alphaletter=E&amp;term=Electricity%20-%20from%20fossil%20fuels" title="Notes and Definitions: Electricity - from fossil fuels"> Electricity - from fossil fuels:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2237.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">88.7% of total installed capacity (2010 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2237rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=81#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 81 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2239&amp;alphaletter=E&amp;term=Electricity%20-%20from%20nuclear%20fuels" title="Notes and Definitions: Electricity - from nuclear fuels"> Electricity - from nuclear fuels:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2239.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">0% of total installed capacity (2010 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2239rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=31#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 31 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2238&amp;alphaletter=E&amp;term=Electricity%20-%20from%20hydroelectric%20plants" title="Notes and Definitions: Electricity - from hydroelectric plants"> Electricity - from hydroelectric plants:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2238.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">0% of total installed capacity (2010 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2238rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=151#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 151 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2240&amp;alphaletter=E&amp;term=Electricity%20-%20from%20other%20renewable%20sources" title="Notes and Definitions: Electricity - from other renewable sources"> Electricity - from other renewable sources:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2240.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">11.3% of total installed capacity (2010 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2240rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=25#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 25 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2241&amp;alphaletter=C&amp;term=Crude%20oil%20-%20production" title="Notes and Definitions: Crude oil - production"> Crude oil - production:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2241.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">2,811 bbl/day (2012 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2241rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=103#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 103 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2242&amp;alphaletter=C&amp;term=Crude%20oil%20-%20exports" title="Notes and Definitions: Crude oil - exports"> Crude oil - exports:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2242.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">0 bbl/day (2010 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2242rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=75#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 75 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2243&amp;alphaletter=C&amp;term=Crude%20oil%20-%20imports" title="Notes and Definitions: Crude oil - imports"> Crude oil - imports:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2243.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">228,800 bbl/day (2010 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2243rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=30#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 30 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2244&amp;alphaletter=C&amp;term=Crude%20oil%20-%20proved%20reserves" title="Notes and Definitions: Crude oil - proved reserves"> Crude oil - proved reserves:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2244.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">0 bbl (1 January 2013 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2244rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=101#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 101 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2245&amp;alphaletter=R&amp;term=Refined%20petroleum%20products%20-%20production" title="Notes and Definitions: Refined petroleum products - production"> Refined petroleum products - production:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2245.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">234,200 bbl/day (2010 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2245rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=50#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 50 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2246&amp;alphaletter=R&amp;term=Refined%20petroleum%20products%20-%20consumption" title="Notes and Definitions: Refined petroleum products - consumption"> Refined petroleum products - consumption:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2246.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">5,661 bbl/day (2011 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2246rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=163#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 163 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2247&amp;alphaletter=R&amp;term=Refined%20petroleum%20products%20-%20exports" title="Notes and Definitions: Refined petroleum products - exports"> Refined petroleum products - exports:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2247.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">234,200 bbl/day (2010 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2247rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=26#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 26 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2248&amp;alphaletter=R&amp;term=Refined%20petroleum%20products%20-%20imports" title="Notes and Definitions: Refined petroleum products - imports"> Refined petroleum products - imports:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2248.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">6,725 bbl/day (2010 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2248rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=137#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 137 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2249&amp;alphaletter=N&amp;term=Natural%20gas%20-%20production" title="Notes and Definitions: Natural gas - production"> Natural gas - production:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2249.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">1 cu m (2011 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2249rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=95#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 95 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2250&amp;alphaletter=N&amp;term=Natural%20gas%20-%20consumption" title="Notes and Definitions: Natural gas - consumption"> Natural gas - consumption:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2250.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">1 cu m (2010 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2250rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=115#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 115 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2251&amp;alphaletter=N&amp;term=Natural%20gas%20-%20exports" title="Notes and Definitions: Natural gas - exports"> Natural gas - exports:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2251.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">1 cu m (2011 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2251rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=53#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 53 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2252&amp;alphaletter=N&amp;term=Natural%20gas%20-%20imports" title="Notes and Definitions: Natural gas - imports"> Natural gas - imports:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2252.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">1 cu m (2011 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2252rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=78#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 78 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2253&amp;alphaletter=N&amp;term=Natural%20gas%20-%20proved%20reserves" title="Notes and Definitions: Natural gas - proved reserves"> Natural gas - proved reserves:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2253.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">0 cu m (1 January 2013 est.)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2253rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=107#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 107 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2254&amp;alphaletter=C&amp;term=Carbon%20dioxide%20emissions%20from%20consumption%20of%20energy" title="Notes and Definitions: Carbon dioxide emissions from consumption of energy"> Carbon dioxide emissions from consumption of energy:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2254.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">1.237 million Mt (2011 est.)</div>
								
				<tr>
					<td class="category_data" style="padding-bottom: 5px;"></td>
				</tr>
			</table>
		</ul>
	</div>
	
</div>

<script src="jClocksGMT-master/js/jClocksGMT.js"></script>
<script src="jClocksGMT-master/js/jquery.rotate.js"></script>
<link rel="stylesheet" href="jClocksGMT-master/css/jClocksGMT.css">



		
<script>
	
		$(document).ready(function() { 
			
				$('[id^="CollapsiblePanel1"] h2').css({'background-color':'#cce5e5',"border-bottom":"2px solid white","cursor":"pointer"}); // cam 			
			
		});
	
   </script>



<div id="CollapsiblePanel1_Comm" class="CollapsiblePanel" style="width:100%; ">
<div class="wrapper">
<h2 class="question question-back" ccode="aa" sectiontitle="Communications"><a href="javascript:void();">Communications</span> ::</span><span class="region">Aruba</span></a></h2>
<div class="answer" align="left">
	<div class="box" style="padding: 0px; margin: 0px;">
		<ul style="text-align: left;padding: 0px;margin: 0px;width: 100%;">
			<table border="0" cellspacing="0" cellpadding="0"  style="width: 100%;">
				 
				
					<tr class="cam_light" >
					
					<td width="450" height="20"><div class="category" style="padding-left:5px;" id="field"> 
										<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2150&amp;alphaletter=T&amp;term=Telephones%20-%20main%20lines%20in%20use" title="Notes and Definitions: Telephones - main lines in use"> Telephones - main lines in use</a>:
									 </div></td>
					
						<td align="right">
						
						
						
								<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2150.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"> <img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0"  style="text-decoration:none;"> </a>
								
						</tr>
						
					
					<tr>
					
					<td id="data" colspan="2" style="vertical-align:middle;">
					
					
					
					
						<div class="category_data">43,000 (2012)</div>
					
									<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2150rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=169#aa169" onMouseDown="" title="Country comparison to the world" alt="Country comparison to the world">169</a> </span>
									
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2151&amp;alphaletter=T&amp;term=Telephones%20-%20mobile%20cellular" title="Notes and Definitions: Telephones - mobile cellular"> Telephones - mobile cellular:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2151.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">135,000 (2012)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2151rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=188#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 188 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2124&amp;alphaletter=T&amp;term=Telephone%20system" title="Notes and Definitions: Telephone system"> Telephone system:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2124.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
							<div class="category" style="padding-top: 2px;">
								general assessment:
									
								
								<span class="category_data" style="font-weight:normal; vertical-align:bottom;">modern fully automatic telecommunications system</span></div>
								
								
								
								<div class="category" style="padding-top: 2px;">
									domestic:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">increased competition through privatization; 3 mobile-cellular service providers are now licensed </span></div>
								
								<div class="category" style="padding-top: 2px;">
									international:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">country code - 297; landing site for the PAN-AM submarine telecommunications cable system that extends from the US Virgin Islands through Aruba to Venezuela, Colombia, Panama, and the west coast of South America; extensive interisland microwave radio relay links (2007) </span></div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2213&amp;alphaletter=B&amp;term=Broadcast%20media" title="Notes and Definitions: Broadcast media"> Broadcast media:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2213.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">2 commercial TV stations; cable TV subscription service provides access to foreign channels; about 20 commercial radio stations broadcast (2007)</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2154&amp;alphaletter=I&amp;term=Internet%20country%20code" title="Notes and Definitions: Internet country code"> Internet country code:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2154.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">.aw</div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2184&amp;alphaletter=I&amp;term=Internet%20hosts" title="Notes and Definitions: Internet hosts"> Internet hosts:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2184.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">40,560 (2012)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2184rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=101#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 101 </a> </span>
											
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2153&amp;alphaletter=I&amp;term=Internet%20users" title="Notes and Definitions: Internet users"> Internet users:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2153.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">24,000 (2009)</div>
								
											<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2153rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=188#aa" onMouseDown=""  title="Country comparison to the world" alt="Country comparison to the world"> 188 </a> </span>
											
				<tr>
					<td class="category_data" style="padding-bottom: 5px;"></td>
				</tr>
			</table>
		</ul>
	</div>
	
</div>

<script src="jClocksGMT-master/js/jClocksGMT.js"></script>
<script src="jClocksGMT-master/js/jquery.rotate.js"></script>
<link rel="stylesheet" href="jClocksGMT-master/css/jClocksGMT.css">



		
<script>
	
		$(document).ready(function() { 
			
				$('[id^="CollapsiblePanel1"] h2').css({'background-color':'#cce5e5',"border-bottom":"2px solid white","cursor":"pointer"}); // cam 			
			
		});
	
   </script>



<div id="CollapsiblePanel1_Trans" class="CollapsiblePanel" style="width:100%; ">
<div class="wrapper">
<h2 class="question question-back" ccode="aa" sectiontitle="Transportation"><a href="javascript:void();">Transportation</span> ::</span><span class="region">Aruba</span></a></h2>
<div class="answer" align="left">
	<div class="box" style="padding: 0px; margin: 0px;">
		<ul style="text-align: left;padding: 0px;margin: 0px;width: 100%;">
			<table border="0" cellspacing="0" cellpadding="0"  style="width: 100%;">
				 
				
					<tr class="cam_light" >
					
					<td width="450" height="20"><div class="category" style="padding-left:5px;" id="field"> 
										<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2053&amp;alphaletter=A&amp;term=Airports" title="Notes and Definitions: Airports"> Airports</a>:
									 </div></td>
					
						<td align="right">
						
						
						
								<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2053.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"> <img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0"  style="text-decoration:none;"> </a>
								
						</tr>
						
					
					<tr>
					
					<td id="data" colspan="2" style="vertical-align:middle;">
					
					
					
					
						<div class="category_data">1 (2013)</div>
					
									<span class="category" style="padding-left:7px;">country comparison to the world:</span> <span class="category_data"> <a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/rankorder/2053rank.html?countryname=Aruba&amp;countrycode=aa&amp;regionCode=cam&amp;rank=210#aa210" onMouseDown="" title="Country comparison to the world" alt="Country comparison to the world">210</a> </span>
									
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2030&amp;alphaletter=A&amp;term=Airports%20-%20with%20paved%20runways" title="Notes and Definitions: Airports - with paved runways"> Airports - with paved runways:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2030.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
							<div class="category" style="padding-top: 2px;">
								total:
									
								
								<span class="category_data" style="font-weight:normal; vertical-align:bottom;">1</span></div>
								
								
								
								<div class="category" style="padding-top: 2px;">
									2,438 to 3,047 m:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">1 (2013) </span></div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2120&amp;alphaletter=P&amp;term=Ports%20and%20terminals" title="Notes and Definitions: Ports and terminals"> Ports and terminals:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2120.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
							<div class="category" style="padding-top: 2px;">
								major seaport(s):
									
								
								<span class="category_data" style="font-weight:normal; vertical-align:bottom;">Barcadera, Oranjestad</span></div>
								
								
								
								<div class="category" style="padding-top: 2px;">
									oil terminal(s):
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">Sint Nicolaas </span></div>
								
								<div class="category" style="padding-top: 2px;">
									cruise port(s):
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">Oranjestad </span></div>
								
				<tr>
					<td class="category_data" style="padding-bottom: 5px;"></td>
				</tr>
			</table>
		</ul>
	</div>
	
</div>

<script src="jClocksGMT-master/js/jClocksGMT.js"></script>
<script src="jClocksGMT-master/js/jquery.rotate.js"></script>
<link rel="stylesheet" href="jClocksGMT-master/css/jClocksGMT.css">



		
<script>
	
		$(document).ready(function() { 
			
				$('[id^="CollapsiblePanel1"] h2').css({'background-color':'#cce5e5',"border-bottom":"2px solid white","cursor":"pointer"}); // cam 			
			
		});
	
   </script>



<div id="CollapsiblePanel1_Military" class="CollapsiblePanel" style="width:100%; ">
<div class="wrapper">
<h2 class="question question-back" ccode="aa" sectiontitle="Military"><a href="javascript:void();">Military</span> ::</span><span class="region">Aruba</span></a></h2>
<div class="answer" align="left">
	<div class="box" style="padding: 0px; margin: 0px;">
		<ul style="text-align: left;padding: 0px;margin: 0px;width: 100%;">
			<table border="0" cellspacing="0" cellpadding="0"  style="width: 100%;">
				 
				
					<tr class="cam_light" >
					
					<td width="450" height="20"><div class="category" style="padding-left:5px;" id="field"> 
										<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2055&amp;alphaletter=M&amp;term=Military%20branches" title="Notes and Definitions: Military branches"> Military branches</a>:
									 </div></td>
					
						<td align="right">
						
						
						
								<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2055.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"> <img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0"  style="text-decoration:none;"> </a>
								
						</tr>
						
					
					<tr>
					
					<td id="data" colspan="2" style="vertical-align:middle;">
					
					
					
					
						<div class="category_data">no regular military forces (2011)</div>
					
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2105&amp;alphaletter=M&amp;term=Manpower%20available%20for%20military%20service" title="Notes and Definitions: Manpower available for military service"> Manpower available for military service:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2105.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
							<div class="category" style="padding-top: 2px;">
								males age 16-49:
									
								
								<span class="category_data" style="font-weight:normal; vertical-align:bottom;">24,891</span></div>
								
								
								
								<div class="category" style="padding-top: 2px;">
									females age 16-49:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">26,202 (2010 est.) </span></div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2025&amp;alphaletter=M&amp;term=Manpower%20fit%20for%20military%20service" title="Notes and Definitions: Manpower fit for military service"> Manpower fit for military service:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2025.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
							<div class="category" style="padding-top: 2px;">
								males age 16-49:
									
								
								<span class="category_data" style="font-weight:normal; vertical-align:bottom;">20,527</span></div>
								
								
								
								<div class="category" style="padding-top: 2px;">
									females age 16-49:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">21,493 (2010 est.) </span></div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2026&amp;alphaletter=M&amp;term=Manpower%20reaching%20militarily%20significant%20age%20annually" title="Notes and Definitions: Manpower reaching militarily significant age annually"> Manpower reaching militarily significant age annually:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2026.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
							<div class="category" style="padding-top: 2px;">
								male:
									
								
								<span class="category_data" style="font-weight:normal; vertical-align:bottom;">767</span></div>
								
								
								
								<div class="category" style="padding-top: 2px;">
									female:
									
									<span class="category_data" style="font-weight:normal; vertical-align:top;">743 (2010 est.) </span></div>
								
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2137&amp;alphaletter=M&amp;term=Military%20-%20note" title="Notes and Definitions: Military - note"> Military - note:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2137.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">defense is the responsibility of the Netherlands; the Aruba security services focus on organized crime and terrorism (2011)</div>
								
				<tr>
					<td class="category_data" style="padding-bottom: 5px;"></td>
				</tr>
			</table>
		</ul>
	</div>
	
</div>

<script src="jClocksGMT-master/js/jClocksGMT.js"></script>
<script src="jClocksGMT-master/js/jquery.rotate.js"></script>
<link rel="stylesheet" href="jClocksGMT-master/css/jClocksGMT.css">



		
<script>
	
		$(document).ready(function() { 
			
				$('[id^="CollapsiblePanel1"] h2').css({'background-color':'#cce5e5',"border-bottom":"2px solid white","cursor":"pointer"}); // cam 			
			
		});
	
   </script>



<div id="CollapsiblePanel1_Issues" class="CollapsiblePanel" style="width:100%; ">
<div class="wrapper">
<h2 class="question question-back" ccode="aa" sectiontitle="Transnational Issues"><a href="javascript:void();">Transnational Issues</span> ::</span><span class="region">Aruba</span></a></h2>
<div class="answer" align="left">
	<div class="box" style="padding: 0px; margin: 0px;">
		<ul style="text-align: left;padding: 0px;margin: 0px;width: 100%;">
			<table border="0" cellspacing="0" cellpadding="0"  style="width: 100%;">
				 
				
					<tr class="cam_light" >
					
					<td width="450" height="20"><div class="category" style="padding-left:5px;" id="field"> 
										<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2070&amp;alphaletter=D&amp;term=Disputes%20-%20international" title="Notes and Definitions: Disputes - international"> Disputes - international</a>:
									 </div></td>
					
						<td align="right">
						
						
						
								<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2070.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"> <img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0"  style="text-decoration:none;"> </a>
								
						</tr>
						
					
					<tr>
					
					<td id="data" colspan="2" style="vertical-align:middle;">
					
					
					
					
						<div class="category_data">none</div>
					
							</td>
							
							</tr>
							
							<tr>
								<td height="10"></td>
							</tr>
							
							<tr class="cam_light">
								<td width="450" height="20">
									
									<div class="category" style="padding-left:5px;" id="field"> 
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/docs/notesanddefs.html?fieldkey=2086&amp;alphaletter=I&amp;term=Illicit%20drugs" title="Notes and Definitions: Illicit drugs"> Illicit drugs:</a>
												 </div></td>
								
									
									<td align="right"> 
										
										
										
												<a href="/web/20140902034709/https://www.cia.gov/library/publications/the-world-factbook/fields/2086.html#aa" title="<img src = '../graphics/field_listing_tooltip.gif'>"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/graphics/field_listing_on.gif" border="0" ></a>
												</td>
								
							</tr>
							<tr height="22">
							
							<td colspan="2" id="data">
							
							
									<div class="category_data">transit point for US- and Europe-bound narcotics with some accompanying money-laundering activity; relatively high percentage of population consumes cocaine

</div>
								
				<tr>
					<td class="category_data" style="padding-bottom: 5px;"></td>
				</tr>
			</table>
		</ul>
	</div>
	
</div>

				<div class="wrapper">
					<div style="float:right; margin-top: 0px;" class="expand_all">
						<a href="javascript:void(0)" class="expand">EXPAND ALL</a><a href="javascript:void(0)" class="collapse" style="display: none;">COLLAPSE ALL</a> 
					</div>
				</div>
		
	<div id="flagDialog" style="display: none" title="The World Factbook"></div>
	<div id="photoDialogWindow" style="display:none;" title="The World Factbook"></div>
	

								<!-- InstanceEndEditable --> 
							</td>
						</tr>
						<tr>
							<td style="height:75px;">&nbsp;</td>
						</tr>
					</table>
				</div>
                 </div>
               </article>
             </div>
           </div>
         </section> 
        <footer id="footer"><span class="divider"></span>
       <a href="#" class="logo-2"><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/images/logo-2.png" alt="Central Intelligence Agency"></a>
       <div class="footer-holder">
         <div class="footer-frame">

           <nav class="footer-nav"><div class="info-block">
                   
                   <h3><a href="/web/20140902034709/https://www.cia.gov/about-cia">About CIA</a></h3>
                   <ul><li>
                           <a href="/web/20140902034709/https://www.cia.gov/about-cia/todays-cia">Today's CIA</a>
                        </li>
                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/about-cia/leadership">Leadership</a>

                        </li>
                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/about-cia/cia-vision-mission-values">CIA Vision, Mission &amp; Values</a>
                        </li>
                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/about-cia/headquarters-tour">Headquarters Tour</a>
                        </li>

                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/about-cia/cia-museum">CIA Museum</a>
                        </li>
                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/about-cia/history-of-the-cia">History of the CIA</a>
                        </li>
                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/about-cia/faqs">FAQs</a>

                        </li>
                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/about-cia/no-fear-act">NoFEAR Act</a>
                        </li>
                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/about-cia/site-policies">Site Policies</a>
                        </li>
                    </ul></div>

              <div class="info-block">
                   
                   <h3><a href="/web/20140902034709/https://www.cia.gov/careers">Careers &amp; Internships</a></h3>
                   <ul><li>
                           <a href="/web/20140902034709/https://www.cia.gov/careers/opportunities">Career Opportunities </a>
                        </li>
                        <li>

                           <a href="/web/20140902034709/https://www.cia.gov/careers/student-opportunities">Student Opportunities</a>
                        </li>
                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/careers/application-process">Application Process</a>
                        </li>
                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/careers/life-at-cia">Life at CIA</a>

                        </li>
                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/careers/diversity">Diversity</a>
                        </li>
                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/careers/military-transition">Military Transition</a>
                        </li>
                        <li>

                           <a href="/web/20140902034709/https://www.cia.gov/careers/games-information">Diversions &amp; Information</a>
                        </li>
                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/careers/faq">FAQs</a>
                        </li>
                    </ul><h3><a href="/web/20140902034709/https://www.cia.gov/offices-of-cia">Offices of CIA</a></h3>

                   <ul><li>
                           <a href="/web/20140902034709/https://www.cia.gov/offices-of-cia/intelligence-analysis">Intelligence &amp; Analysis</a>
                        </li>
                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/offices-of-cia/clandestine-service">Clandestine Service</a>
                        </li>
                        <li>

                           <a href="/web/20140902034709/https://www.cia.gov/offices-of-cia/science-technology">Science &amp; Technology</a>
                        </li>
                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/offices-of-cia/mission-support">Support to Mission</a>
                        </li>
                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/offices-of-cia/human-resources">Human Resources</a>

                        </li>
                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/offices-of-cia/public-affairs">Public Affairs</a>
                        </li>
                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/offices-of-cia/general-counsel">General Counsel</a>
                        </li>
                        <li>

                           <a href="/web/20140902034709/https://www.cia.gov/offices-of-cia/equal-employment-opportunity">Equal Employment Opportunity</a>
                        </li>
                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/offices-of-cia/congressional-affairs">Congressional Affairs</a>
                        </li>
                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/offices-of-cia/inspector-general">Inspector General</a>

                        </li>
                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/offices-of-cia/military-affairs">Military Affairs</a>
                        </li>
                    </ul></div>


             <div class="info-block">
                   
                   <h3><a href="/web/20140902034709/https://www.cia.gov/news-information">News &amp; Information</a></h3>

                   <ul><li>
                           <a href="/web/20140902034709/https://www.cia.gov/news-information/press-releases-statements">Press Releases &amp; Statements</a>
                        </li>
                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/news-information/speeches-testimony">Speeches &amp; Testimony</a>
                        </li>

                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/news-information/cia-the-war-on-terrorism">CIA &amp; the War on Terrorism</a>
                        </li>
                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/news-information/featured-story-archive">Featured Story Archive</a>
                        </li>
                        <li>

                           <a href="/web/20140902034709/https://www.cia.gov/news-information/Whats-New-on-CIAgov">What&#8217;s New Archive</a>
                        </li>
                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/news-information/your-news">Your News</a>
                        </li>
                    </ul><h3><a href="/web/20140902034709/https://www.cia.gov/library">Library</a></h3>
                   <ul><li>

                           <a href="/web/20140902034709/https://www.cia.gov/library/publications">Publications</a>
                        </li>
                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/library/center-for-the-study-of-intelligence">Center for the Study of Intelligence</a>
                        </li>
                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/library/foia">Freedom of Information Act Electronic Reading Room</a>

                        </li>
                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/library/kent-center-occasional-papers">Kent Center Occasional Papers</a>
                        </li>
                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/library/intelligence-literature">Intelligence Literature</a>
                        </li>
                        <li>

                           <a href="/web/20140902034709/https://www.cia.gov/library/reports">Reports</a>
                        </li>
                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/library/related-links.html">Related Links</a>
                        </li>
                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/library/video-center">Video Center</a>

                        </li>
                    </ul></div>


               <div class="info-block add">
                   
                   <h3><a href="/web/20140902034709/https://www.cia.gov/kids-page">Kids' Zone</a></h3>
                   <ul><li>
                           <a href="/web/20140902034709/https://www.cia.gov/kids-page/k-5th-grade">K-5th Grade</a>
                        </li>

                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/kids-page/6-12th-grade">6-12th Grade</a>
                        </li>
                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/kids-page/parents-teachers">Parents &amp; Teachers</a>
                        </li>
                        <li>

                           <a href="/web/20140902034709/https://www.cia.gov/kids-page/games">Games</a>
                        </li>
                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/kids-page/related-links">Related Links</a>
                        </li>
                        <li>
                           <a href="/web/20140902034709/https://www.cia.gov/kids-page/privacy-statement">Privacy Statement</a>

                        </li>
                    </ul><h3><a href="/web/20140902034709/https://www.cia.gov/contact-cia">Connect with CIA</a></h3>
                   <ul class="socials-list"><li><a href="/web/20140902034709/http://www.youtube.com/user/ciagov">CIA YouTube</a></li>
                     <li><a class="social-2" href="/web/20140902034709/http://www.flickr.com/photos/ciagov">CIA Flickr PhotoStream</a></li>
                     <li><a class="social-3" href="/web/20140902034709/https://www.cia.gov/news-information/your-news">RSS</a></li>
                     <li><a class="social-4" href="/web/20140902034709/https://www.cia.gov/contact-cia">Contact Us</a></li>

                   </ul></div>


               </nav><div id="plugins" class="info-panel">
                    <h4>* Required plugins</h4>
                    <ul><li data-plugin="swf"><a href="/web/20140902034709/http://get.adobe.com/flashplayer/">Adobe&#174; Flash Player</a></li>
                        <li data-plugin="pdf"><a href="/web/20140902034709/http://get.adobe.com/reader/">Adobe&#174; Reader&#174;</a></li>

                        <li data-plugin="doc"><a href="/web/20140902034709/http://www.microsoft.com/en-us/download/details.aspx?id=4">MS Word Viewer</a></li>
                    </ul></div>
         </div>
       </div>
     </footer>
    </div>
<div class="footer-panel" style="width: 990px;" align="center">
      <nav class="sub-nav" style="width: 100%; text-align: center;" >
        <h3 class="visuallyhidden">Footer Navigation</h3>
        <ul>
			<li><a href="/web/20140902034709/https://www.cia.gov/about-cia/site-policies/#privacy-notice" title="Site Policies">Privacy</a></li>
			<li><a href="/web/20140902034709/https://www.cia.gov/about-cia/site-policies/#copy" title="Site Policies">Copyright</a></li>
			<li><a href="/web/20140902034709/https://www.cia.gov/about-cia/site-policies/" title="Site Policies">Site Policies</a></li>
			<li><a href="/web/20140902034709/http://www.usa.gov/">USA.gov</a></li>
			<li><a href="/web/20140902034709/http://www.foia.cia.gov/">FOIA</a></li>
			<li><a href="/web/20140902034709/http://www.dni.gov/">DNI.gov</a></li>
			<li><a href="/web/20140902034709/https://www.cia.gov/about-cia/no-fear-act/" title="No FEAR Act">NoFEAR Act</a></li>
			<li><a href="/web/20140902034709/https://www.cia.gov/offices-of-cia/inspector-general/">Inspector General</a></li>
			<li><a href="/web/20140902034709/https://www.cia.gov/mobile/">Mobile Site</a></li>
			<li><a href="/web/20140902034709/https://www.cia.gov/contact-cia/">Contact CIA</a></li>
			<li><a href="/web/20140902034709/https://www.cia.gov/sitemap.html">Site Map</a></li>
		</ul>
       <div style="width: 100%;" align="center"><a href="/web/20140902034709/https://www.cia.gov/open/" ><img src="/web/20140902034709im_/https://www.cia.gov/library/publications/the-world-factbook/images/ico-06.png" width="101" height="24" alt="open gov"></a></div>
      </nav>
</div>
<a href="#" class="go-top">GO TOP</a>
  <script>
	$(document).ready(function() {
		// Show or hide the sticky footer button
		$(window).scroll(function() {
			if ($(this).scrollTop() > 350) {
				$('.go-top').fadeIn(100);
			} else {
				$('.go-top').fadeOut(100);
			}
		});
		
		// Animate the scroll to top
		$('.go-top').click(function(event) {
			event.preventDefault();
			
			$('html, body').animate({scrollTop: 350}, 300);
		})
	});
</script>	

</body>
<!-- InstanceEnd --></html> 





<!--
     FILE ARCHIVED ON 3:47:09 Sep 2, 2014 AND RETRIEVED FROM THE
     INTERNET ARCHIVE ON 19:05:04 Jan 23, 2017.
     JAVASCRIPT APPENDED BY WAYBACK MACHINE, COPYRIGHT INTERNET ARCHIVE.

     ALL OTHER CONTENT MAY ALSO BE PROTECTED BY COPYRIGHT (17 U.S.C.
     SECTION 108(a)(3)).
-->

`

const asHtml20170320 = `

<!DOCTYPE html>
<!--[if lt IE 7]> <html class="no-js lt-ie9 lt-ie8 lt-ie7" lang="en"> <![endif]--><!--[if IE 7]>    <html class="no-js lt-ie9 lt-ie8" lang="en"> <![endif]--><!--[if IE 8]>    <html class="no-js lt-ie9" lang="en"> <![endif]--><!--[if gt IE 8]><!--><html xmlns="http://www.w3.org/1999/xhtml" class="no-js" lang="en" xml:lang="en"><!--<![endif]--><head><meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
<!-- <link rel="stylesheet" type="text/css" href="../css/fullscreen-external.css" media="screen" /> -->
<link rel="stylesheet" type="text/css" href="../css/publications.css" />
<link rel="stylesheet" type="text/css" href="../css/publications-detail.css" />
<meta charset="utf-8" /><meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />
<title>The World Factbook — Central Intelligence Agency</title>
<meta name="description" content="" />
<meta name="viewport" content="width=device-width" />
<meta name="LastModified" content="JAN 12, 2017" />
<link rel="stylesheet" href="../css/jobcart.css" />
<link rel="stylesheet" href="../css/smallscreen.css" />
<link rel="stylesheet" href="../css/modules.css" />
<!--[if lt IE 9]><link href="../css/fullscreen.css" rel="stylesheet" type="text/css"><![endif]-->
<link href="../css/fullscreen.css" rel="stylesheet" type="text/css" media="only screen and (min-width: 993px)" />
<script src="../js/modernizr.custom.73407.js"></script>
<script type="text/javascript" src="../js/jquery-1.8.3.min.js"></script>
<script type="text/javascript" src="../js/response.min.js"></script>
<script type="text/javascript" src="../js/rwd.js"></script>
<script type="text/javascript" src="../js/connect.js"></script>
<!-- css files for popup --><script>var isIElt9 = false;</script>
<link href="../css/bootstrap_customized.css" rel="stylesheet" />
<link rel="stylesheet" type="text/css" href="../styles/wfb_styles.css" media="screen" />
<link rel="stylesheet" type="text/css" href="../css/expandcollapse.css" media="screen" />

<script type="text/javascript" src="../js/jquery.min.js"></script>
<script type="text/javascript" src="../js/bootstrap.min.js"></script>

<script type="text/javascript" src="../js/jquery.main.js"></script>

<!--[if lt IE 9]>

    <script async src="../js/respond.min.js"></script>

    <script async src="../js/html5shiv.min.js"></script>

    <script>var isIElt9 = true;</script>

<![endif]--><!--   jquery code added for popup --><script async="" src="../js/bootstrap_customized.min.js"></script><!-- our public key --><script async="" src="/js2/pubkey.js"></script><!-- image preloading --><script>
  prePic = new Image();
  prePic.src ="images/separator-2.gif";
  prePic2 = new Image();
  prePic2.src="images/arrow-5.gif";
  prePic3 = new Image();
  prePic3.src="images/bg-heading-panel.jpg";
  prePic4 = new Image();
  prePic4.src="images/heading-banners/CHI12052_Headers_contact.jpg";
 </script><script type="text/javascript" src="../js/jquery.ba-throttle-debounce.min.js"></script>
        <script type="text/javascript" src="../js/expandcollapse.js"></script>
    </head><body class="template-document_view_publications_detail portaltype-agencypage site-CIAgov section-library icons-on"><div><noscript>Javascript must be enabled for the correct page display</noscript><div id="wrapper">



 <!-- Modal for Landing Page-->
 <div aria-hidden="true" aria-labelledby="myModalLabel" class="modal fade" id="popupLandingModal" role="dialog" style="z-index:2147483647;overflow:auto" tabindex="-1">
  <div class="modal-dialog">
   <div class="modal-content">
    <div class="modal-body">
      <div class="section-contact-cia main-block-popup">
      <button aria-hidden="true" aria-label="Close" class="cfclose" data-dismiss="modal" type="button">&#215;</button><br><section id="mainLanding"><div class="section-contact-cia heading-panel" style="width:93%">
        <!--  <h1>Contact CIA</h1> -->
        </div>
        <div class="main-holder threecol" style="width:100%">
         <div>
          
	    <ul class="breadcrumbs"><li><a href="/">Home</a></li><li><a href="/library">Library</a></li><li><a href="/library/publications">Publications</a></li><li><li>The World Factbook</li>
	    </ul><article class="description-box"><a id="main-content" tabindex="-1">&#160;</a>
           <div>
            <div class="">
             <div id="content"><h1>Contact CIA</h1>
              <div>
               <div class="" id="parent-fieldname-text-55a5766e45f049b42fce0fdbf970b9ae">
                <p>The Office of Public Affairs (OPA) is the single point of contact for all inquiries about the Central Intelligence Agency (CIA).</p>
                <p>We read every letter, fax, or e-mail we receive, and we will convey your comments to CIA officials outside OPA as appropriate. However, with limited staff and resources, we simply cannot respond to all who write to us.</p>
                <hr class="centerline" styple="margin-top:2px;margin-bottom:2px;"><h3>Contact Information</h3>
                <p><a href="#" class="contactTriggerFromLanding btn btn-info btn-block" role="button">Submit questions or comments online</a>

                </p><p><b>By postal mail:</b><br> Central Intelligence Agency<br> Office of Public Affairs<br> Washington, D.C. 20505</p>
                <p><b>By phone:<br></b>(703) 482-0623<br>Open during normal business hours.</p>
                <p><b>By fax:<br></b>(571) 204-3800<i><br>(please include a phone number where we may call you)</i></p>
                <p><br>Contact the <a href="/offices-of-cia/inspector-general/index.html" class="internal-link">Office of Inspector General</a></p>
                <p>Contact the <a href="/contact-cia/employment-verification-office.html" class="internal-link">Employment Verification Office</a></p>
                <p>&#160;</p>
                <hr class="centerline" styple="margin-top:2px;margin-bottom:2px;"><p><b>Before contacting us:</b></p>
                <ul><li>Please check our <a href="/sitemap.html" title="Site Map"><b>site map</b></a>, <b>search</b> feature, or <b>our site navigation on the left </b>to locate the information you seek. We do not routinely respond to questions for which answers are found within this Web site.<br><br></li>
                <li><a href="/careers/index.html" title="Careers"><b>Employment</b></a><b>: </b>We do not routinely answer questions about employment beyond the information on this Web site, and we do not routinely answer inquiries about the status of job applications. Recruiting will contact applicants within 45 days if their qualifications meet our needs. <br><br>Because of safety concerns for the prospective applicant, as well as security and communication issues, the CIA Recruitment Center does not accept resumes, nor can we return phone calls, e-mails or other forms of communication, from US citizens living outside of the US. When you return permanently to the US (not on vacation or leave), please visit the <a href="/careers/opportunities/index.html" title="Careers at CIA">CIA Careers page</a> and apply online for the position of interest.<br><br>To verify an employee's employment, please contact the <a href="/contact-cia/employment-verification-office.html" class="internal-link">Employment Verification Office</a>.<br><br></li>
                <li><b>Solicitations to transfer large sums of money to your bank account:</b> If you receive a solicitation to transfer a large amount of money from an African nation to your bank account in exchange for a payment of millions of dollars, go to the US Secret Service Web site for information about the Nigerian Advance Fee Fraud or "4-1-9" Fraud scheme.<br><br></li>
                <li>If you have information which you believe might be of interest to the CIA in pursuit of the CIA's foreign intelligence mission, you may use our <a href="#" class="contactTriggerFromLanding"><b>e-mail form</b></a>. We will carefully protect all information you provide, including your identity. The CIA, as a foreign intelligence agency, does not engage in US domestic law enforcement.<br><br></li>
                <li>If you have information relating to Iraq which you believe might be of interest to the US Government, please contact us through the <a href="/about-cia/iraqi-rewards-program.html" title="Iraqi Rewards Program"><b>Iraqi Rewards Program</b></a> &#8212; <a href="/about-cia/iraqi-rewards-program.html" title="Iraqi Rewards Program"><img src="/contact-cia/secondary_page_link.gif/image.gif" alt="&#1576;&#1585;&#1606;&#1575;&#1605;&#1580; &#1605;&#1603;&#1575;&#1601;&#1570;&#1578; &#1575;&#1604;&#1593;&#1585;&#1575;&#1602;" class="image-inline" title="secondary_page_link.gif"></a></li>
                </ul></div>
              </div>
             </div>
            </div>
           </div>
          </article></div>
        </div>
       </section></div>
    </div>
   </div>
  </div>
 </div>
 <!-- Modal for Contact Page-->
 <div aria-hidden="true" aria-labelledby="myModalLabel" class="modal fade" id="popupContactModal" role="dialog" style="z-index:2147483647;overflow:auto" tabindex="-1">
  <div class="modal-dialog">
   <div class="modal-content">
    <div class="modal-body">
      <div class="section-contact-cia main-block-popup">
        <button aria-hidden="true" aria-label="Close" class="cfclose" data-dismiss="modal" type="button">&#215;</button><br><div id="mainContact">
        <div class="section-contact-cia heading-panel" style="width:93%">
         <h1>Library</h1>
        </div>
        <div class="main-holder threecol" style="width:100%">
         <div id="contentContact">
          
	    <ul class="breadcrumbs"><li><a href="/">Home</a></li><li><a href="/library">Library</a></li><li><a href="/library/publications">Publications</a></li><li><a href="/library/publications/resources">Resources</a></li><li>The World Factbook</li>
	    </ul><div class="description-box">
           <a id="main-content" tabindex="-1">
           </a>
           <div class="text-holderPopup" style="width:100%">
            <div id="viewlet-above-content">
            </div>
            <div class="">
             <div id="contentContact">
              <h1>
               <div id="popupBanner" style="margin-left:5%;margin-right:5%;color:#5f1d1d;font-size:0.6em">Contact Us Form</div>
              </h1>
               <div>
                 <div class="" id="parent-fieldname-text-b7a1aa432b8946fc8bc268300efe896d">
                  <div id="popupBodyFull"></div>
                   <!-- popup form here -->
                   <div id="popupBody">
                    <meta content="telephone=no" name="format-detection"><form accept-charset="utf-8" id="cf" name="cf" role="form" class="form-horizontal">

                     <div class="form-group" style="margin-left:6%;width:95%;margin-bottom:7px;">
                       <label for="message" id="mlabel" type="label">Message:<span style="color:red">*</span></label>
                       <textarea autocapitalize="off" autocomplete="off" autocorrect="off" class="form-control" cols="70" id="message" name="message" rows="10" style="resize:none;width:90%" required="true" placeholder="Message Text"></textarea><div id="messageErr"></div>
                     </div>

                     <div class="form-group" style="margin-left:6%;width:95%;margin-bottom:7px;">
                       <label for="email" id="elabel">Email:<span style="color:red">*</span></label>
                       <input autocapitalize="off" autocomplete="off" autocorrect="off" class="form-control" id="email" maxlength="40" name="email" size="40" type="email" placeholder="user@email.com" pattern="[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z0-9-]+)*$" required="" style="width:90%"></div>
                     <div class="form-group" style="margin-left:6%;width:95%;margin-bottom:7px;">
                       <label for="name">Name: </label>
                       <input autocapitalize="off" autocomplete="off" autocorrect="off" class="form-control" id="name" maxlength="40" name="sender" size="40" type="text" placeholder="Firstname Lastname" style="width:90%"></div>

                     <div class="form-inline" style="margin-left:6%;margin-bottom:7px;">
                       <label>Phone Number:</label><br><input autocomplete="off" class="form-control" id="phone_cc" maxlength="5" name="ph_cc" placeholder="Country Code" size="12" type="cc" style="max-width:55%"><input autocomplete="off" class="form-control" id="phone" maxlength="20" name="phone_num" placeholder="Number" size="40" type="phone" style="max-width:90%"></div>

                     <div class="form-inline" style="margin-left:6%;margin-bottom:7px;">
                       <label>Mobile/Cell:</label><br><input autocomplete="off" class="form-control" id="cellcc" maxlength="5" name="mobilecc" placeholder="Country Code" size="12" type="cc" style="max-width:55%"><input autocomplete="off" class="form-control" id="cellnumber" maxlength="20" name="mobilenum" placeholder="Number" size="40" type="phone" style="max-width:90%"></div>

                    <br><p align="center">
                      </p><div style="margin-left:5%" id="refId"></div>
                     

                      <br><span id="submitFormButton">
                        <input onclick="javascript:confirm_submission();" class="btn-lg btn-block btn-primary center-block" type="button" value="Send"></span>
                        <br><input name="Reset" class="btn-lg btn-block center-block" type="reset" value="Reset"><p align="left"><span style="color:red;margin-left:5%">*</span> = required</p>
                    </form>
                   </div>
                </div>
                </div>
               </div>
              </div>
             </div>
           </div>
         </div>
        </div>
       </div>
      </div>
     </div>
    </div>
   </div>
  </div>
<!-- Modal for Report Page-->
 <div aria-hidden="true" aria-labelledby="myModalLabel" class="modal fade" id="popupReportModal" role="dialog" style="z-index:2147483647;overflow:auto" tabindex="-1">
  <div class="modal-dialog">
   <div class="modal-content">
    <div class="modal-body">
      <div class="section-contact-cia main-block-popup">
        <button aria-hidden="true" aria-label="Close" class="cfclose" data-dismiss="modal" type="button">&#215;</button><br><section id="mainReport"><div class="section-contact-cia heading-panel" style="width:93%">
         <h1>Library</h1>
        </div>
        <div class="main-holder threecol" style="width:100%">
         <div>
          
	    <ul class="breadcrumbs"><li><a href="/">Home</a></li><li><a href="/library">Library</a></li><li><a href="/library/publications">Publications</a></li><li><a href="/library/publications/resources">Resources</a></li><li>The World Factbook</li>
	    </ul><article class="description-box"><a id="main-content" tabindex="-1">&#160;</a>
           <div>
            <div class="">
             <div>
              <h1>Report Threats</h1>
              <div>
        <div class="" id="parent-fieldname-text-84043b8c4f22633413a10d8a8009e18c">
            <p>The United States and its partners continue to face a 
growing number of global threats and challenges. The CIA&#8217;s mission 
includes collecting and analyzing information about high priority 
national security issues such as international terrorism, the 
proliferation of weapons of mass destruction, cyber attacks, 
international organized crime and narcotics trafficking, regional 
conflicts, counterintelligence threats, and the effects of environmental
 and natural disasters.</p>
<p>These challenges are international in scope and are priorities for 
the Central Intelligence Agency. If you have information about these or 
other national security challenges, please provide it through our secure
 online form.  The information you provide will be protected and 
confidential. The CIA is particularly interested in information about 
imminent or planned terrorist attacks.  In cases where an imminent 
threat exists, immediately contact your local law enforcement agencies 
and provide them with the threat information.</p>
<p align="center"><a href="#" class="contactTriggerFromReport">To contact the Central Intelligence Agency click here.</a></p>
            
        </div>
              </div>
             </div>
            </div>
           </div>
          </article></div>
        </div>
       </section></div>
    </div>
    <div>
    </div>
   </div>
  </div>
 </div>

    <header id="header"><div class="header-holder">
       	<a class="skip" accesskey="S" href="#main-content">skip to content</a>
        <span class="bg-globe"></span>
        <div class="header-panel">
          <hgroup><h1 class="logo"><a href="/"><img src="/++theme++contextual.agencytheme/images/logo.png" alt="Central Intelligence Agency"><span>Central Intelligence Agency</span></a></h1>
              <h2 class="work-text">The Work Of A Nation. The Center of Intelligence.</h2>
          </hgroup><div class="search-form">
            <div class="row">
              <div class="add-nav">
                <ul><li><a title="Report Threats" class="active reportTrigger" href="/contact-cia/report-threats.html">Report Threats</a></li>
                  <li><a title="Iraqi Rewards Program" href="/about-cia/iraqi-rewards-program.html">&#1585;&#1593;&#1585;&#1576;&#1610;&#1593;&#1585;&#1576;&#1610;</a></li>
                  <li><a title="A single point of contact for all CIA inquiries." class="landingTrigger" href="/contact-cia">Contact</a></li>
                </ul></div>
            </div>
            <div class="row">
              <form id="ciaSearchForm" method="get" action="/search">
                <fieldset><legend class="visuallyhidden">Search CIA.gov</legend>
                  <label class="visuallyhidden" for="q">Search</label>
                  <input name="q" type="text" class="text" id="q" maxlength="2047" placeholder="Search CIA.gov..."><input type="hidden" name="site" value="CIA"><input type="hidden" name="output" value="xml_no_dtd"><input type="hidden" name="client" value="CIA"><input type="hidden" name="myAction" value="/search"><input type="hidden" name="proxystylesheet" value="CIA"><input type="hidden" name="submitMethod" value="get"><input type="submit" value="Search" class="submit"></fieldset></form>
            </div>
            <div class="row">              
              <ul class="lang-list"><li lang="ar" xml:lang="ar"><a title="Arabic" href="/ar">&#1593;&#1585;&#1576;&#1610;</a></li>
                <li lang="zh-cn" xml:lang="zh-cn"><a title="Chinese" href="/zh">&#20013;&#25991;</a></li>
                <li lang="en" xml:lang="en"><a title="English" href="/">English</a></li>
                <li lang="fr" xml:lang="fr"><a title="French" href="/fr">Fran&#231;ais</a></li>
                <li lang="ru" xml:lang="ru"><a title="Russian" href="/ru">&#1056;&#1091;&#1089;&#1089;&#1082;&#1080;&#1081;</a></li>
                <li lang="es" xml:lang="es"><a title="Spanish" href="/es">Espa&#241;ol</a></li>
                <li><a title="More Languages" alt="More Languages" class="more" href="/foreign-languages">More<span class="visuallyhidden"> Languages</span></a></li>
             </ul></div>
          </div>   
        </div>
        <nav id="nav"><h3 class="visuallyhidden">Navigation</h3>
          <ul><li class=""><a class="" href="/" title="">
        
        <span>Home</span>
    </a></li>


<li class=""><a class="hasChildrens" href="/about-cia" title="About CIA Description">
        
        <span>About CIA</span>
    </a><div class="drop"><ul class="globalSectionsLevel1 show-static"><li class="plain ">

    

    <a class="" href="/about-cia/todays-cia" title="">
        
        <span>Today's CIA</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/about-cia/leadership" title="">
        
        <span>Leadership</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/about-cia/cia-vision-mission-values" title="">
        
        <span>CIA Vision, Mission, Ethos &amp; Challenges</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/about-cia/headquarters-tour" title="">
        
        <span>Headquarters Tour</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/about-cia/cia-museum" title="">
        
        <span>CIA Museum</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/about-cia/history-of-the-cia" title="">
        
        <span>History of the CIA</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/about-cia/publications-review-board" title="Publications Review Board">
        
        <span>Publications Review Board</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/about-cia/accessibility" title="Accessibility ">
        
        <span>Accessibility</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/about-cia/faqs" title="">
        
        <span>FAQs</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/about-cia/no-fear-act" title="Note - link to the No Fear Act parent page, which is housed under EEO">
        
        <span>NoFEAR Act</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/about-cia/site-policies" title="">
        
        <span>Site Policies</span>
    </a>
    
    
        
        
        
    
    

</li>




        </ul><div class="static"><img class="image-inline" src="/images/section-banners/about-menu/image.jpg" title="About Menu" alt="About Menu"><h4><a href="/about-cia/">About CIA</a></h4>

<p>Discover the CIA <a href="/about-cia/">history, mission, vision and values</a>.</p></div><div class="visualClear"><!-- --></div></div></li>


<li class=""><a class="hasChildrens" href="/careers" title="">
        
        <span>Careers &amp; Internships</span>
    </a><div class="drop"><ul class="globalSectionsLevel1 show-static"><li class="plain ">

    

    <a class="" href="/careers/opportunities" title="This is an overview of all career opportunities at the CIA. ">
        
        <span>Career Opportunities </span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/careers/student-opportunities" title="This is the student profile page for candidates looking for jobs/ job listings at the CIA. Student Opportunities - Student Profiles">
        
        <span>Student Opportunities</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/careers/application-process" title="How to apply to the CIA.">
        
        <span>Application Process</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/careers/life-at-cia" title="This is the about CIA section of the Careers Site">
        
        <span>Life at CIA</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/careers/benefits.html" title="Every career at the CIA is a rewarding one. In addition to serving your country with the certainty that your work makes a difference, the CIA offers a comprehensive benefits package to reflect the dedication and contributions of our employees.">
        
        <span>Benefits</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/careers/diversity" title="This is the diversity information for the Careers Site">
        
        <span>Diversity</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/careers/military-transition" title="Your prior military service could qualify you to continue to serve your nation at the Central Intelligence Agency. Opportunities for qualified applicants are available in the U.S. and abroad.">
        
        <span>Military Transition</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/careers/games-information" title="">
        
        <span>Tools and Challenges</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/careers/faq" title="Frequently Asked Questions/ FAQ for a Career at the CIA in the Careers Section">
        
        <span>FAQs</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/careers/video-center" title="Repository of CIA videos">
        
        <span>Video Center</span>
    </a>
    
    
        
        
        
    
    

</li>




        </ul><div class="static"><img class="image-inline" src="/images/section-banners/careers-menu/image.jpg" title="Careers Menu" alt="Careers Menu"><h4><a href="/careers/">Careers &amp; Internships</a></h4>

<p>Your talent. Your diverse skills. Our mission. Learn more about <a href="/careers/">Career Opportunities at CIA</a>.</p></div><div class="visualClear"><!-- --></div></div></li>


<li class=""><a class="hasChildrens" href="/offices-of-cia" title="">
        
        <span>Offices of CIA</span>
    </a><div class="drop"><ul class="globalSectionsLevel1 show-static"><li class="plain ">

    

    <a class="" href="/offices-of-cia/intelligence-analysis" title="">
        
        <span>Intelligence &amp; Analysis</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/offices-of-cia/clandestine-service" title="">
        
        <span>Clandestine Service</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/offices-of-cia/science-technology" title="">
        
        <span>Science &amp; Technology</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/offices-of-cia/mission-support" title="">
        
        <span>Support to Mission</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/offices-of-cia/digital-innovation" title="">
        
        <span>Digital Innovation</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/offices-of-cia/mission-centers" title="">
        
        <span>Mission Centers</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/offices-of-cia/human-resources" title="">
        
        <span>Human Resources</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/offices-of-cia/public-affairs" title="Public Affairs">
        
        <span>Public Affairs</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/offices-of-cia/general-counsel" title="">
        
        <span>General Counsel</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/offices-of-cia/equal-employment-opportunity" title="">
        
        <span>Equal Employment Opportunity</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/offices-of-cia/congressional-affairs" title="Office of Congressional Affairs">
        
        <span>Congressional Affairs</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/offices-of-cia/inspector-general" title="Inspector General">
        
        <span>Inspector General</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/offices-of-cia/military-affairs" title="Military Affairs">
        
        <span>Military Affairs</span>
    </a>
    
    
        
        
        
    
    

</li>




        </ul><div class="static"><img class="image-inline" src="/images/section-banners/offices-menu/image.jpg" title="Offices Menu" alt="Offices Menu"><h4><a href="/offices-of-cia/">Offices of CIA</a></h4>

<p><a href="/offices-of-cia/">Learn how the CIA is organized</a> into directorates and key offices, responsible for securing our nation.</p></div><div class="visualClear"><!-- --></div></div></li>


<li class=""><a class="hasChildrens" href="/news-information" title="News &amp; Information Description">
        
        <span>News &amp; Information</span>
    </a><div class="drop"><ul class="globalSectionsLevel1 show-static"><li class="plain ">

    

    <a class="" href="/news-information/blog" title="">
        
        <span>Blog</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/news-information/press-releases-statements" title="">
        
        <span>Press Releases &amp; Statements</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/news-information/speeches-testimony" title="">
        
        <span>Speeches &amp; Testimony</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/news-information/cia-the-war-on-terrorism" title="">
        
        <span>CIA &amp; the War on Terrorism</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/news-information/featured-story-archive" title="index for featured story">
        
        <span>Featured Story Archive</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/news-information/your-news" title="">
        
        <span>Your News</span>
    </a>
    
    
        
        
        
    
    

</li>




        </ul><div class="static"><img class="image-inline" src="/images/section-banners/news-menu/image.jpg" title="News Menu" alt="News Menu"><h4><a href="/news-information/">News &amp; Information</a></h4>

<p>The most up-to-date CIA <a href="/news-information/">news, press releases, information and more</a>.</p></div><div class="visualClear"><!-- --></div></div></li>


<li class="
                   active
               "><a class="hasChildrens" href="/library" title="">
        
        <span>Library</span>
    </a><span class="arrow"></span><div class="drop right"><ul class="globalSectionsLevel1 show-static"><li class="selected ">

    

    <a class="" href="/library/publications" title="">
        
        <span>Publications</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/library/center-for-the-study-of-intelligence" title="CSI section">
        
        <span>Center for the Study of Intelligence</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/library/foia" title="">
        
        <span>Freedom of Information Act Electronic Reading Room</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/library/kent-center-occasional-papers" title="">
        
        <span>Kent Center Occasional Papers</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/library/intelligence-literature" title="">
        
        <span>Intelligence Literature</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/library/reports" title="Reports">
        
        <span>Reports</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/library/related-links.html" title="Related Links">
        
        <span>Related Links</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/library/video-center" title="Repository of CIA videos">
        
        <span>Video Center</span>
    </a>
    
    
        
        
        
    
    

</li>




        </ul><div class="static"><img class="image-inline" src="/images/section-banners/library-menu/image.jpg" title="Library Menu" alt="Library Menu"><h4><a href="/library/">Library</a></h4>

<p>Our <a href="/library/">open-source library</a> houses the thousands of documents, periodicals, maps and reports released to the public.</p></div><div class="visualClear"><!-- --></div></div></li>


<li class=""><a class="hasChildrens" href="/kids-page" title="Kids' Page Description">
        
        <span>Kids' Zone</span>
    </a><div class="drop right"><ul class="globalSectionsLevel1 show-static"><li class="plain ">

    

    <a class="" href="/kids-page/k-5th-grade" title="K-5th Grade">
        
        <span>K-5th Grade</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/kids-page/6-12th-grade" title="">
        
        <span>6-12th Grade</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/kids-page/parents-teachers" title="">
        
        <span>Parents &amp; Teachers</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/kids-page/games" title="">
        
        <span>Games</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/kids-page/related-links" title="">
        
        <span>Related Links</span>
    </a>
    
    
        
        
        
    
    

</li>


<li class="plain ">

    

    <a class="" href="/kids-page/privacy-statement" title="">
        
        <span>Privacy Statement</span>
    </a>
    
    
        
        
        
    
    

</li>




        </ul><div class="static"><p><img class="image-inline" src="/images/section-banners/kids-menu/image.jpg" title="Kids Menu" alt="Kids Menu"></p>

<h4><a href="/kids-page/">Kids' Zone</a></h4>

<p><a href="/kids-page/">Learn more about the Agency</a> and find some top secret things you won't see anywhere else.</p></div><div class="visualClear"><!-- --></div></div></li></ul></nav></div>
    </header><div class="main-block">
      <div id="wfb-main">
     	<section id="main"><div class="heading-panel">
             <h1>Library</h1>
           </div>           
           <div class="main-holder threecol">
             <div id="sidebar">
               <nav class="sidebar-nav"><h2 class="visuallyhidden">Secondary Navigation</h2>
                 <ul><li class="navTreeItem navTreeTopNode">
                
                   <a class="contenttype-agencyfolder" href="/library" title="">
                   
                   
                   Library
                   </a>
                
            </li>
            



<li class="navTreeItem visualNoMarker navTreeItemInPath navTreeFolderish section-publications">

    


        <a class="state-published navTreeItemInPath navTreeFolderish contenttype-agencyfolder" href="/library/publications" title="">
            Publications
        </a>

        

    
</li>


<li class="navTreeItem visualNoMarker navTreeFolderish section-center-for-the-study-of-intelligence">

    


        <a class="state-published navTreeFolderish contenttype-agencyfolder" href="/library/center-for-the-study-of-intelligence" title="CSI section">
            Center for the Study of Intelligence
        </a>

        

    
</li>


<li class="navTreeItem visualNoMarker navTreeFolderish section-foia">

    


        <a class="state-published navTreeFolderish contenttype-agencyfolder" href="/library/foia" title="">
            Freedom of Information Act Electronic Reading Room
        </a>

        

    
</li>


<li class="navTreeItem visualNoMarker navTreeFolderish section-kent-center-occasional-papers">

    


        <a class="state-published navTreeFolderish contenttype-agencyfolder" href="/library/kent-center-occasional-papers" title="">
            Kent Center Occasional Papers
        </a>

        

    
</li>


<li class="navTreeItem visualNoMarker navTreeFolderish section-intelligence-literature">

    


        <a class="state-published navTreeFolderish contenttype-agencyfolder" href="/library/intelligence-literature" title="">
            Intelligence Literature
        </a>

        

    
</li>


<li class="navTreeItem visualNoMarker navTreeFolderish section-reports">

    


        <a class="state-published navTreeFolderish contenttype-agencyfolder" href="/library/reports" title="Reports">
            Reports
        </a>

        

    
</li>


<li class="navTreeItem visualNoMarker section-related-links">

    


        <a class="state-published contenttype-agencypage" href="/library/related-links.html" title="Related Links">
            Related Links
        </a>

        

    
</li>


<li class="navTreeItem visualNoMarker navTreeFolderish section-video-center">

    


        <a class="state-published navTreeFolderish contenttype-agencyfolder" href="/library/video-center" title="Repository of CIA videos">
            Video Center
        </a>

        

    
</li>




        </ul></nav><div class="portletWrapper kssattr-portlethash-706c6f6e652e6c656674636f6c756d6e0a636f6e746578740a2f434941676f760a6e617669676174696f6e" id="portletwrapper-706c6f6e652e6c656674636f6c756d6e0a636f6e746578740a2f434941676f760a6e617669676174696f6e" data-portlethash="706c6f6e652e6c656674636f6c756d6e0a636f6e746578740a2f434941676f760a6e617669676174696f6e">
</div></div>
             <div id="content">
               
	    <ul class="breadcrumbs"><li><a href="/">Home</a></li><li><a href="/library">Library</a></li><li><a href="/library/publications">Publications</a></li><li><a href="/library/publications/resources">Resources</a></li><li>The World Factbook</li>
	    </ul><article class="description-box"><a id="main-content" tabindex="-1">&#160;</a>
                 <div class="text-holder-full">

    <div id="content">
                             <div id="content-core">
                             <!-- content goes here -->                      


                   <div id="wfb-text-holder">
<!-- dfl                   <div class="text-box">   -->
				   <!-- Page Content go here -->
<!--                <div class="wfbLogo"><img src="../images/banner_ext2.png" alt="World Factbook Banner" /></div>  -->

                <div id="fbHeader" >
                  <div id="fbLogo">
                    <img title="World Factbook Title" src="../images/banner_ext2.png"></div>
                  <div id="cntrySelect">

                 <script>
                 $(document).ready(function() {
                 // $(".selecter_links").selecter({
                 // defaultLabel: "Please Select a Country to View",
                 // links: true
                 // });
                 $( ".selecter_links" ).change(function(e) {
                 if (this.form.selecter_links.selectedIndex > 0) {
                 window.location = this.form.selecter_links.options[this.form.selecter_links.selectedIndex].value;
                 }
                 });
                 });
                 </script><form action="#" method="GET">
                         <select name="selecter_links" class="selecter_links" onchange="document.location.href= this.value;"><option value="">Please select a country to view</option><option value="../geos/xx.html"> World </option><option value="../geos/af.html"> Afghanistan </option><option value="../geos/ax.html"> Akrotiri </option><option value="../geos/al.html"> Albania </option><option value="../geos/ag.html"> Algeria </option><option value="../geos/aq.html"> American Samoa </option><option value="../geos/an.html"> Andorra </option><option value="../geos/ao.html"> Angola </option><option value="../geos/av.html"> Anguilla </option><option value="../geos/ay.html"> Antarctica </option><option value="../geos/ac.html"> Antigua and Barbuda </option><option value="../geos/xq.html"> Arctic Ocean </option><option value="../geos/ar.html"> Argentina </option><option value="../geos/am.html"> Armenia </option><option value="../geos/aa.html"> Aruba </option><option value="../geos/at.html"> Ashmore and Cartier Islands </option><option value="../geos/zh.html"> Atlantic Ocean </option><option value="../geos/as.html"> Australia </option><option value="../geos/au.html"> Austria </option><option value="../geos/aj.html"> Azerbaijan </option><option value="../geos/bf.html"> Bahamas, The </option><option value="../geos/ba.html"> Bahrain </option><option value="../geos/um.html"> Baker Island </option><option value="../geos/bg.html"> Bangladesh </option><option value="../geos/bb.html"> Barbados </option><option value="../geos/bo.html"> Belarus </option><option value="../geos/be.html"> Belgium </option><option value="../geos/bh.html"> Belize </option><option value="../geos/bn.html"> Benin </option><option value="../geos/bd.html"> Bermuda </option><option value="../geos/bt.html"> Bhutan </option><option value="../geos/bl.html"> Bolivia </option><option value="../geos/bk.html"> Bosnia and Herzegovina </option><option value="../geos/bc.html"> Botswana </option><option value="../geos/bv.html"> Bouvet Island </option><option value="../geos/br.html"> Brazil </option><option value="../geos/io.html"> British Indian Ocean Territory </option><option value="../geos/vi.html"> British Virgin Islands </option><option value="../geos/bx.html"> Brunei </option><option value="../geos/bu.html"> Bulgaria </option><option value="../geos/uv.html"> Burkina Faso </option><option value="../geos/bm.html"> Burma </option><option value="../geos/by.html"> Burundi </option><option value="../geos/cv.html"> Cabo Verde </option><option value="../geos/cb.html"> Cambodia </option><option value="../geos/cm.html"> Cameroon </option><option value="../geos/ca.html"> Canada </option><option value="../geos/cj.html"> Cayman Islands </option><option value="../geos/ct.html"> Central African Republic </option><option value="../geos/cd.html"> Chad </option><option value="../geos/ci.html"> Chile </option><option value="../geos/ch.html"> China </option><option value="../geos/kt.html"> Christmas Island </option><option value="../geos/ip.html"> Clipperton Island </option><option value="../geos/ck.html"> Cocos (Keeling) Islands </option><option value="../geos/co.html"> Colombia </option><option value="../geos/cn.html"> Comoros </option><option value="../geos/cg.html"> Congo, Democratic Republic of the </option><option value="../geos/cf.html"> Congo, Republic of the </option><option value="../geos/cw.html"> Cook Islands </option><option value="../geos/cr.html"> Coral Sea Islands </option><option value="../geos/cs.html"> Costa Rica </option><option value="../geos/iv.html"> Cote d'Ivoire </option><option value="../geos/hr.html"> Croatia </option><option value="../geos/cu.html"> Cuba </option><option value="../geos/cc.html"> Curacao </option><option value="../geos/cy.html"> Cyprus </option><option value="../geos/ez.html"> Czechia </option><option value="../geos/da.html"> Denmark </option><option value="../geos/dx.html"> Dhekelia </option><option value="../geos/dj.html"> Djibouti </option><option value="../geos/do.html"> Dominica </option><option value="../geos/dr.html"> Dominican Republic </option><option value="../geos/ec.html"> Ecuador </option><option value="../geos/eg.html"> Egypt </option><option value="../geos/es.html"> El Salvador </option><option value="../geos/ek.html"> Equatorial Guinea </option><option value="../geos/er.html"> Eritrea </option><option value="../geos/en.html"> Estonia </option><option value="../geos/et.html"> Ethiopia </option><option value="../geos/fk.html"> Falkland Islands (Islas Malvinas) </option><option value="../geos/fo.html"> Faroe Islands </option><option value="../geos/fj.html"> Fiji </option><option value="../geos/fi.html"> Finland </option><option value="../geos/fr.html"> France </option><option value="../geos/fp.html"> French Polynesia </option><option value="../geos/fs.html"> French Southern and Antarctic Lands </option><option value="../geos/gb.html"> Gabon </option><option value="../geos/ga.html"> Gambia, The </option><option value="../geos/gz.html"> Gaza Strip </option><option value="../geos/gg.html"> Georgia </option><option value="../geos/gm.html"> Germany </option><option value="../geos/gh.html"> Ghana </option><option value="../geos/gi.html"> Gibraltar </option><option value="../geos/gr.html"> Greece </option><option value="../geos/gl.html"> Greenland </option><option value="../geos/gj.html"> Grenada </option><option value="../geos/gq.html"> Guam </option><option value="../geos/gt.html"> Guatemala </option><option value="../geos/gk.html"> Guernsey </option><option value="../geos/gv.html"> Guinea </option><option value="../geos/pu.html"> Guinea-Bissau </option><option value="../geos/gy.html"> Guyana </option><option value="../geos/ha.html"> Haiti </option><option value="../geos/hm.html"> Heard Island and McDonald Islands </option><option value="../geos/vt.html"> Holy See (Vatican City) </option><option value="../geos/ho.html"> Honduras </option><option value="../geos/hk.html"> Hong Kong </option><option value="../geos/um.html"> Howland Island </option><option value="../geos/hu.html"> Hungary </option><option value="../geos/ic.html"> Iceland </option><option value="../geos/in.html"> India </option><option value="../geos/xo.html"> Indian Ocean </option><option value="../geos/id.html"> Indonesia </option><option value="../geos/ir.html"> Iran </option><option value="../geos/iz.html"> Iraq </option><option value="../geos/ei.html"> Ireland </option><option value="../geos/im.html"> Isle of Man </option><option value="../geos/is.html"> Israel </option><option value="../geos/it.html"> Italy </option><option value="../geos/jm.html"> Jamaica </option><option value="../geos/jn.html"> Jan Mayen </option><option value="../geos/ja.html"> Japan </option><option value="../geos/um.html"> Jarvis Island </option><option value="../geos/je.html"> Jersey </option><option value="../geos/um.html"> Johnston Atoll </option><option value="../geos/jo.html"> Jordan </option><option value="../geos/kz.html"> Kazakhstan </option><option value="../geos/ke.html"> Kenya </option><option value="../geos/um.html"> Kingman Reef </option><option value="../geos/kr.html"> Kiribati </option><option value="../geos/kn.html"> Korea, North </option><option value="../geos/ks.html"> Korea, South </option><option value="../geos/kv.html"> Kosovo </option><option value="../geos/ku.html"> Kuwait </option><option value="../geos/kg.html"> Kyrgyzstan </option><option value="../geos/la.html"> Laos </option><option value="../geos/lg.html"> Latvia </option><option value="../geos/le.html"> Lebanon </option><option value="../geos/lt.html"> Lesotho </option><option value="../geos/li.html"> Liberia </option><option value="../geos/ly.html"> Libya </option><option value="../geos/ls.html"> Liechtenstein </option><option value="../geos/lh.html"> Lithuania </option><option value="../geos/lu.html"> Luxembourg </option><option value="../geos/mc.html"> Macau </option><option value="../geos/mk.html"> Macedonia </option><option value="../geos/ma.html"> Madagascar </option><option value="../geos/mi.html"> Malawi </option><option value="../geos/my.html"> Malaysia </option><option value="../geos/mv.html"> Maldives </option><option value="../geos/ml.html"> Mali </option><option value="../geos/mt.html"> Malta </option><option value="../geos/rm.html"> Marshall Islands </option><option value="../geos/mr.html"> Mauritania </option><option value="../geos/mp.html"> Mauritius </option><option value="../geos/mx.html"> Mexico </option><option value="../geos/fm.html"> Micronesia, Federated States of </option><option value="../geos/um.html"> Midway Islands </option><option value="../geos/md.html"> Moldova </option><option value="../geos/mn.html"> Monaco </option><option value="../geos/mg.html"> Mongolia </option><option value="../geos/mj.html"> Montenegro </option><option value="../geos/mh.html"> Montserrat </option><option value="../geos/mo.html"> Morocco </option><option value="../geos/mz.html"> Mozambique </option><option value="../geos/wa.html"> Namibia </option><option value="../geos/nr.html"> Nauru </option><option value="../geos/bq.html"> Navassa Island </option><option value="../geos/np.html"> Nepal </option><option value="../geos/nl.html"> Netherlands </option><option value="../geos/nc.html"> New Caledonia </option><option value="../geos/nz.html"> New Zealand </option><option value="../geos/nu.html"> Nicaragua </option><option value="../geos/ng.html"> Niger </option><option value="../geos/ni.html"> Nigeria </option><option value="../geos/ne.html"> Niue </option><option value="../geos/nf.html"> Norfolk Island </option><option value="../geos/cq.html"> Northern Mariana Islands </option><option value="../geos/no.html"> Norway </option><option value="../geos/mu.html"> Oman </option><option value="../geos/zn.html"> Pacific Ocean </option><option value="../geos/pk.html"> Pakistan </option><option value="../geos/ps.html"> Palau </option><option value="../geos/um.html"> Palmyra Atoll </option><option value="../geos/pm.html"> Panama </option><option value="../geos/pp.html"> Papua New Guinea </option><option value="../geos/pf.html"> Paracel Islands </option><option value="../geos/pa.html"> Paraguay </option><option value="../geos/pe.html"> Peru </option><option value="../geos/rp.html"> Philippines </option><option value="../geos/pc.html"> Pitcairn Islands </option><option value="../geos/pl.html"> Poland </option><option value="../geos/po.html"> Portugal </option><option value="../geos/rq.html"> Puerto Rico </option><option value="../geos/qa.html"> Qatar </option><option value="../geos/ro.html"> Romania </option><option value="../geos/rs.html"> Russia </option><option value="../geos/rw.html"> Rwanda </option><option value="../geos/tb.html"> Saint Barthelemy </option><option value="../geos/sh.html"> Saint Helena, Ascension, and Tristan da Cunha </option><option value="../geos/sc.html"> Saint Kitts and Nevis </option><option value="../geos/st.html"> Saint Lucia </option><option value="../geos/rn.html"> Saint Martin </option><option value="../geos/sb.html"> Saint Pierre and Miquelon </option><option value="../geos/vc.html"> Saint Vincent and the Grenadines </option><option value="../geos/ws.html"> Samoa </option><option value="../geos/sm.html"> San Marino </option><option value="../geos/tp.html"> Sao Tome and Principe </option><option value="../geos/sa.html"> Saudi Arabia </option><option value="../geos/sg.html"> Senegal </option><option value="../geos/ri.html"> Serbia </option><option value="../geos/se.html"> Seychelles </option><option value="../geos/sl.html"> Sierra Leone </option><option value="../geos/sn.html"> Singapore </option><option value="../geos/sk.html"> Sint Maarten </option><option value="../geos/lo.html"> Slovakia </option><option value="../geos/si.html"> Slovenia </option><option value="../geos/bp.html"> Solomon Islands </option><option value="../geos/so.html"> Somalia </option><option value="../geos/sf.html"> South Africa </option><option value="../geos/oo.html"> Southern Ocean </option><option value="../geos/sx.html"> South Georgia and South Sandwich Islands </option><option value="../geos/od.html"> South Sudan </option><option value="../geos/sp.html"> Spain </option><option value="../geos/pg.html"> Spratly Islands </option><option value="../geos/ce.html"> Sri Lanka </option><option value="../geos/su.html"> Sudan </option><option value="../geos/ns.html"> Suriname </option><option value="../geos/sv.html"> Svalbard </option><option value="../geos/wz.html"> Swaziland </option><option value="../geos/sw.html"> Sweden </option><option value="../geos/sz.html"> Switzerland </option><option value="../geos/sy.html"> Syria </option><option value="../geos/tw.html"> Taiwan </option><option value="../geos/ti.html"> Tajikistan </option><option value="../geos/tz.html"> Tanzania </option><option value="../geos/th.html"> Thailand </option><option value="../geos/tt.html"> Timor-Leste </option><option value="../geos/to.html"> Togo </option><option value="../geos/tl.html"> Tokelau </option><option value="../geos/tn.html"> Tonga </option><option value="../geos/td.html"> Trinidad and Tobago </option><option value="../geos/ts.html"> Tunisia </option><option value="../geos/tu.html"> Turkey </option><option value="../geos/tx.html"> Turkmenistan </option><option value="../geos/tk.html"> Turks and Caicos Islands </option><option value="../geos/tv.html"> Tuvalu </option><option value="../geos/ug.html"> Uganda </option><option value="../geos/up.html"> Ukraine </option><option value="../geos/ae.html"> United Arab Emirates </option><option value="../geos/uk.html"> United Kingdom </option><option value="../geos/us.html"> United States </option><option value="../geos/um.html"> United States Pacific Island Wildlife Refuges </option><option value="../geos/uy.html"> Uruguay </option><option value="../geos/uz.html"> Uzbekistan </option><option value="../geos/nh.html"> Vanuatu </option><option value="../geos/ve.html"> Venezuela </option><option value="../geos/vm.html"> Vietnam </option><option value="../geos/vq.html"> Virgin Islands </option><option value="../geos/wq.html"> Wake Island </option><option value="../geos/wf.html"> Wallis and Futuna </option><option value="../geos/we.html"> West Bank </option><option value="../geos/wi.html"> Western Sahara </option><option value="../geos/ym.html"> Yemen </option><option value="../geos/za.html"> Zambia </option><option value="../geos/zi.html"> Zimbabwe </option><option value="../geos/ee.html"> European Union </option></select></form>
                </div>  <!-- cntrySelect -->

                      <div class="fbnav">
                        <div class="fbNavBox">
                         <script type="text/javascript">
                         var timeout         = 500;
                         var closetimer          = 0;
                         var ddmenuitem      = 0;

                         function wfbNav_open()
                         {       wfbNav_canceltimer();
                                 wfbNav_close();
                                 ddmenuitem = $(this).find('ul').eq(0).css('visibility', 'visible');}

                         function wfbNav_close()
                         {       if(ddmenuitem) ddmenuitem.css('visibility', 'hidden');}

                         function wfbNav_timer()
                         {       closetimer = window.setTimeout(wfbNav_close, timeout);}

                         function wfbNav_canceltimer()
                         {       if(closetimer)
                                 {       window.clearTimeout(closetimer);
                                         closetimer = null;}}


                         $(document).ready(function()
                         {       $('#wfbNav > li').bind('mouseover', wfbNav_open);
                                 $('#wfbNav > li').bind('mouseout',  wfbNav_timer);});


                         document.onclick = wfbNav_close;
                         </script><ul id="wfbNav" style="z-index: 9999;"><li style="border-bottom: 2px solid #CCCCCC; "><a href="../" style="width:20px; height: 12px;" title="The World Factbook Home"><img src="../graphics/home_on.png" border="0"></a></li>
                                        <li style="border-bottom: 2px solid #CCCCCC; "><a href="javascript:void(0);" style="width:65px;" title="About">ABOUT</a>
                                                <ul class="sub_menu"><li><a href="../docs/history.html">History</a></li>
                                                        <li><a href="../docs/contributor_copyright.html">Copyright and Contributors</a></li>
                                                        <li><a href="../docs/purchase_info.html">Purchasing</a></li>
                                                        <li><a href="../docs/didyouknow.html">Did You Know?</a></li>

                                                </ul></li>
                                        <li style="border-bottom: 2px solid #CCCCCC; "><a href="javascript:void(0);" style="width:95px;" title="References">REFERENCES</a>
                                                <ul class="sub_menu"><li><a href="../docs/refmaps.html">Regional and World Maps</a></li>
                                                        <li><a href="../docs/flagsoftheworld.html">Flags of the World</a></li>
                                                        <li><a href="../docs/gallery.html">Gallery of Covers</a></li>
                                                        <li><a href="../docs/notesanddefs.html">Definitions and Notes</a></li>
                                                        <li><a href="../docs/profileguide.html">Guide to Country Profiles</a></li>
                                                        <li><a href="../rankorder/rankorderguide.html">Guide to Country Comparisons</a></li>

                                                        <li><a href="../docs/guidetowfbook.html">The World Factbook Users Guide</a></li>
                                                </ul></li>
                                        <li style="border-bottom: 2px solid #CCCCCC; "><a href="javascript:void(0);" title="Appendices">APPENDICES</a>
                                                <ul class="sub_menu"><li><a href="../appendix/appendix-a.html">A: abbreviations</a></li>
                                                        <li><a href="../appendix/appendix-b.html">B: international organizations and groups</a></li>
                                                        <li><a href="../appendix/appendix-c.html">C: selected international environmental agreements</a></li>
                                                        <li><a href="../appendix/appendix-d.html">D: cross-reference list of country data codes</a></li>
                                                        <li><a href="../appendix/appendix-e.html">E: cross-reference list of hydrographic data codes</a></li>
                                                        <li><a href="../appendix/appendix-f.html">F: cross-reference list of geographic names</a></li>
                                                        <li><a href="../appendix/appendix-g.html">G: weights and measures</a></li>
                                                </ul></li>
                                        <li id="faqs" style="border-bottom: 2px solid #CCCCCC; "><a href="../docs/faqs.html" style="cursor:pointer;width:50px;">FAQ<span style="text-transform:lowercase;">s</span></a></li>

                                        <li id="contact" style="border-bottom: 2px solid #CCCCCC; "> <a href="/contact-cia/index.html" title="Contact" style="cursor:pointer;width:73px;"> CONTACT </a> </li>

                                </ul></div>   <!-- wfbnav -->
                     <div class="lowband_dl">
                       <ul><li class="low-band"><a href="../print/textversion.html"><img src="../graphics/view_text.jpg"></a></li>
                         <li class="fb-download"><a href="/library/publications/download/index.html"><img src="../graphics/download_pub.jpg"></a></li>
                       </ul>
                     </div>
                        </div>  <!--  fbNavBox  -->

 
             </div><!-- fbHeader -->


                   <div class="wfb-text-box">

<!-- generated content -->
<div class="region1 geos_title aus_dark">Australia-Oceania  <strong>::</strong> <span class="region_name1 countryName ">AUSTRALIA</span><a class="printVersion" href="print_as.html"><img style="padding: 3px;" src="../graphics/print.gif"></a></div>

     <div class="img-block-wrapper">
         <div class="flag-photo-holder">
           <div class="flag-location" style="background-image:url(../graphics/aus_lgmap_bkgrnd.jpg); background-repeat: repeat-x; background-position: top ;">
              <div class="lastMod">Page last updated on January 12, 2017</div>
              <div class="flagBox"><a data-toggle="modal" href="#cntryFlagModal"><img src="../graphics/flags/large/as-lgflag.gif"></a></div>
        <div class="modal fade" id="cntryFlagModal" role="dialog">
           <div class="wfb-modal-dialog">
              <div class="modal-content"  >
                 <div class="wfb-modal-header" style="border-radius: 4px; font-family: Verdana,Arial,sans-serif; font-size: 14px !important; font-weight: bold;  padding: 0.4em 16px 0.4em 1em; background: #cccccc url("..images/ui-bg_highlight-soft_75_cccccc_1x100.png") repeat-x scroll 50% 50%;" >
                    <span style="font-size: 14px !important; margin: 0.1em 16px 0.1em 0;" class="modal-title wfb-title">The World Factbook</span><span style="float: right; margin-top: -4px;">
                    <button type="button" class="close" title="close" data-dismiss="modal">&times;</button></span>
                 </div>
                 <div class="wfb-modal-body">
                    <div class="region1 geos_title aus_dark">Australia-Oceania   <strong>::</strong><span class="region_name1 countryName ">AUSTRALIA</span>
                       <a class="printVersion" href="../graphics/flags/large/as-lgflag.gif"><img src="../graphics/print.gif" style="padding: 3px;"></a>
                    </div>
           <div class="modal-map-holder" style="background-image:url(../graphics/aus_lgmap_bkgrnd.jpg); background-repeat: repeat-x; background-position: top ;">
             <div class="modalFlagBox" style="width: 50%; padding: 33px 0 0 0; margin-left: 30px;">
                  <img class="aus_lgflagborder" src="../graphics/flags/large/as-lgflag.gif">
             </div>
             <div class="modalFlagDesc"  >
                <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Flag Description </div>
                <div class="photogallery_captiontext" style="height: 375px;background-color:white;">
                  <span class="flag_description_text">blue with the flag of the UK in the upper hoist-side quadrant and a large seven-pointed star in the lower hoist-side quadrant known as the Commonwealth or Federation Star, representing the federation of the colonies of Australia in 1901; the star depicts one point for each of the six original states and one representing all of Australia's internal and external territories; on the fly half is a representation of the Southern Cross constellation in white with one small, five-pointed star and four larger, seven-pointed stars</span>
                </div>
             </div>
           </div>
                 </div>
              </div>
            </div>
        </div>

              <div class="locatorBox"><a data-toggle="modal" href="#cntryLocatorModal"><img src="../graphics/locator/aus/as_large_locator.gif"></a></div>
           </div>
        <div class="modal fade" id="cntryLocatorModal" role="dialog">
           <div class="wfb-modal-dialog">
              <div class="modal-content"  >
                 <div class="wfb-modal-header" style="border-radius: 4px; font-family: Verdana,Arial,sans-serif; font-size: 14px !important; font-weight: bold;  padding: 0.4em 16px 0.4em 1em; background: #cccccc url("..images/ui-bg_highlight-soft_75_cccccc_1x100.png") repeat-x scroll 50% 50%;" >
                    <span style="font-size: 14px !important; margin: 0.1em 16px 0.1em 0;" class="modal-title wfb-title">The World Factbook</span><span style="float: right; margin-top: -4px;">
                    <button type="button" class="close" title="close" data-dismiss="modal">&times;</button></span>
                 </div>
                 <div class="wfb-modal-body">
                    <div class="region1 geos_title aus_dark">Australia-Oceania   <strong>::</strong><span class="region_name1 countryName ">AUSTRALIA</span>
                       <a class="printVersion" href="../graphics/locator/aus/as_large_locator.gif"><img src="../graphics/print.gif" style="padding: 3px;"></a>
                    </div>
           <div class="modal-map-holder" style="background-image:url(../graphics/aus_lgmap_bkgrnd.jpg); background-repeat: repeat-x; background-position: top ;">
             <div class="mapBox">
                  <img class="aus_lgflagborder" src="../graphics/locator/aus/as_large_locator.gif">
             </div>
           </div>
                 </div>
              </div>
            </div>
        </div> 

           <div class="photo_bkgrnd_static">
               <div class="gallery_icon_holder">
                  <a data-toggle="modal" href="#wfbPhotoGalleryModal"><img id="photoDialog" width="123" height="81" border="0" style="padding-top:10px;" src="../graphics/photo_on.gif"></a>
               </div>
               <div class="smalltext_nav photoDialogText">View <a data-toggle="modal" href="#wfbPhotoGalleryModal"><strong>23 photos</strong></a> of <br />AUSTRALIA</div>
 
             </div>
          </div>
        <div class="modal fade" id="wfbPhotoGalleryModal" role="dialog">
           <div class="wfb-modal-dialog">
              <div class="modal-content"  >
                 <div class="wfb-modal-header" style="border-radius: 4px; font-family: Verdana,Arial,sans-serif; font-size: 14px !important; font-weight: bold;  padding: 0.4em 16px 0.4em 1em; background: #cccccc url("..images/ui-bg_highlight-soft_75_cccccc_1x100.png") repeat-x scroll 50% 50%;" >
                    <span style="font-size: 14px !important; margin: 0.1em 16px 0.1em 0;" class="modal-title wfb-title">The World Factbook</span><span style="float: right; margin-top: -4px;">
                    <button type="button" class="close" title="close" data-dismiss="modal">&times;</button></span>
                 </div>
                 <div class="wfb-modal-body">
                    <div class="region1 geos_title aus_dark">Australia-Oceania   <strong>::</strong><span class="region_name1 countryName ">AUSTRALIA</span>
                    </div>
           <div class="gallery-modal-map-holder" >
             <div class="gallery-mapBox">

<div id="myCarousel" class="carousel carousel-fade" data-ride="carousel" data-interval="false" >

  <!-- Wrapper for slides --> 
  <div class="carousel-inner" role="listbox">

    <div class="item active">
      <img src="../photo_gallery/as/images/AS_062.jpg" alt="AUSTRALIA">
<div class="galleria-counter" style="opacity: 1; transition: none 0s ease 0s ;">
<span class="galleria-current">1</span>
/
<span class="galleria-total">23</span>
</div>
      <div class="carousel-caption">
        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Caption </div>
        <div class="photoCaption"><p class="flag_description_text">Exterior view of Sydney Aquarium on the eastern side of Darling Harbour.</p></div>
      </div>
      <div class="carousel-photo-info">
        <div class="photoInfo" style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699; line-height: 24px;">
           <div>Dimensions</div><div>File Size</div><div>Download</div>
        </div>
        <div class="photoInfo">
           <div class="flag_description_text">1500X1125</div><div class="flag_description_text">297.0 KB</div><div><a href="../photo_gallery/as/images/large/AS_062_large.jpg"><img class="wfbDownload" src="../graphics/download_image.png" style="padding:0; margin-left: auto; margin-right: auto;" /></a></div>
        </div>
      
        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Usage </div>
        <p class="flag_description_text" style="padding-left: 5px;">Factbook photos - obtained from a variety of sources - are in the public domain and are copyright free. <br /> <a href=/about-cia/site-policies/#copy title=Agency Copyright Notice>Agency Copyright Notice</a></p>
      </div>
   </div>

    <div class="item">
      <img src="../photo_gallery/as/images/AS_061.jpg" alt="AUSTRALIA">

<div class="galleria-counter" style="opacity: 1; transition: none 0s ease 0s ;">
<span class="galleria-current">2</span>
/
<span class="galleria-total">23</span>
</div>


      <div class="carousel-caption">
        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Caption </div>
        <div class="photoCaption"><p class="flag_description_text">Night view of Sydney Harbor Bridge.</p></div>
      </div>
      <div class="carousel-photo-info">
        <div class="photoInfo" style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699; line-height: 24px;">
           <div>Dimensions</div><div>File Size</div><div>Download</div>
        </div>
        <div class="photoInfo">
           <div class="flag_description_text">1500X997</div><div class="flag_description_text">280.6 KB</div><div><a href="../photo_gallery/as/images/large/AS_061_large.jpg"><img class="wfbDownload" src="../graphics/download_image.png" style="padding:0; margin-left: auto; margin-right: auto;" /></a></div>
        </div>

        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Usage </div>
        <p class="flag_description_text" style="padding-left: 5px;">Factbook photos - obtained from a variety of sources - are in the public domain and are copyright free. <br /> <a href=/about-cia/site-policies/#copy title=Agency Copyright Notice>Agency Copyright Notice</a></p>
      </div>
   </div>

    <div class="item">
      <img src="../photo_gallery/as/images/AS_060.jpg" alt="AUSTRALIA">

<div class="galleria-counter" style="opacity: 1; transition: none 0s ease 0s ;">
<span class="galleria-current">3</span>
/
<span class="galleria-total">23</span>
</div>


      <div class="carousel-caption">
        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Caption </div>
        <div class="photoCaption"><p class="flag_description_text">Sydney skyline as seen from a Darling Harbour footbridge.</p></div>
      </div>
      <div class="carousel-photo-info">
        <div class="photoInfo" style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699; line-height: 24px;">
           <div>Dimensions</div><div>File Size</div><div>Download</div>
        </div>
        <div class="photoInfo">
           <div class="flag_description_text">1500X1002</div><div class="flag_description_text">201.3 KB</div><div><a href="../photo_gallery/as/images/large/AS_060_large.jpg"><img class="wfbDownload" src="../graphics/download_image.png" style="padding:0; margin-left: auto; margin-right: auto;" /></a></div>
        </div>

        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Usage </div>
        <p class="flag_description_text" style="padding-left: 5px;">Factbook photos - obtained from a variety of sources - are in the public domain and are copyright free. <br /> <a href=/about-cia/site-policies/#copy title=Agency Copyright Notice>Agency Copyright Notice</a></p>
      </div>
   </div>

    <div class="item">
      <img src="../photo_gallery/as/images/AS_059.jpg" alt="AUSTRALIA">

<div class="galleria-counter" style="opacity: 1; transition: none 0s ease 0s ;">
<span class="galleria-current">4</span>
/
<span class="galleria-total">23</span>
</div>


      <div class="carousel-caption">
        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Caption </div>
        <div class="photoCaption"><p class="flag_description_text">HM Bark Endeavour, a replica ship anchored at the Australian National Maritime Museum in Darling Harbour, Sydney. The original Endeavour was commanded by Lt. James Cook during his first voyage of discovery (1768-1771) where he mapped the New Zealand coast and explored the eastern coast of Australia.</p></div>
      </div>
      <div class="carousel-photo-info">
        <div class="photoInfo" style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699; line-height: 24px;">
           <div>Dimensions</div><div>File Size</div><div>Download</div>
        </div>
        <div class="photoInfo">
           <div class="flag_description_text">1500X1123</div><div class="flag_description_text">228.6 KB</div><div><a href="../photo_gallery/as/images/large/AS_059_large.jpg"><img class="wfbDownload" src="../graphics/download_image.png" style="padding:0; margin-left: auto; margin-right: auto;" /></a></div>
        </div>

        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Usage </div>
        <p class="flag_description_text" style="padding-left: 5px;">Factbook photos - obtained from a variety of sources - are in the public domain and are copyright free. <br /> <a href=/about-cia/site-policies/#copy title=Agency Copyright Notice>Agency Copyright Notice</a></p>
      </div>
   </div>

    <div class="item">
      <img src="../photo_gallery/as/images/AS_058.jpg" alt="AUSTRALIA">

<div class="galleria-counter" style="opacity: 1; transition: none 0s ease 0s ;">
<span class="galleria-current">5</span>
/
<span class="galleria-total">23</span>
</div>


      <div class="carousel-caption">
        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Caption </div>
        <div class="photoCaption"><p class="flag_description_text">The retired naval destroyer HMAS Vampire at the Australian National Maritime Museum in Darling Harbour, Sydney.</p></div>
      </div>
      <div class="carousel-photo-info">
        <div class="photoInfo" style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699; line-height: 24px;">
           <div>Dimensions</div><div>File Size</div><div>Download</div>
        </div>
        <div class="photoInfo">
           <div class="flag_description_text">1500X994</div><div class="flag_description_text">221.7 KB</div><div><a href="../photo_gallery/as/images/large/AS_058_large.jpg"><img class="wfbDownload" src="../graphics/download_image.png" style="padding:0; margin-left: auto; margin-right: auto;" /></a></div>
        </div>

        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Usage </div>
        <p class="flag_description_text" style="padding-left: 5px;">Factbook photos - obtained from a variety of sources - are in the public domain and are copyright free. <br /> <a href=/about-cia/site-policies/#copy title=Agency Copyright Notice>Agency Copyright Notice</a></p>
      </div>
   </div>

    <div class="item">
      <img src="../photo_gallery/as/images/AS_057.jpg" alt="AUSTRALIA">

<div class="galleria-counter" style="opacity: 1; transition: none 0s ease 0s ;">
<span class="galleria-current">6</span>
/
<span class="galleria-total">23</span>
</div>


      <div class="carousel-caption">
        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Caption </div>
        <div class="photoCaption"><p class="flag_description_text">Rescue boat docked in Darling Harbour, Sydney.</p></div>
      </div>
      <div class="carousel-photo-info">
        <div class="photoInfo" style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699; line-height: 24px;">
           <div>Dimensions</div><div>File Size</div><div>Download</div>
        </div>
        <div class="photoInfo">
           <div class="flag_description_text">1500X1123</div><div class="flag_description_text">263.2 KB</div><div><a href="../photo_gallery/as/images/large/AS_057_large.jpg"><img class="wfbDownload" src="../graphics/download_image.png" style="padding:0; margin-left: auto; margin-right: auto;" /></a></div>
        </div>

        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Usage </div>
        <p class="flag_description_text" style="padding-left: 5px;">Factbook photos - obtained from a variety of sources - are in the public domain and are copyright free. <br /> <a href=/about-cia/site-policies/#copy title=Agency Copyright Notice>Agency Copyright Notice</a></p>
      </div>
   </div>

    <div class="item">
      <img src="../photo_gallery/as/images/AS_056.jpg" alt="AUSTRALIA">

<div class="galleria-counter" style="opacity: 1; transition: none 0s ease 0s ;">
<span class="galleria-current">7</span>
/
<span class="galleria-total">23</span>
</div>


      <div class="carousel-caption">
        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Caption </div>
        <div class="photoCaption"><p class="flag_description_text">The Sydney Tower is the tallest free-standing structure in the city and the second-tallest in Australia. The tower stands 309 m (1,014 ft) above the central business district.</p></div>
      </div>
      <div class="carousel-photo-info">
        <div class="photoInfo" style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699; line-height: 24px;">
           <div>Dimensions</div><div>File Size</div><div>Download</div>
        </div>
        <div class="photoInfo">
           <div class="flag_description_text">1500X1122</div><div class="flag_description_text">110.9 KB</div><div><a href="../photo_gallery/as/images/large/AS_056_large.jpg"><img class="wfbDownload" src="../graphics/download_image.png" style="padding:0; margin-left: auto; margin-right: auto;" /></a></div>
        </div>

        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Usage </div>
        <p class="flag_description_text" style="padding-left: 5px;">Factbook photos - obtained from a variety of sources - are in the public domain and are copyright free. <br /> <a href=/about-cia/site-policies/#copy title=Agency Copyright Notice>Agency Copyright Notice</a></p>
      </div>
   </div>

    <div class="item">
      <img src="../photo_gallery/as/images/AS_055.jpg" alt="AUSTRALIA">

<div class="galleria-counter" style="opacity: 1; transition: none 0s ease 0s ;">
<span class="galleria-current">8</span>
/
<span class="galleria-total">23</span>
</div>


      <div class="carousel-caption">
        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Caption </div>
        <div class="photoCaption"><p class="flag_description_text">A close up of the Three Sisters sandstone rock formation in the Blue Mountains.</p></div>
      </div>
      <div class="carousel-photo-info">
        <div class="photoInfo" style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699; line-height: 24px;">
           <div>Dimensions</div><div>File Size</div><div>Download</div>
        </div>
        <div class="photoInfo">
           <div class="flag_description_text">1500X1065</div><div class="flag_description_text">223.3 KB</div><div><a href="../photo_gallery/as/images/large/AS_055_large.jpg"><img class="wfbDownload" src="../graphics/download_image.png" style="padding:0; margin-left: auto; margin-right: auto;" /></a></div>
        </div>

        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Usage </div>
        <p class="flag_description_text" style="padding-left: 5px;">Factbook photos - obtained from a variety of sources - are in the public domain and are copyright free. <br /> <a href=/about-cia/site-policies/#copy title=Agency Copyright Notice>Agency Copyright Notice</a></p>
      </div>
   </div>

    <div class="item">
      <img src="../photo_gallery/as/images/AS_054.jpg" alt="AUSTRALIA">

<div class="galleria-counter" style="opacity: 1; transition: none 0s ease 0s ;">
<span class="galleria-current">9</span>
/
<span class="galleria-total">23</span>
</div>


      <div class="carousel-caption">
        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Caption </div>
        <div class="photoCaption"><p class="flag_description_text">The Three Sisters sandstone rock formation in the Blue Mountains west of Sydney.</p></div>
      </div>
      <div class="carousel-photo-info">
        <div class="photoInfo" style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699; line-height: 24px;">
           <div>Dimensions</div><div>File Size</div><div>Download</div>
        </div>
        <div class="photoInfo">
           <div class="flag_description_text">1500X1012</div><div class="flag_description_text">185.0 KB</div><div><a href="../photo_gallery/as/images/large/AS_054_large.jpg"><img class="wfbDownload" src="../graphics/download_image.png" style="padding:0; margin-left: auto; margin-right: auto;" /></a></div>
        </div>

        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Usage </div>
        <p class="flag_description_text" style="padding-left: 5px;">Factbook photos - obtained from a variety of sources - are in the public domain and are copyright free. <br /> <a href=/about-cia/site-policies/#copy title=Agency Copyright Notice>Agency Copyright Notice</a></p>
      </div>
   </div>

    <div class="item">
      <img src="../photo_gallery/as/images/AS_053.jpg" alt="AUSTRALIA">

<div class="galleria-counter" style="opacity: 1; transition: none 0s ease 0s ;">
<span class="galleria-current">10</span>
/
<span class="galleria-total">23</span>
</div>


      <div class="carousel-caption">
        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Caption </div>
        <div class="photoCaption"><p class="flag_description_text">A view of the countryside in the state of Victoria - approaching Bendigo from Melbourne.</p></div>
      </div>
      <div class="carousel-photo-info">
        <div class="photoInfo" style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699; line-height: 24px;">
           <div>Dimensions</div><div>File Size</div><div>Download</div>
        </div>
        <div class="photoInfo">
           <div class="flag_description_text">1500X1125</div><div class="flag_description_text">187.6 KB</div><div><a href="../photo_gallery/as/images/large/AS_053_large.jpg"><img class="wfbDownload" src="../graphics/download_image.png" style="padding:0; margin-left: auto; margin-right: auto;" /></a></div>
        </div>

        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Usage </div>
        <p class="flag_description_text" style="padding-left: 5px;">Factbook photos - obtained from a variety of sources - are in the public domain and are copyright free. <br /> <a href=/about-cia/site-policies/#copy title=Agency Copyright Notice>Agency Copyright Notice</a></p>
      </div>
   </div>

    <div class="item">
      <img src="../photo_gallery/as/images/AS_052.JPG" alt="AUSTRALIA">

<div class="galleria-counter" style="opacity: 1; transition: none 0s ease 0s ;">
<span class="galleria-current">11</span>
/
<span class="galleria-total">23</span>
</div>


      <div class="carousel-caption">
        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Caption </div>
        <div class="photoCaption"><p class="flag_description_text">The historic Shamrock Hotel is one of the finest examples of Victorian-era architecture in Bendigo, a city renowned for its 19th century buildings. </p></div>
      </div>
      <div class="carousel-photo-info">
        <div class="photoInfo" style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699; line-height: 24px;">
           <div>Dimensions</div><div>File Size</div><div>Download</div>
        </div>
        <div class="photoInfo">
           <div class="flag_description_text">1500X1125</div><div class="flag_description_text">252.3 KB</div><div><a href="../photo_gallery/as/images/large/AS_052_large.JPG"><img class="wfbDownload" src="../graphics/download_image.png" style="padding:0; margin-left: auto; margin-right: auto;" /></a></div>
        </div>

        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Usage </div>
        <p class="flag_description_text" style="padding-left: 5px;">Factbook photos - obtained from a variety of sources - are in the public domain and are copyright free. <br /> <a href=/about-cia/site-policies/#copy title=Agency Copyright Notice>Agency Copyright Notice</a></p>
      </div>
   </div>

    <div class="item">
      <img src="../photo_gallery/as/images/AS_051.JPG" alt="AUSTRALIA">

<div class="galleria-counter" style="opacity: 1; transition: none 0s ease 0s ;">
<span class="galleria-current">12</span>
/
<span class="galleria-total">23</span>
</div>


      <div class="carousel-caption">
        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Caption </div>
        <div class="photoCaption"><p class="flag_description_text">Another view of the Bendigo Shamrock Hotel and its unique two-story veranda. </p></div>
      </div>
      <div class="carousel-photo-info">
        <div class="photoInfo" style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699; line-height: 24px;">
           <div>Dimensions</div><div>File Size</div><div>Download</div>
        </div>
        <div class="photoInfo">
           <div class="flag_description_text">1125X1500</div><div class="flag_description_text">256.5 KB</div><div><a href="../photo_gallery/as/images/large/AS_051_large.JPG"><img class="wfbDownload" src="../graphics/download_image.png" style="padding:0; margin-left: auto; margin-right: auto;" /></a></div>
        </div>

        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Usage </div>
        <p class="flag_description_text" style="padding-left: 5px;">Factbook photos - obtained from a variety of sources - are in the public domain and are copyright free. <br /> <a href=/about-cia/site-policies/#copy title=Agency Copyright Notice>Agency Copyright Notice</a></p>
      </div>
   </div>

    <div class="item">
      <img src="../photo_gallery/as/images/AS_050.jpg" alt="AUSTRALIA">

<div class="galleria-counter" style="opacity: 1; transition: none 0s ease 0s ;">
<span class="galleria-current">13</span>
/
<span class="galleria-total">23</span>
</div>


      <div class="carousel-caption">
        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Caption </div>
        <div class="photoCaption"><p class="flag_description_text">A monument in Bendigo commemorating the Australian men who fell in the South African (Boer) War (1899-1902).</p></div>
      </div>
      <div class="carousel-photo-info">
        <div class="photoInfo" style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699; line-height: 24px;">
           <div>Dimensions</div><div>File Size</div><div>Download</div>
        </div>
        <div class="photoInfo">
           <div class="flag_description_text">1500X1125</div><div class="flag_description_text">143.1 KB</div><div><a href="../photo_gallery/as/images/large/AS_050_large.jpg"><img class="wfbDownload" src="../graphics/download_image.png" style="padding:0; margin-left: auto; margin-right: auto;" /></a></div>
        </div>

        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Usage </div>
        <p class="flag_description_text" style="padding-left: 5px;">Factbook photos - obtained from a variety of sources - are in the public domain and are copyright free. <br /> <a href=/about-cia/site-policies/#copy title=Agency Copyright Notice>Agency Copyright Notice</a></p>
      </div>
   </div>

    <div class="item">
      <img src="../photo_gallery/as/images/AS_049.JPG" alt="AUSTRALIA">

<div class="galleria-counter" style="opacity: 1; transition: none 0s ease 0s ;">
<span class="galleria-current">14</span>
/
<span class="galleria-total">23</span>
</div>


      <div class="carousel-caption">
        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Caption </div>
        <div class="photoCaption"><p class="flag_description_text">View of a Blue Mountains cable car as seen from the bottom of the funicular railway.</p></div>
      </div>
      <div class="carousel-photo-info">
        <div class="photoInfo" style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699; line-height: 24px;">
           <div>Dimensions</div><div>File Size</div><div>Download</div>
        </div>
        <div class="photoInfo">
           <div class="flag_description_text">1500X1125</div><div class="flag_description_text">217.9 KB</div><div><a href="../photo_gallery/as/images/large/AS_049_large.JPG"><img class="wfbDownload" src="../graphics/download_image.png" style="padding:0; margin-left: auto; margin-right: auto;" /></a></div>
        </div>

        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Usage </div>
        <p class="flag_description_text" style="padding-left: 5px;">Factbook photos - obtained from a variety of sources - are in the public domain and are copyright free. <br /> <a href=/about-cia/site-policies/#copy title=Agency Copyright Notice>Agency Copyright Notice</a></p>
      </div>
   </div>

    <div class="item">
      <img src="../photo_gallery/as/images/AS_048.jpg" alt="AUSTRALIA">

<div class="galleria-counter" style="opacity: 1; transition: none 0s ease 0s ;">
<span class="galleria-current">15</span>
/
<span class="galleria-total">23</span>
</div>


      <div class="carousel-caption">
        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Caption </div>
        <div class="photoCaption"><p class="flag_description_text">Warning sign posted at the Blue Mountains funicular railway.</p></div>
      </div>
      <div class="carousel-photo-info">
        <div class="photoInfo" style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699; line-height: 24px;">
           <div>Dimensions</div><div>File Size</div><div>Download</div>
        </div>
        <div class="photoInfo">
           <div class="flag_description_text">1500X1125</div><div class="flag_description_text">188.0 KB</div><div><a href="../photo_gallery/as/images/large/AS_048_large.jpg"><img class="wfbDownload" src="../graphics/download_image.png" style="padding:0; margin-left: auto; margin-right: auto;" /></a></div>
        </div>

        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Usage </div>
        <p class="flag_description_text" style="padding-left: 5px;">Factbook photos - obtained from a variety of sources - are in the public domain and are copyright free. <br /> <a href=/about-cia/site-policies/#copy title=Agency Copyright Notice>Agency Copyright Notice</a></p>
      </div>
   </div>

    <div class="item">
      <img src="../photo_gallery/as/images/AS_047.jpg" alt="AUSTRALIA">

<div class="galleria-counter" style="opacity: 1; transition: none 0s ease 0s ;">
<span class="galleria-current">16</span>
/
<span class="galleria-total">23</span>
</div>


      <div class="carousel-caption">
        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Caption </div>
        <div class="photoCaption"><p class="flag_description_text">A view of the Blue Mountains as seen from a cable car.</p></div>
      </div>
      <div class="carousel-photo-info">
        <div class="photoInfo" style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699; line-height: 24px;">
           <div>Dimensions</div><div>File Size</div><div>Download</div>
        </div>
        <div class="photoInfo">
           <div class="flag_description_text">1500X1125</div><div class="flag_description_text">208.1 KB</div><div><a href="../photo_gallery/as/images/large/AS_047_large.jpg"><img class="wfbDownload" src="../graphics/download_image.png" style="padding:0; margin-left: auto; margin-right: auto;" /></a></div>
        </div>

        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Usage </div>
        <p class="flag_description_text" style="padding-left: 5px;">Factbook photos - obtained from a variety of sources - are in the public domain and are copyright free. <br /> <a href=/about-cia/site-policies/#copy title=Agency Copyright Notice>Agency Copyright Notice</a></p>
      </div>
   </div>

    <div class="item">
      <img src="../photo_gallery/as/images/AS_046.jpg" alt="AUSTRALIA">

<div class="galleria-counter" style="opacity: 1; transition: none 0s ease 0s ;">
<span class="galleria-current">17</span>
/
<span class="galleria-total">23</span>
</div>


      <div class="carousel-caption">
        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Caption </div>
        <div class="photoCaption"><p class="flag_description_text">Orphan Rock in the Blue Mountains as viewed from a cable car.</p></div>
      </div>
      <div class="carousel-photo-info">
        <div class="photoInfo" style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699; line-height: 24px;">
           <div>Dimensions</div><div>File Size</div><div>Download</div>
        </div>
        <div class="photoInfo">
           <div class="flag_description_text">1500X1125</div><div class="flag_description_text">275.8 KB</div><div><a href="../photo_gallery/as/images/large/AS_046_large.jpg"><img class="wfbDownload" src="../graphics/download_image.png" style="padding:0; margin-left: auto; margin-right: auto;" /></a></div>
        </div>

        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Usage </div>
        <p class="flag_description_text" style="padding-left: 5px;">Factbook photos - obtained from a variety of sources - are in the public domain and are copyright free. <br /> <a href=/about-cia/site-policies/#copy title=Agency Copyright Notice>Agency Copyright Notice</a></p>
      </div>
   </div>

    <div class="item">
      <img src="../photo_gallery/as/images/AS_016.jpg" alt="AUSTRALIA">

<div class="galleria-counter" style="opacity: 1; transition: none 0s ease 0s ;">
<span class="galleria-current">18</span>
/
<span class="galleria-total">23</span>
</div>


      <div class="carousel-caption">
        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Caption </div>
        <div class="photoCaption"><p class="flag_description_text">Incipient sunset in the Outback.</p></div>
      </div>
      <div class="carousel-photo-info">
        <div class="photoInfo" style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699; line-height: 24px;">
           <div>Dimensions</div><div>File Size</div><div>Download</div>
        </div>
        <div class="photoInfo">
           <div class="flag_description_text">1500X1200</div><div class="flag_description_text">299.6 KB</div><div><a href="../photo_gallery/as/images/large/AS_016_large.jpg"><img class="wfbDownload" src="../graphics/download_image.png" style="padding:0; margin-left: auto; margin-right: auto;" /></a></div>
        </div>

        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Usage </div>
        <p class="flag_description_text" style="padding-left: 5px;">Factbook photos - obtained from a variety of sources - are in the public domain and are copyright free. <br /> <a href=/about-cia/site-policies/#copy title=Agency Copyright Notice>Agency Copyright Notice</a></p>
      </div>
   </div>

    <div class="item">
      <img src="../photo_gallery/as/images/AS_041.jpg" alt="AUSTRALIA">

<div class="galleria-counter" style="opacity: 1; transition: none 0s ease 0s ;">
<span class="galleria-current">19</span>
/
<span class="galleria-total">23</span>
</div>


      <div class="carousel-caption">
        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Caption </div>
        <div class="photoCaption"><p class="flag_description_text">Some of the granite boulders at Devils Marbles Conservation Reserve near Wauchope, in the Northern Territory. The marbles were formed through various geological processes including chemical and mechanical weathering.</p></div>
      </div>
      <div class="carousel-photo-info">
        <div class="photoInfo" style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699; line-height: 24px;">
           <div>Dimensions</div><div>File Size</div><div>Download</div>
        </div>
        <div class="photoInfo">
           <div class="flag_description_text">1500X1200</div><div class="flag_description_text">286.5 KB</div><div><a href="../photo_gallery/as/images/large/AS_041_large.jpg"><img class="wfbDownload" src="../graphics/download_image.png" style="padding:0; margin-left: auto; margin-right: auto;" /></a></div>
        </div>

        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Usage </div>
        <p class="flag_description_text" style="padding-left: 5px;">Factbook photos - obtained from a variety of sources - are in the public domain and are copyright free. <br /> <a href=/about-cia/site-policies/#copy title=Agency Copyright Notice>Agency Copyright Notice</a></p>
      </div>
   </div>

    <div class="item">
      <img src="../photo_gallery/as/images/AS_033.jpg" alt="AUSTRALIA">

<div class="galleria-counter" style="opacity: 1; transition: none 0s ease 0s ;">
<span class="galleria-current">20</span>
/
<span class="galleria-total">23</span>
</div>


      <div class="carousel-caption">
        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Caption </div>
        <div class="photoCaption"><p class="flag_description_text">142 million years ago, an asteroid or comet slammed into what is now the Missionary Plains in the Northern Territory, forming a crater 24 km (15 mi) in diameter. Due to erosion, the crater rim has been reduced to only 5 km (3 mi). Today, like a bull&apos;s eye, the circular ring of hills that defines Gosses Bluff (image center) stands as a stark reminder of the event shown in this high-resolution satellite photo. Image courtesy of USGS.</p></div>
      </div>
      <div class="carousel-photo-info">
        <div class="photoInfo" style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699; line-height: 24px;">
           <div>Dimensions</div><div>File Size</div><div>Download</div>
        </div>
        <div class="photoInfo">
           <div class="flag_description_text">1392X1500</div><div class="flag_description_text">594.4 KB</div><div><a href="../photo_gallery/as/images/large/AS_033_large.jpg"><img class="wfbDownload" src="../graphics/download_image.png" style="padding:0; margin-left: auto; margin-right: auto;" /></a></div>
        </div>

        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Usage </div>
        <p class="flag_description_text" style="padding-left: 5px;">Factbook photos - obtained from a variety of sources - are in the public domain and are copyright free. <br /> <a href=/about-cia/site-policies/#copy title=Agency Copyright Notice>Agency Copyright Notice</a></p>
      </div>
   </div>

    <div class="item">
      <img src="../photo_gallery/as/images/AS_038.jpg" alt="AUSTRALIA">

<div class="galleria-counter" style="opacity: 1; transition: none 0s ease 0s ;">
<span class="galleria-current">21</span>
/
<span class="galleria-total">23</span>
</div>


      <div class="carousel-caption">
        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Caption </div>
        <div class="photoCaption"><p class="flag_description_text">Boab trees along the Plenty Highway in the Outback. These trees store water in their swollen trunks and shed their leaves during the dry season. Indigenous Australians used them as a source of water and food, and utilized the leaves medicinally.</p></div>
      </div>
      <div class="carousel-photo-info">
        <div class="photoInfo" style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699; line-height: 24px;">
           <div>Dimensions</div><div>File Size</div><div>Download</div>
        </div>
        <div class="photoInfo">
           <div class="flag_description_text">1500X1200</div><div class="flag_description_text">373.5 KB</div><div><a href="../photo_gallery/as/images/large/AS_038_large.jpg"><img class="wfbDownload" src="../graphics/download_image.png" style="padding:0; margin-left: auto; margin-right: auto;" /></a></div>
        </div>

        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Usage </div>
        <p class="flag_description_text" style="padding-left: 5px;">Factbook photos - obtained from a variety of sources - are in the public domain and are copyright free. <br /> <a href=/about-cia/site-policies/#copy title=Agency Copyright Notice>Agency Copyright Notice</a></p>
      </div>
   </div>

    <div class="item">
      <img src="../photo_gallery/as/images/AS_036.jpg" alt="AUSTRALIA">

<div class="galleria-counter" style="opacity: 1; transition: none 0s ease 0s ;">
<span class="galleria-current">22</span>
/
<span class="galleria-total">23</span>
</div>


      <div class="carousel-caption">
        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Caption </div>
        <div class="photoCaption"><p class="flag_description_text">Although covered with an intimidating array of spikes, the thorny devil lizards that inhabit the scrub and desert of western Australia are actually quite gentle. Their main diet consists of ants. They can grow up to 20 cm (8 in) and can live up to 20 years.</p></div>
      </div>
      <div class="carousel-photo-info">
        <div class="photoInfo" style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699; line-height: 24px;">
           <div>Dimensions</div><div>File Size</div><div>Download</div>
        </div>
        <div class="photoInfo">
           <div class="flag_description_text">1500X1200</div><div class="flag_description_text">340.1 KB</div><div><a href="../photo_gallery/as/images/large/AS_036_large.jpg"><img class="wfbDownload" src="../graphics/download_image.png" style="padding:0; margin-left: auto; margin-right: auto;" /></a></div>
        </div>

        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Usage </div>
        <p class="flag_description_text" style="padding-left: 5px;">Factbook photos - obtained from a variety of sources - are in the public domain and are copyright free. <br /> <a href=/about-cia/site-policies/#copy title=Agency Copyright Notice>Agency Copyright Notice</a></p>
      </div>
   </div>

    <div class="item">
      <img src="../photo_gallery/as/images/AS_040.jpg" alt="AUSTRALIA">

<div class="galleria-counter" style="opacity: 1; transition: none 0s ease 0s ;">
<span class="galleria-current">23</span>
/
<span class="galleria-total">23</span>
</div>


      <div class="carousel-caption">
        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Caption </div>
        <div class="photoCaption"><p class="flag_description_text">Wonderfully original sign in the Outback.</p></div>
      </div>
      <div class="carousel-photo-info">
        <div class="photoInfo" style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699; line-height: 24px;">
           <div>Dimensions</div><div>File Size</div><div>Download</div>
        </div>
        <div class="photoInfo">
           <div class="flag_description_text">1500X1200</div><div class="flag_description_text">487.3 KB</div><div><a href="../photo_gallery/as/images/large/AS_040_large.jpg"><img class="wfbDownload" src="../graphics/download_image.png" style="padding:0; margin-left: auto; margin-right: auto;" /></a></div>
        </div>

        <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Usage </div>
        <p class="flag_description_text" style="padding-left: 5px;">Factbook photos - obtained from a variety of sources - are in the public domain and are copyright free. <br /> <a href=/about-cia/site-policies/#copy title=Agency Copyright Notice>Agency Copyright Notice</a></p>
      </div>
   </div>

</div> 
  <!-- Indicators -->
  <ol class="carousel-indicators" style="z-index:1 !important">
   <li data-target="#myCarousel" data-slide-to="${depCount}" class="active"></li>
    <li data-target="#myCarousel" data-slide-to="${depCount}"></li>
    <li data-target="#myCarousel" data-slide-to="${depCount}"></li>
    <li data-target="#myCarousel" data-slide-to="${depCount}"></li>
    <li data-target="#myCarousel" data-slide-to="${depCount}"></li>
    <li data-target="#myCarousel" data-slide-to="${depCount}"></li>
    <li data-target="#myCarousel" data-slide-to="${depCount}"></li>
    <li data-target="#myCarousel" data-slide-to="${depCount}"></li>
    <li data-target="#myCarousel" data-slide-to="${depCount}"></li>
    <li data-target="#myCarousel" data-slide-to="${depCount}"></li>
    <li data-target="#myCarousel" data-slide-to="${depCount}"></li>
    <li data-target="#myCarousel" data-slide-to="${depCount}"></li>
    <li data-target="#myCarousel" data-slide-to="${depCount}"></li>
    <li data-target="#myCarousel" data-slide-to="${depCount}"></li>
    <li data-target="#myCarousel" data-slide-to="${depCount}"></li>
    <li data-target="#myCarousel" data-slide-to="${depCount}"></li>
    <li data-target="#myCarousel" data-slide-to="${depCount}"></li>
    <li data-target="#myCarousel" data-slide-to="${depCount}"></li>
    <li data-target="#myCarousel" data-slide-to="${depCount}"></li>
    <li data-target="#myCarousel" data-slide-to="${depCount}"></li>
    <li data-target="#myCarousel" data-slide-to="${depCount}"></li>
    <li data-target="#myCarousel" data-slide-to="${depCount}"></li>
    <li data-target="#myCarousel" data-slide-to="${depCount}"></li>

  </ol>

  <!-- Controls -->
  <a class="left carousel-control" href="#myCarousel" role="button" data-slide="prev">
    <span class="glyphicon glyphicon-chevron-left" aria-hidden="true"></span>
    <span class="sr-only">Previous</span>
  </a>
  <a class="right carousel-control" href="#myCarousel" role="button" data-slide="next">
    <span class="glyphicon glyphicon-chevron-right" aria-hidden="true"></span>
    <span class="sr-only">Next</span>
  </a>
</div>


             </div>
           </div>
                 </div>
              </div>
            </div>
        </div>

        <div class="map-holder" style="background-image:url(../graphics/aus_lgmap_bkgrnd.jpg); background-repeat: repeat-x; background-position: top ;">
              <div class="mapBox"><a data-toggle="modal" href="#cntryMapModal"><img src="../graphics/maps/as-map.gif"></a></div>
        </div>
        <div class="modal fade" id="cntryMapModal" role="dialog">
           <div class="wfb-modal-dialog">
              <div class="modal-content"  >
                 <div class="wfb-modal-header" style="border-radius: 4px; font-family: Verdana,Arial,sans-serif; font-size: 14px !important; font-weight: bold;  padding: 0.4em 16px 0.4em 1em; background: #cccccc url("..images/ui-bg_highlight-soft_75_cccccc_1x100.png") repeat-x scroll 50% 50%;" >
                    <span style="font-size: 14px !important; margin: 0.1em 16px 0.1em 0;" class="modal-title wfb-title">The World Factbook</span><span style="float: right; margin-top: -4px;">
                    <button type="button" class="close" title="close" data-dismiss="modal">&times;</button></span>
                 </div>
                 <div class="wfb-modal-body">
                    <div class="region1 geos_title aus_dark">Australia-Oceania   <strong>::</strong><span class="region_name1 countryName ">AUSTRALIA</span>
                       <a class="printVersion" href="../graphics/maps/as-map.gif"><img src="../graphics/print.gif" style="padding: 3px;"></a>
                    </div>
           <div class="modal-map-holder" style="background-image:url(../graphics/aus_lgmap_bkgrnd.jpg); background-repeat: repeat-x; background-position: top ;">
             <div class="mapBox">
                  <img class="aus_lgflagborder mapFit1" src="../graphics/maps/as-map.gif">
             </div>
           </div>
                 </div>
              </div>
            </div>
        </div>





     </div>
<ul class="expandcollapse">
<li><h2 class='question aus_med' sectiontitle='Introduction' ccode='as' style='border-bottom: 2px solid white; cursor: pointer;'>Introduction ::  <span class='region'>AUSTRALIA </span></h2></li><li>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2028&term=Background'>Background:</a><a href='../fields/2028.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>Prehistoric settlers arrived on the continent from Southeast Asia at least 40,000 years before the first Europeans began exploration in the 17th century. No formal territorial claims were made until 1770, when Capt. James COOK took possession of the east coast in the name of Great Britain (all of Australia was claimed as British territory in 1829 with the creation of the colony of Western Australia). Six colonies were created in the late 18th and 19th centuries; they federated and became the Commonwealth of Australia in 1901. The new country took advantage of its natural resources to rapidly develop agricultural and manufacturing industries and to make a major contribution to the Allied effort in World Wars I and II.</div>
<div class=category_data>In recent decades, Australia has become an internationally competitive, advanced market economy due in large part to economic reforms adopted in the 1980s and its location in one of the fastest growing regions of the world economy. Long-term concerns include an aging population, pressure on infrastructure, and environmental issues such as floods, droughts, and bushfires. Australia is the driest inhabited continent on earth, making it particularly vulnerable to the challenges of climate change. Australia is home to 10 per cent of the world's biodiversity, and a great number of its flora and fauna exist nowhere else in the world.</div>
</li>
<li><h2 class='question aus_med' sectiontitle='Geography' ccode='as' style='border-bottom: 2px solid white; cursor: pointer;'>Geography ::  <span class='region'>AUSTRALIA </span></h2></li><li>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2144&term=Location'>Location:</a><a href='../fields/2144.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>Oceania, continent between the Indian Ocean and the South Pacific Ocean</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2011&term=Geographic coordinates'>Geographic coordinates:</a><a href='../fields/2011.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>27 00 S, 133 00 E</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2145&term=Map references'>Map references:</a><a href='../fields/2145.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>Oceania</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2147&term=Area'>Area:</a><a href='../fields/2147.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>total: </span><span class=category_data>7,741,220 sq km</span></div>
<div><span class=category>land: </span><span class=category_data>7,682,300 sq km</span></div>
<div><span class=category>water: </span><span class=category_data>58,920 sq km</span></div>
<div><span class=category>note: </span><span class=category_data>includes Lord Howe Island and Macquarie Island</span></div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2147rank.html#as'>6</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2023&term=Area - comparative'>Area - comparative:</a><a href='../fields/2023.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>slightly smaller than the US contiguous 48 states</div>
<div class='disTable areaComp'><span class='category tCell' style='margin-bottom:0px; vertical-align:bottom;'>Area comparison map:</span>
<span class="tCell"><a data-toggle="modal" href="#areaCompModal"><img src="../graphics/areacomparison_icon.jpg" border="0" style="cursor:pointer; border: 0px solid #CCC;"></a></span></div>

        <div class="modal fade" id="areaCompModal" role="dialog">
             <div class="wfb-modal-dialog">
                <div class="modal-content"  >
                   <div class="wfb-modal-header" style="border-radius: 4px; font-family: Verdana,Arial,sans-serif; font-size: 14px !important; font-weight: bold;  padding: 0.4em 16px 0.4em 1em; background: #cccccc url("..images/ui-bg_highlight-soft_75_cccccc_1x100.png") repeat-x scroll 50% 50%;" >
                      <span style="font-size: 14px !important; margin: 0.1em 16px 0.1em 0;" class="modal-title wfb-title">The World Factbook</span><span style="float: right; margin-top: -4px;">
                      <button type="button" class="close" title="close" data-dismiss="modal">&times;</button></span>
                    </div>
                   <div class="wfb-modal-body">
                      <div class="region1 geos_title aus_dark">Australia-Oceania   <strong>::</strong><span class="region_name1 countryName ">AUSTRALIA</span>
                         <a class="printVersion" href="../graphics/areacomparison/AS_area.jpg"><img src="../graphics/print.gif" style="padding: 3px;"></a>
                      </div>
                      <div class="areaComp-modal-map-holder" style="background-image:url(../graphics/aus_lgmap_bkgrnd.jpg); background-repeat: repeat-x; background-position: top ;">
                          <div class="modalFlagBox" style="width: 50%; padding: 33px 0 0 0; margin-left: 30px;">
                             <img class="aus_lgflagborder" src="../graphics/areacomparison/AS_area.jpg">
                          </div>
                          <div class="modalFlagDesc"  >
                              <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Area Comparison </div>
                              <div class="photogallery_captiontext" style="height: 375px;background-color:white;">
                                 <span class="flag_description_text">slightly smaller than the US contiguous 48 states</span>
                              </div>
                           </div>
                      </div>
                   </div>
                </div>
             </div>
     </div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2096&term=Land boundaries'>Land boundaries:</a><a href='../fields/2096.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>0 km</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2060&term=Coastline'>Coastline:</a><a href='../fields/2060.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>25,760 km</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2106&term=Maritime claims'>Maritime claims:</a><a href='../fields/2106.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>territorial sea: </span><span class=category_data>12 nm</span></div>
<div><span class=category>contiguous zone: </span><span class=category_data>24 nm</span></div>
<div><span class=category>exclusive economic zone: </span><span class=category_data>200 nm</span></div>
<div><span class=category>continental shelf: </span><span class=category_data>200 nm or to the edge of the continental margin</span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2059&term=Climate'>Climate:</a><a href='../fields/2059.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>generally arid to semiarid; temperate in south and east; tropical in north</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2125&term=Terrain'>Terrain:</a><a href='../fields/2125.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>mostly low plateau with deserts; fertile plain in southeast</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2020&term=Elevation'>Elevation:</a><a href='../fields/2020.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>mean elevation: </span><span class=category_data>330 m</span></div>
<div><span class=category>elevation extremes: </span><span class=category_data>lowest point: Lake Eyre -15 m</span></div>
<div class=category_data>highest point: Mount Kosciuszko 2,229 m</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2111&term=Natural resources'>Natural resources:</a><a href='../fields/2111.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>bauxite, coal, iron ore, copper, tin, gold, silver, uranium, nickel, tungsten, rare earth elements, mineral sands, lead, zinc, diamonds, natural gas, petroleum</div>
<div><span class=category>note: </span><span class=category_data>Australia is the world's largest net exporter of coal accounting for 29% of global coal exports</span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2097&term=Land use'>Land use:</a><a href='../fields/2097.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>agricultural land: </span><span class=category_data>53.4%</span></div>
<div class=category_data>arable land 6.2%; permanent crops 0.1%; permanent pasture 47.1%</div>
<div><span class=category>forest: </span><span class=category_data>19.3%</span></div>
<div><span class=category>other: </span><span class=category_data>27.3% (2011 est.)</span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2146&term=Irrigated land'>Irrigated land:</a><a href='../fields/2146.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>25,500 sq km (2012)</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2021&term=Natural hazards'>Natural hazards:</a><a href='../fields/2021.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>cyclones along the coast; severe droughts; forest fires</div>
<div><span class=category>volcanism: </span><span class=category_data>volcanic activity on Heard and McDonald Islands</span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2032&term=Environment - current issues'>Environment - current issues:</a><a href='../fields/2032.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>soil erosion from overgrazing, industrial development, urbanization, and poor farming practices; soil salinity rising due to the use of poor quality water; desertification; clearing for agricultural purposes threatens the natural habitat of many unique animal and plant species; the Great Barrier Reef off the northeast coast, the largest coral reef in the world, is threatened by increased shipping and its popularity as a tourist site; limited natural freshwater resources</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2033&term=Environment - international agreements'>Environment - international agreements:</a><a href='../fields/2033.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>party to: </span><span class=category_data>Antarctic-Environmental Protocol, Antarctic-Marine Living Resources, Antarctic Seals, Antarctic Treaty, Biodiversity, Climate Change, Climate Change-Kyoto Protocol, Desertification, Endangered Species, Environmental Modification, Hazardous Wastes, Law of the Sea, Marine Dumping, Marine Life Conservation, Ozone Layer Protection, Ship Pollution, Tropical Timber 83, Tropical Timber 94, Wetlands, Whaling</span></div>
<div><span class=category>signed, but not ratified: </span><span class=category_data>none of the selected agreements</span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2113&term=Geography - note'>Geography - note:</a><a href='../fields/2113.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>world's smallest continent but sixth-largest country; the largest country in Oceania, the largest country entirely in the Southern Hemisphere, and the largest country without land borders; the only continent without glaciers; population concentrated along the eastern and southeastern coasts; the invigorating sea breeze known as the "Fremantle Doctor" affects the city of Perth on the west coast and is one of the most consistent winds in the world</div>
</li>
<li><h2 class='question aus_med' sectiontitle='People and Society' ccode='as' style='border-bottom: 2px solid white; cursor: pointer;'>People and Society ::  <span class='region'>AUSTRALIA </span></h2></li><li>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2119&term=Population'>Population:</a><a href='../fields/2119.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>22,992,654 (July 2016 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2119rank.html#as'>56</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2110&term=Nationality'>Nationality:</a><a href='../fields/2110.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>noun: </span><span class=category_data>Australian(s)</span></div>
<div><span class=category>adjective: </span><span class=category_data>Australian</span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2075&term=Ethnic groups'>Ethnic groups:</a><a href='../fields/2075.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>English 25.9%, Australian 25.4%, Irish 7.5%, Scottish 6.4%, Italian 3.3%, German 3.2%, Chinese 3.1%, Indian 1.4%, Greek 1.4%, Dutch 1.2%, other 15.8% (includes Australian aboriginal .5%), unspecified 5.4%</div>
<div><span class=category>note: </span><span class=category_data>data represents self-identified ancestry, over a third of respondents reported two ancestries (2011 est.)</span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2098&term=Languages'>Languages:</a><a href='../fields/2098.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>English 76.8%, Mandarin 1.6%, Italian 1.4%, Arabic 1.3%, Greek 1.2%, Cantonese 1.2%, Vietnamese 1.1%, other 10.4%, unspecified 5% (2011 est.)</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2122&term=Religions'>Religions:</a><a href='../fields/2122.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>Protestant 30.1% (Anglican 17.1%, Uniting Church 5.0%, Presbyterian and Reformed 2.8%, Baptist, 1.6%, Lutheran 1.2%, Pentecostal 1.1%, other Protestant 1.3%), Catholic 25.3% (Roman Catholic 25.1%, other Catholic 0.2%), other Christian 2.9%, Orthodox 2.8%, Buddhist 2.5%, Muslim 2.2%, Hindu 1.3%, other 1.3%, none 22.3%, unspecified 9.3% (2011 est.)</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2010&term=Age structure'>Age structure:</a><a href='../fields/2010.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>0-14 years: </span><span class=category_data>17.84% (male 2,105,433/female 1,997,433)</span></div>
<div><span class=category>15-24 years: </span><span class=category_data>12.96% (male 1,528,993/female 1,451,340)</span></div>
<div><span class=category>25-54 years: </span><span class=category_data>41.55% (male 4,862,591/female 4,691,975)</span></div>
<div><span class=category>55-64 years: </span><span class=category_data>11.82% (male 1,347,780/female 1,369,501)</span></div>
<div><span class=category>65 years and over: </span><span class=category_data>15.82% (male 1,684,339/female 1,953,269) (2016 est.)</span></div>
<div class='disTable popPyramid'><span class='category tCell' style='margin-bottom:0px; vertical-align:bottom;'>population pyramid:</span>
<span class="tCell"><a data-toggle="modal" href="#popPyramidModal"><img title="" src="../graphics/poppyramid_icon.jpg" style="cursor:pointer; border: 0px solid #CCC;"></span></a></div>

        <div class="modal fade" id="popPyramidModal" role="dialog">
             <div class="wfb-modal-dialog">
                <div class="modal-content"  >
                   <div class="wfb-modal-header" style="border-radius: 4px; font-family: Verdana,Arial,sans-serif; font-size: 14px !important; font-weight: bold;  padding: 0.4em 16px 0.4em 1em; background: #cccccc url("..images/ui-bg_highlight-soft_75_cccccc_1x100.png") repeat-x scroll 50% 50%;" >
                      <span style="font-size: 14px !important; margin: 0.1em 16px 0.1em 0;" class="modal-title wfb-title">The World Factbook</span><span style="float: right; margin-top: -4px;">
                      <button type="button" class="close" title="close" data-dismiss="modal">&times;</button></span>
                    </div>
                   <div class="wfb-modal-body">

                      <div class="region1 geos_title aus_dark">Australia-Oceania   <strong>::</strong><span class="region_name1 countryName ">AUSTRALIA</span>

                         <a class="printVersion" href="../graphics/population/AS_popgraph 2016.bmp"><img src="../graphics/print.gif" style="padding: 3px;"></a>
                      </div>
                      <div class="modal-map-holder" style="background-image:url(../graphics/aus_lgmap_bkgrnd.jpg); background-repeat: repeat-x; background-position: top ;">
                          <div class="modalFlagBox" style="width: 50%; padding: 33px 0 0 0; margin-left: 30px;">
                             <img class="aus_lgflagborder" src="../graphics/population/AS_popgraph 2016.bmp">
                          </div>
                          <div class="modalFlagDesc"  >
                              <div style="font-size:11px; padding-left:5px; font-weight:bold; border:1px solid #FFFFFF; background-image: url(../graphics/aus_medium.jpg); color: #006699;text-align: left; line-height: 24px;">Population Pyramid </div>
                              <div class="photogallery_captiontext" style="height: 375px;background-color:white;">
                                 <span class="flag_description_text">A population pyramid illustrates the age and sex structure of a country's population and may provide insights about political and social stability, as well as economic development. The population is distributed along the horizontal axis, with males shown on the left and females on the right. The male and female populations are broken down into 5-year age groups represented as horizontal bars along the vertical axis, with the youngest age groups at the bottom and the oldest at the top. The shape of the population pyramid gradually evolves over time based on fertility, mortality, and international migration trends.<br /><br />For additional information, please see the entry for Population pyramid on the Definitions and Notes page under the References tab.</span>
                              </div>
                           </div>
                      </div>
                   </div>
                </div>
             </div>
     </div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2261&term=Dependency ratios'>Dependency ratios:</a><a href='../fields/2261.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>total dependency ratio: </span><span class=category_data>50.9%</span></div>
<div><span class=category>youth dependency ratio: </span><span class=category_data>28.2%</span></div>
<div><span class=category>elderly dependency ratio: </span><span class=category_data>22.7%</span></div>
<div><span class=category>potential support ratio: </span><span class=category_data>4.4% (2015 est.)</span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2177&term=Median age'>Median age:</a><a href='../fields/2177.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>total: </span><span class=category_data>38.6 years</span></div>
<div><span class=category>male: </span><span class=category_data>37.8 years</span></div>
<div><span class=category>female: </span><span class=category_data>39.4 years (2016 est.)</span></div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2177rank.html#as'>58</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2002&term=Population growth rate'>Population growth rate:</a><a href='../fields/2002.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>1.05% (2016 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2002rank.html#as'>113</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2054&term=Birth rate'>Birth rate:</a><a href='../fields/2054.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>12.1 births/1,000 population (2016 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2054rank.html#as'>163</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2066&term=Death rate'>Death rate:</a><a href='../fields/2066.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>7.2 deaths/1,000 population (2016 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2066rank.html#as'>124</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2112&term=Net migration rate'>Net migration rate:</a><a href='../fields/2112.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>5.6 migrant(s)/1,000 population (2016 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2112rank.html#as'>22</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2212&term=Urbanization'>Urbanization:</a><a href='../fields/2212.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>urban population: </span><span class=category_data>89.4% of total population (2015)</span></div>
<div><span class=category>rate of urbanization: </span><span class=category_data>1.47% annual rate of change (2010-15 est.)</span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2219&term=Major urban areas - population'>Major urban areas - population:</a><a href='../fields/2219.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>Sydney 4.505 million; Melbourne 4.203 million; Brisbane 2.202 million; Perth 1.861 million; Adelaide 1.256 million; CANBERRA (capital) 423,000 (2015)</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2018&term=Sex ratio'>Sex ratio:</a><a href='../fields/2018.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>at birth: </span><span class=category_data>1.06 male(s)/female</span></div>
<div><span class=category>0-14 years: </span><span class=category_data>1.05 male(s)/female</span></div>
<div><span class=category>15-24 years: </span><span class=category_data>1.05 male(s)/female</span></div>
<div><span class=category>25-54 years: </span><span class=category_data>1.04 male(s)/female</span></div>
<div><span class=category>55-64 years: </span><span class=category_data>0.98 male(s)/female</span></div>
<div><span class=category>65 years and over: </span><span class=category_data>0.86 male(s)/female</span></div>
<div><span class=category>total population: </span><span class=category_data>1.01 male(s)/female (2016 est.)</span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2256&term=Mother's mean age at first birth'>Mother's mean age at first birth:</a><a href='../fields/2256.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>30.5 (2006 est.)</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2223&term=Maternal mortality rate'>Maternal mortality rate:</a><a href='../fields/2223.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>6 deaths/100,000 live births (2015 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2223rank.html#as'>164</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2091&term=Infant mortality rate'>Infant mortality rate:</a><a href='../fields/2091.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>total: </span><span class=category_data>4.3 deaths/1,000 live births</span></div>
<div><span class=category>male: </span><span class=category_data>4.6 deaths/1,000 live births</span></div>
<div><span class=category>female: </span><span class=category_data>4 deaths/1,000 live births (2016 est.)</span></div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2091rank.html#as'>188</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2102&term=Life expectancy at birth'>Life expectancy at birth:</a><a href='../fields/2102.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>total population: </span><span class=category_data>82.2 years</span></div>
<div><span class=category>male: </span><span class=category_data>79.8 years</span></div>
<div><span class=category>female: </span><span class=category_data>84.8 years (2016 est.)</span></div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2102rank.html#as'>15</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2127&term=Total fertility rate'>Total fertility rate:</a><a href='../fields/2127.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>1.77 children born/woman (2016 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2127rank.html#as'>158</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2258&term=Contraceptive prevalence rate'>Contraceptive prevalence rate:</a><a href='../fields/2258.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>72.3%</div>
<div><span class=category>note: </span><span class=category_data>percent of women aged 18-44 (2005)</span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2225&term=Health expenditures'>Health expenditures:</a><a href='../fields/2225.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>9.4% of GDP (2014)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2225rank.html#as'>37</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2226&term=Physicians density'>Physicians density:</a><a href='../fields/2226.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>3.27 physicians/1,000 population (2011)</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2227&term=Hospital bed density'>Hospital bed density:</a><a href='../fields/2227.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>3.9 beds/1,000 population (2010)</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2216&term=Drinking water source'>Drinking water source:</a><a href='../fields/2216.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>improved: </span><span class=category_data></span></div>
<div class=category_data>urban: 100% of population</div>
<div class=category_data>rural: 100% of population</div>
<div class=category_data>total: 100% of population</div>
<div><span class=category>unimproved: </span><span class=category_data></span></div>
<div class=category_data>urban: 0% of population</div>
<div class=category_data>rural: 0% of population</div>
<div class=category_data>total: 0% of population (2015 est.)</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2217&term=Sanitation facility access'>Sanitation facility access:</a><a href='../fields/2217.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>improved: </span><span class=category_data></span></div>
<div class=category_data>urban: 100% of population</div>
<div class=category_data>rural: 100% of population</div>
<div class=category_data>total: 100% of population</div>
<div><span class=category>unimproved: </span><span class=category_data></span></div>
<div class=category_data>urban: 0% of population</div>
<div class=category_data>rural: 0% of population</div>
<div class=category_data>total: 0% of population (2015 est.)</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2155&term=HIV/AIDS - adult prevalence rate'>HIV/AIDS - adult prevalence rate:</a><a href='../fields/2155.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>0.16% (2015 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2155rank.html#as'>100</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2156&term=HIV/AIDS - people living with HIV/AIDS'>HIV/AIDS - people living with HIV/AIDS:</a><a href='../fields/2156.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>26,900 (2015 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2156rank.html#as'>74</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2157&term=HIV/AIDS - deaths'>HIV/AIDS - deaths:</a><a href='../fields/2157.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>200 (2015 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2157rank.html#as'>105</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2228&term=Obesity - adult prevalence rate'>Obesity - adult prevalence rate:</a><a href='../fields/2228.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>29.9% (2014)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2228rank.html#as'>44</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2224&term=Children under the age of 5 years underweight'>Children under the age of 5 years underweight:</a><a href='../fields/2224.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>0.2% (2007)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2224rank.html#as'>138</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2206&term=Education expenditures'>Education expenditures:</a><a href='../fields/2206.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>5.3% of GDP (2013)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2206rank.html#as'>56</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2205&term=School life expectancy (primary to tertiary education)'>School life expectancy (primary to tertiary education):</a><a href='../fields/2205.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>total: </span><span class=category_data>20 years</span></div>
<div><span class=category>male: </span><span class=category_data>20 years</span></div>
<div><span class=category>female: </span><span class=category_data>21 years (2013)</span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2229&term=Unemployment, youth ages 15-24'>Unemployment, youth ages 15-24:</a><a href='../fields/2229.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>total: </span><span class=category_data>13.3%</span></div>
<div><span class=category>male: </span><span class=category_data>14.1%</span></div>
<div><span class=category>female: </span><span class=category_data>12.5% (2014 est.)</span></div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2229rank.html#as'>93</a></span></div>
</li>
<li><h2 class='question aus_med' sectiontitle='Government' ccode='as' style='border-bottom: 2px solid white; cursor: pointer;'>Government ::  <span class='region'>AUSTRALIA </span></h2></li><li>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2142&term=Country name'>Country name:</a><a href='../fields/2142.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>conventional long form: </span><span class=category_data>Commonwealth of Australia</span></div>
<div><span class=category>conventional short form: </span><span class=category_data>Australia</span></div>
<div><span class=category>etymology: </span><span class=category_data>the name Australia derives from the Latin "australis" meaning "southern"; the Australian landmass was long referred to as "Terra Australis" or the Southern Land</span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2128&term=Government type'>Government type:</a><a href='../fields/2128.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>parliamentary democracy (Federal Parliament) under a constitutional monarchy; a Commonwealth realm</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2057&term=Capital'>Capital:</a><a href='../fields/2057.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>name: </span><span class=category_data>Canberra</span></div>
<div><span class=category>geographic coordinates: </span><span class=category_data>35 16 S, 149 08 E</span></div>
<div><span class=category>time difference: </span><span class=category_data>UTC+10 (15 hours ahead of Washington, DC, during Standard Time)</span></div>
<div><span class=category>daylight saving time: </span><span class=category_data>+1hr, begins first Sunday in October; ends first Sunday in April</span></div>
<div><span class=category>note: </span><span class=category_data>Australia has three time zones</span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2051&term=Administrative divisions'>Administrative divisions:</a><a href='../fields/2051.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>6 states and 2 territories*; Australian Capital Territory*, New South Wales, Northern Territory*, Queensland, South Australia, Tasmania, Victoria, Western Australia</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2068&term=Dependent areas'>Dependent areas:</a><a href='../fields/2068.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>Ashmore and Cartier Islands, Christmas Island, Cocos (Keeling) Islands, Coral Sea Islands, Heard Island and McDonald Islands, Norfolk Island</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2088&term=Independence'>Independence:</a><a href='../fields/2088.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>1 January 1901 (from the federation of UK colonies)</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2109&term=National holiday'>National holiday:</a><a href='../fields/2109.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>Australia Day (commemorates the arrival of the First Fleet of Australian settlers), 26 January (1788); ANZAC Day (commemorates the anniversary of the landing of troops of the Australian and New Zealand Army Corps during World War I at Gallipoli, Turkey), 25 April (1915)</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2063&term=Constitution'>Constitution:</a><a href='../fields/2063.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>9 July 1900; effective 1 January 1901; amended several times, last in 1977; note - a referendum to amend the constitution to reflect the Aboriginal and Torres Strait Islander Peoples Recognition Act 2013 is planned for early 2017 (2016)</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2100&term=Legal system'>Legal system:</a><a href='../fields/2100.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>common law system based on the English model</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2220&term=International law organization participation'>International law organization participation:</a><a href='../fields/2220.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>accepts compulsory ICJ jurisdiction with reservations; accepts ICCt jurisdiction</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2263&term=Citizenship'>Citizenship:</a><a href='../fields/2263.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>citizenship by birth: </span><span class=category_data>no</span></div>
<div><span class=category>citizenship by descent only: </span><span class=category_data>at least one parent must be a citizen or permanent resident of Australia</span></div>
<div><span class=category>dual citizenship recognized: </span><span class=category_data>yes</span></div>
<div><span class=category>residency requirement for naturalization: </span><span class=category_data>4 years</span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2123&term=Suffrage'>Suffrage:</a><a href='../fields/2123.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>18 years of age; universal and compulsory</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2077&term=Executive branch'>Executive branch:</a><a href='../fields/2077.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>chief of state: </span><span class=category_data>Queen of Australia ELIZABETH II (since 6 February 1952); represented by Governor General Sir Peter COSGROVE (since 28 March 2014)</span></div>
<div><span class=category>head of government: </span><span class=category_data>Prime Minister Malcolm TURNBULL (since 15 September 2015); Deputy Prime Minister Barnaby JOYCE (since 18 February 2016)</span></div>
<div><span class=category>cabinet: </span><span class=category_data>Cabinet nominated by the prime minister from among members of Parliament and sworn in by the governor general</span></div>
<div><span class=category>elections/appointments: </span><span class=category_data>the monarchy is hereditary; governor general appointed by the monarch on the recommendation of the prime minister; following legislative elections, the leader of the majority party or majority coalition is sworn in as prime minister by the governor general</span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2101&term=Legislative branch'>Legislative branch:</a><a href='../fields/2101.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>description: </span><span class=category_data>bicameral Federal Parliament consists of the Senate (76 seats; 12 members from each of the 6 states and 2 each from the 2 mainland territories; members directly elected in multi-seat constituencies by proportional representation vote; members serve 6-year terms with one-half of state membership renewed every 3 years and territory membership renewed every 3 years) and the House of Representatives (150 seats; members directly elected in single-seat constituencies by majority preferential vote; members serve terms of up to 3 years)</span></div>
<div><span class=category>elections: </span><span class=category_data>Senate - last held on 2 July 2016; House of Representatives - last held on 2 July 2016; this election represents a rare double dissolution where all 226 seats in both the Senate and House of Representatives are up for reelection</span></div>
<div><span class=category>election results: </span><span class=category_data>Senate - percent of vote by party NA - awaiting final results; seats by party NA - awaiting final results; House of Representatives - percent of vote by party Liberal/National Coalition 42.14%, ALP 34.91%, The Greens 9.93%, Katter's Australian Party 0.55%, Nick Xenophon Team 1.86%, independents 2.85%; seats by party Liberal/National Coalition 77, ALP 68, The Greens 1, Katter's Australian Party 1, Nick Xenophon Team 1, independents 2</span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2094&term=Judicial branch'>Judicial branch:</a><a href='../fields/2094.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>highest court(s): </span><span class=category_data>High Court of Australia (consists of 7 justices, including the chief justice); note - each of the 6 states, 2 territories, and Norfolk Island has a Supreme Court; the High Court is the final appellate court beyond the state and territory supreme courts</span></div>
<div><span class=category>judge selection and term of office: </span><span class=category_data>justices appointed by the governor-general in council for life with mandatory retirement at age 70</span></div>
<div><span class=category>subordinate courts: </span><span class=category_data>subordinate courts: subordinate courts at the federal level: Federal Court; Federal Magistrates' Courts of Australia; Family Court; subordinate courts at the state and territory level: Local Court - New South Wales; Magistrates' Courts &ndash; Victoria, Queensland, South Australia, Western Australia, Tasmania, Northern Territory, Australian Capital Territory; District Courts &ndash; New South Wales, Queensland, South Australia, Western Australia; County Court &ndash; Victoria; Family Court &ndash; Western Australia; Court of Petty Sessions &ndash; Norfolk Island</span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2118&term=Political parties and leaders'>Political parties and leaders:</a><a href='../fields/2118.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>Australian Greens Party [Richard DI NATALE]</div>
<div class=category_data>Australian Labor Party [Bill SHORTEN]</div>
<div class=category_data>Country Liberal Party or CLP [Gary HIGGINS]</div>
<div class=category_data>Family First Party [Bob DAY]</div>
<div class=category_data>Katter's Australian Party [Bob KATTER]</div>
<div class=category_data>Liberal National Party of Queensland or LNP [Timothy NICHOLLS]</div>
<div class=category_data>Liberal Party [Malcolm TURNBULL]</div>
<div class=category_data>National Party of Australia [Barnaby JOYCE]</div>
<div class=category_data>Palmer United Party or PUP [Clive PALMER]</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2115&term=Political pressure groups and leaders'>Political pressure groups and leaders:</a><a href='../fields/2115.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>other: </span><span class=category_data>business groups, environmental groups, social groups, trade unions</span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2107&term=International organization participation'>International organization participation:</a><a href='../fields/2107.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>ADB, ANZUS, APEC, ARF, ASEAN (dialogue partner), Australia Group, BIS, C, CD, CP, EAS, EBRD, EITI (implementing country), FAO, FATF, G-20, IAEA, IBRD, ICAO, ICC (national committees), ICCt, ICRM, IDA, IEA, IFC, IFRCS, IHO, ILO, IMF, IMO, IMSO, Interpol, IOC, IOM, IPU, ISO, ITSO, ITU, ITUC (NGOs), MIGA, NEA, NSG, OECD, OPCW, OSCE (partner), Pacific Alliance (observer), Paris Club, PCA, PIF, SAARC (observer), SICA (observer), Sparteca, SPC, UN, UN Security Council (temporary), UNCTAD, UNESCO, UNHCR, UNMISS, UNMIT, UNRWA, UNTSO, UNWTO, UPU, WCO, WFTU (NGOs), WHO, WIPO, WMO, WTO, ZC</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2149&term=Diplomatic representation in the US'>Diplomatic representation in the US:</a><a href='../fields/2149.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>chief of mission: </span><span class=category_data>Ambassador Joseph Benedict HOCKEY (since 28 January 2016)</span></div>
<div><span class=category>chancery: </span><span class=category_data>1601 Massachusetts Avenue NW, Washington, DC 20036</span></div>
<div><span class=category>telephone: </span><span class=category_data>[1] (202) 797-3000</span></div>
<div><span class=category>FAX: </span><span class=category_data>[1] (202) 797-3168</span></div>
<div><span class=category>consulate(s) general: </span><span class=category_data>Atlanta, Chicago, Honolulu, Houston, Los Angeles, New York, San Francisco</span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2007&term=Diplomatic representation from the US'>Diplomatic representation from the US:</a><a href='../fields/2007.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>chief of mission: </span><span class=category_data>Ambassador Morrell John BERRY (since 25 September 2013)</span></div>
<div><span class=category>embassy: </span><span class=category_data>Moonah Place, Yarralumla, Canberra, Australian Capital Territory 2600</span></div>
<div><span class=category>mailing address: </span><span class=category_data>APO AP 96549</span></div>
<div><span class=category>telephone: </span><span class=category_data>[61] (02) 6214-5600</span></div>
<div><span class=category>FAX: </span><span class=category_data>[61] (02) 6214-5970</span></div>
<div><span class=category>consulate(s) general: </span><span class=category_data>Melbourne, Perth, Sydney</span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2081&term=Flag description'>Flag description:</a><a href='../fields/2081.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>blue with the flag of the UK in the upper hoist-side quadrant and a large seven-pointed star in the lower hoist-side quadrant known as the Commonwealth or Federation Star, representing the federation of the colonies of Australia in 1901; the star depicts one point for each of the six original states and one representing all of Australia's internal and external territories; on the fly half is a representation of the Southern Cross constellation in white with one small, five-pointed star and four larger, seven-pointed stars</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2230&term=National symbol(s)'>National symbol(s):</a><a href='../fields/2230.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>Southern Cross constellation (composed of five stars: four large seven-pointed stars, one small five-pointed star), kangaroo, emu; national colors: green, gold</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2218&term=National anthem'>National anthem:</a><a href='../fields/2218.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>name: </span><span class=category_data>"Advance Australia Fair"</span></div>
<div><span class=category>lyrics/music: </span><span class=category_data>Peter Dodds McCORMICK</span></div>
<div><span class=category>note: </span><span class=category_data>adopted 1984; although originally written in the late 19th century, the anthem was not used for all official occasions until 1984; as a Commonwealth country, in addition to the national anthem, "God Save the Queen" is also played at Royal functions (see United Kingdom)</span></div>
<div class='wrap'><div class='audio-player'><audio id='audio-player-1' class='my-audio-player' src='../anthems/AS.mp3' type='audio/mp3' controls='controls'></audio></div></div>
</li>
<li><h2 class='question aus_med' sectiontitle='Economy' ccode='as' style='border-bottom: 2px solid white; cursor: pointer;'>Economy ::  <span class='region'>AUSTRALIA </span></h2></li><li>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2116&term=Economy - overview'>Economy - overview:</a><a href='../fields/2116.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>Following two decades of continuous growth, low unemployment, contained inflation, very low public debt, and a strong and stable financial system, Australia enters 2016 facing a range of growth constraints, principally driven by a sharp fall in global prices of key export commodities. Demand for resources and energy from Asia and especially China has stalled and sharp drops in current prices have impacted growth.</div>
<div class=category_data></div>
<div class=category_data>The services sector is the largest part of the Australian economy, accounting for about 70% of GDP and 75% of jobs. Australia was comparatively unaffected by the global financial crisis as the banking system has remained strong and inflation is under control.</div>
<div class=category_data></div>
<div class=category_data>Australia benefited from a dramatic surge in its terms of trade in recent years, although this trend has reversed due to falling global commodity prices. Australia is a significant exporter of natural resources, energy, and food. Australia's abundant and diverse natural resources attract high levels of foreign investment and include extensive reserves of coal, iron, copper, gold, natural gas, uranium, and renewable energy sources. A series of major investments, such as the US$40 billion Gorgon Liquid Natural Gas project, will significantly expand the resources sector.</div>
<div class=category_data></div>
<div class=category_data>Australia is an open market with minimal restrictions on imports of goods and services. The process of opening up has increased productivity, stimulated growth, and made the economy more flexible and dynamic. Australia plays an active role in the World Trade Organization, APEC, the G20, and other trade forums. Australia&rsquo;s free trade agreement (FTA) with China entered into force in 2015, adding to existing FTAs with the Republic of Korea, Japan, Chile, Malaysia, New Zealand, Singapore, Thailand, and the US, and a regional FTA with ASEAN and New Zealand. Australia continues to negotiate bilateral agreements with India and Indonesia, as well as larger agreements with its Pacific neighbors and the Gulf Cooperation Council countries, and an Asia-wide Regional Comprehensive Economic Partnership that includes the ten ASEAN countries and China, Japan, Korea, New Zealand and India. Australia is also working on the Trans-Pacific Partnership Agreement with Brunei, Canada, Chile, Japan, Malaysia, Mexico, New Zealand, Peru, Singapore, the US, and Vietnam.</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2001&term=GDP (purchasing power parity)'>GDP (purchasing power parity):</a><a href='../fields/2001.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>$1.189 trillion (2016 est.)</div>
<div class=category_data>$1.156 trillion (2015 est.)</div>
<div class=category_data>$1.128 trillion (2014 est.)</div>
<div><span class=category>note: </span><span class=category_data>data are in 2016 dollars</span></div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2001rank.html#as'>20</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2195&term=GDP (official exchange rate)'>GDP (official exchange rate):</a><a href='../fields/2195.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>$1.257 trillion (2015 est.)</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2003&term=GDP - real growth rate'>GDP - real growth rate:</a><a href='../fields/2003.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>2.9% (2016 est.)</div>
<div class=category_data>2.4% (2015 est.)</div>
<div class=category_data>2.7% (2014 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2003rank.html#as'>102</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2004&term=GDP - per capita (PPP)'>GDP - per capita (PPP):</a><a href='../fields/2004.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>$48,800 (2016 est.)</div>
<div class=category_data>$48,300 (2015 est.)</div>
<div class=category_data>$47,800 (2014 est.)</div>
<div><span class=category>note: </span><span class=category_data>data are in 2016 dollars</span></div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2004rank.html#as'>26</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2260&term=Gross national saving'>Gross national saving:</a><a href='../fields/2260.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>21.7% of GDP (2016 est.)</div>
<div class=category_data>22.1% of GDP (2015 est.)</div>
<div class=category_data>23.7% of GDP (2014 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2260rank.html#as'>72</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2259&term=GDP - composition, by end use'>GDP - composition, by end use:</a><a href='../fields/2259.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>household consumption: </span><span class=category_data>58.5%</span></div>
<div><span class=category>government consumption: </span><span class=category_data>18.7%</span></div>
<div><span class=category>investment in fixed capital: </span><span class=category_data>24.3%</span></div>
<div><span class=category>investment in inventories: </span><span class=category_data>0%</span></div>
<div><span class=category>exports of goods and services: </span><span class=category_data>19.4%</span></div>
<div><span class=category>imports of goods and services: </span><span class=category_data>-20.9% (2016 est.)</span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2012&term=GDP - composition, by sector of origin'>GDP - composition, by sector of origin:</a><a href='../fields/2012.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>agriculture: </span><span class=category_data>3.6%</span></div>
<div><span class=category>industry: </span><span class=category_data>28.2%</span></div>
<div><span class=category>services: </span><span class=category_data>68.2% (2016 est.)</span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2052&term=Agriculture - products'>Agriculture - products:</a><a href='../fields/2052.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>wheat, barley, sugarcane, fruits; cattle, sheep, poultry</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2090&term=Industries'>Industries:</a><a href='../fields/2090.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>mining, industrial and transportation equipment, food processing, chemicals, steel</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2089&term=Industrial production growth rate'>Industrial production growth rate:</a><a href='../fields/2089.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>2% (2016 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2089rank.html#as'>104</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2095&term=Labor force'>Labor force:</a><a href='../fields/2095.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>12.63 million (2016 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2095rank.html#as'>45</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2048&term=Labor force - by occupation'>Labor force - by occupation:</a><a href='../fields/2048.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>agriculture: </span><span class=category_data>3.6%</span></div>
<div><span class=category>industry: </span><span class=category_data>21.1%</span></div>
<div><span class=category>services: </span><span class=category_data>75.3% (2009 est.)</span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2129&term=Unemployment rate'>Unemployment rate:</a><a href='../fields/2129.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>5.8% (2016 est.)</div>
<div class=category_data>6.1% (2015 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2129rank.html#as'>62</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2046&term=Population below poverty line'>Population below poverty line:</a><a href='../fields/2046.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>NA%</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2047&term=Household income or consumption by percentage share'>Household income or consumption by percentage share:</a><a href='../fields/2047.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>lowest 10%: </span><span class=category_data>2%</span></div>
<div><span class=category>highest 10%: </span><span class=category_data>25.4% (1994)</span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2172&term=Distribution of family income - Gini index'>Distribution of family income - Gini index:</a><a href='../fields/2172.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>30.3 (2008)</div>
<div class=category_data>35.2 (1994)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2172rank.html#as'>120</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2056&term=Budget'>Budget:</a><a href='../fields/2056.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>revenues: </span><span class=category_data>$420.5 billion</span></div>
<div><span class=category>expenditures: </span><span class=category_data>$446.4 billion (2016 est.)</span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2221&term=Taxes and other revenues'>Taxes and other revenues:</a><a href='../fields/2221.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>33.5% of GDP (2016 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2221rank.html#as'>65</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2222&term=Budget surplus (+) or deficit (-)'>Budget surplus (+) or deficit (-):</a><a href='../fields/2222.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>-2.1% of GDP (2016 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2222rank.html#as'>70</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2186&term=Public debt'>Public debt:</a><a href='../fields/2186.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>46.1% of GDP (2016 est.)</div>
<div class=category_data>44.2% of GDP (2015 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2186rank.html#as'>98</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2080&term=Fiscal year'>Fiscal year:</a><a href='../fields/2080.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>1 July - 30 June</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2092&term=Inflation rate (consumer prices)'>Inflation rate (consumer prices):</a><a href='../fields/2092.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>1.4% (2016 est.)</div>
<div class=category_data>1.5% (2015 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2092rank.html#as'>87</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2207&term=Central bank discount rate'>Central bank discount rate:</a><a href='../fields/2207.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>3% (28 February 2013)</div>
<div class=category_data>4.35% (31 December 2010)</div>
<div><span class=category>note: </span><span class=category_data>this is the Reserve Bank of Australia's "cash rate target," or policy rate</span></div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2207rank.html#as'>106</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2208&term=Commercial bank prime lending rate'>Commercial bank prime lending rate:</a><a href='../fields/2208.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>5.1% (31 December 2016 est.)</div>
<div class=category_data>5.58% (31 December 2015 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2208rank.html#as'>138</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2214&term=Stock of narrow money'>Stock of narrow money:</a><a href='../fields/2214.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>$245.3 billion (31 December 2016 est.)</div>
<div class=category_data>$223.2 billion (31 December 2015 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2214rank.html#as'>18</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2215&term=Stock of broad money'>Stock of broad money:</a><a href='../fields/2215.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>$1.661 trillion (31 December 2013 est.)</div>
<div class=category_data>$1.648 trillion (31 December 2012 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2215rank.html#as'>11</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2211&term=Stock of domestic credit'>Stock of domestic credit:</a><a href='../fields/2211.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>$2.158 trillion (31 December 2016 est.)</div>
<div class=category_data>$1.986 trillion (31 December 2015 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2211rank.html#as'>11</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2200&term=Market value of publicly traded shares'>Market value of publicly traded shares:</a><a href='../fields/2200.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>$1.187 trillion (31 December 2015 est.)</div>
<div class=category_data>$1.289 trillion (31 December 2014 est.)</div>
<div class=category_data>$1.366 trillion (31 December 2013 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2200rank.html#as'>13</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2187&term=Current account balance'>Current account balance:</a><a href='../fields/2187.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>-$43.85 billion (2016 est.)</div>
<div class=category_data>-$57.98 billion (2015 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2187rank.html#as'>194</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2078&term=Exports'>Exports:</a><a href='../fields/2078.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>$184.3 billion (2016 est.)</div>
<div class=category_data>$188.3 billion (2015 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2078rank.html#as'>26</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2049&term=Exports - commodities'>Exports - commodities:</a><a href='../fields/2049.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>coal, iron ore, gold, meat, wool, alumina, wheat, machinery and transport equipment</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2050&term=Exports - partners'>Exports - partners:</a><a href='../fields/2050.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>China 32.2%, Japan 15.9%, South Korea 7.1%, US 5.4%, India 4.2% (2015)</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2087&term=Imports'>Imports:</a><a href='../fields/2087.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>$203.1 billion (2016 est.)</div>
<div class=category_data>$207.7 billion (2015 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2087rank.html#as'>21</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2058&term=Imports - commodities'>Imports - commodities:</a><a href='../fields/2058.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>machinery and transport equipment, computers and office machines, telecommunication equipment and parts; crude oil and petroleum products</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2061&term=Imports - partners'>Imports - partners:</a><a href='../fields/2061.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>China 23%, US 11.2%, Japan 7.4%, South Korea 5.5%, Thailand 5.1%, Germany 4.6% (2015)</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2188&term=Reserves of foreign exchange and gold'>Reserves of foreign exchange and gold:</a><a href='../fields/2188.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>$54.3 billion (31 December 2016 est.)</div>
<div class=category_data>$49.27 billion (31 December 2015 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2188rank.html#as'>36</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2079&term=Debt - external'>Debt - external:</a><a href='../fields/2079.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>$1.692 trillion (31 December 2016 est.)</div>
<div class=category_data>$1.524 trillion (31 December 2015 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2079rank.html#as'>12</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2198&term=Stock of direct foreign investment - at home'>Stock of direct foreign investment - at home:</a><a href='../fields/2198.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>$614.5 billion (31 December 2016 est.)</div>
<div class=category_data>$582.6 billion (31 December 2015 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2198rank.html#as'>16</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2199&term=Stock of direct foreign investment - abroad'>Stock of direct foreign investment - abroad:</a><a href='../fields/2199.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>$441.9 billion (31 December 2016 est.)</div>
<div class=category_data>$437.8 billion (31 December 2015 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2199rank.html#as'>18</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2076&term=Exchange rates'>Exchange rates:</a><a href='../fields/2076.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>Australian dollars (AUD) per US dollar -</div>
<div class=category_data>1.352 (2016 est.)</div>
<div class=category_data>1.3291 (2015 est.)</div>
<div class=category_data>1.3291 (2014 est.)</div>
<div class=category_data>1.1094 (2013 est.)</div>
<div class=category_data>0.97 (2012 est.)</div>
</li>
<li><h2 class='question aus_med' sectiontitle='Energy' ccode='as' style='border-bottom: 2px solid white; cursor: pointer;'>Energy ::  <span class='region'>AUSTRALIA </span></h2></li><li>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2232&term=Electricity - production'>Electricity - production:</a><a href='../fields/2232.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>235 billion kWh (2014 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2232rank.html#as'>21</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2233&term=Electricity - consumption'>Electricity - consumption:</a><a href='../fields/2233.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>224 billion kWh (2014 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2233rank.html#as'>18</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2234&term=Electricity - exports'>Electricity - exports:</a><a href='../fields/2234.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>0 kWh (2013 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2234rank.html#as'>211</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2235&term=Electricity - imports'>Electricity - imports:</a><a href='../fields/2235.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>0 kWh (2013 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2235rank.html#as'>213</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2236&term=Electricity - installed generating capacity'>Electricity - installed generating capacity:</a><a href='../fields/2236.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>67 million kW (2014 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2236rank.html#as'>17</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2237&term=Electricity - from fossil fuels'>Electricity - from fossil fuels:</a><a href='../fields/2237.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>78.5% of total installed capacity (2012 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2237rank.html#as'>95</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2239&term=Electricity - from nuclear fuels'>Electricity - from nuclear fuels:</a><a href='../fields/2239.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>0% of total installed capacity (2012 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2239rank.html#as'>204</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2238&term=Electricity - from hydroelectric plants'>Electricity - from hydroelectric plants:</a><a href='../fields/2238.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>12.7% of total installed capacity (2012 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2238rank.html#as'>107</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2240&term=Electricity - from other renewable sources'>Electricity - from other renewable sources:</a><a href='../fields/2240.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>7.6% of total installed capacity (2012 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2240rank.html#as'>49</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2241&term=Crude oil - production'>Crude oil - production:</a><a href='../fields/2241.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>322,300 bbl/day (2015 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2241rank.html#as'>32</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2242&term=Crude oil - exports'>Crude oil - exports:</a><a href='../fields/2242.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>248,400 bbl/day (2015 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2242rank.html#as'>28</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2243&term=Crude oil - imports'>Crude oil - imports:</a><a href='../fields/2243.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>332,800 bbl/day (2015 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2243rank.html#as'>26</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2244&term=Crude oil - proved reserves'>Crude oil - proved reserves:</a><a href='../fields/2244.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>1.2 billion bbl (1 January 2016 es)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2244rank.html#as'>40</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2245&term=Refined petroleum products - production'>Refined petroleum products - production:</a><a href='../fields/2245.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>481,800 bbl/day (2015 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2245rank.html#as'>37</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2246&term=Refined petroleum products - consumption'>Refined petroleum products - consumption:</a><a href='../fields/2246.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>1.116 million bbl/day (2015 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2246rank.html#as'>21</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2247&term=Refined petroleum products - exports'>Refined petroleum products - exports:</a><a href='../fields/2247.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>42,730 bbl/day (2015 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2247rank.html#as'>60</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2248&term=Refined petroleum products - imports'>Refined petroleum products - imports:</a><a href='../fields/2248.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>546,700 bbl/day (2015 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2248rank.html#as'>15</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2249&term=Natural gas - production'>Natural gas - production:</a><a href='../fields/2249.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>62.64 billion cu m (2014 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2249rank.html#as'>15</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2250&term=Natural gas - consumption'>Natural gas - consumption:</a><a href='../fields/2250.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>38.51 billion cu m (2014 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2250rank.html#as'>25</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2251&term=Natural gas - exports'>Natural gas - exports:</a><a href='../fields/2251.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>31.61 billion cu m (2014 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2251rank.html#as'>12</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2252&term=Natural gas - imports'>Natural gas - imports:</a><a href='../fields/2252.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>6.938 billion cu m (2014 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2252rank.html#as'>30</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2253&term=Natural gas - proved reserves'>Natural gas - proved reserves:</a><a href='../fields/2253.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>860.8 billion cu m (1 January 2016 es)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2253rank.html#as'>26</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2254&term=Carbon dioxide emissions from consumption of energy'>Carbon dioxide emissions from consumption of energy:</a><a href='../fields/2254.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>385 million Mt (2013 est.)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2254rank.html#as'>18</a></span></div>
</li>
<li><h2 class='question aus_med' sectiontitle='Communications' ccode='as' style='border-bottom: 2px solid white; cursor: pointer;'>Communications ::  <span class='region'>AUSTRALIA </span></h2></li><li>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2150&term=Telephones - fixed lines'>Telephones - fixed lines:</a><a href='../fields/2150.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>total subscriptions: </span><span class=category_data>9.08 million</span></div>
<div><span class=category>subscriptions per 100 inhabitants: </span><span class=category_data>40 (July 2015 est.)</span></div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2150rank.html#as'>22</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2151&term=Telephones - mobile cellular'>Telephones - mobile cellular:</a><a href='../fields/2151.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>total: </span><span class=category_data>31.77 million</span></div>
<div><span class=category>subscriptions per 100 inhabitants: </span><span class=category_data>140 (July 2015 est.)</span></div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2151rank.html#as'>39</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2124&term=Telephone system'>Telephone system:</a><a href='../fields/2124.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>general assessment: </span><span class=category_data>excellent domestic and international service</span></div>
<div><span class=category>domestic: </span><span class=category_data>domestic satellite system; significant use of radiotelephone in areas of low population density; rapid growth of mobile telephones</span></div>
<div><span class=category>international: </span><span class=category_data>country code - 61; landing point for the SEA-ME-WE-3 optical telecommunications submarine cable with links to Asia, the Middle East, and Europe; the Southern Cross fiber-optic submarine cable provides links to NZ and the US; satellite earth stations - 10 (2015)</span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2213&term=Broadcast media'>Broadcast media:</a><a href='../fields/2213.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>the Australian Broadcasting Corporation (ABC) runs multiple national and local radio networks and TV stations, as well as Australia Network, a TV service that broadcasts throughout the Asia-Pacific region and is the main public broadcaster; Special Broadc (2008)</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2154&term=Internet country code'>Internet country code:</a><a href='../fields/2154.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>.au</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2153&term=Internet users'>Internet users:</a><a href='../fields/2153.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>total: </span><span class=category_data>19.238 million</span></div>
<div><span class=category>percent of population: </span><span class=category_data>84.6% (July 2015 est.)</span></div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2153rank.html#as'>28</a></span></div>
</li>
<li><h2 class='question aus_med' sectiontitle='Transportation' ccode='as' style='border-bottom: 2px solid white; cursor: pointer;'>Transportation ::  <span class='region'>AUSTRALIA </span></h2></li><li>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2269&term=National air transport system'>National air transport system:</a><a href='../fields/2269.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>number of registered air carriers: </span><span class=category_data>11</span></div>
<div><span class=category>inventory of registered aircraft operated by air carriers: </span><span class=category_data>175</span></div>
<div><span class=category>annual passenger traffic on registered air carriers: </span><span class=category_data>69,294,187</span></div>
<div><span class=category>annual freight traffic on registered air carriers: </span><span class=category_data>1,887,295,820 mt-km (2015)</span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2270&term=Civil aircraft registration country code prefix'>Civil aircraft registration country code prefix:</a><a href='../fields/2270.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>VH (2016)</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2053&term=Airports'>Airports:</a><a href='../fields/2053.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>480 (2013)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2053rank.html#as'>16</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2030&term=Airports - with paved runways'>Airports - with paved runways:</a><a href='../fields/2030.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>total: </span><span class=category_data>349</span></div>
<div><span class=category>over 3,047 m: </span><span class=category_data>11</span></div>
<div><span class=category>2,438 to 3,047 m: </span><span class=category_data>14</span></div>
<div><span class=category>1,524 to 2,437 m: </span><span class=category_data>155</span></div>
<div><span class=category>914 to 1,523 m: </span><span class=category_data>155</span></div>
<div><span class=category>under 914 m: </span><span class=category_data>14 (2013)</span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2031&term=Airports - with unpaved runways'>Airports - with unpaved runways:</a><a href='../fields/2031.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>total: </span><span class=category_data>131</span></div>
<div><span class=category>1,524 to 2,437 m: </span><span class=category_data>16</span></div>
<div><span class=category>914 to 1,523 m: </span><span class=category_data>101</span></div>
<div><span class=category>under 914 m: </span><span class=category_data>14 (2013)</span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2019&term=Heliports'>Heliports:</a><a href='../fields/2019.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>1 (2013)</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2117&term=Pipelines'>Pipelines:</a><a href='../fields/2117.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>condensate/gas 637 km; gas 30,054 km; liquid petroleum gas 240 km; oil 3,609 km; oil/gas/water 110 km; refined products 72 km (2013)</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2121&term=Railways'>Railways:</a><a href='../fields/2121.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>total: </span><span class=category_data>36,967.5 km</span></div>
<div><span class=category>broad gauge: </span><span class=category_data>3,727 km 1.600-m gauge (372 km electrified)</span></div>
<div><span class=category>standard gauge: </span><span class=category_data>18,727 km 1.435-m gauge (650 km electrified)</span></div>
<div><span class=category>narrow gauge: </span><span class=category_data>14,513.5 km 1.067-m gauge (2,075.5 km electrified) (2014)</span></div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2121rank.html#as'>7</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2085&term=Roadways'>Roadways:</a><a href='../fields/2085.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>total: </span><span class=category_data>823,217 km</span></div>
<div><span class=category>paved: </span><span class=category_data>356,343 km</span></div>
<div><span class=category>unpaved: </span><span class=category_data>466,874 km (2011)</span></div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2085rank.html#as'>9</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2093&term=Waterways'>Waterways:</a><a href='../fields/2093.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>2,000 km (mainly used for recreation on Murray and Murray-Darling river systems) (2011)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2093rank.html#as'>42</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2108&term=Merchant marine'>Merchant marine:</a><a href='../fields/2108.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>total: </span><span class=category_data>41</span></div>
<div><span class=category>by type: </span><span class=category_data>bulk carrier 8, cargo 7, liquefied gas 4, passenger 6, passenger/cargo 6, petroleum tanker 5, roll on/roll off 5</span></div>
<div><span class=category>foreign-owned: </span><span class=category_data>17 (Canada 5, Germany 2, Singapore 2, South Africa 1, UK 5, US 2)</span></div>
<div><span class=category>registered in other countries: </span><span class=category_data>25 (Bahamas 1, Dominica 1, Fiji 2, Liberia 1, Netherlands 1, Panama 4, Singapore 12, Tonga 1, UK 1, US 1) (2010)</span></div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2108rank.html#as'>75</a></span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2120&term=Ports and terminals'>Ports and terminals:</a><a href='../fields/2120.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>major seaport(s): </span><span class=category_data>Brisbane, Cairns, Darwin, Fremantle, Geelong, Gladstone, Hobart, Melbourne, Newcastle, Port Adelaide, Port Kembla, Sydney</span></div>
<div><span class=category>dry bulk cargo port(s): </span><span class=category_data>Dampier (iron ore), Dalrymple Bay (coal), Hay Point (coal), Port Hedland (iron ore), Port Walcott (iron ore)</span></div>
<div><span class=category>container port(s) (TEUs): </span><span class=category_data>Brisbane (1,004,983), Melbourne (2,467,967), Sydney (2,028,074)(2011)</span></div>
<div><span class=category>LNG terminal(s) (export): </span><span class=category_data>Darwin, Karratha, Burrup, Curtis Island</span></div>
</li>
<li><h2 class='question aus_med' sectiontitle='Military and Security' ccode='as' style='border-bottom: 2px solid white; cursor: pointer;'>Military and Security ::  <span class='region'>AUSTRALIA </span></h2></li><li>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2055&term=Military branches'>Military branches:</a><a href='../fields/2055.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>Australian Defense Force (ADF): Australian Army, Royal Australian Navy (includes Naval Aviation Force), Royal Australian Air Force, Joint Operations Command (JOC) (2013)</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2024&term=Military service age and obligation'>Military service age and obligation:</a><a href='../fields/2024.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>17 years of age for voluntary military service (with parental consent); no conscription; women allowed to serve in most combat roles, except the Army special forces (2013)</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2034&term=Military expenditures'>Military expenditures:</a><a href='../fields/2034.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>1.71% of GDP (2012)</div>
<div class=category_data>1.84% of GDP (2011)</div>
<div class=category_data>1.71% of GDP (2010)</div>
<div><span class='category'>country comparison to the world:  </span><span class='category_data'><a href='../rankorder/2034rank.html#as'>50</a></span></div>
</li>
<li><h2 class='question aus_med' sectiontitle='Transnational Issues' ccode='as' style='border-bottom: 2px solid white; cursor: pointer;'>Transnational Issues ::  <span class='region'>AUSTRALIA </span></h2></li><li>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2070&term=Disputes - international'>Disputes - international:</a><a href='../fields/2070.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>In 2007, Australia and Timor-Leste agreed to a 50-year development zone and revenue sharing arrangement and deferred a maritime boundary; Australia asserts land and maritime claims to Antarctica; Australia's 2004 submission to the Commission on the Limits of the Continental Shelf extends its continental margins over 3.37 million square kilometers, expanding its seabed roughly 30 percent beyond its claimed EEZ; all borders between Indonesia and Australia have been agreed upon bilaterally, but a 1997 treaty that would settle the last of their maritime and EEZ boundary has yet to be ratified by Indonesia's legislature; Indonesian groups challenge Australia's claim to Ashmore Reef; Australia closed parts of the Ashmore and Cartier reserve to Indonesian traditional fishing</div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2194&term=Refugees and internally displaced persons'>Refugees and internally displaced persons:</a><a href='../fields/2194.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div><span class=category>refugees (country of origin): </span><span class=category_data>7,785 (Afghanistan); 5,201 (Iran) (2015)</span></div>
<div id='field' class='category aus_light' style='padding-left:5px;'><a href='../docs/notesanddefs.html?fieldkey=2086&term=Illicit drugs'>Illicit drugs:</a><a href='../fields/2086.html#as'><img style='float: right;padding: 4px; border: 0; text-decoration:none;' src='../graphics/field_listing_on.gif'></a></div>
<div class=category_data>Tasmania is one of the world's major suppliers of licit opiate products; government maintains strict controls over areas of opium poppy cultivation and output of poppy straw concentrate; major consumer of cocaine and amphetamines</div>
</li>
</ul>


<!-- end generated content -->

				   <!-- End Page Content -->
                    
                   </div>  <!--  wfb-text-box -->
              
                   </div>  <!--  wfb-text-holder -->





                             <!-- END added content -->
                             </div>
                             

                            
                        </div>
            
        </div>
               </article></div>
               <aside class="content-box-2"></aside></div>
         </section><footer id="footer"><span class="divider"></span>
           <a href="#" class="logo-2"><img src="/++theme++contextual.agencytheme/images/logo-2.png" width="61" height="61" alt="Central Intelligence Agency"></a>
           <div class="footer-holder">
             <div class="footer-frame">
               <nav class="footer-nav"><div class="info-block">
                   
                   <h3><a href="/about-cia">About CIA</a></h3>
                   <ul><li>
                           <a href="/about-cia/todays-cia">Today's CIA</a>
                        </li>
                        <li>
                           <a href="/about-cia/leadership">Leadership</a>
                        </li>
                        <li>
                           <a href="/about-cia/cia-vision-mission-values">CIA Vision, Mission, Ethos &amp; Challenges</a>
                        </li>
                        <li>
                           <a href="/about-cia/headquarters-tour">Headquarters Tour</a>
                        </li>
                        <li>
                           <a href="/about-cia/cia-museum">CIA Museum</a>
                        </li>
                        <li>
                           <a href="/about-cia/history-of-the-cia">History of the CIA</a>
                        </li>
                        <li>
                           <a href="/about-cia/publications-review-board">Publications Review Board</a>
                        </li>
                        <li>
                           <a href="/about-cia/accessibility">Accessibility</a>
                        </li>
                        <li>
                           <a href="/about-cia/faqs">FAQs</a>
                        </li>
                        <li>
                           <a href="/about-cia/no-fear-act">NoFEAR Act</a>
                        </li>
                        <li>
                           <a href="/about-cia/site-policies">Site Policies</a>
                        </li>
                    </ul></div>

              <div class="info-block">
                   
                   <h3><a href="/careers">Careers &amp; Internships</a></h3>
                   <ul><li>
                           <a href="/careers/opportunities">Career Opportunities </a>
                        </li>
                        <li>
                           <a href="/careers/student-opportunities">Student Opportunities</a>
                        </li>
                        <li>
                           <a href="/careers/application-process">Application Process</a>
                        </li>
                        <li>
                           <a href="/careers/life-at-cia">Life at CIA</a>
                        </li>
                        <li>
                           <a href="/careers/benefits.html">Benefits</a>
                        </li>
                        <li>
                           <a href="/careers/diversity">Diversity</a>
                        </li>
                        <li>
                           <a href="/careers/military-transition">Military Transition</a>
                        </li>
                        <li>
                           <a href="/careers/games-information">Tools and Challenges</a>
                        </li>
                        <li>
                           <a href="/careers/faq">FAQs</a>
                        </li>
                        <li>
                           <a href="/careers/video-center">Video Center</a>
                        </li>
                    </ul><h3><a href="/offices-of-cia">Offices of CIA</a></h3>
                   <ul><li>
                           <a href="/offices-of-cia/intelligence-analysis">Intelligence &amp; Analysis</a>
                        </li>
                        <li>
                           <a href="/offices-of-cia/clandestine-service">Clandestine Service</a>
                        </li>
                        <li>
                           <a href="/offices-of-cia/science-technology">Science &amp; Technology</a>
                        </li>
                        <li>
                           <a href="/offices-of-cia/mission-support">Support to Mission</a>
                        </li>
                        <li>
                           <a href="/offices-of-cia/digital-innovation">Digital Innovation</a>
                        </li>
                        <li>
                           <a href="/offices-of-cia/mission-centers">Mission Centers</a>
                        </li>
                        <li>
                           <a href="/offices-of-cia/human-resources">Human Resources</a>
                        </li>
                        <li>
                           <a href="/offices-of-cia/public-affairs">Public Affairs</a>
                        </li>
                        <li>
                           <a href="/offices-of-cia/general-counsel">General Counsel</a>
                        </li>
                        <li>
                           <a href="/offices-of-cia/equal-employment-opportunity">Equal Employment Opportunity</a>
                        </li>
                        <li>
                           <a href="/offices-of-cia/congressional-affairs">Congressional Affairs</a>
                        </li>
                        <li>
                           <a href="/offices-of-cia/inspector-general">Inspector General</a>
                        </li>
                        <li>
                           <a href="/offices-of-cia/military-affairs">Military Affairs</a>
                        </li>
                    </ul></div>


             <div class="info-block">
                   
                   <h3><a href="/news-information">News &amp; Information</a></h3>
                   <ul><li>
                           <a href="/news-information/blog">Blog</a>
                        </li>
                        <li>
                           <a href="/news-information/press-releases-statements">Press Releases &amp; Statements</a>
                        </li>
                        <li>
                           <a href="/news-information/speeches-testimony">Speeches &amp; Testimony</a>
                        </li>
                        <li>
                           <a href="/news-information/cia-the-war-on-terrorism">CIA &amp; the War on Terrorism</a>
                        </li>
                        <li>
                           <a href="/news-information/featured-story-archive">Featured Story Archive</a>
                        </li>
                        <li>
                           <a href="/news-information/your-news">Your News</a>
                        </li>
                    </ul><h3><a href="/library">Library</a></h3>
                   <ul><li>
                           <a href="/library/publications">Publications</a>
                        </li>
                        <li>
                           <a href="/library/center-for-the-study-of-intelligence">Center for the Study of Intelligence</a>
                        </li>
                        <li>
                           <a href="/library/foia">Freedom of Information Act Electronic Reading Room</a>
                        </li>
                        <li>
                           <a href="/library/kent-center-occasional-papers">Kent Center Occasional Papers</a>
                        </li>
                        <li>
                           <a href="/library/intelligence-literature">Intelligence Literature</a>
                        </li>
                        <li>
                           <a href="/library/reports">Reports</a>
                        </li>
                        <li>
                           <a href="/library/related-links.html">Related Links</a>
                        </li>
                        <li>
                           <a href="/library/video-center">Video Center</a>
                        </li>
                    </ul></div>


               <div class="info-block add">
                   
                   <h3><a href="/kids-page">Kids' Zone</a></h3>
                   <ul><li>
                           <a href="/kids-page/k-5th-grade">K-5th Grade</a>
                        </li>
                        <li>
                           <a href="/kids-page/6-12th-grade">6-12th Grade</a>
                        </li>
                        <li>
                           <a href="/kids-page/parents-teachers">Parents &amp; Teachers</a>
                        </li>
                        <li>
                           <a href="/kids-page/games">Games</a>
                        </li>
                        <li>
                           <a href="/kids-page/related-links">Related Links</a>
                        </li>
                        <li>
                           <a href="/kids-page/privacy-statement">Privacy Statement</a>
                        </li>
                    </ul><h3><a href="/contact-cia">Connect with CIA</a></h3>
                   <ul class="socials-list"><li><a class="social-6" href="https://twitter.com/CIA">CIA Twitter</a></li>
					 <li><a class="social-5" href="https://www.facebook.com/Central.Intelligence.Agency">CIA Facebook</a></li>
					 <li><a href="http://www.youtube.com/user/ciagov">CIA YouTube</a></li>
					 <li><a class="social-2" href="http://www.flickr.com/photos/ciagov">CIA Flickr PhotoStream</a></li>
                     <li><a class="social-3" href="/news-information/your-news">RSS</a></li>
                     <li><a class="social-4 landingTrigger" href="/contact-cia">Contact Us</a></li>
                   </ul></div>


               </nav><div id="plugins" class="info-panel">
                    <h4>* Required plugins</h4>
                    <ul><li data-plugin="swf"><a href="http://get.adobe.com/flashplayer/">Adobe&#174; Flash Player</a></li>
                        <li data-plugin="pdf"><a href="http://get.adobe.com/reader/">Adobe&#174; Reader&#174;</a></li>
                        <li data-plugin="doc"><a href="http://www.microsoft.com/en-us/download/details.aspx?id=4">MS Word Viewer</a></li>
                    </ul></div>
             </div>
           </div>
         </footer></div>
    <div class="footer-panel">
      <nav class="sub-nav"><h3 class="visuallyhidden">Footer Navigation</h3>
        <ul><li><a href="/about-cia/site-policies/#privacy-notice" title="Site Policies">Privacy</a></li>
          <li><a href="/about-cia/site-policies/#copy" title="Site Policies">Copyright</a></li>
          <li><a href="/about-cia/site-policies/" title="Site Policies">Site Policies</a></li>
          <li><a href="http://www.usa.gov/">USA.gov</a></li>
          <li><a href="http://www.foia.cia.gov/">FOIA</a></li>
          <li><a href="http://www.dni.gov/">DNI.gov</a></li>
          <li><a href="/about-cia/no-fear-act/" title="No FEAR Act">NoFEAR Act</a></li>
          <li><a href="/offices-of-cia/inspector-general/">Inspector General</a></li>
          <!--<li><a tal:attributes="href string:$portal_url/mobile/" href="#" >Mobile Site</a></li>-->
          <li><a class="landingTrigger" href="/contact-cia/">Contact CIA</a></li>
          <li><a href="/sitemap.html">Site Map</a></li>
        </ul><a href="/open/" class="footer-logo"><img src="/++theme++contextual.agencytheme/images/ico-06.png" alt="open gov"></a>
      </nav></div>
    <div class="skip"><a href="#wrapper">back to top</a></div>
  </div></div><script async src="/js2/verification.js"></script></body></html>

`

var aaCountryList20140902 = []string{
	"xx.html",
	"af.html",
	"ax.html",
	"al.html",
	"ag.html",
	"aq.html",
	"an.html",
	"ao.html",
	"av.html",
	"ay.html",
	"ac.html",
	"xq.html",
	"ar.html",
	"am.html",
	"aa.html",
	"at.html",
	"zh.html",
	"as.html",
	"au.html",
	"aj.html",
	"bf.html",
	"ba.html",
	"bg.html",
	"bb.html",
	"bo.html",
	"be.html",
	"bh.html",
	"bn.html",
	"bd.html",
	"bt.html",
	"bl.html",
	"bk.html",
	"bc.html",
	"bv.html",
	"br.html",
	"io.html",
	"vi.html",
	"bx.html",
	"bu.html",
	"uv.html",
	"bm.html",
	"by.html",
	"cv.html",
	"cb.html",
	"cm.html",
	"ca.html",
	"cj.html",
	"ct.html",
	"cd.html",
	"ci.html",
	"ch.html",
	"kt.html",
	"ip.html",
	"ck.html",
	"co.html",
	"cn.html",
	"cg.html",
	"cf.html",
	"cw.html",
	"cr.html",
	"cs.html",
	"iv.html",
	"hr.html",
	"cu.html",
	"cc.html",
	"cy.html",
	"ez.html",
	"da.html",
	"dx.html",
	"dj.html",
	"do.html",
	"dr.html",
	"ec.html",
	"eg.html",
	"es.html",
	"ek.html",
	"er.html",
	"en.html",
	"et.html",
	"fk.html",
	"fo.html",
	"fj.html",
	"fi.html",
	"fr.html",
	"fp.html",
	"gb.html",
	"ga.html",
	"gz.html",
	"gg.html",
	"gm.html",
	"gh.html",
	"gi.html",
	"gr.html",
	"gl.html",
	"gj.html",
	"gq.html",
	"gt.html",
	"gk.html",
	"gv.html",
	"pu.html",
	"gy.html",
	"ha.html",
	"hm.html",
	"vt.html",
	"ho.html",
	"hk.html",
	"hu.html",
	"ic.html",
	"in.html",
	"xo.html",
	"id.html",
	"ir.html",
	"iz.html",
	"ei.html",
	"im.html",
	"is.html",
	"it.html",
	"jm.html",
	"jn.html",
	"ja.html",
	"je.html",
	"jo.html",
	"kz.html",
	"ke.html",
	"kr.html",
	"kn.html",
	"ks.html",
	"kv.html",
	"ku.html",
	"kg.html",
	"la.html",
	"lg.html",
	"le.html",
	"lt.html",
	"li.html",
	"ly.html",
	"ls.html",
	"lh.html",
	"lu.html",
	"mc.html",
	"mk.html",
	"ma.html",
	"mi.html",
	"my.html",
	"mv.html",
	"ml.html",
	"mt.html",
	"rm.html",
	"mr.html",
	"mp.html",
	"mx.html",
	"fm.html",
	"md.html",
	"mn.html",
	"mg.html",
	"mj.html",
	"mh.html",
	"mo.html",
	"mz.html",
	"wa.html",
	"nr.html",
	"bq.html",
	"np.html",
	"nl.html",
	"nc.html",
	"nz.html",
	"nu.html",
	"ng.html",
	"ni.html",
	"ne.html",
	"nf.html",
	"cq.html",
	"no.html",
	"mu.html",
	"zn.html",
	"pk.html",
	"ps.html",
	"pm.html",
	"pp.html",
	"pf.html",
	"pa.html",
	"pe.html",
	"rp.html",
	"pc.html",
	"pl.html",
	"po.html",
	"rq.html",
	"qa.html",
	"ro.html",
	"rs.html",
	"rw.html",
	"tb.html",
	"sh.html",
	"sc.html",
	"st.html",
	"rn.html",
	"sb.html",
	"vc.html",
	"ws.html",
	"sm.html",
	"tp.html",
	"sa.html",
	"sg.html",
	"ri.html",
	"se.html",
	"sl.html",
	"sn.html",
	"sk.html",
	"lo.html",
	"si.html",
	"bp.html",
	"so.html",
	"sf.html",
	"oo.html",
	"sx.html",
	"od.html",
	"sp.html",
	"pg.html",
	"ce.html",
	"su.html",
	"ns.html",
	"sv.html",
	"wz.html",
	"sw.html",
	"sz.html",
	"sy.html",
	"tw.html",
	"ti.html",
	"tz.html",
	"th.html",
	"tt.html",
	"to.html",
	"tl.html",
	"tn.html",
	"td.html",
	"ts.html",
	"tu.html",
	"tx.html",
	"tk.html",
	"tv.html",
	"ug.html",
	"up.html",
	"ae.html",
	"uk.html",
	"us.html",
	"uy.html",
	"uz.html",
	"nh.html",
	"ve.html",
	"vm.html",
	"vq.html",
	"wq.html",
	"wf.html",
	"we.html",
	"wi.html",
	"ym.html",
	"za.html",
	"zi.html",
	"ee.html",
}
