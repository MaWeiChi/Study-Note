

```bash
$ sudo nano /usr/share/applications/google-chrome.desktop

#Exec=/usr/bin/google-chrome-stable %U
Exec=/usr/bin/google-chrome-stable --force-device-scale-factor=1.5 %U

StartupNotify=true
Terminal=false
Icon=google-chrome
Type=Application
Categories=Network;WebBrowser;
MimeType=application/pdf;application/rdf+xml;application/rss+xml;application/xhtml+xml;application/xhtml_xml;applic>
Actions=new-window;new-private-window;

```

