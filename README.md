# Chrome extension to play music in tikplay

* Put com.tik.tikhost.json  to ~/.config/chromium/NativeMessagingHosts (Or ~/.config/chrome/NativeMessagingHosts if you use chrome)
* restart chrom(e/ium)

* Put sendToTikplay in /usr/bin and make it excutable and accessable to the user running chrome
* load chrome-extension (this folder) from chrom(e/ium) ( Extensions, developer mode, load unpacked extension ) 

## Requirements:
* Python
* SSH with kekkonen configured as a alias and with passwordless key
* tikplay configured as alias in kekkonen
