# ghostgame
A simple RNG-based game, originally written in Python, re-written in Go.
# Building
1. Install the latest version of [Go](https://golang.org/dl)
2. Clone the source repository if you haven't already:
```
git clone https://github.com/TheSonicMaster/ghostgame.git
```
3. Change to the source directory if you aren't already in it:
```
cd ghostgame
```
4. Build it.
```
make
```
5. (Optional) strip unnecessary debug symbols from the executable:
```
make strip
```
6. Install it.
```
sudo make install
```
**Note: By default, ghostgame gets installed to the `/usr/local` prefix. This
can be overridden using the following commands:**
```
sudo make PREFIX=/somewhere/on/your/system install
```
**If you are going to distribute it as a package for a package manager, use
`DESTDIR` instead:**
```
sudo make DESTDIR=/path/to/package/directory install
```
