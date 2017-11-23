function getUrlParam(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)"); 
    var r = window.location.search.substr(1).match(reg); 
    if (r != null) return unescape(r[2]); return ""; 
}

function change(url){
	$("#iframeContent").attr("src",url)
}

$(document).ready(function(){
	var req_url = getUrlParam("requrl")
	if (req_url == "" ) {
		req_url = "/";
	}
});
