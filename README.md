# golibrepo

problem - I have a RaspberryPi Cluster that I run python apps on,
and the data-intensive & distrubuted stuff is becoming slow and 
annoying to code via Python

my solution - Instead of re-writing entire servers in Go, start
breaking out libs into this monorepo and upload them to my cluster
as a c-shared file, which can be dynamically loaded into py3 code

todo
- include inner-package dependencies
- clean up CI (ex: don't try buildchanges when there's none)
