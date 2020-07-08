Guillaume Martres 2016-10-04 20:12:41 UTC

```
Alright, I got it, let me know if there's a less terrible way of doing this:
- Edit /usr/share/plasma/plasmoids/org.kde.plasma.private.systemtray/contents/config/main.xml
- Under "<entry name="iconSize" type="Int">" you'll see:
  <default>1</default>
- Replace 1 by 2 or something bigger, for me 2 was exactly right.
```