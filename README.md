# Chrome extension to play music in tikplay

* Put com.tik.tikhost.json  to ~/.config/chromium/NativeMessagingHosts (Or ~/.config/chrome/NativeMessagingHosts if you use chrome)

* Put sendToTikplay in /usr/bin and make it excutable and accessable to the user running chrome
* load chrome-extension (this folder) from chrom(e/ium) ( Extensions, developer mode, load unpacked extension ) 
* Change com.tik.tikhos.json allowed-origins id to match your id from chrome extensions:

"allowed_origins": [
    "chrome-extension://<id from extensions>/"
  ]

* restart chrom(e/ium)
 
## Requirements:
* Python
* SSH with kekkonen configured as a alias and with passwordless key
* tikplay configured as alias in kekkonen
