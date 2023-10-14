package encoding

import (
	"encoding/base64"
	"github.com/axgle/mahonia"
	exurl "github.com/qiniu/api.v6/url"
	"github.com/robertkrimen/otto"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"math/rand"
	"time"

	"net/url"
	"strconv"
	"strings"
)

// URLEscape URL编码
func URLEscape(path string, escapeSlash bool) string {
	result := url.PathEscape(path)
	if !escapeSlash {
		result = strings.Replace(result, "%2F", "/", -1)
	}
	return result
}

// URLUnescape URL解码
func URLUnescape(path string) string {
	result, err := url.PathUnescape(path)
	if err != nil {
		return path
	}
	return result
}

func Escape(value string) (string, error) {
	vm := otto.New()
	_, err := vm.Run(jsEscape)
	if err != nil {
		return "", err
	}
	res, err := vm.Call("es", nil, value)
	if err != nil {
		return "", err
	}
	return res.String(), err
}

var jsEscape = `
function es(s) {
	return escape(s);
}
`

func Unescape(value string) (string, error) {
	vm := otto.New()
	_, err := vm.Run(jsUnescape)
	if err != nil {
		return "", err
	}
	res, err := vm.Call("unes", nil, value)
	if err != nil {
		return "", err
	}
	return res.String(), err
}

var jsUnescape = `
function unes(s) {
	return unescape(s);
}
`

func URICompCoding(value string) string {
	return exurl.EscapeEx(value, exurl.EncodeQueryComponent)
}

func URICompDecoding(value string) (string, error) {
	return exurl.UnescapeEx(value, exurl.EncodeQueryComponent)
}

func EncodeURI(value string) string {
	return exurl.EscapeEx(value, exurl.EncodeFragment)
}

func DecodeURI(value string) (string, error) {
	return exurl.UnescapeEx(value, exurl.EncodeFragment)
}

func Utf8ToGBK(utf8str string) (string, error) {
	result, _, err := transform.String(simplifiedchinese.GBK.NewEncoder(), utf8str)
	bres := []byte(result)
	var res string
	for i := range bres {
		t, _ := DecHex(int64(bres[i]))
		res += "%" + t
	}
	return res, err
}

func GBKToUtf8(gbkStr string) (string, error) {
	bs := strings.Split(gbkStr, "%")
	bs2 := make([]byte, 0, len(bs))
	for i := range bs {
		if bs[i] != "" {
			num, err := HexDec(bs[i])
			if err != nil {
				return "", err
			}
			bs2 = append(bs2, byte(num))
		}
	}
	dec := mahonia.NewDecoder("GBK")
	res := dec.ConvertString(string(bs2))
	return res, nil
}

// HexStr 十六进制字符串
func HexStr(str string) (string, string, string) {
	enc := mahonia.NewEncoder("GBK")
	res := enc.ConvertString(str)
	bs := []byte(res)
	var str1 = "0x"
	var str2 = "0x"
	var str3 string
	for i := range bs {
		t, _ := DecHex(int64(bs[i]))
		str1 += t
		str2 += "00" + t
		str3 += "\\x" + t
	}
	return str1, str2, str3
}

func StrHex(str string) (string, error) {
	str = strings.TrimLeft(str, "0x")
	strs := strings.Split(str, "\\x")
	var hexs []string
	if len(strs) == 1 {
		rs := []rune(strs[0])
		hexs = make([]string, len(rs)/2)
		for i := range hexs {
			temp := string(rs[i*2 : (i+1)*2])
			hexs[i] = temp
		}
	} else if len(strs) > 1 {
		hexs = make([]string, 0, len(strs))
		for i := range strs {
			if strs[i] != "" {
				hexs = append(hexs, strs[i])
			}
		}
	}
	hexs2 := make([]byte, 0, len(hexs))
	for i := range hexs {
		if hexs[i] != "00" {
			if num, err := HexDec(hexs[i]); err != nil {
				return "", err
			} else {
				hexs2 = append(hexs2, byte(num))
			}
		}
	}
	dec := mahonia.NewDecoder("GBK")
	res := dec.ConvertString(string(hexs2))
	return res, nil
}

func EncodeBase64(value string) string {
	return base64.StdEncoding.EncodeToString([]byte(value))
}

func DecodeBase64(value string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(value)
	return string(data), err
}

// Convert 编码转换将src编码转换为dst编码
//https://www.gnu.org/software/libiconv/
//func Convert(dst, src string, reader io.Reader) ([]byte, error) {
//	if dst == "" {
//		dst = "utf-8"
//	}
//	if src == "" {
//		src = "utf-8"
//	}
//	if dst == src {
//		return ioutil.ReadAll(reader)
//	}
//	cd, err := iconv.Open(dst, src) // convert src to dst
//	if err != nil {
//		return nil, err
//	}
//	defer cd.Close()
//	r := iconv.NewReader(cd, reader, 0)
//	var buffer = make([]byte, 0)
//	for {
//		buf := make([]byte, 1024)
//		n, err := r.Read(buf)
//		if err != nil && err != io.EOF {
//			io.Copy(ioutil.Discard, r) //释放掉未获取的数据
//			return nil, err
//		}
//		if n == 0 {
//			break
//		}
//		buffer = append(buffer, buf[:n]...)
//		if len(buffer) > MaxResponseBytes {
//			io.Copy(ioutil.Discard, r) //释放掉未获取的数据
//			break
//		}
//	}
//	return buffer, nil
//}

