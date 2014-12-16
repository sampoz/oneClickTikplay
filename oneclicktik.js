
function sendCurrentUrlToTikPlay() {
    chrome.tabs.query({active:true, lastFocusedWindow:true}, function(tab) {
        chrome.runtime.sendNativeMessage(
            "com.tik.tikhost",
            {youtubeUrl: tab[0].url}
        );
        document.querySelector("#MessageDiv").textContent = "URL sent!"
    });
}

document.addEventListener('DOMContentLoaded', function () {
    document.querySelector("#TikPlayButton").addEventListener('click', sendCurrentUrlToTikPlay);
});