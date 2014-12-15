chrome.tabs.getSelected(null, function(tab) {
    myFunction(tab.url);
});

function myFunction(tablink) {
   var testobj = {};
   testobj.youtubeURL = tablink;
   chrome.runtime.sendNativeMessage("com.tik.tikhost",testobj);
}