type charset string

const (
	UTF8    = charset("UTF-8")
	GB18030 = charset("GB18030")
)

// ConvertByte 将b charset转换为UTF-8编码
func ConvertByte(b []byte, charset charset) []byte {
	var byt = make([]byte, 0)
	switch charset {
	case GB18030:
		byt, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(b)
	case UTF8:
		fallthrough
	default:
		byt = b
	}
	return byt
}

// EncodePacked 将内容编码为
func EncodePacked(value string) (string, error) {
	vm := otto.New()
	_, err := vm.Run(jsPackedEncode)
	if err != nil {
		return "", err
	}
	res, err := vm.Call("pack", nil, value, 62, true, false)
	if err != nil {
		return "", err
	}
	return res.String(), err
}

func DecodePacked(value string) (string, error) {
	vm := otto.New()
	_, err := vm.Run(jsPackedDecode)
	if err != nil {
		return "", err
	}
	res, err := vm.Call("decodeScript", nil, value)
	if err != nil {
		return "", err
	}
	return res.String(), err
}

var jsPackedEncode = `
function pack(_7,_0,_2,_8){var I="$1";_7+="\n";_0=Math.min(parseInt(_0),95);function _15(s){var i,p;for(i=0;(p=_6[i]);i++){s=p(s)}return s};var _25=function(p,a,c,k,e,d){while(c--)if(k[c])p=p.replace(new RegExp('\\b'+e(c)+'\\b','g'),k[c]);return p};var _26=function(){if(!''.replace(/^/,String)){while(c--)d[e(c)]=k[c]||e(c);k=[function(e){return d[e]}];e=function(){return'\\w+'};c=1}};var _6=[];function _4(p){_6[_6.length]=p};function _18(s){var p=new ParseMaster;p.escapeChar="\\";p.add(/'[^'\n\r]*'/,I);p.add(/"[^"\n\r]*"/,I);p.add(/\/\/[^\n\r]*[\n\r]/," ");p.add(/\/\*[^*]*\*+([^\/][^*]*\*+)*\//," ");p.add(/\s+(\/[^\/\n\r\*][^\/\n\r]*\/g?i?)/,"$2");p.add(/[^\w\x24\/'"*)\?:]\/[^\/\n\r\*][^\/\n\r]*\/g?i?/,I);if(_8)p.add(/;;;[^\n\r]+[\n\r]/);p.add(/\(;;\)/,I);p.add(/;+\s*([};])/,"$2");s=p.exec(s);p.add(/(\b|\x24)\s+(\b|\x24)/,"$2 $3");p.add(/([+\-])\s+([+\-])/,"$2 $3");p.add(/\s+/,"");return p.exec(s)};function _17(s){var p=new ParseMaster;p.add(/((\x24+)([a-zA-Z_]+))(\d*)/,function(m,o){var l=m[o+2].length;var s=l-Math.max(l-m[o+3].length,0);return m[o+1].substr(s,l)+m[o+4]});var r=/\b_[A-Za-z\d]\w*/;var k=_13(s,_9(r),_21);var e=k.e;p.add(r,function(m,o){return e[m[o]]});return p.exec(s)};function _16(s){if(_0>62)s=_20(s);var p=new ParseMaster;var e=_12(_0);var r=(_0>62)?/\w\w+/ :/\w+/;k=_13(s,_9(r),e);var e=k.e;p.add(r,function(m,o){return e[m[o]]});return s&&_27(p.exec(s),k)};function _13(s,r,e){var a=s.match(r);var so=[];var en={};var pr={};if(a){var u=[];var p={};var v={};var c={};var i=a.length,j=0,w;do{w="$"+a[--i];if(!c[w]){c[w]=0;u[j]=w;p["$"+(v[j]=e(j))]=j++}c[w]++}while(i);i=u.length;do{w=u[--i];if(p[w]!=null){so[p[w]]=w.slice(1);pr[p[w]]=true;c[w]=0}}while(i);u.sort(function(m1,m2){return c[m2]-c[m1]});j=0;do{if(so[i]==null)so[i]=u[j++].slice(1);en[so[i]]=v[i]}while(++i<u.length)}return{s:so,e:en,p:pr}};function _27(p,k){var E=_10("e\\(c\\)","g");p="'"+_5(p)+"'";var a=Math.min(k.s.length,_0)||1;var c=k.s.length;for(var i in k.p)k.s[i]="";k="'"+k.s.join("|")+"'.split('|')";var e=_0>62?_11:_12(a);e=String(e).replace(/_0/g,"a").replace(/arguments\.callee/g,"e");var i="c"+(a>10?".toString(a)":"");if(_2){var d=_19(_26);if(_0>62)d=d.replace(/\\\\w/g,"[\\xa1-\\xff]");else if(a<36)d=d.replace(E,i);if(!c)d=d.replace(_10("(c)\\s*=\\s*1"),"$1=0")}var u=String(_25);if(_2){u=u.replace(/\{/,"{"+d+";")}u=u.replace(/"/g,"'");if(_0>62){u=u.replace(/'\\\\b'\s*\+|\+\s*'\\\\b'/g,"")}if(a>36||_0>62||_2){u=u.replace(/\{/,"{e="+e+";")}else{u=u.replace(E,i)}u=pack(u,0,false,true);var p=[p,a,c,k];if(_2){p=p.concat(0,"{}")}return"eval("+u+"("+p+"))\n"};function _12(a){return a>10?a>36?a>62?_11:_22:_23:_24};var _24=function(c){return c};var _23=function(c){return c.toString(36)};var _22=function(c){return(c<_0?'':arguments.callee(parseInt(c/_0)))+((c=c%_0)>35?String.fromCharCode(c+29):c.toString(36))};var _11=function(c){return(c<_0?'':arguments.callee(c/_0))+String.fromCharCode(c%_0+161)};var _21=function(c){return"_"+c};function _5(s){return s.replace(/([\\'])/g,"\\$1")};function _20(s){return s.replace(/[\xa1-\xff]/g,function(m){return"\\x"+m.charCodeAt(0).toString(16)})};function _10(s,f){return new RegExp(s.replace(/\$/g,"\\$"),f)};function _19(f){with(String(f))return slice(indexOf("{")+1,lastIndexOf("}"))};function _9(r){return new RegExp(String(r).slice(1,-1),"g")};_4(_18);if(_8)_4(_17);if(_0)_4(_16);return _15(_7)};
function ParseMaster(){var E=0,R=1,L=2;var G=/\(/g,S=/\$\d/,I=/^\$\d+$/,T=/(['"])[1]\+(.*)\+[1][1]$/,ES=/\\./g,Q=/'/,DE=/\x01[^\x01]*\x01/g;var self=this;this.add=function(e,r){if(!r)r="";var l=(_14(String(e)).match(G)||"").length+1;if(S.test(r)){if(I.test(r)){r=parseInt(r.slice(1))-1}else{var i=l;var q=Q.test(_14(r))?'"':"'";while(i)r=r.split("$"+i--).join(q+"+a[o+"+i+"]+"+q);r=new Function("a,o","return"+q+r.replace(T,"$1")+q)}}_33(e||"/^$/",r,l)};this.exec=function(s){_3.length=0;return _30(_5(s,this.escapeChar).replace(new RegExp(_1,this.ignoreCase?"gi":"g"),_31),this.escapeChar).replace(DE,"")};this.reset=function(){_1.length=0};var _3=[];var _1=[];var _32=function(){return"("+String(this[E]).slice(1,-1)+")"};_1.toString=function(){return this.join("|")};function _33(){arguments.toString=_32;_1[_1.length]=arguments}function _31(){if(!arguments[0])return"";var i=1,j=0,p;while(p=_1[j++]){if(arguments[i]){var r=p[R];switch(typeof r){case"function":return r(arguments,i);case"number":return arguments[r+i]}var d=(arguments[i].indexOf(self.escapeChar)==-1)?"":"\x01"+arguments[i]+"\x01";return d+r}else i+=p[L]}};function _5(s,e){return e?s.replace(new RegExp("\\"+e+"(.)","g"),function(m,c){_3[_3.length]=c;return e}):s};function _30(s,e){var i=0;return e?s.replace(new RegExp("\\"+e,"g"),function(){return e+(_3[i++]||"")}):s};function _14(s){return s.replace(ES,"")}};ParseMaster.prototype={constructor:ParseMaster,ignoreCase:false,escapeChar:""};
`

var jsPackedDecode = `
function decodeScript(jsstr){
	try{
		return eval(jsstr.slice(4));
	}catch(e){
		
	}
};
`

// HtmlEscape HTML转义
func HtmlEscape(htmlStr string) string {
	rs := []rune(htmlStr)
	var html_ string
	for _, r := range rs {
		html_ += "&#" + strconv.Itoa(int(r)) + ";" // 网页
	}
	return html_
}

// HtmlUnescape HTML反转义
func HtmlUnescape(str string) (string, error) {
	strs := strings.Split(str, ";")
	rs := make([]rune, 0, len(strs))
	for i := range strs {
		if strs[i] != "" {
			num, err := strconv.Atoi(strings.TrimLeft(strs[i], "&#"))
			if err != nil {
				return "", err
			}
			rs = append(rs, rune(num))
		}
	}
	return string(rs), nil
}

// GenRandomString 随机生成字符串
func GenRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var result []byte
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
