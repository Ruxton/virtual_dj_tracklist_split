Mixcloud CLI Uploader
====================================

Overview
--------

A go CLI application that splits Virtual DJ's tracklists.txt into multiple files by session/date

Binaries
---------

Compiled [binaries are available on dropbox](https://www.dropbox.com/sh/bqrajt2q74vn3jx/AABBFzI4327haGgjpHQrwKHHa) for:
  * [Windows 64bit](https://www.dropbox.com/s/4hb25ooz2p3h38d/mixcloud.exe)
  * [Linux 32bit](https://www.dropbox.com/s/p2ny4njqfm966z0/mixcloud.linux)
  * [Linux 64bit](https://www.dropbox.com/s/g6p5fg7bnn9x5o2/mixcloud.linux64)
  * [OSX 64bit](https://www.dropbox.com/s/27j48w0bete7xkt/mixcloud.osx)

Source Requirements
------------

* GoLang > 1.2.1

Using Mixcloud CLI Uploader From Source
--------------------------------

  1. Run bin/build
  1. Run the built packages as below from pkg/

Using Pre-compiled Packages
---------------------------

Using compiled packages:

  `vdj_tracklist_split tracklist.txt`

Meta
----

* Code: `git clone git://github.com/ruxton/virtual_dj_tracklist_split.git`
